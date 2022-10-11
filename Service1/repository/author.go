package repository

import (
	"fmt"
	"service1/config"
	"service1/dto"
	"service1/errs"
	"service1/proto"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthService1Repo interface {
	LoginGRPC(request *proto.LoginRequest) (*proto.LoginResponse, error)
	FindBookByIdAuthor(request *proto.FindBookByIdAuthorRequest) ([]*proto.FindBookByIdAuthorResponse, error)
	FindAuthorByUsername(username string) (string, *errs.AppError)
	Logout(request *proto.LogoutRequest) (*proto.LogoutResponse, error)
}

type DefaultAuthService1Repo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthService1Repo(db *gorm.DB, redis *redis.Client) AuthService1Repo {
	return &DefaultAuthService1Repo{
		db:    db,
		redis: redis,
	}
}

func (r DefaultAuthService1Repo) FindAuthorByUsername(username string) (string, *errs.AppError) {
	var res dto.LoginAuthorResponse
	if err := r.db.Table("author").Select("author.Username").Where("author.Username = ?", username).Find(&res).Error; err != nil {
		return "false", errs.ErrorGetData()
	}
	if res.Username == "" {
		return "false", nil
	} else {
		return res.Username, nil
	}
}

func (r DefaultAuthService1Repo) FindBookByIdAuthor(request *proto.FindBookByIdAuthorRequest) ([]*proto.FindBookByIdAuthorResponse, error) {
	response := []*proto.FindBookByIdAuthorResponse{}
	var res []dto.Book
	fmt.Print(request)
	if err := r.db.Table("book").Select("book.IdBook,book.Name").Where("book_author.idAuthor = ?", request.IdAuthor).Joins("JOIN book_author on book.idBook = book_author.idBook").Find(&res).Error; err != nil {
		return response, err
	}
	for _, book := range res {
		bookProto := &proto.FindBookByIdAuthorResponse{}
		bookProto.IdBook = book.IdBook
		bookProto.NameBook = book.Name
		response = append(response, bookProto)
	}
	return response, nil
}

func (r DefaultAuthService1Repo) LoginGRPC(request *proto.LoginRequest) (*proto.LoginResponse, error) {
	response := &proto.LoginResponse{}
	var res dto.LoginAuthorResponse
	if err := r.db.Table("author").Select("author.Username").Where("author.Username = ? and author.Password = ?", request.Username, request.Password).Find(&res).Error; err != nil {
		return response, err
	}
	response.Status = "False"
	if res.Username != "" {
		jwtToken, token, e := config.NewJWTToken(res.Username)
		if e != nil {
			return response, e
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		expiredAt := claims["exp"].(int64)
		res := &proto.LoginResponse{
			Status:   "Success",
			Username: res.Username,
			Token:    *token,
			ExpireAt: expiredAt,
		}
		return res, nil
	}
	return response, nil
}

func (r DefaultAuthService1Repo) Logout(request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	err := r.redis.Set(request.Token, "logout", 0).Err()
	if err != nil {
		return &proto.LogoutResponse{Status: "false"}, err
	}
	return &proto.LogoutResponse{Status: "true"}, nil
}
