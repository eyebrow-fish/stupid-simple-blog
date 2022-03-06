package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func init() {
	host := env("POSTGRES_HOST", "localhost")
	username := env("POSTGRES_USER", "postgres")
	password := env("POSTGRES_PASSWORD", "password")

	var err error
	s := "postgres://" + username + ":" + password + "@" + host + "/?sslmode=disable"
	DB, err = sql.Open("postgres", s)
	if err != nil {
		panic(err)
	}
}

func env(k string, f string) string {
	x := os.Getenv("POSTGRES_HOST")
	if x == "" {
		return f
	} else {
		return x
	}
}
