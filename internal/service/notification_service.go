package service

import (
	"backend/internal/domain/dto"
	"context"
	"errors"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type INotificationService interface {
	SendNotification(notificationData dto.SendNotificationBody) (int, int, error)
}

type notificationService struct {
	FirebaseApp         *firebase.App
	userFCMTokenService IUserFCMTokenService
}

func NewNotificationService(
	userFCMTokenService IUserFCMTokenService,

) INotificationService {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json") // Path to your service account key
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}
	return &notificationService{
		userFCMTokenService: userFCMTokenService,
		FirebaseApp:         app,
	}
}

func (s *notificationService) SendNotification(notiData dto.SendNotificationBody) (int, int, error) {
	ctx := context.Background()

	// Obtain a messaging.Client from the initialized Firebase App.
	client, err := s.FirebaseApp.Messaging(ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
		return 0, 0, errors.New("failed to get Firebase Messaging client")
	}

	// This registration token comes from the client FCM SDKs.

	registrationTokens, err := s.userFCMTokenService.GetUserFCMToken(notiData.UserID)
	if err != nil {
		log.Printf("error getting user FCM token: %v\n", err)
		return 0, 0, errors.New("failed to get user FCM token")
	}

	if len(registrationTokens) == 0 {
		return 0, 0, errors.New("no registration tokens found")
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
		return 0, 0, errors.New("failed to send multicast message")
	}

	// Log and return the number of messages successfully sent
	successCount := notiResponses.SuccessCount
	failureCount := notiResponses.FailureCount
	fmt.Printf("Successfully sent message to %d devices, %d messages failed\n", successCount, failureCount)

	return successCount, failureCount, nil

}
