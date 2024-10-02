package dto

type CreateUserBody struct {
	ID        string  `json:"id" validate:"required"`
	Username  string  `json:"username" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Surname   string  `json:"surname" validate:"required"`
	Email     string  `json:"email" validate:"required,email"`
	Phone     string  `json:"phone" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Image     string  `json:"image"`
}

type UserInfoResponse struct {
	ID        string          `json:"id"`
	Username  string          `json:"username"`
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	Address   string          `json:"address"`
	Latitude  float64         `json:"latitude"`
	Longitude float64         `json:"longitude"`
	Image     string          `json:"image"`
	Pets      []CreatePetBody `json:"pets"`
}