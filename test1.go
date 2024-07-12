package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/sha3"
	"log"
	"os"
)

func main() {

	fileHandler, err := os.Create("./test.txt")
	if nil != err {
		panic(err)
	}
	defer fileHandler.Close()

	for i := 0; i <= 2000; i++ {

		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		pKeys := hexutil.Encode(privateKeyBytes)[2:]
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		hash := sha3.NewLegacyKeccak256()
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		hash.Write(publicKeyBytes[1:])
		address = hexutil.Encode(hash.Sum(nil)[12:])

		//使用WriteString
		_, err = fileHandler.WriteString(address + "=========" + pKeys + "\n")
		fmt.Println("已生成" + gconv.String(i) + "个账号")
	}

}

func checkErr(err error) {
	if nil != err {
		panic(err)
	}
}
