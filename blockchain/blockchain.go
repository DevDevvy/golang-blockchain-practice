package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

// NewBlockchain creates a new blockchain
func NewBlockchain(blockchainAddress string) *Blockchain {
	block := &Block{}
	blockchain := new(Blockchain)
	blockchain.blockchainAddress = blockchainAddress
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

// CopyTransactionPool copies the transaction pool before it's cleared
func (blockchain *Blockchain) CopyTransactionPool() []*Transaction {
	//creates a slice of pointers for the new transactions
	transactions := make([]*Transaction, 0)
	for _, transaction := range blockchain.transactionPool {
		transactions = append(transactions,
			NewTransaction(transaction.senderBlockchainAddress,
				transaction.recipientBlockchainAddress,
				transaction.value,
			),
		)
	}
	return transactions
}

// ValidProof calculates the difficulty
func (blockchain *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHash := guessBlock.Hash()
	guessString := fmt.Sprintf("%x", guessHash)
	fmt.Println(guessString)
	return guessString[:difficulty] == zeroes
}

// ProofOfWork cycles through solutions to find the nonce
func (blockchain *Blockchain) ProofOfWork() int {
	transactions := blockchain.CopyTransactionPool()
	previousHash := blockchain.LastBlock().Hash()
	nonce := 0
	for !blockchain.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}

// MarshalJSON makes readable output
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
