package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nahuelvf17/proyectgo/bd"
	"github.com/nahuelvf17/proyectgo/models"
)

// Email mail del usuario
var Email string

// Usuario del token
var Usuario string

//ProcesoToken para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			Usuario = claims.ID.Hex()
		}

		return claims, encontrado, ID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido")
	}

	return claims, false, string(""), err
}
