package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"trabalho-bd-2/internal/config"
	"trabalho-bd-2/internal/models"
)

func InsereFornecedor(data *models.Fornecedor) error {

	// adicionando id para o fornecedor.
	data.Id = bson.NewObjectId()

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// insere no banco de dados(mongo-db).
	err := db.C("fornecedor").Insert(data)
	if err != nil {
		return fmt.Errorf("Error ao inserir o fornecedor no mongodb")
	}

	// inserção ocorreu com sucesso.
	return nil
}

func BuscaFornecedor(data *models.Fornecedor, id string) error {

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// insere no banco de dados(mongo-db).
	err := db.C("fornecedor").FindId(bson.ObjectIdHex(id)).One(data)
	if err != nil {
		return fmt.Errorf("Id do fornecedor não encontrado no mongodb")
	}

	// busca por id ocorreu com sucesso.
	return nil
}
