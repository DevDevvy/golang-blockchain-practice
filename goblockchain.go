package main

import (
	"goblockchain/blockchain"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}
func main() {
	blockchain := blockchain.NewBlockchain()
	blockchain.Print()

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()

	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()

}
