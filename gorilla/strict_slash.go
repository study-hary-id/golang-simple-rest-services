package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var (
		r    = mux.NewRouter()
		PORT = ":8000"
	)

	r.StrictSlash(true)
	// GET /articles == GET /articles/
	r.Path("/articles/").Handler(http.HandlerFunc(articleHandler))

	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
