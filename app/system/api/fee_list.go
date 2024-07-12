// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/api/fee_list.go
// 生成人：gfast
// ==========================================================================

package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type feeList struct {
	SystemBase
}

var FeeList = new(feeList)

// List 列表
func (c *feeList) List(r *ghttp.Request) {
	var req *dao.FeeListSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
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

// Add 添加
func (c *feeList) Add(r *ghttp.Request) {
	var req *dao.FeeListAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.FeeList.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *feeList) Get(r *ghttp.Request) {
	id := r.GetUint64("id")
	if id == 0 {
		c.FailJsonExit(r, "参数错误")
	}
	info, err := service.FeeList.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *feeList) Edit(r *ghttp.Request) {
	var req *dao.FeeListEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.FeeList.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *feeList) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.FeeList.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

// ChangeStatus 修改状态
func (c *feeList) ChangeStatus(r *ghttp.Request) {
	var req *dao.FeeListStatusReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if err := service.FeeList.ChangeStatus(r.GetCtx(), req); err != nil {
		c.FailJsonExit(r, err.Error())
	} else {
		c.SusJsonExit(r, "状态设置成功")
	}
}
