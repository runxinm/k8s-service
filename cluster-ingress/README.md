## 介绍
k8s 的 nginx-ingress 控制器
nginx-ingress的暴露方式是NodePort，所以配置DNS解析时，可以绑定集群后面的所有（具有公网IP的）node节点的IP。
- nginx-ingress控制器的启动
- Ingress
- 修改配置文件  使得其支持服务发现

## 服务暴露路径
前提是已经运行了ingress-nginx
--> 后端pod(deployment)
--> service 
--> ingress规则 
--> 写入到Ingress-nginx-controller配置文件并自动重载使更改生效 
--> 对Ingress-nginx创建service(自建集群使用NodePort方式，云集群一般可以使用Loadbalance方式)
--> 实现client无论通过哪个K8s节点的IP+端口都可以访问到后端pod

服务暴露方式也说明了部署应用时应有的顺序

## Ingress注意事项
修改service暴露方式,不能使用NodePort方式暴露，而是改用ClusterIP+ExternalName方式
如果pod/deployment和ingress在同一命名空间，则不需要使用externalname方式（ExternalName相当于在这一命名空间为另一命名空间的服务起了个别名）
rule规则中，host使用域名(APPID?) paths的backend使用 service-name service-port
## 各yaml文件作用和含义
- deployv1.2.yaml 官网1.2版本的nginx-ingress
- deployv1.7.yaml 官网1.7版本的nginx-ingress
- 0-namespaces.yaml create命名空间 nginx-ingress
- 1-nginx-ingress.yaml create RBAC/deployment/service等,修改暴露方式为NodePort,且固定在30080,30443管理http和https。
- 1-nginx-ingress-backup.yaml 直接使用hostnetwork的80和443方式，不推荐，测试可用，实际生产最好不这么使用，可能会造成网络混乱和其他不可控的问题。
