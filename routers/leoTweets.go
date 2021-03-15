package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/nahuelvf17/proyectgo/bd"
)

// LeoTweets leo los tweets
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro Pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro Pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweet(ID, pag)
	log.Println(respuesta, correcto)
	if correcto == false {
		http.Error(w, "Error al buscar los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
