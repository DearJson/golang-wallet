package consumer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"gfast/app/system/model"
	"gfast/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

const bhaRechargeEventABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"orderHash","type":"bytes32"},{"indexed":true,"internalType":"address","name":"user","type":"address"},{"indexed":false,"internalType":"string","name":"orderNo","type":"string"},{"indexed":false,"internalType":"address[]","name":"tokens","type":"address[]"},{"indexed":false,"internalType":"address[]","name":"receivers","type":"address[]"},{"indexed":false,"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"indexed":false,"internalType":"uint256","name":"expireAt","type":"uint256"}],"name":"DepositSuccess","type":"event"}]`

type bhaRechargeEvent struct {
	OrderNo   string
	Tokens    []common.Address
	Receivers []common.Address
	Amounts   []*big.Int
	ExpireAt  *big.Int
}

type bhaRechargeRecord struct {
	CoinToken        string
	CoinToken1       string
	Amount           string
	Amount1          string
	ContractAddress  string
	ContractAddress1 string
	Remarks          string
	RechargeType     int8
	CustomeUser      string
	CustomeCoin      string
	CustomeAmount    string
}

func parseBHARechargeDepositEvent(logs []rpc.BscTransferDetailLog, rechargeAddress, fromAddress string, coinAddress map[string]*model.Currency) (*bhaRechargeRecord, error) {
	contractABI, err := abi.JSON(strings.NewReader(bhaRechargeEventABI))
	if err != nil {
		return nil, err
	}
	depositEvent := contractABI.Events["DepositSuccess"]

	for _, receiptLog := range logs {
		if !strings.EqualFold(receiptLog.Address, rechargeAddress) || len(receiptLog.Topics) == 0 || !strings.EqualFold(receiptLog.Topics[0], depositEvent.ID.Hex()) {
			continue
		}
		if len(receiptLog.Topics) < 3 {
			return nil, fmt.Errorf("DepositSuccess事件topics数量错误: %d", len(receiptLog.Topics))
		}

		user := common.BytesToAddress(common.HexToHash(receiptLog.Topics[2]).Bytes()[12:]).Hex()
		if !strings.EqualFold(user, fromAddress) {
			return nil, fmt.Errorf("DepositSuccess事件用户%s与交易发送方%s不一致", user, fromAddress)
		}

		data, err := hex.DecodeString(strings.TrimPrefix(receiptLog.Data, "0x"))
		if err != nil {
			return nil, fmt.Errorf("解析DepositSuccess事件数据失败: %w", err)
		}
		var event bhaRechargeEvent
		if err = contractABI.UnpackIntoInterface(&event, "DepositSuccess", data); err != nil {
			return nil, fmt.Errorf("解码DepositSuccess事件失败: %w", err)
		}
		return buildBHARechargeRecord(&event, coinAddress)
	}

	return nil, nil
}

func buildBHARechargeRecord(event *bhaRechargeEvent, coinAddress map[string]*model.Currency) (*bhaRechargeRecord, error) {
	if len(event.Tokens) == 0 || len(event.Tokens) != len(event.Receivers) || len(event.Tokens) != len(event.Amounts) {
		return nil, fmt.Errorf("DepositSuccess事件数组长度不一致")
	}

	uniqueTokens := make([]string, 0, len(event.Tokens))
	totals := make(map[string]*big.Int, len(event.Tokens))
	tokenNames := make(map[string]string, len(event.Tokens))
	customCoins := make([]string, 0, len(event.Tokens))
	customUsers := make([]string, 0, len(event.Receivers))
	customAmounts := make([]string, 0, len(event.Amounts))

	for index, token := range event.Tokens {
		tokenAddress := strings.ToLower(token.Hex())
		currency, ok := coinAddress[tokenAddress]
		if !ok {
			return nil, fmt.Errorf("DepositSuccess事件包含未配置币种: %s", tokenAddress)
		}
		if event.Amounts[index] == nil || event.Amounts[index].Sign() <= 0 {
			return nil, fmt.Errorf("DepositSuccess事件包含无效数量")
		}
		if _, ok = totals[tokenAddress]; !ok {
			uniqueTokens = append(uniqueTokens, tokenAddress)
			totals[tokenAddress] = new(big.Int)
			tokenNames[tokenAddress] = currency.Name
		}
		totals[tokenAddress].Add(totals[tokenAddress], event.Amounts[index])
		customCoins = append(customCoins, tokenAddress)
		customUsers = append(customUsers, strings.ToLower(event.Receivers[index].Hex()))
		customAmounts = append(customAmounts, event.Amounts[index].String())
	}
	if len(uniqueTokens) > 127 {
		return nil, fmt.Errorf("DepositSuccess事件币种数量过多: %d", len(uniqueTokens))
	}

	record := &bhaRechargeRecord{
		Remarks:      event.OrderNo,
		RechargeType: int8(len(uniqueTokens)),
	}
	for index, tokenAddress := range uniqueTokens {
		currency := coinAddress[tokenAddress]
		amount := decimal.NewFromBigInt(totals[tokenAddress], -int32(currency.Decimals)).String()
		switch index {
		case 0:
			record.CoinToken = tokenNames[tokenAddress]
			record.Amount = amount
			record.ContractAddress = tokenAddress
		case 1:
			record.CoinToken1 = tokenNames[tokenAddress]
			record.Amount1 = amount
			record.ContractAddress1 = tokenAddress
		}
	}

	var err error
	if record.CustomeCoin, err = marshalBHARechargeArray(customCoins); err != nil {
		return nil, err
	}
	if record.CustomeUser, err = marshalBHARechargeArray(customUsers); err != nil {
		return nil, err
	}
	if record.CustomeAmount, err = marshalBHARechargeArray(customAmounts); err != nil {
		return nil, err
	}
	return record, nil
}

func marshalBHARechargeArray(values []string) (string, error) {
	data, err := json.Marshal(values)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
