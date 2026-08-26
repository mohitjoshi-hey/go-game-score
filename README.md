# go-game-score

A Go application built iteratively with TDD, exploring HTTP servers, interfaces, dependency injection, persistence, CLI, scheduling, WebSockets, and more.

## Current Progress

### Game Score Server

- [x] HTTP server using Go's standard library
- [x] `PlayerStore` interface
- [x] In-memory player store
- [x] GET player scores
- [x] POST player wins
- [x] Unit tests
- [x] Integration tests
- [x] Manual HTTP testing

### Architecture

```text
HTTP Request
     ↓
PlayerServer
     ↓
PlayerStore
     ↓
InMemoryPlayerStore
     ↓
map[string]int
