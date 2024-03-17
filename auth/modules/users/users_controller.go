package users

import "github.com/gofiber/fiber/v2"

type UsersController struct {
	usersService *UsersService
}

func InitUserController(usersService *UsersService) *UsersController {
	return &UsersController{usersService}
}

func (uc *UsersController) GetCurrentUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
