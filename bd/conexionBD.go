package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN database connection */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dferreiro:da%232457fe_M0n90d8@cluster0.ssnkxgu.mongodb.net/?retryWrites=true&w=majority")

/* ConectarBD para conectar a BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa a la BD")
	return client
}

/* CheckConnection ping to database */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
