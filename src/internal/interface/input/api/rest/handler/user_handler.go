package handler

import (
	"encoding/json"
	"movie-crud-application/src/internal/config"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/internal/usecase"
	"net/http"
)

type UserHandlerImpl interface {
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request, config *config.Config)
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

func (uh UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request, config *config.Config) {

	//1. Decode the request body
	var req models.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//2. Login user and get users and tokens

	loginResponse, err := uh.userService.LoginUser(req, config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//3. set cookies

	accessTokenCookie := http.Cookie{
		Name:     "accessToken",
		Value:    loginResponse.TokenString,
		Expires:  loginResponse.TokenExpire,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	sessionCookie := http.Cookie{
		Name:     "sessionCookie",
		Value:    loginResponse.Session.TokenHash,
		Expires:  loginResponse.Session.ExpiresAt,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &accessTokenCookie)
	http.SetCookie(w, &sessionCookie)

	//4. send the response back to client

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-user", loginResponse.FoundUser.Username)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully"})
}
