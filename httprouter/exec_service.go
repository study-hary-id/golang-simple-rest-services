package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

// getCommandOutput executes system commands and returns the results as a string.
func getCommandOutput(command string, arguments ...string) string {
	var (
		out    bytes.Buffer
		stderr bytes.Buffer
		cmd    = exec.Command(command, arguments...)
	)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.String()
}

// goVersion serves the version of Go installed on the machine.
func goVersion(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprintf(w, getCommandOutput(
		"/usr/local/go/bin/go", // The location of go binary
		"version",
	))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getFileContent shows the content within file which is requested.
func getFileContent(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	_, err := fmt.Fprintf(w, getCommandOutput(
		"/bin/cat", // The location of cat (program that show text file) binary.
		"../data/"+params.ByName("name"),
	))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var (
		router = httprouter.New()
		PORT   = ":8000"
	)

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(":8000", router))
}
