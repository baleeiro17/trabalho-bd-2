package models

import "gopkg.in/mgo.v2/bson"

type Produto struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Nome       string        `bson:"nome,omitempty" json:"nome,omitempty"`
	PrecoCusto float64       `bson:"precoCusto,omitempty" json:"precoCusto,omitempty"`
	PrecoVenda float64       `bson:"precoVenda,omitempty" json:"precoVenda,omitempty"`
	Fornecedor *Fornecedor   `bson:"fornecedor" json:"fornecedor"`
	// a flag omitempty so adiciona ao banco de dados se tiver um campo.
}
