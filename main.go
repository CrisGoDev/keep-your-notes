package main

import (
	"net/http"
	"strconv"
	"time"

	note "github.com/CrisGoDev/keep-your-notes/internal/Note"
	"github.com/CrisGoDev/keep-your-notes/pkg/boostrap"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Initial Configuration
func main() {

	router := mux.NewRouter()
	_ = godotenv.Load()
	l := boostrap.InitLogger()

	db, err := boostrap.DBConnection()

	if err != nil {
		l.Fatal(err)
	}

	noteRepo := note.NewRepo(l, db)
	noteService := note.NewService(l, noteRepo)
	noteEnd := note.MakeEndpoint(noteService)

	router.HandleFunc("/notes", noteEnd.GetAll).Methods("GET")
	router.HandleFunc("/notes", noteEnd.Create).Methods("POST")
	router.HandleFunc("/notes/{id}", noteEnd.Update).Methods("PATCH")
	router.HandleFunc("/notes/{id}", noteEnd.Delete).Methods("DELETE")
	router.HandleFunc("/notes/{id}", noteEnd.Get).Methods("GET")

	port, err := boostrap.GetPort()

	if err != nil {
		l.Fatal("Please specify the port")
		return
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	l.Fatal(srv.ListenAndServe())
}
