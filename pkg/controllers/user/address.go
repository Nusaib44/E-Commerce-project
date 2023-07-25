package user

import (
	"net/http"
	"project/initializers"
	"project/pkg/function"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
)

func AddNewAddress(g *gin.Context) {
	userid := function.GetUserId(g)
	var body struct {
		HouseName   string
		Street      string
		AddressLine string
		City        string
		State       string
		Pincode     int
		Country     string
		IsDefault   bool
	}
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}

	new := models.Address{
		UserId:      userid,
		HouseName:   body.HouseName,
		Street:      body.Street,
		AddressLine: body.AddressLine,
		City:        body.City,
		State:       body.State,
		Pincode:     body.Pincode,
		Country:     body.State,
		IsDefault:   body.IsDefault,
	}
	result := initializers.DB.Create(&new)
	if result.Error != nil {
		response.ErrorMessage(g, "failed to add address... try again", result.Error.Error(), 502)
		return
	}
	response.SurcessMessage(g, "address added surcessfully", new)
}
