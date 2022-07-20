package app

import (
	"book-sto/config"
	"book-sto/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func RunApp() {
	config.InitDatabase()
	router := gin.Default()
	routes.BookRoute(router)
	routes.AuthorRoute(router)
	routes.CategoryRoute(router)

	log.Println("Server is running on PORT ", config.ConnectPort())
	router.Run(config.ConnectPort())
}
