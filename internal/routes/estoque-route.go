package routes

import (
	"github.com/gin-gonic/gin"
	"trabalho-bd-2/internal/controllers"
)

func DefineRotasParaEstoque(route *gin.RouterGroup) {
	route.POST("/estoque", controllers.CriaEstoque)
	route.GET("/estoque/:id", controllers.DevolveEstoque)
}
