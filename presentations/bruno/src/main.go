package main

import (
	"catalog/api"
	"catalog/config"
	"catalog/repository"
	"catalog/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	jwtware "github.com/gofiber/contrib/jwt"
)

func main() {

	conf := config.NewConfig()
	repo, _ := repository.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	productService := service.NewProductService(repo)
	loginService := service.NewLoginService()
	handler := api.NewHandler(productService, loginService)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Post("/login", handler.Login)
	app.Get("/products/:code", handler.Get)
	app.Get("/products", handler.GetAll)

	// Add jwt middleware for write routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte("secret"), // Dumb algorithm for testing purposes
		},
	}))
	app.Delete("/products/:code", handler.Delete)
	app.Put("/products", handler.Put)
	app.Post("/products", handler.Post)
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
