package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/response"
	"backend/internal/service"
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type NotificationHandler struct {
	FirebaseApp         *firebase.App
	userFCMTokenService service.IUserFCMTokenService
}

// NewNotificationHandler initializes a new NotificationHandler with a Firebase App.
func NewNotificationHandler(userFCMTokenService service.IUserFCMTokenService,
) NotificationHandler {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json") // Path to your service account key
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}

	return NotificationHandler{
		userFCMTokenService: userFCMTokenService,
		FirebaseApp:         app,
	}
}

// SendNotification sends a notification using Firebase Cloud Messaging.
func (h NotificationHandler) SendNotification(c *fiber.Ctx) error {
	body := new(dto.SendNotificationBody)
	ctx := context.Background()

	// Obtain a messaging.Client from the initialized Firebase App.
	client, err := h.FirebaseApp.Messaging(ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Data:    "Failed to get Firebase Messaging client",
		})
	}

	// This registration token comes from the client FCM SDKs.

	registrationTokens, err := h.userFCMTokenService.GetUserFCMToken(body.UserID)
	if err != nil {
		log.Printf("error getting user FCM token: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Data:    "Failed to get user FCM token",
		})
	}

	if len(registrationTokens) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Data:    "No registered FCM tokens found",
		})
	}

	// Define a message payload.
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "New Alert",
			Body:  "This is a test notification sent to all users from the server.",
		},
		Tokens: registrationTokens, // Set the list of FCM tokens here
	}

	// Send a multicast message to all devices corresponding to the provided registration tokens.
	notiResponses, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Printf("Failed to send multicast message: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Data:    "Failed to send notification to all users",
		})
	}

	// Log and return the number of messages successfully sent
	successCount := notiResponses.SuccessCount
	failureCount := notiResponses.FailureCount
	fmt.Printf("Successfully sent message to %d devices, %d messages failed\n", successCount, failureCount)

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    fmt.Sprintf("Notification sent successfully to %d devices, %d failures", successCount, failureCount),
	})
}
