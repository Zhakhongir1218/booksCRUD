package book

import (
	"booksCRUD/internal/configuration"
	"fmt"
	"log"
)

const insertStatement string = "insert into book (name, price) values ($1, $2)"
const findByIdStatement string = "select * from book where id=$1"
const updateStatement string = "update book set name = $1, price = $2 where id = $3"
const getAllStatement string = "select * from book"

var conn = configuration.InitialConnectionPostgres

type DbConfiguration struct {
}

func (db *DbConfiguration) Insert(book Book) {
	_, err := conn.Exec(insertStatement, book.Name, book.Price)
	if err != nil {
		panic(err)
	}
	log.Println("Book was created")
}

func (db *DbConfiguration) FindBookById(id int) Book {
	row := conn.QueryRow(findByIdStatement, id)
	b := Book{}
	err := row.Scan(&b.Id, &b.Name, &b.Price)
	if err != nil {
		log.Println(err)
		return b
	}
	return b
}

func (db *DbConfiguration) UpdateBook(book Book) {
	_, err := conn.Exec(updateStatement, book.Name, book.Price, book.Id)
	if err != nil {
		panic(err)
	}
}

func (db *DbConfiguration) GetBooks() []Book {
	var books []Book
	rows, err := conn.Query(getAllStatement)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}
	return books
}
