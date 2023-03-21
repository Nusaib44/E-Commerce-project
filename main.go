package main

import (
	"project/initializers"
	"project/pkg/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	route := gin.Default()
	routes.AdminRoute(route)
	routes.UserRoute(route)
	routes.CommonRoutes(route)
	route.Run()
}
