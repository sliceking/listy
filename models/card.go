package models

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Category    string `gorm:"not null"`
	Tasks       []Task
	Completed   bool
}
