# HTTP Response Time Application

This application runs and hits the https://gitlab.com endpoint for 5 minutes
meanwhile printing the response times. After 5 minutes, the application 
reports back the number of requests made within that 5 minutes.

## Dependencies

[Go 1.9.2](https://golang.org/doc/install) |
[Dep](https://github.com/golang/dep#dep) |
[Docker](https://docs.docker.com/engine/installation/)

*Note: All instructions are written for OSX*

## Install Dev Environment
```
$ go get github.com/berryhill/http-response-time
$ cd $GOPATH/src/github/berryhill/http-response-time
$ dep ensure
```

## Up Dev Environment

#### Simple local build
```
$ cd $GOPATH/src/github/berryhill/http-response-time
$ go build
$ ./http-response-time
```

#### Containerized build via docker-compose
```
$ cd $GOPATH/src/github/berryhill/http-response-time
$ docker-compose up --build
```

## Testing

### Run Unit Tests Locally
```
$ go test ./...
```

### Benchmark with Memory Profiling

Create memprofile from Benchmark
``` 
$ go test -run none -bench . -benchtime 10s -benchmem -memprofile mem.out
```

View Benchmark memory profile with pprof
``` 
$ go tool pprof -alloc_space ./http-response-time.test mem.out
```


### Profiling 

Build and run program
```
$ go build
$ ./http-response-time
```

At beginning of program execution, create a snapshot of the heap
``` 
$ curl -s http://localhost:8765/debug/pprof/heap > base.heap
```

Later on during program execution, create another snapshot of the heap
```
$ curl -s http://localhost:8765/debug/pprof/heap > current.heap
```

Compare the two heaps with pprof
``` 
$ go tool pprof --base base.heap ./http-response-time current.heap
```

Also, while the program is executing, one can visit: [http://localhost:8765/debug/pprof/](http://localhost:8765/debug/pprof/)
