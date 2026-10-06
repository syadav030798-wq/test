# Hotstar — interview Go scaffold

Minimal Go HTTP service structured with **hexagonal architecture** (ports and adapters). Use it as a starting point for AI-assisted coding exercises.

## Layout

```text
cmd/server/                 # composition root (wire dependencies, run HTTP server)
internal/
  domain/                   # entities and domain errors (no HTTP/DB imports)
  application/greeting/     # use cases + port interfaces
  adapter/
    http/                   # driving adapter (REST handlers, router)
    memory/                 # driven adapter (in-memory repository)
```

**Dependency rule:** adapters and `cmd` depend on application/domain; domain does not depend on adapters.

## Prerequisites

- Go 1.22+ ([install Go](https://go.dev/doc/install))

## Run locally

```bash
go run ./cmd/server
```

Optional port:

```bash
ADDR=:3000 go run ./cmd/server
```

## Try the APIs

Health:

```bash
curl -s http://localhost:8080/health
```

Hello (use case + in-memory repo):

```bash
curl -s "http://localhost:8080/api/v1/hello?name=Sonu"
```

Validation:

```bash
curl -s "http://localhost:8080/api/v1/hello?name="
```

## Tests

```bash
go test ./...
```

## Where to extend during an interview

| Task | Typical location |
|------|------------------|
| New business rule | `internal/domain/` |
| New use case | `internal/application/...` + new port interfaces |
| New REST route | `internal/adapter/http/` |
| Postgres / Redis | new package under `internal/adapter/` implementing a port |
| Second entrypoint (worker, CLI) | new `cmd/...` calling the same use cases |

## Build binary

```bash
go build -o bin/server ./cmd/server
./bin/server
```
