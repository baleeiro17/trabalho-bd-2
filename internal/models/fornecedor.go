package models

type Fornecedor struct {
	Id    int    `json:"id" binding:"required"`
	Nome  string `json:"nome" binding:"required"`
	Cnpj  string `json:"cnpj" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
