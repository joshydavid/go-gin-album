
# Go (Golang) Gin Album API

A small example REST API written in Go using the Gin web framework. This project demonstrates basic CRUD-style endpoints for managing a simple in-memory collection of music albums.

## Features

- Lightweight HTTP server using Go (Golang) Gin
- In-memory album store (example data in `internal/service/album_service.go`)
- Endpoints for listing, retrieving, creating, and deleting albums

## Getting Started

1. Clone the repository

```sh
git clone https://github.com/joshydavid/go-gin-album.git
cd go-gin-album
```

2. Run the server

```sh
go run ./cmd
```

By default Gin's `router.Run()` will start the server on port 8080 at `http://localhost:8080`.

## Build

Create a binary with:

```sh
go build -o bin/go-gin-album ./cmd
```

Then run it:

```sh
./bin/go-gin-album
```

## API Endpoints

The routes are registered in `api/routes.go`. Available endpoints:

- GET `/healthcheck` - simple health check
- GET `/albums` - return all albums
- GET `/albums/:id` - return a single album by ID
- POST `/albums` - add a new album (JSON body)
- DELETE `/albums/:id` - delete an album by ID

## Example Album JSON

```json
{
 "id": "4",
 "title": "Kind of Blue",
 "artist": "Miles Davis",
 "price": 29.99
}
```

## cURL examples

List all albums

```sh
curl -s http://localhost:8080/albums | jq
```

Get album by id

```sh
curl -s http://localhost:8080/albums/1 | jq
```

Add an album

```sh
curl -X POST http://localhost:8080/albums \
 -H "Content-Type: application/json" \
 -d '{"id":"4","title":"Kind of Blue","artist":"Miles Davis","price":29.99}'
```

Delete an album

```sh
curl -X DELETE http://localhost:8080/albums/4
```

## Project Structure

Key files and folders:

- `cmd/` — application entrypoint (`cmd/main.go`)
- `api/routes.go` — route registration
- `internal/handler/` — Gin handlers for each endpoint
- `internal/service/` — simple in-memory service layer and sample data
- `internal/model/` — domain models (e.g., `Album`)
- `internal/constant/` — route constants

## Data Model

The album model is defined in `internal/model/album.go`:

```go
type Album struct {
  ID     string  `json:"id"`
  Title  string  `json:"title"`
  Artist string  `json:"artist"`
  Price  float64 `json:"price"`
}
```

## Notes

- This project keeps albums in memory (slice in `internal/service/album_service.go`). For production use replace the in-memory store with a persistent database and add validation, logging, and tests.
- Add request validation and better error responses.
- Add unit/integration tests and a CI workflow.
