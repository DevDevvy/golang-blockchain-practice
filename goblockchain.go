package main

import (
	"goblockchain/blockchain"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}
func main() {
	block := blockchain.NewBlock(0, "init hash")
	block.Print()
}
