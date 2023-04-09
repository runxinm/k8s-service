## 介绍
solo 开源，方便/快速的搭建个人博客

## 注意事项
除此使用时需要进入到mysql中，添加solo数据库 类似于下列命令
create database solo
create database solo DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

## 原始mysql命令
```
docker run -itd  --name mysql-master \
-v /data/master/data/mysql-master:/var/lib/mysql \
-v /data/master/master:/etc/mysql/conf.d \
-e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 mysql
```

```
-v 参数将mysql的数据以及配置文件持久化，挂载到本地
-e 参数给数据库赋值密码与用户
-p 参数映射3306端口
```

## 参考链接
从docker迁移到k8s[solo博客](https://huaweicloud.csdn.net/63310f9cd3efff3090b50c87.html?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Eactivity-1-122017149-blog-128852169.235%5Ev27%5Epc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Eactivity-1-122017149-blog-128852169.235%5Ev27%5Epc_relevant_default&utm_relevant_index=2#devmenu8)