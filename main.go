package main

import (
	"log"

	"github.com/ascendere/resources/bd"
	"github.com/ascendere/resources/handlers"
)

func main (){
	if bd.ChequeoConnection() == 0{
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}