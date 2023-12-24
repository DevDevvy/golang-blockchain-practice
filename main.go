package main

import (
	"fmt"
	"goblockchain/blockchain"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}
func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	transaction := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	blockchain := blockchain.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(
		walletA.BlockchainAddress(), walletB.BlockchainAddress(),
		1.0, walletA.PublicKey(), transaction.GenerateSignature(),
	)
	fmt.Println("added?", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("walletA balance: %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("walletB balance: %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("walletM balance: %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
