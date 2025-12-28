package handlers

import (
	"encoding/json"
	"net/http"
	"pg_crud/internals/models"
)

func(c * ConnPool) GetHandler(w http.ResponseWriter, r *http.Request){

	query := "SELECT id, title, author, category FROM books"

	rows, err := c.Db.Query(r.Context(), query)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	books := []models.BookData{}

	for rows.Next(){
		var b models.BookData
		err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.Category)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(books)

}