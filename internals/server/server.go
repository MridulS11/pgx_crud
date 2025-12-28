package server

import (
	"log"
	"net/http"
	"pg_crud/configs"
	"pg_crud/internals/handlers"
	"pg_crud/internals/middleware"
)

func Server(s *handlers.ConnPool){
	mux := http.NewServeMux()
	log.Println("Server Started")
	mux.HandleFunc("POST /books", s.PostHandler)
	mux.HandleFunc("GET /books", s.GetHandler)
	mux.HandleFunc("GET /books/{id}", s.GetByIdHandler)
	mux.HandleFunc("DELETE /books/{id}", s.DeleteHandler)
	mux.HandleFunc("PUT /books/{id}", s.PutHandler)
	mux.HandleFunc("PATCH /books/{id}", s.PatchHandler)
	middle := middleware.SetHeader(middleware.SecLayer(mux))
	if err := http.ListenAndServe(":8080", middle); err != nil{
		log.Println(configs.ErrString + err.Error())
	}
}