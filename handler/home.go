package handler

import (
	"fmt"
	"net/http"
)

//Home is the Handler for "/"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
