## 来源
由于我们的一部分服务器，分布在不同的公共网络环境。而且，这些服务器程序，发布都使用了docker，但并没有使用k8s这样的工具。处于离散的管理状态。
也考虑过rancher这样的工具，但对于我们来说还是太重了。我需要一个更轻量的管理工具。因此docker-manager和docker-agent就诞生了。

## docker-manager
基于docker的多主机容器web管理，依赖mysql，你只需配置好数据库连接参数，数据库表会自动创建和更新。

## 使用方法

```
docker pull xiaojun207/docker-manager:1.0.5

docker ps -aq --filter "name=docker-manager" | grep -q . && docker stop docker-manager && docker rm -fv docker-manager

docker run -d --name docker-manager -p 8068:8068 -e driveName=mysql -e dataSourceUrl='root:Abc123@(dbhost:3306)/dbname?charset=utf8' xiaojun207/docker-manager:1.0.5

```

## 登录账号
初次启动，程序会自动创建管理员账号，用户名密码，会打印到日志输出中。（仅显示一次，请做好备份）

## 客户端
需配合xiaojun207/docker-agent镜像使用，docker-agent的具体使用方法，请参见其说明

## 联系邮箱
email: xiaojun207@126.com
