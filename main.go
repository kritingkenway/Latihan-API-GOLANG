package main

import (
	"coba/model"

	"coba/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Koneksi ke database
	
	model.DatabaseInit()
	
	app := fiber.New()
	
	
	routes.Setup(app)
	

	app.Listen(":3000")
}