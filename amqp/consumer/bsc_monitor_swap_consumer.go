package consumer

import (
	"encoding/hex"
	"encoding/json"
	"gfast/app/common/service"
	"gfast/library"
	"gfast/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type BscMonitorSwapConsumer struct{}

// Consumer 实现币安扫块消费者
func (b *BscMonitorSwapConsumer) Consumer(dataByte []byte, key uint64) error {
	transfer := rpc.BscTransactions{}
	err := json.Unmarshal(dataByte, &transfer)
	if err != nil {
		return err
	} //json解析到结构体里面
	g.Log().File("monitor_consumer.{Y-m-d}.log").Printf("消费 %v", transfer.Hash)

	bscClient := rpc.BscClient{}
	transferReceipt, err := bscClient.Init().GetTransferStatus(transfer.Hash)
	transferDetail := rpc.BscTransferDetail{}
	err = mapstructure.Decode(transferReceipt, &transferDetail)
	if err != nil || transferDetail.Status == "" {
		g.Log().Printf("查询交易%v失败，%v# \n", transfer.Hash, err)
		return err
	}
	if transferDetail.Status == "0x0" {
		g.Log().Printf("交易状态是失败的，%v#不处理  %v#\n", transfer.Hash, err)
		return nil
	}
	if reader, err := os.Open("./abi/swap/swap.abi"); err == nil {
		defer reader.Close()
		if tokenAbi, err := abi.JSON(reader); err == nil {
			decodedSig, _ := hex.DecodeString(transfer.Input[2:10])
			method, _ := tokenAbi.MethodById(decodedSig)
			decodedData, _ := hex.DecodeString(transfer.Input[10:])
			type FunctionInputs struct {
				Path []common.Address
				To   common.Address
			}
			var data FunctionInputs
			inputMap := make(map[string]interface{}, 0)
			err = method.Inputs.UnpackIntoMap(inputMap, decodedData)
			if err != nil {
				return err
			}
			err = mapstructure.Decode(inputMap, &data)
			if err != nil {
				return err
			}
			var (
				amount           string
				amount1          string
				amount0String    string
				amount1String    string
				typeIndex        uint8
				contractAddress  string
				contractAddress1 string
			)
			monitorSwapCoinAddress, _ := service.MonitorCurrency.GetAddress("bsc")
			for keys, newPath := range data.Path {
				if keys == 0 {
					contractAddress = strings.ToLower(newPath.String())
					if library.ElementIsInSlice(contractAddress, monitorSwapCoinAddress) {
						typeIndex = 2
					}
				}
				if keys == len(data.Path)-1 {
					contractAddress1 = strings.ToLower(newPath.String())
					if library.ElementIsInSlice(contractAddress1, monitorSwapCoinAddress) {
						typeIndex = 1
					}
				}
			}
			blockHeight := strconv.FormatUint(gconv.Uint64(library.HexToBigInt(transfer.BlockNumber)), 10)

			for _, log := range transferDetail.Logs {
				if log.Topics[0] == "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822" {
					logData := log.Data[2:]
					if logData[:64] == "0000000000000000000000000000000000000000000000000000000000000000" {
						amount0String = "0x" + strings.TrimLeft(logData[64:128], "0")
						amount1String = "0x" + strings.TrimLeft(logData[128:192], "0")
					} else {
						amount0String = "0x" + strings.TrimLeft(logData[:64], "0")
						amount1String = "0x" + strings.TrimLeft(logData[192:], "0")
					}
					amount = gconv.String(library.HexToBigInt(amount0String))
					amount1 = gconv.String(library.HexToBigInt(amount1String))
				}
			}

			exists, _ := g.Model("monitor_swap_trade").Where("hash", transfer.Hash).Count()
			if exists <= 0 {
				g.Model("monitor_swap_trade").Data(g.Map{"main_chain": "bsc", "block_hash": transfer.BlockHash, "from_address": transfer.From,
					"contract_address": contractAddress, "contract_address1": contractAddress1, "amount": amount, "amount1": amount1, "path": contractAddress + "," + contractAddress1,
					"to_address": data.To.String(), "type": typeIndex, "hash": transfer.Hash, "block_height": blockHeight, "created_at": time.Now().Format("2006-01-02 15:04:05")}).Insert()

				callbackUrl, _ := service.SysConfig.GetConfigByKey("sys.rechargeCallbackUrl")
				if callbackUrl.ConfigValue == "" {
					g.Log().File("callback.{Y-m-d}.log").Printf("未配置回调地址,不发送请求 %v \n", transfer.Hash)
					return nil
				}
				postForm := url.Values{
					"main_chain":        {"bsc"},
					"block_hash":        {transfer.BlockHash},
					"from_address":      {transfer.From},
					"contract_address":  {contractAddress},
					"contract_address1": {contractAddress1},
					"amount":            {amount},
					"amount1":           {amount1},
					"path":              {contractAddress + "," + contractAddress1},
					"to_address":        {data.To.String()},
					"type":              {gconv.String(typeIndex)},
					"hash":              {transfer.Hash},
					"block_height":      {blockHeight},
				}
				resp, _ := g.Client().PostForm(callbackUrl.ConfigValue, postForm)
				defer resp.Body.Close()
				g.Log().File("callback.{Y-m-d}.log").Printf("发送监控交易回调请求 请求域名:【%v】 请求参数:【%v】 返回code码【%v】", callbackUrl.ConfigValue, postForm.Encode(), resp.StatusCode)

			}
		}
	}
	return err
}
