package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo/exception"

	_ "github.com/lib/pq"
)

func NewDb(config Config) *sql.DB {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Get("DB_HOST"), config.Get("DB_USER"), config.Get("DB_PASSWORD"), config.Get("DB_NAME")))

	exception.PanicIfErr(err)
	return db
}

func NewDbContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
