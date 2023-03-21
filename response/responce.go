package response

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Status  bool
	Message string
	Errors  interface{}
	Data    interface{}
}

func ErrorMessage(g *gin.Context, msg string, errors interface{}, code int) {
	responce := response{
		Status:  false,
		Message: msg,
		Errors:  errors,
	}
	g.JSON(code, responce)
}
func SurcessMessage(g *gin.Context, msg string, data interface{}) {
	responce := response{
		Status:  true,
		Message: msg,
		Errors:  nil,
		Data:    data,
	}
	g.JSON(200, responce)
}

// response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
// response.SurcessMessage(g, "address added surcessfully", new)
// response.ErrorMessage(g, "failed to hash passsword", hash_err.Error(), 500)
