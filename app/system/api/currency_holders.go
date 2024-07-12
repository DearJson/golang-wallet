// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/api/currency_holders.go
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

type currencyHolders struct {
	SystemBase
}

var CurrencyHolders = new(currencyHolders)

// List 列表
func (c *currencyHolders) List(r *ghttp.Request) {
	var req *dao.CurrencyHoldersSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.CurrencyHolders.GetList(req)
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
func (c *currencyHolders) Add(r *ghttp.Request) {
	var req *dao.CurrencyHoldersAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if req.MainChain == "bsc" {
		req.ContractAddress = strings.ToLower(req.ContractAddress)
	}
	err := service.CurrencyHolders.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *currencyHolders) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	info, err := service.CurrencyHolders.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *currencyHolders) Edit(r *ghttp.Request) {
	var req *dao.CurrencyHoldersEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if req.MainChain == "bsc" {
		req.ContractAddress = strings.ToLower(req.ContractAddress)
	}
	err := service.CurrencyHolders.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *currencyHolders) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.CurrencyHolders.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}
