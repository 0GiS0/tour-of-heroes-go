package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	AlterEgo    string `json:"alterEgo"`
	Description string `json:"description"`
}
