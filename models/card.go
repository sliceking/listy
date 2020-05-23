package models

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Active      bool   ``
}
