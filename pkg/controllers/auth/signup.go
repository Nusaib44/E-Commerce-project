package auth
import (
	"net/http"
	"project/initializers"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
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
