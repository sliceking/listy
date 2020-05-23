package service

import "github.com/sliceking/listy/models"

type CardsService interface {
	Save(models.Card) models.Card
	FindAll() []models.Card
}

type cardsService struct {
	cards []models.Card
}

func New() CardsService {
	return &cardsService{}
}

func (service *cardsService) Save(card models.Card) models.Card {
	service.cards = append(service.cards, card)
	return card
}

func (service *cardsService) FindAll() []models.Card {
	return service.cards
}
