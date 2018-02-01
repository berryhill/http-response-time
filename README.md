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
$ go run main.go
```

#### Containerized build via docker-compose

```
$ cd $GOPATH/src/github/berryhill/http-response-time
$ docker-compose up --build
```
