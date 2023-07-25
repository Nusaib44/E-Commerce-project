package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func usertoken(email string, g *gin.Context) {
	// create tocken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Can't creare token"})
	}

	// sending token
	g.SetSameSite(http.SameSiteLaxMode)
	g.SetCookie("coookie", tokenString, 36000, "", "", false, true)

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Can't creare token"})
	}
	g.JSON(http.StatusOK, gin.H{
		"message":       "user logged in surcessfully",
		"user token":    tokenString,
		"refresh token": rt,
	})
}

func admintoken(email string, g *gin.Context) {
	// create tocken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Can't creare token"})
	}

	// sending token
	g.SetSameSite(http.SameSiteLaxMode)
	g.SetCookie("admincoookie", tokenString, 36000, "", "", false, true)

	g.JSON(http.StatusOK, gin.H{
		"":           "admin logged in surcessfully",
		"admintoken": tokenString,
	})

}

// func Refreshtoken(c *gin.Context) {

// 	userid := function.GetUserId(c)

// 	type tokenReqBody struct {
// 		RefreshToken string `json:"refresh_token"`
// 	}
// 	tokenReq := tokenReqBody{}
// 	c.Bind(&tokenReq)

// 	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte("secret"), nil
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't creare token"})
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// Get the user record from database or
// 		var mail string
// 		initializers.DB.Raw("select email fron userdata where id=?", userid).Scan(&mail)
// 		// run through your business logic to verify if the user can log in

// 		if claims["sub"] == mail {

// 			println("token created")
// 			usertoken(mail, c)
// 		}
// 	}
// }
