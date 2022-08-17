package model

import "testing"

func TestLevelDB(t *testing.T) {
	var test = LevelDB{Path: "./db/test"}
	test.Init()
	data := map[string]interface{}{}
	data["name"] = "test"
	test.Store("test", data)
}
