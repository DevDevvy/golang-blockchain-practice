package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	block := &Block{}
	blockchain := new(Blockchain)
	blockchain.CreateBlock(0, block.Hash())
	return blockchain
}

// Print prints the blockchain
func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (transaction *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("_", 30))
	fmt.Printf("senderBlockchainAddress: %s\n", transaction.senderBlockchainAddress)
	fmt.Printf("recipientBlockchainAddress: %s\n", transaction.recipientBlockchainAddress)
	fmt.Printf("value: %.1f\n", transaction.value)
}

func (blockchain *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	transaction := NewTransaction(sender, recipient, value)
	blockchain.transactionPool = append(blockchain.transactionPool, transaction)
}

func (transaction *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    transaction.senderBlockchainAddress,
		RecipientBlockchainAddress: transaction.recipientBlockchainAddress,
		Value:                      transaction.value,
	})
}
