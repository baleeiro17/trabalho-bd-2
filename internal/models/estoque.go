package models

type Estoque struct {
	Id         int `json:"id" binding:"required"`
	Quantidade int `json:"quantidade" binding:"required"`
	ProdutoId  int `json:"produtoId" binding:"required"`
}
