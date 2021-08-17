package main

import (
	"log"
	"net/http"

	"github.com/boydmeyer/bond-go/router"
)

func main() {
	r := router.Init()
	log.Fatal(http.ListenAndServe(":8080", r))
}
