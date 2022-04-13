package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// serveText writes given text as a response to http request.
func serveText(text string, w http.ResponseWriter) error {
	_, err := w.Write([]byte(text + "\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func main() {
	var (
		newMux = http.NewServeMux()
		port   = ":8000"
	)

	newMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			if err := serveText("NewServeMux", w); err != nil {
				fmt.Printf("/: error: %s\n", err.Error())
			}
			return
		}

		// Response with 404 - not Found if access other endpoints except `/`.
		http.NotFound(w, r)
		return
	})

	newMux.HandleFunc("/random-int", func(w http.ResponseWriter, r *http.Request) {
		text := strconv.Itoa(rand.Int())
		if err := serveText(text, w); err != nil {
			fmt.Printf("/random-int: error: %s\n", err.Error())
		}
		return
	})

	newMux.HandleFunc("/random-float", func(w http.ResponseWriter, r *http.Request) {
		text := strconv.FormatFloat(rand.Float64(), 'E', -1, 64)
		if err := serveText(text, w); err != nil {
			fmt.Printf("/random-float: error: %s\n", err.Error())
		}
		return
	})

	fmt.Printf("Server listening at http://localhost%s\n\n", port)
	log.Fatal(http.ListenAndServe(port, newMux))
}
