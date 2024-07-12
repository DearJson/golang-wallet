package dao

import "github.com/ethereum/go-ethereum/common"

// LpInfoReq LP合约地址
type LpInfoReq struct {
	ContractAddress string `p:"contract_address" v:"required#合约地址不能为空"`
}

type BalanceOfReq struct {
	ContractAddress string `p:"contract_address" v:"required#合约地址不能为空"`
	Address         string `p:"address" v:"required#地址不能为空"`
}

type SetAddressReq struct {
	UserId  string `p:"user_id" v:"required#用户标识不能为空"`
	Address string `p:"address" v:"required#地址不能为空"`
}

type GetCoinPriceReq struct {
	RouteAddress     string           `p:"route_address" v:"required#路由地址不能为空"`
	Path             []common.Address `p:"path" v:"required#转换路径不能为空"`
	AmountInDecimals int              `p:"amount_in_decimals" v:"required#币种精度不能为空"`
}
type V3GetCoinPriceReq struct {
	Token1 string `p:"token1" v:"required#token1不能为空"`
	Token2 string `p:"token2" v:"required#token2不能为空"`
	Fee    int64  `p:"fee"  v:"required#手续费不能为空"`
}

type GetUserNftReq struct {
	ContractAddress string `p:"contract_address" v:"required#NFT合约地址不能为空"`
	Address         string `p:"address" v:"required#地址不能为空"`
}
