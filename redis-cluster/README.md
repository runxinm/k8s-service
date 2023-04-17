# redis集群模式

## 初始化集群
<!-- StatefulSet创建完毕后，可以看到6个pod已经启动了，但这时候整个redis集群还没有初始化，需要使用官方提供的redis-trib工具。它是官方提供的redis-cluster管理工具，可以实现redis集群的创建、更新等功能，在早期的redis版本中，它是以源码包里redis-trib.rb这个ruby脚本的方式来运作的，现在（我使用的5.0.3）已经被官方集成进redis-cli中。
我们当然可以在任意一个redis节点上运行对应的工具来初始化整个集群，但这么做显然有些不太合适，我们希望每个节点的职责尽可能地单一，所以最好单独起一个pod来运行整个集群的管理工具。
首先在k8s上创建一个ubuntu的pod，用来作为管理节点。
```
kubectl run -i --tty redis-cluster-manager --image=ubuntu --restart=Never /bin/bash
```
进入pod内部先安装一些工具，包括wget,dnsutils，然后下载和安装redis。 -->
### 查看 redis 副本
```
$ kg pod -n redis -o wide
NAME      READY   STATUS    RESTARTS   AGE   IP              NODE          NOMINATED NODE   READINESS GATES
redis-0   1/1     Running   0          26s   10.180.12.181   m-test-c1w1   <none>           <none>
redis-1   1/1     Running   0          23s   10.180.64.213   m-k8s-fedc1   <none>           <none>
redis-2   1/1     Running   0          21s   10.180.71.165   m-test-c1w2   <none>           <none>
redis-3   1/1     Running   0          18s   10.180.12.182   m-test-c1w1   <none>           <none>
redis-4   1/1     Running   0          15s   10.180.64.211   m-k8s-fedc1   <none>           <none>
redis-5   1/1     Running   0          12s   10.180.71.164   m-test-c1w2   <none>           <none>
```

### 进入到redis-0容器
```
kubectl exec -it redis-0 /bin/bash -n redis
```
### 使用0、2、4作为master，1、3、5作为slave
```
redis-cli --cluster create 10.180.12.181:6379 10.180.71.165:6379 10.180.64.211:6379
```
结果如下,注意上面的master节点，会生成对应节点id：xxx、xxx、xxx，用于创建slave节点。

```
···
>>> Performing Cluster Check (using node 10.180.12.181:6379)
M: 4f5b307xxxxx273d8a1bcf52f075e5 10.180.12.181:6379
   slots:[0-5460] (5461 slots) master
M: 5f041826xxxxx8f6f0cc846c13dd83 10.180.71.165:6379
   slots:[5461-10922] (5462 slots) master
M: 96f9d4f31xxxxf439698c41a62 10.180.64.211:6379
   slots:[10923-16383] (5461 slots) master
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
root@redis-0:/data# 
```

给master分别附加对应的Slave节点，这里的cluster-master-id在上一步创建的时候会给出
```
redis-cli --cluster add-node 10.180.64.213:6379 10.180.12.181:6379 --cluster-slave --cluster-master-id 4f5b30702afade1b455cbe273d8a1bcf52f075e5
redis-cli --cluster add-node 10.180.12.182:6379 10.180.71.165:6379 --cluster-slave --cluster-master-id 5f041826de74e7d7caf781a8f6f0cc846c13dd83
redis-cli --cluster add-node 10.180.71.164:6379 10.180.64.211:6379 --cluster-slave --cluster-master-id 96f9d4f31dfa185bba2ba3ed64cf439698c41a62
```

集群初始化后，随意进入一个节点检查一下集群信息
```
redis-cli -c
cluster info
cluster nodes
```

尝试在集群外部连接redis。集群模式下，-c不能缺失。
```
redis-cli -c -h 10.160.100.101 -p 31003
```

## 一些问题
Node 10.180.12.181:6379 replied with error:
ERR Client sent AUTH, but no password is set

原因：
```
这个错误信息表示一个客户端尝试连接到 Redis 服务器，位于 IP 地址为 10.180.12.181，端口为 6379，但是没有设置密码进行身份验证。

Redis 支持身份验证来保护访问服务器的安全性。当客户端连接到 Redis 时，它可以发送一个带有密码的 AUTH 命令来进行身份验证。在这种情况下，客户端发送了 AUTH 命令，但是 Redis 拒绝了身份验证尝试，因为没有设置密码。

要解决这个错误，您需要确保客户端使用 AUTH 命令发送正确的密码。如果您是 Redis 服务器的管理员，您可以通过修改 Redis 配置文件来设置密码，通常位于 /etc/redis/redis.conf，然后重新启动 Redis 服务。您可以通过添加或修改 "requirepass" 配置指令来设置密码，例如：

requirepass mypassword

将 "mypassword" 替换为实际要使用的密码。设置密码后，您需要更新客户端代码，以便使用 AUTH 命令发送正确的密码。
```

## 参考
[k8s搭建redis集群3主3从](https://juejin.cn/post/7202272345833914428#heading-8)
[k8s搭建redis集群](https://zhuanlan.zhihu.com/p/451935324)