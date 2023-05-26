package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type LoginParams struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func login(ctx *fiber.Ctx) error {
	params := new(LoginParams)
	ctx.BodyParser(params)
	fmt.Println(params)
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Hello World",
	})
}
