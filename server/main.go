package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}

// to run the server run: "go run main.go blockchain_server.go"
func main() {
	port := flag.Uint("port", 5000, "TCP port number for blockchain server")
	flag.Parse()
	fmt.Println("port:", *port)
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
