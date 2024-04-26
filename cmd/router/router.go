package router

import (
	"scalex/internal/handler"
	"scalex/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	router.POST("/login", handler.Login)
	router.GET("/home", middleware.AuthMiddleware(), handler.Home)
	router.POST("/addBook", middleware.AuthMiddleware(), handler.AddBook)
	router.DELETE("/deleteBook", middleware.AuthMiddleware(), handler.DeleteBook)

	router.Run(":8080")

}
