package auth

import (
	"net/http"
	"project/initializers"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
