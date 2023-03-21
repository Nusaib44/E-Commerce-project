package routes

import (
	"project/middleware"
	"project/pkg/controllers"

	// "project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {

	r.GET("/cart", middleware.RequireAuth, controllers.Cart)
	r.POST("/addtocart", middleware.RequireAuth, controllers.AddToCart)
	r.POST("/updatecartitemquantity", middleware.RequireAuth, controllers.UpdateCartItemQuantity)
	r.DELETE("/cartitemdelete", middleware.RequireAuth, controllers.CartItemDelete)

	r.POST("/addaddress", middleware.RequireAuth, controllers.AddNewAddress)
	r.POST("/checkout", middleware.RequireAuth, controllers.Checkout)

	r.POST("/placeorder", middleware.RequireAuth, controllers.PlaceOrder)
	r.GET("/listorder", middleware.RequireAuth, controllers.ListOrder)
	r.POST("/ordercancelation", middleware.RequireAuth, controllers.OrderCancelation)

	r.GET("/userinfo", middleware.RequireAuth, controllers.UserInfo)
	r.POST("/edituserinfo", middleware.RequireAuth, controllers.EditUserInfo)
	r.POST("/editpassword", middleware.RequireAuth, controllers.ChangePassword)
	r.POST("/stripe", middleware.RequireAuth, controllers.StripePayment)
	r.POST("/validatecoupon", middleware.RequireAuth, controllers.ValidateCoupen)

}
