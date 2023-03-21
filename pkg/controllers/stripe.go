package controllers

import (
	"net/http"
	"os"
	"strings"

	"project/initializers"
	"project/pkg/function"
	"project/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func StripePayment(c *gin.Context) {
	userid := function.GetUserId(c)

	var payment models.Charge
	c.BindJSON(&payment)

	var total int64
	var walet int
	initializers.DB.Raw("SELECT sum(total) FROM shopping_cart_items WHERE cart_id=?", userid).Scan(&total)
	initializers.DB.Raw("select walet from userdata where ID=?", userid).Scan(&walet)
	grandtotal := total - int64(walet)
	c.JSON(http.StatusOK, gin.H{"Total  amount paid": grandtotal})

	var product []string
	initializers.DB.Raw("select product_name from shopping_cart_items where cart_id=?", userid).Scan(&product)

	var mail string
	initializers.DB.Raw("select email from userdata where id=?", userid).Scan(&mail)
	println(total)
	println(mail)

	descript := strings.Join(product, ",")
	apiKey := os.Getenv("STRIPE_SECRET_KEY")
	stripe.Key = apiKey
	println(descript)

	_, err := charge.New(&stripe.ChargeParams{
		Amount:       &grandtotal,
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  &descript,
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: &mail,
	})

	if err != nil {
		c.String(http.StatusBadRequest, "Payment Unsuccessfull")
		return
	}

	err1 := SavePayment(&payment)
	if err1 != nil {
		c.String(http.StatusNotFound, "error occured")
	} else {
		c.JSON(200, gin.H{
			"message": "payment succesfull",
		})
	}

}

func SavePayment(charge *models.Charge) (err error) {
	if err = initializers.DB.Create(charge).Error; err != nil {
		return err
	}
	return nil

}
