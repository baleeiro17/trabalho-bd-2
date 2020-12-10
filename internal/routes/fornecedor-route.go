package routes

import (
	"github.com/gin-gonic/gin"
	"trabalho-bd-2/internal/controllers"
)

func DefineRotasParaFornecedor(route *gin.RouterGroup) {
	route.POST("/fornecedor", controllers.CriaFornecedor)
	route.GET("/fornecedor/:id", controllers.DevolveFornecedor)
}
