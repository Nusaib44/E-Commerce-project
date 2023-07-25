package routes

import (
	"project/middleware"
	"project/pkg/controllers/common"
	"project/pkg/controllers/user"
	// "project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {

	users := r.Group("/user")
	users.Use(middleware.RequireAuth)

	users.GET("/product", common.ViewProduct)
	users.GET("/displaycategory", common.DisplayCategory)

	users.GET("/cart", user.Cart)
	users.POST("/addtocart", user.AddToCart)
	users.POST("/updatecartitemquantity", user.UpdateCartItemQuantity)
	users.DELETE("/cartitemdelete", user.CartItemDelete)

	users.POST("/addaddress", user.AddNewAddress)
	users.POST("/checkout", user.Checkout)

	users.POST("/placeorder", user.PlaceOrder)
	users.GET("/listorder", user.ListOrder)
	users.POST("/ordercancelation", user.OrderCancelation)

	users.GET("/userinfo", user.UserInfo)
	users.POST("/edituserinfo", user.EditUserInfo)
	users.POST("/editpassword", user.ChangePassword)
	users.POST("/stripe", user.StripePayment)
	users.POST("/validatecoupon", user.ValidateCoupen)
	users.POST("/return", user.RetutnOrder)

}
