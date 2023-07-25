package common

import (
	"project/initializers"
	"project/response"

	"github.com/gin-gonic/gin"
)

func ViewProduct(g *gin.Context) {

	product := response.Product{}

	product_id := g.Query("id")
	initializers.DB.Raw("select *from products where id=?", product_id).Scan(&product)
	view := product.View + 1
	initializers.DB.Raw("update products set view=? where id=?", view, product_id).Scan(&product)

	response.SurcessMessage(g, "displaying selected product", product)

}
