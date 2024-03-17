package main

import (
	"go-auth/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(prefix string, app *fiber.App){
	user.
}

func main() {
	client := store.ConnectDB()
	if client == nil {
		panic("Failed to connect to MongoDB")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // any request body size larger than 100mb will be rejected
		Immutable: true,              // will set static files as immutable, download them, cache them and use them in future
	})

	API_PREFIX := "/api"
	SetupRoutes(API_PREFIX, app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	error := app.Listen(":3000")
	if error != nil {
		panic(error)
	}
}
