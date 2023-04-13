package main

import(
	"fmt"
	// "bytes"
	// "text/template"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"CNCCTC/CNtemplate"
)


type Person struct {
    Name string
    Age  int
}


func main() {
	data := Person{Name: "Alice", Age: 30}
	tmpl := "name: {{.Name}}\nage: {{.Age}}\n"
	result, err := CNtemplate.GenerateYAML(data, tmpl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// nsYaml := CNtemplate.GenerateNamespaceYaml("my-namespace")
	// fmt.Println(string(nsYaml))

	DeploymentYaml,err := CNtemplate.GenerateDeploymentYAML("my-deployment", "my-namespace", 3, "my-pod-image", "my-pod", "arg1 arg2 arg3", map[string]string{"ENV_KEY": "ENV_VALUE"}, map[string]int{"http": 80, "https": 443}, map[string]string{"memory": "1Gi", "cpu": "1"}, map[string]string{"memory": "2Gi", "cpu": "2"}, []string{"/mnt/data"}, "my-pvc")
	// fmt.Println(string(DeploymentYaml))
	// fmt.Println(string(DeploymentYaml))

	// svcYaml := CNtemplate.GenerateServiceYaml("my-namespace", "my-service", "my-pod", 8080, 80)
	// fmt.Println(string(svcYaml))

	// ingressYaml := CNtemplate.GenerateIngressYaml("my-namespace", "my-ingress", "example.com", "my-app-ingress", 80)
	// fmt.Println(string(ingressYaml))

	// 打印生成的 YAML 字符串

	// // TODO 1 - 将 YAML 字符串转换为 Go 结构体
	// var namespace struct {
	// 	ApiVersion string `yaml:"apiVersion"`
	// 	Kind       string
	// 	Metadata   struct {
	// 		Name string
	// 	}
	// }
	// err := yaml.Unmarshal(nsYaml, &namespace)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("%+v\n", DeploymentYaml)

	// // TODO 2 - 部署该yaml
	// 
	k8sconfig, err := clientcmd.BuildConfigFromFlags("", "./config.yaml")
	if err != nil {
		log.Fatal(err)   
	}
	K8sclient, _ := kubernetes.NewForConfig(k8sconfig)
	fmt.Printf("%+v\n\t",K8sclient)


}