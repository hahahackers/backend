package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "ping\n")
	if err != nil {
		return
	}
}

func chgk(w http.ResponseWriter, req *http.Request) {
	tourID := strings.TrimPrefix(req.URL.Path, "/chgk/")
	_, _ = fmt.Fprintf(w, "tourID: %s\n", tourID)
}

func main() {

	println("Starting server on port 80")

	http.HandleFunc("/ping", hello)
	http.HandleFunc("/chgk/", chgk)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		println("Error starting server: ", err.Error())
		return
	}
}
