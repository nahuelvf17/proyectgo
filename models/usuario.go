package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usuario es la structura en la base
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id, omitepty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitepmty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitepmty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitepmty"`
	Banner          string             `bson:"banner" json:"banner,omitepmty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitepmty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitepmty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitepmty"`
}
