package baseWs

import (
	"github.com/go-basic/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsHandler(w http.ResponseWriter, r *http.Request, id string, group *WsConnectionGroup) {
	//log.Println("coming:", getCount())
	//	w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *Connection
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		log.Println("WsHandler.upgrade.err:", err)
		return
	}

	if id == "" {
		id = uuid.New()
	}

	if conn, err = NewConnection(id, wsConn); err != nil {
		log.Println("WsHandler.InitConnection.error", err)
		goto ERR
	}
	conn.Headers["AppId"] = r.Header.Get("AppId")
	group.AddConn(id, conn)

	for {
		if data, err = conn.ReadMessage(); err != nil {
			log.Println("WsHandler.ReadMessage.error", err)
			goto ERR
		}
		msg := ToWsMsg(data)

		if err = group.MsgHandler(msg, conn); err != nil {
			log.Println("WsHandler.WsMsgHandler.error", err)
			goto ERR
		}
	}
ERR:
	conn.Close()

}
