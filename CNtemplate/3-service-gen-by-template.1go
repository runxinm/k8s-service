// package CNtemplate

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"text/template"
// 	"gopkg.in/yaml.v2"
// )


// func GenerateServiceYaml(nsName string, svcName string, podName string, exposePortPod int32, exposePortSvc int32) []byte {
// 	svcTemplate := `
// apiVersion: v1
// kind: Service
// metadata:
//   name: {{ .Name }}
//   namespace: {{ .Namespace }}
//   labels:
//     app: {{ .PodName }}
// spec:
//   ports:
//   - targetPort: {{ .ExposePortPod }}
//     port: {{ .ExposePortSvc }}
//     protocol: [TCP](poe://www.poe.com/_api/key_phrase?phrase=TCP&prompt=Tell%20me%20more%20about%20TCP.)
//   selector:
//     app: {{ .PodName }}
// `
// 	tmpl, err := template.New("service").Parse(svcTemplate)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var svc Service
// 	svc.Namespace = nsName
// 	svc.Name = svcName
// 	svc.PodName = podName
// 	svc.ExposePortPod = exposePortPod
// 	svc.ExposePortSvc = exposePortSvc

// 	var result []byte
// 	err = tmpl.Execute(&result, svc)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result
// }