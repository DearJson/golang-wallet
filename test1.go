package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"gfast/library"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	//amountString := "0x" + strings.TrimLeft(library.SubStr("0xa9059cbb0000000000000000000000001b86a64098daac00842e0c8d1cc053978a502cd70000000000000000000000000000000000000000000000030461545f7e995f09", -64, -1), "0")

	//amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).Div(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[contractAddress].Decimals)))).RoundFloor(int32(coinAddress[contractAddress].Decimals)).String()
	//dapp-sus
	//amount = decimal.NewFromBigInt(library.HexToBigInt(amountString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, 18)), 18).Mul(decimal.NewFromFloat32(0.97)).Truncate(18).String()
	//amount := decimal.NewFromBigInt(library.HexToBigInt(amountString), int32(-(18))).Truncate(18).String()

	//fmt.Printf("%s", amount)
	//
	//a, _ := library.DecryptByAes("9c13668e4475788246bb82fb363ff84f715b48e6685150c676070a93d51ee93ed6cda2fd49d1e91fbddf0e57964954de48e8adf2c18e4800ce273b9159d240eb2a1b67c85269146f55a8bfd283180e0c812a6af84bf89b6636e14e3b")
	//
	//fmt.Printf("%s\n", a)

	fmt.Printf("%s\n", library.NewGoogleAuth().GetSecret())
	a, b, err := generateAddress()
	fmt.Printf("地址： %v \n 私钥： %v \n 错误信息: %v", a, b, err)
}

func generateAddress() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])

	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}
