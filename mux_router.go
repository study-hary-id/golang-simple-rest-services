package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// articleHandler handles category and id of the articles.
func articleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := fmt.Fprintf(
		w,
		"%v in Category: %v\n",
		vars["id"],
		vars["category"],
	)
	if err != nil {
		return
	}
}

// detailsHandler handles details of particular articles.
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := fmt.Fprintf(w, "Details of ID: %v\n", vars["id"])
	if err != nil {
		return
	}
}

// settingsHandler handles settings of particular articles.
func settingsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := fmt.Fprintf(w, "Settings of ID: %v\n", vars["id"])
	if err != nil {
		return
	}
}

func main() {
	var (
		r    = mux.NewRouter()
		PORT = ":8000"
	)

	// Create route and define the handler with original style.
	//r.HandleFunc(
	//	"/articles/{category}/{id:[0-9]+}",
	//	articleHandler,
	//).Name("articleRoute")

	sr := r.PathPrefix("/articles").Subrouter()
	sr.HandleFunc("/{category}/{id:[0-9]+}", articleHandler).
		Name("articleHandler")
	sr.HandleFunc("/{category}/{id:[0-9]+}/details", detailsHandler).
		Name("detailsArticle")
	sr.HandleFunc("/{category}/{id:[0-9]+}/settings", settingsHandler).
		Name("settingsArticle")

	// Get the resource from any predefined routes.
	//url, err := r.Get("articleRoute").URL("category", "books", "id", "123")
	//fmt.Println(url.URL) 						// prints /articles/books/123

	s := &http.Server{
		Addr:         PORT,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(s.ListenAndServe())
}
