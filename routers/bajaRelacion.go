package routers

import (
	"net/http"

	"github.com/nahuelvf17/proyectgo/bd"
	"github.com/nahuelvf17/proyectgo/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un errorr al borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(http.StatusCreated)
}
