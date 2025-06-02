package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"gfast/abi/erc20"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/sha3"
	"math/rand"
	"strings"
	"time"
)

type eth struct {
	WebBase
}

var Eth = new(eth)

// GenerateAddress 生成ETH地址
func (b *eth) GenerateAddress(r *ghttp.Request) {
	var req *dao.AddressAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//判断一下该用户标识的bsc链是否存在
	exists, err := service.Address.GetInfoByUserId(r.GetCtx(), req.UserId, "eth")
	if exists != nil {
		result := make(map[string]string)
		result["user_id"] = exists.UserId
		result["address"] = exists.Address
		b.SusJsonExit(r, result)
	}

	//开始生成钱包地址
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		b.FailJsonExit(r, "生成失败")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		b.FailJsonExit(r, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])

	req.MainChain = "eth"
	req.PrivateKey, _ = library.EncryptByAes(gconv.Bytes(hexutil.Encode(privateKeyBytes)[2:]))
	req.Address = address
	err = service.Address.Add(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	result := make(map[string]string)
	result["user_id"] = req.UserId
	result["address"] = req.Address
	b.SusJsonExit(r, result)
}

// Withdraw 提现代币
func (b *eth) Withdraw(r *ghttp.Request) {
	var req *dao.WithdrawAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "eth")
	if coinInfo == nil {
		b.FailJsonExit(r, "暂未配置该币种,无法转账")
	}

	minAmount, _ := cservice.SysConfig.GetConfigByKey("sys.minWithdrawAudit")
	if req.Amount <= gconv.Float64(minAmount.ConfigValue) {
		req.Status = 2
	} else {
		req.Status = 1
	}
	req.MainChain = "eth"
	req.CoinToken = coinInfo.Name
	req.Address = strings.ToLower(req.Address)
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, req.Status, req.Nonce1)
	err := service.Withdraw.Add(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	b.SusJsonExit(r)
}

func (b *eth) ResetHash(r *ghttp.Request) {
	var (
		ethClient = rpc.EthClient{}
	)
	hash := gconv.String(r.GetPost("hash"))
	blockData, _ := ethClient.Init().GetTransactionByHash(hash)
	ethStruct := rpc.EthTransactions{}
	err := mapstructure.Decode(blockData, &ethStruct)
	if err != nil {
		g.Log().Printf("ETH解析块失败，%v %v", hash, err)
		return
	}
	//解析监控交易的地址
	contractRecharge := g.Cfg().GetBool("eth.contract_recharge")
	contractRechargeAddress := g.Cfg().GetString("eth.contract_address")
	addressRecharge := g.Cfg().GetBool("eth.address_recharge")
	ctx := gctx.New()
	userAddress, _ := service.Address.GetEthAllAddress(ctx)
	coinAddress, _ := service.Currency.GetEthCoinAddress(ctx)
	queueExchange := &amqp.QueueExchange{
		QuName: _const.EthSweepQuName,
		RtKey:  _const.EthSweepRtKey,
		ExName: _const.EthSweepExName,
		ExType: _const.EthSweepExType,
	}
	mq := amqp.New(queueExchange)

	bs := false
	if contractRecharge && ethStruct.To == contractRechargeAddress {
		bs = true
	}
	inCoinAddress := false
	if _, ok := coinAddress[ethStruct.To]; ok {
		inCoinAddress = true
	}
	if addressRecharge && (library.ElementIsInSlice(ethStruct.To, userAddress) || inCoinAddress) {
		bs = true
	}
	if bs {
		jsonByte, _ := json.Marshal(ethStruct)
		t := &(producer.BscProducer{Msg: string(jsonByte)})
		mq.RegisterProducer(t)
		g.Log().File("producer.{Y-m-d}.log").Printf("生产 %v", ethStruct.Hash)
	}
	mq.Start()
	b.SusJsonExit(r)
}

// BalanceOf 查询代币余额
func (b *eth) BalanceOf(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("eth_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)

	if r.GetPost("contract_address") == "0x1000000000000000000000000000000000000000" {
		balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(gconv.String(r.GetPost("address"))), nil)
		result := make(map[string]string)
		result["balance"] = gconv.String(balance.String())
		b.SusJsonExit(r, result)
	} else {
		erc20Contract, err := erc20.NewErc20AbiCaller(common.HexToAddress(gconv.String(r.GetPost("contract_address"))), client)
		if err != nil {
			b.FailJsonExit(r, "查询失败")
		}
		transactOpts := &bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		}
		balance, _ := erc20Contract.BalanceOf(transactOpts, common.HexToAddress(gconv.String(r.GetPost("address"))))
		result := make(map[string]string)
		result["balance"] = balance.String()
		b.SusJsonExit(r, result)
	}
}
