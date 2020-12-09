package models

type Fornecedor struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome" binding:"required"`
	Cnpj  string `json:"cnpj" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
