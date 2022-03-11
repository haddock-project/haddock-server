package events

import "encoding/json"

type Event struct {
	Name string          `json:"name"`
	Args json.RawMessage `json:"args"`
}
