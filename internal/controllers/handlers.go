package controllers

import (
	_ "booksCRUD/docs"
	"booksCRUD/internal/book"
	"booksCRUD/internal/book/database"
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
	headerMap := c.GetReqHeaders()
	var jwtToken string = headerMap["Jwt"]
	if jwtToken == "" {
		return fmt.Errorf("did't find out a token")
	}
	u := user.ReadToken(jwtToken)
	fmt.Println(u)
	msg := c.Params("id")
	log.Println(msg)
	idAsInt, _ := strconv.Atoi(msg)
	var b book.Book
	b = database.FindBookById(idAsInt)

	return c.JSON(b)
}
