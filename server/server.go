package main

import (
	"net/http"

	"golang.org/x/net/context"
)

var mux = createGRPCHandler(context.Background())

func init() {
	http.Handle("/", mux)
}

func main() {
	http.ListenAndServe(":8080", nil)
}
