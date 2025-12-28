package handlers

import (
	"encoding/json"
	"net/http"
	"pg_crud/configs"
	"pg_crud/internals/models"
)

func (c *ConnPool) PostHandler(w http.ResponseWriter, r * http.Request){

	var book models.BookData

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil{
		http.Error(w, configs.ErrString, http.StatusBadRequest)
		return
	}

	query := "INSERT into books (title, author, category) VALUES ($1, $2, $3) RETURNING id"

	err := c.Db.QueryRow(r.Context(), query, book.Title, book.Author, book.Category).Scan(&book.Id)
	if err != nil{
		http.Error(w, configs.ErrString + err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}