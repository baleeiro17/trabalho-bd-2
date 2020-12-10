package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
)

func CriaVenda(ctx *gin.Context) {
	var venda models.Venda
	err := ctx.ShouldBindJSON(&venda)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insere a venda no banco de dados.

	// retorna o status http Ok.
	ctx.JSON(http.StatusOK, gin.H{"data": venda})

}

func DevolveVenda(ctx *gin.Context) {
	//var venda models.venda

	//id do venda.
	id := ctx.Params.ByName("id")

	// busca a venda no banco de dados.

	// retorna  o id da venda.
	ctx.JSON(http.StatusOK, gin.H{"data": id})
}
