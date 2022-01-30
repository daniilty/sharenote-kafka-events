package events

const (
	EventTypeUserDelete = "user_delete"
)

type Event struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type UserDeleteEvent struct {
	ID string `json:"id"`
}
