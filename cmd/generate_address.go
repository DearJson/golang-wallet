package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"gfast/hdwallet"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 生成助记词
	mnemonic := hdwallet.RandSeed()

	// 获取私钥字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:] // 移除0x前缀

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换公钥")
	}

	// 生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// 输出结果
	fmt.Println("BSC地址生成成功:")
	fmt.Println("地址:", address)
	fmt.Println("私钥:", privateKeyHex)
	fmt.Println("助记词:", mnemonic)
	fmt.Println("\n注意: 请妥善保管私钥，不要泄露给任何人！")
}
