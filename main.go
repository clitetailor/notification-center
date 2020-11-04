package main

import (
	"flag"
	"log"
	"net/http"

	"notification-center/center"
	"notification-center/handler"
	"notification-center/hub/chat"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	nc := &center.NotificationCenter{}

	nc.RegisterHubType("Chat", func() center.Hub {
		return &chat.ChatHub{}
	})

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

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
