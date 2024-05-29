package server

import (
	"log"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

// запуск сервера. localhost:8888
func MainServer() {

	app := fiber.New()

	handlers(app)

	log.Fatal(app.Listen(":8888"))
}

func handlers(instance *fiber.App) {

	instance.Delete("/delete", services.CarDelete)
	instance.Get("/show", services.CarShow)
	instance.Post("/insert", services.CarCreate)

	// instance.Delete("/delete", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

}
