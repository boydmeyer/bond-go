package handler

import (
	"log"
	"net/http"

	"github.com/boydmeyer/bond-go/event"
	"github.com/boydmeyer/bond-go/websocket"
)

func Websocket(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.New(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	/*
		Send other user a message
	*/
	ws.On("message", func(e *event.Event) {
		log.Printf("Message received: %s", e.Data.(string))
		ws.Out <- (&event.Event{
			Name: "response",
			Data: e.Data,
		}).Raw()
	})

	/*
		Send other user a heart
	*/
	ws.On("heart", func(e *event.Event) {
		log.Printf("heart: %s", e.Data.(string))
	})
}
