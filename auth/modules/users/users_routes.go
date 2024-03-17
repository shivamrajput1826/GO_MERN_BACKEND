package users

import (
	"go-auth/store"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(prefix string, app *fiber.App) {
	// service
	usersService := &UsersService{Client: store.ConnectDB()}
	// Controller
	usersController := InitUserController(usersService)

	applicationRoutes := app.Group(prefix + "/users")
	applicationRoutes.Get("/currentUser")
	applicationRoutes.Post("/signin")
	applicationRoutes.Post("/login")
	applicationRoutes.Post("/signout")
}
