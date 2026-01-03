package handlers

import (
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnPool struct {
	Db *pgxpool.Pool
	Validate *validator.Validate
}

func TitleCheck(fl validator.FieldLevel) bool{
	title := fl.Field().String()
	return !strings.Contains(strings.ToLower(title), "forbidden")
}

func NewConnPool(Db *pgxpool.Pool) *ConnPool{
	v := validator.New()
	v.RegisterValidation("forbidden", TitleCheck)

	return &ConnPool{
		Db: Db,
		Validate: v,
	}
}