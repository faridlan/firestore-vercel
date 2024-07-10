package config

import (
	"github.com/faridlan/firestore-vercel/controller"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	UserController controller.UserController
}

func NewRouter(router Router) *fiber.App {

	app := fiber.New()

	app.Post("/api/users", router.UserController.Save)
	app.Get("/api/users", router.UserController.Find)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app

}
