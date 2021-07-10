package data

import (
	"docker-manager/model"
)

var (
	dbPath = "./db"

	// config
	Config = model.NewLevelDB(dbPath + "/config") // 客户端配置
	// docker
	Servers            = model.NewLevelDB(dbPath + "/servers")         // {serverName: {DockerInfo}}
	Containers         = model.NewLevelDB(dbPath + "/containers")      // {serverName: [{ContainerInfo}]}
	ContainerServerMap = model.NewLevelDB(dbPath + "/containerServer") // {ContainerId: serverName}
	Stats              = model.NewLevelDB(dbPath + "/stats")           // {ContainerId: {ContainerStatsInfo}}
	Images             = model.NewLevelDB(dbPath + "/images")          // {serverName: [{ImageInfo}]}
	// user
	SampleUser  = model.NewLevelDB(dbPath + "/users") // {user:password}
	SampleToken = model.NewLevelDB(dbPath + "/token") // {user:password}
	// task
	Task = model.NewLevelDB(dbPath + "/task") // {"id":{"id":"223456622", }, }
	Log  = model.NewLevelDB(dbPath + "/log")  // {containerId, stats}
	// App
	AppGroup = model.NewLevelDB(dbPath + "/appGroup") // {appname: [serverName]}
	AppInfos = model.NewLevelDB(dbPath + "/appInfos") // {appname: {Image: "", "ContainerName":"", Ports: [], Volumes: [], Mounts: [], env: []}}
)

func Close() {
	Config.Close()
	Servers.Close()
	Containers.Close()
	ContainerServerMap.Close()
	Stats.Close()
	Images.Close()
	SampleUser.Close()
	SampleToken.Close()
	Task.Close()
	Log.Close()
	AppGroup.Close()
	AppInfos.Close()
}

func init() {
	//Init("sqlite3", "./db/data.db")
	//InitDB("mysql", "root:Abc123@(127.0.0.1:3306)/docker-manager?charset=utf8")
	//DBEngine.Sync2(new(table.User), new(table.UserApi), new(table.Servers), new(table.Services))

	if Config.Size() == 0 {
		Config.Store("agentConfig", map[string]interface{}{
			"TaskFrequency": 5 * 60,
		})
	}

	if SampleUser.Size() == 0 {
		SampleUser.StoreStr("xiaojun", "b24543877bae1cc3")
		SampleUser.StoreStr("admin", "72d60f075a2341ab")
	}

	if SampleToken.Size() == 0 {
		SampleToken.StoreStr("agentToken", "a8ceec2ef7294f46f6fd4fea32c0aa4d6d76dd27c86b4f7c470ee5b568382057")
		SampleToken.StoreStr("apiToken", "41033229b540ca1b790b6ebb2bf9e889b3e96a6f2f28018da8857a275f89e60f")
	}

	//AppInfos.Store("docker-agent", map[string]interface{}{"Name": "docker-agent", "memory": 128*1024*1024, "Image": "xiaojun207/docker-agent:latest", "Ports": []string{}, "Volumes": []string{}, "Mounts": []string{}, "env": []string{}})
	// 所有机器上都要安装
	//AppGroup.Store("docker-agent", []string{"cicd.super.com", "kdax-robot-manager", "kdax-trade-4", "kdax-robot-2", "kdax-nginx-hk-1", "kcx-openresty-01", "kdax-market-2", "kdax-market-1", "kdax-trade-1", "kdax-nginx-hk-2", "nmax-nacos-03", "nmax-future-03", "nmax-future-01", "kcx-middle-01", "kdax-robot-1", "kdax-gateway-02", "kdax-agent-1", "kdax-ansible-jms", "nmax-future-02", "kdax-nginx-jp-1", "nmax-nacos-01", "kdax-trade-2", "nmax-es-01", "kdax-market-5", "docker-desktop", "kdax-nginx-jp-2", "kdax-gateway-01", "kdax-market-4", "kdax-trade-3", "nmax-nacos-02", "kdax-api-wallet-1"})
}
