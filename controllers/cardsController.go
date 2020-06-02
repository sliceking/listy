package controllers

import (
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

// Index will display all cards
// GET /card
func (c *CardsController) Index() func(ctx *gin.Context) {
	var backlog []models.Card
	var working []models.Card
	var done []models.Card

	c.DB.Where("category = ?", "backlog").Find(&backlog)
	c.DB.Where("category = ?", "working").Find(&working)
	c.DB.Where("category = ?", "done").Find(&done)

	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"backlog": backlog,
			"working": working,
			"done":    done,
		})
	}
}

// Show will display an individual card
// GET /card/:id
func (c *CardsController) Show() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var card models.Card
		id := ctx.Param("id")

		c.DB.Where("ID = ?", id).Preload("Tasks").Find(&card)

		ctx.HTML(http.StatusOK, "show.html", gin.H{
			"card": card,
		})
	}
}

// Update will update a card record
// PUT /card/:id
func (c *CardsController) Update() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var card models.Card
		id := ctx.Param("id")

		// DO UPDATE HERE
	}
}

// Create will make a new card record
// POST /card
func (c *CardsController) Create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var card models.Card

		// DO CREATE HERE
	}
}
