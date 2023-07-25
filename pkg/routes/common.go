package routes

import (
	"project/pkg/controllers/auth"
	"project/pkg/controllers/common"

	"github.com/gin-gonic/gin"
)

func CommonRoutes(r *gin.Engine) {
	r.POST("/signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.GET("/listproduct", common.DisplayProduct)
	r.GET("/displaycategory", common.DisplayCategory)
	r.POST("/sendotp", common.SendOtp)
	r.POST("/checkotp", common.Checkotp)
}
