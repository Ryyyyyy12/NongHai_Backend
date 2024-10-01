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

	//Repositories
	trackingRepo := repository.NewTrackingRepository(*DB)
	userRepo := repository.NewUserRepository(DB)
	petRepo := repository.NewPetRepository(*DB)
	chatRepo := repository.NewChatRepository(*DB)
	tokenRepo := repository.NewTokenRepository(*DB)

	//Services
	trackingService := service.NewTrackingService(trackingRepo, userRepo, petRepo)
	userService := service.NewUserService(userRepo)
	chatService := service.NewChatService(chatRepo)
	userFCMTokenService := service.NewUserFCMTokenService(tokenRepo)
	notificationService := service.NewNotificationService(userFCMTokenService)

	//Handlers
	trackingHandler := handler.NewTrackingHandler(trackingService, userService)
	ChatHandler := handler.NewChatHandler(chatService)
	userHandler := handler.NewUserHandler(userService)
	userTokenHandler := handler.NewUserTokenHandler(userFCMTokenService)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	app := InitFiber()

	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Success": true,
			"Message": "NongHai API💫",
		})
	})

	app.Use(middleware.TokenMiddleWare)
	apiGroup := app.Group("/api")

	trackingGroup := apiGroup.Group("/tracking")
	trackingGroup.Post("/createTracking", trackingHandler.CreateTracking)
	trackingGroup.Get("/getTracking", trackingHandler.GetTracking)

	chatGroup := apiGroup.Group("/chat")
	chatGroup.Post("/createChatRoom", ChatHandler.CreateChatRoom)
	chatGroup.Get("/getChatRoom", ChatHandler.GetChatRoom)
	chatGroup.Get("/getCurrentUserChatRoom", ChatHandler.GetCurrentUserChatRoom)
	chatGroup.Post("/setRead", ChatHandler.ReadChat)
	chatGroup.Post("/setUnread", ChatHandler.SetUnread)

	tokenGroup := apiGroup.Group("/token")
	tokenGroup.Post("/createUserToken", userTokenHandler.CreateUserFCMToken)
	// tokenGroup.Get("/checkIfTokenExist", userTokenHandler.CheckIfTokenExist)
	tokenGroup.Delete("/removeUserToken", userTokenHandler.RemoveUserFCMToken)

	notificationGroup := apiGroup.Group("/notification")
	notificationGroup.Post("/sendNotification", notificationHandler.SendNotification)

	userGroup := apiGroup.Group("/user")
	userGroup.Post("/createUser", userHandler.CreateUser)
	userGroup.Get("/:id", userHandler.GetUser)

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
