package main

// import k8s to gen deployment,service,ingress,namespace struct
// import go-restful to gen api
import (
	"fmt"
	"context"
	"time"
	// "sync"
	// "bytes"
	// "text/template"
	// k8s "k8s.io/client-go/kubernetes"
	// "gopkg.in/yaml.v2"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// corev1 "k8s.io/client-go/kubernetes/typed/core/v1" // namespace service
	coreV1 "k8s.io/api/core/v1" // namespace service
	appsV1 "k8s.io/api/apps/v1" // deployment
	"k8s.io/apimachinery/pkg/util/intstr"
	// appv1 "k8s.io/client-go/kubernetes/typed/apps/v1" // deployment
	// "k8s.io/client-go/kubernetes/typed/networking/v1beta1" // ingress
)

// 其实没啥用,只是为了方便打印struct信息
type CNNamespace struct {
	Name string
}

type CNDeployment struct {
	DeployName       string
	Namespace        string
	ReplicasCount    int
	PodImage         string
	PodName          string
	ArgsListValue    []string
	EnvMap           map[string]string
	ContainerPortMap map[string]int32
	Requests         map[string]string
	Limits           map[string]string
	MountPathList    []string
	PVCName          string
}

type CNService struct {
	Namespace     string
	ServiceName   string
	PodName       string
	ExposePortPod int
	ExposePortSvc int
	NodePort      int
}


type CNApplication struct {
	Username string
	ApplicatioName string
	k8sns *coreV1.Namespace
	k8sdeploy *appsV1.Deployment
	k8ssvc *coreV1.Service
	nsstate string // namespace state :  error delete pending running
	deploystate string
	svcstate string
}

type CNApplications map[string]*CNApplication // key is Appid



// GenerateYAML generate yaml from template and data
// Init function will init k8s deployment,service,ingress,namespace struct

// queue , item is CNApplication TODO: use redis queue
var APPqueue = make(chan *CNApplication, 1000)



func main() {
	// 1-init connect to k8s 
	// 2-gen namespace/deploy/svc information
	// 3-create k8s namespace struct
	// 3-create k8s deployment struct
	// 3-create k8s service struct
	// 4-apply k8s namespace/deploy/svc
	// 5-wait for k8s namespace/deploy/svc running
	// 6-delete k8s svc/deploy/namespace


	fmt.Println("init CNtemplate")

	// Todo connect to redis
	

	// 1-init connect to k8s
	// 将生成的k8s namespace/svc/deploy/ingress 对象部署到k8s集群
	k8sconfig, err := clientcmd.BuildConfigFromFlags("", "./config.yaml")
	if err != nil {
		panic(err)
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		panic(err)
	}	
	fmt.Println(clientset)

	// 2-gen namespace/deploy/svc information
	// from function args
	// 变量username ,作为CNnamespace struct's name
	username := "test"
	appname := "testapp"
	// imagename := "nginx:latest"
	imagename := "httpd:latest"
	containerPortMap := map[string]int32{"http": int32(80)}
	replicasCount := 1
	requests := map[string]string{"cpu": "100m", "memory": "100Mi"}
	limits := map[string]string{"cpu": "200m", "memory": "200Mi"}

	exposePortPod := 80
	exposePortSvc := 80 // from nodeport map[] gen one  对于用户来说相当于随机端口30000-50000 
	nodeport := 40080

	tmpns := CNNamespace{
		Name: username,
	}

	tmpdeploy := CNDeployment{
		DeployName: appname,
		Namespace: username,
		PodName: appname,
		ReplicasCount: replicasCount,
		PodImage: imagename,
		// ArgsListValue: []string{},
		// EnvMap: map[string]string{},
		ContainerPortMap: containerPortMap,
		Requests: requests,
		Limits: limits,
		// MountPathList: []string{},
		// PVCName: "",
	}

	tmpservice := CNService{
		Namespace: username,
		ServiceName: appname,
		PodName: appname,
		ExposePortPod: exposePortPod,
		ExposePortSvc: exposePortSvc,
		NodePort: nodeport,
	}

	// 3-create namespace and apply to k8s
	tmpk8sns := Createnamespace(tmpns)
	err = Applynamespace(&tmpk8sns,clientset)
	if err != nil {
		fmt.Println(err)
	}

	// 3-create deployment and apply to k8s
	tmpk8sdeploy := CreateDeployment(tmpdeploy)
	err = ApplyDeployment(&tmpk8sdeploy,clientset)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", tmpk8sdeploy)
	
	// 3-create service and apply to k8s
	tmpk8ssvc := CreateService(tmpservice)
	err = ApplyService(&tmpk8ssvc,clientset)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", tmpk8ssvc)


	// sleep 30s
	fmt.Println("sleep 30s")
	time.Sleep(time.Second*30)

	// delete all k8s object just now created
	err = DeleteService(&tmpk8ssvc,clientset)
	if err != nil {
		fmt.Println(err)
	}

	err = DeleteDeployment(&tmpk8sdeploy,clientset)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sleep 1s")
	time.Sleep(time.Second)	
	err = Deletenamespace(&tmpk8sns,clientset)
	if err != nil {
		fmt.Println(err)
	}
	
}

// create ns,参数为k8s namespace struct 和 k8s clientset ,返回值为k8s namespace struct
// func Createnamespace(ns string,clientset kubernetes.Interface){
// func Createnamespace(ns string,clientset kubernetes.Interface) corev1.Namespace{
// func Createnamespace(ns CNNamespace,clientset kubernetes.Interface) coreV1.Namespace{
func Createnamespace(ns CNNamespace) coreV1.Namespace{
	fmt.Println("create namespace")
	k8sns := &coreV1.Namespace{
		ObjectMeta: metaV1.ObjectMeta{
			Name: ns.Name,
		},
	}
	return *k8sns
}

func Applynamespace(k8sns *coreV1.Namespace, clientset kubernetes.Interface) error {
	fmt.Println("apply namespace")
	_, err := clientset.CoreV1().Namespaces().Create(context.Background(),k8sns,metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("create namespace error")
		return err
	}
	return nil
}

func Deletenamespace(k8sns *coreV1.Namespace, clientset kubernetes.Interface) error {
	fmt.Println("delete namespace")
	err := clientset.CoreV1().Namespaces().Delete(context.Background(),k8sns.Name,metaV1.DeleteOptions{})
		if err != nil {
			fmt.Println("delete namespace error")
			return err
		}
	return nil
	// }
}

func CreateDeployment(tmpdeploy CNDeployment) appsV1.Deployment{
	fmt.Println("create deployment")
	var replicas int32 = int32(tmpdeploy.ReplicasCount)

	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
				Name: tmpdeploy.DeployName,
				Labels: map[string]string{
						"app": tmpdeploy.DeployName,
				},
				Namespace: tmpdeploy.Namespace,
		},
		Spec: appsV1.DeploymentSpec{
				Replicas: &replicas,   //这里要传递的是指针
				Selector: &metaV1.LabelSelector{
						MatchLabels: map[string]string{
								"app": tmpdeploy.DeployName,
						},
				},
				Template: coreV1.PodTemplateSpec{
						ObjectMeta: metaV1.ObjectMeta{
								Name: tmpdeploy.PodName,
								Labels: map[string]string{
										"app": tmpdeploy.PodName,								
								},
						},
						Spec: coreV1.PodSpec{
								Containers: []coreV1.Container{
										{
												Name:  tmpdeploy.PodName,
												Image: tmpdeploy.PodImage,
												Ports: []coreV1.ContainerPort{
														{
																Name:          "http",
																Protocol:      coreV1.ProtocolTCP,
																// ContainerPort: int32(tmpdeploy.ContainerPortMap["http"]),
																ContainerPort: tmpdeploy.ContainerPortMap["http"],
														},
												},
										},
								},
						},
				},
		},
	}
	return *deployment
}

func ApplyDeployment(k8sdeploy *appsV1.Deployment,clientset kubernetes.Interface) error {
	fmt.Println("apply deployment")
	_, err := clientset.AppsV1().Deployments(k8sdeploy.Namespace).Create(context.Background(),k8sdeploy,metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("create deployment error")
		return err
	}
	return nil
}

func UpdateDeployment(k8sdeploy *appsV1.Deployment,clientset kubernetes.Interface) error{
	fmt.Println("update deployment")
	_, err := clientset.AppsV1().Deployments(k8sdeploy.Namespace).Update(context.Background(),k8sdeploy,metaV1.UpdateOptions{})
	if err != nil {
		fmt.Println("update deployment error")
		return err
	}
	return nil
}

func DeleteDeployment(k8sdeploy *appsV1.Deployment,clientset kubernetes.Interface) error {
	fmt.Println("delete deployment")
	err := clientset.AppsV1().Deployments(k8sdeploy.Namespace).Delete(context.Background(),k8sdeploy.Name,metaV1.DeleteOptions{})
	if err != nil {
		fmt.Println("delete deployment error")
		return err
	}
	return nil
}

func CreateService(tmpsvc CNService) coreV1.Service{
	fmt.Println("create service")
	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
				Name: tmpsvc.ServiceName,
				Labels: map[string]string{
						"app": tmpsvc.PodName,
				},
				Namespace: tmpsvc.Namespace,
		},
		Spec: coreV1.ServiceSpec{
				Type: coreV1.ServiceTypeNodePort,
				Selector: map[string]string{
						"app": tmpsvc.PodName,
				},
				Ports: []coreV1.ServicePort{
						{
								Name: "http",
								Port: int32(tmpsvc.ExposePortSvc),
								TargetPort: intstr.FromInt(int(tmpsvc.ExposePortPod)),
								Protocol: coreV1.ProtocolTCP,
								NodePort: int32(tmpsvc.NodePort),
						},
				},
		},
	}
	return *service
}

func ApplyService(k8ssvc *coreV1.Service,clientset kubernetes.Interface) error {
	fmt.Println("apply service")
	_, err := clientset.CoreV1().Services(k8ssvc.Namespace).Create(context.Background(),k8ssvc,metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("create service error")
		return err
	}
	return nil
}

func UpdateService(k8ssvc *coreV1.Service,clientset kubernetes.Interface) error{
	fmt.Println("update service")
	_, err := clientset.CoreV1().Services(k8ssvc.Namespace).Update(context.Background(),k8ssvc,metaV1.UpdateOptions{})
	if err != nil {
		fmt.Println("update service error")
		return err
	}
	return nil
}

func DeleteService(k8ssvc *coreV1.Service,clientset kubernetes.Interface) error {
	fmt.Println("delete service")
	err := clientset.CoreV1().Services(k8ssvc.Namespace).Delete(context.Background(),k8ssvc.Name,metaV1.DeleteOptions{})
	if err != nil {
		fmt.Println("delete service error")
		return err
	}
	return nil
}

// func Gendemond(appname string, appimage string, appport int32, appdomain string, appexposeport int32){
// 	fmt.Println("gen demond")
// 	// gen namespace struct
// 	tmpns.Name = appname
// 	// gen deployment struct
// 	tmpdeploy.Name = appname
// 	tmpdeploy.Spec.Template.Spec.Containers[0].Name = appname
// 	tmpdeploy.Spec.Template.Spec.Containers[0].Image = appimage
// 	tmpdeploy.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort = appport
// 	// gen service struct
// 	tmpsvc.Name = appname
// 	tmpsvc.Spec.Ports[0].Port = appport
// 	// gen ingress struct
// 	tmping.Name = appname
// 	tmping.Spec.Rules[0].Host = appdomain
// 	tmping.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName = appname
// 	tmping.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntVal = appexposeport
// }