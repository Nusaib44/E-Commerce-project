package controllers

import (
	"net/http"
	"project/initializers"
	"project/pkg/models"
	"project/response"

	"github.com/gin-gonic/gin"
)

func AddProductOffer(g *gin.Context) {
	var body struct {
		Product int
		Offer   int
	}
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}

	// add offer
	productOffer := models.ProductOffers{
		Product:    body.Product,
		OfferPrice: body.Offer,
	}
	OfferResult := initializers.DB.Create(&productOffer)
	if OfferResult.Error != nil {
		response.ErrorMessage(g, "failed to add Offer", OfferResult.Error.Error(), 500)
		return
	}

	var price int
	var products models.Product
	initializers.DB.Raw("SELECT total from products where id=?", body.Product).Scan(&price)
	newprice := price - body.Offer
	println(price, "price")
	println("new ", newprice)
	initializers.DB.Raw("update products set total=?,product_offer=? where id=?", newprice, body.Offer, body.Product).Scan(&products)
	response.SurcessMessage(g, "product order added", products)
}

func AddCategoryOffer(g *gin.Context) {

	var body struct {
		Category int
		Offer    int
	}
	if v := g.Bind(&body); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}

	Offer := models.CategoryOffers{
		Category:   body.Category,
		OfferPrice: body.Offer,
	}
	OfferResult := initializers.DB.Create(&Offer)
	if OfferResult.Error != nil {
		response.ErrorMessage(g, "failed to add Offer", OfferResult.Error.Error(), 500)
		return
	}
	var id []int
	initializers.DB.Raw("SELECT id from products where category=?", body.Category).Scan(&id)
	for _, v := range id {
		var price int
		var products models.Product
		initializers.DB.Raw("SELECT total from products where id=?", v).Scan(&price)
		newprice := price - ((price * body.Offer) / 100)
		initializers.DB.Raw("update products set total=?,category_offer=? where id=?", newprice, body.Offer, v).Scan(&products)
		response.SurcessMessage(g, "category offer added", products)
	}

}
