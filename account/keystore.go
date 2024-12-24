package account

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func GenKeystore() {
	// 生成新的 keystore account
	ks := keystore.NewKeyStore("./account/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x77633A0F12053F1656dB95BF199e2edFF98Ad4AE
}

func ImportKeystore() {
	// 从 keystore 中导入
	file := "./account/keystore/UTC--2024-12-24T08-34-31.496992000Z--cb265b8ce8fae1919fe17afd698ca4f2c19bd539"
	ks := keystore.NewKeyStore("./account/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("ReadFile error: ", err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	// 如果已经存在，会报：account already exists
	if err != nil {
		log.Fatal("Import error: ", err)
	}

	fmt.Println(account.Address.Hex()) // 0x77633A0F12053F1656dB95BF199e2edFF98Ad4AE

	if err := os.Remove(file); err != nil {
		log.Fatal("Remove error: ", err)
	}
}
