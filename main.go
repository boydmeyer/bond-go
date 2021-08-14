package main

import (
	"log"
	"net/http"

	"github.com/boydmeyer/bond/router"
	"github.com/gorilla/mux"
)

func main() {
	mr := mux.NewRouter()
	router.Setup(mr)
	log.Fatal(http.ListenAndServe(":8080", mr))
}
