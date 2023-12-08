package main

import (
	"fmt"
	"log"
	"net"
)

const port = 8080

func main() {
	server := newServer()
	go server.executeCommand()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
	defer listener.Close()
	log.Printf("server listening on port %v", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err)
			continue
		}
		go server.newClient(conn)
	}
}
