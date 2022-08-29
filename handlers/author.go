package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/proto"
	"book-sto/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AuthorHandler struct {
	services service.AuthorServices
}

func NewAuthorHandler(services service.AuthorServices) AuthorHandler {

	return AuthorHandler{

		services: services,
	}
}

func (a AuthorHandler) LoginAuthor() gin.HandlerFunc {
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
			WriteRespon(ctx, http.StatusOK, dto.LoginSuccess("author"))
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:    "token",
				Value:   response.Token,
				Expires: time.Now().Add(30 * 24 * time.Hour),
			})
			fmt.Println(response.Token)

		} else {
			WriteRespon(ctx, http.StatusOK, dto.LoginFalse())
		}
	}
}

func (a AuthorHandler) GetListAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		res, err := a.services.ListAuthor()
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (a AuthorHandler) CreateAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var author dto.CreateAutherRequest
		err := ctx.BindJSON(&author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		_, e := a.services.CreateAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageCreateSuccess("Author"))
	}
}

func (a AuthorHandler) SearchAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var author dto.SearchAuthorRequest
		err := ctx.BindJSON(&author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (a AuthorHandler) ShowBookByAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(string)
		res, e := a.services.ShowBookByAuthor(user)
		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}

func FindBookByIdAuthor(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client := proto.NewAddServiceClient(conn)
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
