package test

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
)

func SimulateClient() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	testAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	sim := simulated.NewBackend(types.GenesisAlloc{
		testAddr: {Balance: balance},
	})
	defer sim.Close()

	client := sim.Client()

	chainid, _ := client.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	head, _ := client.HeaderByNumber(context.Background(), nil) // Should be child's, good enough
	gasPrice := new(big.Int).Add(head.BaseFee, big.NewInt(params.GWei))

	// 使用下面的会报错：max priority fee per gas higher than max fee per gas
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainid,
		Nonce:     nonce,
		GasTipCap: big.NewInt(params.GWei),
		GasFeeCap: gasPrice,
		Gas:       21000,
		To:        &toAddress,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainid), privateKey)
	if err != nil {
		log.Fatal("types.SignTx, error: ", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("SendTransaction error: ", err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51

	sim.Commit()

	block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
	if err != nil {
		log.Fatalf("could not get block at height 1: %v", err)
	}
	fmt.Printf("block hash: %v\n", block.Transactions()[0].Hash())
	fmt.Printf("signedTx hash: %v\n", signedTx.Hash())

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	fmt.Printf("status: %v\n", receipt.Status) // status: 1

	// Output:
	// tx sent: 0x31db33a141987637386719f414a78dc1ca49ed2c16adf3e959bf094a53b9a2fc
	// block hash: 0x31db33a141987637386719f414a78dc1ca49ed2c16adf3e959bf094a53b9a2fc
	// signedTx hash: 0x31db33a141987637386719f414a78dc1ca49ed2c16adf3e959bf094a53b9a2fc
	// status: 1
}
