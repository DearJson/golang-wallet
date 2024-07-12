// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-09-11 16:09:11
// 生成路径: gfast/app/system/service/authorize_address.go
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

type authorizeAddress struct {
}

var AuthorizeAddress = new(authorizeAddress)

// GetList 获取任务列表
func (s *authorizeAddress) GetList(req *dao.AuthorizeAddressSearchReq) (total, page int, list []*model.AuthorizeAddress, err error) {
	m := dao.AuthorizeAddress.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.MainChain+" = ?", req.MainChain)
	}
	if req.ContractAddress != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.ContractAddress+" = ?", req.ContractAddress)
	}
	if req.Address != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.Address+" = ?", req.Address)
	}
	if req.CoinName != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.CoinName+" like ?", "%"+req.CoinName+"%")
	}
	if req.CoinDecimals != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.CoinDecimals+" = ?", req.CoinDecimals)
	}
	if req.CoinAddress != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.CoinAddress+" = ?", req.CoinAddress)
	}
	if req.Num != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.Num+" = ?", req.Num)
	}
	if req.Balance != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.Balance+" = ?", req.Balance)
	}
	if req.AuthorizeNum != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.AuthorizeNum+" = ?", req.AuthorizeNum)
	}
	if req.AuthorizeHash != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.AuthorizeHash+" = ?", req.AuthorizeHash)
	}
	if req.AuthorizeBlock != "" {
		m = m.Where(dao.AuthorizeAddress.Columns.AuthorizeBlock+" = ?", req.AuthorizeBlock)
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
func (s *authorizeAddress) GetInfoById(ctx context.Context, id int64) (info *model.AuthorizeAddress, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.AuthorizeAddress.Ctx(ctx).Where(dao.AuthorizeAddress.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *authorizeAddress) Add(ctx context.Context, req *dao.AuthorizeAddressAddReq) (err error) {
	_, err = dao.AuthorizeAddress.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *authorizeAddress) Edit(ctx context.Context, req *dao.AuthorizeAddressEditReq) error {
	_, err := dao.AuthorizeAddress.Ctx(ctx).FieldsEx(dao.AuthorizeAddress.Columns.Id, dao.AuthorizeAddress.Columns.CreatedAt).Where(dao.AuthorizeAddress.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *authorizeAddress) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.AuthorizeAddress.Ctx(ctx).Delete(dao.AuthorizeAddress.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// GetInfoByHash 通过Hash查找充值记录
func (s *authorizeAddress) GetInfoByHash(ctx context.Context, authorizeHash string) (info *model.AuthorizeAddress, err error) {
	err = dao.AuthorizeAddress.Ctx(ctx).Where(dao.AuthorizeAddress.Columns.AuthorizeHash, authorizeHash).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		return nil, nil
	}
	return
}
