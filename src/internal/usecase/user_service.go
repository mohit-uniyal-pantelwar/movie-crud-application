package usecase

import (
	"errors"
	"fmt"
	"movie-crud-application/src/internal/adapters/persistance"
	"movie-crud-application/src/internal/config"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/pkg"
	"time"
)

type UserServiceImpl interface {
	CreateUser(user models.User) (models.User, error)
	LoginUser(user models.User, config *config.Config) (LoginResponse, error)
}

type UserService struct {
	userRepo    persistance.UserRepoImpl
	sessionRepo persistance.SessionRepoImpl
}

func NewUserService(userRepo persistance.UserRepoImpl, sessionRepo persistance.SessionRepoImpl) UserServiceImpl {
	return UserService{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (us UserService) CreateUser(user models.User) (models.User, error) {
	createdUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return createdUser, errors.New("failed to create user, try again later")
	}

	return createdUser, nil
}

type LoginResponse struct {
	FoundUser   models.User
	TokenString string
	TokenExpire time.Time
	Session     models.Session
}

func (us UserService) LoginUser(user models.User, config *config.Config) (LoginResponse, error) {
	loginResponse := LoginResponse{}

	//1. find user by username and verify the password

	foundUser, err := us.userRepo.FindUserByUsername(user.Username)
	if err != nil {
		return loginResponse, fmt.Errorf("invalid username")
	}

	if err := pkg.CheckPassword(foundUser.Password, user.Password); err != nil {
		return loginResponse, fmt.Errorf("invalid passsword")
	}

	//2. Create token

	tokenString, tokenExpiration, err := pkg.GenerateJWT(foundUser.Id, config.JWT_SECRET)
	if err != nil {
		return loginResponse, fmt.Errorf("failed to generate jwt")
	}

	//3. Create session

	session, err := pkg.GenerateSession(foundUser.Id)
	if err != nil {
		return loginResponse, fmt.Errorf("failed to generate session")
	}

	err = us.sessionRepo.CreateSession(session)
	if err != nil {
		return loginResponse, fmt.Errorf("failed to create session")
	}

	//4. return the login response
	loginResponse.FoundUser = foundUser
	loginResponse.TokenString = tokenString
	loginResponse.TokenExpire = tokenExpiration
	loginResponse.Session = session

	return loginResponse, nil
}
