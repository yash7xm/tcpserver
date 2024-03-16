package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":2323")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("waiting for client to connect")
	conn, err := listener.Accept()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(conn)
	log.Println("client connected")
}
