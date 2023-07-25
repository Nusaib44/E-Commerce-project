package admin

import (
	"net/http"
	"project/initializers"
	"project/pkg/models"
	"project/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddCoupen(g *gin.Context) {
	var body struct {
		Code   string
		Value  int
		Limit  int64
		IsUsed bool
	}

	// binding json to go response writter
	if v := g.Bind(&body); v != nil {
		println("errrrr......not binded")
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}
	coupen := models.Coupen{
		Code:   body.Code,
		Value:  body.Value,
		Limit:  body.Limit,
		IsUsed: false,
		Expire: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	add_coupen := initializers.DB.Create(&coupen)

	if add_coupen.Error != nil {
		response.ErrorMessage(g, "failed to add coupon", add_coupen.Error.Error(), 500)
		return
	}
}

func ListCoupon(g *gin.Context) {
	pagestring := g.Query("page")
	page, _ := strconv.Atoi(pagestring)
	offset := (page - 1) * 3
	var coupon []models.Coupen
	initializers.DB.Limit(3).Offset(offset).Find(&coupon)
	response.SurcessMessage(g, "listing coupon", coupon)
}

func EditCoupon(g *gin.Context) {

	params := g.Query("id")
	id, _ := strconv.Atoi(params)
	var coupon models.Coupen
	initializers.DB.First(&coupon, id)

	if coupon.ID < 1 {
		response.ErrorMessage(g, "enter valid coupon code", "Coupon not found", 400)
		return
	}

	var body struct {
		Code   string
		Value  int
		Limit  int64
		IsUsed bool
	}
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}

}
