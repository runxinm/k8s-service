// package CNtemplate

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"text/template"
// 	"gopkg.in/yaml.v2"
// )


// func GenerateIngressYaml(nsName string, ingressName string, domainName string, appIngressExternalName string, exposePortSvc int32) []byte {
// 	ingressTemplate := `
// kind: Service
// apiVersion: v1
// metadata:
//   name: {{ .AppIngressExternalName }}
//   namespace: ingress-nginx
// spec:
//   type: ExternalName
//   externalName: {{ .Name }}.{{ .Namespace }}.svc.cluster.local
// ---
// apiVersion: networking.k8s.io/v1
// kind: Ingress
// metadata:
//   name: {{ .Name }}
//   namespace: ingress-nginx
// spec:
//   ingressClassName: nginx
//   rules:
//   - host: {{ .DomainName }}
//     http:
//       paths:
//       - backend:
//           service:
//             name: {{ .AppIngressExternalName }}
//             port:
//               number: {{ .ExposePortSvc }}
//         path: /
//         pathType: Prefix
// `
// 	tmpl, err := template.New("ingress").Parse(ingressTemplate)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var ingress Ingress
// 	ingress.Namespace = nsName
// 	ingress.Name = ingressName
// 	ingress.DomainName = domainName
// 	ingress.AppIngressName = appIngressExternalName
// 	ingress.ExposePortSvc = exposePortSvc

// 	var result []byte
// 	err = tmpl.Execute(&result, ingress)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result
// }
