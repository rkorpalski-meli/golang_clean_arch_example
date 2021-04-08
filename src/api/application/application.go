package application

import (
	"github.com/gin-gonic/gin"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/application/middleware"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/application/registry"
)

func StartServer() {
	bookHandler := registry.CreateBookHandler()

	router := gin.Default()
	router.GET("/book/find/:title", bookHandler.FindBook)
	router.POST("/book/find", bookHandler.InsertBook)

	router.Use(middleware.Cors())

	router.Run()
}
