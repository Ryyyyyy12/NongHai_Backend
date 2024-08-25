package handler

import "backend/internal/service"

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userSvc service.IUserService) UserHandler {
	return UserHandler{
		userService: userSvc,
	}
}
