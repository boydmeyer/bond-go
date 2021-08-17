package event

import "encoding/json"

/*
EXAMPLE EVENT
{
	"event": "message",
	"data": "data"
}
*/

//Handler handles events on the Websocket.
type Handler func(*Event)

//Event is implemented with Name and Data. It holds information sent over websockets.
type Event struct {
	Name string      `json:"event"`
	Data interface{} `json:"data"`
}

//New creates a new Event struct from raw json data.
func New(rawData []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(rawData, &event)
	return event, err
}

//Raw creates raw json data from an array of bytes.
func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
