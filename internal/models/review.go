package models

type Review struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Message  string `json:"description"`
	ImageURL string `json:"image_url"`
}
