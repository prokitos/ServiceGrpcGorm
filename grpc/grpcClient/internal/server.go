package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func MainServer() {

	app := fiber.New()

	handlers(app)

	log.Fatal(app.Listen(":6666"))
}

func handlers(instance *fiber.App) {

	instance.Get("/login", LoginRoute)
	instance.Get("/register", RegisterRoute)

}
