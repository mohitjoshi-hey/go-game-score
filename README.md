# go-game-score

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26.1-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go version"/>
  <img src="https://img.shields.io/badge/JSON-Persistence-000000?style=for-the-badge&logo=json&logoColor=white" alt="JSON persistence"/>
  <img src="https://img.shields.io/badge/Testing-TDD-25A162?style=for-the-badge&logo=testinglibrary&logoColor=white" alt="Test-driven development"/>
  <img src="https://img.shields.io/badge/HTTP-Server-00599C?style=for-the-badge&logo=fastapi&logoColor=white" alt="HTTP server"/>
</p>

A Go application built iteratively with TDD, exploring HTTP servers, interfaces, dependency injection, persistence, CLI, scheduling, WebSockets, and more.

## Current Progress

### Game Score Server

- [x] HTTP server using Go's standard library
- [x] `PlayerStore` interface
- [x] In-memory player store
- [x] File-system player store
- [x] JSON-based data persistence
- [x] GET player scores
- [x] POST player wins
- [x] League table sorted by wins
- [x] File handling for empty databases
- [x] `tape` wrapper to truncate and rewrite database files
- [x] Unit tests
- [x] Integration tests
- [x] Manual HTTP testing

## Architecture

```text
HTTP Request
     ↓
PlayerServer
     ↓
PlayerStore (interface)
     ↓
FileSystemPlayerStore
     ↓
JSON Encoder
     ↓
tape (file writer wrapper)
     ↓
db.file.json
```

## Concepts Explored

- **Test-Driven Development (TDD):** Building functionality through tests.
- **Interfaces:** Using `PlayerStore` to abstract player storage.
- **Dependency Injection:** Passing a store into `PlayerServer`.
- **HTTP Handling:** Processing requests and returning player scores and league data.
- **JSON Encoding and Decoding:** Loading and saving player data.
- **File Persistence:** Keeping player scores between application runs.
- **Sorting:** Displaying players in descending order of wins.
- **Composition:** Wrapping file operations with `tape` to truncate the database before writing.

## Database

Player data is stored in `db.file.json` as a JSON array:

```json
[
  {"Name": "MJ", "Wins": 14},
  {"Name": "SN", "Wins": 5}
]
```

The application loads the existing league when it starts and writes updated scores to the database file when a player wins.
