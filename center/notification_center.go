package center

type NotificationCenter struct {
	hubs map[string]map[string]Hub

	hubFactories map[string]func() Hub
}

func (center *NotificationCenter) RegisterHubType(hubType string, factory func() Hub) {
	center.hubFactories[hubType] = factory
}

func (center *NotificationCenter) JoinHub(hubType string, hubID string, user *User) {
	center.hubs[hubType][hubID].AddUser(user)
}

func (center *NotificationCenter) LeaveHub(hubType string, hubID string, user *User) {
	center.hubs[hubType][hubID].RemoveUser(user)
}

func (center *NotificationCenter) SendHubData(hubType string, hubID string, data interface{}) {
	// TODO: Send JSON message.
}
