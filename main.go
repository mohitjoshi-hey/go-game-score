package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return 123
}

func main() {

    server := &PlayerServer{&InMemoryPlayerStore{}}

	// prepares your PlayerServer function to be used as an HTTP handler, and: starts a server on port that sends incoming requests to that handler.
    log.Fatal(http.ListenAndServe(":5000", server))
}