package center

import "github.com/gorilla/websocket"

type User struct {
	Conn *websocket.Conn
}
