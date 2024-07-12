package task

import (
	"gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/library"
	"gfast/rpc"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"net/url"
)

//检查
func checkStatusTask() {
	if g.Config().GetBool("bsc.contract_recharge") || g.Config().GetBool("bsc.address_recharge") {
		bnbCheckStatus()
	}
	if g.Config().GetBool("tron.contract_recharge") || g.Config().GetBool("tron.address_recharge") {
		tronCheckStatus()
	}
	if g.Config().GetBool("heco.contract_recharge") || g.Config().GetBool("heco.address_recharge") {
		hecoCheckStatus()
	}
	if g.Config().GetBool("wemix.contract_recharge") || g.Config().GetBool("wemix.address_recharge") {
		wemixCheckStatus()
	}
	if g.Config().GetBool("nac.address_recharge") {
		nacCheckStatus()
	}
}

//bnbCheckStatus 检查币安相关任务
func bnbCheckStatus() {
	//查询所有待出金的任务
	var withdrawList []*model.Withdraw
	g.Model("withdraw").Where("main_chain", "bsc").Where("status", 3).Scan(&withdrawList)
	client := rpc.BscClient{}
	bscStruct := rpc.BscTransferDetail{}

	for _, withdraw := range withdrawList {
		data, err := client.Init().GetTransferStatus(withdraw.Hash)
		//检查，如果超过两分钟未找到hash，可以判定为失败了
		if err == nil && data == nil {
			if gtime.Timestamp()-withdraw.UpdatedAt.Timestamp() >= 120 {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		} else {
			err = mapstructure.Decode(data, &bscStruct)
			if err != nil || bscStruct.Status == "" {
				continue
			}
			if bscStruct.Status == "0x1" {
				g.Model("withdraw").Data(g.Map{"status": 5}).Where("id", withdraw.Id).Update()
				if withdraw.NotifyUrl != "" {
					data := url.Values{
						"address":          {withdraw.Address},
						"contract_address": {gconv.String(withdraw.ContractAddress)},
						"amount":           {gconv.String(withdraw.Amount)},
						"status":           {gconv.String(5)},
						"remarks":          {withdraw.Remarks},
						"hash":             {withdraw.Hash},
					}
					resp, _ := g.Client().PostForm(withdraw.NotifyUrl, data)
					g.Log().File("callback.{Y-m-d}.log").Printf("发送提现回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", withdraw.NotifyUrl, data.Encode(), resp.StatusCode)
					resp.Body.Close()
				}
			} else {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		}
	}

	var feeList []*model.FeeList
	g.Model("fee_list").Where("main_chain", "bsc").Where("status", 1).Scan(&feeList)
	for _, fee := range feeList {
		data, _ := client.Init().GetTransferStatus(fee.Hash)
		err := mapstructure.Decode(data, &bscStruct)
		if err != nil || bscStruct.Status == "" {
			continue
		}
		if bscStruct.Status == "0x1" {
			g.Model("fee_list").Data(g.Map{"status": 2}).Where("id", fee.Id).Update()
			g.Model("recharge").Data(g.Map{"status": 1}).Where("id", fee.RechargeId).Update()
		} else {
			g.Model("fee_list").Data(g.Map{"status": 3}).Where("id", fee.Id).Update()
		}
	}

	var rechargeList []*model.Recharge
	g.Model("recharge").Where("main_chain", "bsc").Where("status", 2).Scan(&rechargeList)

	for _, recharge := range rechargeList {
		data, _ := client.Init().GetTransferStatus(recharge.ImputationHash)
		err := mapstructure.Decode(data, &bscStruct)
		if err != nil || bscStruct.Status == "" {
			continue
		}
		if bscStruct.Status == "0x1" {
			g.Model("recharge").Data(g.Map{"status": 3}).Where("id", recharge.Id).Update()
			//归集成功发送通知
			recharge.Status = 3
			sendNotify(recharge)
		} else {
			g.Model("fee_list").Data(g.Map{"status": 4}).Where("id", recharge.Id).Update()
		}
	}
}

//hecoCheckStatus 检查火币相关任务
func hecoCheckStatus() {
	//查询所有待出金的任务
	var withdrawList []*model.Withdraw
	g.Model("withdraw").Where("main_chain", "heco").Where("status", 3).Scan(&withdrawList)
	client := rpc.HecoClient{}
	bscStruct := rpc.HecoTransferDetail{}

	for _, withdraw := range withdrawList {
		data, err := client.Init().GetTransferStatus(withdraw.Hash)
		//检查，如果超过两分钟未找到hash，可以判定为失败了
		if err == nil && data == nil {
			if gtime.Timestamp()-withdraw.UpdatedAt.Timestamp() >= 120 {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		} else {
			err = mapstructure.Decode(data, &bscStruct)
			if err != nil || bscStruct.Status == "" {
				continue
			}
			if bscStruct.Status == "0x1" {
				g.Model("withdraw").Data(g.Map{"status": 5}).Where("id", withdraw.Id).Update()
				if withdraw.NotifyUrl != "" {
					data := url.Values{
						"address":          {withdraw.Address},
						"contract_address": {gconv.String(withdraw.ContractAddress)},
						"amount":           {gconv.String(withdraw.Amount)},
						"status":           {gconv.String(5)},
						"remarks":          {withdraw.Remarks},
					}
					resp, _ := g.Client().PostForm(withdraw.NotifyUrl, data)
					g.Log().File("callback.{Y-m-d}.log").Printf("发送提现回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", withdraw.NotifyUrl, data.Encode(), resp.StatusCode)
					resp.Body.Close()
				}
			} else {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		}
	}

	var feeList []*model.FeeList
	g.Model("fee_list").Where("main_chain", "heco").Where("status", 1).Scan(&feeList)
	for _, fee := range feeList {
		data, _ := client.Init().GetTransferStatus(fee.Hash)
		err := mapstructure.Decode(data, &bscStruct)
		if err != nil || bscStruct.Status == "" {
			continue
		}
		if bscStruct.Status == "0x1" {
			g.Model("fee_list").Data(g.Map{"status": 2}).Where("id", fee.Id).Update()
			g.Model("recharge").Data(g.Map{"status": 1}).Where("id", fee.RechargeId).Update()
		} else {
			g.Model("fee_list").Data(g.Map{"status": 3}).Where("id", fee.Id).Update()
		}
	}

	var rechargeList []*model.Recharge
	g.Model("recharge").Where("main_chain", "heco").Where("status", 2).Scan(&rechargeList)

	for _, recharge := range rechargeList {
		data, _ := client.Init().GetTransferStatus(recharge.ImputationHash)
		err := mapstructure.Decode(data, &bscStruct)
		if err != nil || bscStruct.Status == "" {
			continue
		}
		if bscStruct.Status == "0x1" {
			g.Model("recharge").Data(g.Map{"status": 3}).Where("id", recharge.Id).Update()
			//归集成功发送通知
			recharge.Status = 3
			sendNotify(recharge)
		} else {
			g.Model("fee_list").Data(g.Map{"status": 4}).Where("id", recharge.Id).Update()
		}
	}
}

//tronCheckStatus 检查波场相关任务
func tronCheckStatus() {
	//查询所有待出金的任务
	var withdrawList []*model.Withdraw
	g.Model("withdraw").Where("main_chain", "tron").Where("status", 3).Scan(&withdrawList)
	var cache = service.Cache.New()
	var (
		err         error
		transaction *core.Transaction
	)
	tronClient, err := rpc.NewClient(gconv.String(cache.Get("tron_rpc_url")))
	if err != nil {
		g.Log().Printf("波场节点链接失败！！！%v", err)
		return
	}
	for _, withdraw := range withdrawList {
		transaction, err = tronClient.GRPC.GetTransactionByID(withdraw.Hash)
		if err == nil && transaction == nil {
			if gtime.Timestamp()-withdraw.UpdatedAt.Timestamp() >= 120 {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		} else {

			if transaction.Ret[0].ContractRet == core.Transaction_Result_SUCCESS {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 5, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 5, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
				if withdraw.NotifyUrl != "" {
					data := url.Values{
						"address":          {withdraw.Address},
						"contract_address": {gconv.String(withdraw.ContractAddress)},
						"amount":           {gconv.String(withdraw.Amount)},
						"status":           {gconv.String(5)},
						"remarks":          {withdraw.Remarks},
						"hash":             {withdraw.Hash},
						"token_id":         {gconv.String(withdraw.TokenId)},
						"url":              {withdraw.Url},
					}
					resp, _ := g.Client().PostForm(withdraw.NotifyUrl, data)
					g.Log().File("callback.{Y-m-d}.log").Printf("发送提现回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", withdraw.NotifyUrl, data.Encode(), resp.StatusCode)
					resp.Body.Close()
				}
			} else {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		}
	}

	var feeList []*model.FeeList
	g.Model("fee_list").Where("main_chain", "tron").Where("status", 1).Scan(&feeList)
	for _, fee := range feeList {
		transaction, err = tronClient.GRPC.GetTransactionByID(fee.Hash)
		if err != nil {
			return
		}

		if transaction.Ret[0].ContractRet == core.Transaction_Result_SUCCESS {
			g.Model("fee_list").Data(g.Map{"status": 2}).Where("id", fee.Id).Update()
			g.Model("recharge").Data(g.Map{"status": 1}).Where("id", fee.RechargeId).Update()
		} else {
			g.Model("fee_list").Data(g.Map{"status": 3}).Where("id", fee.Id).Update()
		}
	}

	var rechargeList []*model.Recharge
	g.Model("recharge").Where("main_chain", "tron").Where("status", 2).Scan(&rechargeList)

	for _, recharge := range rechargeList {
		transaction, err = tronClient.GRPC.GetTransactionByID(recharge.ImputationHash)
		if err != nil && transaction == nil {
			if gtime.Timestamp()-recharge.UpdatedAt.Timestamp() >= 120 {
				g.Model("recharge").Data(g.Map{"status": 1}).Where("id", recharge.Id).Update()
			}
		} else {
			if transaction.Ret[0].ContractRet == core.Transaction_Result_SUCCESS {
				g.Model("recharge").Data(g.Map{"status": 3}).Where("id", recharge.Id).Update()
				//归集成功发送通知
				recharge.Status = 3
				sendNotify(recharge)
			} else {
				g.Model("fee_list").Data(g.Map{"status": 4}).Where("id", recharge.Id).Update()
			}
		}
	}
}

//nacCheckStatus 检查nac相关任务
func nacCheckStatus() {
	//查询所有待出金的任务
	var withdrawList []*model.Withdraw
	g.Model("withdraw").Where("main_chain", "nac").Where("status", 3).Scan(&withdrawList)

	for _, withdraw := range withdrawList {
		txDetail, err := rpc.GetTxDetail(withdraw.Hash)
		if err != nil {
			continue
		}
		//检查，如果超过两分钟未找到hash，可以判定为失败了
		//if err == nil && txDetail == nil {
		//	if gtime.Timestamp()-withdraw.UpdatedAt.Timestamp() >= 120 {
		//		hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
		//		g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
		//	}
		//} else {
		if txDetail.RawTx.Status == 1 {
			g.Model("withdraw").Data(g.Map{"status": 5}).Where("id", withdraw.Id).Update()
			if withdraw.NotifyUrl != "" {
				data := url.Values{
					"address":          {withdraw.Address},
					"contract_address": {gconv.String(withdraw.ContractAddress)},
					"amount":           {gconv.String(withdraw.Amount)},
					"status":           {gconv.String(5)},
					"remarks":          {withdraw.Remarks},
					"hash":             {withdraw.Hash},
				}
				resp, _ := g.Client().PostForm(withdraw.NotifyUrl, data)
				g.Log().File("callback.{Y-m-d}.log").Printf("发送提现回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", withdraw.NotifyUrl, data.Encode(), resp.StatusCode)
				resp.Body.Close()
			}
		} else {
			hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
			g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
		}
	}

	var feeList []*model.FeeList
	g.Model("fee_list").Where("main_chain", "nac").Where("status", 1).Scan(&feeList)
	for _, fee := range feeList {
		txDetail, err := rpc.GetTxDetail(fee.Hash)
		if err != nil {
			continue
		}
		if txDetail.RawTx.Status == 1 {
			g.Model("fee_list").Data(g.Map{"status": 2}).Where("id", fee.Id).Update()
			g.Model("recharge").Data(g.Map{"status": 1}).Where("id", fee.RechargeId).Update()
		} else {
			g.Model("fee_list").Data(g.Map{"status": 3}).Where("id", fee.Id).Update()
		}
	}

	var rechargeList []*model.Recharge
	g.Model("recharge").Where("main_chain", "nac").Where("status", 2).Scan(&rechargeList)

	for _, recharge := range rechargeList {
		txDetail, err := rpc.GetTxDetail(recharge.ImputationHash)
		if err != nil {
			continue
		}
		if txDetail.RawTx.Status == 1 {
			g.Model("recharge").Data(g.Map{"status": 3}).Where("id", recharge.Id).Update()
			//归集成功发送通知
			recharge.Status = 3
			sendNotify(recharge)
		} else {
			g.Model("fee_list").Data(g.Map{"status": 4}).Where("id", recharge.Id).Update()
		}
	}
}

//wemixCheckStatus 检查WEMIX相关任务
func wemixCheckStatus() {
	//查询所有待出金的任务
	var withdrawList []*model.Withdraw
	g.Model("withdraw").Where("main_chain", "wemix").Where("status", 3).Scan(&withdrawList)
	client := rpc.WemixClient{}
	bscStruct := rpc.WemixTransferDetail{}

	for _, withdraw := range withdrawList {
		data, err := client.Init().GetTransferStatus(withdraw.Hash)
		//检查，如果超过五分钟未找到hash，可以判定为失败了
		if err == nil && data == nil {
			if gtime.Timestamp()-withdraw.UpdatedAt.Timestamp() >= 300 {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		} else {
			err = mapstructure.Decode(data, &bscStruct)
			if err != nil || bscStruct.Status == "" {
				continue
			}
			if bscStruct.Status == "0x1" {
				g.Model("withdraw").Data(g.Map{"status": 5}).Where("id", withdraw.Id).Update()
				if withdraw.NotifyUrl != "" {
					data := url.Values{
						"address":          {withdraw.Address},
						"contract_address": {gconv.String(withdraw.ContractAddress)},
						"amount":           {gconv.String(withdraw.Amount)},
						"status":           {gconv.String(5)},
						"remarks":          {withdraw.Remarks},
					}
					resp, _ := g.Client().PostForm(withdraw.NotifyUrl, data)
					g.Log().File("callback.{Y-m-d}.log").Printf("发送提现回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", withdraw.NotifyUrl, data.Encode(), resp.StatusCode)
					resp.Body.Close()
				}
			} else {
				hashKey := library.Md5Data(withdraw.Address, withdraw.ContractAddress, withdraw.Amount, 4, withdraw.Nonce1)
				g.Model("withdraw").Data(g.Map{"status": 4, "hashKey": hashKey}).Where("id", withdraw.Id).Update()
			}
		}
	}

	var feeList []*model.FeeList
	g.Model("fee_list").Where("main_chain", "wemix").Where("status", 1).Scan(&feeList)
	for _, fee := range feeList {
		data, _ := client.Init().GetTransferStatus(fee.Hash)
		err := mapstructure.Decode(data, &bscStruct)
		if err != nil || bscStruct.Status == "" {
			continue
		}
		if bscStruct.Status == "0x1" {
			g.Model("fee_list").Data(g.Map{"status": 2}).Where("id", fee.Id).Update()
			g.Model("recharge").Data(g.Map{"status": 1}).Where("id", fee.RechargeId).Update()
		} else {
			g.Model("fee_list").Data(g.Map{"status": 3}).Where("id", fee.Id).Update()
		}
	}

	var rechargeList []*model.Recharge
	g.Model("recharge").Where("main_chain", "wemix").Where("status", 2).Scan(&rechargeList)

	for _, recharge := range rechargeList {
		data, _ := client.Init().GetTransferStatus(recharge.ImputationHash)
		err := mapstructure.Decode(data, &bscStruct)
		if err != nil || bscStruct.Status == "" {
			continue
		}
		if bscStruct.Status == "0x1" {
			g.Model("recharge").Data(g.Map{"status": 3}).Where("id", recharge.Id).Update()
			//归集成功发送通知
			recharge.Status = 3
			sendNotify(recharge)
		} else {
			g.Model("fee_list").Data(g.Map{"status": 4}).Where("id", recharge.Id).Update()
		}
	}
}

func sendNotify(recharge *model.Recharge) {
	//检测是否配置了充值回调地址
	callbackUrl, _ := service.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
	if callbackUrl.ConfigValue == "" {
		g.Log().Path("data/log/callback.{Y-m-d}.log").Printf("未配置回调地址,不发送请求 %v \n", recharge)
		return
	}
	data := url.Values{
		"main_chain":        {recharge.MainChain},
		"block_hash":        {recharge.BlockHash},
		"recharge_type":     {gconv.String(recharge.RechargeType)},
		"from_address":      {recharge.FromAddress},
		"to_address":        {recharge.ToAddress},
		"coin_token":        {recharge.CoinToken},
		"coin_token1":       {recharge.CoinToken1},
		"contract_address":  {recharge.ContractAddress},
		"contract_address1": {recharge.ContractAddress1},
		"amount":            {gconv.String(recharge.Amount)},
		"amount1":           {gconv.String(recharge.Amount1)},
		"hash":              {recharge.Hash},
		"imputation_hash":   {""},
		"remarks":           {recharge.Remarks},
		"status":            {gconv.String(recharge.Status)},
		"token_id":          {recharge.TokenId},
	}
	resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, data)
	defer resp.Body.Close()
	g.Log().File("callback.{Y-m-d}.log").Printf("发送充值回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", callbackUrl.ConfigValue, data.Encode(), resp.StatusCode)
	return
}
