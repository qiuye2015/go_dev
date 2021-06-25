module github.com/qiuye2015/go_dev/go-websocket

go 1.14

require (
	github.com/gorilla/websocket v1.4.2
	github.com/qiuye2015/go_dev/go-websocket/impl v0.0.0
)

replace github.com/qiuye2015/go_dev/go-websocket/impl => ./impl
