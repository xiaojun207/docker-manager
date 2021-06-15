package ws

import (
	"docker-manager/model"
	"encoding/json"
	"errors"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
	"sync"
	"time"
)

type WsConnectionManager struct {
	connectionMap model.SyncMap
	lock          sync.Mutex
	LastData      model.SyncMap
}

var wsConnectManager = WsConnectionManager{}

func init() {
	wsConnectManager.init()
}

func (e *WsConnectionManager) init() {
	// 定时清理掉线的连接
	utils.NewFixedDelayJob(5*time.Second, e.Clean)
	// 启动心跳线程
	utils.NewFixedDelayJob(1*time.Minute, e.pingAll)
}

func (e *WsConnectionManager) Clean() {
	e.connectionMap.Range(func(key, value interface{}) bool {
		conn := value.(*Connection)
		if conn != nil && conn.isClosed {
			e.removeConnection(key)
		}
		return true
	})
}

func (e *WsConnectionManager) AddConn(id string, conn *Connection) {
	log.Println("AddConn, id:", id)
	e.connectionMap.Store(id, conn)
}

func (e *WsConnectionManager) GetList() model.SyncMap {
	return e.connectionMap
}

func (e *WsConnectionManager) GetCount() int {
	return e.connectionMap.Size()
}

func (e *WsConnectionManager) removeConnection(id interface{}) {
	e.connectionMap.Delete(id)
}

func (e *WsConnectionManager) Load(id string) (error, *Connection) {
	conn, ok := e.connectionMap.Load(id)
	if !ok {
		return errors.New("Conn is not exists , id:" + id), nil
	}
	return nil, conn.(*Connection)
}

func (e *WsConnectionManager) IsConnected(id string) bool {
	_, res := e.connectionMap.Load(id)
	return res
}

func (e *WsConnectionManager) pingAll() {
	e.connectionMap.Range(func(key, value interface{}) bool {
		conn := value.(*Connection)
		if conn == nil || conn.isClosed {
			e.removeConnection(key)
		}
		err := conn.Ping()
		if err != nil && !conn.isClosed {
			conn.Close()
			e.removeConnection(key)
		}
		return true
	})
}

func (e *WsConnectionManager) PushResp(id string, msg WsMsg) error {
	err, conn := e.Load(id)
	if err != nil {
		log.Println("PushResp, id:", id, ",err:", err)
		return errors.New("push resp conn is not exists , id:" + id)
	}
	//log.Println("PushResp, id:", id, ",msg:", msg)

	if conn == nil || conn.isClosed {
		e.removeConnection(id)
		return errors.New("push resp conn is isClosed , id:" + id)
		//} else if _, ok := conn.channels[resp.Channel]; ok {
		//	data, _ := json.Marshal(resp)
		//	return conn.WriteMessage(data)
	} else {
		data, _ := json.Marshal(msg)
		return conn.WriteMessage(data)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////

func addConn(conn *Connection) {
	wsConnectManager.AddConn(conn.id, conn)
}

func Push(id, chancel string, data interface{}) error {
	resp := NewWsMsg(chancel)
	resp.Data = data
	return wsConnectManager.PushResp(id, resp)
}

func LastData(channel string) (interface{}, bool) {
	return wsConnectManager.LastData.Load(channel)
}

func IsConnected(id string) bool {
	return wsConnectManager.IsConnected(id)
}
