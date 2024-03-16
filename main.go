package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Processing the request")
	time.Sleep(5 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":2323")

	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("waiting for client to connect")
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")
		go do(conn)
	}
}
