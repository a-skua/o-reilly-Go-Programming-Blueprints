package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている一人のユーザーを表します。
type client struct {
	// socket はこのクライアントのためのWebSocketです
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです。
	send chan []byte
	// room はこのクライアントが参加しているチャットルームです。
	room *room
}
