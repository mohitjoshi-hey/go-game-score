package main

import (
	poker "github.com/mohitjoshi-hey/texas-holdem-tournament-engine"
	"log"
	"net/http"
)

const dbFileName = "db.file.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	// prepares your PlayerServer function to be used as an HTTP handler, and: starts a server on port that sends incoming requests to that handler.
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
