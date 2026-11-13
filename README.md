# Go Practice Project — No Database

A hands-on Go project designed to refresh fundamentals and practice real code.
Work through the packages in order. Each exercise has TODOs, expected behavior,
and tests where applicable.

## How to use

```bash
go mod init github.com/you/golang-practice
go test ./...
go run ./cmd/todo
```

## Learning path

1. `internal/basics` — variables, types, constants, strings, slices, maps
2. `internal/control` — if/switch, loops, range, early returns
3. `internal/functions` — functions, variadic, closures, recursion
4. `internal/structs` — structs, methods, embedding, pointers
5. `internal/interfaces` — interfaces, polymorphism, type switches
6. `internal/errors` — errors, wrapping, sentinel errors, custom errors
7. `internal/concurrency` — goroutines, channels, select, sync, context
8. `internal/generics` — type parameters and constraints
9. `internal/testing` — table-driven tests, benchmarks
10. `internal/projects` — larger practice components
11. `cmd/todo` — in-memory CLI todo app

## Rules

- Do not use a database.
- Implement each TODO yourself before looking at hints.
- Run `go vet ./...` and `go test ./...` often.
- Prefer small functions, clear names, and explicit error handling.
