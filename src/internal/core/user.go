package models

import (
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

type Session struct {
	Id        uuid.UUID
	UserId    int
	TokenHash string
	ExpiresAt time.Time
	IssuedAt  time.Time
}
