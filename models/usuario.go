package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Modelo para usuario*/
type Usuario struct {
	ID              primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	Nombre          string             `bson: "nombre json: "nombre,omitempty"`
	Apellidos       string             `bson: "apellidos" json:"appellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson: "email" json:"email,omitempty"`
	Password        string             `bson: "password" json:"password,omitempty"`
	Avatar          string             `bson: "avatar" json:"avatar,omitempty"`
	Banner          string             `bson: "banner" json:"banner,omitempty"`
	Ubicacion       string             `bson: "ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson: "sitoWeb" json:"sitioWeb,omitempty"`
}
