package main

import (
	"log"
	"net/http"
)

func main(){
	handler := http.HandlerFunc(PlayerServer) // prepares your PlayerServer function to be used as an HTTP handler, and: starts a server on port 5000 that sends incoming requests to that handler.
	log.Fatal(http.ListenAndServe(":5000", handler))
}