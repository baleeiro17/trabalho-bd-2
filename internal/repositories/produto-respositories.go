package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"trabalho-bd-2/internal/config"
	"trabalho-bd-2/internal/models"
)

func InsereProduto(data *models.Produto) error {

	// adicionando id para o produto.
	data.Id = bson.NewObjectId()

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// insere no banco de dados(mongo-db).
	err := db.C("produto").Insert(data)
	if err != nil {
		return fmt.Errorf("Error ao inserir o produto no mongodb")
	}

	// inserção ocorreu com sucesso.
	return nil
}

func BuscaProduto(data *models.Produto, id string) error {

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// busca produto no banco de dados(mongo-db).
	err := db.C("produto").FindId(bson.ObjectIdHex(id)).One(data)
	if err != nil {
		return fmt.Errorf("Id do produto não encontrado no mongodb")
	}

	// busca o fornecedor que faz parte do produto para mostrar seus dados.
	err = db.C("fornecedor").FindId(data.Fornecedor.Id).One(data.Fornecedor)
	if err != nil {
		return fmt.Errorf("Id do fornecedor colocado no produto não encontrado no mongodb")
	}

	// busca por id ocorreu com sucesso.
	return nil
}
