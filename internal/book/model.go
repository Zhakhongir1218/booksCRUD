package book

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
