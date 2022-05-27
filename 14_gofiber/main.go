package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to my App!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("This is a Post Request!")
	})

	app.Get("/download", func(c *fiber.Ctx) error {
		return c.SendFile("gopher_.svg")
	})

	app.Static("/", "./files")
	// => http://localhost:3000/samplefile.txt

	fmt.Println("Starting the Web Server...")
	log.Fatal(app.Listen(":3000"))

}
