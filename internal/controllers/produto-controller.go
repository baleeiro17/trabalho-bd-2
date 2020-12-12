package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
	"trabalho-bd-2/internal/repositories"
)

func CriaProduto(ctx *gin.Context) {
	var produto models.Produto

	err := ctx.ShouldBindJSON(&produto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// inserir produto no banco de dados.
	err = repositories.InsereProduto(&produto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// produto criado com sucesso.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": produto})

}

func DevolveProduto(ctx *gin.Context) {
	var produto models.Produto

	// id do produto.
	id := ctx.Params.ByName("id")

	// busca o produto no banco de dados.
	err := repositories.BuscaProduto(&produto, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retorna as informações do produto.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": produto})
}
