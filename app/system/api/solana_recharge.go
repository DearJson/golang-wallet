package api

import (
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"gfast/app/system/service"
	"net/url"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type solanaRecharge struct {
	SystemBase
}

var SolanaRecharge = new(solanaRecharge)

// List 列表
func (c *solanaRecharge) List(r *ghttp.Request) {
	var req *dao.RechargeSearchReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	req.Ctx = r.GetCtx()
	coin1Total, coin2Total, total, page, list, err := service.Recharge.GetList(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
		"coin1Total":  coin1Total,
		"coin2Total":  coin2Total,
	}
	c.SusJsonExit(r, result)
}

// Get 获取
func (c *solanaRecharge) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.Recharge.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Callback 发起回调
func (c *solanaRecharge) Callback(r *ghttp.Request) {
	ids := r.GetArray("ids")
	if len(ids) <= 0 {
		c.FailJsonExit(r, "参数错误")
	}
	callbackUrl, _ := cservice.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
	if callbackUrl.ConfigValue == "" {
		c.FailJsonExit(r, "未配置回调地址")
	}

	var list []*model.Recharge
	dao.Recharge.Ctx(r.GetCtx()).WhereIn(dao.Recharge.Columns.Id, ids).Scan(&list)
	if len(list) == 0 {
		c.SusJsonExit(r)
	}
	for _, value := range list {
		cleanRemarks := strings.ReplaceAll(value.Remarks, "\x00", "")
		data := url.Values{
			"main_chain":       {value.MainChain},
			"block_hash":       {value.BlockHash},
			"recharge_type":    {gconv.String(value.RechargeType)},
			"from_address":     {value.FromAddress},
			"to_address":       {value.ToAddress},
			"coin_token":       {value.CoinToken},
			"contract_address": {value.ContractAddress},
			"amount":           {gconv.String(value.Amount)},
			"hash":             {value.Hash},
			"imputation_hash":  {value.ImputationHash},
			"remarks":          {cleanRemarks},
			"status":           {gconv.String(value.Status)},
		}
		resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, data)
		if resp != nil {
			g.Log().File("callback.{Y-m-d}.log").Printf("发送Solana充值回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", callbackUrl.ConfigValue, data.Encode(), resp.StatusCode)
			resp.Body.Close()
		}
	}
	c.SusJsonExit(r)
}
