package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		mux  = &http.ServeMux{}
		port = ":8000"
	)

	/*
		A path without a trailing backslash (/), refers to an explicit path.
	*/
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "%q %v", r.Method, r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	})

	/*
		Subtrees match the start of a path, and include the trailing /.
		`GET /articles` is the same as `GET /articles/`.
	*/
	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "%q %v", r.Method, r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	})

	fmt.Printf("Server listening at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
