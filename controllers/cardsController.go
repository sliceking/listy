package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sliceking/listy/models"
	"github.com/sliceking/listy/service"
)

type CardsController interface {
	FindAll() []models.Card
	Save(*gin.Context) models.Card
}

type controller struct {
	service service.CardsService
}

func New(service service.CardsService) CardsController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Card {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Card {
	var card models.Card
	ctx.BindJSON(&card)
	c.service.Save(card)
	return card
}
