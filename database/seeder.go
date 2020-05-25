package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sliceking/listy/models"
)

func Seeder(db *gorm.DB) {
	db.LogMode(true)
	db.DropTableIfExists(&models.Card{}, &models.Task{})
	db.AutoMigrate(&models.Card{}, &models.Task{})

	newCard := models.Card{
		Title:       "I am the First",
		Description: "i am the",
		Tasks: []models.Task{
			{
				Description: "do it",
			},
			{
				Description: "what eva",
			},
		},
	}

	secondNewCard := models.Card{
		Title:       "I am the second",
		Description: "i am the",
		Tasks: []models.Task{
			{
				Description: "sumthangt",
			},
			{
				Description: "bloppblopb",
			},
		},
	}

	db.Create(&newCard)
	db.Create(&secondNewCard)
}
