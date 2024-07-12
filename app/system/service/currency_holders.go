// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-27 04:23:13
// 生成路径: gfast/app/system/service/currency_holders.go
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

type currencyHolders struct {
}

var CurrencyHolders = new(currencyHolders)

// GetList 获取任务列表
func (s *currencyHolders) GetList(req *dao.CurrencyHoldersSearchReq) (total, page int, list []*model.CurrencyHolders, err error) {
	m := dao.CurrencyHolders.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.CurrencyHolders.Columns.MainChain+" = ?", req.MainChain)
	}
	if req.ContractAddress != "" {
		m = m.Where(dao.CurrencyHolders.Columns.ContractAddress+" = ?", req.ContractAddress)
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
func (s *currencyHolders) GetInfoById(ctx context.Context, id int) (info *model.CurrencyHolders, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.CurrencyHolders.Ctx(ctx).Where(dao.CurrencyHolders.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *currencyHolders) Add(ctx context.Context, req *dao.CurrencyHoldersAddReq) (err error) {
	_, err = dao.CurrencyHolders.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *currencyHolders) Edit(ctx context.Context, req *dao.CurrencyHoldersEditReq) error {
	_, err := dao.CurrencyHolders.Ctx(ctx).FieldsEx(dao.CurrencyHolders.Columns.Id, dao.CurrencyHolders.Columns.CreatedAt).Where(dao.CurrencyHolders.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *currencyHolders) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.CurrencyHolders.Ctx(ctx).Delete(dao.CurrencyHolders.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}
