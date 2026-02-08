package rpc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gfast/app/common/service"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strconv"
)

type EthClient struct {
	client *rpc.Client
}

func (b *EthClient) Init() *EthClient {
	//获取连接与eth客户端
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("eth_rpc_url"))
	b.client, _ = rpc.Dial(rpcUrl)
	return b
}

// GetNowBlock  获取最新板块
func (b *EthClient) GetNowBlock() (block int64, err error) {
	var _block string
	err = b.client.Call(&_block, "eth_blockNumber")
	if err != nil {
		return -1, err
	}
	block, _ = strconv.ParseInt(_block, 0, 64)
	return block, nil
}

// GetBlockByNumber 根据块号获取块所有交易信息
func (b *EthClient) GetBlockByNumber(blockNumber int64) (block interface{}, err error) {
	var _block interface{}
	_blockNumber := "0x" + strconv.FormatInt(blockNumber, 16)
	err = b.client.Call(&_block, "eth_getBlockByNumber", _blockNumber, true)
	if err != nil {
		return _block, err
	}
	return _block, nil
}

// GetTransferStatus 根据hash获取交易详情
func (b *EthClient) GetTransferStatus(hash string) (transferData interface{}, err error) {
	var _transferData interface{}
	err = b.client.Call(&_transferData, "eth_getTransactionReceipt", hash)
	if err != nil {
		return _transferData, err
	}
	return _transferData, err
}

// GetTransactionByHash 根据hash获取交易详情
func (b *EthClient) GetTransactionByHash(hash string) (transferData interface{}, err error) {
	var _transferData interface{}
	err = b.client.Call(&_transferData, "eth_getTransactionByHash", hash)
	if err != nil {
		return _transferData, err
	}
	return _transferData, err
}

// GetTransactionCount 查询指定地址的交易量
func (b *EthClient) GetTransactionCount(address string) (nonce interface{}, err error) {
	var _nonce interface{}
	err = b.client.Call(&_nonce, "eth_getTransactionCount", address, "latest")
	if err != nil {
		return _nonce, err
	}
	return _nonce, err
}

// TransferEth 转账eth (支持EIP-1559 Type 2交易)
func TransferEth(privateKeys string, amount *big.Int, toAddress string, MaxNonce uint64) (transferData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("eth_rpc_url"))
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
		return nil, 0, fmt.Errorf("public key assertion failed")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

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

	// Get actual chain ID from RPC
	actualChainID, err := client.NetworkID(context.Background())
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to get network ID: %v", err)
		return nil, 0, err
	}

	chainID := actualChainID // Use actual chain ID instead of config

	// 获取最新的区块头以检查是否支持EIP-1559
	head, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to get latest header: %v", err)
		return nil, 0, err
	}

	toAddresss := common.HexToAddress(toAddress)

	var tx *types.Transaction

	// 如果区块支持EIP-1559 (BaseFee存在)，使用Type 2交易
	if head.BaseFee != nil {
		// 获取Gas费用信息
		gasTipCap, err := client.SuggestGasTipCap(context.Background())
		if err != nil {
			g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to get gas tip cap: %v", err)
			return nil, 0, err
		}

		// 计算gasFeeCap = baseFee * 2 + gasTipCap
		gasFeeCap := new(big.Int).Add(
			gasTipCap,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)

		// 估算Gas Limit
		msg := ethereum.CallMsg{
			From:      fromAddress,
			To:        &toAddresss,
			GasFeeCap: gasFeeCap,
			GasTipCap: gasTipCap,
			Value:     amount,
			Data:      nil,
		}
		gasLimit, err := client.EstimateGas(context.Background(), msg)
		if err != nil {
			g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to estimate gas: %v", err)
			gasLimit = 21000 // 默认值
		} else {
			gasLimit = gasLimit * 110 / 100 // 增加10%缓冲
		}

		// 创建EIP-1559交易 (Type 2)
		tx = types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Gas:       gasLimit,
			To:        &toAddresss,
			Value:     amount,
			Data:      nil,
		})
	} else {
		// 回退到Legacy交易 (Type 0)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to get gas price: %v", err)
			return nil, 0, err
		}
		gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(12))
		gasPrice = new(big.Int).Div(gasPrice, big.NewInt(10))

		// 估算Gas Limit
		msg := ethereum.CallMsg{
			From:     fromAddress,
			To:       &toAddresss,
			GasPrice: gasPrice,
			Value:    amount,
			Data:     nil,
		}
		gasLimit, err := client.EstimateGas(context.Background(), msg)
		if err != nil {
			g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to estimate gas: %v", err)
			gasLimit = 21000 // 默认值
		} else {
			gasLimit = gasLimit * 110 / 100 // 增加10%缓冲
		}

		tx = types.NewTransaction(nonce, toAddresss, amount, gasLimit, gasPrice, nil)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("transaction signing failed: %v", err)
		return nil, 0, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("交易广播失败: %v", err)
		return nil, 0, err
	}

	return signedTx.Hash().Hex(), nonce, nil
}

// TransferEthToken 转账其他代币
func TransferEthToken(privateKeys string, amount *big.Int, toAddress string, tokenAddress string, MaxNonce uint64) (transferData interface{}, currentNonce uint64, err error) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("eth_rpc_url"))
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

	// 动态获取Gas Price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to get gas price: %v", err)
		return nil, 0, err
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(12))
	gasPrice = new(big.Int).Div(gasPrice, big.NewInt(10))

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

	// 动态估算Gas Limit
	bnbValue := big.NewInt(0)
	msg := ethereum.CallMsg{
		From:     fromAddress,
		To:       &tokenAddresss,
		GasPrice: gasPrice,
		Value:    bnbValue,
		Data:     data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		g.Log().File("withdraw.{Y-m-d}.log").Printf("failed to estimate gas: %v", err)
		gasLimit = 100000 // ERC20 transfer默认值
	} else {
		gasLimit = gasLimit * 110 / 100 // 增加10%缓冲
	}

	tx := types.NewTransaction(nonce, tokenAddresss, bnbValue, gasLimit, gasPrice, data)
	chainID := big.NewInt(g.Cfg().GetInt64("eth.chain_id"))
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
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
