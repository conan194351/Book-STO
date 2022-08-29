package routes

import (
	"book-sto/handlers"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Service(router *gin.Engine, conn *grpc.ClientConn) {
	route := router.Group("/api/service/author")
	{
		route.GET("/:a/:b", handlers.FindBookByIdAuthor(conn))
	}
}
