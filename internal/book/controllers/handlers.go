package controllers

import (
	"booksCRUD/internal/book"
	"booksCRUD/internal/database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func RouterCreation() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/books/:id", findBookById)
	router.GET("/api/books", getAll)
	router.POST("/api/books", insertBook)
	router.PUT("/api/books", updateBook)
	return router
}

func getAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := database.Init()
	writer.Header().Set("Content-Type", "application/json")
	books := database.GetAllBooks(db)
	json.NewEncoder(writer).Encode(books)
}

func updateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := database.Init()
	writer.Header().Set("Content-Type", "application/json")
	var b book.Book
	_ = json.NewDecoder(request.Body).Decode(&b)
	database.UpdateBook(b, db)
	err := db.Close()
	if err != nil {
		log.Println("Closing connection error")
	}
}

func insertBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := database.Init()
	writer.Header().Set("Content-Type", "application/json")
	var b book.Book
	_ = json.NewDecoder(request.Body).Decode(&b)
	database.Create(b, db)
	err := db.Close()
	if err != nil {
		log.Println("Closing connection error")
	}
}

func findBookById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := database.Init()
	writer.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	var b book.Book
	b = database.FindBookById(idAsInt, db)
	json.NewEncoder(writer).Encode(b)
}
