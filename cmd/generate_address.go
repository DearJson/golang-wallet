package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/sha3"
)

func main() {
	//fmt.Printf("%v", library.NewGoogleAuth().GetSecret())

	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])
	PrivateKey := gconv.String(gconv.Bytes(hexutil.Encode(privateKeyBytes)[2:]))
	fmt.Printf("地址： %v \n", address)
	fmt.Printf("私钥： %v \n", PrivateKey)
}
