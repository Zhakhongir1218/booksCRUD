package main

import (
	"booksCRUD/internal/configuration"
	"booksCRUD/internal/controllers"
)

// @title books-crud-service
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com
func main() {
	router := controllers.RouterCreation()
	configuration.CreateServer(router)
}
