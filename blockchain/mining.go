package blockchain

import "log"

// Mining mines a block
func (blockchain *Blockchain) Mining() bool {
	//add the reward transaction
	blockchain.AddTransaction(MINING_SENDER, blockchain.blockchainAddress, MINING_REWARD, nil, nil)
	nonce := blockchain.ProofOfWork()
	//get the previous hash
	previousHash := blockchain.LastBlock().Hash()
	//create a new block
	blockchain.CreateBlock(nonce, previousHash)
	// return true to indicate that the mining was successful
	log.Println("action: mining, status: success")
	return true
}
