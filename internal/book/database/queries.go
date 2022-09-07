package database

import (
	"booksCRUD/internal/book"
	"booksCRUD/internal/configuration"
	"fmt"
	"log"
)

var db = configuration.InitialConnection

const insertStatement string = "insert into books (name, price) values ($1, $2)"
const findByIdStatement string = "select * from books where id=$1"
const updateStatement string = "update books set name = $1, price = $2 where id = $3"
const getAllStatement string = "select * from books"

func Insert(book book.Book) {
	_, err := db.Exec(insertStatement, book.Name, book.Price)
	if err != nil {
		panic(err)
	}
	log.Println("Book was created")
}

func FindBookById(id int) book.Book {
	row := db.QueryRow(findByIdStatement, id)
	b := book.Book{}
	err := row.Scan(&b.Id, &b.Name, &b.Price)
	if err != nil {
		log.Println(err)
		return b
	}
	return b
}

func UpdateBook(book book.Book) {
	_, err := db.Exec(updateStatement, book.Name, book.Price, book.Id)
	if err != nil {
		panic(err)
	}
}

func GetAllBooks() []book.Book {
	var books []book.Book
	rows, err := db.Query(getAllStatement)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		b := book.Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}
	return books
}
