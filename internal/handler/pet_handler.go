package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/domain/response"
	"backend/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// PetHandler handles pet-related API requests
type PetHandler struct {
	petService service.IPetService
}

// NewPetHandler creates a new PetHandler instance
func NewPetHandler(petService service.IPetService) PetHandler {
	return PetHandler{
		petService: petService,
	}
}

// GetPet retrieves pet information based on ID and returns it with age
func (h PetHandler) GetPet(c *fiber.Ctx) error {
	petId := c.Params("id")
	if petId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Message: "Pet ID is required",
		})
	}

	// Retrieve pet information from the service
	pet, err := h.petService.GetPetInfo(petId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.InfoResponse{
			Success: false,
			Message: "Pet not found: " + err.Error(),
		})
	}

	// Respond with the pet data including age
	return c.Status(fiber.StatusOK).JSON(response.InfoResponse{
		Success: true,
		Data:    pet,
		Message: "Pet retrieved successfully",
	})
}

// CreatePet creates a new pet using the data from the request body
func (h PetHandler) CreatePet(c *fiber.Ctx) error {
	body := new(dto.CreatePetBody)
	if err := c.BodyParser(body); err != nil {
		// Log the error and request body for debugging
		fmt.Println("Error parsing body:", err)
		bodyStr := c.Body()
		fmt.Println("Received body:", string(bodyStr))

		return c.Status(fiber.StatusBadRequest).JSON(response.InfoResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Create model.Pet without age
	newPet := model.Pet{
		UserID:      body.UserID,
		Name:        body.Name,
		AnimalType:  body.AnimalType,
		Breed:       body.Breed,
		DateOfBirth: body.DateOfBirth.Time, // Extract time from CustomDate
		Sex:         body.Sex,
		Weight:      body.Weight,
		HairColor:   body.HairColor,
		BloodType:   body.BloodType,
		Eyes: body.Eyes,
		Status: body.Status,
		Note:        body.Note,
		Image:       body.Image,
	}

	// Call the service to create the pet
	createdPet, err := h.petService.Create(&newPet)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.InfoResponse{
			Success: false,
			Message: "Failed to create pet: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.InfoResponse{
		Success: true,
		Data:    createdPet,
		Message: "Pet created successfully",
	})
}
