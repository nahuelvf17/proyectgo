package bd

import (
	"context"
	"time"

	"github.com/nahuelvf17/proyectgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario chequea si existe el usuario antes de insertarlo
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex() // lo convierto en hex para manejarlo mejor

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
