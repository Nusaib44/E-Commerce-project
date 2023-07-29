package common

import (
	"fmt"
	"net/http"
	"os"
	"project/pkg"
	"project/pkg/models"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var TWILIO_ACCOUNT_SID string = pkg.TWILIO_ACCOUNT_SID
var TWILIO_AUTH_TOKEN string = pkg.TWILIO_AUTH_TOKEN
var VERIFY_SERVICE_SID string = pkg.VERIFY_SERVICE_SID
var to string
var plus string

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TWILIO_ACCOUNT_SID,
	Password: TWILIO_AUTH_TOKEN})

func SendOtp(g *gin.Context) {
	g.JSON(200, TWILIO_ACCOUNT_SID)
	g.JSON(200, TWILIO_AUTH_TOKEN)
	g.JSON(200, VERIFY_SERVICE_SID)

	to = g.Query("phone")
	plus = "+91" + to
	params := &openapi.CreateVerificationParams{}
	params.SetTo(plus)
	params.SetChannel("whatsapp")

	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		g.JSON(http.StatusOK, gin.H{"": "OTP send"})
	}
}

func Checkotp(g *gin.Context) {
	var user models.Userdata

	fmt.Println("Please check your phone and enter the code:")
	code := g.Query("code")

	params := &openapi.CreateVerificationCheckParams{}
	fmt.Println(plus)
	params.SetTo(plus)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else if *resp.Status == "approved" {
		g.JSON(http.StatusOK, gin.H{"": "OTP Verified"})
		// create tocken

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.Email,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "Can't creare token"})
		}

		// sending token
		g.SetSameSite(http.SameSiteLaxMode)
		g.SetCookie("coookie", tokenString, 3600*24*30, "", "", false, true)

		g.JSON(http.StatusOK, gin.H{"message": "login surcessful"})

	} else {
		fmt.Println("Incorrect!")
	}
}
