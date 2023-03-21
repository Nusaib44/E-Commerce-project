package middleware

import (
	"fmt"
	"net/http"
	"os"
	"project/initializers"
	"project/pkg/models"
	"project/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(g *gin.Context) {

	// geting cookie
	tokenString, err := g.Cookie("coookie")

	if err != nil {
		response.ErrorMessage(g, "acess denied", err.Error(), 500)
		g.AbortWithStatus(500)
		return
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the expm
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			response.ErrorMessage(g, "loging to acess", "admin token expired", 500)
			return
		}
		// Find the user with token sub
		var user models.Userdata
		// initializers.DB.First(&user, claims["sub"])
		initializers.DB.First(&user, "email = ?", claims["sub"])

		if user.ID == 0 {
			response.ErrorMessage(g, "user not found", "user not found ", 404)
			g.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func AdminAuth(g *gin.Context) {
	// geting cookie
	tokenString, err := g.Cookie("admincoookie")
	if err != nil {
		response.ErrorMessage(g, "acess denied", err.Error(), http.StatusUnauthorized)
		g.AbortWithStatus(500)
		return
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the expm
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			response.ErrorMessage(g, "acess denied", "no admin token found", http.StatusUnauthorized)
			g.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user with token sub
		var user models.Userdata
		// initializers.DB.First(&user, claims["sub"])
		initializers.DB.First(&user, "email = ?", claims["sub"])

		if user.ID == 0 {
			response.ErrorMessage(g, "admin not found", "admin not found ", 404)
			g.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
