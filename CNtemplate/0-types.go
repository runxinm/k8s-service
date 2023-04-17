package CNtemplate

type Namespace struct {
	Name string
}

type Service struct {
	Namespace     string
	Name          string
	PodName       string
	ExposePortPod int32
	ExposePortSvc int32
}

type Ingress struct {
	Namespace         string
	Name              string
	DomainName        string
	AppIngressName    string
	ExposePortSvc     int32
}