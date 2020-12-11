package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
	"trabalho-bd-2/internal/repositories"
)

func CriaFornecedor(ctx *gin.Context) {
	var fornecedor models.Fornecedor
	err := ctx.ShouldBindJSON(&fornecedor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insere o fornecedor no banco de dados(mongodb).
	err = repositories.InsereFornecedor(&fornecedor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retorna http ok.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": fornecedor})

}

func DevolveFornecedor(ctx *gin.Context) {
	var fornecedor models.Fornecedor

	//id do fornecedor.
	id := ctx.Params.ByName("id")

	// busca o fornecedor no banco de dados.
	err := repositories.BuscaFornecedor(&fornecedor, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retorna  o id do fornecedor.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": fornecedor})
}
