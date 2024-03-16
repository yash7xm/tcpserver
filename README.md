# Multi-threaded TCP Server

## How to Run

```
$ go run main.go
```

Fire the following commands on another terminal to simulate
multiple concurrent requests.

```
$ curl http://localhost:2323 &
$ curl http://localhost:2323 &
$ curl http://localhost:2323 &
```
