package middleware

import (
	"log"
	"mikhael-project-go/internal/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Ini Lewat Headers
		token := c.GetHeader("token")

		log.Println("token : ", token)

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		log.Println("Token new : ", token)
		// todo -> cek JWT
		res, err := jwt.ParseWithClaims(token, &utils.Claims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		log.Println("Midleware : ", res)
		if err != nil || !res.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": err.Error()})
			c.Abort()
			return
		}

		c.Next()

	}
}
