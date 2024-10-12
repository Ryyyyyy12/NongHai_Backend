package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/repository"
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
	CreateNotificationObject(notiData dto.CreateNotificationObjectBody) error
	GetNotificationObject(userID string) ([]*model.Notification, error)
	ReadNotificationObject(notiID string) error
	GetNotificationObjectByNotiID(notiID string) (*model.Notification, error)
}

type notificationService struct {
	FirebaseApp         *firebase.App
	userFCMTokenService IUserFCMTokenService
	notiRepo            repository.INotificationRepository
	petRepo             repository.IPetRepository
}

func NewNotificationService(
	userFCMTokenService IUserFCMTokenService,
	notiRepo repository.INotificationRepository,
	petRepo repository.IPetRepository,

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
		notiRepo:            notiRepo,
		petRepo:             petRepo,
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

	registrationTokens, err := s.userFCMTokenService.GetUserFCMToken(*notiData.SentTO)
	if err != nil {
		log.Printf("error getting user FCM token: %v\n", err)
		return 0, 0, errors.New("failed to get user FCM token")
	}

	if len(registrationTokens) == 0 {
		return 0, 0, errors.New("no registration tokens found")
	}

	// Define a message payload.
	// message := &messaging.MulticastMessage{
	// 	Notification: &messaging.Notification{
	// 		Title: "New Alert",
	// 		Body:  "This is a test notification sent to all users from the server.",
	// 	},
	// 	Tokens: registrationTokens, // Set the list of FCM tokens here
	// }

	// // Send a multicast message to all devices corresponding to the provided registration tokens.
	// notiResponses, err := client.SendMulticast(ctx, message)
	// if err != nil {
	// 	log.Printf("Failed to send multicast message: %v\n", err)
	// 	return 0, 0, errors.New("failed to send multicast message")
	// }

	// // Log and return the number of messages successfully sent
	// successCount := notiResponses.SuccessCount
	// failureCount := notiResponses.FailureCount
	// fmt.Printf("Successfully sent message to %d devices, %d messages failed\n", successCount, failureCount)

	// return successCount, failureCount, nil

	// Temp Fix muliticast message Mmulticast discontinued
	successCount := 0
	failureCount := 0
	for _, token := range registrationTokens {
		fmt.Println()
		fmt.Println("senting to Token: ", token)
		fmt.Println()
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: *notiData.Title,
				Body:  *notiData.Body,
			},
			Data: map[string]string{
				"navigate_to": *notiData.NotificationData.Navigateto,
				"chat_with":   *notiData.NotificationData.ChatWith,
			},
			Token: token,
		}
		// Send a message to the device corresponding to the provided registration token.
		responseID, err := client.Send(ctx, message)
		if err != nil {
			log.Printf("Failed to send message: %v\n", err)
			failureCount++
			return 0, 0, errors.New("failed to send message")
		}
		// Successfully sent message
		fmt.Printf("Successfully sent message: %s\n", responseID)
		successCount++
	}
	return successCount, failureCount, nil

}

func (s *notificationService) CreateNotificationObject(notiData dto.CreateNotificationObjectBody) error {
	// Check if pet exists
	petData, err := s.petRepo.FindById(*notiData.PetID)
	if err != nil {
		return err
	}

	if err := s.notiRepo.CreateNotificationObject(model.Notification{
		UserID:     petData.UserID,
		PetID:      *notiData.PetID,
		TrackingID: notiData.TrackingID,
		IsRead:     false,
	}); err != nil {
		return err
	}
	return nil
}

func (s *notificationService) GetNotificationObject(userID string) ([]*model.Notification, error) {
	notiObject, err := s.notiRepo.GetNotificationObjectByUserID(userID)
	if err != nil {
		return nil, err
	}
	return notiObject, nil
}

func (s *notificationService) ReadNotificationObject(notiID string) error {
	notiObject, err := s.notiRepo.GetNotificationObjectByNotiID(notiID)
	if err != nil {
		return err
	}
	notiObject.IsRead = true
	if err := s.notiRepo.UpdateNotificationObject(*notiObject); err != nil {
		return err
	}

	return nil
}

func (s *notificationService) GetNotificationObjectByNotiID(notiID string) (*model.Notification, error) {
	notiObject, err := s.notiRepo.GetNotificationObjectByNotiID(notiID)
	if err != nil {
		return nil, err
	}
	return notiObject, nil
}
