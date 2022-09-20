package book

type CrudRepository interface {
	Insert(book Book)
	FindBookById(id int) Book
	UpdateBook(book Book)
	GetBooks() []Book
}
