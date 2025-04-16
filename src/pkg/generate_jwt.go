package pkg

import (
	models "movie-crud-application/src/internal/core"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Uid int
	jwt.StandardClaims
}

func GenerateJWT(uid int, jwtKey string) (string, time.Time, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil

}

func GenerateSession(userId int) (models.Session, error) {
	tokenId := uuid.New()
	expiresAt := time.Now().Add(2 * time.Hour)
	issuedAt := time.Now()

	hashToken, err := bcrypt.GenerateFromPassword([]byte(tokenId.String()), bcrypt.DefaultCost)

	if err != nil {
		return models.Session{}, err
	}

	session := models.Session{
		Id:        tokenId,
		UserId:    userId,
		TokenHash: string(hashToken),
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
	}

	return session, nil
}
