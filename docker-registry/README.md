## 介绍
私有镜像仓库

## 注意事项
修改/etc/docker/daemon.json,增加insecure-registries
```
{
    "registry-mirrors": [
      "http://hub-mirror.c.163.com",
      "https://ustc-edu-cn.mirror.aliyuncs.com"
    ],
    "insecure-registries":["10.160.100.101:31001"],
    "live-restore":true,
    "exec-opts": ["native.cgroupdriver=systemd"]
}
```

## 测试
- 本地镜像
查看本地具有的nginx镜像 docker images
```
nginx                                                    latest    904b8cb13b93   6 weeks ago     142MB
```

- 重新打标签
```
docker tag 904b8cb13b93 10.160.100.101:31001/nginx-2
```

- 测试
```
docker push 10.160.100.101:31001/nginx-2

docker rmi 10.160.100.101:31001/nginx-2

docker images | grep "nginx"

docker pull 10.160.100.101:31001/nginx-2
```


也可以通过浏览器访问私有镜像仓库，通过http://[ip地址]:[端口]/v2/_catalog 访问

## TODO
设置用户权限，比如登录名和账号。

在k8s集群中拉取本地私有镜像时，也需要先创建Secret，然后才能够成功拉取自定义镜像。