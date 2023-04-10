package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
	"gopkg.in/yaml.v2"
)


func applyYaml(yaml []byte) {
	tmpfile, err := ioutil.TempFile("", "temp-yaml-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(yaml); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	cmd := exec.Command("kubectl", "apply", "-f", tmpfile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
