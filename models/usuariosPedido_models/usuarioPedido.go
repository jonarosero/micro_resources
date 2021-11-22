package usuariopedidomodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsuarioPedido struct {
	UsuarioID primitive.ObjectID `bson:"usuarioid" json:"usuarioId"`
	PedidoID primitive.ObjectID `bson:"pedidoid" json:"pedidoid"`
}
