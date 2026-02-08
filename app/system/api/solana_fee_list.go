package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type solanaFeeList struct {
	SystemBase
}

var SolanaFeeList = new(solanaFeeList)

// List 列表
func (c *solanaFeeList) List(r *ghttp.Request) {
	var req *dao.FeeListSearchReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	req.Ctx = r.GetCtx()
	total, page, list, err := service.FeeList.GetList(req)
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
