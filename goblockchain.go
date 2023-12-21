package main

import (
	"goblockchain/blockchain"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}
func main() {
	bcAddress := "1234567890"
	blockchain := blockchain.NewBlockchain(bcAddress)
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	blockchain.Mining()
	blockchain.Print()

}
