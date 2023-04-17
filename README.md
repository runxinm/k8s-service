# 介绍
一些可以基于k8s部署的项目。
来源：
- 从docker run 改变为 k8s的yaml，自行编写yaml文件。
- 从docker-compose改变为k8s的yaml，修改kompose转换成的yaml文件。
- 自己开发的

目的：
  方便后续同学学习/使用/部署/迁移
  提取k8s部署应用所需的模板、流程
  
# 服务暴露路径
前提是已经运行了ingress-nginx、sc存储类(可以自行搭建NFS并创建sc，也可以借助openebs)

docker镜像

--> 后端pod(deployment)

--> service 

--> ingress规则 

--> 写入到Ingress-nginx-controller配置文件并自动重载使更改生效 

--> 对Ingress-nginx创建service(自建集群使用NodePort方式，云集群一般可以使用Loadbalance方式)

--> 实现client无论通过哪个K8s节点的IP+端口都可以访问到后端pod

服务暴露方式也说明了部署应用时的顺序:
- 1-创建 namespace
- 2-(如果需要存储) 
  - 创建存储PVC,指定存储的大小，访问方式（单节点读、多节点读、多节点读写）
  - 创建Configmap(存放配置信息，或者通过pod的环境变量方式传递信息，推荐使用前者)
- 2-创建 要部署的pod的一些依赖，例如 其他pod(例如某些web需要使用mysql)
- 3-创建 deployment(pod)  
- 3-创建 service(clusterIP或者headless类型)
- 4-创建 Ingress(指定ingressclass名称和规则)
- 测试应用是否部署成功

# k8s-service-List(Done)

## 基础服务(核心)
- nfs-sc(存储后端)
  - 设置为默认存储类
  - 动态创建PV
  - ns:nfs-provisioner

- openebs(另一种存储后端)
  - 没有NFS时的简单替代
  - 只支持 RWO 单节点

- docker-registry
  - 31001
  - 私有docker镜像仓库(TODO增加用户权限)
  - ns:docker-registry

- Prometheus
  - 31002
  - Prometheus监控
  - ns:monitoring

- redis-cluster
  - 31003
  - redis集群
  - redis-cli -c -h 10.160.100.101 -p 31003

- redis单节点
  - 31004
  - 单机模式的redis
  - redis-cli -h 10.160.100.101 -p 31004
  - auth password

- cluster-ingress
  - 30080/30443
  - 基于nginx-ingress 的负载均衡,支持七层负载均衡，服务暴露。
  - nodeport暴露方式的一种替代，但并不完美，最好使用LB方式。
  - ns:ingress-nginx

## 应用
- solo
  - 41001
  - 个人博客-支持静态和动态
  - ns:solo
- spug
  -  41002
  - 基于vue和python的运维平台，支持主机管理、监控告警等
  - ns:spug
- memos
  - 41003
  - 具有知识管理和社交功能的开源自托管备忘录中心
  - ns:memos
- vaultwarden
  - 41004
  - 密码管理
  - ns:vw

## 注意事项
如果使用ingress，需要自行配置DNS解析到ingress-controller

# TODO
## k8s-service-List(TODO)
<!-- - [prometheus](#)
  - prometheus监控组件，一系列监控指标 
  - 在之前基础上 增加ingress服务暴露 -->

- [gitlab](#)
  - 代码仓库

<!-- - [docker-register](#)
  - docker私有仓库 -->

- [openfaas](#)
  - faas平台

- [mysql](#)
  - mysql数据库
  - 存在一些bug

- [ETCD](#)
  - 数据库/注册中心

<!-- - [redis](#)
  - redis单节点 或 redis集群 -->

- [homer](#)
  - 主页

- [seafile](#)
  - 文件网盘

<!-- - [vaultwarden](#)
  - 密码管理 -->

- [send](#)
  - 另一个密码管理

- [drone](https://github.com/drone/drone)
  - 4100x
  - 基于 Docker 的持续集成平台，使用 Go 语言编写
  - Drone is a continuous delivery system built on container technology. Drone uses a simple YAML build file, to define and execute build pipelines inside Docker containers

- [zookeeper集群](https://www.jianshu.com/p/e0f9bfa6a998)
  - 41010

- [gochat](https://github.com/LockGit/gochat)
  - 一个基于go实现的轻量级im系统，具有docker
  - 需要改为使用 k8s 方式部署

- [discuz](https://zhuanlan.zhihu.com/p/398073277)
  - 类似于bbs？

- [ChatGPT web](#)
  - 开源ChatGPT

- [python3](#)
  - 运行环境
  
- [Nodejs](#)
  - 运行环境

## 项目
- templete
  - 模板提取
    - pvc.yaml     存储(目前只支持数据存储)
    - deployment.yaml   app
    - service.yaml   服务暴露
    - config.yaml (后期)
- Tools
  - 生成模板
    - gen-yaml.go
