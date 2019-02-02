package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	port       = ":8080"
	staticPath = "./frontend/build"
)

func main() {
	fs := http.FileServer(http.Dir(staticPath))

	http.Handle("/", fs)

	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
