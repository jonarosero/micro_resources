package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChequeoExistenRecursos(id primitive.ObjectID, cantidadPedida int) (string, error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recurso")
	col := db.Collection("recurso")

	var resultado recursomodels.Recurso
	var nombre string

	error := col.FindOne(ctx, bson.M{"_id":id}).Decode(&resultado)

	if error != nil {
		return nombre, error, "No se encuentra el recurso"
	}

	if resultado.CantidadDisponible == 0 {
		return nombre, error, "El recurso no se encuentra disponible"
	}

	if cantidadPedida > resultado.CantidadDisponible {
		return nombre, error, "No se dispone de tantos recursos"
	}

	if cantidadPedida > resultado.CantidadExistente {
		return nombre, error, "No existen tantos recursos"
	}

	if cantidadPedida < resultado.CantidadDisponible && cantidadPedida > 0 {
		resultado.CantidadDisponible = resultado.CantidadDisponible - cantidadPedida

		registro := make(map[string]interface{})

		registro["cantidadDisponible"] = resultado.CantidadDisponible
		updtString := bson.M{
			"$set": registro,
		}
	
		filtro := bson.M{"_id": bson.M{"$eq": resultado.ID}}
	
		_, err := col.UpdateOne(ctx, filtro, updtString)

		if err != nil {
			return "", error, err.Error()
		}

		nombre = resultado.NombreRecurso
	
	}

	return nombre, error, id.String()
}