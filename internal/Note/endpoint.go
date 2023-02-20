package note

import (
	"encoding/json"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Get Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
	}
)

func MakeEndpoint(repo Repository) Endpoints {

	return Endpoints{
		Get: makeGetEndpoint(repo),
	}
}

func makeGetEndpoint(repo Repository) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := repo.Get()

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: users})
	}
}
