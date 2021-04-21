package main

import (
	"log"

	"github.com/oflores/twitor-olf/bd"
	"github.com/oflores/twitor-olf/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
