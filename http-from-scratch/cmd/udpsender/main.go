package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Resolve the UDP address (localhost:42069)
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}
	// Prepare the UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Failed to dial UDP: %v", err)
	}
	defer conn.Close()

	// Create a new reader stdin
	reader := bufio.NewReader(os.Stdin)

	// Start an infinite loop for user input
	for {
		fmt.Printf("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Erro reading input: %v", err)
			continue
		}

		// Write line to UDP connection
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Printf("Error writing to UDP: %v", err)
			continue
		}
	}
}
