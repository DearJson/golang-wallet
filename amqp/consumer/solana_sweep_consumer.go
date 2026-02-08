package consumer

import (
	"encoding/json"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	sservice "gfast/app/system/service"
	"gfast/rpc"
	"net/url"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
)

// SolanaSweepConsumer Solana扫块消费者
type SolanaSweepConsumer struct{}

// Consumer 实现Solana扫块消费者
func (s *SolanaSweepConsumer) Consumer(dataByte []byte, key uint64) error {
	solTx := &rpc.SolanaTransaction{}
	err := json.Unmarshal(dataByte, &solTx)
	if err != nil {
		return err
	}

	g.Log().File("solana-consumer.{Y-m-d}.log").Printf("消费 %v", solTx.Signature)

	ctx := gctx.New()

	// 检查如果hash已存在，直接跳过
	exists, err := sservice.Recharge.GetInfoByHash(ctx, solTx.Signature)
	if err != nil {
		return err
	}
	if exists != nil {
		g.Log().File("solana-consumer.{Y-m-d}.log").Printf("交易已存在，跳过 %v", solTx.Signature)
		return nil
	}

	// 检查来源地址，如果是手续费地址或提币地址则跳过
	solFeeAddress, _ := service.SysConfig.GetConfigByKey("sys.solanaFeeAddress")
	if solFeeAddress != nil && solTx.FromAddress == solFeeAddress.ConfigValue {
		return nil
	}
	solWithdrawAddress, _ := service.SysConfig.GetConfigByKey("sys.solanaWithdrawAddress")
	if solWithdrawAddress != nil && solTx.FromAddress == solWithdrawAddress.ConfigValue {
		return nil
	}

	var (
		coinToken       string
		contractAddress string
		amount          string
		rechargeType    int8
		status          int8
	)

	if solTx.IsToken {
		// SPL Token转账
		coinAddress, err := sservice.Currency.GetSolanaCoinAddress(ctx)
		if err != nil {
			return err
		}
		// 检查Mint是否在我们的币种列表中
		currency, ok := coinAddress[solTx.Mint]
		if !ok {
			g.Log().File("solana-consumer.{Y-m-d}.log").Printf("未知的SPL Token Mint: %v, 跳过", solTx.Mint)
			return nil
		}
		coinToken = currency.Name
		contractAddress = solTx.Mint
		// 金额已经在Helius回调中计算好了（带精度）
		amount = solTx.Amount
		rechargeType = 1
		status = 1
	} else {
		// SOL原生转账
		coinAddress, err := sservice.Currency.GetSolanaCoinAddress(ctx)
		if err != nil {
			return err
		}
		currency, ok := coinAddress[rpc.SOLNativeMint]
		if !ok {
			g.Log().File("solana-consumer.{Y-m-d}.log").Printf("SOL币种未配置，跳过")
			return nil
		}
		coinToken = currency.Name
		contractAddress = rpc.SOLNativeMint
		// Helius回调的amount是lamports，需要转换
		lamports, _ := decimal.NewFromString(solTx.Amount)
		amount = lamports.Div(decimal.NewFromInt(rpc.LamportsPerSOL)).String()
		rechargeType = 1
		status = 1
	}

	// 合约充值订单号写入备注
	remarks := ""
	if solTx.OrderId != "" {
		remarks = solTx.OrderId
	}

	// 写入充值记录
	data := dao.RechargeAddReq{
		MainChain:       "solana",
		BlockHash:       solTx.BlockHash,
		CoinToken:       coinToken,
		FromAddress:     solTx.FromAddress,
		ToAddress:       solTx.ToAddress,
		Amount:          amount,
		ContractAddress: contractAddress,
		Hash:            solTx.Signature,
		BlockHeight:     gconv.String(solTx.Slot),
		Status:          status,
		RechargeType:    rechargeType,
		Remarks:         remarks,
	}
	err = sservice.Recharge.Add(ctx, &data)
	if err != nil {
		g.Log().File("solana-consumer.{Y-m-d}.log").Printf("插入交易失败 %v", err)
		return err
	}

	// 发送回调通知
	solanaSendNotify(&data)
	return nil
}

// solanaSendNotify 发送充值回调通知
func solanaSendNotify(recharge *dao.RechargeAddReq) {
	callbackUrl, _ := service.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
	if callbackUrl == nil || callbackUrl.ConfigValue == "" {
		g.Log().File("callback.{Y-m-d}.log").Printf("未配置回调地址,不发送请求 %v", recharge)
		return
	}
	data := url.Values{
		"main_chain":       {recharge.MainChain},
		"block_hash":       {recharge.BlockHash},
		"recharge_type":    {gconv.String(recharge.RechargeType)},
		"from_address":     {recharge.FromAddress},
		"to_address":       {recharge.ToAddress},
		"coin_token":       {recharge.CoinToken},
		"contract_address": {recharge.ContractAddress},
		"amount":           {recharge.Amount},
		"hash":             {recharge.Hash},
		"imputation_hash":  {""},
		"remarks":          {recharge.Remarks},
		"status":           {gconv.String(recharge.Status)},
	}
	resp, err := g.Client().PostForm(callbackUrl.ConfigValue, data)
	if err != nil {
		g.Log().File("callback.{Y-m-d}.log").Printf("发送充值回调请求失败: %v", err)
		return
	}
	defer resp.Body.Close()
	g.Log().File("callback.{Y-m-d}.log").Printf("发送充值回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】",
		callbackUrl.ConfigValue, data.Encode(), resp.StatusCode)
}
