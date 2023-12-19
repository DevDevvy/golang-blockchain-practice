package blockchain

import (
	"fmt"
	"time"
)

// Block is a block in the blockchain
type Block struct {
	nonce        int
	previousHash string
	transactions []string
	timestamp    int64
}

// NewBlock Creates a new block
func NewBlock(nonce int, previousHash string) *Block {
	block := new(Block)
	block.timestamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash
	return block
}

// Print prints the block"
func (block *Block) Print() {
	fmt.Printf("nonce: %d\n", block.nonce)
	fmt.Printf("previousHash: %s\n", block.previousHash)
	fmt.Printf("timestamp: %d\n", block.timestamp)
	fmt.Printf("transactions: %s\n", block.transactions)
}
