# go-game-score

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26.1-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go version"/>
  <img src="https://img.shields.io/badge/JSON-Persistence-000000?style=for-the-badge&logo=json&logoColor=white" alt="JSON persistence"/>
  <img src="https://img.shields.io/badge/Testing-TDD-25A162?style=for-the-badge&logo=testinglibrary&logoColor=white" alt="Test-driven development"/>
  <img src="https://img.shields.io/badge/HTTP-Server-00599C?style=for-the-badge" alt="HTTP server"/>
</p>

A Go application built iteratively with TDD, exploring HTTP servers, interfaces, dependency injection, persistence, CLI, scheduling, WebSockets, and more.

## 🚀 Current Progress

### Game Score Server

- [x] HTTP server using Go's standard library
- [x] `PlayerStore` interface
- [x] In-memory player store
- [x] File-system player store
- [x] JSON-based data persistence
- [x] GET player scores
- [x] POST player wins
- [x] League table sorted by wins
- [x] Empty database file handling
- [x] `tape` writer wrapper
- [x] Unit tests
- [x] Integration tests
- [x] Manual HTTP testing

## 🏗️ Architecture

```text
HTTP Request
     ↓
PlayerServer
     ↓
PlayerStore
   /       \
  /         \
InMemory    FileSystem
 Store         Store
                 ↓
            JSON Encoder
                 ↓
               tape
                 ↓
            db.file.json
```

## 🛠️ Tech Stack

<p>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"/>
  <img src="https://img.shields.io/badge/JSON-000000?style=flat-square&logo=json&logoColor=white" alt="JSON"/>
  <img src="https://img.shields.io/badge/HTTP-005C84?style=flat-square" alt="HTTP"/>
  <img src="https://img.shields.io/badge/Git-F05032?style=flat-square&logo=git&logoColor=white" alt="Git"/>
  <img src="https://img.shields.io/badge/GitHub-181717?style=flat-square&logo=github&logoColor=white" alt="GitHub"/>
</p>

## 🧩 Concepts Explored

- **Test-Driven Development (TDD):** Building functionality through tests.
- **Interfaces:** Abstracting storage with `PlayerStore`.
- **Dependency Injection:** Passing a store into `PlayerServer`.
- **HTTP Handling:** Processing player-score and win requests.
- **JSON Encoding and Decoding:** Loading and saving player data.
- **File Persistence:** Retaining scores between application runs.
- **Sorting:** Displaying players by wins in descending order.
- **Composition:** Using `tape` to truncate the database before writing.

## 💾 Database

Player data is stored in `db.file.json` as a JSON array:

```json
[
  {"Name": "MJ", "Wins": 14},
  {"Name": "SN", "Wins": 5}
]
```

The application loads the existing league on startup and saves updated scores when a player wins.
