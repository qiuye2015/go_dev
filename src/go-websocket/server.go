package main

import (
	"github.com/gorilla/websocket"
	"github.com/qiuye2015/go_dev/go-websocket/impl"
	"net/http"
	"time"
)

var (
	upprader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8888", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)
	wsConn, err = upprader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	if conn, err = impl.InitConn(wsConn); err != nil {
		goto ERR
	}
	go func() {
		for {
			if err := conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
