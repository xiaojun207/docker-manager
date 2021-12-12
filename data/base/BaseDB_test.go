package base

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {

	path := "/tmp/docker-manager/database.db"

	log.Println(filepath.Base(path))
	log.Println(filepath.Dir(path))
	log.Println(filepath.Base(path))

	os.MkdirAll(filepath.Dir(path), os.ModePerm)
}
