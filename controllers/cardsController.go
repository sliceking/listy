package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sliceking/listy/models"
)

func NewCardsController(db *gorm.DB) *CardsController {
	return &CardsController{DB: db}
}

type CardsController struct {
	DB *gorm.DB
}

func (c *CardsController) Index() func(ctx *gin.Context) {
	var cards []models.Card
	c.DB.Find(&cards)
	fmt.Println(cards)

	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Cards": cards,
		})
	}
}
