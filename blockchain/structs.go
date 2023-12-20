package blockchain

// Block is a block in the blockchain
type Block struct {
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
	timestamp    int64
}

// Blockchain is the blockchain
type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}
