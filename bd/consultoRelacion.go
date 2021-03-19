package bd

import (
	"context"
	"log"
	"time"

	"github.com/nahuelvf17/proyectgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta relacion de usuarios
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	log.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}
