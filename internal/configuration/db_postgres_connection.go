package configuration

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var InitialConnectionPostgres = Init()

func Init() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=zhakhongir1218 dbname=go-crud sslmode=disable")
	if err != nil {
		panic(err)
	}
	log.Println("Connection success")
	return db
}
