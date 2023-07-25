package payment

import (
	"net/http"
	"project/initializers"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
)

func AddPaymentMethod(g *gin.Context) {

	var body struct {
		Paymentmethod string
	}
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}
	println("qwewqw", body.Paymentmethod)
	new := models.Payment{
		PaymentMethod: body.Paymentmethod,
	}
	result := initializers.DB.Create(&new)
	if result.Error != nil {
		response.ErrorMessage(g, "failed to add payment method", result.Error.Error(), 500)
		return
	}
	response.SurcessMessage(g, "payment added surcessfully", new)
}
