package function

import (
	"fmt"
	"os"
	"project/initializers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetUserId(g *gin.Context) int {
	tokenString, _ := g.Cookie("coookie")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	claim := token.Claims.(jwt.MapClaims)
	GetMail := claim["sub"]
	ToStr := GetMail.(string)

	var userID int
	initializers.DB.Raw("SELECT id FROM userdata WHERE email=?", ToStr).Scan(&userID)

	return userID
}

func ProductQuantity(ProductItemID int) int {
	var currentQuantity int
	initializers.DB.Raw("SELECT quantity FROM products WHERE id=?", ProductItemID).Scan(&currentQuantity)
	println(currentQuantity)
	return currentQuantity
}

// func RedeemCoupon(code string) {

// 	var coupen models.Coupen
// 	initializers.DB.Raw("SELECT *FROM coupens where code=?", code).Scan(&coupen)
// 	if coupen.ID == 0 {
// 		g.JSON(400, gin.H{"error": "enter valid coupen code"})
// 	}
// 	if !coupen.IsUsed && coupen.Limit < int64(total) {
// 		total = total - coupen.Value
// 	}
// }
