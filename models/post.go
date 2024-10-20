package models

type Post struct{
	ID uint `json: "id" gorm:"primary_key"`
	Title string `json:"title"`
	Body string `json:"body"`
	UserID uint `json: "user_id"`
}