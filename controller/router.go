package controller

import (
	"github.com/alobe/seawill/controller/middleware"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func InitRouter() {
	app = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET",
		}),
		middleware.Log(),
	)

	api := app.Group("/api")

	user := api.Group("/user")

	user.Post("/register", register)
	user.Post("/login", login)
	user.Get("/list", middleware.CheckCookie, getUserList)

	app.Listen(":3000")
}
