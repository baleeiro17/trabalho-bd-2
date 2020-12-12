package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
	"trabalho-bd-2/internal/repositories"
)

func CriaVenda(ctx *gin.Context) {
	var venda models.Venda

	err := ctx.ShouldBindJSON(&venda)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insere a venda no banco de dados.
	err = repositories.InsereVenda(&venda)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retorna o status http Ok.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": venda})
}

func DevolveVenda(ctx *gin.Context) {
	var venda models.Venda

	//id do venda.
	id := ctx.Params.ByName("id")

	// busca a venda no banco de dados.
	err := repositories.BuscaVenda(&venda, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// retorna as informações do venda.
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": venda})
}
