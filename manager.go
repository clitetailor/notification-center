package main

import (
	"encoding/json"
	"notification-center/center"

	"github.com/gorilla/websocket"
)

type HandlerManager struct {
	handlers map[string]func(cent center.Center, u *center.User, data interface{})

	User   *center.User
	Center center.Center
}

func (manager *HandlerManager) AddHandler(action string, hdl func(cent center.Center, u *center.User, data interface{})) {
	manager.handlers[action] = hdl
}

func (manager *HandlerManager) Start() {
	for {
		messageType, message, err := manager.User.Conn.ReadMessage()
		if err != nil {
			return
		}

		if messageType == websocket.TextMessage {
			var v = map[string]interface{}{}

			err = json.Unmarshal(message, v)
			if err != nil {
				continue
			}

			_, ok := v["action"]
			if !ok {
				continue
			}

			action, ok := v["action"].(string)
			if !ok {
				continue
			}

			data, ok := v["data"]
			if !ok {
				continue
			}

			hdl, ok := manager.handlers[action]
			if !ok {
				continue
			}

			hdl(manager.Center, manager.User, data)
		}
	}
}
