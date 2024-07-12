package rpc

import (
	"encoding/json"
	"errors"
	service "gfast/app/common/service"
	sservice "gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

// NewAccount 生成新账号
func NewAccount(name string) (account NacNewAccountData, error error) {
	//开始生成钱包地址
	cache := service.Cache.New()
	password := "SFIIkq740VF72wHXEDsc4srtYVYdNp@"
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/account/new", "name="+name+"&password="+password)
	defer resp.Close()
	account = NacNewAccountData{}
	if err != nil {
		return account, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err := json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return account, err
	}
	if rpcBaseMessage.Flag == false {
		return account, err
	}
	if err := gconv.Struct(rpcBaseMessage.Data, &account); err != nil {
		return account, err
	}
	return account, nil
}

// GetLastHeight 查找最后一个块
func GetLastHeight() (blockNumber int64, error error) {
	cache := service.Cache.New()
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/block/lastHeight", "instanceId=1")
	defer resp.Close()
	blockNumber = 0
	if err != nil {
		return blockNumber, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err := json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return blockNumber, err
	}
	if rpcBaseMessage.Flag == false {
		return blockNumber, err
	}
	return gconv.Int64(rpcBaseMessage.Data), nil
}

// GetBlockDetail 获取块信息
func GetBlockDetail(blockNumber int64) (blockDetail NacBlockDetailData, error error) {
	cache := service.Cache.New()
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/block/detail", "instanceId=1&height="+gconv.String(blockNumber))
	defer resp.Close()
	//blockDetail = NacBlockDetailData{}
	if err != nil {
		return blockDetail, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err := json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return blockDetail, err
	}
	if rpcBaseMessage.Flag == false {
		return blockDetail, err
	}
	if err := gconv.Struct(rpcBaseMessage.Data, &blockDetail); err != nil {
		return blockDetail, err
	}
	return blockDetail, nil
}

// GetTxDetail 获取交易信息
func GetTxDetail(hash string) (txDetail NacTxDetail, error error) {
	cache := service.Cache.New()
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/tx/detail", "instanceId=1&hash="+hash)
	defer resp.Close()
	if err != nil {
		return txDetail, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err := json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return txDetail, err
	}
	if rpcBaseMessage.Flag == false {
		return txDetail, errors.New("未找到数据")
	}
	if err := gconv.Struct(rpcBaseMessage.Data, &txDetail); err != nil {
		return txDetail, err
	}
	return txDetail, nil
}

// GetFee 获取交易手续费
func GetFee() (fee *big.Float, err error) {
	cache := service.Cache.New()
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/gas/fee", "instanceId=1")
	defer resp.Close()
	fee = big.NewFloat(0)
	if err != nil {
		return fee, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err := json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return fee, err
	}
	if rpcBaseMessage.Flag == false {
		return fee, err
	}
	return big.NewFloat(gconv.Float64(rpcBaseMessage.Data)), nil
}

// GetBalance 获取余额
func GetBalance(address string, tokenId string) (amount decimal.Decimal, err error) {
	ctx := gctx.New()
	coinAddress, err := sservice.Currency.GetNacCoinAddress(ctx)
	inCoinAddress := false
	if _, ok := coinAddress[tokenId]; ok {
		inCoinAddress = true
	}
	if inCoinAddress == false {
		return decimal.Decimal{}, errors.New("未配置币种,无法查询")
	}
	cache := service.Cache.New()
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/account/balance", "instanceId=1&address="+address)
	defer resp.Close()
	if err != nil {
		return amount, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err = json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return amount, err
	}
	data := gconv.Map(rpcBaseMessage.Data)
	data1 := gconv.Map(data["tokenBalanceMap"])
	for k, v := range data1 {
		if k == tokenId {
			amount, _ = decimal.NewFromString(gconv.String(v))
			amount = amount.DivRound(decimal.NewFromFloat(math.Pow(10, float64(coinAddress[tokenId].Decimals))), 18)
		}
	}
	return amount, err
}

// NacWithdraw 提现
func NacWithdraw(fromAddress string, toAddress string, tokenId string, amount string, remark string) (hash string, err error) {
	cache := service.Cache.New()
	password := "SFIIkq740VF72wHXEDsc4srtYVYdNp@"
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/account/send", "instanceId=1&fromAddress="+fromAddress+"&password="+password+"&remark="+remark+"&toAddress="+toAddress+"&token="+tokenId+"&value="+amount)
	defer resp.Close()
	if err != nil {
		return hash, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err = json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return hash, err
	}
	if rpcBaseMessage.Flag == false {
		return hash, err
	}
	nacWithdrawData := NacWithdrawData{}
	if err := gconv.Struct(rpcBaseMessage.Data, &nacWithdrawData); err != nil {
		return hash, err
	}
	return nacWithdrawData.Hash, err
}

// NacGetGasFee 获取手续费率
func NacGetGasFee() (fee decimal.Decimal, err error) {
	cache := service.Cache.New()
	resp, err := g.Client().Get(gconv.String(cache.Get("nac_rpc_url"))+"/station/api/gas/fee", "instanceId=1")
	defer resp.Close()
	if err != nil {
		return fee, err
	}
	content := resp.ReadAllString()
	rpcBaseMessage := new(NacBaseMessage)
	if err = json.Unmarshal(gconv.Bytes(content), &rpcBaseMessage); err != nil {
		return fee, err
	}
	return decimal.NewFromString(gconv.String(rpcBaseMessage.Data))
}
