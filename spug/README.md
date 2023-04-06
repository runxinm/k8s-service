## 介绍
使用 Python+Vue 实现的开源运维平台，前后端分离方便二次开发。该项目基于 Docker 镜像发布部署，方便安装和升级。支持运维常见功能：主机管理、任务计划管理、发布部署、监控告警等
## 参考链接
https://github.com/openspug/spug
将官网从docker compose 方式部署改为使用 k8s方式

## 要求
k8s集群要有默认存储类,本集群使用 openebs


## 注意事项
初次使用时，需要使用
`kubectl exec -it <pod_name> init_spug admin spug.dev`
命令创建用户，admin 和 spug.dev

成功之后结果如下
```
Operations to perform:
  Apply all migrations: account, alarm, app, config, deploy, exec, home, host, monitor, notify, repository, schedule, setting
Running migrations:
  Applying account.0001_initial... OK
  Applying alarm.0001_initial... OK
  Applying config.0001_initial... OK
  Applying app.0001_initial... OK
  Applying repository.0001_initial... OK
  Applying deploy.0001_initial... OK
  Applying exec.0001_initial... OK
  Applying home.0001_initial... OK
  Applying host.0001_initial... OK
  Applying monitor.0001_initial... OK
  Applying notify.0001_initial... OK
  Applying schedule.0001_initial... OK
  Applying setting.0001_initial... OK
初始化/更新成功
创建用户成功
```

之后可以使用admin  spug.dev进行登录,创建新的用户和角色


## TODO
初始化pod时直接初始化创建admin用户(目前通过在yaml文件中增加args项并没有成功进行初始化，需进一步排查原因)
