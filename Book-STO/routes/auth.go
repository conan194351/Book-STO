package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/repository"
	"book-sto/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	handler := handlers.NewAuthHandler(service.NewAuthServices(repository.NewAuthRepository(config.DB)))
	route := router.Group("/api/auth")
	{
		route.POST("/login", handler.LoginAuthor())
	}
}
