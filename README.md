# go-game-score

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26.1-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go version"/>
  <img src="https://img.shields.io/badge/JSON-Persistence-000000?style=for-the-badge&logo=json&logoColor=white" alt="JSON persistence"/>
  <img src="https://img.shields.io/badge/Testing-TDD-25A162?style=for-the-badge&logo=testinglibrary&logoColor=white" alt="Test-driven development"/>
  <img src="https://img.shields.io/badge/HTTP-Server-00599C?style=for-the-badge&logo=fastapi&logoColor=white" alt="HTTP server"/>
  <img src="https://img.shields.io/badge/CLI-Application-4EAA25?style=for-the-badge&logo=gnometerminal&logoColor=white" alt="Command-line interface"/>
</p>

A Go application built iteratively with TDD, exploring HTTP servers, interfaces, dependency injection, persistence, command-line applications, package structure, scheduling, WebSockets, and more.

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

### Command-Line Application

- [x] CLI for recording player wins
- [x] Read user input using `bufio.Scanner`
- [x] Use `io.Reader` for testable input
- [x] Reuse the `PlayerStore` interface
- [x] Persist player scores between CLI executions
- [x] Separate executable entry points for CLI and HTTP server
- [x] External-package tests using `package poker_test`
- [x] Shared testing utilities

## Project Structure

```text
go-game-score/
├── cmd/
│   ├── cli/
│   │   └── main.go
│   └── webserver/
│       └── main.go
├── CLI.go
├── CLI_test.go
├── file_system_store.go
├── file_system_store_test.go
├── in_memory_player_store.go
├── league.go
├── server.go
├── server_test.go
├── server_integration_test.go
├── tape.go
├── tape_test.go
├── testing.go
├── go.mod
└── README.md
```

## Architecture

### HTTP Server

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

### CLI Application

```text
User Input
    ↓
bufio.Scanner
    ↓
CLI
    ↓
PlayerStore (interface)
    ↓
FileSystemPlayerStore
    ↓
JSON Persistence
```

## Concepts Explored

- **Test-Driven Development (TDD):** Building functionality through tests.
- **Interfaces:** Using `PlayerStore` to abstract player storage.
- **Dependency Injection:** Passing dependencies into application components.
- **HTTP Handling:** Processing requests and returning player scores and league data.
- **JSON Encoding and Decoding:** Loading and saving player data.
- **File Persistence:** Keeping player scores between application runs.
- **Sorting:** Displaying players in descending order of wins.
- **Composition:** Wrapping file operations with `tape` to truncate the database before writing.
- **Command-Line Applications:** Building a CLI that reads user input and records player wins.
- **Package Structure:** Separating reusable application logic from executable entry points.
- **Exported and Unexported Identifiers:** Controlling access to package functionality.
- **External Package Testing:** Testing public APIs through `package poker_test`.
- **Test Helpers:** Sharing testing utilities across test files.

## Database

Player data is stored in JSON files as arrays of player records:

```json
[
  {"Name": "MJ", "Wins": 14},
  {"Name": "SN", "Wins": 5}
]
```

The application loads existing player data and writes updated scores back to the database.

The HTTP server and CLI currently use separate database filenames, so their data is persisted independently.

## Running the Applications

**HTTP Server**

```bash
go run ./cmd/webserver
```

**CLI**

```bash
go run ./cmd/cli
```

## Running Tests

```bash
go test ./...
```
