package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
}

func main() {

	// Address()

	// Transaction()

	// TransferToken()

	SubscribeBlock()
}
