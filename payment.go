package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gfast/abi/dappaco_recharge"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/util/gconv"
	"math"
	"math/big"
)

//dapp-aco每日凌晨打aco一次到提币地址钱包
func main() {
	bnbPrivateKey := "bf04a33c656711f45213cbc36ed6ac5ade54d866919d62fb15078e69c1feba6f"
	amount := big.NewFloat(28)
	tenDecimal := big.NewFloat(math.Pow(10, float64(18)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	//提币到哪个地址
	withdrawAddress := "0xA7A40e7D286443407b75275f3c173D8E8aDf8270"
	AcoAddress := "0x3cedaCB7b39B10a6fC8f656EB3cC07DFc12Ae35c"

	//合约地址
	contractAddress := "0x1B3940feB5a6923A4Cec44ddc2B1F34431e3fD55"

	rpcUrl := "https://bscrpc.com/"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Printf("连接节点失败 %v", err)
		return
	}
	defer client.Close()

	dappAcoContract, err := dappaco_recharge.NewDappaco(common.HexToAddress(contractAddress), client)
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
	val, err := dappAcoContract.ContractTransfer(transactOptss, common.HexToAddress(AcoAddress), common.HexToAddress(withdrawAddress), convertAmount)
	if err != nil {
		fmt.Printf("发起转账失败 %v", err)
		return
	}
	hashResult := val.Hash().Hex()

	if hashResult != "" {
		fmt.Println("定时任务提币到" + withdrawAddress + " 数量为" + gconv.String(amount) + "hash ===" + gconv.String(hashResult))
	}
	return
}
