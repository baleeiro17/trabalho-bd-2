package models

type Produto struct {
	Id           int     `json:"id" binding:"required"`
	Nome         string  `json:"nome" binding:"required"`
	PrecoCusto   float64 `json:"precoCusto" binding:"required"`
	PrecoVenda   float64 `json:"precoVenda" binding:"required"`
	FornecedorId int     `json:"fornecedorId" binding:"required"`
}
