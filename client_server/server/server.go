package main

import (
	"bufio"
	"log"
	"net"
)

const (
	port string = "8080"
)

func initialize_server() {
	log.Println("Starting server...")

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s", port)
		return
	}
	log.Printf("Listening on port %s\n", port)

	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("Failed to accept port %s", port)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln("Client terminated")
		}
		log.Print("Message Received:", string(message))
	}

}

func main() {
	initialize_server()
}
