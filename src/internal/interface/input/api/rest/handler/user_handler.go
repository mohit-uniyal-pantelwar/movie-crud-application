package handler

import (
	"encoding/json"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/internal/usecase"
	"net/http"
)

type UserHandlerImpl interface {
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	userService usecase.UserServiceImpl
}

func NewUserHandler(userService usecase.UserServiceImpl) UserHandlerImpl {
	return UserHandler{
		userService: userService,
	}
}

func (uh UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	insertedUser, err := uh.userService.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertedUser)
}
