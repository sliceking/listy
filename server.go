package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sliceking/listy/controllers"
	"github.com/sliceking/listy/service"
)

var (
	cardsService    service.CardsService        = service.New()
	cardsController controllers.CardsController = controllers.New(cardsService)
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("views/*")

	// controllers can return funcs that are gin.HandlerFuncs: func(ctx *gin.Context) {}
	server.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"message": "ok!",
			"ok":      "no",
		})
	})

	server.GET("/cards", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cardsController.FindAll())
	})

	// get post to work with postman
	server.POST("/cards", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cardsController.Save(ctx))
	})

	server.Run(":8080")
}
