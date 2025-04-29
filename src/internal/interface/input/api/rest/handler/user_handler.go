package handler

import (
	"encoding/json"
	"movie-crud-application/src/internal/config"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/pkg"
	"net/http"
	"time"
)

type UserHandler struct {
	config      *config.Config
	userService models.UserServiceImpl
}

func NewUserHandler(config *config.Config, userService models.UserServiceImpl) UserHandler {
	return UserHandler{
		config:      config,
		userService: userService,
	}
}

func (uh UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusBadRequest,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	insertedUser, err := uh.userService.CreateUser(user)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Data: insertedUser,
	}
	response.Set()
}

func (uh UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	//1. Decode the request body
	var req models.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusBadRequest,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	//2. Login user and get users and tokens

	loginResponse, err := uh.userService.LoginUser(req, uh.config)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
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
		Value:    loginResponse.Session.Id.String(),
		Expires:  loginResponse.Session.ExpiresAt,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &accessTokenCookie)
	http.SetCookie(w, &sessionCookie)

	//4. send the response back to client

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"x-user":       loginResponse.FoundUser.Username,
		},
		Message: "Logged in successfully",
	}
	response.Set()
}

func (uh UserHandler) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	//1. get user id from context

	userId, ok := r.Context().Value("user").(int)
	if !ok {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusUnauthorized,
			Error:          "user not found in context",
		}
		response.Set()
		return
	}

	//2. fetch user profile from id

	user, err := uh.userService.GetUserById(userId)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	//3. return the response

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"x-user":       user.Username,
		},
		Data: user,
	}
	response.Set()

}

func (uh UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//1. get user id from context

	userId, ok := r.Context().Value("user").(int)
	if !ok {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusUnauthorized,
			Error:          "user not found in context",
		}
		response.Set()
		return
	}

	//2. logout user

	err := uh.userService.LogoutUser(userId)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	//3. set cookies

	accessTokenCookie := http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Now(),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	sessionCookie := http.Cookie{
		Name:     "sessionCookie",
		Value:    "",
		Expires:  time.Now(),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &accessTokenCookie)
	http.SetCookie(w, &sessionCookie)

	//4. return the response

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Message: "Logged out successfully",
	}
	response.Set()
}

func (uh UserHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	//1. Get the session cookie

	cookie, err := r.Cookie("sessionCookie")
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusUnauthorized,
			Error:          "user not found in context",
		}
		response.Set()
		return
	}

	//2. Get the new token from session cookie, if invalid return the error

	token, expirationTime, err := uh.userService.GetJWTFromSessionId(cookie.Value)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	//3. set the new token as cookie

	accessTokenCookie := http.Cookie{
		Name:     "accessToken",
		Value:    token,
		Expires:  expirationTime,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &accessTokenCookie)

	//4. return the response

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Message: "token refreshed successfully",
	}
	response.Set()
}
