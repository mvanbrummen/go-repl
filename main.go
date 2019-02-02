package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend"))

	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
