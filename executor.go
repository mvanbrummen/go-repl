package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
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

func (e *GoExecutor) Execute(code string) (string, error) {
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

func (e *RubyExecutor) Execute(code string) (string, error) {
	tmpFile, err := writeTmpFile(code, "tmp.rb")

	if err != nil {
		return "", err
	}
	return executeCommand("ruby", tmpFile)
}

func (e *RubyExecutor) Version() (string, error) {
	return executeCommand("ruby", "-v")
}

func executeCommand(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		return "", err
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
