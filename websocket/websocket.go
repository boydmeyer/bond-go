package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/boydmeyer/bond-go/event"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//WebSocket represents a WebSocket connection.
type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]event.Handler
}

//New creates a new Websocket struct
func New(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error while upgrading connection: %v", err)
		return nil, err
	}

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]event.Handler),
	}

	go ws.Reader()
	go ws.Writer()

	return ws, nil
}

//Reader reads messages sent over a websocket connection
func (ws *WebSocket) Reader() {
	defer func() {
		ws.Conn.Close()
	}()

	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WS Message Error: %v", err)
			}
		}

		event, err := event.New(message)
		if err != nil {
			log.Printf("Message parse error: %v", err)
		}

		log.Printf("Message: %v", event)

		if action, ok := ws.Events[event.Name]; ok {
			action(event)
		}
	}

}

//Writer writes messages over the websocket connection
func (ws *WebSocket) Writer() {
	for {
		select {
		case message, ok := <-ws.Out:
			if !ok {
				ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				return
			}
			w, err := ws.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			w.Close()
		}
	}
}

//On implements an Event on a websocket connection
func (ws *WebSocket) On(eventName string, action event.Handler) *WebSocket {
	ws.Events[eventName] = action
	return ws
}
