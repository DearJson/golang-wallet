package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"gfast/abi/erc20"
	"gfast/abi/lp"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"gfast/hdwallet"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/sha3"
	"math/rand"
	"strings"
	"time"
)

type tron struct {
	WebBase
}

var Tron = new(tron)

// TronGenerateAddress 生成波场地址
func (t *tron) TronGenerateAddress(r *ghttp.Request) {
	var req *dao.AddressAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		t.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//判断一下该用户标识的bsc链是否存在
	exists, err := service.Address.GetInfoByUserId(r.GetCtx(), req.UserId, "tron")
	if exists != nil {
		result := make(map[string]string)
		result["user_id"] = exists.UserId
		result["address"] = exists.Address
		t.SusJsonExit(r, result)
	}

	//开始生成钱包地址
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.FailJsonExit(r, "生成失败")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.FailJsonExit(r, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])
	tronAddress, _ := hdwallet.FromHexAddress("41" + address[2:])

	req.MainChain = "tron"
	req.PrivateKey, _ = library.EncryptByAes(gconv.Bytes(hexutil.Encode(privateKeyBytes)[2:]))
	req.Address = tronAddress
	err = service.Address.Add(r.GetCtx(), req)
	if err != nil {
		t.FailJsonExit(r, err.Error())
	}
	result := make(map[string]string)
	result["user_id"] = req.UserId
	result["address"] = req.Address
	t.SusJsonExit(r, result)
}

// TronWithdraw 提现代币
func (t *tron) TronWithdraw(r *ghttp.Request) {
	var req *dao.WithdrawAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		t.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "tron")
	if coinInfo == nil {
		t.FailJsonExit(r, "暂未配置该币种,无法转账")
	}

	minAmount, _ := cservice.SysConfig.GetConfigByKey("sys.minWithdrawAudit")
	if req.Amount <= gconv.Float64(minAmount.ConfigValue) {
		req.Status = 2
	} else {
		req.Status = 1
	}
	req.MainChain = "tron"
	req.CoinToken = coinInfo.Name
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, req.Status, req.Nonce1)
	err := service.Withdraw.Add(r.GetCtx(), req)
	if err != nil {
		t.FailJsonExit(r, err.Error())
	}
	t.SusJsonExit(r)
}

// TronWithdrawNft 提现NFT
func (t *tron) TronWithdrawNft(r *ghttp.Request) {
	var req *dao.WithdrawNftAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		t.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "tron")
	if coinInfo == nil {
		t.FailJsonExit(r, "暂未配置该币种,无法转账")
	}
	//该币种不是NFT
	if coinInfo.Decimals != 0 {
		t.FailJsonExit(r, "该币种非NFT,无法转账")
	}
	req.Amount = 1
	req.MainChain = "tron"
	req.CoinToken = coinInfo.Name
	req.Address = strings.ToLower(req.Address)
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	var status int8
	status = 1
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, status, req.Nonce1)
	err := service.Withdraw.AddNft(r.GetCtx(), req)
	if err != nil {
		t.FailJsonExit(r, err.Error())
	}
	t.SusJsonExit(r)
}

// TronTokenInfo 查询币种信息
func (t *tron) TronTokenInfo(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	erc20Contract, err := erc20.NewErc20AbiCaller(common.HexToAddress(gconv.String(r.GetPost("contract_address"))), client)
	if err != nil {
		t.FailJsonExit(r, "查询失败")
	}
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	name, _ := erc20Contract.Name(transactOpts)
	decimal, _ := erc20Contract.Decimals(transactOpts)
	totalSupply, _ := erc20Contract.TotalSupply(transactOpts)
	t.SusJsonExit(r, g.Map{
		"name":        name,
		"decimals":    decimal,
		"totalSupply": totalSupply.String(),
	})
}

// TronLpInfo 查询lP信息
func (t *tron) TronLpInfo(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	lpContract, err := lp.NewLpAbi(common.HexToAddress(gconv.String(r.GetPost("contract_address"))), client)
	if err != nil {
		t.FailJsonExit(r, "查询失败")
	}
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	reserves, err := lpContract.GetReserves(transactOpts)
	TotalSupply, err := lpContract.TotalSupply(transactOpts)

	result := make(map[string]string)
	result["reserve0"] = reserves.Reserve0.String()
	result["reserve1"] = reserves.Reserve1.String()
	result["totalSupply"] = gconv.String(TotalSupply.String())
	t.SusJsonExit(r, result)
}

// TronBalanceOf 查询代币余额
func (t *tron) TronBalanceOf(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)

	if r.GetPost("contract_address") == "0x1000000000000000000000000000000000000000" {
		balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(gconv.String(r.GetPost("address"))), nil)
		result := make(map[string]string)
		result["balance"] = gconv.String(balance.String())
		t.SusJsonExit(r, result)
	} else {
		erc20Contract, err := erc20.NewErc20AbiCaller(common.HexToAddress(gconv.String(r.GetPost("contract_address"))), client)
		if err != nil {
			t.FailJsonExit(r, "查询失败")
		}
		transactOpts := &bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		}
		balance, _ := erc20Contract.BalanceOf(transactOpts, common.HexToAddress(gconv.String(r.GetPost("address"))))
		result := make(map[string]string)
		result["balance"] = gconv.String(balance.String())
		t.SusJsonExit(r, result)
	}
}

// TronSetAddress 设置监控地址
func (t *tron) TronSetAddress(r *ghttp.Request) {
	address := gconv.String(r.GetPost("address"))
	count, _ := g.Model("address").Where("main_chain", "tron").Where("address", address).FindCount()
	if count > 0 {
		t.FailJsonExit(r, "地址已存在")
	}
	g.Model("address").Data(g.Map{"main_chain": "tron", "user_id": gconv.String(r.GetPost("user_id")), "address": address}).Insert()
	cache := cservice.Cache.New()
	cache.Remove(global.TronUserAddressList)
	t.SusJsonExit(r)
}

func (t *tron) ResetBlock(r *ghttp.Request) {
	var (
		cache = cservice.Cache.New()
	)
	tronClient, err := rpc.NewClient(gconv.String(cache.Get("tron_rpc_url")))
	if err != nil {
		g.Log().Printf("波场节点链接失败！！！%v", err)
		return
	}
	ctx := gctx.New()
	userAddress, _ := service.Address.GetTronAllAddress(ctx)
	coinAddress, _ := service.Currency.GetTronCoinAddress(ctx)
	queueExchange := &amqp.QueueExchange{
		QuName: _const.TronSweepQuName,
		RtKey:  _const.TronSweepQuName,
		ExName: _const.TronSweepQuName,
		ExType: _const.TronSweepQuName,
	}
	mq := amqp.New(queueExchange)
	var (
		current       int64
		bs            bool
		data          string
		blockData     *api.BlockExtention
		inCoinAddress bool
		toAddress     string
		method        string
	)
	contractRecharge := g.Cfg().GetBool("tron.contract_recharge")
	contractRechargeAddress := g.Cfg().GetString("tron.contract_address")
	addressRecharge := g.Cfg().GetBool("tron.address_recharge")
	current = gconv.Int64(r.GetPost("block_number"))
	blockData, err = tronClient.GetBlockByNum(current)
	if err != nil {
		g.Log().Printf("波场解析块失败，%v %v", current, err)
		return
	}
	for _, value := range blockData.GetTransactions() {

		if value.Result == nil || !value.Result.Result {
			continue
		}
		rets := value.Transaction.Ret
		if len(rets) < 1 || rets[0].ContractRet != core.Transaction_Result_SUCCESS {
			continue
		}
		bs = false
		inCoinAddress = false
		if value.Transaction.GetRawData().GetContract()[0].Type == core.Transaction_Contract_TriggerSmartContract {
			unObj := &core.TriggerSmartContract{}
			err = proto.Unmarshal(value.GetTransaction().GetRawData().GetContract()[0].GetParameter().GetValue(), unObj)
			if err != nil {
				continue
			}
			contractAddress := hdwallet.EncodeCheck(unObj.GetContractAddress())

			if contractRecharge && contractAddress == contractRechargeAddress {
				bs = true
			}

			if _, ok := coinAddress[contractAddress]; ok {
				inCoinAddress = true
			}
			if addressRecharge && inCoinAddress {
				data = hex.EncodeToString(unObj.GetData())
				method = data[0:8]
				if (method == "a9059cbb" || method == "23b872dd") && len(data) >= 136 {
					if method == "a9059cbb" {
						toAddress, _ = hdwallet.FromHexAddress("41" + data[32:72])
					} else {
						toAddress, _ = hdwallet.FromHexAddress("41" + data[96:136])
					}
					if library.ElementIsInSlice(toAddress, userAddress) {
						bs = true
					}
				}
			}
		}
		//这里只有TRX充值才能进来
		if addressRecharge && value.Transaction.GetRawData().GetContract()[0].Type == core.Transaction_Contract_TransferContract {
			unObj1 := &core.TransferContract{}
			err = proto.Unmarshal(value.GetTransaction().GetRawData().GetContract()[0].GetParameter().GetValue(), unObj1)
			if err != nil {
				continue
			}
			toAddress := hdwallet.EncodeCheck(unObj1.ToAddress)
			if library.ElementIsInSlice(toAddress, userAddress) {
				bs = true
			}
		}

		if bs {
			aa, _ := proto.Marshal(value)
			msg := rpc.TronTransaction{
				Transaction: aa,
				BlockHeader: blockData.BlockHeader.RawData.Number,
				BlockHash:   hex.EncodeToString(blockData.Blockid),
			}
			dd, _ := json.Marshal(msg)
			t := &(producer.TronProducer{
				Msg: string(dd),
			})
			mq.RegisterProducer(t)
			g.Log().File("tron-producer.{Y-m-d}.log").Printf("生产 %v", hex.EncodeToString(value.Txid))
		}
	}
	mq.Start()
	t.SusJsonExit(r)
}
