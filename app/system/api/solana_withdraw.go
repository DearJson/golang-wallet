package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"gfast/app/system/service"
	"gfast/library"
	"net/url"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type solanaWithdraw struct {
	SystemBase
}

var SolanaWithdraw = new(solanaWithdraw)

// List 列表
func (c *solanaWithdraw) List(r *ghttp.Request) {
	var req *dao.WithdrawSearchReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	req.Ctx = r.GetCtx()
	withdrawTotal, total, page, list, err := service.Withdraw.GetList(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
		"coinTotal":   withdrawTotal,
	}
	c.SusJsonExit(r, result)
}

// Get 获取
func (c *solanaWithdraw) Get(r *ghttp.Request) {
	id := r.GetUint64("id")
	if id == 0 {
		c.FailJsonExit(r, "参数错误")
	}
	info, err := service.Withdraw.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// AuditSuccess 审核通过
func (c *solanaWithdraw) AuditSuccess(r *ghttp.Request) {
	ids := r.GetArray("ids")
	if len(ids) <= 0 {
		c.FailJsonExit(r, "参数错误")
	}
	array, _ := dao.Withdraw.Ctx(r.GetCtx()).WhereIn(dao.Withdraw.Columns.Id, ids).WhereIn(dao.Withdraw.Columns.Status, [2]int{1, 4}).Array(dao.Withdraw.Columns.Id)
	if len(array) == 0 {
		c.SusJsonExit(r)
	}
	var list []*model.Withdraw
	g.Model("withdraw").WhereIn("id", array).Scan(&list)
	for _, value := range list {
		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, 2, value.Nonce1)
		g.Model("withdraw").Data(g.Map{"status": 2, "hashKey": hashKey}).Where("id", value.Id).Update()
	}
	c.SusJsonExit(r)
}

// AuditFail 审核拒绝
func (c *solanaWithdraw) AuditFail(r *ghttp.Request) {
	ids := r.GetArray("ids")
	if len(ids) <= 0 {
		c.FailJsonExit(r, "参数错误")
	}
	array, _ := dao.Withdraw.Ctx(r.GetCtx()).WhereIn(dao.Withdraw.Columns.Id, ids).WhereIn(dao.Withdraw.Columns.Status, [2]int{1, 4}).Array(dao.Withdraw.Columns.Id)
	if len(array) == 0 {
		c.SusJsonExit(r)
	}
	var list []*model.Withdraw
	g.Model("withdraw").WhereIn("id", array).Scan(&list)
	for _, value := range list {
		hashKey := library.Md5Data(value.Address, value.ContractAddress, value.Amount, 6, value.Nonce1)
		g.Model("withdraw").Data(g.Map{"status": 6, "hashKey": hashKey}).Where("id", value.Id).Update()

		if value.NotifyUrl != "" {
			data := url.Values{
				"main_chain":       {value.MainChain},
				"coin_token":       {value.CoinToken},
				"address":          {value.Address},
				"contract_address": {gconv.String(value.ContractAddress)},
				"amount":           {gconv.String(value.Amount)},
				"status":           {gconv.String(6)},
				"remarks":          {value.Remarks},
			}
			resp, _ := g.Client().PostForm(value.NotifyUrl, data)
			if resp != nil {
				g.Log().File("callback.{Y-m-d}.log").Printf("发送Solana提现拒绝回调 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", value.NotifyUrl, data.Encode(), resp.StatusCode)
				resp.Body.Close()
			}
		}
	}
	c.SusJsonExit(r)
}

// Withdraw 手动提现
func (c *solanaWithdraw) Withdraw(r *ghttp.Request) {
	id := r.GetInt64("id")
	hash := r.GetString("hash")
	withdrawAddress := r.GetString("address")
	info, err := service.Withdraw.GetInfoById(r.GetCtx(), uint64(id))
	if err != nil || (info.Status == 5 || info.Status == 6) {
		c.FailJsonExit(r, "参数错误")
	}
	hashKey := library.Md5Data(info.Address, info.ContractAddress, info.Amount, 3, info.Nonce1)
	g.Model("withdraw").Data(g.Map{"status": 3, "hashKey": hashKey, "hash": hash, "withdraw_address": withdrawAddress}).Where("id", info.Id).Update()
	c.SusJsonExit(r)
}

// Callback 发起回调
func (c *solanaWithdraw) Callback(r *ghttp.Request) {
	ids := r.GetArray("ids")
	if len(ids) <= 0 {
		c.FailJsonExit(r, "参数错误")
	}
	var list []*model.Withdraw
	dao.Withdraw.Ctx(r.GetCtx()).WhereIn(dao.Withdraw.Columns.Id, ids).Scan(&list)
	if len(list) == 0 {
		c.SusJsonExit(r)
	}
	for _, value := range list {
		if value.NotifyUrl != "" {
			data := url.Values{
				"address":          {value.Address},
				"contract_address": {gconv.String(value.ContractAddress)},
				"amount":           {gconv.String(value.Amount)},
				"split_address":    {value.SplitAddress},
				"split_amount":     {value.SplitAmount},
				"status":           {gconv.String(value.Status)},
				"remarks":          {value.Remarks},
				"hash":             {value.Hash},
			}
			resp, _ := g.Client().PostForm(value.NotifyUrl, data)
			if resp != nil {
				g.Log().File("callback.{Y-m-d}.log").Printf("发送Solana提现回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", value.NotifyUrl, data.Encode(), resp.StatusCode)
				resp.Body.Close()
			}
		}
	}
	c.SusJsonExit(r)
}
