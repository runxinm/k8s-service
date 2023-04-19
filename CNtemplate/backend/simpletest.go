package main

import (
        "context"
        "fmt"
		"time"
		"log"
		"strconv" //整数转化字符串

		appsV1 "k8s.io/api/apps/v1"
        coreV1 "k8s.io/api/core/v1"
        metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"   //重命名
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/tools/clientcmd"

)

func main(){
	list_one()
	list_multi()
	configPath := "etc/C1_config"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	clientset, _ := kubernetes.NewForConfig(config)
	for true{
		createDeploy(clientset)
		time.Sleep(time.Duration(15)*time.Second)

		createService(clientset)
		time.Sleep(time.Duration(15)*time.Second)

		editDeployByClient(clientset,1)
		time.Sleep(time.Duration(15)*time.Second)
		
		editReplicasByClient(clientset,2)
		time.Sleep(time.Duration(15)*time.Second)
		
		deleteResource(clientset)
		time.Sleep(time.Duration(15)*time.Second)
	}
}
// 一个集群
func list_one() {
        //fmt.Println("client-go")
        configPath := "etc/C1_config" //etc/C1_config  //使用API编程k8s，需要kubeconfig文件来验证
        config, err := clientcmd.BuildConfigFromFlags("", configPath)   //使用clientcmd的BuildConfigFromFlags，加载配置文件生成Config对象，该对象包含apiserver的地址和用户名和密码和token等，""那里是值masterurl，就是完整路径使用，这里将文件放到main.go同一个的cg目录下，直接空就行，然后configPath写etc/kube.conf即可，config用来变成使用api操作集群验证用
        if err != nil {
                log.Fatal(err)   //接收err这个error类型作为参数，只开启Fatal级别的日志，日志级别有LOG DEBUG INFO WARN ERROR FATAL
        }
        clientset, err := kubernetes.NewForConfig(config)   //使用kubernetes包下的NewForConfig来初始化clientset中的每个client，参数是配置文件，返回clientset对象的，是多个client的集合，就像是kubernetes的客户端
        fmt.Println(clientset)
		if err != nil {
                log.Fatal(err)
        }
        nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("node:")
        for _, node := range nodeList.Items {
                fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
                        node.Name,
                        node.Status.Addresses,
                        node.Status.Phase,
                        node.Status.NodeInfo.OSImage,
                        node.Status.NodeInfo.KubeletVersion,
                        node.Status.NodeInfo.OperatingSystem,
                        node.Status.NodeInfo.Architecture,
                        node.CreationTimestamp,
                )
        }
        fmt.Println("namespace")
        namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
        for _, namespace := range namespaceList.Items {
                fmt.Println(namespace.Name, namespace.CreationTimestamp, namespace.Status.Phase)
        }
        serviceList, _ := clientset.CoreV1().Services("").List(context.TODO(), metaV1.ListOptions{})
        fmt.Println("services")
        for _, service := range serviceList.Items {
                fmt.Println(service.Name, service.Spec.Type, service.CreationTimestamp, service.Spec.Ports, service.Spec.ClusterIP)

        }
        deploymentList, _ := clientset.AppsV1().Deployments("default").List(context.TODO(), metaV1.ListOptions{})
        fmt.Println("deployment")
        for _, deployment := range deploymentList.Items {
                fmt.Println(deployment.Name, deployment.Namespace, deployment.CreationTimestamp, deployment.Labels, deployment.Spec.Selector.MatchLabels, deployment.Status.Replicas, deployment.Status.AvailableReplicas)

        }

}

// 多个集群
func list_multi() {
	//fmt.Println("client-go")
	var all_client [2]*kubernetes.Clientset

	for i:=0; i< 2;i++ {
		tmp := strconv.Itoa(i+1)
		configPath := "etc/C" + tmp + "_config"                       //使用API编程k8s，需要kubeconfig文件来验证
		config, err := clientcmd.BuildConfigFromFlags("", configPath) //使用clientcmd的BuildConfigFromFlags，加载配置文件生成Config对象，该对象包含apiserver的地址和用户名和密码和token等，""那里是值masterurl，就是完整路径使用，这里将文件放到main.go同一个的cg目录下，直接空就行，然后configPath写etc/kube.conf即可，config用来变成使用api操作集群验证用
		// fmt.Println(reflect.TypeOf(config))
		if err != nil {
			log.Fatal(err) //接收err这个error类型作为参数，只开启Fatal级别的日志，日志级别有LOG DEBUG INFO WARN ERROR FATAL
		}
		clientset, err := kubernetes.NewForConfig(config) //使用kubernetes包下的NewForConfig来初始化clientset中的每个client，参数是配置文件，返回clientset对象的，是多个client的集合，就像是kubernetes的客户端
		all_client[i] = clientset
		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println(reflect.TypeOf(clientset))
	for key, clientset := range all_client {
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("Cluster%d info :",key)
		nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("*********************************")
		fmt.Println("node:")
		for _, node := range nodeList.Items {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				node.Name,
				node.Status.Addresses,
				node.Status.Phase,
				node.Status.NodeInfo.OSImage,
				node.Status.NodeInfo.KubeletVersion,
				node.Status.NodeInfo.OperatingSystem,
				node.Status.NodeInfo.Architecture,
				node.CreationTimestamp,
			)
		}
		fmt.Println("*********************************")
		fmt.Println("namespace")
		namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
		for _, namespace := range namespaceList.Items {
			fmt.Println(namespace.Name, namespace.CreationTimestamp, namespace.Status.Phase)
		}
		serviceList, _ := clientset.CoreV1().Services("").List(context.TODO(), metaV1.ListOptions{})
		fmt.Println("*********************************")
		fmt.Println("services")
		for _, service := range serviceList.Items {
			fmt.Println(service.Name, service.Spec.Type, service.CreationTimestamp, service.Spec.Ports, service.Spec.ClusterIP)

		}
		deploymentList, _ := clientset.AppsV1().Deployments("default").List(context.TODO(), metaV1.ListOptions{})
		fmt.Println("*********************************")
		fmt.Println("deployment")
		for _, deployment := range deploymentList.Items {
			fmt.Println(deployment.Name, deployment.Namespace, deployment.CreationTimestamp, deployment.Labels, deployment.Spec.Selector.MatchLabels, deployment.Status.Replicas, deployment.Status.AvailableReplicas)

		}
	}
}

// func createDeploy() {
func createDeploy(clientset *kubernetes.Clientset) {
		// configPath := "etc/C1_config"
		fmt.Println("Deploy start")
		// fmt.Println(configPath)
        namespace := "default"
        var replicas int32 = 3
        deployment := &appsV1.Deployment{
                ObjectMeta: metaV1.ObjectMeta{
                        Name: "nginx",
                        Labels: map[string]string{
                                "app": "nginx",
                                "env": "dev",
                        },
                },
                Spec: appsV1.DeploymentSpec{
                        Replicas: &replicas,   //这里要传递的是指针
                        Selector: &metaV1.LabelSelector{
                                MatchLabels: map[string]string{
                                        "app": "nginx",
                                        "env": "dev",
                                },
                        },
                        Template: coreV1.PodTemplateSpec{
                                ObjectMeta: metaV1.ObjectMeta{
                                        Name: "nginx",
                                        Labels: map[string]string{
                                                "app": "nginx",
                                                "env": "dev",
                                        },
                                },
                                Spec: coreV1.PodSpec{
                                        Containers: []coreV1.Container{
                                                {
                                                        Name:  "nginx",
                                                        Image: "nginx:1.20.1",
                                                        Ports: []coreV1.ContainerPort{
                                                                {
                                                                        Name:          "http",
                                                                        Protocol:      coreV1.ProtocolTCP,
                                                                        ContainerPort: 80,
                                                                },
                                                                {
                                                                        Name:          "https",
                                                                        Protocol:      coreV1.ProtocolTCP,
                                                                        ContainerPort: 443,
                                                                },
                                                        },
                                                },
                                        },
                                },
                        },
                },
        }
        deployment, _ = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metaV1.CreateOptions{})
        fmt.Println(deployment)
}

// func createService(configPath string) {
func createService(clientset *kubernetes.Clientset) {
	// configPath := "etc/kube.conf"
	// configPath := "etc/C1_config"
	// config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	// clientset, _ := kubernetes.NewForConfig(config)
	fmt.Println("sevice start")
	namespace := "default"
	service := &coreV1.Service{
			ObjectMeta: metaV1.ObjectMeta{
					Name: "nginx-service",
					Labels: map[string]string{
							"env": "dev",
					},
			},
			Spec: coreV1.ServiceSpec{
					Type: coreV1.ServiceTypeNodePort,
					Selector: map[string]string{
							"env": "dev",
							"app": "nginx",
					},
					Ports: []coreV1.ServicePort{
							{
									Name:     "http",
									Port:     80,
									Protocol: coreV1.ProtocolTCP,
							},
					},
			},
	}
	service, _ = clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metaV1.CreateOptions{})
	fmt.Println(service)
	fmt.Println("sevice finish")

}


// 更新镜像
func editDeploy(configPath string,replicas int32) {
        // configPath := "etc/kube.conf"
		fmt.Println("***********************************************************")
		fmt.Println("update deployment:start")
		config, _ := clientcmd.BuildConfigFromFlags("", configPath)
        clientset, _ := kubernetes.NewForConfig(config)
        namespace:="default"
        // var replicas int32=1
        name:="nginx"
        deployment, _ := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
        
		deployment.Spec.Replicas=&replicas
        deployment.Spec.Template.Spec.Containers[0].Image="nginx:1.10"
        deployment,err:=clientset.AppsV1().Deployments(namespace).Update(context.TODO(),deployment,metaV1.UpdateOptions{})

        fmt.Println(deployment,err)
        fmt.Println("update deployment:finish")
		fmt.Println("***********************************************************")
}

func editDeployByClient(clientset *kubernetes.Clientset,replicas int32) {
	// configPath := "etc/kube.conf"
	fmt.Println("***********************************************************")
	fmt.Println("update deployment:start")
	namespace:="default"
	// var replicas int32=1
	name:="nginx"
	deployment, _ := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	
	deployment.Spec.Replicas=&replicas
	deployment.Spec.Template.Spec.Containers[0].Image="nginx:1.10"
	deployment,err:=clientset.AppsV1().Deployments(namespace).Update(context.TODO(),deployment,metaV1.UpdateOptions{})

	fmt.Println(deployment,err)
	fmt.Println("update deployment:finish")
	fmt.Println("***********************************************************")
}

// 更新副本数
func editReplicas(configPath string,replicas int32) {
		fmt.Println("***********************************************************")
		fmt.Println("update deployment-replicas:start")
        config, _ := clientcmd.BuildConfigFromFlags("", configPath)
        clientset, _ := kubernetes.NewForConfig(config)
        // var replicas int32=1
        name:="nginx"
        namespace:="default"
        scale, _ := clientset.AppsV1().Deployments(namespace).GetScale(context.TODO(),name, metaV1.GetOptions{})
        // fmt.Println(scale)
		scale.Spec.Replicas=2
        scale,err:=clientset.AppsV1().Deployments(namespace).UpdateScale(context.TODO(),name,scale,metaV1.UpdateOptions{})
        fmt.Println(scale,err)
		fmt.Println("update deployment-replicas:finish")
		fmt.Println("***********************************************************")
}

func editReplicasByClient(clientset *kubernetes.Clientset,replicas int32) {
	fmt.Println("***********************************************************")
	fmt.Println("update deployment-replicas:start")
	// var replicas int32=1
	name:="nginx"
	namespace:="default"
	scale, _ := clientset.AppsV1().Deployments(namespace).GetScale(context.TODO(),name, metaV1.GetOptions{})
	// fmt.Println(scale)
	scale.Spec.Replicas=replicas
	scale,err:=clientset.AppsV1().Deployments(namespace).UpdateScale(context.TODO(),name,scale,metaV1.UpdateOptions{})
	fmt.Println(scale,err)
	fmt.Println("update deployment-replicas:finish")
	fmt.Println("***********************************************************")


}


func deleteResource(clientset *kubernetes.Clientset) {
	// configPath := "etc/C1_config"
	// config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	// clientset, _ := kubernetes.NewForConfig(config)
	name,servicename:="nginx","nginx-service"
	namespace:="default"
	clientset.AppsV1().Deployments(namespace).Delete(context.TODO(),name, metaV1.DeleteOptions{})
	clientset.CoreV1().Services(namespace).Delete(context.TODO(),servicename,metaV1.DeleteOptions{})
	//clientset.ExtensionsV1beta1().Ingress()
	fmt.Println("delete Finish")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}
