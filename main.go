package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/pdyraga/ethereum-playground/contracts"
)

const simpleGreeterAddress = "0xe5570894211b99563a6d6172c4fe33f52b41676e"
const echoGreeterAddress = "0xe2e90ca56834b64df24b22f6995b8e901362b3b4"

const ipcURL = "/Users/piotr/Library/Ethereum/rinkeby/geth.ipc"

func main() {
	conn, err := ethclient.Dial(ipcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	callSimpleGreeter(conn)
	callEchoGreeter(conn)
}

func callSimpleGreeter(conn *ethclient.Client) {
	simpleGreeter, err := contracts.NewSimpleGreeter(common.HexToAddress(simpleGreeterAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a SimpleGreeter contract: %v", err)
	}

	greeting, err := simpleGreeter.Greet(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve greeting: %v", err)
	}

	fmt.Println("Simple greeter greeting:", greeting)
}

func callEchoGreeter(conn *ethclient.Client) {
	echoGreeter, err := contracts.NewEchoGreeter(common.HexToAddress(echoGreeterAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a EchoGreeter contract: %v", err)
	}

	greeting, err := echoGreeter.Greet(nil, "yummy cookie")
	if err != nil {
		log.Fatalf("Failed to retrieve greeting:", err)
	}

	fmt.Println("Echo greeter greeting:", greeting)
}
