package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"trabalho-bd-2/internal/config"
	"trabalho-bd-2/internal/models"
)

func InsereVenda(data *models.Venda) error {

	// adicionando id para a venda.
	data.Id = bson.NewObjectId()

	// adicionando time para a venda.
	data.Data = time.Now()

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// insere no banco de dados(mongo-db).
	err := db.C("venda").Insert(data)
	if err != nil {
		return fmt.Errorf("Error ao inserir a venda no mongodb")
	}

	// inserção ocorreu com sucesso.
	return nil
}

func BuscaVenda(data *models.Venda, id string) error {

	// variavel de conexão do banco de dados.
	db := config.GetMongoDB()

	// busca a venda no banco de dados(mongo-db).
	err := db.C("venda").FindId(bson.ObjectIdHex(id)).One(data)
	if err != nil {
		return fmt.Errorf("Id do estoque não encontrado no banco de dados")
	}

	// busca os produtos que faz parte da venda para mostrar seus dados.
	// ou poderiamos colocar data.quantidade.
	for i := 0; i < len(data.Lista); i++ {

		err = db.C("produto").FindId(data.Lista[i].Id).One(&data.Lista[i])
		if err != nil {
			return fmt.Errorf("Id do produto colocado na venda não encontrado no banco de dados")
		}

		// busca o fornecedor que faz parte do produto para mostrar seus dados.
		err = db.C("fornecedor").FindId(data.Lista[i].Fornecedor.Id).One(data.Lista[i].Fornecedor)
		if err != nil {
			return fmt.Errorf("Id do fornecedor colocado no produto não encontrado no banco de dados")
		}
	}

	// busca pela venda ocorreu com sucesso.
	return nil
}
