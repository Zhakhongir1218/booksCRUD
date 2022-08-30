package main

import (
	"booksCRUD/internal/book/controllers"
	"booksCRUD/internal/configuration"
)

func main() {
	router := controllers.RouterCreation()
	configuration.CreateServer(router)
}
