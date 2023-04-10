package main

import (
	"fmt"
	"io/ioutil"
	"text/template"
	"gopkg.in/yaml.v2"
)


func generateNamespaceYaml(nsName string) []byte {
	nsTemplate := `
apiVersion: v1
kind: Namespace
metadata:
  name: {{ .Name }}
`
	tmpl, err := template.New("namespace").Parse(nsTemplate)
	if err != nil {
		panic(err)
	}

	var ns Namespace
	ns.Name = nsName

	var result []byte
	err = tmpl.Execute(&result, ns)
	if err != nil {
		panic(err)
	}

	return result
}