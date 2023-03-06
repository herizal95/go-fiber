package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {

	router := fiber.New()
	app := fiber.New()

	app.Mount("/api", router)
	router.Get("/data", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"code":    http.StatusOK,
			"status":  "success",
			"message": "welcome to golang, fiber, and gorm",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
