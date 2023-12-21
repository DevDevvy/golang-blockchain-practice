package blockchain

func (blockchain *Blockchain) Mining() bool {
	//add the reward transaction
	blockchain.AddTransaction(MINING_SENDER, blockchain.blockchainAddress, MINING_REWARD)
	nonce := blockchain.ProofOfWork()
	//get the previous hash
	previousHash := blockchain.LastBlock().Hash()
	//create a new block
	blockchain.CreateBlock(nonce, previousHash)
	// return true to indicate that the mining was successful
	return true
}
