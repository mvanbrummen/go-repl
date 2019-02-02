package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	port       = ":8080"
	staticPath = "./frontend/build"
)

type Language int

func (l Language) String() string {
	languages := []string{
		"golang",
		"ruby",
		"javascript",
		"python",
	}
	return languages[l]
}

const (
	Golang Language = iota
	Ruby
	JavaScript
	Python
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
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

	executor, err := getExecutor(req.Language)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, _ := executor.Execute(req.Code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NewCodeResponse(result))
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	executor, err := getExecutor(lang)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	version, err := executor.Version()
	if err != nil {
		log.Error(err)
	}
	json.NewEncoder(w).Encode(NewCodeResponse(version))
}

func getExecutor(lang string) (Executor, error) {
	var executor Executor

	switch lang {
	case Golang.String():
		executor = NewGoExecutor()
	case Ruby.String():
		executor = NewRubyExecutor()
	case JavaScript.String():
		executor = NewJavascriptExecutor()
	case Python.String():
		executor = NewPythonExecutor()
	default:
		return nil, fmt.Errorf("Unsupport lanaguage: %s", lang)
	}

	return executor, nil
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
