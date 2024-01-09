package main

import (
	"goblockchain/wallet"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/rs/cors"
)

const templateDir = "templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (walletServer *WalletServer) Port() uint16 {
	return walletServer.port
}

func (walletServer *WalletServer) Gateway() string {
	return walletServer.gateway
}

func (walletServer *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(templateDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Println("invalid method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid http method")
	}
}

func (walletServer *WalletServer) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", walletServer.Index)
	mux.HandleFunc("/wallet", walletServer.Wallet)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(walletServer.Port())), handler))
}
