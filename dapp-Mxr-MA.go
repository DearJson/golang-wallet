package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"gfast/abi/dappMxrPayment"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

// MXR项目自动转账
func main() {
	var num = flag.Float64("n", 0, "转出数量")
	flag.Parse()

	if *num <= 0 {
		fmt.Println("请输入转出数量")
		return
	}
	bnbPrivateKey := "224efd3f102fee370de6d07e33527f8f9f26c5ca23a35e1ef956cb89827bd680"
	amount := decimal.NewFromFloat(*num)
	tenDecimal := decimal.NewFromInt(1000000000000000000)
	convertAmount := amount.Mul(tenDecimal).BigInt()

	rpcUrl := "https://bscrpc.com/"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Printf("连接节点失败 %v", err)
		return
	}
	defer client.Close()

	dappMxrMa, err := dappMxrPayment.NewDappMxrPayment(common.HexToAddress("0x7254DbF82077e6a0aBEE983f9C01669138E61359"), client)
	if err != nil {
		fmt.Printf("实例化合约失败 %v", err)
		return
	}
	privateKey, err := crypto.HexToECDSA(bnbPrivateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("cannot assert type: publicKey is not of type *ecdsa.PublicKey， %v", ok)
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Printf("查询nonce失败 %v", err)
		return
	}
	chainID := 56
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainID)))
	if err != nil {
		fmt.Printf("绑定链ID失败 %v", err)
		return
	}

	gasLimit := gconv.Uint64(350000) // in units
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(6) * gconv.Int64(math.Pow(10, 9)))

	transactOptss := &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    big.NewInt(gconv.Int64(nonce)),
		Signer:   auth.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}
	val, err := dappMxrMa.WithdrawToCoinAddress(transactOptss, common.HexToAddress("0xFA9aA7FB5A781F35127944dCCA50ca492d829c9C"), convertAmount)
	if err != nil {
		fmt.Printf("转账失败")
		return
	}
	hashResult := val.Hash().Hex()
	if hashResult != "" {
		fmt.Printf("%v", gconv.String(hashResult))
	}
	return
}
