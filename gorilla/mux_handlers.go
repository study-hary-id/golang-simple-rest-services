package main

import (
	"fmt"
	"net/http"

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
