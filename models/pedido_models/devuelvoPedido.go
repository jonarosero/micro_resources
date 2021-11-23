package pedidomodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoPedido struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FechaPedido       time.Time          `bson:"fechaPedido" json:"fechaPedido,omitempty"`
	FechaDevolucion   time.Time          `bson:"fechaDevolucion" json:"fechaDevolucion,omitempty"`
	IdProyecto        primitive.ObjectID `bson:"idProyecto,omitempty" json:"idProyecto"`
	InformePedido     string             `bson:"informePedido" json:"informePedido,omitempty"`
	InformeDevolucion string             `bson:"informeDevolucion" json:"informeDevolucion,omitempty"`
	Estado            bool               `bson:"estado" json:"estado,omitempty"`
	Recurso           []struct {
		RecursoID      primitive.ObjectID `bson:"recursoid,omitempty" json:"recursoid"`
		NombreRecurso  string             `bson:"nombreRecurso" json:"nombreRecurso,omitempty"`
		CantidadPedida int                `bson:"cantidadpedida" json:"cantidadpedida,omitempty"`
	}
	Usuario			struct {
		UsuarioID	primitive.ObjectID `bson:"usuarioid,omitempty" json:"usuarioid"`
		Email  string             `bson:"email" json:"email,omitempty"`
	}
}
