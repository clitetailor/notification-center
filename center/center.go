package center

type Center interface {
	RegisterHubType(hubType string, factory func() Hub)
	JoinHub(hubType string, hubID string, user *User)
	LeaveHub(hubType string, hubID string, user *User)
	SendHubData(hubType string, hubID string, data interface{})
}
