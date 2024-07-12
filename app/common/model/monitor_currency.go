package model

import "github.com/gogf/gf/os/gtime"

// MonitorCurrency 监控币种
type MonitorCurrency struct {
	Id              int64       `orm:"id,primary" json:"id"`                    //
	MainChain       string      `orm:"main_chain" json:"mainChain"`             // 主链
	Name            string      `orm:"name" json:"name"`                        // 币种名称
	ContractAddress string      `orm:"contract_address" json:"contractAddress"` // 合约地址
	Decimals        int         `orm:"decimals" json:"decimals"`                // 精度
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`             //
	UpdatedAt       *gtime.Time `orm:"updated_at" json:"updatedAt"`             //
}
