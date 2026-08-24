package main

import (
	"log"
	"net/http"
)

func main() {

    server := &PlayerServer{NewInMemoryPlayerStore()}

	// prepares your PlayerServer function to be used as an HTTP handler, and: starts a server on port that sends incoming requests to that handler.
    log.Fatal(http.ListenAndServe(":5000", server))
}