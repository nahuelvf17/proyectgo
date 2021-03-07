package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de coneccion de la base de datos
var MongoCN = ConectarDB()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://nahuel21:bauty21@twitter.myoly.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConectarDB es la funcion que me permite conectar a la base de datos
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Conexcion exitosa con la DB")

	return client
}

// ChequeoConnection chequear si esta levantada la base
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
