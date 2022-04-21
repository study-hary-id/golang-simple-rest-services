package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var (
		r    = mux.NewRouter()
		PORT = ":8000"
	)

	// Create subrouter for /articles endpoint: Best Practices.
	sr := r.PathPrefix("/articles").Subrouter()
	sr.HandleFunc("/{category}/{id:[0-9]+}", articleHandler).
		Name("articleHandler")
	sr.HandleFunc("/{category}/{id:[0-9]+}/details", detailsHandler).
		Name("detailsArticle")
	sr.HandleFunc("/{category}/{id:[0-9]+}/settings", settingsHandler).
		Name("settingsArticle")

	// Implement handlers.LoggingHandler from gorilla.
	loggedRouter := handlers.LoggingHandler(os.Stdout, sr)

	s := &http.Server{
		Addr:         PORT,
		Handler:      loggedRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(s.ListenAndServe())
}
