package account

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenerageWallet() {
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

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0xdcFee110209d0fAc223a88D40b16F8Da02846f36
}
