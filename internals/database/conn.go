package database

import (
	"context"
	"fmt"
	"log"
	"pg_crud/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Conn() *pgxpool.Pool{

	pool, err := pgxpool.New(context.Background(), configs.Dsn)
	if err != nil{
		log.Println(configs.ErrString, err)
	}
	err = pool.Ping(context.Background())
	if err != nil{
		log.Println(configs.ErrString, err)
	}
	fmt.Println("Successfully Connected!")

	return pool

}