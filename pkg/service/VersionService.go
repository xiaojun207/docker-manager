package service

import (
	"encoding/json"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
	"strings"
	"time"
)

// export VERSION=$(git describe --tags `git rev-list --tags --max-count=1`)
// export DATE=`date -u '+%Y-%m-%d_%I:%M:%S%p-GMT'`
// export COMMIT_HASH=`git rev-parse HEAD`

// go build -a -ldflags "-X docker-manager/service.Version=$VERSION" -o docker-manager
var (
	Version    = "1.5.0"
	Date       = "20220815"
	CommitHash = ""
)

type HubTag struct {
	Count    int            `json:"count"`
	Next     interface{}    `json:"next"`
	Previous interface{}    `json:"previous"`
	Results  []HubTagResult `json:"results"`
}

type HubTagResult struct {
	Creator             int           `json:"creator"`
	Id                  int           `json:"id"`
	Images              []HubTagImage `json:"images"`
	LastUpdated         time.Time     `json:"last_updated"`
	LastUpdater         int           `json:"last_updater"`
	LastUpdaterUsername string        `json:"last_updater_username"`
	Name                string        `json:"name"`
	Repository          int           `json:"repository"`
	FullSize            int           `json:"full_size"`
	V2                  bool          `json:"v2"`
	TagStatus           string        `json:"tag_status"`
	TagLastPulled       time.Time     `json:"tag_last_pulled"`
	TagLastPushed       time.Time     `json:"tag_last_pushed"`
	MediaType           string        `json:"media_type"`
}

type HubTagImage struct {
	Architecture string      `json:"architecture"`
	Features     string      `json:"features"`
	Variant      interface{} `json:"variant"`
	Digest       string      `json:"digest"`
	Os           string      `json:"os"`
	OsFeatures   string      `json:"os_features"`
	OsVersion    interface{} `json:"os_version"`
	Size         int         `json:"size"`
	Status       string      `json:"status"`
	LastPulled   time.Time   `json:"last_pulled"`
	LastPushed   time.Time   `json:"last_pushed"`
}

func init() {
	log.Println("BuildInfo, Version:", Version, ", Date:", Date, ", CommitHash:", CommitHash)
}

func GetLatestTag() string {
	image := "xiaojun207/docker-manager"
	// https://hub.docker.com/v2/repositories/xiaojun207/docker-manager/tags/?page_size=5&page=1&ordering=last_updated
	url := "https://hub.docker.com/v2/repositories/" + image + "/tags/?page_size=5&page=1&ordering=last_updated"
	resp := utils.Get(url)
	var d map[string]interface{}
	err := json.Unmarshal([]byte(resp), &d)
	if err != nil {
		log.Println("JsonToMap err: ", err, url)
	}
	lastVersion := "0.0.0"
	results := (d["results"]).([]interface{})
	for _, m := range results {
		v := m.(map[string]interface{})["name"].(string)
		if v == "latest" {
			continue
		}
		if VersionCompare(lastVersion, v) < 0 {
			lastVersion = v
		}
	}
	return lastVersion
}

// eg.: 1.0.1
// return： -1 0 1
func VersionCompare(ver1, ver2 string) int {
	arrVer1 := strings.Split(ver1, ".")
	arrVer2 := strings.Split(ver2, ".")
	len1, len2 := len(arrVer1), len(arrVer2)
	minLen := utils.If(len1 < len2, len1, len2).(int)
	for i := 0; i < minLen; i++ {
		v1 := utils.StrToInt(arrVer1[i])
		v2 := utils.StrToInt(arrVer2[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	if len1 > len2 {
		return 1
	} else if len1 < len2 {
		return -1
	} else {
		return 0
	}
}
