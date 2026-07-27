package api

import (
	"fmt"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"gfast/hdwallet"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"github.com/shopspring/decimal"
)

type solanaCurrency struct {
	SystemBase
}

var SolanaCurrency = new(solanaCurrency)

func validateSolanaWithdrawSplit(enabled int8, address string, amount float64, bps, decimals int) error {
	if enabled != 0 && enabled != 1 {
		return fmt.Errorf("提现分账开关只能为0或1")
	}
	if enabled == 0 {
		return nil
	}
	addressBytes, err := hdwallet.SolanaBase58Decode(address)
	if err != nil || len(addressBytes) != 32 {
		return fmt.Errorf("提现分账地址不是有效的Solana地址")
	}
	if bps < 0 || bps >= 10000 {
		return fmt.Errorf("提现分账比例必须在1到9999基点之间")
	}
	splitAmount := decimal.NewFromFloat(amount)
	if splitAmount.IsNegative() {
		return fmt.Errorf("提现固定分账数量不能小于0")
	}
	hasFixedAmount := splitAmount.IsPositive()
	hasRate := bps > 0
	if hasFixedAmount == hasRate {
		return fmt.Errorf("提现固定分账数量和分账比例必须二选一")
	}
	if decimals < 0 || decimals > 255 {
		return fmt.Errorf("币种精度无效")
	}
	if hasFixedAmount {
		scaledAmount := splitAmount.Mul(decimal.New(1, int32(decimals)))
		if !scaledAmount.Equal(scaledAmount.Truncate(0)) {
			return fmt.Errorf("提现分账数量不能超过币种精度")
		}
	}
	return nil
}

// List 列表
func (c *solanaCurrency) List(r *ghttp.Request) {
	var req *dao.CurrencySearchReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	req.Ctx = r.GetCtx()
	total, page, list, err := service.Currency.GetList(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	c.SusJsonExit(r, result)
}

// Add 添加
func (c *solanaCurrency) Add(r *ghttp.Request) {
	var req *dao.CurrencyAddReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	if err := validateSolanaWithdrawSplit(req.WithdrawSplitEnabled, req.WithdrawSplitAddress, req.WithdrawSplitAmount, req.WithdrawSplitBps, req.Decimals); err != nil {
		c.FailJsonExit(r, err.Error())
	}
	err := service.Currency.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *solanaCurrency) Get(r *ghttp.Request) {
	id := r.GetUint64("id")
	if id == 0 {
		c.FailJsonExit(r, "参数错误")
	}
	info, err := service.Currency.GetInfoById(r.GetCtx(), int64(id))
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *solanaCurrency) Edit(r *ghttp.Request) {
	var req *dao.CurrencyEditReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	if err := validateSolanaWithdrawSplit(req.WithdrawSplitEnabled, req.WithdrawSplitAddress, req.WithdrawSplitAmount, req.WithdrawSplitBps, req.Decimals); err != nil {
		c.FailJsonExit(r, err.Error())
	}
	err := service.Currency.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *solanaCurrency) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.Currency.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}
