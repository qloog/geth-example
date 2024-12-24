package account

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountBalance() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

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
}
