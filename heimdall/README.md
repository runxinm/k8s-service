# 介绍
Application Dashboard 

应用程序仪表板和启动器

Heimdall Application Dashboard 是您所有 Web 应用程序的仪表板。不过，它不需要局限于应用程序，您可以添加指向任何您喜欢的内容的链接。


# 部署

`kubectl apply -f .`

docker部署命令
```
docker run -d \
  --name=heimdall \
  -e PUID=1026 \
  -e PGID=100 \
  -e TZ=Asia/Shanghai \
  -p 9780:80 \
  -p 9443:443 \
  -v /volume1/docker/heimdall:/config \
  --restart unless-stopped \
  linuxserver/heimdall:latest
```

# 注意事项
直接使用时,https证书存在问题


# 参考链接
[官网](https://heimdall.site)
[开源代码](https://github.com/linuxserver/Heimdall)