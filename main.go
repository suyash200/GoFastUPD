package main

import (
	"fmt"
	"github.com/suyash200/GoFastUPD/client"
	"github.com/suyash200/GoFastUPD/server"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		return
	}

	switch os.Args[1] {
	case "server":
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		fmt.Println("Unknown command. Use 'server' or 'client'.")
	}
}
