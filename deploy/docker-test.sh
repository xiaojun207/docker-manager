
containerName="docker-manager"

docker ps -aq --filter "name=${containerName}" | grep -q . && docker stop ${containerName} && docker rm -fv ${containerName}
docker run -d --name docker-manager -p 8068:8068 -e driveName=mysql -e dataSourceUrl='root:Abc123@(192.168.3.67:3306)/docker-manager?charset=utf8' xiaojun207/docker-manager:latest

