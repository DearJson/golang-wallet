package contract

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// WithdrawParams 提现参数
type WithdrawParams struct {
	AbiKey   string        // 提现方法的标识
	DataABI  string        // 用于编码 data 参数的 ABI JSON 字符串
	DataArgs []interface{} // data 参数的实际值
}

// WithdrawContract 提现合约交互
type WithdrawContract struct {
	client       *ethclient.Client
	contractAddr common.Address
	auth         *bind.TransactOpts
}

// NewWithdrawContract 创建提现合约实例
func NewWithdrawContract(rpcURL string, privateKey string, contractAddress string) (*WithdrawContract, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(contractAddress)
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(56)) // BSC主网chainID
	if err != nil {
		return nil, err
	}

	return &WithdrawContract{
		client:       client,
		contractAddr: contractAddr,
		auth:         auth,
	}, nil
}

// Withdraw 执行提现操作
func (w *WithdrawContract) Withdraw(ctx context.Context, params WithdrawParams) (string, error) {
	// 1. 解析 data 参数的 ABI
	parsedABI, err := abi.JSON(strings.NewReader(params.DataABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse data ABI: %v", err)
	}

	// 2. 编码 data 参数
	data, err := parsedABI.Pack("", params.DataArgs...)
	if err != nil {
		return "", fmt.Errorf("failed to pack data arguments: %v", err)
	}

	// 3. 构造调用数据
	methodID := crypto.Keccak256([]byte("withdraw(string,bytes)"))[:4]
	abiKeyBytes := []byte(params.AbiKey)
	abiKeyLen := big.NewInt(int64(len(abiKeyBytes)))
	abiKeyPadded := common.LeftPadBytes(abiKeyBytes, 32)
	dataLen := big.NewInt(int64(len(data)))
	dataPadded := common.LeftPadBytes(data, 32)

	// 4. 组合所有参数
	var input []byte
	input = append(input, methodID...)
	input = append(input, common.LeftPadBytes(abiKeyLen.Bytes(), 32)...)
	input = append(input, abiKeyPadded...)
	input = append(input, common.LeftPadBytes(dataLen.Bytes(), 32)...)
	input = append(input, dataPadded...)

	// 5. 发送交易
	tx := types.NewTransaction(
		w.auth.Nonce.Uint64(),
		w.contractAddr,
		big.NewInt(0), // value
		300000,        // gas limit
		w.auth.GasPrice,
		input,
	)

	// 6. 签名交易
	signedTx, err := w.auth.Signer(w.auth.From, tx)
	if err != nil {
		return "", err
	}

	// 7. 发送交易
	err = w.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// Close 关闭客户端连接
func (w *WithdrawContract) Close() {
	if w.client != nil {
		w.client.Close()
	}
}
