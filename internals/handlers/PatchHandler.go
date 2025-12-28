package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func(c * ConnPool) PatchHandler(w http.ResponseWriter, r * http.Request){

	id := r.PathValue("id")

	var updates map[string]any

	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(updates) == 0{
		http.Error(w, "Nothing To Update\n", http.StatusBadRequest)
		return
	}

	query := "UPDATE books SET "
	args := []any{}
	argId := 1

	for col, val := range updates{

		if col != "title" && col != "author" && col != "category"{
			continue
		}

		query += fmt.Sprintf("%s = $%d, ", col, argId)
		args = append(args, val)
		argId++

	}

	query = query[:len(query)-1]

	query += fmt.Sprintf("WHERE id = $%d", argId)
	args = append(args, id)

	res, err := c.Db.Exec(r.Context(), query, args...)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.RowsAffected() == 0{
		http.Error(w, "No Rows Affected", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Book Patched!")

}