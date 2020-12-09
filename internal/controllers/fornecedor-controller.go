package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trabalho-bd-2/internal/models"
)

func CriaFornecedor(ctx *gin.Context) {
	var fornecedor models.Fornecedor
	err := ctx.ShouldBindJSON(&fornecedor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cria um fornecedor.
	ctx.JSON(http.StatusOK, gin.H{"data": fornecedor})

}
