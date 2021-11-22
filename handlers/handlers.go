package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/resources/middlew"
	recursorouters "github.com/ascendere/resources/routers/recurso_routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	//Llamada al CRUD de Recursos
	router.HandleFunc("/registroRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.RegistroRecurso))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}