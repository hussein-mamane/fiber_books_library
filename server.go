package main

import (
	"awesomeProject/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func init() {

}
func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ServerHeader:          "gunicorn",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: nil,
		Next:             nil,
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})

	})
	books := app.Group("/books")
	books.Get("/", controllers.GetAllBooks)
	books.Post("/", controllers.PostBook)
	books.Delete("/delete/:id", controllers.DeleteBookById)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
