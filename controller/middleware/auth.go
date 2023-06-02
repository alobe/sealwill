package middleware

import (
	"context"
	"fmt"

	"github.com/alobe/seawill/lib"
	"github.com/gofiber/fiber/v2"
)

func CheckCookie(ctx *fiber.Ctx) error {
	key := ctx.Cookies("x-seawill")
	fmt.Println(key)
	if key == "" {
		return fiber.ErrUnauthorized
	}

	if _, err := lib.Rds.Get(context.Background(), key).Result(); err != nil {
		return fiber.ErrUnauthorized
	}

	ctx.Next()
	return nil
}
