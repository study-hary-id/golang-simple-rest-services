package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.ServeFiles(
		// TODO: Replace according to the directory you want to save static files.
		"/www/*filepath",                   // Endpoint to access the static files.
		http.Dir("/home/haryansyah/Public/www")) // The location of the file to be served.
	log.Fatal(http.ListenAndServe(":8000", router))
}
