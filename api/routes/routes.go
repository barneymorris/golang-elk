package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	HandleHello    func(*fiber.Ctx) error
	HandleGenError func(*fiber.Ctx) error
}

func GetRoutes() *Routes {
	return &Routes{
		HandleHello:    HandleHello,
		HandleGenError: HandleGenError,
	}
}
