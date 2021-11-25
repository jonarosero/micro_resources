package routers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ascendere/resources/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Email string
var IDUsuario string
var Nombre string
var Tk string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersDelUniverso")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		var usuario struct {
			ID primitive.ObjectID `json:"id"`
		}

		client := &http.Client{}

		req, errRequest := http.NewRequest("GET", "http://34.123.95.33/verPerfil?id="+claims.Id, nil)

		if errRequest != nil {
			return claims, false, "", errRequest
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+tk)

		resp, error := client.Do(req)

		if error != nil {
			return claims, false, "", error
		}

		defer resp.Body.Close()

		bodyBytes, errorBytes := ioutil.ReadAll(resp.Body)

		if errorBytes != nil {
			return claims, false, "", error
		}

		json.Unmarshal(bodyBytes, &usuario)

		if len(usuario.ID) > 0 {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
			Nombre = claims.Nombre + " " + claims.Apellidos
			Tk = tk
		}
		return claims, true, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}

