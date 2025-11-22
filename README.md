# Demystifying Go’s Concurrency and Performance (GDG Golang, Nov, 2025)

## Max CPUs
### Mutexes Demo
```bash
go run ./cmd/demo1/gomaxprocs.go -mode=mutex #implicitly GOMAXPROCS=0
```
vs.
```bash
GOMAXPROCS=1 go run ./cmd/demo1/gomaxprocs.go -mode=mutex
```

###  Atomics Demo

```bash
go run ./cmd/demo1/gomaxprocs.go -mode=atomic #implicitly GOMAXPROCS=0
```
vs.
```bash
GOMAXPROCS=1 go run ./cmd/demo1/gomaxprocs.go -mode=atomic
```

## benchstat

```bash
go install golang.org/x/perf/cmd/benchstat@latest
go test -run='^$' -bench=BenchmarkStartService -benchtime=10000x  -count=10 > current.txt
GOEXPERIMENT=greenteagc go test -run='^$' -bench=BenchmarkStartService -benchtime=10000x  -count=10 > greentea.txt
benchstat current.txt greentea.txt 
```

