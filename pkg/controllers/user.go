package controllers

import (
	"net/http"
	"project/initializers"
	"project/pkg/function"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserInfo(g *gin.Context) {

	var user models.Userdata
	var add []models.Address
	var order []models.Order
	userID := function.GetUserId(g)

	initializers.DB.Raw("SELECT *FROM orders WHERE user_id=?", userID).Scan(&order)
	initializers.DB.Raw("SELECT *FROM addresses WHERE user_id=?", userID).Scan(&add)
	initializers.DB.Raw("SELECT *FROM userdata WHERE user_id=?", userID).Scan(&user)

	userinfo := response.UserInfo{
		UserDetails: user,
		Address:     add,
		Order:       order,
	}
	response.SurcessMessage(g, "displaying user info", userinfo)
}

func EditUserInfo(g *gin.Context) {

	var user models.Userdata
	id := function.GetUserId(g)
	var Body struct {
		Username    string
		Email       string
		PhoneNumber string
	}

	if v := g.Bind(&Body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}
	initializers.DB.First(&user, id)
	if user.ID < 1 {
		response.ErrorMessage(g, "failed to get user", " user is not on database", 502)
		return
	}

	err := initializers.DB.Model(&user).Updates(models.Userdata{
		Username:    Body.Username,
		Email:       Body.Email,
		PhoneNumber: Body.PhoneNumber,
	})
	if err != nil {
		response.ErrorMessage(g, "can not update user", err, 500)
	}
	response.SurcessMessage(g, "userinfo updated", user)
}

func ChangePassword(g *gin.Context) {
	var user models.Userdata
	var Body struct{ pass string }
	println("vkyvcjgfhvc", Body.pass)
	id := function.GetUserId(g)
	pass := g.Query("password")

	if v := g.Bind(&Body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}

	if Body.pass != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.pass))
		if err != nil {
			response.ErrorMessage(g, "invalid password", err.Error(), 400)
		} else {
			hash, hash_err := bcrypt.GenerateFromPassword([]byte(pass), 10)
			if hash_err != nil {
				response.ErrorMessage(g, "failed to hash passsword", err, 500)
				return
			}
			initializers.DB.Raw("update userdata SET password=? WHERE id=?", hash, id).Scan(&user)
			response.SurcessMessage(g, "password  updated surcessfully", user)

		}
	}

}
