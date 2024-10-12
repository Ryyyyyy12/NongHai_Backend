package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/domain/response"
	"backend/internal/service"
	"backend/internal/util/text"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UserHandler handles user-related API requests
type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService service.IUserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

// GetUser retrieves user information based on ID, including pets
func (h UserHandler) GetUser(c *fiber.Ctx) error {
	// Get the user ID from the URL parameters
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Message: "User ID is required",
		})
	}

	// Retrieve user information from the service, along with associated pets
	user, err := h.userService.GetUserInfo(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.InfoResponse{
			Success: false,
			Message: "User not found: " + err.Error(),
		})
	}

	// Respond with the user data
	return c.Status(fiber.StatusOK).JSON(response.InfoResponse{
		Success: true,
		Data:    user,
		Message: "User retrieved successfully",
	})
}

// CreateUser creates a new user using the data from the request body
func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	// Parse request body into CreateUserBody DTO
	body := new(dto.CreateUserBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Validate the request body fields
	if err := text.Validator.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.InfoResponse{
			Success: false,
			Message: "Validation failed: " + err.Error(),
		})
	}

	// Use the Firebase UID (passed from Flutter) as the user ID
	newUser := model.User{
		ID:        body.ID,        // Firebase UID
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  body.Username,
		Name:      body.Name,
		Surname:   body.Surname,
		Email:     body.Email,
		Phone:     body.Phone,
		Address:   body.Address,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
		Image:     body.Image,
	}

	// Call the service to create the user
	createdUser, err := h.userService.Create(&newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Message: "Failed to create user: " + err.Error(),
		})
	}

	// Respond with the created user data
	return c.Status(fiber.StatusCreated).JSON(response.InfoResponse{
		Success: true,
		Data:    createdUser,
		Message: "User created successfully",
	})
}

// UpdateUser updates an existing user based on the request body
func (h UserHandler) UpdateUser(c *fiber.Ctx) error {
	// Get the user ID from the URL parameters
	log.Printf("Request Method: %s, URL: %s", c.Method(), c.OriginalURL())
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Message: "User ID is required",
		})
	}

	// Parse request body into UpdateUserBody DTO
	body := new(dto.UpdateUserBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Call the service to update the user
	updatedUser, err := h.userService.Update(userId, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Message: "Failed to update user: " + err.Error(),
		})
	}

	// Respond with the updated user data
	return c.Status(fiber.StatusOK).JSON(response.InfoResponse{
		Success: true,
		Data:    updatedUser,
		Message: "User updated successfully",
	})
}

