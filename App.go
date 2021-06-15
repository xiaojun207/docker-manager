package main

import (
	"docker-manager/data"
	"docker-manager/web"
)

func main() {

	web.Start("8068", "/dockerMgrApi/")

	defer data.Close()

	select {}
}
