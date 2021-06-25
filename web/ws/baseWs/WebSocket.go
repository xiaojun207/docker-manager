package baseWs

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	id           string
	wsConnect    *websocket.Conn
	inChan       chan []byte
	outChan      chan []byte
	closeChan    chan byte
	channels     map[string]string
	mutex        sync.Mutex // 对closeChan关闭上锁
	Closed       bool       // 防止closeChan被关闭多次
	LastPongTime int64      // 毫秒级
}

func NewConnection(id string, wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
		channels:  map[string]string{},
		id:        id,
	}

	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {

	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("r connection is closed")
	}
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {

	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("w connection is closed")
	}
	return
}

func (conn *Connection) Push(resp WsMsg) (err error) {
	return conn.WriteMessage(resp.ToBytes())
}

func (conn *Connection) Ping() (err error) {
	resp := NewWsMsg(CH_PING)
	return conn.WriteMessage(resp.ToBytes())
}

func (conn *Connection) Pong() (err error) {
	resp := NewWsMsg(CH_PONG)
	return conn.WriteMessage(resp.ToBytes())
}

func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.Closed {
		close(conn.closeChan)
		conn.Closed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
