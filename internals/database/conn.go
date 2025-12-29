package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"pg_crud/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Conn() *pgxpool.Pool{

	godotenv.Load()
	dsn := os.Getenv("DB_DSN")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil{
		log.Fatalf(configs.ErrString + "%v", err)
	}
	err = pool.Ping(context.Background())
	if err != nil{
		log.Fatalf(configs.ErrString + "%v", err)
	}
	fmt.Println("Successfully Connected!")

	return pool

}