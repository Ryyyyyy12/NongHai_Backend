package loaders

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/loaders/config"
	"backend/loaders/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InitRoutes() {

	trackingRepo := repository.NewTrackingRepository(*DB)
	userRepo := repository.NewUserRepository(*DB)
	petRepo := repository.NewPetRepository(*DB)

	trackingService := service.NewTrackingService(trackingRepo, userRepo, petRepo)
	userService := service.NewUserService(userRepo)
	trackingHandler := handler.NewTrackingHandler(trackingService, userService)

	app := InitFiber()

	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Success": true,
			"Message": "NongHai API💫",
		})
	})

	app.Use(middleware.TokenMiddleWare)
	apiGroup := app.Group("/api")
	apiGroup.Post("tracking/createTracking", trackingHandler.CreateTracking)
	apiGroup.Get("tracking/getTracking", trackingHandler.GetTracking)
	apiGroup.Use(middleware.Cors())

	Serve(app, config.Conf.Address)

}
func InitFiber() *fiber.App {

	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		Network:       fiber.NetworkTCP4,
		AppName:       "NongHai API💫",
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	return app
}
func Serve(app *fiber.App, address string) {
	err := app.Listen(address)
	if err != nil {
		logrus.Fatal("Error starting server: ", err)
	}
}
