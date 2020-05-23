package models

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Tasks       []Task
	Completed   bool
}

type Task struct {
	gorm.Model
	CardID      uint
	Description string `gorm:"not null"`
}
