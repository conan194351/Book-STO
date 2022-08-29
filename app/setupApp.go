package app

import (
	"book-sto/config"
	"book-sto/routes"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RunApp() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	config.InitDatabase()
	router := gin.Default()
	routes.Service(router, conn)
	routes.BookRoute(router)
	routes.AuthorRoute(router)
	routes.CategoryRoute(router)
	log.Println("Server is running on PORT ", config.ConnectPort())
	router.Run(config.ConnectPort())
}
