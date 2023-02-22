package note

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/CrisGoDev/keep-your-notes/domain"
	"github.com/CrisGoDev/keep-your-notes/pkg/meta"
	"github.com/gorilla/mux"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		GetAll Controller
		Create Controller
		Get    Controller
		Update Controller
		Delete Controller
	}

	CreateRequest struct {
		Title string `json:"title"`
		Body  string `json:"body" `
	}

	UpdateRequest struct {
		Title *string `json:"title"`
		Body  *string `json:"body"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

func MakeEndpoint(s Service) Endpoints {

	return Endpoints{
		GetAll: makeGetAllEndpoint(s),
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		Update: makeUdpateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func fieldsOfNote() map[string]int {
	fields := make(map[string]int)
	t := reflect.TypeOf(domain.Note{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("json")

		fields[field] = i + 1
	}

	return fields
}

var orders = map[string]int{
	"asc":  1,
	"desc": 2,
}

func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		queryVariable := r.URL.Query()
		fields := fieldsOfNote()
		order := queryVariable.Get("order")
		fieldToOrder := queryVariable.Get("order_by")

		if fieldToOrder != "" && fields[fieldToOrder] < 1 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: "Incorrect field"})
			return
		}

		if order != "" && orders[order] < 1 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: "Incorrect ordering mode"})
			return
		}

		filters := Filters{
			Body:      queryVariable.Get("body"),
			Title:     queryVariable.Get("title"),
			CreatedAt: queryVariable.Get("created_at"),
			OrderBy:   fieldToOrder,
			Order:     order,
		}

		limit, _ := strconv.Atoi(queryVariable.Get("limit"))
		page, _ := strconv.Atoi(queryVariable.Get("page"))

		count, err := s.Count(filters)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}
		meta, err := meta.New(page, limit, count)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}

		users, err := s.GetAll(filters, meta.Offset(), meta.Limit())

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: users, Meta: meta})
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Inavlid request format"})
			return
		}

		if req.Body == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "body required"})
			return
		}

		if req.Title == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "title required"})
			return
		}

		note, err := s.Create(req.Title, req.Body)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 201, Data: note})
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)

		id := path["id"]

		note, err := s.Get(id)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: `note doesn't exist`})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: note})
	}
}

func makeUdpateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Inavlid request format"})
			return
		}

		if req.Title != nil && *req.Title == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Inavlid request format"})
			return
		}

		if req.Body != nil && *req.Body == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Inavlid request format"})
			return
		}

		path := mux.Vars(r)
		id := path["id"]

		err := s.Update(id, req.Title, req.Body)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: `note doesn't exist`})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "succes"})
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		path := mux.Vars(r)

		id := path["id"]

		err := s.Delete(id)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: `note doesn't exist`})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "succes"})
	}
}
