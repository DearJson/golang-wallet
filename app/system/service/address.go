// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 03:32:45
// 生成路径: gfast/app/system/service/address.go
// 生成人：gfast
// ==========================================================================

package service

import (
	"context"
	"gfast/app/common/global"
	comModel "gfast/app/common/model"
	"gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type address struct {
}

var Address = new(address)

// GetList 获取任务列表
func (s *address) GetList(req *dao.AddressSearchReq) (total, page int, list []*model.Address, err error) {
	m := dao.Address.Ctx(req.Ctx)
	if req.UserId != "" {
		m = m.Where(dao.Address.Columns.UserId+" = ?", req.UserId)
	}
	if req.Address != "" {
		m = m.Where(dao.Address.Columns.Address+" = ?", req.Address)
	}
	if req.MainChain != "" {
		m = m.Where(dao.Address.Columns.MainChain+" = ?", req.MainChain)
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
func (s *address) GetInfoById(ctx context.Context, id uint64) (info *model.Address, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Address.Ctx(ctx).Where(dao.Address.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// GetInfoByUserId 通过用户标识获取
func (s *address) GetInfoByUserId(ctx context.Context, userId string, mainChain string) (info *model.Address, err error) {
	if userId == "" {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Address.Ctx(ctx).Where("main_chain", mainChain).Where(dao.Address.Columns.UserId, userId).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	return
}

// Add 添加
func (s *address) Add(ctx context.Context, req *dao.AddressAddReq) (err error) {
	_, err = dao.Address.Ctx(ctx).Insert(req)
	cache := service.Cache.New()
	cache.Remove(global.BscUserAddressList)
	cache.Remove(global.TronUserAddressList)
	cache.Remove(global.HecoUserAddressList)
	cache.Remove(global.NacUserAddressList)
	cache.Remove(global.WemixUserAddressList)
	cache.Remove(global.EthUserAddressList)
	return
}

// Edit 修改
func (s *address) Edit(ctx context.Context, req *dao.AddressEditReq) error {
	_, err := dao.Address.Ctx(ctx).FieldsEx(dao.Address.Columns.Id, dao.Address.Columns.CreatedAt).Where(dao.Address.Columns.Id, req.Id).
		Update(req)
	cache := service.Cache.New()
	cache.Remove(global.BscUserAddressList)
	cache.Remove(global.TronUserAddressList)
	cache.Remove(global.HecoUserAddressList)
	cache.Remove(global.NacUserAddressList)
	cache.Remove(global.WemixUserAddressList)
	cache.Remove(global.EthUserAddressList)
	return err
}

// DeleteByIds 删除
func (s *address) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.Address.Ctx(ctx).Delete(dao.Address.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	cache := service.Cache.New()
	cache.Remove(global.BscUserAddressList)
	cache.Remove(global.TronUserAddressList)
	cache.Remove(global.HecoUserAddressList)
	cache.Remove(global.NacUserAddressList)
	cache.Remove(global.WemixUserAddressList)
	cache.Remove(global.EthUserAddressList)
	return
}

func (s *address) GetBnbAllAddress(ctx context.Context) (list []string, err error) {
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.BscUserAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := dao.Address.Ctx(ctx).Where(dao.Address.Columns.MainChain, "bsc").FindArray(dao.Address.Columns.Address)
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	cache.Set(global.BscUserAddressList, list, 0)
	return
}

func (s *address) GetEthAllAddress(ctx context.Context) (list []string, err error) {
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.EthUserAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := dao.Address.Ctx(ctx).Where(dao.Address.Columns.MainChain, "eth").FindArray(dao.Address.Columns.Address)
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	cache.Set(global.EthUserAddressList, list, 0)
	return
}

func (s *address) GetNacAllAddress(ctx context.Context) (list []string, err error) {
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.NacUserAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := dao.Address.Ctx(ctx).Where(dao.Address.Columns.MainChain, "nac").FindArray(dao.Address.Columns.Address)
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	cache.Set(global.NacUserAddressList, list, 0)
	return
}

func (s *address) GetHecoAllAddress(ctx context.Context) (list []string, err error) {
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.HecoUserAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := dao.Address.Ctx(ctx).Where(dao.Address.Columns.MainChain, "heco").FindArray(dao.Address.Columns.Address)
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	cache.Set(global.HecoUserAddressList, list, 0)
	return
}

func (s *address) GetWemixAllAddress(ctx context.Context) (list []string, err error) {
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.WemixUserAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := dao.Address.Ctx(ctx).Where(dao.Address.Columns.MainChain, "wemix").FindArray(dao.Address.Columns.Address)
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	cache.Set(global.WemixUserAddressList, list, 0)
	return
}

func (s *address) GetTronAllAddress(ctx context.Context) (list []string, err error) {
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.TronUserAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := dao.Address.Ctx(ctx).Where(dao.Address.Columns.MainChain, "tron").FindArray(dao.Address.Columns.Address)
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	cache.Set(global.TronUserAddressList, list, 0)
	return
}
