package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func generateDeploymentYaml(deployName string, nsName string, podName string, podImage string, replicasCount int, containerPort int, argsList []string, envMap map[string]string, requests map[string]string, limits map[string]string, mountpathList []string, pvcName string) []byte {
	deploymentTemplate := `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: %s
  name: %s
  namespace: %s
spec:
  replicas: %d
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
        - args:
            %s
          env:
            - name: %s
              value: %s
          image: %s
          name: %s
          ports:
            - containerPort: %d
          resources:
            requests:
              memory: %s
              cpu: %s
            limits:
              memory: %s
              cpu: %s
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: %s
              name: %s
      volumes:
        - name: %s
          persistentVolumeClaim:
            claimName: %s
`
	deploymentYaml := fmt.Sprintf(deploymentTemplate, deployName, deployName, nsName, replicasCount, deployName, podName, joinArgs(argsList), getEnvMapKey(envMap), getEnvMapValue(envMap), podImage, podName, containerPort, requests["memory"], requests["cpu"], limits["memory"], limits["cpu"], joinMountPathList(mountpathList), pvcName, pvcName, pvcName)

	return []byte(deploymentYaml)
}

func joinArgs(argsList []string) string {
	return fmt.Sprintf("'%s'", joinList(argsList, "' '"))
}

func joinMountPathList(mountpathList []string) string {
	return joinList(mountpathList, ",")
}

func joinList(list []string, sep string) string {
	if len(list) == 0 {
		return ""
	}
	return sep + sep.Join(list) + sep
}

func getEnvMapKey(envMap map[string]string) string {
	for k, _ := range envMap {
		return k
	}
	return ""
}

func getEnvMapValue(envMap map[string]string) string {
	for _, v := range envMap {
		return v
	}
	return ""
}

func main() {
	deployName := "my-deployment"
	nsName := "my-namespace"
	podName := "my-pod"
	podImage := "my-pod-image:latest"
	replicasCount := 3
	containerPort := 8080
	argsList := []string{"arg1", "arg2", "arg3"}
	envMap := map[string]string{"ENV_KEY": "ENV_VALUE"}
	requests := map[string]string{"memory": "1Gi", "cpu": "1"}
	limits := map[string]string{"memory": "2Gi", "cpu": "2"}
	mountpathList := []string{"/mnt/data", "/mnt/config"}
	pvcName := "my-pvc"

	deploymentYaml := generateDeploymentYaml(deployName, nsName, podName, podImage, replicasCount, containerPort, argsList, envMap, requests, limits, mountpathList, pvcName)

	// 打印生成的 Deployment YAML
	fmt.Println(string(deploymentYaml))

	// 将 YAML 字符串转换为 Go 结构体
	var deployment interface{}
	err := yaml.Unmarshal(deploymentYaml, &deployment)
	if err != nil {
		panic(err)
	}

	// 打印生成的 Deployment 对象
	fmt.Printf("%+v\n", deployment)
}

// generateDeploymentYaml 函数接受多个参数，包括 Deployment 的名称、命名空间、Pod 的名称、镜像名称、副本数、容器端口、命令行参数列表、环境变量映射、资源限制、挂载路径列表和 PVC 的名称。该函数使用这些参数填充 Deployment 的 YAML 模板，并返回生成的 YAML 字符串。