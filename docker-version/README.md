# 因简单/重复功能 而未使用k8s部署的一些应用
### 1.Camus
Camus 是一个群组视频聊天应用程序，它使用 WebRTC 进行直接的点对点通信。用户可以创建和加入房间、使用麦克风和网络摄像头传输音频和视频、共享屏幕以及发送短信.

`docker run -d -p 50001:5000 camuschat/camus:latest`

仅需要映射端口，但界面相对来说简陋，使用体验不好，已经两年未更新。

- [deployment.yaml](./camus.yaml)

[开源代码](https://github.com/camuschat/camus)

### 2.Zdir
使用Golang + Vue3开发的轻量级目录列表程序，支持Linux、Docker、Windows部署，支持视频、音频、代码等常规文件预览，适合个人或初创公司文件分享使用，亦可作为轻量级网盘使用

`docker run -d --name="zdir" -v /volume1/docker/zdir:/data/apps/zdir/data -v /volume1/docker/zdir:/data/apps/zdir/data/public -p 50002:6080 --restart=always helloz/zdir:latest`

- [pod.yaml](./zdir.yaml) (使用主机目录)

[开源代码](https://github.com/helloxz/zdir)


### 3.