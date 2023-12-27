package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// NewBlock Creates a new block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	block := new(Block)
	block.timestamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash
	block.transactions = transactions
	return block
}

// CreateBlock creates a new block on the chain
func (blockchain *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash, blockchain.transactionPool)
	blockchain.chain = append(blockchain.chain, block)
	blockchain.transactionPool = []*Transaction{}
	return block
}

// Hash returns the hash of the block
func (block *Block) Hash() [32]byte {
	m, _ := json.Marshal(block)
	return sha256.Sum256([]byte(m))
}

// MarshalJSON converts the block to a readable JSON string
func (block *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PreviousHash string         `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
		Timestamp    int64          `json:"timestamp"`
	}{
		Nonce:        block.nonce,
		PreviousHash: fmt.Sprintf("%x", block.previousHash),
		Transactions: block.transactions,
		Timestamp:    block.timestamp,
	})
}

// LastBlock returns the last block in the chain
// in order to get the previous hash for each block
func (blockchain *Blockchain) LastBlock() *Block {
	return blockchain.chain[len(blockchain.chain)-1]
}

// Print prints the block"
func (block *Block) Print() {
	fmt.Printf("nonce: %d\n", block.nonce)
	fmt.Printf("previousHash: %x\n", block.previousHash)
	fmt.Printf("timestamp: %d\n", block.timestamp)
	for _, t := range block.transactions {
		t.Print()
	}
}
