# Concurrency Talk (GDG Golang, Nov, 2025)


## Max CPUs
### Mutexes Demo
```go
go run gomaxprocs/demo.go -mode=mutex #implicitly GOMAXPROCS=0
```
vs.
```go
GOMAXPROCS=1 go run gomaxprocs/demo.go -mode=mutex
```

###  Atomics Demo

```go
go run gomaxprocs/demo.go -mode=atomic #implicitly GOMAXPROCS=0
```
vs.
```go
GOMAXPROCS=1 go run gomaxprocs/demo.go -mode=atomic
```



