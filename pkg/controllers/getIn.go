package controllers

import (
	"fmt"
	"net/http"
	"os"
	"project/initializers"
	"project/pkg"
	"project/pkg/models"
	"project/response"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(g *gin.Context) {
	// data storage of user from body
	var body struct {
		Username    string
		Email       string
		PhoneNumber string
		Password    string
		Status      bool
		Isadmin     bool
	}
	var err []string

	// binding json to go response writter
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}
	flag := 0
	// checking the user
	var user models.Userdata

	initializers.DB.First(&user, "email = ?", body.Email)
	println(user.Email, "login mail")
	println(user.ID)
	if user.ID > 0 {
		flag = 1
		err = append(err, "email alreaddy exist try another email")
	}
	// checking phone number
	println(" phone number", body.PhoneNumber)
	if len(body.PhoneNumber) != 10 {
		flag = 1
		err = append(err, "invalid phone number")
	}
	// checking phonenumber already exist
	var phone models.Userdata
	initializers.DB.First(&phone, "phone_number = ?", body.PhoneNumber)
	if phone.ID > 0 {
		flag = 1
		err = append(err, "phone number alreaddy exist try another number or login with otp")
	}
	if flag == 1 {
		response.ErrorMessage(g, "login failed", err, 404)
		return
	}
	// Hashing the password
	hash, hash_err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if hash_err != nil {
		response.ErrorMessage(g, "failed to hash passsword", hash_err.Error(), 500)
		return
	}

	println("hash", hash)

	// create a user with mail pass
	sign_up_user := models.Userdata{
		Username:    body.Username,
		Email:       body.Email,
		Password:    string(hash),
		PhoneNumber: body.PhoneNumber,
		Status:      body.Status,
		Isadmin:     body.Isadmin,
	}
	result := initializers.DB.Create(&sign_up_user)

	if result.Error != nil {
		response.ErrorMessage(g, "failed to create user", result.Error.Error(), 500)
		return
	}
	// respond
	response.SurcessMessage(g, "signed up surcessfully", sign_up_user)

	var userid int
	initializers.DB.Raw("SELECT ID FROM userdata WHERE email=?", body.Email).Scan(&userid)

	cart := models.Cart{UserId: userid}
	walet := models.Walet{UserId: userid}
	result2 := initializers.DB.Create(&cart)
	result3 := initializers.DB.Create(&walet)
	if result2.Error != nil {
		response.ErrorMessage(g, "failed to create cart", result2.Error.Error(), 500)
		return
	}

	if result3.Error != nil {
		response.ErrorMessage(g, "failed to create walet", result3.Error.Error(), 500)
		return
	}
}

func Login(g *gin.Context) {

	// data storage of user from body
	var body struct {
		Email    string
		Password string
	}
	// binding json to go response writter
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}
	// checking the user
	var user models.Userdata

	initializers.DB.First(&user, "email = ?", body.Email)
	println(user.Email, "login mail")
	if user.ID == 0 {
		response.ErrorMessage(g, "invalid email check email or try signup", "", 400)
		return
	}

	// compairing password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		response.ErrorMessage(g, "invalid password", err.Error(), 400)
		return
	}

	// check block
	if !user.Status {
		response.ErrorMessage(g, "u are blocked", err.Error(), 400)
		return
	}

	// for user____________________________________________________
	if !user.Isadmin && user.Status {

		usertoken(user.Email, g)

	} else if user.Status && user.Isadmin {

		// for admin____________________________________________________
		admintoken(user.Email, g)

	} else {
		g.JSON(http.StatusBadRequest, gin.H{"error": "You are blocked"})
	}

}

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
