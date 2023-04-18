package main

// import k8s to gen deployment,service,ingress,namespace struct
// import go-restful to gen api
import (
	"fmt"
	"context"
	"time"
	// "bytes"
	// "text/template"
	// k8s "k8s.io/client-go/kubernetes"
	// "gopkg.in/yaml.v2"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// corev1 "k8s.io/client-go/kubernetes/typed/core/v1" // namespace service
	corev1 "k8s.io/api/core/v1"
	// appv1 "k8s.io/client-go/kubernetes/typed/apps/v1" // deployment
	// "k8s.io/client-go/kubernetes/typed/networking/v1beta1" // ingress
)

// GenerateYAML generate yaml from template and data
// Init function will init k8s deployment,service,ingress,namespace struct
func main() {
	fmt.Println("init CNtemplate")
	// gen k8s namespace struct
	k8sns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "123",
		},
	}
	// tmpns := corev1.Namespace{} // undefined: corev1.Namespace
	// 打印namespace struct 信息,并转换为yaml类型再次打印
	fmt.Printf("%+v\n", k8sns)
	// nayaml,err := yaml.Marshal(k8sns)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(nayaml))

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
	
	// create namespace
	Createnamespace(k8sns,clientset)
	
	// create deployment
	// _, err = clientset.AppsV1().Deployments("default").Create(k8sdeploy)
	// fmt.Printf("%+v\n", nayaml)
	// tmpdeploy = v1.Deployment{}
	// tmpsvc = v1.Service{}
	// tmping = v1beta1.Ingress{}
}

// create ns,参数为k8s namespace struct 和 k8s clientset
func Createnamespace(k8sns *corev1.Namespace,clientset kubernetes.Interface){
	fmt.Println("create namespace")
	// 循环10次创建namespace
	for i := 0; i < 50; i++ {
		k8sns.Name = fmt.Sprintf("test-%d",i)
		_, err := clientset.CoreV1().Namespaces().Create(context.Background(),k8sns,metav1.CreateOptions{})
		// _, err = clientset.CoreV1().Namespaces().Create(k8sns)
		if err != nil {
			fmt.Println("create namespace error")
		}
	}
	// sleep 10s
	fmt.Println("sleep 10s")
	time.Sleep(10 * time.Second)
	
	// delete namespace
	fmt.Println("delete namespace")
	for i := 0; i < 50; i++ {
		k8sns.Name = fmt.Sprintf("test-%d",i)
		err := clientset.CoreV1().Namespaces().Delete(context.Background(),k8sns.Name,metav1.DeleteOptions{})
		if err != nil {
			fmt.Println("delete namespace error")
		}
	}
}

func Deletenamespace(){
	fmt.Println("delete namespace")
}

func CreateDeployment(){
	fmt.Println("create deployment")
}

func UpdateDeployment(){
	fmt.Println("update deployment")
}

func DeleteDeployment(){
	fmt.Println("delete deployment")
}

func CreateService(){
	fmt.Println("create service")
}

func UpdateService(){
	fmt.Println("update service")
}

func DeleteService(){	
	fmt.Println("delete service")
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