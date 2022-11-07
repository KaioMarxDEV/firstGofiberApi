package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcomeController(c *fiber.Ctx) error {
	return c.SendString("Welcome http api")
}

func main() {
	var app = fiber.New()

	app.Get("/", welcomeController)

	log.Fatal(app.Listen(":3000"))
}
