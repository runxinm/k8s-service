package CNtemplate

import (
  "bytes"
	"fmt"
	// "os"
	"text/template"
)


func GenerateYAML(data interface{}, tmpl string) (string, error) {
  // 创建一个新的模板并解析模板字符串
  t, err := template.New("mytemplate").Parse(tmpl)
  if err != nil {
      return "", err
  }

  // 创建一个缓冲区
  var buf bytes.Buffer

  // 填充数据到模板并将结果写入缓冲区
  err = t.Execute(&buf, data)
  if err != nil {
      return "", err
  }

  // 将缓冲区中的内容转换为字符串并返回
  return buf.String(), nil
}

// TODO
// ArgsListValue修改
func GenerateDeploymentYAML(deployName string, nsName string, replicasCount int, podImage string, podName string, argsListValue string, envMap map[string]string, containerPortMap map[string]int, requests map[string]string, limits map[string]string, mountPathList []string, pvcName string) (string, error) {
	// Define the YAML template.
	deploymentYAML := `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: {{ .DeployName }}
  name: {{ .DeployName }}
  namespace: {{ .Namespace }}
spec:
  replicas: {{ .ReplicasCount }}
  selector:
    matchLabels:
      app: {{ .DeployName }}
  template:
    metadata:
      labels:
        app: {{ .PodName }}
    spec:
      containers:
        - args:
          {{ .ArgsListValue }}
          env: {{ range $key, $value := .EnvMap }}
            - name: {{ $key }}
              value: {{ $value }} {{ end }}
          image: {{ .PodImage }}
          name: {{ .PodName }}
          ports: {{ range $key, $value := .ContainerPortMap }}
            - containerPort: {{ $value }}
              name: {{ $key }} {{ end }}
          resources:
            requests:
              memory: "{{ .Requests.Memory }}"
              cpu: "{{ .Requests.CPU }}"
            limits:
              memory: "{{ .Limits.Memory }}"
              cpu: "{{ .Limits.CPU }}"
          securityContext:
            privileged: true
          volumeMounts:{{ range $index, $value := .MountPathList }}
            - mountPath: {{ $value }}
              name: pvc-{{ $index }}{{ end }}
      volumes:
        - name: {{ .PVCName }}
          persistentVolumeClaim:
            claimName: {{ .PVCName }}
`

	// Parse the template.
	tmpl, err := template.New("deployment").Parse(deploymentYAML)
  fmt.Println(err)
  fmt.Println(tmpl)
	if err != nil {
		return "", err
	}

	// Define the data to be passed to the template.
	data := struct {
		DeployName       string
		Namespace        string
		ReplicasCount    int
		PodImage         string
		PodName          string
		ArgsListValue    string
		EnvMap           map[string]string
		ContainerPortList []int
		ContainerPortMap map[string]int
		Requests         map[string]string
		Limits           map[string]string
		MountPathList    []string
		PVCName          string
	}{
		DeployName:       deployName,
		Namespace:        nsName,
		ReplicasCount:    replicasCount,
		PodImage:         podImage,
		PodName:          podName,
		ArgsListValue:    argsListValue,
		EnvMap:           envMap,
		// ContainerPortList: containerPortList,
		ContainerPortMap: containerPortMap,
		Requests:         requests,
		Limits:           limits,
		MountPathList:    mountPathList,
		PVCName:          pvcName,
	}

	// Execute the template with the data.
  
  // 创建一个缓冲区
  var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
  fmt.Println(tmpl)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 这将生成以下 YAML 文件：
/*

			*/