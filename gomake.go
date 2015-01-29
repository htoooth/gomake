// Copyright 2015 One Off Code
/*
Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

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
