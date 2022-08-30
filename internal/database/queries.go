package database

import (
	"booksCRUD/internal/book"
	"database/sql"
	"log"
)

func Create(book book.Book, db *sql.DB) {
	sqlStatement := "insert into books (name, price) values ($1, $2)"
	_, err := db.Exec(sqlStatement, book.Name, book.Price)
	if err != nil {
		panic(err)
	}
	log.Println("Book was created")
}

func FindBookById(id int, db *sql.DB) book.Book {
	findByIdStatement := "select * from books where id=$1"
	row := db.QueryRow(findByIdStatement, id)
	b := book.Book{}
	err := row.Scan(&b.ID, &b.Name, &b.Price)
	if err != nil {
		log.Println(err)
		return b
	}
	return b
}
