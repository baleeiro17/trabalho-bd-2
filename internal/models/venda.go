package models

type Venda struct {
	Id            int     `json:"id" binding:"required"`
	ProdutoId     int     `json:"produtoId" binding:"required"`
	Quantidade    int     `json:"quantidade" binding:"required"`
	NumNotaFiscal string  `json:"numNotaFiscal" binding:"required"`
	ValorTotal    float64 `json:"valorTotal" binding:"required"`
	Data          string  `json:"data" binding:"required"`
}
