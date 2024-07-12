package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type liquidityCurrency struct {
}

var LiquidityCurrency = new(liquidityCurrency)

func (s *liquidityCurrency) GetAddress(mainChain string) (list []string, err error) {
	//从缓存获取
	key := "LiquidityCurrencyList" + mainChain
	iList := Cache.New().Get(key)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := g.Model("liquidity_currency").Where("main_chain", "bsc").FindArray("contract_address")
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	Cache.New().Set(key, list, 0)
	return
}
