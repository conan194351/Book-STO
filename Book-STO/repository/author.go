package repository

import (
	"book-sto/config"
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/model"
	"database/sql"

	"github.com/golang-jwt/jwt"
)

type AuthorRepository interface {
	List() ([]model.Author, *errs.AppError)
	Create(model.Author) (model.Author, *errs.AppError)
	SearchAuthor(req string) ([]model.Author, *errs.AppError)
	Login(dto.LoginAuthorRequest) (dto.LoginAuthorResponse, *errs.AppError)
	FindAuthorByUsername(username string) (string, *errs.AppError)
	ShowBookByAuthor(req string) ([]model.Book, *errs.AppError)
}

type DefaultAuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {

	return DefaultAuthorRepository{

		db: db,
	}
}

func (a DefaultAuthorRepository) Login(req dto.LoginAuthorRequest) (dto.LoginAuthorResponse, *errs.AppError) {
	author := model.Author{
		Username: req.Username,
		Password: req.Password,
	}
	var username string
	var response dto.LoginAuthorResponse
	res, err := a.db.Query("select author.username from longphu.author as author where author.username = ? and author.password = ?", author.Username, author.Password)
	for res.Next() {
		err = res.Scan(&username)
		if err != nil {
			response.Status = "False"
			return response, errs.ErrorReadData()
		}
	}
	response.Status = "False"
	if username != "" {
		jwtToken, token, e := config.NewJWTToken(username)
		if e != nil {
			return response, e
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		expiredAt := claims["exp"].(int64)
		res := dto.LoginAuthorResponse{
			Status:   "Success",
			Username: username,
			Token:    *token,
			ExpireAt: expiredAt,
		}
		return res, nil
	}
	return response, nil
}

func (a DefaultAuthorRepository) FindAuthorByUsername(username string) (string, *errs.AppError) {
	res, err := a.db.Query("select author.username from longphu.author as author where author.username = ? ", username)
	if err != nil {

		return "false", errs.ErrorGetData()
	}
	var user string
	for res.Next() {
		err = res.Scan(&user)
		if err != nil {

			return "false", errs.ErrorReadData()
		}
	}
	if user == "" {
		return "false", nil
	} else {
		return user, nil
	}
}

func (a DefaultAuthorRepository) List() ([]model.Author, *errs.AppError) {

	res, err := a.db.Query("SELECT author.idauthor, author.Name, author.NativeLand FROM longphu.author as author")

	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var authors []model.Author
	var author model.Author
	for res.Next() {
		err = res.Scan(&author.IdAuthor, &author.Name, &author.NativeLand)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a DefaultAuthorRepository) Create(author model.Author) (model.Author, *errs.AppError) {

	newAuthor, err := a.db.Prepare("INSERT INTO longphu.author(Name, NativeLand) VALUES (? ,?) ")
	if err != nil {

		return author, errs.ErrorInsertData()
	}
	newAuthor.Exec(author.Name, author.NativeLand)
	return author, nil
}

func (a DefaultAuthorRepository) SearchAuthor(req string) ([]model.Author, *errs.AppError) {
	var authors []model.Author
	if req == "" {
		return authors, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT author.Name, author.NativeLand FROM longphu.author as author WHERE author.Name LIKE ?", bodyString)

	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {

		var author model.Author
		err = res.Scan(&author.Name, &author.NativeLand)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a DefaultAuthorRepository) ShowBookByAuthor(req string) ([]model.Book, *errs.AppError) {
	var books = []model.Book{}
	res, err := a.db.Query("select book.idBook, book.Name from longphu.book as book , longphu.author as author, longphu.book_author as b_a where book.idBook = b_a.idBook and b_a.idAuthor = author.idAuthor and author.username = ?", req)
	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {
		var book model.Book
		err = res.Scan(&book.Name, &book.Category)
		if err != nil {
			return nil, errs.ErrorReadData()
		}
		books = append(books, book)
	}
	return books, nil
}
