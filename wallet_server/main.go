package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Wallet Server:")
}

func main() {
	port := flag.Uint("port", 8080, "TCP port number for wallet server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain gateway")
	flag.Parse()
	fmt.Println("port:", *port)
	fmt.Println("gateway:", *gateway)
	app := NewWalletServer(uint16(*port), *gateway)
	app.Run()
}
