package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/pdyraga/ethereum-playground/contracts"
)

const contractAddress = "0xe5570894211b99563a6d6172c4fe33f52b41676e"
const ipcURL = "/Users/piotr/Library/Ethereum/rinkeby/geth.ipc"

func main() {
	conn, err := ethclient.Dial(ipcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	token, err := contracts.NewSimpleGreeter(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	name, err := token.Greet(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve greeting: %v", err)
	}

	fmt.Println("Token greeting:", name)
}
