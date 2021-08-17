package router

import (
	"github.com/boydmeyer/bond-go/handler"
	"github.com/gorilla/mux"
)

//Init initiates a mux router and implements endpoints.
func Init() *mux.Router {
	mr := mux.NewRouter()

	mr.HandleFunc("/", handler.Home)
	mr.HandleFunc("/ws", handler.Websocket)

	return mr
}
