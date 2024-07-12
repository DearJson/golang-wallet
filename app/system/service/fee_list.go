// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/service/fee_list.go
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

type feeList struct {
}

var FeeList = new(feeList)

// GetList 获取任务列表
func (s *feeList) GetList(req *dao.FeeListSearchReq) (total, page int, list []*model.FeeList, err error) {
	m := dao.FeeList.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.FeeList.Columns.MainChain+" = ?", req.MainChain)
	}
	if req.Address != "" {
		m = m.Where(dao.FeeList.Columns.Address+" = ?", req.Address)
	}
	if req.Hash != "" {
		m = m.Where(dao.FeeList.Columns.Hash+" = ?", req.Hash)
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
func (s *feeList) GetInfoById(ctx context.Context, id uint64) (info *model.FeeList, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.FeeList.Ctx(ctx).Where(dao.FeeList.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *feeList) Add(ctx context.Context, req *dao.FeeListAddReq) (err error) {
	_, err = dao.FeeList.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *feeList) Edit(ctx context.Context, req *dao.FeeListEditReq) error {
	_, err := dao.FeeList.Ctx(ctx).FieldsEx(dao.FeeList.Columns.Id, dao.FeeList.Columns.CreatedAt).Where(dao.FeeList.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *feeList) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.FeeList.Ctx(ctx).Delete(dao.FeeList.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// ChangeStatus 修改状态
func (s *feeList) ChangeStatus(ctx context.Context, req *dao.FeeListStatusReq) error {
	_, err := dao.FeeList.Ctx(ctx).WherePri(req.Id).Update(g.Map{
		dao.FeeList.Columns.Status: req.Status,
	})
	return err
}
