## 介绍
k8s 的 nginx-ingress 控制器
- nginx-ingress控制器的启动
- Ingress
- 修改配置文件  使得其支持服务发现

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
