package router

import (
	"gfast/app/web/api"
	"gfast/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/v1", func(group *ghttp.RouterGroup) {
		//context拦截器dd
		group.Middleware(middleware.CheckIp)

		group.Group("/bnb", func(group *ghttp.RouterGroup) {
			//创建币安地址
			group.POST("/generateAddress", api.Bsc.GenerateAddress)
			//币安提现代币
			group.POST("/withdraw", api.Bsc.Withdraw)
			//币安提现NFT
			group.POST("/withdrawNft", api.Bsc.WithdrawNft)
			//修改提现状态
			group.POST("/updateWithdrawStatus", api.Bsc.UpdateWithdrawStatus)

			//查询合约信息
			group.POST("/tokenInfo", api.Bsc.TokenInfo)
			//查询LP token
			group.POST("/lpInfo", api.Bsc.LpInfo)
			//查询LP3 token
			group.POST("/lp3Info", api.Bsc.Lp3Info)
			////查询余额
			group.POST("/balanceOf", api.Bsc.BalanceOf)
			////写入要监控的地址
			group.POST("/setAddress", api.Bsc.SetAddress)
			//单独归集某个账户的
			group.POST("/mergeAmount", api.Bsc.MergeAmount)
			//按照现有要求执行一下某个块
			group.POST("/resetBlock", api.Bsc.ResetBlock)
			//按照现有要求执行一下某个hash
			group.POST("/resetHash", api.Bsc.ResetHash)
			//新增币种
			group.POST("/insertCurrency", api.Bsc.InsertCurrency)
			//查询v2代币价格
			group.POST("/getCoinPrice", api.Bsc.GetCoinPrice)
			//查询v3代币价格
			//group.POST("/getCoinPrice", api.Bsc.V3GetCoinPrice)
			//查询用户nft持有情况
			group.POST("/getUserNft", api.Bsc.GetUserNft)
			//查询某笔交易详情
			group.POST("/transferDetail", api.Bsc.GetTransferDetail)
			//查询充值合约自动买币自动添加流动性LP数量
			group.POST("/getAutoAddLiquidityDetail", api.Bsc.GetAutoAddLiquidityDetail)
			//查询充值合约自动买币的数量
			group.POST("/getAutoTradeDetail", api.Bsc.GetAutoTradeDetail)
			//签名校验工具
			group.POST("/signVerify", api.Bsc.SignVerify)
			//根据hash查询交易详情
			group.POST("/getDetailByHash", api.Bsc.ResetHash)
		})

		group.Group("/eth", func(group *ghttp.RouterGroup) {
			//创建以太坊地址
			group.POST("/generateAddress", api.Eth.GenerateAddress)
			//以太坊提现代币
			group.POST("/withdraw", api.Eth.Withdraw)
			//按照现有要求执行一下某个hash
			group.POST("/resetHash", api.Eth.ResetHash)
		})

		group.Group("/tron", func(group *ghttp.RouterGroup) {
			//创建波场地址
			group.POST("/generateAddress", api.Tron.TronGenerateAddress)
			//波场提现代币
			group.POST("/withdraw", api.Tron.TronWithdraw)
			//波场提现NFT
			group.POST("/withdrawNft", api.Tron.TronWithdrawNft)
			//查询合约信息
			group.POST("/tokenInfo", api.Tron.TronTokenInfo)
			//查询LP token
			group.POST("/lpInfo", api.Tron.TronLpInfo)
			////查询余额
			group.POST("/balanceOf", api.Tron.TronBalanceOf)
			////写入要监控的地址
			group.POST("/setAddress", api.Tron.TronSetAddress)

			//按照现有要求执行一下某个块
			group.POST("/resetBlock", api.Tron.ResetBlock)
		})

		group.Group("/heco", func(group *ghttp.RouterGroup) {
			//创建波场地址
			group.POST("/generateAddress", api.Heco.GenerateAddress)
			//波场提现代币
			group.POST("/withdraw", api.Heco.Withdraw)
			////写入要监控的地址
			group.POST("/setAddress", api.Heco.SetAddress)

		})

		group.Group("/wemix", func(group *ghttp.RouterGroup) {
			//创建WEMIX地址
			group.POST("/generateAddress", api.Wemix.GenerateAddress)
			//WEMIX提现代币
			group.POST("/withdraw", api.Wemix.Withdraw)
		})

		group.Group("/nac", func(group *ghttp.RouterGroup) {
			//创建NAC地址
			group.POST("/generateAddress", api.Nac.GenerateAddress)
			//NAC提现代币
			group.POST("/withdraw", api.Nac.Withdraw)
			//NAC查询余额
			group.POST("/getBalance", api.Nac.GetBalance)
			//按照现有要求执行一下某个块
			group.POST("/resetBlock", api.Nac.ResetBlock)
		})

	})

}
