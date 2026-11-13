# Go Cheatsheet

## Basics

```go
var name string = "Ada"
age := 30
const Pi = 3.14159
```

## Slices and maps

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)

users := map[string]int{"ada": 1}
value, ok := users["grace"]
```

## Structs and methods

```go
type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}
```

## Interfaces

```go
type Shape interface {
	Area() float64
}
```

## Errors

```go
if err != nil {
	return fmt.Errorf("loading user: %w", err)
}
```

## Goroutines and channels

```go
ch := make(chan int, 1)
go func() { ch <- 42 }()

value := <-ch
```

## Testing

```bash
go test ./...
go test -cover ./...
go test -bench=.
```

## Useful commands

```bash
go run ./cmd/todo
go build ./...
go vet ./...
go fmt ./...
```
