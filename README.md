[中文说明](#zh) | 
<a name="en">English</a>
## Foreword
Because some of our servers are distributed in different network environments. And, these server programs use docker for publishing, but do not use k8s such tools. In a discrete management state.
I've also considered tools like rancher, but it's still too heavy for us. I need a lighter management tool. Therefore, docker manager and docker agent were born.

## Function description
* Overview, including: number of servers, number of running containers, total number of containers, number of applications, number of distributed tasks, and number of real-time logs opened.
* Display of all containers, including container name, ID, server, image used, port and creation time.
* You can start container, delete container, restart container.
* Container status information, including CPU usage, memory usage, and network usage.
* Publish a new container to the target server.
* The container real-time log (if exists) is like to: docker logs - F -- tail 10 containername. It is more resource consuming. It is better to only view the log temporarily (this function does not support cluster deployment).
* Display of server assets, mainly including: total number of containers, number of running containers, CPU usage, memory usage, docker version, and whether docker agent is online (this function does not support cluster deployment).
* User management, administrator and docker agent account, password and status management.


## docker-manager
Docker based multi host container web management, data is stored as sqlit3 by default
It can also be mysql. You only need to configure the database connection parameters, and the database tables will be created and updated automatically.

## Start

```
docker pull xiaojun207/docker-manager:latest

docker ps -aq --filter "name=docker-manager" | grep -q . && docker stop docker-manager && docker rm -fv docker-manager

docker run -d --name docker-manager -p 8068:8068 -e driveName=mysql -e dataSourceUrl='root:Abc123@(dbhost:3306)/dbname?charset=utf8' xiaojun207/docker-manager:latest

```

Parameter Description:

Parameter | required | default value | description
---|----------|---------------|--- 
driveName | no       | sqlite3  | The default is "sqlite3", or "mysql". If MySQL is used, the "datasourceurl" must be configured.
dataSourceUrl | no       | -     | Database connection URL, <br>when driveName is "sqlite3"，"dataSourceUrl" default：data/database.db（/app/data/database.db），<br>when driveName is "mysql"，the "dataSourceUrl" is required，such as：-e dataSourceUrl='root:Abc123@(dbhost:3306)/dbname?charset=utf8' 
useCache | no       | false         | whether to enable local cache. It can be enabled in stand-alone deployment, but not in cluster deployment


## Login account
Upon initial startup, the program will automatically create an administrator account (admin), a client account (agent, password, token), and a user name and password, which will be printed into the log output. (only displayed once, please make a backup)

## Client Agent
```
docker pull xiaojun207/docker-agent:latest

docker run -d --name docker-agent -v /var/run/docker.sock:/var/run/docker.sock -e DockerServer="http://192.168.1.200:8068/dockerMgrApi/agent" -e Token="12345678" xiaojun207/docker-agent:latest

```
It should be used in conjunction with "xiaojun207/docker-agent" image. Please refer to the description for the specific usage of docker agent

Special note:
    The hostname of each server (the host of docker agent) must be unique

## Contact email
If you have any ideas or suggestions, please send an email to the following email:

email: xiaojun207@126.com


## 中文说明

<a name="zh">中文说明</a> | [English](#en)

## 前言
由于我们的一部分服务器，分布在不同的公共网络环境。而且，这些服务器程序，发布都使用了docker，但并没有使用k8s这样的工具。处于离散的管理状态。
也考虑过rancher这样的工具，但对于我们来说还是太重了。我需要一个更轻量的管理工具。因此docker-manager和docker-agent就诞生了。

## 功能简述
* 总览，包括：服务器数量、运行容器数量、容器总数量、应用数量、下发任务数量、实时日志开启数量
* 所有容器展示，包括：容器名称，id、所在服务器、使用镜像、端口、创建时间
* 启动容器、删除容器、重新启动
* 容器状态信息，包括：cpu使用、内存使用、网络使用
* 发布新的容器到目标服务器
* 容器实时日志（如果有的话），相当于docker logs -f --tail 10 容器名，比较耗资源，仅临时查看日志用比较好（该功能不支持集群部署）
* 服务器资产展示，主要包括：容器总数量、运行容器数量、cpu使用、内存使用、docker版本、docker-agent是否在线（该功能不支持集群部署）
* 用户管理，对管理员和docker-agent账号、密码、状态管理
*

## docker-manager
基于docker的多主机容器web管理，数据默认存储为sqlit3
也可以是mysql，你只需配置好数据库连接参数，数据库表会自动创建和更新。


## 使用方法

```
docker pull xiaojun207/docker-manager:latest

docker ps -aq --filter "name=docker-manager" | grep -q . && docker stop docker-manager && docker rm -fv docker-manager

docker run -d --name docker-manager -p 8068:8068 -e driveName=mysql -e dataSourceUrl='root:Abc123@(dbhost:3306)/dbname?charset=utf8' xiaojun207/docker-manager:latest

```

参数说明:

参数 | 是否必填 | 默认值     | 说明
---|------|---------|--- 
driveName | 否    | sqlite3 | 也可以是mysql，如果是mysql，则dataSourceUrl必须配置
dataSourceUrl | 否    | -       | 数据库连接url<br>，当driveName为sqlite3时，dataSourceUrl默认为：data/database.db（即/app/data/database.db），<br>当driveName为mysql时，dataSourceUrl则必填，例如：-e dataSourceUrl='root:Abc123@(dbhost:3306)/dbname?charset=utf8'
useCache | 否    | false   | 是否启用本地缓存，单机部署的时候启用，集群部署请不要启用


## 登录账号
初次启动，程序会自动创建管理员账号(admin)、客户端账号(agent, 密码即TOKEN)，用户名密码，会打印到日志输出中。（仅显示一次，请做好备份）

## 客户端
```
docker pull xiaojun207/docker-agent:latest

docker run -d --name docker-agent -v /var/run/docker.sock:/var/run/docker.sock -e DockerServer="http://192.168.1.200:8068/dockerMgrApi/agent" -e Token="12345678" xiaojun207/docker-agent:latest

```
需配合xiaojun207/docker-agent镜像使用，docker-agent的具体使用方法，请参见其说明

特别说明:
    每台服务器(docker-agent的宿主机)的hostname，必须唯一

## 联系邮箱
如果，你有什么想法或建议，请你发送邮件到下面的邮箱：

email: xiaojun207@126.com
