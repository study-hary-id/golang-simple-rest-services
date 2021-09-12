package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.ServeFiles(
		"/www/*filepath",							// Endpoint to access the static files.
		http.Dir("/home/haryansyah/Public/www"))		// The location of the file to be served.
	log.Fatal(http.ListenAndServe(":8000", router))
}
