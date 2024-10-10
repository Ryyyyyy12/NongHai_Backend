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
	petRepo := repository.NewPetRepository(DB)
	chatRepo := repository.NewChatRepository(*DB)
	tokenRepo := repository.NewTokenRepository(*DB)
	notificationRepo := repository.NewNotificationRepository(*DB)

	//Services
	trackingService := service.NewTrackingService(trackingRepo, userRepo, petRepo)
	userService := service.NewUserService(userRepo)
	chatService := service.NewChatService(chatRepo)
	userFCMTokenService := service.NewUserFCMTokenService(tokenRepo)
	notificationService := service.NewNotificationService(userFCMTokenService, notificationRepo, petRepo)
	petService := service.NewPetService(petRepo)

	//Handlers
	trackingHandler := handler.NewTrackingHandler(trackingService, userService, notificationService, petService)
	ChatHandler := handler.NewChatHandler(chatService)
	userHandler := handler.NewUserHandler(userService)
	userTokenHandler := handler.NewUserTokenHandler(userFCMTokenService)
	notificationHandler := handler.NewNotificationHandler(notificationService)
	petHandler := handler.NewPetHandler(petService)

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
	trackingGroup.Get("/getTrackingById", trackingHandler.GetTrackingByID)

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
	// notificationGroup.Post("/createNotificationObject", notificationHandler.CreateNotificationObject)
	notificationGroup.Get("/getNotificationObject", notificationHandler.GetNotificationObject)
	notificationGroup.Post("/setNotificationRead", notificationHandler.SetNotificationRead)
	notificationGroup.Get("/getNotification", notificationHandler.GetNotificationObjectByNotiID)

	userGroup := apiGroup.Group("/user")
	userGroup.Post("/createUser", userHandler.CreateUser)
	userGroup.Get("/:id", userHandler.GetUser)

	petGroup := apiGroup.Group("/pet")
	petGroup.Post("/createPet", petHandler.CreatePet)
	petGroup.Get("/:id", petHandler.GetPet)

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
