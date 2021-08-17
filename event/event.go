package event

import "encoding/json"

/*
EXAMPLE EVENT
{
	"event": "message",
	"data": "data"
}
*/

type EventHandler func(*Event)

type Event struct {
	Name string      `json:"event"`
	Data interface{} `json:"data"`
}

func New(rawData []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(rawData, &event)
	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
