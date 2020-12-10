package models

import "gopkg.in/mgo.v2/bson"

type Fornecedor struct {
	Id    bson.ObjectId `bson:"_id"" json:"id" `
	Nome  string        `bson:nome json:"nome" binding:"required"`
	Cnpj  string        `bson:cnpj json:"cnpj" binding:"required"`
	Email string        `bson: email json:"email" binding:"required"`
	Phone string        `bson:phone json:"phone" binding:"required"`
}
