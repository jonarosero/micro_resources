package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoRecursos() ([]*recursomodels.DevuelvoRecurso, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	var results []*recursomodels.DevuelvoRecurso

	var resultadoRecurso []*recursomodels.Recurso

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, err
	}

	for cur.Next(ctx) {
		var s recursomodels.Recurso
		err := cur.Decode(&s)
		if err != nil {
			return results, err
		}
		resultadoRecurso = append(resultadoRecurso, &s)
	}

	for _, recurso := range resultadoRecurso {
		colTipo := db.Collection("tipoRecurso")
		var tipo recursomodels.TipoRecurso

		errTipo := colTipo.FindOne(ctx, bson.M{"_id": recurso.TipoID}).Decode(&tipo)

		if errTipo != nil {
			return results, err
		}

		var aux recursomodels.DevuelvoRecurso

		aux.ID = recurso.ID
		aux.NombreRecurso = recurso.NombreRecurso
		aux.Imagen = recurso.Imagen
		aux.CantidadDisponible = recurso.CantidadDisponible
		aux.CantidadExistente = recurso.CantidadExistente
		aux.TipoRecurso.ID = tipo.ID
		aux.TipoRecurso.NombreTipo = tipo.NombreTipo
		aux.TipoRecurso.DescripcionTipo = tipo.DescripcionTipo

		results = append(results, &aux)
	}

	if err != nil {
		return results, err
	}
	return results, err
}
