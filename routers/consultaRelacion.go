package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nahuelvf17/proyectgo/bd"
	"github.com/nahuelvf17/proyectgo/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application/type")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
