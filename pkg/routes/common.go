package routes

import (
	"project/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func CommonRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/listproduct", controllers.DisplayProduct)
	r.GET("/displaycategory", controllers.DisplayCategory)
	r.POST("/sendotp", controllers.SendOtp)
	r.POST("/checkotp", controllers.Checkotp)
}
