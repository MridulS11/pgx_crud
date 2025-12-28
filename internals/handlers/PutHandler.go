package handlers

import (
	"encoding/json"
	"net/http"
	"pg_crud/internals/models"
)

func(c * ConnPool) PutHandler(w http.ResponseWriter, r * http.Request){

	id := r.PathValue("id")

	var book models.BookData

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE books SET title = $1, author = $2, category = $3 WHERE id = $4 RETURNING id"

	err := c.Db.QueryRow(r.Context(), query, book.Title, book.Author, book.Category, id).Scan(&book.Id)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}