package routes

import (
	"project/middleware"
	"project/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoute(r *gin.Engine) {

	// <--------------------------------USER MANAGMENT-------------------------------->
	r.GET("/listuser", middleware.AdminAuth, controllers.ListUser)
	r.PATCH("/block/:id", middleware.AdminAuth, controllers.BlockUser)
	r.PATCH("/unblock/:id", middleware.AdminAuth, controllers.Unblock)

	// <--------------------------------PRODUCT MANAGMENT<-------------------------------->
	r.POST("/addproduct", middleware.AdminAuth, controllers.AddProduct)
	r.PATCH("/editproduct", middleware.AdminAuth, controllers.EditProduct)
	r.DELETE("/deleteproduct", middleware.AdminAuth, controllers.DeleteProduct)

	// <--------------------------------CATEGORY MANAGMENT<-------------------------------->
	r.POST("/addcategory", middleware.AdminAuth, controllers.AddCategory)
	r.PATCH("/editcategory", middleware.AdminAuth, controllers.EditCategory)
	r.PUT("/deletecategory", middleware.AdminAuth, controllers.DeleteCategory)

	r.POST("/addpayment", middleware.AdminAuth, controllers.AddPaymentMethod)
	r.POST("/addcoupen", middleware.AdminAuth, controllers.AddCoupen)
	r.GET("/listcoupon", middleware.AdminAuth, controllers.ListCoupon)

	r.POST("/addproductoffer", middleware.AdminAuth, controllers.AddProductOffer)
	r.POST("/addcategoryoffer", middleware.AdminAuth, controllers.AddCategoryOffer)
	r.POST("/return", middleware.RequireAuth, controllers.RetutnOrder)

}
