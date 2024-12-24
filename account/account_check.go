package account

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckAccount() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// address check
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("address is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("address is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	// check if address is an Account or a Smart Contract
	// 0x Protocol Token (ZRX) smart contract address
	// 将 十六进制字符串 转换为以太坊的地址
	ethAddress2 := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), ethAddress2, nil) // nil is latest block
	if err != nil {
		log.Fatal("CodeAt 1, error: ", err)
	}
	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	ethAddress3 := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), ethAddress3, nil) // nil is latest block
	if err != nil {
		log.Fatal("CodeAt 2, error: ", err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
