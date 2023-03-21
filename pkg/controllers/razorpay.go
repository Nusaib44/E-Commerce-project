package controllers

import (
	"project/initializers"
	"project/pkg/function"
	// "project/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

func Razorpay(g *gin.Context) {

	userid := function.GetUserId(g)
	var sumtotal int
	// var user models.Userdata
	initializers.DB.Raw("SELECT sum(total) FROM shopping_cart_items where cart_id=?", userid).Scan(&sumtotal)
	razpayvalue := sumtotal / 100

	client := razorpay.NewClient("YOUR_KEY_ID", "YOUR_SECRET")
	data := map[string]interface{}{
		"amount":   razpayvalue,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	value := body["id"]

	if err != nil {
		g.JSON(404, gin.H{
			"msg": "error creating order",
		})
	}
	g.HTML(200, "app.html", gin.H{

		"UserID":       userid,
		"total_price":  sumtotal,
		"total":        razpayvalue,
		"orderid":      value,
		"amount":       sumtotal,
	})

}
