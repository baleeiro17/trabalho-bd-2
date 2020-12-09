package internal

import (
	"github.com/gin-gonic/gin"
	"trabalho-bd-2/internal/routes"
)

func DefineRoutes(router *gin.Engine, rotas string) {

	rotasV1 := router.Group(rotas)

	// define as rotas para o fornecedor.
	routes.DefineRotasParaFornecedor(rotasV1)

	// define as rotas para o produto.

	// define as rotas para o estoque.

}
