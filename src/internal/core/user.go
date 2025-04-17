package models

import (
	"movie-crud-application/src/internal/config"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	FoundUser   User
	TokenString string
	TokenExpire time.Time
	Session     Session
}

type UserRepoImpl interface {
	CreateUser(user User) (User, error)
	FindUserByUsername(username string) (User, error)
	FindUserById(userId int) (User, error)
}

type UserServiceImpl interface {
	CreateUser(user User) (User, error)
	LoginUser(user User, config *config.Config) (LoginResponse, error)
	GetUserById(userId int) (User, error)
	LogoutUser(userId int) error
	GetJWTFromSessionId(sessionId string) (string, time.Time, error)
}

type Session struct {
	Id        uuid.UUID
	UserId    int
	TokenHash string
	ExpiresAt time.Time
	IssuedAt  time.Time
}

type SessionRepoImpl interface {
	CreateSession(session Session) error
	GetSessionById(sessionId string) (Session, error)
	DeleteSession(id int) error
}
