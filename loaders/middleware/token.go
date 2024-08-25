package middleware

import (
	"backend/internal/domain/response"
	"backend/loaders/config"
	"github.com/gofiber/fiber/v2"
)

func TokenMiddleWare(c *fiber.Ctx) error {
	token := config.Conf.Token
	if c.Get("Authorization") != "Bearer "+token {
		return c.JSON(response.Error(false, "Token is not valid"))
	}
	return c.Next()

}
