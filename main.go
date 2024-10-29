package main

import (
	"fmt"
	"log"
)

func main() {
	server := NewServer(":8080")
	fmt.Println("Server running on port", server.listenAddr)
	if err := server.Start(); err!= nil {
		log.Fatal("Failed to start server", err)
	}

}

