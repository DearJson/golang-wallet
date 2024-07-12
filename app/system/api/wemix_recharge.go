// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/api/recharge.go
// 生成人：gfast
// ==========================================================================

package api

import (
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"io/ioutil"
	"net/url"
)

type wemixRecharge struct {
	SystemBase
}

var WemixRecharge = new(wemixRecharge)

// List 列表
func (c *wemixRecharge) List(r *ghttp.Request) {
	var req *dao.RechargeSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "wemix"
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
func (c *wemixRecharge) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.Recharge.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Callback 发起回调
func (c *wemixRecharge) Callback(r *ghttp.Request) {
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
	data := url.Values{}
	for _, value := range list {
		data = url.Values{
			"main_chain":        {"heco"},
			"block_hash":        {value.BlockHash},
			"recharge_type":     {gconv.String(value.RechargeType)},
			"from_address":      {value.FromAddress},
			"to_address":        {value.ToAddress},
			"coin_token":        {value.CoinToken},
			"coin_token1":       {value.CoinToken1},
			"contract_address":  {value.ContractAddress},
			"contract_address1": {value.ContractAddress1},
			"amount":            {gconv.String(value.Amount)},
			"amount1":           {gconv.String(value.Amount1)},
			"hash":              {value.Hash},
			"imputation_hash":   {""},
			"remarks":           {value.Remarks},
			"status":            {gconv.String(value.Status)},
			"token_id":          {value.TokenId},
		}
		resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, data)
		resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			g.Log().File("callback.{Y-m-d}.log").Printf("发送充值回调请求失败 %v 请求域名:【%v】 \n", body, callbackUrl.ConfigValue)
			return
		}
		g.Model("recharge").Data(g.Map{"call_number": value.CallNumber + 1}).Where(dao.Recharge.Columns.Id, value.Id).Update()

	}
	c.SusJsonExit(r)
}
