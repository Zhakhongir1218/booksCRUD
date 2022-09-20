package book

type CrudService interface {
	GetBookById(idAsInt int, db CrudRepository) Book
	CreateNewBook(b Book, db CrudRepository)
	GetAllBooks(db CrudRepository) []Book
	UpdateBook(b Book, db CrudRepository) Book
}
