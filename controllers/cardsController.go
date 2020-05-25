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
	var backlog []models.Card
	var working []models.Card
	var done []models.Card
	c.DB.Where("category = ?", "backlog").Find(&backlog)
	c.DB.Where("category = ?", "working").Find(&working)
	c.DB.Where("category = ?", "done").Find(&done)

	fmt.Println(backlog)

	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"backlog": backlog,
			"working": working,
			"done":    done,
		})
	}
}
