package router

import (
	"github.com/go-jwt-test/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(userController *controller.UserController) *fiber.App {

	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to golang, fiber, and gorm",
		})

	})

	router.Route("/users", func(router fiber.Router) {
		router.Post("/", userController.Create)
		router.Get("", userController.FindAll)
	})

	router.Route("users/:uid", func(router fiber.Router) {
		router.Delete("", userController.Delete)
		router.Get("", userController.FindById)
		router.Patch("", userController.Update)
	})

	return router
}
