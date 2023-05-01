# 介绍
aria2 是一个用于下载文件的实用程序。支持的协议有 HTTP(S)、FTP、SFTP、BitTorrent 和 Metalink。aria2 可以从多个来源/协议下载文件并尝试利用您的最大下载带宽。它支持同时从 HTTP(S)/FTP/SFTP 和 BitTorrent 下载文件，同时从 HTTP(S)/FTP/SFTP 下载的数据上传到 BitTorrent swarm。使用 Metalink 的块校验和，aria2 在下载 BitTorrent 等文件时自动验证数据块.
# 部署

`kubectl apply -f .`

# 注意事项
1. 需要在AriaNg设置中 修改ip和端口为实际使用的，否则会出现未连接错误
2. 资源可以适当增加

测试是否可以下载,随便找一个链接
```
magnet:?xt=urn:btih:74de3c8487a4fb6c7480c505f6d99f192545576b&dn=%e9%98%b3%e5%85%89%e7%94%b5%e5%bd%b1dy.ygdy8.com.%e6%91%87%e6%9b%b3%e9%9c%b2%e8%90%a5%e5%89%a7%e5%9c%ba%e7%89%88.2022.BD.1080P.%e6%97%a5%e8%af%ad%e4%b8%ad%e5%ad%97.mkv&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2fexodus.desync.com%3a6969%2fannounce
```

3. ip:port/rclone 可以使用rclone

默认账号user 默认密码password

4. ip:port/files 可以使用文件管理器

默认账号admin 默认密码admin


# 参考链接
[项目官方代码](https://github.com/aria2/aria2)
