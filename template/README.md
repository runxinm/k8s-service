# 模板填充

- 1-命名空间
  - namespace
    - ns-name(用户名)
- 2-存储空间
  - configMap('可多个')
    - config-name(用户名+app+config+index)
    - ns-name(用户名)
    - data(map,key为配置项名称，value为配置项值)
  - pvc('可多个')
    - pvc-name(用户名+app+pvc+index)
    - namespace(用户名)
- 3-deployment
  - replicas-count (副本数量,默认为1)
  - deploy-name(应用名称)
  - pod-name(pod名称)
  - args-list (暂时省略)
  - env-map  (暂时省略)
  - app-image (应用镜像名称 暂时考虑单pod)
  - containerPortmap
  - requests-map
  - limits-map
  - mountpath-list  (暂时省略)
  - pvc-name  (暂时省略)
- 4-service
  - svc-name(应用名称-service)
  - ns-name
  - pod-name (用于service select后端的pod)
    - 更进一步的,使用用户自定义的key-value 填充 selector 结构体
  - expose-port-pod
  - expose-port-service

- 5-ingress 
  - ExternalName
    - app-ingress-externalname (deploy-name)
    - svc-name
    - ns-name
  - Ingress
    - ingress-name (ns-name + deploy-name)
    - rules
      - domain-name (用于7层负载均衡)
      - app-ingress-externalname
      - expose-port-service
      - 进一步的详细rule参考nginx.
        - path
        - pathType
        - https时需要提供额外的认证

# 所需参数 (Input)


- 用户名
  - 用户名称
    - 生成命名空间
    - 标识应用归属

- 应用名称
  - deployment-name
  - pod信息(pod)
    - 镜像名称(pod-image)
    - args-list (暂时省略)
    - env-list (暂时省略)
    - containerPort(list) (容器需要暴露的端口 port)
      - 例如
        - mysql 服务需要/默认暴露 3306 端口，
        - redis (默认)暴露 6379 端口,
        - 其他自开发的应用,可能需要暴露的端口.
      - 可能需要暴露多个端口
    - resources( requests + limits )
    - mountPath(暂时省略)

- (+域名)
  - 用于配置7层负载均衡
  - 无域名,则使用  svc-name.ns-name.svc.cluster.local