package admin

import (
	"fmt"
	"project/initializers"
	"project/pkg/models"
	"project/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderStatus(g *gin.Context) {

	params := g.Query("id")
	prm := g.Query("status")
	id, _ := strconv.Atoi(params)

	var order models.Order

	initializers.DB.Raw("update orders SET status=? WHERE id=?", prm, id).Scan(order)

}

func Listorders(g *gin.Context) {

	var orders []models.Order
	initializers.DB.Find(&orders)
	response.SurcessMessage(g, "displaying Orders", orders)

}

func ReturnManangement(g *gin.Context) {

	params := g.Query("id")
	returnId, _ := strconv.Atoi(params)
	status := g.Query("status")
	var returrndb models.Return

	println(returnId)
	initializers.DB.Raw("update returns set status=? where id=?", status, 1).Scan(&returrndb)
	fmt.Println(returrndb)
	response.SurcessMessage(g, "updated", returrndb)

}

func ListReturn(g *gin.Context) {

	var returndb []models.Return
	initializers.DB.Find(&returndb)
	response.SurcessMessage(g, "displaying Returns", returndb)

}

func Cashback(g *gin.Context) {

	params := g.Query("id")
	returnId, _ := strconv.Atoi(params)

	var rtn models.Return

	initializers.DB.Raw("SELECT price, user_id FROM orders WHERE id=?", returnId).Scan(&rtn)

	var walet models.Walet
	var balance int
	initializers.DB.Raw("select balance from walets where id=?", rtn.UserID).Scan(&balance)
	newbalance := balance + rtn.Price
	initializers.DB.Raw("UPDATE walets SET balance=? where id=?", newbalance, rtn.UserID).Scan(&walet)

	response.SurcessMessage(g, "amount added to users to wallet", walet)
}
