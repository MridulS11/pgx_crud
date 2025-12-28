package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pg_crud/internals/models"
	"strconv"
)

func(c * ConnPool) GetHandler(w http.ResponseWriter, r *http.Request){

	limitstr := r.URL.Query().Get("limit")
	offsetstr := r.URL.Query().Get("offset")

	limit := 10
	if l, err := strconv.Atoi(limitstr); err == nil && l > 0{
		limit = l
	}

	offset := 0
	if o, err := strconv.Atoi(offsetstr); err == nil && o > 0{
		offset = o
	}

	var total int
	err := c.Db.QueryRow(r.Context(), "SELECT COUNT(*) FROM books").Scan(&total)

	if total <= limit{
		fmt.Fprintf(w, "Viewed All Results: %d\n", total)
	}else{
		fmt.Fprintf(w, "Viewed %d Results out of %d.\n", limit, total)
	}

	query := "SELECT id, title, author, category FROM books Order BY id LIMIT $1 OFFSET $2"	//Select itself is returning

	rows, err := c.Db.Query(r.Context(), query, limit, offset)
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