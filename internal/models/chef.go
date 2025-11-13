package models

import "gorm.io/gorm"

type Chef struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url"`
	Contacts    string `json:"contacts"`
}
