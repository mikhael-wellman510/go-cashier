package utils

import (
	"mikhael-project-go/internal/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId *entities.User) (string, error) {

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	claims := Claims{
		UserId:   userId.ID,
		Username: userId.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 24 Jam
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),

			// todo -> 20 detik
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 20)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
