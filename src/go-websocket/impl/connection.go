package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	WsConn    *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	CloseChan chan byte

	mutex    sync.Mutex
	IsClosed bool
}

func InitConn(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		WsConn:    wsConn,
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}
	//读协程
	go conn.readLoop()
	//写协程
	go conn.writeLoop()
	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.InChan:
	case <-conn.CloseChan:
		err = errors.New("connection is closed!")
	}

	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.OutChan <- data:
	case <-conn.CloseChan:
		err = errors.New("connection is closed!")
	}
	return
}
func (conn *Connection) Close() {
	//线程安全可重入的
	conn.WsConn.Close()
	//保证只执行一次
	conn.mutex.Lock()
	if !conn.IsClosed {
		close(conn.CloseChan)
		conn.IsClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.WsConn.ReadMessage(); err != nil {
			goto ERR
		}
		//inChan满了就会阻塞在这
		select {
		case conn.InChan <- data:
		case <-conn.CloseChan:
			//closeChan被关闭的时候
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
		case data = <-conn.OutChan:
		case <-conn.CloseChan:
			goto ERR
		}
		if err = conn.WsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
