package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type solanaAddress struct {
	SystemBase
}

var SolanaAddress = new(solanaAddress)

// List 列表
func (c *solanaAddress) List(r *ghttp.Request) {
	var req *dao.AddressSearchReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.MainChain = "solana"
	req.Ctx = r.GetCtx()
	total, page, list, err := service.Address.GetList(req)
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
