package usecase

import (
	"errors"
	"fmt"
	"movie-crud-application/src/internal/config"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/pkg"
	"time"
)

type UserService struct {
	userRepo    models.UserRepoImpl
	sessionRepo models.SessionRepoImpl
	jwtKey      string
}

func NewUserService(userRepo models.UserRepoImpl, sessionRepo models.SessionRepoImpl, jwtKey string) models.UserServiceImpl {
	return UserService{userRepo: userRepo, sessionRepo: sessionRepo, jwtKey: jwtKey}
}

func (us UserService) CreateUser(user models.User) (models.User, error) {
	createdUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return createdUser, errors.New("failed to create user, try again later")
	}

	return createdUser, nil
}

func (us UserService) LoginUser(user models.User, config *config.Config) (models.LoginResponse, error) {
	loginResponse := models.LoginResponse{}

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

func (us UserService) GetUserById(userId int) (models.User, error) {
	user, err := us.userRepo.FindUserById(userId)
	if err != nil {
		return user, fmt.Errorf("cannot get user, try again later")
	}

	return user, nil
}

func (us UserService) LogoutUser(userId int) error {
	err := us.sessionRepo.DeleteSession(userId)
	if err != nil {
		return fmt.Errorf("failed to logout user, please try again later")
	}

	return nil
}

func (us UserService) GetJWTFromSessionId(sessionId string) (string, time.Time, error) {
	//1. Get session by session ID

	session, err := us.sessionRepo.GetSessionById(sessionId)
	if err != nil {
		return "", time.Now(), fmt.Errorf("error fetching session")
	}

	//2. Check if session is not expired

	if time.Now().After(session.ExpiresAt) {
		return "", time.Now(), fmt.Errorf("session expired, login again")
	}

	//3. generate new JWT token

	token, expirationTime, err := pkg.GenerateJWT(session.UserId, us.jwtKey)
	if err != nil {
		return token, expirationTime, fmt.Errorf("error generating token")
	}

	//4 return the result

	return token, expirationTime, nil
}
