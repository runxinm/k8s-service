package main

import (
	"context"
	"fmt"
	"io/ioutil"

	
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 创建 Kubernetes 客户端
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 读取 YAML 文件
	yamlFile, err := ioutil.ReadFile("path/to/your/yaml/file.yaml")
	if err != nil {
		panic(err)
	}

	// 应用 YAML
	if err := applyYaml(clientset, yamlFile); err != nil {
		panic(err)
	}
}

func applyYaml(clientset *kubernetes.Clientset, yamlFile []byte) error {
	// 解析 YAML 文件
	obj, err := parseYaml(yamlFile)
	if err != nil {
		return err
	}

	// 获取对象的 GroupVersionKind 信息
	gvk := obj.GetObjectKind().GroupVersionKind()

	// 根据 GroupVersionKind 创建资源接口
	resource := schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: gvk.Kind + "s",
	}

	// 获取命名空间
	namespace, _, err := unstructured.NestedString(obj.Object, "metadata", "namespace")
	if err != nil {
		return err
	}

	// 获取资源接口
	resourceIf := clientset.Resource(resource).Namespace(namespace)

	// 应用 YAML
	_, err = resourceIf.Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("%s/%s created\n", gvk.Kind, obj.GetName())
	return nil
}

func parseYaml(yamlFile []byte) (*unstructured.Unstructured, error) {
	// 解析 YAML
	obj := &unstructured.Unstructured{}
	err := yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		return nil, err
	}

	// 设置默认值
	err = setDefaultValues(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func setDefaultValues(obj *unstructured.Unstructured) error {
	// 设置默认命名空间
	if obj.GetNamespace() == "" {
		obj.SetNamespace("default")
	}

	// 设置默认标签
	labels := obj.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	labels["app"] = obj.GetName()
	obj.SetLabels(labels)

	return nil
}