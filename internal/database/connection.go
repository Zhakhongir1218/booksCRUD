package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Init() *sql.DB {
	connStr := "user=postgres password=zhakhongir1218 dbname=go-crud sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	log.Println("Connection success")
	return db
}
