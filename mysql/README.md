# 介绍
mysql:5.7
# 部署

`kubectl apply -f .`

docker部署命令
```
docker run --restart=always --name=mysql5.7 -p 3306:3306 \
--mount type=bind,src=/home/data/mysql/conf/my.cnf,dst=/etc/my.cnf \
--mount type=bind,src=/home/data/mysql/data,dst=/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

# 注意事项
#### 连接数据库
```
mysql -h 10.160.100.101 -P 31005 -u root -p
```

#### 创建用户
创建用户test_user ,可以从任意ip连接数据库，密码11111
```
grant all privileges on *.* to 'user1'@'%' identified by '111111';
```

#### 查看
```
SELECT user, host, authentication_string FROM mysql.user;
```

结果
```
 SELECT user, host, authentication_string FROM mysql.user;
+---------------+-----------+-------------------------------------------+
| user          | host      | authentication_string                     |
+---------------+-----------+-------------------------------------------+
| root          | localhost | *2470C0C06DEE42FD1xxxxxxxxxxxxxxxxxxxxxxx |
| mysql.session | localhost | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE |
| mysql.sys     | localhost | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE |
| root          | %         | *2470C0C06DEE42FD16xxxxxxxxxxxxxxxxxxxxxx |
| user1         | %         | *FD571203974BA9AFE2xxxxxxxxxxxxxxxxxxxxxx |
+---------------+-----------+-------------------------------------------+
```

# 参考链接
[mysql容器化部署](https://blog.csdn.net/NeverLate_gogogo/article/details/114406303)