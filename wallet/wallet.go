package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"goblockchain/utils"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

// Wallet is a wallet
type Wallet struct {
	//ecdsa is the elliptic curve digital signature algorithm
	//used to generate public and private keys
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

// Transaction is a transaction
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
}

// NewWallet creates a new wallet with a public and private key
func NewWallet() *Wallet {
	wallet := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	wallet.privateKey = privateKey
	wallet.publicKey = &wallet.privateKey.PublicKey
	h2 := sha256.New()
	//Perform a SHA-256 hash on the public key
	h2.Write([]byte(wallet.publicKey.X.Bytes()))
	h2.Write([]byte(wallet.publicKey.Y.Bytes()))
	digest2 := h2.Sum(nil)
	//Perform a SHA-3 hash on the result of SHA-256
	h3 := sha3.NewLegacyKeccak256()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	//Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	//Perform SHA-256 hash on the extended RIPEMD-160 result
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	//Perform SHA-256 hash on the result of the previous SHA-256 hash
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	//Take the first 4 bytes of the second SHA-256 hash. This is the address checksum
	chsum := digest6[:4]
	//Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4.
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	//Convert the result from a byte string into a base58 string using Base58Check encoding.
	address := base58.Encode(dc8)
	wallet.blockchainAddress = address

	return wallet
}

// BlockchainAddress returns the blockchain address
func (wallet *Wallet) BlockchainAddress() string {
	return wallet.blockchainAddress
}

// MarshalJSON makes readable output as the blockchain address
func (wallet *Wallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"private_key"`
		PublicKey         string `json:"public_key"`
		BlockchainAddress string `json:"blockchain_address"`
	}{
		PrivateKey:        wallet.PrivateKeyStr(),
		PublicKey:         wallet.PublicKeyStr(),
		BlockchainAddress: wallet.blockchainAddress,
	})
}

// PrivateKey returns the private key
func (wallet *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return wallet.privateKey
}

// PrivateKeyStr returns the private key as a readable string
func (wallet *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", wallet.privateKey.D.Bytes())
}

// PublicKey returns the public key
func (wallet *Wallet) PublicKey() *ecdsa.PublicKey {
	return wallet.publicKey
}

// PublicKeyStr returns the public key as a readable string
func (wallet *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", wallet.publicKey.X.Bytes(), wallet.publicKey.Y.Bytes())
}

// NewTransaction creates a new transaction
func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value, privateKey, publicKey}
}

func (transaction *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(transaction)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, transaction.senderPrivateKey, h[:])
	return &utils.Signature{R: r, S: s}
}

// MarshalJSON makes readable output as the signature
func (transaction *Transaction) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    transaction.senderBlockchainAddress,
		Recipient: transaction.recipientBlockchainAddress,
		Value:     transaction.value,
	})
}
