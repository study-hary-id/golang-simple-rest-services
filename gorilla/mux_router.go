package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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

	// Create subrouter for /articles endpoint: Best Practices.
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
