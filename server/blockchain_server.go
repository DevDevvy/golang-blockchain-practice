package main

import (
	"goblockchain/blockchain"
	"goblockchain/wallet"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (server *BlockchainServer) Port() uint16 {
	return server.port
}

func (server *BlockchainServer) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minersWallet.BlockchainAddress(), server.Port())
		cache["blockchain"] = bc
		log.Println("blockchain created, private key:", minersWallet.PrivateKeyStr())
	}
	return bc
}

func (server *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		blockchain := server.GetBlockchain()
		m, _ := blockchain.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Println("invalid method")
	}
}

func (server *BlockchainServer) Run() {
	http.HandleFunc("/", server.GetChain)
	http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(server.Port())), nil)
}
