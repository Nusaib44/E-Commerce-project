package user

import (
	"project/initializers"
	"project/pkg/function"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
)

func ValidateCoupen(g *gin.Context) {

	userid := function.GetUserId(g)
	// coupen check
	coupen_code := g.Query("code")
	println("cccc", coupen_code)
	var coupen models.Coupen
	initializers.DB.Raw("SELECT *FROM coupens where code=?", coupen_code).Scan(&coupen)
	if len(coupen_code) < 5 || coupen.ID == 0 {
		println("1")
		response.ErrorMessage(g, "enter valid coupon code", "length of the coupon is kless than 5", 400)
		return
	}
	if coupen.IsUsed {
		println("2")
		response.ErrorMessage(g, "enter valid coupon code", "coupon  already used", 400)
		return
	}
	var total int
	println("3")
	initializers.DB.Raw("select total from carts where id=? ", userid).Scan(&total)
	if !coupen.IsUsed && coupen.Limit < int64(total) {
		var walet models.Walet
		initializers.DB.Raw("select balance from walets where id=? ", userid).Scan(&walet)
		walet.Balance += coupen.Value
		initializers.DB.Raw("update walets set balance=? where id=?", walet.Balance, userid).Scan(&walet)

		initializers.DB.Raw("UPDATE coupens SET is_used=? WHERE id=?", true, coupen.ID).Scan(&coupen)
		response.SurcessMessage(g, "coupen added surcessfully", coupen)
	}

}
