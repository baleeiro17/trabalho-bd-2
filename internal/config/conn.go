package config

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var db *mgo.Database

// the predefined init() function sets off a piece of code
// to run before any other part of your package.
func init() {

	// host do mongodb
	host := os.Getenv("MONGO_HOST")

	// nome do database do mongodb.
	dbName := os.Getenv("MONGO_DB_NAME")

	// cria uma session no mongodb.
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}

	// variavel com a sessão do mongo db.
	db = session.DB(dbName)

}

// retorna a sessão com mongodb.
func GetMongoDB() *mgo.Database {
	return db
}
