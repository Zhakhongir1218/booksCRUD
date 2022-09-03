package book

type Book struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
