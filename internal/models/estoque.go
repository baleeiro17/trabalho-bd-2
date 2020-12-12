package models

import "gopkg.in/mgo.v2/bson"

type Estoque struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Quantidade float64       `bson:"quantidade" json:"quantidade" binding:"required"`
	ProdutoId  string        `json:"produtoId" binding:"required"`
	Produto    *Produto      `bson:"-" json:"produto"` // não insere no banco de dados, somente é usado no GET.
}
