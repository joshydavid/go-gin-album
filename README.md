# 🐹 Go (Golang) Gin Album API

A small example REST API written in Go using the Gin web framework. This project provides API endpoints for managing a simple collection of music albums.

## 🚀 Features

- Lightweight HTTP server using Go (Golang) Gin
- Endpoints for listing, retrieving, creating, and deleting albums
- Redis for caching
- Rate Limiting for API endpoints to prevent abuse
- Data stored in PostgreSQL
- API documentation via Swagger

## ⚙️ Getting Started

Clone the repository

```sh
git clone https://github.com/joshydavid/go-gin-album.git
cd go-gin-album
```

Run the server

```sh
go run ./cmd
```

By default Gin's `router.Run()` will start the server on port 8080 at `http://localhost:8080`.

## 🚧 Build

Create a binary with:

```sh
go build -o bin/go-gin-album ./cmd
```

Then run it:

```sh
./bin/go-gin-album
```

## 📁 API documentation
`http://localhost:8080/docs/index.html`

## 💬 API Endpoints

The routes are registered in `api/routes.go`.

```http
GET /api/v1/healthcheck
GET /api/v1/albums
POST /api/v1/albums
GET /api/v1/albums/:id
DELETE /api/v1/albums/:id
```

## ✏️ Example Album JSON

```json
{
 "title": "Kind of Blue",
 "artist": "Miles Davis",
 "price": 29.99
}
```

## ✏️ cURL examples

List all albums

```sh
curl -s http://localhost:8080/api/v1/albums | jq
```

Get album by id

```sh
curl -s http://localhost:8080/api/v1/albums/1 | jq
```

Add an album

```sh
curl -X POST http://localhost:8080/api/v1/albums \
 -H "Content-Type: application/json" \
 -d '{"title":"Kind of Blue","artist":"Miles Davis","price":29.99}'
```

Delete an album

```sh
curl -X DELETE http://localhost:8080/api/v1/albums/4
```

## 📁 Project Structure

Key files and folders:

- `cmd/` - application entrypoint (`cmd/main.go`)
- `api/routes.go` - route registration
- `internal/config/` - database configurations
- `internal/db/` - database set up
- `internal/dto/` - map model to client response
- `internal/server/` - server set up
- `internal/repository/` - repository interfaces and concrete implementations
- `internal/handler/` - Gin handlers for each endpoint
- `internal/service/` - service layer where it contains business logic
- `internal/model/` - domain models (e.g., `Album`)
- `internal/constant/` - route constants

## 💽 Data Model

The album model is defined in `internal/model/album.go`:

```go
type Album struct {
  gorm.Model
  Title  string  `json:"title"`
  Artist string  `json:"artist"`
  Price  float64 `json:"price"`
}
```

## 🚀 Acknowledgement

<a href="https://www.linkedin.com/in/joshydavid/">
  <img src="https://github.com/user-attachments/assets/4dfe0c89-8ced-4e08-bcf3-6261bdbb956d" width="80">
</a> &nbsp;

Developed by [Joshua](https://joshydavid.com)
