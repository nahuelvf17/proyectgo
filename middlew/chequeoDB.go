package middlew

import (
	"net/http"

	"github.com/nahuelvf17/proyectgo/bd"
)

// ChequeoDB es el middlew que me permite conocer el estado de la DB
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida a la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
