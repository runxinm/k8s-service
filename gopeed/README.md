# 介绍
Gopeed（全称 Go Speed），是一款由Golang+Flutter开发的高速下载器，支持（HTTP、BitTorrent、Magnet）协议下载，并且支持全平台使用。


# 部署

`kubectl apply -f .`

docker部署命令
```
docker run -d -p 9999:9999 -v /path/to/download:/download liwei2633/gopeed
```

# 注意事项
尽量多分配些资源

# 参考链接
[官网链接](https://github.com/GopeedLab/gopeed)