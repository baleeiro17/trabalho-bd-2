package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
)

func CriaProduto(ctx *gin.Context) {
	var produto models.Produto

	err := ctx.ShouldBindJSON(&produto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// inserir produto em um banco de dados.

	// cria um fornecedor.
	ctx.JSON(http.StatusOK, gin.H{"data": produto})

}

func DevolveProduto(ctx *gin.Context) {
	//var fornecedor models.Fornecedor

	//id do produto.
	id := ctx.Params.ByName("id")

	// busca o produto no banco de dados.

	// retorna  o id do fornecedor.
	ctx.JSON(http.StatusOK, gin.H{"data": id})
}
