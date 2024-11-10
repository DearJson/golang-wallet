package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"gfast/abi/erc20"
	"gfast/abi/lp"
	"gfast/abi/lp3"
	"gfast/abi/swap"
	"gfast/abi/uniswapv3"
	"gfast/amqp"
	_const "gfast/amqp/const"
	"gfast/amqp/producer"
	"gfast/app/common/global"
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	webDao "gfast/app/web/dao"
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
	"github.com/shopspring/decimal"
	"github.com/snowbai/uniswapv3-sdk/examples/helper"
	"github.com/storyicon/sigverify"
	"golang.org/x/crypto/sha3"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

type bsc struct {
	WebBase
}

var Bsc = new(bsc)

// GenerateAddress 生成币安地址
func (b *bsc) GenerateAddress(r *ghttp.Request) {
	var req *dao.AddressAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//判断一下该用户标识的bsc链是否存在
	exists, err := service.Address.GetInfoByUserId(r.GetCtx(), req.UserId, "bsc")
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

	req.MainChain = "bsc"
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
func (b *bsc) Withdraw(r *ghttp.Request) {
	var req *dao.WithdrawAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "bsc")
	if coinInfo == nil {
		b.FailJsonExit(r, "暂未配置该币种,无法转账")
	} else {
		if req.Amount <= coinInfo.MinWithdraw {
			req.Status = 2
		} else {
			req.Status = 1
		}
		req.MainChain = "bsc"
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
}

// WithdrawNft 提现NFT
func (b *bsc) WithdrawNft(r *ghttp.Request) {
	var req *dao.WithdrawNftAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "bsc")
	if coinInfo == nil {
		b.FailJsonExit(r, "暂未配置该币种,无法转账")
	}
	//该币种不是NFT
	if coinInfo.Decimals != 0 {
		b.FailJsonExit(r, "该币种非NFT,无法转账")
	}
	req.MainChain = "bsc"
	req.CoinToken = coinInfo.Name
	req.Address = strings.ToLower(req.Address)
	minAmount, _ := cservice.SysConfig.GetConfigByKey("sys.minWithdrawAudit")
	if req.Amount <= gconv.Float64(minAmount.ConfigValue) {
		req.Status = 2
	} else {
		req.Status = 1
	}
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, req.Status, req.Nonce1)
	err := service.Withdraw.AddNft(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	b.SusJsonExit(r)
}

func (b *bsc) UpdateWithdrawStatus(r *ghttp.Request) {
	//id := gconv.Uint64(r.GetPost("id"))
	//if id == 0 {
	//	b.FailJsonExit(r, "未传递记录ID")
	//}
	//var withdraw *
	//g.Model("withdraw").Where("main_chain","bsc").Where("id",id)

}

// TokenInfo 查询币种信息
func (b *bsc) TokenInfo(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
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
	name, _ := erc20Contract.Name(transactOpts)
	decimal, _ := erc20Contract.Decimals(transactOpts)
	totalSupply, _ := erc20Contract.TotalSupply(transactOpts)
	b.SusJsonExit(r, g.Map{
		"name":        name,
		"decimals":    decimal,
		"totalSupply": totalSupply.String(),
	})
}

// LpInfo 查询lP信息
func (b *bsc) LpInfo(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	lpContract, err := lp.NewLpAbi(common.HexToAddress(gconv.String(r.GetPost("contract_address"))), client)
	if err != nil {
		b.FailJsonExit(r, "查询失败")
	}
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	reserves, err := lpContract.GetReserves(transactOpts)
	TotalSupply, err := lpContract.TotalSupply(transactOpts)
	token0, err := lpContract.Token0(transactOpts)
	token1, err := lpContract.Token1(transactOpts)

	result := make(map[string]string)
	result["reserve0"] = reserves.Reserve0.String()
	result["reserve1"] = reserves.Reserve1.String()
	result["totalSupply"] = gconv.String(TotalSupply.String())
	result["token0"] = gconv.String(token0.String())
	result["token1"] = gconv.String(token1.String())

	b.SusJsonExit(r, result)
}

// Lp3Info 查询lP3信息
func (b *bsc) Lp3Info(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	lp3Contract, err := lp3.NewSwap(common.HexToAddress(gconv.String(r.GetPost("contract_address"))), client)
	if err != nil {
		b.FailJsonExit(r, "查询失败")
	}
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	fee, err := lp3Contract.ProtocolFees(transactOpts)
	token0, err := lp3Contract.Token0(transactOpts)
	token1, err := lp3Contract.Token1(transactOpts)

	result := make(map[string]string)
	result["token0Fee"] = fee.Token0.String()
	result["token1Fee"] = fee.Token1.String()
	result["token0"] = gconv.String(token0.String())
	result["token1"] = gconv.String(token1.String())
	b.SusJsonExit(r, result)
}

// BalanceOf 查询代币余额
func (b *bsc) BalanceOf(r *ghttp.Request) {
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
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
		result["balance"] = gconv.String(balance.String())
		b.SusJsonExit(r, result)
	}
}

// SetAddress 设置监控地址
func (b *bsc) SetAddress(r *ghttp.Request) {
	address := strings.ToLower(gconv.String(r.GetPost("address")))
	count, _ := g.Model("address").Where("main_chain", "bsc").Where("address", address).FindCount()
	if count > 0 {
		b.FailJsonExit(r, "地址已存在")
	}
	g.Model("address").Data(g.Map{"main_chain": "bsc", "user_id": gconv.String(r.GetPost("user_id")), "address": address}).Insert()
	cache := cservice.Cache.New()
	cache.Remove(global.BscUserAddressList)
	cache.Remove(global.TronUserAddressList)
	b.SusJsonExit(r)
}

// MergeAmount 单独归集某个地址
func (b *bsc) MergeAmount(r *ghttp.Request) {
	requestAddress := strings.ToLower(gconv.String(r.GetPost("address")))
	requestAmount := gconv.Float64(r.GetPost("amount"))
	contractAddress := gconv.String(r.GetPost("contract_address"))

	//查询币种
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), contractAddress, "bsc")
	if coinInfo == nil {
		b.FailJsonExit(r, "暂未配置该币种,无法转账")
	}

	g.Model("recharge").Data(g.Map{
		"main_chain":       "bsc",
		"block_hash":       0,
		"coin_token":       coinInfo.Name,
		"from_address":     "0x0",
		"to_address":       requestAddress,
		"amount":           requestAmount,
		"contract_address": coinInfo.ContractAddress,
		"hash":             "0x0" + gconv.String(time.Now().UnixNano()),
		"remarks":          "系统归集手续费",
	}).Insert()
	b.SusJsonExit(r)
}

func (b *bsc) ResetBlock(r *ghttp.Request) {
	var (
		bscClient = rpc.BscClient{}
	)
	current := gconv.Int64(r.GetPost("block_number"))
	blockData, _ := bscClient.Init().GetBlockByNumber(current)
	bscStruct := rpc.BscBlock{}
	err := mapstructure.Decode(blockData, &bscStruct)
	if err != nil {
		g.Log().Printf("币安解析块失败，%v %v", current, err)
		return
	}
	//解析监控交易的地址
	monitorSwap := g.Cfg().GetBool("bsc.monitor_swap")
	monitorSwapRouteAddress := g.Cfg().GetString("bsc.monitor_swap_route_address")
	monitorSwapCoinAddress, _ := cservice.MonitorCurrency.GetAddress("bsc")
	contractRecharge := g.Cfg().GetBool("bsc.contract_recharge")
	contractRechargeAddress := g.Cfg().GetString("bsc.contract_address")
	addressRecharge := g.Cfg().GetBool("bsc.address_recharge")
	ctx := gctx.New()
	userAddress, _ := service.Address.GetBnbAllAddress(ctx)
	coinAddress, _ := service.Currency.GetBnbCoinAddress(ctx)
	queueExchange := &amqp.QueueExchange{
		QuName: _const.BscSweepQuName,
		RtKey:  _const.BscSweepRtKey,
		ExName: _const.BscSweepExName,
		ExType: _const.BscSweepExType,
	}
	mq := amqp.New(queueExchange)

	monitorSwapExchange := &amqp.QueueExchange{
		QuName: _const.BscMonitorSwapQuName,
		RtKey:  _const.BscMonitorSwapRtKey,
		ExName: _const.BscMonitorSwapExName,
		ExType: _const.BscMonitorSwapExType,
	}
	monitorMq := amqp.New(monitorSwapExchange)
	for _, value := range bscStruct.Transactions {
		bs := false
		if contractRecharge && value.To == contractRechargeAddress {
			bs = true
		}
		inCoinAddress := false
		if _, ok := coinAddress[value.To]; ok {
			inCoinAddress = true
		}
		if addressRecharge && (library.ElementIsInSlice(value.To, userAddress) || inCoinAddress) {
			bs = true
		}

		if bs {
			jsonByte, _ := json.Marshal(value)
			t := &(producer.BscProducer{Msg: string(jsonByte)})
			mq.RegisterProducer(t)
			g.Log().File("producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
		}

		//监控swap交易
		if monitorSwap && monitorSwapRouteAddress == value.To {
			for _, monitorCoinAddress := range monitorSwapCoinAddress {

				if value.Hash == "0x2406397303da83cbf06dc4db073a35ab3cf48f64882f01226d90f1c7aaf2f469" {
					fmt.Printf("%v \n %v\n", value.Input, monitorCoinAddress[2:])
				}

				if strings.Contains(value.Input, monitorCoinAddress[2:]) {
					fmt.Println("进来了")
					jsonByte, _ := json.Marshal(value)
					t := &(producer.BscMonitorSwapProducer{Msg: string(jsonByte)})
					monitorMq.RegisterProducer(t)
					g.Log().File("monitor_producer.{Y-m-d}.log").Printf("生产 %v", value.Hash)
				}
			}
		}
	}
	mq.Start()
	monitorMq.Start()
	b.SusJsonExit(r)
}

func (b *bsc) ResetHash(r *ghttp.Request) {
	var (
		bscClient = rpc.BscClient{}
	)
	hash := gconv.String(r.GetPost("hash"))
	blockData, _ := bscClient.Init().GetTransactionByHash(hash)
	bscStruct := rpc.BscTransactions{}
	err := mapstructure.Decode(blockData, &bscStruct)
	if err != nil {
		g.Log().Printf("币安解析块失败，%v %v", hash, err)
		return
	}
	//解析监控交易的地址
	monitorSwap := g.Cfg().GetBool("bsc.monitor_swap")
	monitorSwapRouteAddress := g.Cfg().GetString("bsc.monitor_swap_route_address")
	monitorSwapCoinAddress, _ := cservice.MonitorCurrency.GetAddress("bsc")
	contractRecharge := g.Cfg().GetBool("bsc.contract_recharge")
	contractRechargeAddress := g.Cfg().GetString("bsc.contract_address")
	addressRecharge := g.Cfg().GetBool("bsc.address_recharge")
	ctx := gctx.New()
	userAddress, _ := service.Address.GetBnbAllAddress(ctx)
	coinAddress, _ := service.Currency.GetBnbCoinAddress(ctx)
	queueExchange := &amqp.QueueExchange{
		QuName: _const.BscSweepQuName,
		RtKey:  _const.BscSweepRtKey,
		ExName: _const.BscSweepExName,
		ExType: _const.BscSweepExType,
	}
	mq := amqp.New(queueExchange)

	monitorSwapExchange := &amqp.QueueExchange{
		QuName: _const.BscMonitorSwapQuName,
		RtKey:  _const.BscMonitorSwapRtKey,
		ExName: _const.BscMonitorSwapExName,
		ExType: _const.BscMonitorSwapExType,
	}
	monitorMq := amqp.New(monitorSwapExchange)

	bs := false
	if contractRecharge && bscStruct.To == contractRechargeAddress {
		bs = true
	}
	inCoinAddress := false
	if _, ok := coinAddress[bscStruct.To]; ok {
		inCoinAddress = true
	}
	if addressRecharge && (library.ElementIsInSlice(bscStruct.To, userAddress) || inCoinAddress) {
		bs = true
	}

	if bs {
		jsonByte, _ := json.Marshal(bscStruct)
		t := &(producer.BscProducer{Msg: string(jsonByte)})
		mq.RegisterProducer(t)
		g.Log().File("producer.{Y-m-d}.log").Printf("生产 %v", bscStruct.Hash)
	}

	//监控swap交易
	if monitorSwap && monitorSwapRouteAddress == bscStruct.To {
		for _, monitorCoinAddress := range monitorSwapCoinAddress {

			if bscStruct.Hash == "0x2406397303da83cbf06dc4db073a35ab3cf48f64882f01226d90f1c7aaf2f469" {
				fmt.Printf("%v \n %v\n", bscStruct.Input, monitorCoinAddress[2:])
			}

			if strings.Contains(bscStruct.Input, monitorCoinAddress[2:]) {
				fmt.Println("进来了")
				jsonByte, _ := json.Marshal(bscStruct)
				t := &(producer.BscMonitorSwapProducer{Msg: string(jsonByte)})
				monitorMq.RegisterProducer(t)
				g.Log().File("monitor_producer.{Y-m-d}.log").Printf("生产 %v", bscStruct.Hash)
			}
		}
	}

	mq.Start()
	monitorMq.Start()
	b.SusJsonExit(r)
}

// InsertCurrency 新增币种
func (b *bsc) InsertCurrency(r *ghttp.Request) {
	contractAddress := strings.ToLower(gconv.String(r.GetPost("contract_address")))
	count, _ := g.Model("currency").Where("main_chain", "bsc").Where("contract_address", contractAddress).Count()
	if count > 0 {
		b.SusJsonExit(r)
	}
	name := gconv.Float64(r.GetPost("name"))
	g.Model("currency").Data(g.Map{
		"main_chain":       "bsc",
		"name":             name,
		"contract_address": contractAddress,
		"decimals":         0,
	}).Insert()
	b.SusJsonExit(r)
}

// GetCoinPrice 获取币种价格
func (b *bsc) GetCoinPrice(r *ghttp.Request) {
	var req *webDao.GetCoinPriceReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	swapContract, _ := swap.NewSwap(common.HexToAddress(req.RouteAddress), client)
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	result, _ := swapContract.GetAmountsOut(transactOpts, big.NewInt(int64(math.Pow10(req.AmountInDecimals))), req.Path)

	tokenPrice := decimal.NewFromBigInt(result[1], 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(req.AmountInDecimals))), int32(req.AmountInDecimals)).String()
	b.SusJsonExit(r, g.Map{
		"price": tokenPrice,
	})
}

// V3GetCoinPrice 查询v3价格
func (b *bsc) V3GetCoinPrice(r *ghttp.Request) {
	//var req *webDao.V3GetCoinPriceReq
	cache := cservice.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, err := ethclient.Dial(rpcUrl)
	uniSwapV3Contract, err := uniswapv3.NewUniswapv3(common.HexToAddress("0xB048Bbc1Ee6b733FFfCFb9e9CeF7375518e25997"), client)
	if err != nil {
		b.FailJsonExit(r, "查询失败")
	}
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	token1 := common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	token2 := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	fee := big.NewInt(100)
	amountIn := helper.FloatStringToBigInt("1.00", 18)
	sqrtPriceLimitX96 := big.NewInt(0)
	var out []interface{}
	rawCaller := uniswapv3.Uniswapv3Raw{Contract: uniSwapV3Contract}
	fmt.Printf("%v", uniswapv3.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           token1,
		TokenOut:          token2,
		AmountIn:          amountIn,
		Fee:               fee,
		SqrtPriceLimitX96: sqrtPriceLimitX96,
	})
	err = rawCaller.Call(transactOpts, &out, "quoteExactInputSingle", uniswapv3.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           token1,
		TokenOut:          token2,
		Fee:               fee,
		AmountIn:          amountIn,
		SqrtPriceLimitX96: sqrtPriceLimitX96,
	})
	if err != nil {
		b.FailJsonExit(r, "查询失败"+err.Error())
	} else {
		//price := ConverDeci
		//
		//price := helper.ConvertDecimal(out[0].(*big.Int), consts.Erc20Decimal)
		//fmt.Println("price get :", price)
		b.SusJsonExit(r, g.Map{
			"price": out[0].(*big.Int),
		})
	}
}

// GetUserNft 获取用户持有NFT信息
//func (b *bsc) GetUserNft(r *ghttp.Request) {
//	var req *webDao.GetUserNftReq
//	//获取参数
//	if err := r.Parse(&req); err != nil {
//		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
//	}
//
//	cache := cservice.Cache.New()
//	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
//	client, err := ethclient.Dial(rpcUrl)
//	nftContract, err := nft.NewNftAbi(common.HexToAddress(req.ContractAddress), client)
//	if err != nil {
//		b.FailJsonExit(r, "查询失败")
//	}
//	transactOpts := &bind.CallOpts{
//		Pending:     false,
//		From:        common.Address{},
//		BlockNumber: nil,
//		Context:     nil,
//	}
//	//查询用户NFT数量
//	nftCount, _ := nftContract.GetUserCardCount(transactOpts, common.HexToAddress(req.Address))
//	var (
//		nftList []string
//		i       int64
//	)
//	if nftCount.Int64() > 0 {
//		for i = 0; i < nftCount.Int64(); i++ {
//			tokenId, _ := nftContract.TokenOfOwnerByIndex(transactOpts, common.HexToAddress(req.Address), big.NewInt(i))
//			nftList = append(nftList, tokenId.String())
//		}
//	}
//	b.SusJsonExit(r, g.Map{"count": nftCount, "list": nftList})
//}

// GetTransferDetail 根据hash获取交易详情
func (b *bsc) GetTransferDetail(r *ghttp.Request) {
	hash := strings.ToLower(gconv.String(r.GetPost("hash")))
	if hash == "" {
		b.FailJsonExit(r, "请传递hash值")
	}
	client := rpc.BscClient{}
	bscStruct := rpc.BscTransferDetail{}
	data, err := client.Init().GetTransferStatus(hash)
	if err != nil {
		b.FailJsonExit(r, "查询失败")
	}
	err = mapstructure.Decode(data, &bscStruct)
	if err != nil || bscStruct.Status == "" {
		b.FailJsonExit(r, "查询失败")
	}
	nowBlock, err := client.Init().GetNowBlock()
	if err != nil {
		b.FailJsonExit(r, "查询失败")
	}
	b.SusJsonExit(r, g.Map{
		"block_hash":          bscStruct.BlockHash,
		"block_number":        bscStruct.BlockNumber,
		"contract_address":    bscStruct.ContractAddress,
		"cumulative_gas_used": "",
		"from":                bscStruct.From,
		"gas_used":            bscStruct.GasUsed,
		"status":              bscStruct.Status,
		"to":                  bscStruct.To,
		"transaction_hash":    bscStruct.TransactionHash,
		"transaction_index":   bscStruct.TransactionIndex,
		"type":                bscStruct.Type,
		"now_block":           nowBlock,
	})
}

// GetAutoAddLiquidityDetail 查询充值合约自动买币自动添加流动性LP数量
func (b *bsc) GetAutoAddLiquidityDetail(r *ghttp.Request) {
	hash := strings.ToLower(gconv.String(r.GetPost("hash")))
	client := rpc.BscClient{}
	bscStruct := rpc.BscTransferDetail{}
	data, err := client.Init().GetTransferStatus(hash)
	if err != nil {
		b.FailJsonExit(r, "查询失败")
	}
	err = mapstructure.Decode(data, &bscStruct)
	if err != nil || bscStruct.Status == "" {
		b.FailJsonExit(r, "查询失败")
	}
	var (
		amount0 string
		amount1 string
		lpNum   string
	)
	for _, log := range bscStruct.Logs {
		if log.Topics[0] == "0xc3d478e7ae5305a4ec4c9758311fa7659d9f0f67b8d7a894d0a74b7b2f7bc847" {
			amount0String := "0x" + log.Data[2:66]
			amount0 = decimal.NewFromBigInt(library.HexToBigInt(amount0String), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()

			amount1String := "0x" + log.Data[66:130]
			amount1 = decimal.NewFromBigInt(library.HexToBigInt(amount1String), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()

			lpString := "0x" + log.Data[130:]
			lpNum = decimal.NewFromBigInt(library.HexToBigInt(lpString), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
		}
	}
	b.SusJsonExit(r, g.Map{
		"amount0": amount0,
		"amount1": amount1,
		"lp_num":  lpNum,
	})
}

// GetAutoTradeDetail 查询充值合约自动买币的数量
func (b *bsc) GetAutoTradeDetail(r *ghttp.Request) {
	hash := strings.ToLower(gconv.String(r.GetPost("hash")))
	client := rpc.BscClient{}
	bscStruct := rpc.BscTransferDetail{}
	data, err := client.Init().GetTransferStatus(hash)
	if err != nil {
		return
	}
	err = mapstructure.Decode(data, &bscStruct)
	if err != nil || bscStruct.Status == "" {
		return
	}
	data0 := bscStruct.Logs[len(bscStruct.Logs)-1].Data
	data1 := decimal.NewFromBigInt(library.HexToBigInt(data0), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
	b.SusJsonExit(r, g.Map{
		"amount": data1,
	})
}

// SignVerify 签名校验工具
func (b *bsc) SignVerify(r *ghttp.Request) {
	message := gconv.String(r.GetPost("message"))
	address := gconv.String(r.GetPost("address"))
	signMessage := gconv.String(r.GetPost("sign_message"))

	fmt.Printf("收到数据 \nmessage=%v \naddress=%v \nsignMessage=%v", message, address, signMessage)
	//VerifyEllipticCurveHexSignatureEx
	valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(common.HexToAddress(address), []byte(message), signMessage)
	if err != nil {
		b.FailJsonExit(r, "校验失败,请稍后")
	}
	b.SusJsonExit(r, g.Map{
		"valid": valid,
	})
}

func (b *bsc) GetDetailByHash(r *ghttp.Request) {

}
