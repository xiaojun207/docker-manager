package utils

import (
	"log"
	"reflect"
	"strings"
)

func TrimContainerName(s interface{}) string {
	names := ""
	typeName := reflect.TypeOf(s).String()
	if typeName == "[]interface {}" {
		names = s.([]interface{})[0].(string)
	} else if typeName == "string" {
		names = s.(string)
	} else {
		log.Println("typeName:", typeName)
	}

	return strings.TrimLeft(names, "/")
}
