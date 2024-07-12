// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/service/currency_holders_detail.go
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

type currencyHoldersDetail struct {
}

var CurrencyHoldersDetail = new(currencyHoldersDetail)

// GetList 获取任务列表
func (s *currencyHoldersDetail) GetList(req *dao.CurrencyHoldersDetailSearchReq) (total, page int, list []*model.CurrencyHoldersDetail, err error) {
	m := dao.CurrencyHoldersDetail.Ctx(req.Ctx)
	if req.HoldersId != "" {
		m = m.Where(dao.CurrencyHoldersDetail.Columns.HoldersId+" = ?", req.HoldersId)
	}
	if req.Address != "" {
		m = m.Where(dao.CurrencyHoldersDetail.Columns.Address+" = ?", req.Address)
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
	order := "id asc"
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
func (s *currencyHoldersDetail) GetInfoById(ctx context.Context, id int) (info *model.CurrencyHoldersDetail, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.CurrencyHoldersDetail.Ctx(ctx).Where(dao.CurrencyHoldersDetail.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *currencyHoldersDetail) Add(ctx context.Context, req *dao.CurrencyHoldersDetailAddReq) (err error) {
	_, err = dao.CurrencyHoldersDetail.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *currencyHoldersDetail) Edit(ctx context.Context, req *dao.CurrencyHoldersDetailEditReq) error {
	_, err := dao.CurrencyHoldersDetail.Ctx(ctx).FieldsEx(dao.CurrencyHoldersDetail.Columns.Id, dao.CurrencyHoldersDetail.Columns.CreatedAt).Where(dao.CurrencyHoldersDetail.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *currencyHoldersDetail) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.CurrencyHoldersDetail.Ctx(ctx).Delete(dao.CurrencyHoldersDetail.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}
