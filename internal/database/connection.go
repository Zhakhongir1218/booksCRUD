package database

import (
	"booksCRUD/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Init() *sql.DB {
	_, dbConfig := config.LoadConfig(".", "db.properties")
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}
	log.Println("Connection success")
	return db
}
