package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// city constructs the Name and the Area of a city.
type city struct {
	Name string
	Area uint64
}

// filterContentType checks content type as JSON in Header.
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware.")
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := w.Write([]byte("415 unsupported media type, send JSON"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// setServerTimeCookie added server timestamp for response cookie.
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		
		cookie := http.Cookie{
			Name:       "Server-Time(UTC)",
			Value:      strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware.")
	})
}

// mainLogic handles the POST request and log out the data.
func mainLogic(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var (
			tempCity city
			err = json.NewDecoder(r.Body).Decode(&tempCity)
		)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(w, "400 bad request\n" + "%s\n", err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}(r.Body)

		log.Printf("Got %s city with area of %d miles.\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("201 created"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("405 method not allowed\n"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	var (
		PORT = ":8000"
		mainLogicHandler = http.HandlerFunc(mainLogic)
	)
	http.Handle("/cities", filterContentType(setServerTimeCookie(mainLogicHandler)))
	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
