package utils

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	ls2 := []interface{}{
		map[string]interface{}{"Name": 3, "info": "333"},
		map[string]interface{}{"Name": 1, "info": "222"},
		map[string]interface{}{"Name": 2, "info": "111"},
	}
	f := func(i, j *interface{}) bool {
		a := (*i).(map[string]interface{})
		b := (*j).(map[string]interface{})
		return a["Name"].(int) > b["Name"].(int)
	}
	SortInterface(ls2, f)
	log.Println("ls2:", ls2)

	ls := []map[string]interface{}{
		{"Name": 3, "info": "333"},
		{"Name": 1, "info": "222"},
		{"Name": 2, "info": "111"},
	}
	SortArrMap(ls, func(a, b map[string]interface{}) bool {
		return a["Name"].(int) > b["Name"].(int)
	})
	log.Println("ls:", ls)
}
