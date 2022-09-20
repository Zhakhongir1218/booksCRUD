package book

import (
	"fmt"
)

type Book struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func (b *Book) GetBookById(idAsInt int, db CrudRepository) Book {
	return db.FindBookById(idAsInt)
}

func (b *Book) CreateNewBook(book Book, db CrudRepository) {
	if book.Name == "" {
		_ = fmt.Errorf("Book's name can't be empty")
	}
	db.Insert(book)
}

func (b *Book) GetAllBooks(db CrudRepository) []Book {
	return db.GetBooks()
}

func (b *Book) UpdateBook(book Book, db CrudRepository) Book {
	if book.Name == "" {
		_ = fmt.Errorf("Book's name can't be empty")
	}
	db.UpdateBook(book)
	return book
}
