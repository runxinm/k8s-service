# 目的
一些可以基于k8s部署的项目。
来源：
- 从docker run 改变为 k8s的yaml，自行编写yaml文件。
- 从docker-compose改变为k8s的yaml，修改kompose转换成的yaml文件。
- 自己开发的小玩意

# k8s-service-List

- solo
  -  41001
  - 个人博客-支持静态和动态
- spug
  -  41002
  - 基于vue和python的运维平台，支持主机管理、监控告警等


# TODO
- [gitlab](#)
  - 代码仓库

- [docker-register](#)
  - docker私有仓库

- [openfaas](#)
  - faas平台

- [homer](#)
  - 主页

- seafile
  - 文件网盘

- vaultwarden
  - 密码管理

- send
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
  - 代码仓库
  
- [Nodejs](#)
  - 代码仓库

- templete
  - 模板提取
    - pvc.yaml     存储(目前只支持数据存储)
    - deployment.yaml   app
    - service.yaml   服务暴露
    - config.yaml (后期)