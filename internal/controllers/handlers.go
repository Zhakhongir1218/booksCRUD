package controllers

import (
	"booksCRUD/internal/book"
	"booksCRUD/internal/book/database"
	"booksCRUD/internal/user"
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
	router.POST("/api/sign-up", signUp)
	return router
}

func getAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	books := database.GetAllBooks()
	err := json.NewEncoder(writer).Encode(books)
	if err != nil {
		return
	}
}

func updateBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var b book.Book
	_ = json.NewDecoder(request.Body).Decode(&b)
	database.UpdateBook(b)
}

func insertBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var b book.Book
	_ = json.NewDecoder(request.Body).Decode(&b)
	database.Insert(b)
}

func findBookById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	var b book.Book
	b = database.FindBookById(idAsInt)
	json.NewEncoder(writer).Encode(b)
}

func signUp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var u user.User
	_ = json.NewDecoder(request.Body).Decode(&u)
	user.SignUp(u.Name, u.Password)
	writer.WriteHeader(200)
	_, err := writer.Write([]byte(u.Name))
	if err != nil {
		return
	}
}
