package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Address() {
	fmt.Println("--------------------- address ---------------------")
	// get account
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Println("address: ", address.Hex())                        // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println("address hash: ", common.HexToHash(address.Hex())) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println("address bytes: ", address.Bytes())                // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

	// get account balance
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance: ", balance) // 41040912992482858613

	// get account balance by block num
	blockNumber := big.NewInt(5532993)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance by blcok number: ", balance) // 25729324269165216042

	// ETH的最小单位 wei，为了方便阅读，可以转为 ETH， 转换计算：wei / 10^18
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("ETH value: ", ethValue) // 25.729324269165216041

	// get account with pending(submitting or waiting for a transaction to be confirmed)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pending balance: ", pendingBalance) // 41040912992482858613

	// gen private key( used for signing transactions)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// convert to hexadeciaml string and remove prefix 0x
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key, hex string: ", hexutil.Encode(privateKeyBytes)[2:]) // f3221b77267054a11c9bd98e57ac09b8270c547a49e1dfa362fb02aae066ff5f

	// generate public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x and 04
	// 590fe0a9eb3b68036f4f02c0fe04062493a9ca7c59fac7c986c0136b7e9bd6f4cb8633c053193b4a706200806c1d9f8605619582bc3578af27efbadc624b55ed

	// generate public key address
	pubAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("public key address: ", pubAddress) // 0xdcFee110209d0fAc223a88D40b16F8Da02846f36

	// address check
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("address is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("address is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	// check if address is an Account or a Smart Contract
	// 0x Protocol Token (ZRX) smart contract address
	address2 := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address2, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address3 := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address3, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
