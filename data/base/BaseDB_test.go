package base

import (
	"fmt"
	"github.com/go-basic/uuid"
	utils2 "github.com/xiaojun207/go-base-utils/utils"
	"strings"
	"testing"
)

func TestInitDB(t *testing.T) {
	key := utils2.Md5(uuid.New())
	key = strings.ReplaceAll(key, "-", "")
	key = utils2.Substring(key, 0, 8)
	fmt.Println(key)
	ds := utils2.DESEncrypt("value", key)
	fmt.Println(ds)

	es := utils2.DESDecrypt(ds, key)
	fmt.Println(es)

}
