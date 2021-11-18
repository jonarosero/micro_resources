package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
	recursomodels "github.com/ascendere/resources/models/recursos_models"
)

func RegistroRecurso(w http.ResponseWriter, r *http.Request){
	var recurso recursomodels.Recurso

	err := json.NewDecoder(r.Body).Decode(&recurso)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+ err.Error(),400)
		return
	}

	_, status, err := recursobd.RegistroRecurso(recurso)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo Recurso", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar un nuevo Recurso", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}