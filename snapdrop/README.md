# 介绍

# 部署

`kubectl apply -f .`

docker部署命令
```
docker run -d \
-e PUID=1026 \
-e PGID=101 \
-e TZ=Asia/Shanghai \
-p 9471:80 \
-v /volume1/docker/snapdrop:/config \
linuxserver/snapdrop:latest
```

# 注意事项

测试结果
![snapdrop_result](./result/snapdrop_result.png)

# 参考链接
[官网链接]()