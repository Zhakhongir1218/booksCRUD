package main

import (
	"booksCRUD/internal/controllers"
)

// @title Book App
// @version 1.0
// @description This is an API for Book Application
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	controllers.RouterCreation()
}
