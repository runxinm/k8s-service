package main

import (
	"fmt"
	"os"
	"text/template"
)

// TODO
// ArgsListValue修改
func GenerateDeploymentYAML(deployName string, nsName string, replicasCount int, podImage string, podName string, argsListValue string, envMap map[string]string, containerPortList []int, containerPortMap map[string]int, requests map[string]string, limits map[string]string, mountPathList []string, pvcName string) error {
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
          env:
          {{ range $key, $value := .EnvMap }}
            - name: {{ $key }}
              value: {{ $value }}
          {{ end }}
          image: {{ .PodImage }}
          name: {{ .PodName }}
          ports:
          {{ range $index, $value := .ContainerPortList }}
            - containerPort: {{ $value }}
          {{ end }}
          {{ range $key, $value := .ContainerPortMap }}
            - containerPort: {{ $value }}
              name: {{ $key }}
          {{ end }}
          resources:
            requests:
              memory: "{{ .Requests.Memory }}"
              cpu: "{{ .Requests.CPU }}"
            limits:
              memory: "{{ .Limits.Memory }}"
              cpu: "{{ .Limits.CPU }}"
          securityContext:
            privileged: true
          volumeMounts:
          {{ range $index, $value := .Mount[PathList](poe://www.poe.com/_api/key_phrase?phrase=PathList&prompt=Tell%20me%20more%20about%20PathList.) }}
            - mountPath: {{ $value }}
              name: {{ .PVCName }}
          {{ end }}
      volumes:
        - name: {{ .PVCName }}
          persistentVolumeClaim:
            claimName: {{ .PVCName }}
`

	// Parse the template.
	tmpl, err := template.New("deployment").Parse(deploymentYAML)
	if err != nil {
		return err
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
		ContainerPortList: containerPortList,
		ContainerPortMap: containerPortMap,
		Requests:         requests,
		Limits:           limits,
		MountPathList:    mountPathList,
		PVCName:          pvcName,
	}

	// Execute the template with the data and print the output.
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	return nil
}

// 这将生成以下 YAML 文件：
/*
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: my-deployment
  name: my-deployment
  namespace: my-namespace
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-deployment
  template:
    metadata:
      labels:
        app: my-pod
		spec:
      containers:
        - args:
          arg1 arg2 arg3
          env:
            - name: ENV_KEY
              value: ENV_VALUE
          image: my-pod-image
          name: my-pod
          ports:
            - containerPort: 80
            - containerPort: 443
              name: https
          resources:
            requests:
              memory: "1Gi"
              cpu: "1"
            limits:
              memory: "2Gi"
              cpu: "2"
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /mnt/data
              name: my-pvc
      volumes:
        - name: my-pvc
          persistentVolumeClaim:
            claimName: my-pvc
			*/