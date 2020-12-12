package models

import "gopkg.in/mgo.v2/bson"

type Fornecedor struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Nome  string        `bson:"nome,omitempty" json:"nome,omitempty"`
	Cnpj  string        `bson:"cnpj,omitempty" json:"cnpj,omitempty"`
	Email string        `bson:"email,omitempty" json:"email,omitempty"`
	Phone string        `bson:"phone,omitempty" json:"phone,omitempty"`
	// a flag omitempty so adiciona ao banco de dados se tiver um campo.
}
