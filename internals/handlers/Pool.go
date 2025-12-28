package handlers

import "github.com/jackc/pgx/v5/pgxpool"

type ConnPool struct {
	Db *pgxpool.Pool
}
