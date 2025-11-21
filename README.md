# Demystifying Go’s Concurrency and Performance (GDG Golang, Nov, 2025)

[Presentation for basic termonology](https://docs.google.com/presentation/d/1OE8rx3FRXD9M91Ss8segLIn8imwm4CIlrJ8zb2BhnsI/edit?slide=id.g3a0db67e201_2_201#slide=id.g3a0db67e201_2_201)

GC cost:
How often it runs
How long it runs

The memory hops to conduct sweep cost a long time.

What are the edge cases?

gRPC?
Why is called Green Tea?



## Container Awareness GOMAXPROCS

[Release notes](https://go.dev/doc/go1.25#container-aware-gomaxprocs)





[![Beer pouring video](https://img.youtube.com/vi/gqCGbCLxVKQ/0.jpg)](https://www.youtube.com/watch?v=gqCGbCLxVKQ)


## Max CPUs
### Mutexes Demo
```go
go run ./cmd/demo1/gomaxprocs.go -mode=mutex #implicitly GOMAXPROCS=0
```
vs.
```go
GOMAXPROCS=1 go run ./cmd/demo1/gomaxprocs.go -mode=mutex
```

###  Atomics Demo

```go
go run ./cmd/demo1/gomaxprocs.go -mode=atomic #implicitly GOMAXPROCS=0
```
vs.
```go
GOMAXPROCS=1 go run ./cmd/demo1/gomaxprocs.go -mode=atomic
```



