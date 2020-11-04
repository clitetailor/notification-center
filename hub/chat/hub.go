package chat

import "notification-center/center"

type ChatHub struct {
	users []*center.User
}

func (hub *ChatHub) AddUser(user *center.User) {
	// TODO: Add mutex.
	hub.users = append(hub.users, user)
}

func (hub *ChatHub) RemoveUser(user *center.User) {
	// TODO: Add mutex.
	hub.users = []*center.User{}

	for _, hubUser := range hub.users {
		if hubUser != user {
			hub.users = append(hub.users, hubUser)
		}
	}
}
