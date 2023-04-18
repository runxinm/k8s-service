package CNtemplate

// 需要根据用户提交的app名称和image名称，生成对应的yaml文件

// namespace --> deployment --> service --> ingress
// namespace.name --> 用户名,租户隔离
// deployment.name --> 用户提交的app名称
// service.name --> 用户提交的app名称
// ingress.name --> 用户提交的app名称
type Namespace struct {
	Name string
}

type Deployment struct {
	DeployName       string
	Namespace        string
	ReplicasCount    int
	PodImage         string
	PodName          string
	ArgsListValue    []string
	EnvMap           map[string]string
	ContainerPortMap map[string]int
	Requests         map[string]string
	Limits           map[string]string
	MountPathList    []string
	PVCName          string
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