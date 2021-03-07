package main

import (
	"log"

	bd "github.com/nahuelvf17/proyectgo/bd"
	"github.com/nahuelvf17/proyectgo/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin coneccion a la base de datos")
		return
	}

	handlers.Manejadores()

}
