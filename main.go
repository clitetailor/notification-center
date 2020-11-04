package main

import (
	"log"
	"net/http"

	"notification-center/center"
	"notification-center/handler"

	"github.com/gorilla/websocket"
)

func main() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	nc := &center.NotificationCenter{}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		u := &center.User{
			Conn: conn,
		}

		man := &HandlerManager{
			User:   u,
			Center: nc,
		}

		man.AddHandler("join_room", handler.HandleJoinRoom)

		go man.Start()
	})
}
