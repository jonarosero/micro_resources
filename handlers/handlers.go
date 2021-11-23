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
	router.HandleFunc("/eliminarRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.EliminarRecurso))).Methods("DELETE")
	router.HandleFunc("/actualizarRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.ActualizarRecurso))).Methods("PUT")
	router.HandleFunc("/buscarRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.BuscarRecurso))).Methods("GET")
	router.HandleFunc("/listarRecursos", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.ListarRecursos))).Methods("GET")

	//Llamada al CRUD de Tipos de Recursos
	router.HandleFunc("/registrarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.RegistrarTipoRecurso))).Methods("POST")
	router.HandleFunc("/eliminarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.EliminarTipoRecurso))).Methods("DELETE")
	router.HandleFunc("/buscarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.BuscarTipoRecurso))).Methods("GET")
	router.HandleFunc("/listarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.ListarTiposRecurso))).Methods("GET")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}