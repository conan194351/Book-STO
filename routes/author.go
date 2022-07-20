package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/repository"
	"book-sto/service"

	"github.com/gin-gonic/gin"
)

func AuthorRoute(router *gin.Engine) {

	handler := handlers.NewAuthorHandler(service.NewAuthorServices(repository.NewAuthorRepository(config.DB)))
	route := router.Group("/api/author")
	{

		route.GET("/", handler.GetListAuthor())
		route.POST("/create", handler.CreateAuthor())
		route.POST("/search", handler.SearchAuthor())
	}
}
