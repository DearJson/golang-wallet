// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/api/currency.go
// 生成人：gfast
// ==========================================================================

package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"strings"
)

type bscCurrency struct {
	SystemBase
}

var BscCurrency = new(bscCurrency)

// List 列表
func (c *bscCurrency) List(r *ghttp.Request) {
	var req *dao.CurrencySearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "bsc"
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
func (c *bscCurrency) Add(r *ghttp.Request) {
	var req *dao.CurrencyAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.ContractAddress = strings.ToLower(req.ContractAddress)
	req.MainChain = "bsc"
	err := service.Currency.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *bscCurrency) Get(r *ghttp.Request) {
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
func (c *bscCurrency) Edit(r *ghttp.Request) {
	var req *dao.CurrencyEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.ContractAddress = strings.ToLower(req.ContractAddress)
	req.MainChain = "bsc"
	err := service.Currency.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *bscCurrency) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.Currency.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}
