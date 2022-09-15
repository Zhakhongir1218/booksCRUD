package controllers

import (
	_ "booksCRUD/docs"
	"booksCRUD/internal/book"
	"booksCRUD/internal/user"
	"fmt"
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

// @host localhost:8080
func RouterCreation() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/api/books/:id", findBookById)
	app.Get("/api/books", getAllBooks)
	app.Post("/api/books", createNewBook)
	app.Post("/auth/sign-up", signUp)
	app.Post("/auth/sign-in", signIn)

	log.Fatal(app.Listen(":8080"))
}

// sign-in functionality
// @Summary sign-in
// @Description sign-in process
// @Tags auth
// @Accept json
// @Produce json
// @Param user body user.User true "AAA 123"
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /auth/sign-in [post]
func signIn(ctx *fiber.Ctx) error {
	u := new(user.User)
	err := ctx.BodyParser(u)
	if err != nil {
		return err
	}
	token := user.SignIn(u.Name, u.Password)
	return ctx.JSON(ResponseHTTP{true, u.Name, token})
}

// sign-up functionality
// @Summary sign-up
// @Description Sign-Up process
// @Tags auth
// @Accept json
// @Produce json
// @Param user body user.User true "AAA 123"
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /auth/sign-up [post]
func signUp(ctx *fiber.Ctx) error {
	u := new(user.User)
	err := ctx.BodyParser(u)
	if err != nil {
		return err
	}
	user.SignUp(u.Name, u.Password)
	return ctx.JSON(ResponseHTTP{true, u.Name, "signed-up"})
}

// CreateNewBook is a function to create a new book
// @Summary Book's creation
// @Description Create new book
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Param book body book.Book true "1, Some-Book, 200"
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/books [post]
func createNewBook(ctx *fiber.Ctx) error {
	_, err := jwtValidation(ctx)
	if err != nil {
		return err
	}
	b := new(book.Book)
	err = ctx.BodyParser(b)
	book.CreateNewBook(*b)
	return ctx.JSON(ResponseHTTP{true, b.Name, "Book was successfully created"})
}

// GetBooks is a function to get all book
// @Summary Get all
// @Description Get all book
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/books [get]
func getAllBooks(ctx *fiber.Ctx) error {
	return ctx.JSON(book.GetAllBooks())
}

// GetBookByID is a function to get a book by ID
// @Summary Get book by ID
// @Description Get book by ID
// @Security ApiKeyAuth
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/books/{id} [get]
func findBookById(c *fiber.Ctx) error {
	_, err := jwtValidation(c)
	if err != nil {
		return err
	}
	msg := c.Params("id")
	log.Println(msg)
	idAsInt, _ := strconv.Atoi(msg)
	b := book.GetBookById(idAsInt)
	return c.JSON(b)
}

func jwtValidation(c *fiber.Ctx) (user.User, error) {
	headerMap := c.GetReqHeaders()
	var jwtToken = headerMap["Authorization"]
	if jwtToken == "" {
		return user.User{}, fmt.Errorf("did't find out a token")
	}
	u := user.ReadToken(jwtToken)
	fmt.Println(u)
	return u, nil
}
