package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"goblockchain/utils"
	"log"
	"strings"
)

// consts for mining
const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

// NewBlockchain creates a new blockchain
func NewBlockchain(blockchainAddress string, port uint16) *Blockchain {
	block := &Block{}
	blockchain := new(Blockchain)
	blockchain.blockchainAddress = blockchainAddress
	blockchain.port = port
	blockchain.CreateBlock(0, block.Hash())
	return blockchain
}

func (blockchain *Blockchain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chains"`
	}{
		Blocks: blockchain.chain,
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

// NewTransaction creates a new transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Print prints the transaction in json
func (transaction *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("_", 30))
	fmt.Printf("senderBlockchainAddress: %s\n", transaction.senderBlockchainAddress)
	fmt.Printf("recipientBlockchainAddress: %s\n", transaction.recipientBlockchainAddress)
	fmt.Printf("value: %.1f\n", transaction.value)
}

// AddTransaction adds a transaction to the transaction pool
func (blockchain *Blockchain) AddTransaction(
	sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, signature *utils.Signature) bool {
	transaction := NewTransaction(sender, recipient, value)
	if sender == MINING_SENDER {
		blockchain.transactionPool = append(blockchain.transactionPool, transaction)
		return true
	}
	if blockchain.VerifyTransactionSignature(senderPublicKey, signature, transaction) {
		//TODO uncomment to go live after testing
		// if blockchain.CalculateTotalAmount(sender) < value {
		// 	log.Println("ERROR: Not enough balance")
		// 	return false
		// }
		blockchain.transactionPool = append(blockchain.transactionPool, transaction)
		return true
	} else {
		log.Println("ERROR: Invalid transaction signature")
	}
	return false
}

// VerifyTransactionSignature verifies the signature of the transaction
func (blockchain *Blockchain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey, signature *utils.Signature, transaction *Transaction) bool {
	m, _ := json.Marshal(transaction)
	hash := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, hash[:], signature.R, signature.S)
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

// CalculateTotalAmount calculates the total by running through the blockchain
func (blockchain *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, block := range blockchain.chain {
		for _, transaction := range block.transactions {
			if transaction.recipientBlockchainAddress == blockchainAddress {
				totalAmount += transaction.value
			}
			if transaction.senderBlockchainAddress == blockchainAddress {
				totalAmount -= transaction.value
			}
		}
	}
	return totalAmount
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
