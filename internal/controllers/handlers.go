package controllers

import (
	_ "booksCRUD/docs"
	"booksCRUD/internal/book"
	"booksCRUD/internal/book/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
	"strconv"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// @title Book App
// @version 1.0
// @description This is an API for Book Application

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func RouterCreation() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/api/books/:id", findBookById)
	app.Get("/api/books", getAllBooks)

	log.Fatal(app.Listen(":8080"))
}

// GetBooks is a function to get all books
// @Summary Get all
// @Description Get all books
// @Tags books
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/books [get]
func getAllBooks(ctx *fiber.Ctx) error {
	return ctx.JSON(database.GetAllBooks())
}

// GetBookByID is a function to get a book by ID
// @Summary Get book by ID
// @Description Get book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/books/{id} [get]
func findBookById(c *fiber.Ctx) error {
	msg := c.Params("id")
	log.Println(msg)
	idAsInt, _ := strconv.Atoi(msg)
	var b book.Book
	b = database.FindBookById(idAsInt)

	return c.JSON(b)
}
