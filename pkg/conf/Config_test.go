package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	d, err := ioutil.ReadFile("../config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var c Config
	err = yaml.Unmarshal(d, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
}
