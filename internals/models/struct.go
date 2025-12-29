package models

type BookData struct{
	Id 		int `json:"id"`
	Title 	string `json:"title" validate:"required,min=3,max=100,forbidden"`
	Author 	string `json:"author" validate:"required"`
	Category string `json:"category" validate:"required,oneofci=Tech Fantasy Misc. Programming"`
}