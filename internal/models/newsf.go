package models

type News struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}
