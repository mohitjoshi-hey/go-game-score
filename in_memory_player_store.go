package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) { // Increase the player's score when u get it in the POST request
	i.store[name]++;
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int { // Show the player's score when u want to see it's score through the GET request
	return i.store[name]
}