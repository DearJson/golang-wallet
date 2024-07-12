package task

import (
	"context"
	"gfast/abi/erc20"
	commonService "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
	"time"
)

// updateBalance 更新已授权的地址授权量以及余额量
func updateBalance() {
	var pageNum = 1
	var pageSize = 10
	var ctx context.Context
	cache := commonService.Cache.New()
	rpcUrl := gconv.String(cache.Get("bnb_rpc_url"))
	client, _ := ethclient.Dial(rpcUrl)
	for {
		var result []model.AuthorizeAddress
		dao.AuthorizeAddress.Ctx(ctx).Page(pageNum, pageSize).Order("id desc").Scan(&result)
		if len(result) <= 0 {
			break
		}
		for _, value := range result {
			erc20Contract, _ := erc20.NewErc20AbiCaller(common.HexToAddress(value.CoinAddress), client)
			transactOpts := &bind.CallOpts{
				Pending:     false,
				From:        common.Address{},
				BlockNumber: nil,
				Context:     nil,
			}
			balanceTemp, _ := erc20Contract.BalanceOf(transactOpts, common.HexToAddress(value.Address))
			authorizeNum, _ := erc20Contract.Allowance(transactOpts, common.HexToAddress(value.Address), common.HexToAddress(value.ContractAddress))

			g.Model("authorize_address").Data(g.Map{
				"balance":       decimal.NewFromBigInt(balanceTemp, 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(value.CoinDecimals))), int32(value.CoinDecimals)).Truncate(int32(value.CoinDecimals)).String(),
				"authorize_num": decimal.NewFromBigInt(authorizeNum, 0).DivRound(decimal.NewFromFloat(math.Pow(10, float64(value.CoinDecimals))), int32(value.CoinDecimals)).Truncate(int32(value.CoinDecimals)).String(),
				"balance_time":  gtime.New(time.Now()),
			}).Where("id", value.Id).Update()

		}

		pageNum = pageNum + 1
	}

}
