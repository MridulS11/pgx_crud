package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pg_crud/configs"
	"pg_crud/internals/models"

	"github.com/jackc/pgx/v5"
)

func(c *ConnPool) GetByIdHandler(w http.ResponseWriter, r *http.Request){

	id := r.PathValue("id")

	query := "SELECT id, title, author, category FROM books WHERE id = $1"

	var book models.BookData
	err := c.Db.QueryRow(r.Context(), query, id).Scan(&book.Id, &book.Title, &book.Author, &book.Category)
	if err != nil{
		if err == pgx.ErrNoRows{
			log.Println(configs.ErrString, err.Error())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
	
}