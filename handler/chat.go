package handler

import "notification-center/center"

func HandleJoinRoom(cent center.Center, user *center.User, data interface{}) {
	cent.JoinHub("Chat", "hub_id", user)
}
