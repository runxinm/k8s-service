# 介绍
speedtest测速

每小时运行一次速度测试检查并绘制结果图表。后端使用Laravel编写，前端使用React。它使用Ookla 的 speedtest cli包获取数据并使用Chart.js绘制结果

帮助您监视您的网络速度，并提供历史记录和数据可视化功能。

具体来说，henrywhitaker3/Speedtest-Tracker包括以下组件：

Speedtest测速工具：这个工具会定期运行测速测试，包括下载速度、上传速度和延迟。它可以帮助您了解您的网络速度和稳定性。

InfluxDB数据库：这个数据库用于存储您的测速结果和相关的元数据。它可以帮助您跟踪您的网络速度变化，并提供数据分析和可视化功能。



# 部署

`kubectl apply -f .`

docker命令
```
docker run \
--name=speedtest \
-p 8765:80 \
-v /volume1/docker/speedtest:/config \
-e OOKLA_EULA_GDPR=true \
--restart unless-stopped \
henrywhitaker3/speedtest-tracker
```

# 注意事项
可能比较慢或者卡，适当多分配一些资源

# 参考链接
[开源代码](https://github.com/henrywhitaker3/Speedtest-Tracker)