package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"gfast/abi/dapp5aPayment"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/util/gconv"
	"math"
	"math/big"
)

//dapp-5a每日凌晨打5a一次到提币地址钱包
func main() {
	//判断是否有传递参数进来
	var num = flag.Float64("n", 0, "转出数量")
	flag.Parse()

	if *num <= 0 {
		fmt.Println("请输入转出数量")
		return
	}

	bnbPrivateKey := "a489483907a650555a6dbd7da4848da2ba132e7a17715d0baf663da39a699b0c"
	amount := big.NewFloat(*num)
	tenDecimal := big.NewFloat(math.Pow(10, float64(18)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	rpcUrl := "https://bscrpc.com/"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Printf("连接节点失败 %v", err)
		return
	}
	defer client.Close()

	dapp5aContract, err := dapp5aPayment.NewDapp5aPayment(common.HexToAddress("0x761a9C084380d551185cA48C68e0786669811A6D"), client)
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
	val, err := dapp5aContract.WithdrawToCoinAddress(transactOptss, convertAmount)
	if err != nil {
		fmt.Printf("发起转账失败 %v", err)
		return
	}
	hashResult := val.Hash().Hex()
	if hashResult != "" {
		fmt.Println("定时任务转帐 数量为" + gconv.String(amount) + "hash ===" + gconv.String(hashResult))
	}
	return
}
