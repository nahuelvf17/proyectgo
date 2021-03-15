package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nahuelvf17/proyectgo/models"
)

// GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("bauty")

	paylod := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, paylod)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
