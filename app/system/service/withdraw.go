// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:34
// 生成路径: gfast/app/system/service/withdraw.go
// 生成人：gfast
// ==========================================================================

package service

import (
	"context"
	comModel "gfast/app/common/model"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"gfast/library"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type withdraw struct {
}

var Withdraw = new(withdraw)

// GetList 获取任务列表
func (s *withdraw) GetList(req *dao.WithdrawSearchReq) (withdrawTotal float64, total, page int, list []*model.Withdraw, err error) {
	m := dao.Withdraw.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.Withdraw.Columns.MainChain+" = ?", req.MainChain)
	}
	if req.CoinToken != "" {
		m = m.Where(dao.Withdraw.Columns.CoinToken+" = ?", req.CoinToken)
	}
	if req.Status != 0 {
		m = m.Where(dao.Withdraw.Columns.Status+" = ?", req.Status)
	}
	if req.Address != "" {
		m = m.Where(dao.Withdraw.Columns.Address+" = ?", req.Address)
	}
	if req.ContractAddress != "" {
		m = m.Where(dao.Withdraw.Columns.ContractAddress+" = ?", req.ContractAddress)
	}
	if req.Hash != "" {
		m = m.Where(dao.Withdraw.Columns.Hash+" = ?", req.Hash)
	}
	if req.Remarks != "" {
		m = m.Where(dao.Withdraw.Columns.Remarks+" = ?", req.Remarks)
	}
	if req.NotifyUrl != "" {
		m = m.Where(dao.Withdraw.Columns.NotifyUrl+" = ?", req.NotifyUrl)
	}
	if req.TimeRange != "" {
		timeRange := library.Explode(",", req.TimeRange)
		m = m.Where(dao.Recharge.Columns.CreatedAt+" >= ?", timeRange[0]).Where(dao.Recharge.Columns.CreatedAt+" <= ?", timeRange[1])
	}

	total, err = m.Count()
	withdrawTotal, err = m.Sum("amount")
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
	err = m.Page(page, req.PageSize).Order(order).WithAll().Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

// GetInfoById 通过id获取
func (s *withdraw) GetInfoById(ctx context.Context, id uint64) (info *model.Withdraw, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Withdraw.Ctx(ctx).Where(dao.Withdraw.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *withdraw) Add(ctx context.Context, req *dao.WithdrawAddReq) (err error) {
	_, err = dao.Withdraw.Ctx(ctx).Insert(req)
	return
}

// AddNft 添加
func (s *withdraw) AddNft(ctx context.Context, req *dao.WithdrawNftAddReq) (err error) {
	_, err = dao.Withdraw.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *withdraw) Edit(ctx context.Context, req *dao.WithdrawEditReq) error {
	_, err := dao.Withdraw.Ctx(ctx).FieldsEx(dao.Withdraw.Columns.Id, dao.Withdraw.Columns.CreatedAt).Where(dao.Withdraw.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *withdraw) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.Withdraw.Ctx(ctx).Delete(dao.Withdraw.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// ChangeStatus 修改状态
func (s *withdraw) ChangeStatus(ctx context.Context, req *dao.WithdrawStatusReq) error {
	_, err := dao.Withdraw.Ctx(ctx).WherePri(req.Id).Update(g.Map{
		dao.Withdraw.Columns.Status: req.Status,
	})
	return err
}
