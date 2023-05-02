# 介绍
Privacy-respecting, hackable metasearch engine

# 部署

`kubectl apply -f .`

docker部署命令
```
docker run -d --name=searxng \
-p 9147:8080 \
-v /volume1/docker/searxng:/etc/searxng \
--restart always \
searxng/searxng
```

# 注意事项

# 参考链接
[官网链接](https://github.com/searxng/searxng)