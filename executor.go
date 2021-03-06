package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const tmpDir = "./.gorepl"

type Executor interface {
	Execute(code string) (string, error)
	Version() (string, error)
}

type GoExecutor struct{}

func NewGoExecutor() *GoExecutor {
	return &GoExecutor{}
}

func (*GoExecutor) Execute(code string) (string, error) {
	tmpFile, err := writeTmpFile(code, "main.go")

	if err != nil {
		return "", err
	}
	return executeCommand("go", "run", tmpFile)
}

func (e *GoExecutor) Version() (string, error) {
	return executeCommand("go", "version")
}

type RubyExecutor struct{}

func NewRubyExecutor() *RubyExecutor {
	return &RubyExecutor{}
}

func (*RubyExecutor) Execute(code string) (string, error) {
	tmpFile, err := writeTmpFile(code, "tmp.rb")

	if err != nil {
		return "", err
	}
	return executeCommand("ruby", tmpFile)
}

func (*RubyExecutor) Version() (string, error) {
	return executeCommand("ruby", "-v")
}

type JavascriptExecutor struct{}

func NewJavascriptExecutor() *JavascriptExecutor {
	return &JavascriptExecutor{}
}

func (*JavascriptExecutor) Execute(code string) (string, error) {
	tmpFile, err := writeTmpFile(code, "tmp.js")

	if err != nil {
		return "", err
	}
	return executeCommand("node", tmpFile)
}

func (*JavascriptExecutor) Version() (string, error) {
	return executeCommand("node", "--version")
}

type PythonExecutor struct{}

func NewPythonExecutor() *PythonExecutor {
	return &PythonExecutor{}
}

func (*PythonExecutor) Execute(code string) (string, error) {
	tmpFile, err := writeTmpFile(code, "tmp.py")

	if err != nil {
		return "", err
	}
	return executeCommand("python", tmpFile)
}

func (*PythonExecutor) Version() (string, error) {
	return executeCommand("python", "-V")
}

type JavaExecutor struct{}

func NewJavaExecutor() *JavaExecutor {
	return &JavaExecutor{}
}

func (*JavaExecutor) Execute(code string) (string, error) {
	tmpFile, err := writeTmpFile(code, "Main.java")

	if err != nil {
		return "", err
	}
	_, err = executeCommand("javac", tmpFile)
	if err != nil {
		return "", err
	}
	return executeCommand("java", "-cp", tmpDir, "Main")
}

func (*JavaExecutor) Version() (string, error) {
	return executeCommand("java", "-version")
}

func executeCommand(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		log.Error(err)
	}

	return string(out), nil
}

func writeTmpFile(code string, filename string) (string, error) {
	tmpFile := fmt.Sprintf("%s/%s", tmpDir, filename)

	err := ioutil.WriteFile(tmpFile, []byte(code), 0644)
	if err != nil {
		return "", err
	}

	return tmpFile, nil
}
