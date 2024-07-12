package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gfast/abi/dappDlc"
	cservice "gfast/app/common/service"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/util/gconv"
	"math"
	"math/big"
)

func main() {
	contractAddress := "0x02DF08D4C9CC33F24B701BC9eea501E4E43E033A"
	bnbPrivateKey := "a773c583b64c8743d240e03295c68ae4ba5d97d1950b01df95a2f485f60dac8b"

	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	contract, err := dappDlc.NewDappDlc(common.HexToAddress(contractAddress), client)

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

	if err != nil {
		fmt.Println("操作失败")
		return
	}

	gasLimit := gconv.Uint64(350000) // in units
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(6) * gconv.Int64(math.Pow(10, 9)))
	transactOpts := &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    big.NewInt(gconv.Int64(nonce)),
		Signer:   auth.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}
	uAmount := big.NewFloat(1)
	tAmount := big.NewFloat(1)
	tenDecimal := big.NewFloat(math.Pow(10, float64(18)))
	usdtAmount, _ := new(big.Float).Mul(tenDecimal, uAmount).Int(&big.Int{})
	tokenAmount, _ := new(big.Float).Mul(tenDecimal, tAmount).Int(&big.Int{})
	liquidity, err := contract.AddLiquidity(transactOpts, usdtAmount, tokenAmount)
	if err != nil {
		fmt.Printf("添加流动性失败 %v", err)
		return
	}
	hashResult := liquidity.Hash().Hex()

	if hashResult != "" {
		fmt.Printf("流动性添加成功 %v", gconv.String(hashResult))
	}
}
