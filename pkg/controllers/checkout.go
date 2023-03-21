package controllers

import (
	"project/initializers"
	"project/pkg/function"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
)

func Checkout(g *gin.Context) {
	// var slice []string
	userid := function.GetUserId(g)
	//display address
	var address []response.Address
	initializers.DB.Raw("SELECT *FROM addresses WHERE user_id=?", userid).Scan(&address)
	// cartItems
	var order []models.ShoppingCartItem
	initializers.DB.Raw("SELECT *FROM shopping_cart_items WHERE cart_id=?", userid).Scan(&order)
	//grand total
	var total int
	initializers.DB.Raw("SELECT sum(total) FROM shopping_cart_items WHERE cart_id=?", userid).Scan(&total)
	println(total, "checkout total")
	initializers.DB.Raw("update carts SET total=? where id=?", total, userid)
	var grandtotal int
	initializers.DB.Raw("select total from carts where id=? ", userid).Scan(&grandtotal)

	//payment method
	var payment []models.Payment
	initializers.DB.Raw("SELECT *FROM payments").Scan(&payment)

	var balance int
	initializers.DB.Raw("select balance from walets where id=?", userid).Scan(&balance)

	println(grandtotal, "db grand total")

	display := response.Checkout{
		Message: "Displaying Checkout",
		Total:   grandtotal,
		Walet:   balance,
		Items:   order,
		Address: address,
		Payment: payment,
	}
	response.SurcessMessage(g, "Displaying Checkout", display)

}
