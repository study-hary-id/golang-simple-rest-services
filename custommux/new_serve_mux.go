package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

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
		PORT   = ":8000"
	)

	newMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := serveText("NewServeMux", w); err != nil {
			fmt.Printf("/: error: %s\n", err.Error())
		}
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

	fmt.Printf("Server listening at http://localhost%s\n\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, newMux))
}
