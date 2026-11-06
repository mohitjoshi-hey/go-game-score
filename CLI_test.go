package poker_test

import (
	poker "github.com/mohitjoshi-hey/texas-holdem-tournament-engine"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Chris"

		poker.AssertPlayerWin(t, playerStore, want)
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Cleo"

		poker.AssertPlayerWin(t, playerStore, want)
	})
}
