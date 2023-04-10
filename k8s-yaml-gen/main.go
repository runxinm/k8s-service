package main

func main() {
	nsYaml := generateNamespaceYaml("my-namespace")
	DeploymentYaml := GenerateDeploymentYAML("my-deployment", "my-namespace", 3, "my-pod-image", "my-pod", "arg1 arg2 arg3", map[string]string{"ENV_KEY": "ENV_VALUE"}, []int{80, 443}, map[string]int{"http": 80, "https": 443}, map[string]string{"memory": "1Gi", "cpu": "1"}, map[string]string{"memory": "2Gi", "cpu": "2"}, []string{"/mnt/data"}, "my-pvc")
	svcYaml := generateServiceYaml("my-namespace", "my-service", "my-pod", 8080, 80)
	ingressYaml := generateIngressYaml("my-namespace", "my-ingress", "example.com", "my-app-ingress", 80)

	// 打印生成的 YAML
	fmt.Println(string(nsYaml))
	fmt.Println(string(DeploymentYaml))
	fmt.Println(string(svcYaml))
	fmt.Println(string(ingressYaml))

	// // 将 YAML 字符串转换为 Go 结构体
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
	// fmt.Printf("%+v\n", namespace)

}