package routes

import (
	"github.com/gin-gonic/gin"
	"trabalho-bd-2/internal/controllers"
)

func DefineRotasParaProduto(route *gin.RouterGroup) {
	route.POST("/produto", controllers.CriaProduto)
	route.GET("/produto/:id", controllers.DevolveProduto)
}
