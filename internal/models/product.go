package models

type Dessert struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Details string  `json:"ingredients"`
	Price       string `json:"price"`
	ImageURL    string  `json:"image_url"`
}
