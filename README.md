# Queue

Queue data structure using slice (thread-safe)

```go
type queue struct {
	front, rear int
	items       []elementType
	rwLock      sync.RWMutex
}
```

Here `elementType` is `any` type in golang.
```go
type elementType any
```