package main

import (
	"booksCRUD/internal/configuration"
	"booksCRUD/internal/controllers"
)

func main() {
	router := controllers.RouterCreation()
	configuration.CreateServer(router)
}
