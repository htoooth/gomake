package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
)

type Task struct {
	Main              string
	Depends, Packages []string
}

func main() {
	const (
		GoPath     = "GOPATH"
		OptionFile = "gmkfile"
	)

	pwd, _ := os.Getwd()
	os.Setenv(GoPath, pwd)

	option, _ := os.Open(OptionFile)

	// gomake istalll
	// gomake build
	// cli have tow options

	t := parseOption(option)
	getDependPackage(t.Depends)
	buildCustomPackage(t.Packages)
	buildMain(t.Main)
}

func getDependPackage(pkgs []string) {
	for _, v := range pkgs {
		cmd := exec.Command("go", "get", v)
		cmd.Run()
	}
}

func buildCustomPackage(pkgs []string) {
	os.Chdir("src")
	for _, v := range pkgs {
		os.Chdir(v)
		cmd := exec.Command("go", "install")
		cmd.Run()
		os.Chdir("..")
	}
	os.Chdir("..")
}

func parseOption(f *os.File) *Task {
	var t Task
	dec := json.NewDecoder(f)

	if err := dec.Decode(&t); err == io.EOF {
		return nil
	} else if err != nil {
		log.Fatal(err)
	}

	return &t
}

func buildMain(app string) {
	os.Chdir("src")
	os.Chdir(app)

	cmd := exec.Command("go", "install")
	cmd.Run()

	os.Chdir("..")
	os.Chdir("..")
}
