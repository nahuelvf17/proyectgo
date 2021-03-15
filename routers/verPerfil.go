package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nahuelvf17/proyectgo/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	log.Println("aca es ver perfil, voy a buscar buscoPerfil")
	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}
	log.Println("todo ok verperfil")
	w.Header().Set("contect-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
