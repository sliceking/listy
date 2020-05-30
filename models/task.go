package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	CardID      uint
	Description string `gorm:"not null"`
	Completed   bool
}
