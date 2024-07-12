package rpc

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"gfast/hdwallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/protobuf/proto"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"math/big"
	"strings"
	"time"
)

type Client struct {
	node   string
	Conn   *grpc.ClientConn
	Client api.WalletClient

	GRPC *client.GrpcClient
}

func NewClient(node string) (*Client, error) {
	c := new(Client)
	c.node = node
	c.GRPC = client.NewGrpcClient(node)
	err := c.GRPC.Start(grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("grpc client start error: %v", err)
	}
	return c, nil
}

func (c *Client) SetTimeout(timeout time.Duration) error {
	if c == nil {
		return nil
	}
	c.GRPC = client.NewGrpcClientWithTimeout(c.node, timeout)
	err := c.GRPC.Start(grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("grpc start error: %v", err)
	}
	return nil
}

/*
保持连接，如果中途连接失败，就重连
*/
func (c *Client) keepConnect() error {
	_, err := c.GRPC.GetNodeInfo()
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return c.GRPC.Reconnect(c.node)
		}
		return fmt.Errorf("node connect error: %v", err)
	}
	return nil
}

func (c *Client) TransferTrx(privateKey string, from, to string, amount decimal.Decimal, trxRemark string) (string, error) {
	var (
		err                   error
		transferTransactionEx *api.TransactionExtention
		txId                  string
		ownerKey              *ecdsa.PrivateKey
		hash                  []byte
	)
	err = c.keepConnect()
	if err != nil {
		return txId, err
	}
	decimals := decimal.NewFromInt(1000000)
	transferTransactionEx, err = c.GRPC.Transfer(from, to, amount.Mul(decimals).IntPart())
	transferTransaction := transferTransactionEx.Transaction
	if trxRemark != "" {
		transferTransaction.RawData.Data = gconv.Bytes(trxRemark)
	}
	if transferTransaction == nil ||
		len(transferTransaction.GetRawData().GetContract()) == 0 {
		return txId, fmt.Errorf("transfer error: invalid transaction")
	}
	ownerKey, _ = hdwallet.GetPrivateKeyByHexString(privateKey)
	hash, err = signTransaction(transferTransaction, ownerKey)
	if err != nil {
		return txId, err
	}
	txId = hexutil.Encode(hash)
	err = c.broadcastTransaction(transferTransaction)
	if err != nil {
		return txId, err
	}
	return txId, nil
}

func (c *Client) GetTrc10Balance(addr, assetId string) (int64, error) {
	err := c.keepConnect()
	if err != nil {
		return 0, err
	}
	acc, err := c.GRPC.GetAccount(addr)
	if err != nil || acc == nil {
		return 0, fmt.Errorf("get %s account error: %v", addr, err)
	}
	for key, value := range acc.AssetV2 {
		if key == assetId {
			return value, nil
		}
	}
	return 0, fmt.Errorf("%s do not find this assetID=%s amount", addr, assetId)
}

func (c *Client) GetTrxBalance(addr string) (*core.Account, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	return c.GRPC.GetAccount(addr)
}

func (c *Client) GetNowBlock() (*api.BlockExtention, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	return c.GRPC.GetNowBlock()
}

func (c *Client) GetBlockByNum(num int64) (*api.BlockExtention, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	return c.GRPC.GetBlockByNum(num)
}

func (c *Client) GetTrc20Balance(addr, contractAddress string) (*big.Int, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	return c.GRPC.TRC20ContractBalance(addr, contractAddress)
}

func (c *Client) TRC20Call(from, contractAddress string, data string, constant bool, feeLimit int64) (*api.TransactionExtention, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	return c.GRPC.TRC20Call(from, contractAddress, data, constant, feeLimit)
}

func (c *Client) TransferTrc10(from, to, assetId string, amount int64) (*api.TransactionExtention, error) {
	err := c.keepConnect()
	if err != nil {
		return nil, err
	}
	fromAddr, err := address.Base58ToAddress(from)
	if err != nil {
		return nil, fmt.Errorf("from address is not equal")
	}
	toAddr, err := address.Base58ToAddress(to)
	if err != nil {
		return nil, fmt.Errorf("to address is not equal")
	}
	return c.GRPC.TransferAsset(fromAddr.String(), toAddr.String(), assetId, amount)
}

func (c *Client) TransferContract(privateKey string, from, to, contract string, amount *big.Int, feeLimit int64) (string, error) {
	var (
		err                   error
		transferTransactionEx *api.TransactionExtention
		txId                  string
		ownerKey              *ecdsa.PrivateKey
		hash                  []byte
	)
	err = c.keepConnect()
	if err != nil {
		return txId, err
	}
	transferTransactionEx, err = c.GRPC.TRC20Send(from, to, contract, amount, feeLimit)
	transferTransaction := transferTransactionEx.Transaction
	if transferTransaction == nil || len(transferTransaction.GetRawData().GetContract()) == 0 {
		return txId, fmt.Errorf("transfer error: invalid transaction")
	}
	ownerKey, _ = hdwallet.GetPrivateKeyByHexString(privateKey)
	hash, err = signTransaction(transferTransaction, ownerKey)
	if err != nil {
		return txId, err
	}
	txId = hexutil.Encode(hash)
	err = c.broadcastTransaction(transferTransaction)
	if err != nil {
		return txId, err
	}
	return txId, nil
}

// signTransaction 签名交易
func signTransaction(transaction *core.Transaction, key *ecdsa.PrivateKey) ([]byte, error) {
	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	contractList := transaction.GetRawData().GetContract()
	for range contractList {
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return nil, err
		}
		transaction.Signature = append(transaction.Signature, signature)
	}
	return hash, nil
}

// broadcastTransaction 发布
func (c *Client) broadcastTransaction(transaction *core.Transaction) error {
	result, err := c.GRPC.Broadcast(transaction)
	if err != nil {
		return fmt.Errorf("broadcast transaction error: %v", err)
	}
	if result.Code != 0 {
		return fmt.Errorf("bad transaction: %v", string(result.GetMessage()))
	}
	if result.Result == true {
		return nil
	}
	d, _ := json.Marshal(result)
	return fmt.Errorf("tx send fail: %s", string(d))
}
