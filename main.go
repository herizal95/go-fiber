package main

import (
	"fmt"
	"log"

	"github.com/go-jwt-test/config"
	"github.com/go-jwt-test/controller"
	"github.com/go-jwt-test/repository"
	"github.com/go-jwt-test/router"
	"github.com/go-jwt-test/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Print("Run service.....")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load .env variabel", err)
	}

	// Database
	db := config.ConnectDatabase(&loadConfig)
	validate := validator.New()

	// Init Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	// Init Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	// Init Controller
	userController := controller.NewUserController(userService)

	// Router
	routes := router.NewRouter(userController)

	app := fiber.New()

	app.Mount("/api", routes)
	log.Fatal(app.Listen(":8000"))

}
