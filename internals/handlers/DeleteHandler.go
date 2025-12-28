package handlers

import (
	"net/http"
)

func (c * ConnPool) DeleteHandler(w http.ResponseWriter, r * http.Request){
	id := r.PathValue("id")

	query := "DELETE FROM books WHERE id = $1"

	tag, err := c.Db.Exec(r.Context(), query, id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsCount := tag.RowsAffected()
	if rowsCount == 0{
		http.Error(w, "No Rows Affected", http.StatusNotFound)
		return
	}


	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Entry Deleted!")
	w.WriteHeader(http.StatusNoContent)
}