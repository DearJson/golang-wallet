// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/service/height.go
// 生成人：gfast
// ==========================================================================

package service

import (
	"context"
	comModel "gfast/app/common/model"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type height struct {
}

var Height = new(height)

// GetList 获取任务列表
func (s *height) GetList(req *dao.HeightSearchReq) (total, page int, list []*model.Height, err error) {
	m := dao.Height.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.Height.Columns.MainChain+" = ?", req.MainChain)
	}
	total, err = m.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	page = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = comModel.PageSize
	}
	order := "id desc"
	if req.OrderBy != "" {
		order = req.OrderBy
	}
	err = m.Page(page, req.PageSize).Order(order).Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

// GetInfoById 通过id获取
func (s *height) GetInfoById(ctx context.Context, id uint64) (info *model.Height, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Height.Ctx(ctx).Where(dao.Height.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}
