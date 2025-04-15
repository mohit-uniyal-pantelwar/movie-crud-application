package handler

import "movie-crud-application/src/internal/usecase"

type UserHandler struct {
	userService usecase.UserService
}

func NewUserHandler(userService usecase.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
