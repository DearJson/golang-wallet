package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type monitorCurrency struct {
}

var MonitorCurrency = new(monitorCurrency)

func (s *monitorCurrency) GetAddress(mainChain string) (list []string, err error) {
	//从缓存获取
	key := "MonitorCurrencyList" + mainChain
	iList := Cache.New().Get(key)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		return
	}
	stsd, _ := g.Model("monitor_currency").Where("main_chain", "bsc").FindArray("contract_address")
	list = gconv.SliceStr(stsd)
	if err != nil {
		g.Log().Error(err)
	}
	//缓存
	Cache.New().Set(key, list, 0)
	return
}
