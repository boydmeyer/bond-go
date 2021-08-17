package router

import (
	"github.com/boydmeyer/bond-go/handler"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/ws", handler.Websocket)
}
