package center

type Hub interface {
	AddUser(user *User)
	RemoveUser(user *User)
}
