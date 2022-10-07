package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/proto"
	"book-sto/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AuthHandler struct {
	services service.AuthServices
}

func NewAuthHandler(services service.AuthServices) AuthHandler {

	return AuthHandler{
		services: services,
	}
}

func (a AuthHandler) LoginAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var author dto.LoginAuthorRequest
		err := ctx.BindJSON(&author)

		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		response, e := a.services.LoginAuthor(author)
		if err != nil {

			WriteError(ctx, e)
			return
		}
		if response.Status == "Success" {
			WriteRespon(ctx, http.StatusOK, dto.LoginSuccess("author", response.Token))
		} else {
			WriteRespon(ctx, http.StatusOK, dto.LoginFalse())
		}
	}
}

func LoginGRPC(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client := proto.NewAddAuthorServiceClient(conn)
		var author dto.LoginAuthorRequest
		err := ctx.BindJSON(&author)

		if err != nil {
			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter Username Or Password"})
			return
		}

		req := &proto.LoginRequest{Username: author.Username, Password: author.Password}
		if response, err := client.LoginGPRC(ctx, req); err == nil {
			WriteRespon(ctx, http.StatusOK, dto.LoginSuccess("author", response.Token))
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
