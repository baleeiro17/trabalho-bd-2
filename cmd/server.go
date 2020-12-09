package main

import (
	"github.com/gin-gonic/gin"
	app "trabalho-bd-2/internal"
)

func main() {
	router := gin.Default()

	// define as rotas padrões
	app.DefineRoutes(router, "/api/v1")

	// roda o servidor HTTP na porta 8080.
	router.Run(":8080")
}
