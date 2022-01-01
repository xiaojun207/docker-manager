package baseWs

import (
	"docker-manager/model"
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type WsConnectionGroup struct {
	connectionMap  model.SyncMap
	lock           sync.Mutex
	LastData       model.SyncMap
	OnConnected    func(id string, conn *Connection)
	OnDisconnected func(id string)
	DataParse      func(d []byte) (msg *WsMsg)
	MsgHandler     func(wsData []byte, conn *Connection) error
}

func NewWsConnectionGroup() WsConnectionGroup {
	group := WsConnectionGroup{}
	group.init()
	return group
}

func (e *WsConnectionGroup) init() {
	// 定时清理掉线的连接
	//utils.NewFixedDelayJob(5*time.Second, e.Clean)
	//启动心跳线程
	//utils.NewFixedDelayJob(1*time.Minute, e.pingAll)
}

func (e *WsConnectionGroup) Clean() {
	e.connectionMap.Range(func(key, value interface{}) bool {
		conn := value.(*Connection)
		if conn != nil && conn.Closed {
			e.removeConnection(key.(string))
		}
		return true
	})
}

func (e *WsConnectionGroup) AddConn(id string, conn *Connection) {
	log.Println("AddConn.id:", id)
	e.connectionMap.Store(id, conn)
	if e.OnConnected != nil {
		e.OnConnected(id, conn)
	}
}

func (e *WsConnectionGroup) GetList() model.SyncMap {
	return e.connectionMap
}

func (e *WsConnectionGroup) GetCount() int {
	return e.connectionMap.Size()
}

func (e *WsConnectionGroup) removeConnection(id string) {
	e.connectionMap.Delete(id)
	if e.OnDisconnected != nil {
		e.OnDisconnected(id)
	}
}

func (e *WsConnectionGroup) Load(id string) (error, *Connection) {
	conn, ok := e.connectionMap.Load(id)
	if !ok {
		return errors.New("Conn is not exists , id:" + id), nil
	}
	return nil, conn.(*Connection)
}

func (e *WsConnectionGroup) LastDataTime(id string) int64 {
	err, conn := e.Load(id)
	if err != nil {
		return 0
	}
	return conn.LastDataTime
}

func (e *WsConnectionGroup) IsConnected(id string) bool {
	_, res := e.connectionMap.Load(id)
	return res
}

func (e *WsConnectionGroup) pingAll() {
	e.connectionMap.Range(func(key, value interface{}) bool {
		conn := value.(*Connection)
		if conn == nil || conn.Closed {
			e.removeConnection(key.(string))
		}
		err := conn.Ping()
		if err != nil && !conn.Closed {
			conn.Close()
			e.removeConnection(key.(string))
		}
		return true
	})
}

func (e *WsConnectionGroup) PushResp(id string, msg WsMsg) error {
	data, _ := json.Marshal(msg)
	return e.PushData(id, data)
}

func (e *WsConnectionGroup) PushData(id string, data []byte) error {
	err, conn := e.Load(id)
	if err != nil {
		log.Println("PushResp, id:", id, ",err:", err)
		return errors.New("push resp conn is not exists , id:" + id)
	}
	//log.Println("PushResp, id:", id, ",msg:", msg)

	if conn == nil || conn.Closed {
		e.removeConnection(id)
		return errors.New("push resp conn is isClosed , id:" + id)
		//} else if _, ok := conn.channels[resp.Channel]; ok {
		//	data, _ := json.Marshal(resp)
		//	return conn.WriteMessage(data)
	} else {
		return conn.WriteMessage(data)
	}
}

func (e *WsConnectionGroup) Push(id, chancel string, data interface{}) error {
	resp := NewWsMsg(chancel)
	resp.Data = data
	return e.PushResp(id, resp)
}

///////////////////////////////////////////////////////////////////////////////////////////
//
//func LastData(channel string) (interface{}, bool) {
//	return wsConnectManager.LastData.Load(channel)
//}
