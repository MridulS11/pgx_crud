package models

type BookData struct{
	Id 		int `json:"id"`
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	Category string `json:"category"`
}