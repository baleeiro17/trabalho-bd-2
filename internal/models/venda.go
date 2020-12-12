package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Venda struct {
	Id            bson.ObjectId `bson:"_id" json:"id"`
	Quantidade    int           `bson:"quantidade" json:"quantidade" binding:"required"`
	NumNotaFiscal string        `bson:"numNotaFiscal" json:"numNotaFiscal" binding:"required"`
	ValorTotal    float64       `bson:"valorTotal" json:"valorTotal" binding:"required"`
	Data          time.Time     `bson:"data" json:"data"`
	Lista         []Produto     `bson:"listaCompras" json:"listaCompras" binding:"required"`
}
