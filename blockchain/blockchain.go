package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Block is a block in the blockchain
type Block struct {
	nonce        int
	previousHash [32]byte
	transactions []string
	timestamp    int64
}

// Blockchain is the blockchain
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlock Creates a new block
func NewBlock(nonce int, previousHash [32]byte) *Block {
	block := new(Block)
	block.timestamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash
	return block
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	block := &Block{}
	blockchain := new(Blockchain)
	blockchain.CreateBlock(0, block.Hash())
	return blockchain
}

// CreateBlock creates a new block on the chain
func (blockchain *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash)
	blockchain.chain = append(blockchain.chain, block)
	// block.transactions = blockchain.transactionPool
	// blockchain.transactionPool = []string{}
	return block
}

func (block *Block) Hash() [32]byte {
	m, _ := json.Marshal(block)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (block *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
		Timestamp    int64    `json:"timestamp"`
	}{
		Nonce:        block.nonce,
		PreviousHash: block.previousHash,
		Transactions: block.transactions,
		Timestamp:    block.timestamp,
	})
}

// Print prints the blockchain
func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (blockchain *Blockchain) LastBlock() *Block {
	return blockchain.chain[len(blockchain.chain)-1]
}

// Print prints the block"
func (block *Block) Print() {
	fmt.Printf("nonce: %d\n", block.nonce)
	fmt.Printf("previousHash: %x\n", block.previousHash)
	fmt.Printf("timestamp: %d\n", block.timestamp)
	fmt.Printf("transactions: %s\n", block.transactions)
}
