package consumer

import (
	"math/big"
	"strings"
	"testing"

	"gfast/app/system/model"
	"gfast/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestParseBHARechargeDepositEvent(t *testing.T) {
	contractABI, err := abi.JSON(strings.NewReader(bhaRechargeEventABI))
	if err != nil {
		t.Fatal(err)
	}
	depositEvent := contractABI.Events["DepositSuccess"]

	rechargeAddress := common.HexToAddress("0x1000000000000000000000000000000000000001")
	user := common.HexToAddress("0x2000000000000000000000000000000000000002")
	usdt := common.HexToAddress("0x3000000000000000000000000000000000000003")
	usda := common.HexToAddress("0x4000000000000000000000000000000000000004")
	receiver1 := common.HexToAddress("0x5000000000000000000000000000000000000005")
	receiver2 := common.HexToAddress("0x6000000000000000000000000000000000000006")
	receiver3 := common.HexToAddress("0x7000000000000000000000000000000000000007")
	tokens := []common.Address{usdt, usdt, usda}
	receivers := []common.Address{receiver1, receiver2, receiver3}
	amounts := []*big.Int{big.NewInt(40_000_000), big.NewInt(30_000_000), big.NewInt(10_000_000_000_000_000)}
	data, err := depositEvent.Inputs.NonIndexed().Pack("R202607050001", tokens, receivers, amounts, big.NewInt(1_800_000_000))
	if err != nil {
		t.Fatal(err)
	}
	logs := []rpc.BscTransferDetailLog{{
		Address: rechargeAddress.Hex(),
		Topics: []string{
			depositEvent.ID.Hex(),
			crypto.Keccak256Hash([]byte("R202607050001")).Hex(),
			common.BytesToHash(common.LeftPadBytes(user.Bytes(), 32)).Hex(),
		},
		Data: "0x" + common.Bytes2Hex(data),
	}}
	currencies := map[string]*model.Currency{
		strings.ToLower(usdt.Hex()): {Name: "USDT", Decimals: 6},
		strings.ToLower(usda.Hex()): {Name: "USDA", Decimals: 18},
	}

	record, err := parseBHARechargeDepositEvent(logs, rechargeAddress.Hex(), user.Hex(), currencies)
	if err != nil {
		t.Fatal(err)
	}
	if record == nil {
		t.Fatal("未解析到DepositSuccess事件")
	}
	if record.CoinToken != "USDT" || record.Amount != "70" || record.ContractAddress != strings.ToLower(usdt.Hex()) {
		t.Fatalf("USDT汇总错误: %+v", record)
	}
	if record.CoinToken1 != "USDA" || record.Amount1 != "0.01" || record.ContractAddress1 != strings.ToLower(usda.Hex()) {
		t.Fatalf("USDA汇总错误: %+v", record)
	}
	if record.Remarks != "R202607050001" || record.RechargeType != 2 {
		t.Fatalf("订单信息错误: %+v", record)
	}
	if record.CustomeCoin != `["0x3000000000000000000000000000000000000003","0x3000000000000000000000000000000000000003","0x4000000000000000000000000000000000000004"]` {
		t.Fatalf("币种明细错误: %s", record.CustomeCoin)
	}
	if record.CustomeAmount != `["40000000","30000000","10000000000000000"]` {
		t.Fatalf("数量明细错误: %s", record.CustomeAmount)
	}
}

func TestParseBHARechargeDepositEventRejectsMismatchedArrays(t *testing.T) {
	event := &bhaRechargeEvent{
		OrderNo:   "R202607050002",
		Tokens:    []common.Address{common.HexToAddress("0x3000000000000000000000000000000000000003")},
		Receivers: nil,
		Amounts:   []*big.Int{big.NewInt(1)},
	}

	if _, err := buildBHARechargeRecord(event, map[string]*model.Currency{}); err == nil {
		t.Fatal("数组长度不一致时应返回错误")
	}
}
