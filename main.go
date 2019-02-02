package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const (
	port       = ":8080"
	staticPath = "./frontend/build"
	tmpDir     = "./.gorepl"
)

type CodeRequest struct {
	Code string `json:"code"`
}

type CodeResponse struct {
	Result string `json:"result"`
}

func NewCodeResponse(result string) *CodeResponse {
	return &CodeResponse{
		Result: result,
	}
}

func runCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req CodeRequest
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	tmpFile := fmt.Sprintf("%s/main.go", tmpDir)

	err = ioutil.WriteFile(tmpFile, []byte(req.Code), 0644)
	if err != nil {
		panic(err)
	}

	out, _ := exec.Command("go", "run", tmpFile).CombinedOutput()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NewCodeResponse(string(out)))
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NewCodeResponse(string(out)))
}

func main() {
	os.Mkdir(tmpDir, os.ModePerm)

	fs := http.FileServer(http.Dir(staticPath))

	http.Handle("/", fs)
	http.HandleFunc("/code", runCode)
	http.HandleFunc("/version", getVersion)

	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
