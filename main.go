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
	log.Println(listener)
}
