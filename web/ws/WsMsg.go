package ws

import (
	"encoding/json"
	"github.com/xiaojun207/go-base-utils/utils"
	"time"
)

const (
	CH_PING          = "base.ht.ping"
	CH_PONG          = "base.ht.pong"
	CH_USER_LOGIN    = "user.login"
	CH_SUB_PREFIX    = "base.sub"
	CH_SUB_OK        = "base.sub.ok"
	CH_CANCEL_PREFIX = "base.cancel"
	CH_CANCEL_OK     = "base.cancel.ok"

	CH_SUB_FAIL = "base.sub.fail"
)

type WsMsg struct {
	Channel   string      `json:"ch"` // "docker.container.run"
	Timestamp int64       `json:"ts"` // 时间戳
	Data      interface{} `json:"d"`  // 数据
}

func NewWsMsg(ch string) WsMsg {
	return WsMsg{
		Channel:   ch,
		Timestamp: time.Now().UnixNano() / 1e6,
		Data:      nil,
	}
}

func (e *WsMsg) ToString() string {
	return utils.StructToJson(e)
}

func (e *WsMsg) ToBytes() []byte {
	return []byte(e.ToString())
}

type WsReq struct {
	Channel   string `json:"ch"`
	Timestamp int64  `json:"ts"`
	Data      string `json:"d"`
}

func ToWsMsg(d []byte) (msg *WsMsg) {
	msg = &WsMsg{}
	json.Unmarshal(d, &msg)
	return
}
