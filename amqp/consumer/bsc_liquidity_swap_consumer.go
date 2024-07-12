package consumer

import (
	"encoding/json"
	"gfast/abi/erc20"
	"gfast/app/common/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	decoder "github.com/mingjingc/abi-decoder"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"math"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type BscLiquiditySwapConsumer struct{}

// Consumer 实现币安扫块消费者
func (b *BscLiquiditySwapConsumer) Consumer(dataByte []byte, key uint64) error {
	transfer := rpc.BscTransactions{}
	err := json.Unmarshal(dataByte, &transfer)
	if err != nil {
		return err
	} //json解析到结构体里面
	g.Log().File("liquidity_consumer.{Y-m-d}.log").Printf("消费 %v", transfer.Hash)

	bscClient := rpc.BscClient{}
	transferReceipt, err := bscClient.Init().GetTransferStatus(transfer.Hash)
	transferDetail := rpc.BscTransferDetail{}
	err = mapstructure.Decode(transferReceipt, &transferDetail)
	if err != nil || transferDetail.Status == "" {
		g.Log().Printf("查询交易%v失败，%v# \n", transfer.Hash, err)
		return err
	}
	if transferDetail.Status == "0x0" {
		g.Log().Printf("交易状态是失败的，%v#不处理  %v#\n", transfer.Hash, err)
		return nil
	}

	//判断下到底是LP交易转移还是增加流动性还是移除流动性
	abiString, _ := os.ReadFile("./abi/swap/swap.abi")
	txDataDecoder := decoder.NewABIDecoder()
	txDataDecoder.SetABI(string(abiString))
	method, err := txDataDecoder.DecodeMethod(transfer.Input)

	var (
		coin_token        string
		coin1_token       string
		coin_decimals     uint8
		coin1_decimals    uint8
		contract_address  string
		contract_address1 string
		amount            string
		amount1           string
		lp_address        string
		fee_num           string
		lp_num            string
		to_address        string
		typeIndex         int8
	)

	if method.Name == "addLiquidity" {
		//添加流动性
		typeIndex = 1
		for _, value := range method.Params {
			switch value.Name {
			case "tokenA":
				contract_address = strings.ToLower(value.Value)
				break
			case "tokenB":
				contract_address1 = strings.ToLower(value.Value)
				break
			case "to":
				to_address = strings.ToLower(value.Value)
				break
			}
		}
		coin_token, coin_decimals = getContract(contract_address)
		coin1_token, coin1_decimals = getContract(contract_address1)
		amountUint := decimal.NewFromFloat(0)
		amount1Uint := decimal.NewFromFloat(0)
		for _, log := range transferDetail.Logs {
			if log.Topics[0] == "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f" {
				lp_address = log.Address
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address {
				amountUint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin_decimals))), 18).Add(amountUint)
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address1 {
				amount1Uint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin1_decimals))), 18).Add(amount1Uint)
			}
		}
		amount = amountUint.String()
		amount1 = amount1Uint.String()
		for _, log := range transferDetail.Logs {
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == lp_address {
				if log.Topics[2][26:] == strings.ToLower(to_address)[2:] {
					lp_num = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
				} else {
					fee_num = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
				}
			}
		}
	} else if method.Name == "addLiquidityETH" {
		//添加流动性
		typeIndex = 1
		for _, value := range method.Params {
			switch value.Name {
			case "token":
				contract_address = strings.ToLower(value.Value)
				break
			case "to":
				to_address = strings.ToLower(value.Value)
				break
			}
		}
		contract_address1 = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
		coin_token, coin_decimals = getContract(contract_address)
		coin1_token, coin1_decimals = getContract(contract_address1)
		amountUint := decimal.NewFromFloat(0)
		amount1Uint := decimal.NewFromFloat(0)
		for _, log := range transferDetail.Logs {
			if log.Topics[0] == "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f" {
				lp_address = log.Address
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address {
				amountUint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin_decimals))), 18).Add(amountUint)
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address1 {
				amount1Uint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin1_decimals))), 18).Add(amount1Uint)
			}
		}
		amount = amountUint.String()
		amount1 = amount1Uint.String()
		for _, log := range transferDetail.Logs {
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == lp_address {
				if log.Topics[2][26:] == strings.ToLower(to_address)[2:] {
					lp_num = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
				} else {
					fee_num = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
				}
			}
		}
	} else if method.Name == "removeLiquidity" || method.Name == "removeLiquidityWithPermit" {
		//移除流动性
		typeIndex = 2
		for _, value := range method.Params {
			switch value.Name {
			case "tokenA":
				contract_address = strings.ToLower(value.Value)
				break
			case "tokenB":
				contract_address1 = strings.ToLower(value.Value)
				break
			case "to":
				to_address = strings.ToLower(value.Value)
				break
			case "liquidity":
				lp, _ := decimal.NewFromString(value.Value)
				lp_num = lp.DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
			}
		}
		coin_token, coin_decimals = getContract(contract_address)
		coin1_token, coin1_decimals = getContract(contract_address1)
		amountUint := decimal.NewFromFloat(0)
		amount1Uint := decimal.NewFromFloat(0)
		for _, log := range transferDetail.Logs {
			if log.Topics[0] == "0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496" {
				lp_address = log.Address
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address {
				amountUint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin_decimals))), 18).Add(amountUint)
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address1 {
				amount1Uint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin1_decimals))), 18).Add(amount1Uint)
			}
		}
		amount = amountUint.String()
		amount1 = amount1Uint.String()
	} else if method.Name == "removeLiquidityETH" || method.Name == "removeLiquidityETHWithPermit" || method.Name == "removeLiquidityETHSupportingFeeOnTransferTokens" || method.Name == "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens" {
		//移除流动性
		typeIndex = 2
		for _, value := range method.Params {
			switch value.Name {
			case "token":
				contract_address = strings.ToLower(value.Value)
				break
			case "to":
				to_address = strings.ToLower(value.Value)
				break
			case "liquidity":
				lp, _ := decimal.NewFromString(value.Value)
				lp_num = lp.DivRound(decimal.NewFromFloat(math.Pow(10, float64(18))), 18).String()
			}
		}
		contract_address1 = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
		coin_token, coin_decimals = getContract(contract_address)
		coin1_token, coin1_decimals = getContract(contract_address1)
		amountUint := decimal.NewFromFloat(0)
		amount1Uint := decimal.NewFromFloat(0)
		for _, log := range transferDetail.Logs {
			if log.Topics[0] == "0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496" {
				lp_address = log.Address
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address {
				amountUint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin_decimals))), 18).Add(amountUint)
			}
			if log.Topics[0] == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" && log.Address == contract_address1 {
				amount1Uint = decimal.NewFromBigInt(library.HexToBigInt(log.Data), 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(coin1_decimals))), 18).Add(amount1Uint)
			}
		}
		amount = amountUint.String()
		amount1 = amount1Uint.String()
	}

	blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)

	exists, _ := g.Model("liquidity_swap_trade").Where("hash", transfer.Hash).Count()
	if exists <= 0 {
		g.Model("liquidity_swap_trade").Data(g.Map{"main_chain": "bsc", "block_hash": transfer.BlockHash, "from_address": transfer.From,
			"coin_token": coin_token, "coin_token1": coin1_token, "coin_token_decimals": coin_decimals, "coin_token1_decimals": coin1_decimals,
			"contract_address": contract_address, "contract_address1": contract_address1, "amount": amount, "amount1": amount1, "lp_address": lp_address,
			"fee_num": fee_num, "lp_num": lp_num, "to_address": to_address, "type": typeIndex, "hash": transfer.Hash, "block_height": blockHeight,
			"created_at": time.Now().Format("2006-01-02 15:04:05")}).Insert()

		callbackUrl, _ := service.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
		if callbackUrl.ConfigValue == "" {
			g.Log().File("callback.{Y-m-d}.log").Printf("未配置回调地址,不发送请求 %v \n", transfer.Hash)
			return nil
		}

		postForm := url.Values{
			"main_chain":        {"bsc"},
			"callback_type":     {"liquidity"},
			"block_hash":        {transfer.BlockHash},
			"from_address":      {transfer.From},
			"coin_token":        {coin_token},
			"coin1_token":       {coin1_token},
			"contract_address":  {contract_address},
			"contract_address1": {contract_address1},
			"amount":            {amount},
			"amount1":           {amount1},
			"lp_address":        {lp_address},
			"fee_num":           {fee_num},
			"lp_num":            {lp_num},
			"to_address":        {to_address},
			"type":              {gconv.String(typeIndex)},
			"hash":              {transfer.Hash},
			"block_height":      {blockHeight},
		}
		resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, postForm)
		defer resp.Body.Close()
		g.Log().File("callback.{Y-m-d}.log").Printf("发送监控流动性回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", callbackUrl.ConfigValue, postForm.Encode(), resp.StatusCode)
	}

	return nil
}

func getContract(contractAddress string) (contractName string, contractDecimals uint8) {
	cache := service.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	erc20Contract, err := erc20.NewErc20AbiCaller(common.HexToAddress(contractAddress), client)
	if err != nil {
		return "", 0
	}
	transactOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	contractName, _ = erc20Contract.Symbol(transactOpts)
	contractDecimals, _ = erc20Contract.Decimals(transactOpts)
	return contractName, contractDecimals
}
