package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sliceking/listy/models"
)

func Seeder(db *gorm.DB) {
	db.LogMode(true)
	db.DropTableIfExists(&models.Card{}, &models.Task{})
	db.AutoMigrate(&models.Card{}, &models.Task{})
	var seeds []models.Card

	seeds = append(seeds, models.Card{
		Title:       "I am the backlog First",
		Description: "Do some really good stuff and then make sure it works",
		Category:    "backlog",
		Tasks: []models.Task{
			{
				Description: "do it",
			},
			{
				Description: "what eva",
			},
		},
	})
	seeds = append(seeds, models.Card{
		Title:       "I am the backlog second",
		Description: "Do some really good stuff and then make sure it works",
		Category:    "backlog",
		Tasks: []models.Task{
			{
				Description: "do it",
			},
			{
				Description: "what eva",
			},
		},
	})

	seeds = append(seeds, models.Card{
		Title:       "I am the working first",
		Description: "Do some mediocre things and then make it maybe work",
		Category:    "working",
		Tasks: []models.Task{
			{
				Description: "sumthangt",
			},
			{
				Description: "bloppblopb",
			},
		},
	})
	seeds = append(seeds, models.Card{
		Title:       "I am the working second",
		Description: "Do some mediocre things and then make it maybe work",
		Category:    "working",
		Tasks: []models.Task{
			{
				Description: "sumthangt",
			},
			{
				Description: "bloppblopb",
			},
		},
	})

	seeds = append(seeds, models.Card{
		Title:       "I am the done third",
		Description: "Do some bad things and then dont make it maybe work",
		Category:    "done",
		Tasks: []models.Task{
			{
				Description: "sumthangt",
			},
			{
				Description: "bloppblopb",
			},
		},
	})
	seeds = append(seeds, models.Card{
		Title:       "I am the done third",
		Description: "Do some bad things and then dont make it maybe work",
		Category:    "done",
		Tasks: []models.Task{
			{
				Description: "sumthangt",
			},
			{
				Description: "bloppblopb",
			},
		},
	})

	for _, seedCard := range seeds {
		db.Create(&seedCard)
	}
}
