package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	dapp_bamboo "gfast/abi/dappBamboo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/util/gconv"
	"math"
	"math/big"
)

// 竹子项目自动转账
func main() {
	var num = flag.Float64("n", 0, "转出数量")
	flag.Parse()

	if *num <= 0 {
		fmt.Println("请输入转出数量")
		return
	}
	shoubiAddress := common.HexToAddress("0xd1f33e67c8a794a8e31986a8d246e649b1b1103c")
	bnbPrivateKey := "9047ed1fd6d8913100a26286f28f2df06f1cea098c8e045d50c5812a4866bbf5"
	amount := big.NewInt(gconv.Int64(num))

	rpcUrl := "https://bscrpc.com/"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Printf("连接节点失败 %v", err)
		return
	}
	defer client.Close()

	dappBamboo, err := dapp_bamboo.NewDappBamboo(common.HexToAddress("0x3f44496594bee96fd31a78df4d45cf182bce0d70"), client)
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
	val, err := dappBamboo.Mint(transactOptss, shoubiAddress, amount)
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
