// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/service/currency.go
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

type currency struct {
}

var Currency = new(currency)

// GetList 获取任务列表
func (s *currency) GetList(req *dao.CurrencySearchReq) (total, page int, list []*model.Currency, err error) {
	m := dao.Currency.Ctx(req.Ctx)
	if req.MainChain != "" {
		m = m.Where(dao.Currency.Columns.MainChain+" = ?", req.MainChain)
	}
	if req.Name != "" {
		m = m.Where(dao.Currency.Columns.Name+" like ?", "%"+req.Name+"%")
	}
	if req.ContractAddress != "" {
		m = m.Where(dao.Currency.Columns.ContractAddress+" = ?", req.ContractAddress)
	}
	if req.Decimals != "" {
		m = m.Where(dao.Currency.Columns.Decimals+" = ?", req.Decimals)
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
func (s *currency) GetInfoById(ctx context.Context, id int64) (info *model.Currency, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// GetInfoByContractAddress 通过合约地址获取
func (s *currency) GetInfoByContractAddress(ctx context.Context, contractAddress string, mainChain string) (info *model.Currency, err error) {
	if contractAddress == "" {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.ContractAddress, contractAddress).Where(dao.Currency.Columns.MainChain, mainChain).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	return
}

// Add 添加
func (s *currency) Add(ctx context.Context, req *dao.CurrencyAddReq) (err error) {
	_, err = dao.Currency.Ctx(ctx).Insert(req)
	cache := service.Cache.New()
	cache.Remove(global.BscCoinAddressList)
	cache.Remove(global.TronCoinAddressList)
	cache.Remove(global.HecoCoinAddressList)
	cache.Remove(global.WemixCoinAddressList)
	cache.Remove(global.EthCoinAddressList)
	return
}

// Edit 修改
func (s *currency) Edit(ctx context.Context, req *dao.CurrencyEditReq) error {
	_, err := dao.Currency.Ctx(ctx).FieldsEx(dao.Currency.Columns.Id, dao.Currency.Columns.CreatedAt).Where(dao.Currency.Columns.Id, req.Id).
		Update(req)
	cache := service.Cache.New()
	cache.Remove(global.BscCoinAddressList)
	cache.Remove(global.TronCoinAddressList)
	cache.Remove(global.HecoCoinAddressList)
	cache.Remove(global.WemixCoinAddressList)
	cache.Remove(global.EthCoinAddressList)
	return err
}

// DeleteByIds 删除
func (s *currency) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.Currency.Ctx(ctx).Delete(dao.Currency.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	cache := service.Cache.New()
	cache.Remove(global.BscCoinAddressList)
	cache.Remove(global.TronCoinAddressList)
	cache.Remove(global.HecoCoinAddressList)
	cache.Remove(global.WemixCoinAddressList)
	cache.Remove(global.EthCoinAddressList)
	return
}

func (s *currency) GetBnbCoinAddress(ctx context.Context) (list1 map[string]*model.Currency, err error) {
	list1 = make(map[string]*model.Currency)
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.BscCoinAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list1)
		return
	}
	var list []*model.Currency
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.MainChain, "bsc").Scan(&list)

	n := ""
	for _, value := range list {
		n = value.ContractAddress
		list1[n] = value
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	//缓存
	cache.Set(global.BscCoinAddressList, list1, 0)
	return
}

func (s *currency) GetEthCoinAddress(ctx context.Context) (list1 map[string]*model.Currency, err error) {
	list1 = make(map[string]*model.Currency)
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.EthCoinAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list1)
		return
	}
	var list []*model.Currency
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.MainChain, "eth").Scan(&list)

	n := ""
	for _, value := range list {
		n = value.ContractAddress
		list1[n] = value
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	//缓存
	cache.Set(global.EthCoinAddressList, list1, 0)
	return
}

func (s *currency) GetNacCoinAddress(ctx context.Context) (list1 map[string]*model.Currency, err error) {
	list1 = make(map[string]*model.Currency)
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.NacCoinAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list1)
		return
	}
	var list []*model.Currency
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.MainChain, "nac").Scan(&list)

	n := ""
	for _, value := range list {
		n = value.ContractAddress
		list1[n] = value
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	//缓存
	cache.Set(global.NacCoinAddressList, list1, 0)
	return
}

func (s *currency) GetHecoCoinAddress(ctx context.Context) (list1 map[string]*model.Currency, err error) {
	list1 = make(map[string]*model.Currency)
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.HecoCoinAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list1)
		return
	}
	var list []*model.Currency
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.MainChain, "heco").Scan(&list)

	n := ""
	for _, value := range list {
		n = value.ContractAddress
		list1[n] = value
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	//缓存
	cache.Set(global.HecoCoinAddressList, list1, 0)
	return
}

func (s *currency) GetWemixCoinAddress(ctx context.Context) (list1 map[string]*model.Currency, err error) {
	list1 = make(map[string]*model.Currency)
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.WemixCoinAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list1)
		return
	}
	var list []*model.Currency
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.MainChain, "wemix").Scan(&list)

	n := ""
	for _, value := range list {
		n = value.ContractAddress
		list1[n] = value
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	//缓存
	cache.Set(global.WemixCoinAddressList, list1, 0)
	return
}

func (s *currency) GetTronCoinAddress(ctx context.Context) (list1 map[string]*model.Currency, err error) {
	list1 = make(map[string]*model.Currency)
	cache := service.Cache.New()
	//从缓存获取
	iList := cache.Get(global.TronCoinAddressList)
	if iList != nil {
		err = gconv.Struct(iList, &list1)
		return
	}
	var list []*model.Currency
	err = dao.Currency.Ctx(ctx).Where(dao.Currency.Columns.MainChain, "tron").Scan(&list)

	n := ""
	for _, value := range list {
		n = value.ContractAddress
		list1[n] = value
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	//缓存
	cache.Set(global.TronCoinAddressList, list1, 0)
	return
}
