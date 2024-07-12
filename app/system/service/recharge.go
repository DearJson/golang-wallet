// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/service/recharge.go
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

type recharge struct {
}

var Recharge = new(recharge)

// GetList 获取任务列表
func (s *recharge) GetList(req *dao.RechargeSearchReq) (coin1Total float64, coin2Total float64, total, page int, list []*model.Recharge, err error) {
	m := dao.Recharge.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.Recharge.Columns.MainChain+" = ?", req.MainChain)
	}
	if req.BlockHash != "" {
		m = m.Where(dao.Recharge.Columns.BlockHash+" = ?", req.BlockHash)
	}
	if req.FromAddress != "" {
		m = m.Where(dao.Recharge.Columns.FromAddress+" = ?", req.FromAddress)
	}
	if req.ToAddress != "" {
		m = m.Where(dao.Recharge.Columns.ToAddress+" = ?", req.ToAddress)
	}
	if req.ContractAddress != "" {
		m = m.Where(dao.Recharge.Columns.ContractAddress+" = ?", req.ContractAddress)
	}
	if req.ContractAddress1 != "" {
		m = m.Where(dao.Recharge.Columns.ContractAddress1+" = ?", req.ContractAddress1)
	}
	if req.Hash != "" {
		m = m.Where(dao.Recharge.Columns.Hash+" = ?", req.Hash)
	}
	if req.BlockHeight != "" {
		m = m.Where(dao.Recharge.Columns.BlockHeight+" = ?", req.BlockHeight)
	}
	if req.Status != "" {
		m = m.Where(dao.Recharge.Columns.Status+" = ?", req.Status)
	}
	if req.ImputationHash != "" {
		m = m.Where(dao.Recharge.Columns.ImputationHash+" = ?", req.ImputationHash)
	}
	if req.TimeRange != "" {
		timeRange := library.Explode(",", req.TimeRange)
		m = m.Where(dao.Recharge.Columns.CreatedAt+" >= ?", timeRange[0]).Where(dao.Recharge.Columns.CreatedAt+" <= ?", timeRange[1])
	}

	total, err = m.Count()
	coin1Total, err = m.Sum("amount")
	coin2Total, err = m.Sum("amount1")
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
func (s *recharge) GetInfoById(ctx context.Context, id int64) (info *model.Recharge, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Recharge.Ctx(ctx).Where(dao.Recharge.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// GetInfoByHash 通过Hash查找充值记录
func (s *recharge) GetInfoByHash(ctx context.Context, hash string) (info *model.Recharge, err error) {
	if hash == "" {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Recharge.Ctx(ctx).Where(dao.Recharge.Columns.Hash, hash).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		return nil, nil
	}
	return
}

// Add 添加
func (s *recharge) Add(ctx context.Context, req *dao.RechargeAddReq) (err error) {
	_, err = dao.Recharge.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *recharge) Edit(ctx context.Context, req *dao.RechargeEditReq) error {
	_, err := dao.Recharge.Ctx(ctx).FieldsEx(dao.Recharge.Columns.Id, dao.Recharge.Columns.CreatedAt).Where(dao.Recharge.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *recharge) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.Recharge.Ctx(ctx).Delete(dao.Recharge.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// ChangeStatus 修改状态
func (s *recharge) ChangeStatus(ctx context.Context, req *dao.RechargeStatusReq) error {
	_, err := dao.Recharge.Ctx(ctx).WherePri(req.Id).Update(g.Map{
		dao.Recharge.Columns.Status: req.Status,
	})
	return err
}
