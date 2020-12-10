package routes

import (
	"github.com/gin-gonic/gin"
	"trabalho-bd-2/internal/controllers"
)

func DefineRotasParaVenda(route *gin.RouterGroup) {
	route.POST("/venda", controllers.CriaVenda)
	route.GET("/venda/:id", controllers.DevolveVenda)
}
