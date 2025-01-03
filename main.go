package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qloog/geth-example/account"
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

	// account
	// account.GetAccountBalance()
	// account.GenKeystore()
	// account.ImportKeystore()
	account.CheckAccount()

	// Transaction()

	// TransferToken()

	// SubscribeBlock()

	// smart contract
	// contract.LoadContract()
	// contract.ReadContract()
	// contract.WriteContract()
	// contract.DeployContract()

	// erc20 token
	// contract.ReadContactERC20()

	// event
	// event.SubscribeEvent()
	// event.ReadEvent()
	// event.ReadERC20Event()
	// event.Read0xprotocolEvent()

	// singature
	// signature.GenerateSignature()

	// test
	// test.SimulateClient()
}
