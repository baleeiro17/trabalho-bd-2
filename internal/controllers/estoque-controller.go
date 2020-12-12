package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
	"trabalho-bd-2/internal/repositories"
)

func CriaEstoque(ctx *gin.Context) {
	var estoque models.Estoque
	err := ctx.ShouldBindJSON(&estoque)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insere o estoque no banco de dados.
	err = repositories.InsereEstoque(&estoque)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// estoque criado com sucesso.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": estoque})
}

func DevolveEstoque(ctx *gin.Context) {
	var estoque models.Estoque

	//id do estoque.
	id := ctx.Params.ByName("id")

	// busca o estoque no banco de dados.
	err := repositories.BuscaEstoque(&estoque, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retorna as informações do produto.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": estoque})
}
