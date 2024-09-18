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