package routes

import (
	"project/middleware"
	"project/pkg/controllers/admin"
	"project/pkg/controllers/payment"

	"github.com/gin-gonic/gin"
)

func AdminRoute(r *gin.Engine) {
	adminr := r.Group("/admin")
	adminr.Use(middleware.AdminAuth)
	// <--------------------------------USER MANAGMENT-------------------------------->
	adminr.GET("/listuser", admin.ListUser)
	adminr.PATCH("/block/:id", admin.BlockUser)
	adminr.PATCH("/unblock/:id", admin.Unblock)

	// <--------------------------------PRODUCT MANAGMENT<-------------------------------->
	adminr.POST("/addproduct", admin.AddProduct)
	adminr.PATCH("/editproduct", admin.EditProduct)
	adminr.DELETE("/deleteproduct", admin.DeleteProduct)

	// <--------------------------------CATEGORY MANAGMENT<-------------------------------->
	adminr.POST("/addcategory", admin.AddCategory)
	adminr.PATCH("/editcategory", admin.EditCategory)
	adminr.PUT("/deletecategory", admin.DeleteCategory)

	adminr.POST("/addpayment", payment.AddPaymentMethod)
	adminr.POST("/addcoupen", admin.AddCoupen)
	adminr.GET("/listcoupon", admin.ListCoupon)

	adminr.POST("/addproductoffer", admin.AddProductOffer)
	adminr.POST("/addcategoryoffer", admin.AddCategoryOffer)

	adminr.POST("/updateorderstatus", admin.OrderStatus)
	adminr.GET("/listorder", admin.Listorders)
	adminr.GET("/listreturn", admin.ListReturn)
	adminr.POST("/updatereturnstatus", admin.ReturnManangement)
	adminr.POST("/cashback", admin.Cashback)

}
