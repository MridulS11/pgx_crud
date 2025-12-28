package main

import (
	"pg_crud/internals/handlers"
	"pg_crud/internals/database"
	"pg_crud/internals/server"
)


func main(){
	pool := database.Conn()
	defer pool.Close()
	bookhandler := &handlers.ConnPool{Db: pool}
	server.Server(bookhandler)
}