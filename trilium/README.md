# 介绍
一个分层笔记应用程序，专注于构建大型个人知识库。

# 部署

`kubectl apply -f .`

docker方式
```
docker run -d \
--name trilium-cn \
-p 18080:8080 \
-v /volume1/docker/trilium-data:/root/trilium-data \
-e TRILIUM_DATA_DIR=/root/trilium-data \
--restart always \
nriver/trilium-cn
```

# 注意事项

# 参考链接
[官网](https://github.com/zadam/trilium)