package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/nahuelvf17/proyectgo/bd"
)

//ObtenerBanner envia el avatar al HTTP
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	var openFile *os.File
	openFile, err = os.Open("uploads/banners/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Banner no encontrada", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Banner no encontrada", http.StatusBadRequest)
	}
}
