package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/response"
	"backend/internal/service"
	"backend/internal/util/text"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type NotificationHandler struct {
	notificationService service.INotificationService
}

// NewNotificationHandler initializes a new NotificationHandler with a Firebase App.
func NewNotificationHandler(
	notificationService service.INotificationService,

) NotificationHandler {
	return NotificationHandler{
		notificationService: notificationService,
	}
}

// SendNotification sends a notification using Firebase Cloud Messaging.
func (h NotificationHandler) SendNotification(c *fiber.Ctx) error {
	body := new(dto.SendNotificationBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	if err := text.Validator.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	successCount, failureCount, err := h.notificationService.SendNotification(*body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    fmt.Sprintf("Notification sent successfully to %d devices, %d failures", successCount, failureCount),
	})
}
