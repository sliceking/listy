package models

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Title       string `gorm:"not null" form:"title"`
	Description string `gorm:"not null" form:"description"`
	Category    string `gorm:"not null" form:"category"`
	Tasks       []Task
	Completed   bool
}
