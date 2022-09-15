package book

import (
	"fmt"
)

type Book struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func GetBookById(idAsInt int) Book {
	return FindBookById(idAsInt)
}

func CreateNewBook(b Book) {
	if b.Name == "" {
		_ = fmt.Errorf("Book's name can't be empty")
	}
	Insert(b)
}

func GetAllBooks() []Book {
	return GetBooks()
}
