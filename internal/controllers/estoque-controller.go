package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
)

func CriaEstoque(ctx *gin.Context) {
	var estoque models.Estoque
	err := ctx.ShouldBindJSON(&estoque)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insere o estoque no banco de dados.

	// retorna http ok.
	ctx.JSON(http.StatusOK, gin.H{"data": estoque})

}

func DevolveEstoque(ctx *gin.Context) {
	//var estoque models.Estoque

	//id do estoque.
	id := ctx.Params.ByName("id")

	// busca o estoque no banco de dados.

	// retorna  o id do estoque.
	ctx.JSON(http.StatusOK, gin.H{"data": id})

}
