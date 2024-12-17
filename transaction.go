package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Transaction() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------------------- transaction ---------------------")

	// query blocks
	// get header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("header number: ", header.Number.String()) // 21320787

	// read full block info
	// include: block number, block timestamp, block hash, block difficulty,
	// as well as the list of transactions and much much more
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block number: ", block.Number().Uint64())               // 5671744
	fmt.Println("block time: ", block.Time())                            // 1527211625
	fmt.Println("block difficulty: ", block.Difficulty().Uint64())       // 3217000136609065
	fmt.Println("block hash: ", block.Hash().Hex())                      // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("block transaction count: ", block.Transactions().Len()) // 144
	// or
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block transaction count by TransactionCount: ", count) // 144

	// query transactions
	// get chainId for getting sender address
	ctx := context.Background()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID: ", chainID) // 1

	for idx, tx := range block.Transactions() {
		fmt.Println("transaction -----idx----: ", idx)
		fmt.Println("transaction hex hash: ", tx.Hash().Hex())         // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("transaction value: ", tx.Value().String())        // 10000000000000000
		fmt.Println("transaction gas: ", tx.Gas())                     // 105000
		fmt.Println("transaction gas price: ", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("transaction nonce: ", tx.Nonce())                 // 110644
		fmt.Println("transaction data: ", string(tx.Data()))           //
		fmt.Println("transaction to hex: ", tx.To().Hex())             // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		// read sender address
		// AsMessage is deprecated
		// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		//     fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		// }
		senderAddress, err := client.TransactionSender(ctx, tx, block.Hash(), uint(idx))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("sender address: ", senderAddress.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258

		// get receipt info
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receipt status: ", receipt.Status) // 1 1:succes 0:fail
		fmt.Println("receipt logs: ", receipt.Logs)     // []

		// too much data, so equal 2 to break
		if idx == 2 {
			break
		}
	}

	// another way: to iterate over transaction without fetching the block info
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(ctx, block.Hash(), idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("way2 tx hex hash: ", tx.Hash().Hex())
		if idx == 2 {
			break
		}
	}

	// query single transaction
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("single query, tx hex hash: ", tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println("single query, isPending: ", isPending)         // false

	fmt.Println("--------------------- transfer eth ---------------------")
	// transfer ETH from one account to another account
	// a transaction include:
	// amount that want to transfer, gas limit, gas price, nonce, receive address and data(optional)
	// NOTE: transaction must be signed with the private key
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transfer, nonce: ", nonce) // 8

	// set amount
	value := big.NewInt(1000000000000000000) // in wei(=1 eth), 18个0
	gasLimit := uint64(21000)                // in units
	// gasPrice := big.NewInt(30000000000)      // in wei(=30 gwei) 10个0
	// gas prices are always fluctuating based on market demand and what users are willing to pay
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")

	tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	// sign the transaction with private key of the sender
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// begin to send transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e
	// end transfer

	fmt.Println("--------------------- transfer token ---------------------")
}
