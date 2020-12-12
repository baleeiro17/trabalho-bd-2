package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"trabalho-bd-2/internal/config"
	"trabalho-bd-2/internal/models"
)

func InsereEstoque(data *models.Estoque) error {

	// adicionando id para o estoque.
	data.Id = bson.NewObjectId()

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// insere no banco de dados(mongo-db).
	err := db.C("estoque").Insert(data)
	if err != nil {
		return fmt.Errorf("Error ao inserir o produto no mongodb")
	}

	// inserção ocorreu com sucesso.
	return nil
}

func BuscaEstoque(data *models.Estoque, id string) error {

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// busca estoque no banco de dados(mongo-db).
	err := db.C("estoque").FindId(bson.ObjectIdHex(id)).One(data)
	if err != nil {
		return fmt.Errorf("Id do estoque não encontrado no banco de dados")
	}

	// busca o produto que faz parte do estoque para mostrar seus dados.
	err = db.C("produto").FindId(data.Produto.Id).One(data.Produto)
	if err != nil {
		return fmt.Errorf("Id do produto colocado no estoque não encontrado no banco de dados")
	}

	// busca o fornecedor que faz parte do produto para mostrar seus dados.
	err = db.C("fornecedor").FindId(data.Produto.Fornecedor.Id).One(data.Produto.Fornecedor)
	if err != nil {
		return fmt.Errorf("Id do fornecedor colocado no produto não encontrado no banco de dados")
	}

	// busca pelo estoque ocorrre com sucesso.
	return nil
}
