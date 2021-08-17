package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/boydmeyer/bond-go/router"
)

func main() {
	r := router.Init()
	fmt.Printf("Listening on port 8080\t->\thttp://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}
