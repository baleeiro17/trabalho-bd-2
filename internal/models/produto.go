package models

import "gopkg.in/mgo.v2/bson"

type Produto struct {
	Id           bson.ObjectId `bson:"_id" json:"id"`
	Nome         string        `bson:"nome" json:"nome" binding:"required"`
	PrecoCusto   float64       `bson:"precoCusto" json:"precoCusto" binding:"required"`
	PrecoVenda   float64       `bson:"precoVenda" json:"precoVenda" binding:"required"`
	FornecedorId string        `bson:"fornecedorId" json:"fornecedorId" binding:"required"`
	Fornecedor   *Fornecedor   `bson:"-" json:"fornecedor"` // não insere no banco de dados, somente é usado no GET.
}
