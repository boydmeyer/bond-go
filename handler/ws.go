package handler

import (
	"log"
	"net/http"

	"github.com/boydmeyer/bond/event"
	"github.com/boydmeyer/bond/websocket"
)

func Websocket(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.New(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	ws.On("message", func(e *event.Event) {
		log.Printf("Message received: %s", e.Data.(string))
	})
}
