package ws

import (
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

func WsHandler(w http.ResponseWriter, r *http.Request) {
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

	//AppId := r.Header.Get("AppId")
	//Token := r.Header.Get("authorization")
	//ServerId := r.Header.Get("ServerId")
	containerId := r.URL.Query().Get("containerId")
	serverName := r.Header.Get("ServerName")
	if serverName == "" {
		serverName = containerId
	}
	log.Println("WsHandler.coming", ",ServerName:", serverName)

	if conn, err = InitConnection(serverName, wsConn); err != nil {
		log.Println("WsHandler.InitConnection.error", err)
		goto ERR
	}

	for {
		if data, err = conn.ReadMessage(); err != nil {
			log.Println("WsHandler.ReadMessage.error", err)
			goto ERR
		}
		msg := ToWsMsg(data)

		if err = WsMsgHandler(msg, conn); err != nil {
			log.Println("WsHandler.WsMsgHandler.error", err)
			goto ERR
		}
	}
ERR:
	conn.Close()

}
