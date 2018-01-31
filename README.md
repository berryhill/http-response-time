# HTTP Response Time Application

## Dependencies

[Go 1.9.1](https://golang.org/doc/install) |
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

#### Containerized build via Docker

```
$ cd $GOPATH/src/github/berryhill/http-response-time
$ docker-compose up
```
