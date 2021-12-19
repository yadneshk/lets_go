package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	server string = "127.0.0.1"
	port   string = "8080"
)

func connect_server() {
	conn, err := net.Dial("tcp", server+":"+port)
	if err != nil {
		log.Fatalf("Failed to connect to server %s:%s\n", server, port)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message from server: " + message)
	}
}

func main() {
	connect_server()
}
