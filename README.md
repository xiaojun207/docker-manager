## 来源
由于我们的一部分服务器，分布在不同的公共网络环境。而且，这些服务器程序，发布都使用了docker，但并没有使用k8s这样的工具。处于离散的管理状态。
也考虑过rancher这样的工具，但对于我们来说还是太重了。我需要一个更轻量的管理工具。因此docker-manager和docker-agent就诞生了。

## 功能简述
* 总览，包括：服务器数量、运行容器数量、容器总数量、应用数量、下发任务数量、实时日志开启数量
* 所有容器展示，包括：容器名称，id、所在服务器、使用镜像、端口、创建时间
* 启动容器、删除容器、重新启动
* 容器状态信息，包括：cpu使用、内存使用、网络使用
* 发布新的容器到目标服务器
* 容器实时日志（如果有的话），相当于docker logs -f --tail 10 容器名，比较耗资源，仅临时查看日志用比较好
* 服务器资产展示，主要包括：容器总数量、运行容器数量、cpu使用、内存使用、docker版本、docker-agent是否在线（该功能不支持集群部署）
* 用户管理，对管理员和docker-agent账号、密码、状态管理
* 

## docker-manager
基于docker的多主机容器web管理，依赖mysql，你只需配置好数据库连接参数，数据库表会自动创建和更新。

## 使用方法

```
docker pull xiaojun207/docker-manager:1.0.5

docker ps -aq --filter "name=docker-manager" | grep -q . && docker stop docker-manager && docker rm -fv docker-manager

docker run -d --name docker-manager -p 8068:8068 -e driveName=mysql -e dataSourceUrl='root:Abc123@(dbhost:3306)/dbname?charset=utf8' xiaojun207/docker-manager:1.0.5

```

## 登录账号
初次启动，程序会自动创建管理员账号(admin)、客户端账号(agent, 密码即TOKEN)，用户名密码，会打印到日志输出中。（仅显示一次，请做好备份）

## 客户端
```
docker pull xiaojun207/docker-agent:latest

docker run -d --name docker-agent -v /var/run/docker.sock:/var/run/docker.sock -e DockerServer="http://192.168.1.200:8068/dockerMgrApi" -e Token="12345678" xiaojun207/docker-agent:latest

```
需配合xiaojun207/docker-agent镜像使用，docker-agent的具体使用方法，请参见其说明

## 联系邮箱
如果，你有什么想法或建议，请你发送邮件到下面的邮箱：

email: xiaojun207@126.com
