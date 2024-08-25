package middleware

import (
	"backend/loaders/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors = func() fiber.Handler {

	origin := ""
	for i, v := range config.Conf.Cors {
		origin += v
		if i != len(config.Conf.Cors)-1 {
			origin += ","
		}
	}

	//set cors config
	configCors := cors.Config{
		AllowOrigins:     origin,
		AllowCredentials: true,
	}

	return cors.New(configCors)
}
