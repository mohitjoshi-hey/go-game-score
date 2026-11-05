package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/mohitjoshi-hey/go-game-score"
)

const dbFileName = "db.file.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
