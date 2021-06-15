package data

func GetServerNameByContainerId(containerId string) string {
	res, _ := ContainerServerMap.LoadStr(containerId)
	return res
}

func GetFollowSize() int {
	count := 0
	Stats.ForEachArrMap(func(ServerName string, arr []map[string]interface{}) {
		for _, v := range arr {
			container := v
			if container["Follow"].(bool) {
				count++
			}
		}
	})
	return count
}
