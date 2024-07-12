package rpc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gfast/abi/newNft"
	"gfast/abi/swap"
	token2 "gfast/abi/token"
	"gfast/app/common/service"
	"gfast/library"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
	"math"
	"math/big"
	"strconv"
	"time"
)

type BscClient struct {
	client *rpc.Client
}

func (b *BscClient) Init() *BscClient {
	//获取连接与eth客户端
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	b.client, _ = rpc.Dial(rpcUrl)
	return b
}

// GetNowBlock  获取最新板块
func (b *BscClient) GetNowBlock() (block int64, err error) {
	var _block string
	err = b.client.Call(&_block, "eth_blockNumber")
	if err != nil {
		return -1, err
	}
	block, _ = strconv.ParseInt(_block, 0, 64)
	return block, nil
}

// GetBlockByNumber 根据块号获取块所有交易信息
func (b *BscClient) GetBlockByNumber(blockNumber int64) (block interface{}, err error) {
	var _block interface{}
	_blockNumber := "0x" + strconv.FormatInt(blockNumber, 16)
	err = b.client.Call(&_block, "eth_getBlockByNumber", _blockNumber, true)
	if err != nil {
		return _block, err
	}
	return _block, nil
}

// GetTransferStatus 根据hash获取交易详情
func (b *BscClient) GetTransferStatus(hash string) (transferData interface{}, err error) {
	var _transferData interface{}
	err = b.client.Call(&_transferData, "eth_getTransactionReceipt", hash)
	if err != nil {
		return _transferData, err
	}
	return _transferData, err
}

// GetTransactionByHash 根据hash获取交易详情
func (b *BscClient) GetTransactionByHash(hash string) (transferData interface{}, err error) {
	var _transferData interface{}
	err = b.client.Call(&_transferData, "eth_getTransactionByHash", hash)
	if err != nil {
		return _transferData, err
	}
	return _transferData, err
}

// GetTransactionCount 查询指定地址的交易量
func (b *BscClient) GetTransactionCount(address string) (nonce interface{}, err error) {
	var _nonce interface{}
	err = b.client.Call(&_nonce, "eth_getTransactionCount", address, "latest")
	if err != nil {
		return _nonce, err
	}
	return _nonce, err
}

// TransferBnb 转账bnb
func TransferBnb(privateKeys string, amount *big.Int, toAddress string, MaxNonce uint64) (transferData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	defer client.Close()
	privateKey, err := crypto.HexToECDSA(privateKeys)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		g.Log().File("withdraw.{Y-m-d}.log").Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, 0, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("来源地址%v", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	//检查一下，最新的nonce是否大于数据库的最后一条nonce+1,如果是，取这个获取的
	if MaxNonce > 0 {
		NextNonce := gconv.Uint64(MaxNonce) + 1
		if nonce < NextNonce {
			nonce = NextNonce
		}
	}
	gasLimitConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
	gasLimit := gconv.Uint64(gasLimitConfig.ConfigValue) // in units
	gasPriceConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasPrice")
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(gasPriceConfig.ConfigValue) * gconv.Int64(math.Pow(10, 9)))

	toAddresss := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, toAddresss, amount, gasLimit, gasPrice, nil)
	chainID := g.Cfg().GetInt64("bsc.chain_id")
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Println("加密1是失败")
		return nil, 0, err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("交易广播失败" + err.Error())
		return nil, 0, err
	}

	return signedTx.Hash().Hex(), nonce, nil
}

// TransferToken 转账其他代币
func TransferToken(privateKeys string, amount *big.Int, toAddress string, tokenAddress string, MaxNonce uint64) (transferData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	defer client.Close()
	privateKey, err := crypto.HexToECDSA(privateKeys)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		g.Log().File("withdraw.{Y-m-d}.log").Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, 0, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	//检查一下，最新的nonce是否大于数据库的最后一条nonce+1,如果是，取这个获取的
	if gconv.Uint64(MaxNonce) > 0 {
		NextNonce := gconv.Uint64(MaxNonce) + 1
		if nonce < NextNonce {
			nonce = NextNonce
		}
	}
	gasLimitConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
	gasLimit := gconv.Uint64(gasLimitConfig.ConfigValue) // in units
	gasPriceConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasPrice")
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(gasPriceConfig.ConfigValue) * gconv.Int64(math.Pow(10, 9)))

	toAddresss := common.HexToAddress(toAddress)
	tokenAddresss := common.HexToAddress(tokenAddress)
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddresss.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	bnbValue := big.NewInt(0)
	tx := types.NewTransaction(nonce, tokenAddresss, bnbValue, gasLimit, gasPrice, data)
	chainID := g.Cfg().GetInt64("bsc.chain_id")
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Println("加密1是失败")
		return nil, 0, err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Println("交易广播失败 %v", err)
		return nil, 0, err
	}
	return signedTx.Hash().Hex(), nonce, nil
}

// TransferTokenOr 转账特殊OR代币
func TransferTokenOr(privateKeys string, amount *big.Int, toAddress string, tokenAddress string, MaxNonce uint64) (transferData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	privateKey, err := crypto.HexToECDSA(privateKeys)
	defer client.Close()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		g.Log().File("withdraw.{Y-m-d}.log").Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, 0, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	//检查一下，最新的nonce是否大于数据库的最后一条nonce+1,如果是，取这个获取的
	if gconv.Uint64(MaxNonce) > 0 {
		NextNonce := gconv.Uint64(MaxNonce) + 1
		if nonce < NextNonce {
			nonce = NextNonce
		}
	}
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	tokenContract, err := token2.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	chainID := g.Cfg().GetInt64("bsc.chain_id")
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	gasLimitConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
	gasLimit := gconv.Uint64(gasLimitConfig.ConfigValue) // in units
	gasPriceConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasPrice")
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(gasPriceConfig.ConfigValue) * gconv.Int64(math.Pow(10, 9)))

	transactOpts := &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    decimal.NewFromInt(int64(nonce)).BigInt(),
		Signer:   auth.Signer,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Value:    big.NewInt(0),
		Context:  auth.Context,
		NoSend:   false,
	}
	result, err := tokenContract.Wd(transactOpts, common.HexToAddress(toAddress), amount)
	if err != nil {
		return nil, 0, err
	}
	return result.Hash().Hex(), nonce, nil
}

// SwapToken 自动购买代币并转账
func SwapToken(privateKeys string, amount *big.Int, toAddress string, SwapRoute string, SwapPath string, MaxNonce uint64) (transferData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	privateKey, err := crypto.HexToECDSA(privateKeys)
	defer client.Close()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		g.Log().File("withdraw.{Y-m-d}.log").Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, 0, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	//检查一下，最新的nonce是否大于数据库的最后一条nonce+1,如果是，取这个获取的
	if gconv.Uint64(MaxNonce) > 0 {
		NextNonce := gconv.Uint64(MaxNonce) + 1
		if nonce < NextNonce {
			nonce = NextNonce
		}
	}
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}

	fmt.Printf("nonce = %v", nonce)

	swapContract, err := swap.NewSwap(common.HexToAddress(SwapRoute), client)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	chainID := g.Cfg().GetInt64("bsc.chain_id")
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	gasLimitConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
	gasLimit := gconv.Uint64(gasLimitConfig.ConfigValue) // in units
	gasPriceConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasPrice")
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(gasPriceConfig.ConfigValue) * gconv.Int64(math.Pow(10, 9)))

	transactOpts := &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    decimal.NewFromInt(int64(nonce)).BigInt(),
		Signer:   auth.Signer,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Value:    big.NewInt(0),
		Context:  auth.Context,
		NoSend:   false,
	}

	swapPathTemp := library.Explode(",", SwapPath)
	res := make([]common.Address, len(swapPathTemp))
	for index, val := range swapPathTemp {
		res[index] = common.HexToAddress(val)
	}
	tokens, err := swapContract.SwapExactTokensForTokens(transactOpts, amount, decimal.NewFromInt(0).BigInt(), res, common.HexToAddress(toAddress), decimal.NewFromInt(time.Now().Unix()+300).BigInt())
	if err != nil {
		return nil, 0, err
	}
	return tokens.Hash().Hex(), nonce, nil
}

// SafeTransferFrom 提现NFT给某个用户
func SafeTransferFrom(privateKeys string, toAddress string, tokenAddress string, tokenId *big.Int, _num *big.Int, MaxNonce uint64) (nftData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	defer client.Close()
	//nftContract, err := nft.NewNftAbi(common.HexToAddress(tokenAddress), client)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	privateKey, err := crypto.HexToECDSA(privateKeys)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		g.Log().File("withdraw.{Y-m-d}.log").Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, 0, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	chainID := g.Cfg().GetInt64("bsc.chain_id")
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}

	gasLimitConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasLimit")
	gasLimit := gconv.Uint64(gasLimitConfig.ConfigValue) // in units
	gasPriceConfig, _ := service.SysConfig.GetConfigByKey("sys.bnbGasPrice")
	gasPrice := new(big.Int)
	gasPrice = big.NewInt(gconv.Int64(gasPriceConfig.ConfigValue) * gconv.Int64(math.Pow(10, 9)))

	//检查一下，最新的nonce是否大于数据库的最后一条nonce+1,如果是，取这个获取的
	if MaxNonce > 0 {
		NextNonce := gconv.Uint64(MaxNonce) + 1
		if nonce < NextNonce {
			nonce = NextNonce
		}
	}
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
	nftContract, err := newNft.NewNewNft(common.HexToAddress(tokenAddress), client)
	val, err := nftContract.SafeTransferFrom(transactOptss, fromAddress, common.HexToAddress(toAddress), tokenId, _num, nil)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("%v", err)
		return nil, 0, err
	}
	return val.Hash().Hex(), nonce, nil

}
