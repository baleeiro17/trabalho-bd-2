package models

import "gopkg.in/mgo.v2/bson"

type Estoque struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Quantidade float64       `bson:"quantidade" json:"quantidade" binding:"required"`
	Produto    *Produto      `bson:"produto" json:"produto" binding:"required"`
	// a flag omitempty so adiciona ao banco de dados se tiver um campo.
}
