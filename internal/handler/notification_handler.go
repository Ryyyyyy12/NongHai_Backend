package handler

import (
	"context"
	"fmt"
	"log"

	"backend/internal/domain/response"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type NotificationHandler struct {
	FirebaseApp *firebase.App
}

// NewNotificationHandler initializes a new NotificationHandler with a Firebase App.
func NewNotificationHandler() NotificationHandler {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json") // Path to your service account key
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}

	return NotificationHandler{
		FirebaseApp: app,
	}
}

// SendNotification sends a notification using Firebase Cloud Messaging.
func (h NotificationHandler) SendNotification(c *fiber.Ctx) error {
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
	registrationToken := "ciDfKJoQSLKgiaZdK7I4Mg:APA91bHfj0NO0XFJO-LTdm68439ERlDLlbEtCK4oiV68bglAJDwQb66UDoVBF3Nwfch0ona3GZoBTxiGaZQC5iOLDcazp2nvl2hDytplFptOeoEcPJCtRXRyUXvnaedm6mMVRuaJfsw1"

	if registrationToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Data:    "Registration token is required",
		})
	}

	// Define a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "New Alert",
			Body:  "This is a test notification sent from the server.",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided registration token.
	responseID, err := client.Send(ctx, message)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Data:    "Failed to send notification",
		})
	}

	// Successfully sent message
	fmt.Printf("Successfully sent message: %s\n", responseID)
	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    fmt.Sprintf("Notification sent successfully, message ID: %s", responseID),
	})
}
