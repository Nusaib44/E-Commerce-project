package controllers

import (
	"project/initializers"
	"project/pkg/models"
	"project/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DisplayProduct(g *gin.Context) {

	// product listing with pagination , we can input page number to list products in different pages
	pagestring := g.Query("page")
	page, _ := strconv.Atoi(pagestring)
	offset := (page - 1) * 3
	var product []models.Product
	initializers.DB.Limit(3).Offset(offset).Find(&product)
	response.SurcessMessage(g, "listing products", product)

}

func DisplayCategory(g *gin.Context) {
	var displayCategory []models.Category
	initializers.DB.Find(&displayCategory)
	response.SurcessMessage(g, "displaying Category", displayCategory)
}
