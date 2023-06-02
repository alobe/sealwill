package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Log() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		uri := c.Request().URI()
		path := string(uri.Path())
		query := string(uri.QueryString())
		c.Next()
		cost := time.Since(start)
		log.Info().
			Int("status", c.Response().StatusCode()).
			Str("method", string(c.Request().Header.Method())).
			Str("path", path).
			Str("query", query).
			Str("ip", c.IP()).
			Str("ua", string(c.Request().Header.UserAgent())).
			Dur("cost", cost).
			Err(c.Context().Err()).
			Msg("router inspect log")

		return nil
	}
}
