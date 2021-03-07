package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nahuelvf17/proyectgo/bd"
	"github.com/nahuelvf17/proyectgo/models"
)

// Registro es la funcion para crear en la base de datos el registro de la base de datos
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El mail de usuario es requerido", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con este email", 400)
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
