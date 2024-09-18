// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappHorc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DappHorcMetaData contains all meta data concerning the DappHorc contract.
var DappHorcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"_updateDailyUpRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addExtendList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"}],\"name\":\"addJoinToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addMintAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"SwapRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"horc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"DefaultInvitor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Fund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Fund2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Fund3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uplevel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ttype\",\"type\":\"uint256\"}],\"name\":\"ADDPOW\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"calInvite\",\"type\":\"bool\"}],\"name\":\"addUserAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"invitor\",\"type\":\"address\"}],\"name\":\"bindInvitor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ttype\",\"type\":\"uint256\"}],\"name\":\"CLAIMEVENT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"invitor\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num3\",\"type\":\"uint256\"}],\"name\":\"modselfLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num4\",\"type\":\"uint256\"}],\"name\":\"modselfRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invite2Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teamPow\",\"type\":\"uint256\"}],\"name\":\"modUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"set_dailyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"set_extendCurrent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"}],\"name\":\"set_extendList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"set_extendRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"set_isExtend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setAmountDailyUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAutoCalRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAutoCalSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setBurnReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setcanclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"}],\"name\":\"setDefaultInvitor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a3\",\"type\":\"address\"}],\"name\":\"setFundAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a3\",\"type\":\"address\"}],\"name\":\"setFundAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setHighPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setHighPriceSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setInviteFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setJoinToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"}],\"name\":\"setJoinToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isMain\",\"type\":\"bool[]\"}],\"name\":\"setJoinTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setLastAmountRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"setLastDailyUpTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setlimitUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setminAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setNormalRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setpowX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setstarttime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"succ\",\"type\":\"bool\"}],\"name\":\"setstop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setTotalMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setUpLevelAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardMintDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invite2Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teamPow\",\"type\":\"uint256\"}],\"name\":\"setuserinfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"__getTokenUsdtReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_amountDailyUp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_autoCalRewardPerDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_autoCalSellRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_binder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_burnReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_dailyDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_defaultInvitor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_extendCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_extendList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_extendRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fundAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fundAddress2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fundAddress3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_highPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_highPriceSellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_inviteAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_inviteFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_invitor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_isExtend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_joinTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_joinTokensType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastAmountRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastDailyUpTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_mintRewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_normalRewardPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_pauseJoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_teamNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_upLevelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canclaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdtDecimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintRewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintRewardTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDailyReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dailyAmountRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"defaultInvitor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pauseJoin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBinderLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDueInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getInvite\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"}],\"name\":\"getJoinTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getJoinTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenDecimals\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"tokenSymbols\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"poolUsdts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"poolTokens\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPendingMintReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getSelfRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sellRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenUsdtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"joinTokenBalances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"joinTokenAllowances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"teamNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfLimit1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfLimit2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfLimit3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfRate1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfRate2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfRate3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfRate4\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"starttime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewPoolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMintPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastMintTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMintReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"viewUserInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardMintDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invite2Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teamPow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DappHorcABI is the input ABI used to generate the binding from.
// Deprecated: Use DappHorcMetaData.ABI instead.
var DappHorcABI = DappHorcMetaData.ABI

// DappHorc is an auto generated Go binding around an Ethereum contract.
type DappHorc struct {
	DappHorcCaller     // Read-only binding to the contract
	DappHorcTransactor // Write-only binding to the contract
	DappHorcFilterer   // Log filterer for contract events
}

// DappHorcCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappHorcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappHorcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappHorcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappHorcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappHorcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappHorcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappHorcSession struct {
	Contract     *DappHorc         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappHorcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappHorcCallerSession struct {
	Contract *DappHorcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DappHorcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappHorcTransactorSession struct {
	Contract     *DappHorcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DappHorcRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappHorcRaw struct {
	Contract *DappHorc // Generic contract binding to access the raw methods on
}

// DappHorcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappHorcCallerRaw struct {
	Contract *DappHorcCaller // Generic read-only contract binding to access the raw methods on
}

// DappHorcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappHorcTransactorRaw struct {
	Contract *DappHorcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappHorc creates a new instance of DappHorc, bound to a specific deployed contract.
func NewDappHorc(address common.Address, backend bind.ContractBackend) (*DappHorc, error) {
	contract, err := bindDappHorc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappHorc{DappHorcCaller: DappHorcCaller{contract: contract}, DappHorcTransactor: DappHorcTransactor{contract: contract}, DappHorcFilterer: DappHorcFilterer{contract: contract}}, nil
}

// NewDappHorcCaller creates a new read-only instance of DappHorc, bound to a specific deployed contract.
func NewDappHorcCaller(address common.Address, caller bind.ContractCaller) (*DappHorcCaller, error) {
	contract, err := bindDappHorc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappHorcCaller{contract: contract}, nil
}

// NewDappHorcTransactor creates a new write-only instance of DappHorc, bound to a specific deployed contract.
func NewDappHorcTransactor(address common.Address, transactor bind.ContractTransactor) (*DappHorcTransactor, error) {
	contract, err := bindDappHorc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappHorcTransactor{contract: contract}, nil
}

// NewDappHorcFilterer creates a new log filterer instance of DappHorc, bound to a specific deployed contract.
func NewDappHorcFilterer(address common.Address, filterer bind.ContractFilterer) (*DappHorcFilterer, error) {
	contract, err := bindDappHorc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappHorcFilterer{contract: contract}, nil
}

// bindDappHorc binds a generic wrapper to an already deployed contract.
func bindDappHorc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DappHorcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappHorc *DappHorcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappHorc.Contract.DappHorcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappHorc *DappHorcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.Contract.DappHorcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappHorc *DappHorcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappHorc.Contract.DappHorcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappHorc *DappHorcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappHorc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappHorc *DappHorcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappHorc *DappHorcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappHorc.Contract.contract.Transact(opts, method, params...)
}

// GetTokenUsdtReserves is a free data retrieval call binding the contract method 0xb827f734.
//
// Solidity: function __getTokenUsdtReserves(address token) view returns(uint256 rUsdt, uint256 rToken)
func (_DappHorc *DappHorcCaller) GetTokenUsdtReserves(opts *bind.CallOpts, token common.Address) (struct {
	RUsdt  *big.Int
	RToken *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "__getTokenUsdtReserves", token)

	outstruct := new(struct {
		RUsdt  *big.Int
		RToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RUsdt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenUsdtReserves is a free data retrieval call binding the contract method 0xb827f734.
//
// Solidity: function __getTokenUsdtReserves(address token) view returns(uint256 rUsdt, uint256 rToken)
func (_DappHorc *DappHorcSession) GetTokenUsdtReserves(token common.Address) (struct {
	RUsdt  *big.Int
	RToken *big.Int
}, error) {
	return _DappHorc.Contract.GetTokenUsdtReserves(&_DappHorc.CallOpts, token)
}

// GetTokenUsdtReserves is a free data retrieval call binding the contract method 0xb827f734.
//
// Solidity: function __getTokenUsdtReserves(address token) view returns(uint256 rUsdt, uint256 rToken)
func (_DappHorc *DappHorcCallerSession) GetTokenUsdtReserves(token common.Address) (struct {
	RUsdt  *big.Int
	RToken *big.Int
}, error) {
	return _DappHorc.Contract.GetTokenUsdtReserves(&_DappHorc.CallOpts, token)
}

// AmountDailyUp is a free data retrieval call binding the contract method 0x5f84aecd.
//
// Solidity: function _amountDailyUp() view returns(uint256)
func (_DappHorc *DappHorcCaller) AmountDailyUp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_amountDailyUp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDailyUp is a free data retrieval call binding the contract method 0x5f84aecd.
//
// Solidity: function _amountDailyUp() view returns(uint256)
func (_DappHorc *DappHorcSession) AmountDailyUp() (*big.Int, error) {
	return _DappHorc.Contract.AmountDailyUp(&_DappHorc.CallOpts)
}

// AmountDailyUp is a free data retrieval call binding the contract method 0x5f84aecd.
//
// Solidity: function _amountDailyUp() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) AmountDailyUp() (*big.Int, error) {
	return _DappHorc.Contract.AmountDailyUp(&_DappHorc.CallOpts)
}

// AutoCalRewardPerDay is a free data retrieval call binding the contract method 0x4a09bfae.
//
// Solidity: function _autoCalRewardPerDay() view returns(bool)
func (_DappHorc *DappHorcCaller) AutoCalRewardPerDay(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_autoCalRewardPerDay")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutoCalRewardPerDay is a free data retrieval call binding the contract method 0x4a09bfae.
//
// Solidity: function _autoCalRewardPerDay() view returns(bool)
func (_DappHorc *DappHorcSession) AutoCalRewardPerDay() (bool, error) {
	return _DappHorc.Contract.AutoCalRewardPerDay(&_DappHorc.CallOpts)
}

// AutoCalRewardPerDay is a free data retrieval call binding the contract method 0x4a09bfae.
//
// Solidity: function _autoCalRewardPerDay() view returns(bool)
func (_DappHorc *DappHorcCallerSession) AutoCalRewardPerDay() (bool, error) {
	return _DappHorc.Contract.AutoCalRewardPerDay(&_DappHorc.CallOpts)
}

// AutoCalSellRate is a free data retrieval call binding the contract method 0x29f11c45.
//
// Solidity: function _autoCalSellRate() view returns(bool)
func (_DappHorc *DappHorcCaller) AutoCalSellRate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_autoCalSellRate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutoCalSellRate is a free data retrieval call binding the contract method 0x29f11c45.
//
// Solidity: function _autoCalSellRate() view returns(bool)
func (_DappHorc *DappHorcSession) AutoCalSellRate() (bool, error) {
	return _DappHorc.Contract.AutoCalSellRate(&_DappHorc.CallOpts)
}

// AutoCalSellRate is a free data retrieval call binding the contract method 0x29f11c45.
//
// Solidity: function _autoCalSellRate() view returns(bool)
func (_DappHorc *DappHorcCallerSession) AutoCalSellRate() (bool, error) {
	return _DappHorc.Contract.AutoCalSellRate(&_DappHorc.CallOpts)
}

// Binder is a free data retrieval call binding the contract method 0x8593b2b1.
//
// Solidity: function _binder(address , uint256 ) view returns(address)
func (_DappHorc *DappHorcCaller) Binder(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_binder", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Binder is a free data retrieval call binding the contract method 0x8593b2b1.
//
// Solidity: function _binder(address , uint256 ) view returns(address)
func (_DappHorc *DappHorcSession) Binder(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DappHorc.Contract.Binder(&_DappHorc.CallOpts, arg0, arg1)
}

// Binder is a free data retrieval call binding the contract method 0x8593b2b1.
//
// Solidity: function _binder(address , uint256 ) view returns(address)
func (_DappHorc *DappHorcCallerSession) Binder(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DappHorc.Contract.Binder(&_DappHorc.CallOpts, arg0, arg1)
}

// BurnReceiver is a free data retrieval call binding the contract method 0x9ba14907.
//
// Solidity: function _burnReceiver() view returns(address)
func (_DappHorc *DappHorcCaller) BurnReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_burnReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnReceiver is a free data retrieval call binding the contract method 0x9ba14907.
//
// Solidity: function _burnReceiver() view returns(address)
func (_DappHorc *DappHorcSession) BurnReceiver() (common.Address, error) {
	return _DappHorc.Contract.BurnReceiver(&_DappHorc.CallOpts)
}

// BurnReceiver is a free data retrieval call binding the contract method 0x9ba14907.
//
// Solidity: function _burnReceiver() view returns(address)
func (_DappHorc *DappHorcCallerSession) BurnReceiver() (common.Address, error) {
	return _DappHorc.Contract.BurnReceiver(&_DappHorc.CallOpts)
}

// DailyDuration is a free data retrieval call binding the contract method 0x66c67c39.
//
// Solidity: function _dailyDuration() view returns(uint256)
func (_DappHorc *DappHorcCaller) DailyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_dailyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyDuration is a free data retrieval call binding the contract method 0x66c67c39.
//
// Solidity: function _dailyDuration() view returns(uint256)
func (_DappHorc *DappHorcSession) DailyDuration() (*big.Int, error) {
	return _DappHorc.Contract.DailyDuration(&_DappHorc.CallOpts)
}

// DailyDuration is a free data retrieval call binding the contract method 0x66c67c39.
//
// Solidity: function _dailyDuration() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) DailyDuration() (*big.Int, error) {
	return _DappHorc.Contract.DailyDuration(&_DappHorc.CallOpts)
}

// DefaultInvitor is a free data retrieval call binding the contract method 0xab806d23.
//
// Solidity: function _defaultInvitor() view returns(address)
func (_DappHorc *DappHorcCaller) DefaultInvitor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_defaultInvitor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultInvitor is a free data retrieval call binding the contract method 0xab806d23.
//
// Solidity: function _defaultInvitor() view returns(address)
func (_DappHorc *DappHorcSession) DefaultInvitor() (common.Address, error) {
	return _DappHorc.Contract.DefaultInvitor(&_DappHorc.CallOpts)
}

// DefaultInvitor is a free data retrieval call binding the contract method 0xab806d23.
//
// Solidity: function _defaultInvitor() view returns(address)
func (_DappHorc *DappHorcCallerSession) DefaultInvitor() (common.Address, error) {
	return _DappHorc.Contract.DefaultInvitor(&_DappHorc.CallOpts)
}

// ExtendCurrent is a free data retrieval call binding the contract method 0xe14b1c16.
//
// Solidity: function _extendCurrent() view returns(uint256)
func (_DappHorc *DappHorcCaller) ExtendCurrent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_extendCurrent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtendCurrent is a free data retrieval call binding the contract method 0xe14b1c16.
//
// Solidity: function _extendCurrent() view returns(uint256)
func (_DappHorc *DappHorcSession) ExtendCurrent() (*big.Int, error) {
	return _DappHorc.Contract.ExtendCurrent(&_DappHorc.CallOpts)
}

// ExtendCurrent is a free data retrieval call binding the contract method 0xe14b1c16.
//
// Solidity: function _extendCurrent() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) ExtendCurrent() (*big.Int, error) {
	return _DappHorc.Contract.ExtendCurrent(&_DappHorc.CallOpts)
}

// ExtendList is a free data retrieval call binding the contract method 0x7f147c3e.
//
// Solidity: function _extendList(uint256 ) view returns(address)
func (_DappHorc *DappHorcCaller) ExtendList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_extendList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtendList is a free data retrieval call binding the contract method 0x7f147c3e.
//
// Solidity: function _extendList(uint256 ) view returns(address)
func (_DappHorc *DappHorcSession) ExtendList(arg0 *big.Int) (common.Address, error) {
	return _DappHorc.Contract.ExtendList(&_DappHorc.CallOpts, arg0)
}

// ExtendList is a free data retrieval call binding the contract method 0x7f147c3e.
//
// Solidity: function _extendList(uint256 ) view returns(address)
func (_DappHorc *DappHorcCallerSession) ExtendList(arg0 *big.Int) (common.Address, error) {
	return _DappHorc.Contract.ExtendList(&_DappHorc.CallOpts, arg0)
}

// ExtendRate is a free data retrieval call binding the contract method 0x8b68b741.
//
// Solidity: function _extendRate() view returns(uint256)
func (_DappHorc *DappHorcCaller) ExtendRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_extendRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtendRate is a free data retrieval call binding the contract method 0x8b68b741.
//
// Solidity: function _extendRate() view returns(uint256)
func (_DappHorc *DappHorcSession) ExtendRate() (*big.Int, error) {
	return _DappHorc.Contract.ExtendRate(&_DappHorc.CallOpts)
}

// ExtendRate is a free data retrieval call binding the contract method 0x8b68b741.
//
// Solidity: function _extendRate() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) ExtendRate() (*big.Int, error) {
	return _DappHorc.Contract.ExtendRate(&_DappHorc.CallOpts)
}

// FundAddress is a free data retrieval call binding the contract method 0x8230af5a.
//
// Solidity: function _fundAddress() view returns(address)
func (_DappHorc *DappHorcCaller) FundAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_fundAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundAddress is a free data retrieval call binding the contract method 0x8230af5a.
//
// Solidity: function _fundAddress() view returns(address)
func (_DappHorc *DappHorcSession) FundAddress() (common.Address, error) {
	return _DappHorc.Contract.FundAddress(&_DappHorc.CallOpts)
}

// FundAddress is a free data retrieval call binding the contract method 0x8230af5a.
//
// Solidity: function _fundAddress() view returns(address)
func (_DappHorc *DappHorcCallerSession) FundAddress() (common.Address, error) {
	return _DappHorc.Contract.FundAddress(&_DappHorc.CallOpts)
}

// FundAddress2 is a free data retrieval call binding the contract method 0xcdc695b9.
//
// Solidity: function _fundAddress2() view returns(address)
func (_DappHorc *DappHorcCaller) FundAddress2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_fundAddress2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundAddress2 is a free data retrieval call binding the contract method 0xcdc695b9.
//
// Solidity: function _fundAddress2() view returns(address)
func (_DappHorc *DappHorcSession) FundAddress2() (common.Address, error) {
	return _DappHorc.Contract.FundAddress2(&_DappHorc.CallOpts)
}

// FundAddress2 is a free data retrieval call binding the contract method 0xcdc695b9.
//
// Solidity: function _fundAddress2() view returns(address)
func (_DappHorc *DappHorcCallerSession) FundAddress2() (common.Address, error) {
	return _DappHorc.Contract.FundAddress2(&_DappHorc.CallOpts)
}

// FundAddress3 is a free data retrieval call binding the contract method 0x177f15d0.
//
// Solidity: function _fundAddress3() view returns(address)
func (_DappHorc *DappHorcCaller) FundAddress3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_fundAddress3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundAddress3 is a free data retrieval call binding the contract method 0x177f15d0.
//
// Solidity: function _fundAddress3() view returns(address)
func (_DappHorc *DappHorcSession) FundAddress3() (common.Address, error) {
	return _DappHorc.Contract.FundAddress3(&_DappHorc.CallOpts)
}

// FundAddress3 is a free data retrieval call binding the contract method 0x177f15d0.
//
// Solidity: function _fundAddress3() view returns(address)
func (_DappHorc *DappHorcCallerSession) FundAddress3() (common.Address, error) {
	return _DappHorc.Contract.FundAddress3(&_DappHorc.CallOpts)
}

// HighPrice is a free data retrieval call binding the contract method 0xe15c207a.
//
// Solidity: function _highPrice() view returns(uint256)
func (_DappHorc *DappHorcCaller) HighPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_highPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighPrice is a free data retrieval call binding the contract method 0xe15c207a.
//
// Solidity: function _highPrice() view returns(uint256)
func (_DappHorc *DappHorcSession) HighPrice() (*big.Int, error) {
	return _DappHorc.Contract.HighPrice(&_DappHorc.CallOpts)
}

// HighPrice is a free data retrieval call binding the contract method 0xe15c207a.
//
// Solidity: function _highPrice() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) HighPrice() (*big.Int, error) {
	return _DappHorc.Contract.HighPrice(&_DappHorc.CallOpts)
}

// HighPriceSellRate is a free data retrieval call binding the contract method 0xdbdaa44e.
//
// Solidity: function _highPriceSellRate() view returns(uint256)
func (_DappHorc *DappHorcCaller) HighPriceSellRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_highPriceSellRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighPriceSellRate is a free data retrieval call binding the contract method 0xdbdaa44e.
//
// Solidity: function _highPriceSellRate() view returns(uint256)
func (_DappHorc *DappHorcSession) HighPriceSellRate() (*big.Int, error) {
	return _DappHorc.Contract.HighPriceSellRate(&_DappHorc.CallOpts)
}

// HighPriceSellRate is a free data retrieval call binding the contract method 0xdbdaa44e.
//
// Solidity: function _highPriceSellRate() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) HighPriceSellRate() (*big.Int, error) {
	return _DappHorc.Contract.HighPriceSellRate(&_DappHorc.CallOpts)
}

// InviteAmount is a free data retrieval call binding the contract method 0xac86af94.
//
// Solidity: function _inviteAmount(address ) view returns(uint256)
func (_DappHorc *DappHorcCaller) InviteAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_inviteAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteAmount is a free data retrieval call binding the contract method 0xac86af94.
//
// Solidity: function _inviteAmount(address ) view returns(uint256)
func (_DappHorc *DappHorcSession) InviteAmount(arg0 common.Address) (*big.Int, error) {
	return _DappHorc.Contract.InviteAmount(&_DappHorc.CallOpts, arg0)
}

// InviteAmount is a free data retrieval call binding the contract method 0xac86af94.
//
// Solidity: function _inviteAmount(address ) view returns(uint256)
func (_DappHorc *DappHorcCallerSession) InviteAmount(arg0 common.Address) (*big.Int, error) {
	return _DappHorc.Contract.InviteAmount(&_DappHorc.CallOpts, arg0)
}

// InviteFee is a free data retrieval call binding the contract method 0xcef6632d.
//
// Solidity: function _inviteFee(uint256 ) view returns(uint256)
func (_DappHorc *DappHorcCaller) InviteFee(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_inviteFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteFee is a free data retrieval call binding the contract method 0xcef6632d.
//
// Solidity: function _inviteFee(uint256 ) view returns(uint256)
func (_DappHorc *DappHorcSession) InviteFee(arg0 *big.Int) (*big.Int, error) {
	return _DappHorc.Contract.InviteFee(&_DappHorc.CallOpts, arg0)
}

// InviteFee is a free data retrieval call binding the contract method 0xcef6632d.
//
// Solidity: function _inviteFee(uint256 ) view returns(uint256)
func (_DappHorc *DappHorcCallerSession) InviteFee(arg0 *big.Int) (*big.Int, error) {
	return _DappHorc.Contract.InviteFee(&_DappHorc.CallOpts, arg0)
}

// Invitor is a free data retrieval call binding the contract method 0x69da1326.
//
// Solidity: function _invitor(address ) view returns(address)
func (_DappHorc *DappHorcCaller) Invitor(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_invitor", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Invitor is a free data retrieval call binding the contract method 0x69da1326.
//
// Solidity: function _invitor(address ) view returns(address)
func (_DappHorc *DappHorcSession) Invitor(arg0 common.Address) (common.Address, error) {
	return _DappHorc.Contract.Invitor(&_DappHorc.CallOpts, arg0)
}

// Invitor is a free data retrieval call binding the contract method 0x69da1326.
//
// Solidity: function _invitor(address ) view returns(address)
func (_DappHorc *DappHorcCallerSession) Invitor(arg0 common.Address) (common.Address, error) {
	return _DappHorc.Contract.Invitor(&_DappHorc.CallOpts, arg0)
}

// IsExtend is a free data retrieval call binding the contract method 0xbdff0adf.
//
// Solidity: function _isExtend() view returns(bool)
func (_DappHorc *DappHorcCaller) IsExtend(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_isExtend")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExtend is a free data retrieval call binding the contract method 0xbdff0adf.
//
// Solidity: function _isExtend() view returns(bool)
func (_DappHorc *DappHorcSession) IsExtend() (bool, error) {
	return _DappHorc.Contract.IsExtend(&_DappHorc.CallOpts)
}

// IsExtend is a free data retrieval call binding the contract method 0xbdff0adf.
//
// Solidity: function _isExtend() view returns(bool)
func (_DappHorc *DappHorcCallerSession) IsExtend() (bool, error) {
	return _DappHorc.Contract.IsExtend(&_DappHorc.CallOpts)
}

// JoinTokens is a free data retrieval call binding the contract method 0xf764356a.
//
// Solidity: function _joinTokens(uint256 ) view returns(address)
func (_DappHorc *DappHorcCaller) JoinTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_joinTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JoinTokens is a free data retrieval call binding the contract method 0xf764356a.
//
// Solidity: function _joinTokens(uint256 ) view returns(address)
func (_DappHorc *DappHorcSession) JoinTokens(arg0 *big.Int) (common.Address, error) {
	return _DappHorc.Contract.JoinTokens(&_DappHorc.CallOpts, arg0)
}

// JoinTokens is a free data retrieval call binding the contract method 0xf764356a.
//
// Solidity: function _joinTokens(uint256 ) view returns(address)
func (_DappHorc *DappHorcCallerSession) JoinTokens(arg0 *big.Int) (common.Address, error) {
	return _DappHorc.Contract.JoinTokens(&_DappHorc.CallOpts, arg0)
}

// JoinTokensType is a free data retrieval call binding the contract method 0x80e82eec.
//
// Solidity: function _joinTokensType(address ) view returns(bool)
func (_DappHorc *DappHorcCaller) JoinTokensType(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_joinTokensType", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// JoinTokensType is a free data retrieval call binding the contract method 0x80e82eec.
//
// Solidity: function _joinTokensType(address ) view returns(bool)
func (_DappHorc *DappHorcSession) JoinTokensType(arg0 common.Address) (bool, error) {
	return _DappHorc.Contract.JoinTokensType(&_DappHorc.CallOpts, arg0)
}

// JoinTokensType is a free data retrieval call binding the contract method 0x80e82eec.
//
// Solidity: function _joinTokensType(address ) view returns(bool)
func (_DappHorc *DappHorcCallerSession) JoinTokensType(arg0 common.Address) (bool, error) {
	return _DappHorc.Contract.JoinTokensType(&_DappHorc.CallOpts, arg0)
}

// LastAmountRate is a free data retrieval call binding the contract method 0x1703dd12.
//
// Solidity: function _lastAmountRate() view returns(uint256)
func (_DappHorc *DappHorcCaller) LastAmountRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_lastAmountRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAmountRate is a free data retrieval call binding the contract method 0x1703dd12.
//
// Solidity: function _lastAmountRate() view returns(uint256)
func (_DappHorc *DappHorcSession) LastAmountRate() (*big.Int, error) {
	return _DappHorc.Contract.LastAmountRate(&_DappHorc.CallOpts)
}

// LastAmountRate is a free data retrieval call binding the contract method 0x1703dd12.
//
// Solidity: function _lastAmountRate() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) LastAmountRate() (*big.Int, error) {
	return _DappHorc.Contract.LastAmountRate(&_DappHorc.CallOpts)
}

// LastDailyUpTime is a free data retrieval call binding the contract method 0xd5559e21.
//
// Solidity: function _lastDailyUpTime() view returns(uint256)
func (_DappHorc *DappHorcCaller) LastDailyUpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_lastDailyUpTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDailyUpTime is a free data retrieval call binding the contract method 0xd5559e21.
//
// Solidity: function _lastDailyUpTime() view returns(uint256)
func (_DappHorc *DappHorcSession) LastDailyUpTime() (*big.Int, error) {
	return _DappHorc.Contract.LastDailyUpTime(&_DappHorc.CallOpts)
}

// LastDailyUpTime is a free data retrieval call binding the contract method 0xd5559e21.
//
// Solidity: function _lastDailyUpTime() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) LastDailyUpTime() (*big.Int, error) {
	return _DappHorc.Contract.LastDailyUpTime(&_DappHorc.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x064d4d7c.
//
// Solidity: function _minAmount() view returns(uint256)
func (_DappHorc *DappHorcCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x064d4d7c.
//
// Solidity: function _minAmount() view returns(uint256)
func (_DappHorc *DappHorcSession) MinAmount() (*big.Int, error) {
	return _DappHorc.Contract.MinAmount(&_DappHorc.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x064d4d7c.
//
// Solidity: function _minAmount() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) MinAmount() (*big.Int, error) {
	return _DappHorc.Contract.MinAmount(&_DappHorc.CallOpts)
}

// MintRewardToken is a free data retrieval call binding the contract method 0x5ac09125.
//
// Solidity: function _mintRewardToken() view returns(address)
func (_DappHorc *DappHorcCaller) MintRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_mintRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintRewardToken is a free data retrieval call binding the contract method 0x5ac09125.
//
// Solidity: function _mintRewardToken() view returns(address)
func (_DappHorc *DappHorcSession) MintRewardToken() (common.Address, error) {
	return _DappHorc.Contract.MintRewardToken(&_DappHorc.CallOpts)
}

// MintRewardToken is a free data retrieval call binding the contract method 0x5ac09125.
//
// Solidity: function _mintRewardToken() view returns(address)
func (_DappHorc *DappHorcCallerSession) MintRewardToken() (common.Address, error) {
	return _DappHorc.Contract.MintRewardToken(&_DappHorc.CallOpts)
}

// NormalRewardPerDay is a free data retrieval call binding the contract method 0xb2e21962.
//
// Solidity: function _normalRewardPerDay() view returns(uint256)
func (_DappHorc *DappHorcCaller) NormalRewardPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_normalRewardPerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalRewardPerDay is a free data retrieval call binding the contract method 0xb2e21962.
//
// Solidity: function _normalRewardPerDay() view returns(uint256)
func (_DappHorc *DappHorcSession) NormalRewardPerDay() (*big.Int, error) {
	return _DappHorc.Contract.NormalRewardPerDay(&_DappHorc.CallOpts)
}

// NormalRewardPerDay is a free data retrieval call binding the contract method 0xb2e21962.
//
// Solidity: function _normalRewardPerDay() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) NormalRewardPerDay() (*big.Int, error) {
	return _DappHorc.Contract.NormalRewardPerDay(&_DappHorc.CallOpts)
}

// PauseJoin is a free data retrieval call binding the contract method 0xfa14616b.
//
// Solidity: function _pauseJoin() view returns(bool)
func (_DappHorc *DappHorcCaller) PauseJoin(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_pauseJoin")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseJoin is a free data retrieval call binding the contract method 0xfa14616b.
//
// Solidity: function _pauseJoin() view returns(bool)
func (_DappHorc *DappHorcSession) PauseJoin() (bool, error) {
	return _DappHorc.Contract.PauseJoin(&_DappHorc.CallOpts)
}

// PauseJoin is a free data retrieval call binding the contract method 0xfa14616b.
//
// Solidity: function _pauseJoin() view returns(bool)
func (_DappHorc *DappHorcCallerSession) PauseJoin() (bool, error) {
	return _DappHorc.Contract.PauseJoin(&_DappHorc.CallOpts)
}

// SellRate is a free data retrieval call binding the contract method 0x4308fa6a.
//
// Solidity: function _sellRate() view returns(uint256)
func (_DappHorc *DappHorcCaller) SellRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_sellRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRate is a free data retrieval call binding the contract method 0x4308fa6a.
//
// Solidity: function _sellRate() view returns(uint256)
func (_DappHorc *DappHorcSession) SellRate() (*big.Int, error) {
	return _DappHorc.Contract.SellRate(&_DappHorc.CallOpts)
}

// SellRate is a free data retrieval call binding the contract method 0x4308fa6a.
//
// Solidity: function _sellRate() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SellRate() (*big.Int, error) {
	return _DappHorc.Contract.SellRate(&_DappHorc.CallOpts)
}

// Stop is a free data retrieval call binding the contract method 0x6853020f.
//
// Solidity: function _stop(address ) view returns(bool)
func (_DappHorc *DappHorcCaller) Stop(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_stop", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stop is a free data retrieval call binding the contract method 0x6853020f.
//
// Solidity: function _stop(address ) view returns(bool)
func (_DappHorc *DappHorcSession) Stop(arg0 common.Address) (bool, error) {
	return _DappHorc.Contract.Stop(&_DappHorc.CallOpts, arg0)
}

// Stop is a free data retrieval call binding the contract method 0x6853020f.
//
// Solidity: function _stop(address ) view returns(bool)
func (_DappHorc *DappHorcCallerSession) Stop(arg0 common.Address) (bool, error) {
	return _DappHorc.Contract.Stop(&_DappHorc.CallOpts, arg0)
}

// TeamNum is a free data retrieval call binding the contract method 0x93b556ad.
//
// Solidity: function _teamNum(address ) view returns(uint256)
func (_DappHorc *DappHorcCaller) TeamNum(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_teamNum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamNum is a free data retrieval call binding the contract method 0x93b556ad.
//
// Solidity: function _teamNum(address ) view returns(uint256)
func (_DappHorc *DappHorcSession) TeamNum(arg0 common.Address) (*big.Int, error) {
	return _DappHorc.Contract.TeamNum(&_DappHorc.CallOpts, arg0)
}

// TeamNum is a free data retrieval call binding the contract method 0x93b556ad.
//
// Solidity: function _teamNum(address ) view returns(uint256)
func (_DappHorc *DappHorcCallerSession) TeamNum(arg0 common.Address) (*big.Int, error) {
	return _DappHorc.Contract.TeamNum(&_DappHorc.CallOpts, arg0)
}

// UpLevelAddr is a free data retrieval call binding the contract method 0x3ea7ef6a.
//
// Solidity: function _upLevelAddr() view returns(address)
func (_DappHorc *DappHorcCaller) UpLevelAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_upLevelAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpLevelAddr is a free data retrieval call binding the contract method 0x3ea7ef6a.
//
// Solidity: function _upLevelAddr() view returns(address)
func (_DappHorc *DappHorcSession) UpLevelAddr() (common.Address, error) {
	return _DappHorc.Contract.UpLevelAddr(&_DappHorc.CallOpts)
}

// UpLevelAddr is a free data retrieval call binding the contract method 0x3ea7ef6a.
//
// Solidity: function _upLevelAddr() view returns(address)
func (_DappHorc *DappHorcCallerSession) UpLevelAddr() (common.Address, error) {
	return _DappHorc.Contract.UpLevelAddr(&_DappHorc.CallOpts)
}

// Canclaim is a free data retrieval call binding the contract method 0x2d813c47.
//
// Solidity: function canclaim() view returns(bool)
func (_DappHorc *DappHorcCaller) Canclaim(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "canclaim")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Canclaim is a free data retrieval call binding the contract method 0x2d813c47.
//
// Solidity: function canclaim() view returns(bool)
func (_DappHorc *DappHorcSession) Canclaim() (bool, error) {
	return _DappHorc.Contract.Canclaim(&_DappHorc.CallOpts)
}

// Canclaim is a free data retrieval call binding the contract method 0x2d813c47.
//
// Solidity: function canclaim() view returns(bool)
func (_DappHorc *DappHorcCallerSession) Canclaim() (bool, error) {
	return _DappHorc.Contract.Canclaim(&_DappHorc.CallOpts)
}

// GetBaseInfo is a free data retrieval call binding the contract method 0x9551ae44.
//
// Solidity: function getBaseInfo() view returns(address usdt, uint256 usdtDecimals, address mintRewardToken, uint256 mintRewardTokenDecimals, uint256 totalUsdt, uint256 totalAmount, uint256 lastDailyReward, uint256 dailyAmountRate, uint256 minAmount, address defaultInvitor, bool pauseJoin, uint256 rewardTokenPrice, uint256 sellRate)
func (_DappHorc *DappHorcCaller) GetBaseInfo(opts *bind.CallOpts) (struct {
	Usdt                    common.Address
	UsdtDecimals            *big.Int
	MintRewardToken         common.Address
	MintRewardTokenDecimals *big.Int
	TotalUsdt               *big.Int
	TotalAmount             *big.Int
	LastDailyReward         *big.Int
	DailyAmountRate         *big.Int
	MinAmount               *big.Int
	DefaultInvitor          common.Address
	PauseJoin               bool
	RewardTokenPrice        *big.Int
	SellRate                *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getBaseInfo")

	outstruct := new(struct {
		Usdt                    common.Address
		UsdtDecimals            *big.Int
		MintRewardToken         common.Address
		MintRewardTokenDecimals *big.Int
		TotalUsdt               *big.Int
		TotalAmount             *big.Int
		LastDailyReward         *big.Int
		DailyAmountRate         *big.Int
		MinAmount               *big.Int
		DefaultInvitor          common.Address
		PauseJoin               bool
		RewardTokenPrice        *big.Int
		SellRate                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Usdt = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.UsdtDecimals = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MintRewardToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MintRewardTokenDecimals = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalUsdt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastDailyReward = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DailyAmountRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MinAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.DefaultInvitor = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.PauseJoin = *abi.ConvertType(out[10], new(bool)).(*bool)
	outstruct.RewardTokenPrice = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.SellRate = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBaseInfo is a free data retrieval call binding the contract method 0x9551ae44.
//
// Solidity: function getBaseInfo() view returns(address usdt, uint256 usdtDecimals, address mintRewardToken, uint256 mintRewardTokenDecimals, uint256 totalUsdt, uint256 totalAmount, uint256 lastDailyReward, uint256 dailyAmountRate, uint256 minAmount, address defaultInvitor, bool pauseJoin, uint256 rewardTokenPrice, uint256 sellRate)
func (_DappHorc *DappHorcSession) GetBaseInfo() (struct {
	Usdt                    common.Address
	UsdtDecimals            *big.Int
	MintRewardToken         common.Address
	MintRewardTokenDecimals *big.Int
	TotalUsdt               *big.Int
	TotalAmount             *big.Int
	LastDailyReward         *big.Int
	DailyAmountRate         *big.Int
	MinAmount               *big.Int
	DefaultInvitor          common.Address
	PauseJoin               bool
	RewardTokenPrice        *big.Int
	SellRate                *big.Int
}, error) {
	return _DappHorc.Contract.GetBaseInfo(&_DappHorc.CallOpts)
}

// GetBaseInfo is a free data retrieval call binding the contract method 0x9551ae44.
//
// Solidity: function getBaseInfo() view returns(address usdt, uint256 usdtDecimals, address mintRewardToken, uint256 mintRewardTokenDecimals, uint256 totalUsdt, uint256 totalAmount, uint256 lastDailyReward, uint256 dailyAmountRate, uint256 minAmount, address defaultInvitor, bool pauseJoin, uint256 rewardTokenPrice, uint256 sellRate)
func (_DappHorc *DappHorcCallerSession) GetBaseInfo() (struct {
	Usdt                    common.Address
	UsdtDecimals            *big.Int
	MintRewardToken         common.Address
	MintRewardTokenDecimals *big.Int
	TotalUsdt               *big.Int
	TotalAmount             *big.Int
	LastDailyReward         *big.Int
	DailyAmountRate         *big.Int
	MinAmount               *big.Int
	DefaultInvitor          common.Address
	PauseJoin               bool
	RewardTokenPrice        *big.Int
	SellRate                *big.Int
}, error) {
	return _DappHorc.Contract.GetBaseInfo(&_DappHorc.CallOpts)
}

// GetBinderLength is a free data retrieval call binding the contract method 0x1b967ad4.
//
// Solidity: function getBinderLength(address account) view returns(uint256)
func (_DappHorc *DappHorcCaller) GetBinderLength(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getBinderLength", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBinderLength is a free data retrieval call binding the contract method 0x1b967ad4.
//
// Solidity: function getBinderLength(address account) view returns(uint256)
func (_DappHorc *DappHorcSession) GetBinderLength(account common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetBinderLength(&_DappHorc.CallOpts, account)
}

// GetBinderLength is a free data retrieval call binding the contract method 0x1b967ad4.
//
// Solidity: function getBinderLength(address account) view returns(uint256)
func (_DappHorc *DappHorcCallerSession) GetBinderLength(account common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetBinderLength(&_DappHorc.CallOpts, account)
}

// GetDueInfo is a free data retrieval call binding the contract method 0x404b6906.
//
// Solidity: function getDueInfo(address token) view returns(uint256, bool)
func (_DappHorc *DappHorcCaller) GetDueInfo(opts *bind.CallOpts, token common.Address) (*big.Int, bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getDueInfo", token)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetDueInfo is a free data retrieval call binding the contract method 0x404b6906.
//
// Solidity: function getDueInfo(address token) view returns(uint256, bool)
func (_DappHorc *DappHorcSession) GetDueInfo(token common.Address) (*big.Int, bool, error) {
	return _DappHorc.Contract.GetDueInfo(&_DappHorc.CallOpts, token)
}

// GetDueInfo is a free data retrieval call binding the contract method 0x404b6906.
//
// Solidity: function getDueInfo(address token) view returns(uint256, bool)
func (_DappHorc *DappHorcCallerSession) GetDueInfo(token common.Address) (*big.Int, bool, error) {
	return _DappHorc.Contract.GetDueInfo(&_DappHorc.CallOpts, token)
}

// GetInvite is a free data retrieval call binding the contract method 0x61c65e2c.
//
// Solidity: function getInvite(address addr) view returns(address)
func (_DappHorc *DappHorcCaller) GetInvite(opts *bind.CallOpts, addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getInvite", addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInvite is a free data retrieval call binding the contract method 0x61c65e2c.
//
// Solidity: function getInvite(address addr) view returns(address)
func (_DappHorc *DappHorcSession) GetInvite(addr common.Address) (common.Address, error) {
	return _DappHorc.Contract.GetInvite(&_DappHorc.CallOpts, addr)
}

// GetInvite is a free data retrieval call binding the contract method 0x61c65e2c.
//
// Solidity: function getInvite(address addr) view returns(address)
func (_DappHorc *DappHorcCallerSession) GetInvite(addr common.Address) (common.Address, error) {
	return _DappHorc.Contract.GetInvite(&_DappHorc.CallOpts, addr)
}

// GetJoinTokenAmount is a free data retrieval call binding the contract method 0x480ba15f.
//
// Solidity: function getJoinTokenAmount(address token, uint256 usdtAmount) view returns(uint256)
func (_DappHorc *DappHorcCaller) GetJoinTokenAmount(opts *bind.CallOpts, token common.Address, usdtAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getJoinTokenAmount", token, usdtAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJoinTokenAmount is a free data retrieval call binding the contract method 0x480ba15f.
//
// Solidity: function getJoinTokenAmount(address token, uint256 usdtAmount) view returns(uint256)
func (_DappHorc *DappHorcSession) GetJoinTokenAmount(token common.Address, usdtAmount *big.Int) (*big.Int, error) {
	return _DappHorc.Contract.GetJoinTokenAmount(&_DappHorc.CallOpts, token, usdtAmount)
}

// GetJoinTokenAmount is a free data retrieval call binding the contract method 0x480ba15f.
//
// Solidity: function getJoinTokenAmount(address token, uint256 usdtAmount) view returns(uint256)
func (_DappHorc *DappHorcCallerSession) GetJoinTokenAmount(token common.Address, usdtAmount *big.Int) (*big.Int, error) {
	return _DappHorc.Contract.GetJoinTokenAmount(&_DappHorc.CallOpts, token, usdtAmount)
}

// GetJoinTokens is a free data retrieval call binding the contract method 0xeaf9e61a.
//
// Solidity: function getJoinTokens() view returns(address[] tokens, uint256[] tokenDecimals, string[] tokenSymbols, uint256[] poolUsdts, uint256[] poolTokens)
func (_DappHorc *DappHorcCaller) GetJoinTokens(opts *bind.CallOpts) (struct {
	Tokens        []common.Address
	TokenDecimals []*big.Int
	TokenSymbols  []string
	PoolUsdts     []*big.Int
	PoolTokens    []*big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getJoinTokens")

	outstruct := new(struct {
		Tokens        []common.Address
		TokenDecimals []*big.Int
		TokenSymbols  []string
		PoolUsdts     []*big.Int
		PoolTokens    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.TokenDecimals = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.TokenSymbols = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.PoolUsdts = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.PoolTokens = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetJoinTokens is a free data retrieval call binding the contract method 0xeaf9e61a.
//
// Solidity: function getJoinTokens() view returns(address[] tokens, uint256[] tokenDecimals, string[] tokenSymbols, uint256[] poolUsdts, uint256[] poolTokens)
func (_DappHorc *DappHorcSession) GetJoinTokens() (struct {
	Tokens        []common.Address
	TokenDecimals []*big.Int
	TokenSymbols  []string
	PoolUsdts     []*big.Int
	PoolTokens    []*big.Int
}, error) {
	return _DappHorc.Contract.GetJoinTokens(&_DappHorc.CallOpts)
}

// GetJoinTokens is a free data retrieval call binding the contract method 0xeaf9e61a.
//
// Solidity: function getJoinTokens() view returns(address[] tokens, uint256[] tokenDecimals, string[] tokenSymbols, uint256[] poolUsdts, uint256[] poolTokens)
func (_DappHorc *DappHorcCallerSession) GetJoinTokens() (struct {
	Tokens        []common.Address
	TokenDecimals []*big.Int
	TokenSymbols  []string
	PoolUsdts     []*big.Int
	PoolTokens    []*big.Int
}, error) {
	return _DappHorc.Contract.GetJoinTokens(&_DappHorc.CallOpts)
}

// GetLevel is a free data retrieval call binding the contract method 0x817c8966.
//
// Solidity: function getLevel(address addr) view returns(uint256 level)
func (_DappHorc *DappHorcCaller) GetLevel(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getLevel", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x817c8966.
//
// Solidity: function getLevel(address addr) view returns(uint256 level)
func (_DappHorc *DappHorcSession) GetLevel(addr common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetLevel(&_DappHorc.CallOpts, addr)
}

// GetLevel is a free data retrieval call binding the contract method 0x817c8966.
//
// Solidity: function getLevel(address addr) view returns(uint256 level)
func (_DappHorc *DappHorcCallerSession) GetLevel(addr common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetLevel(&_DappHorc.CallOpts, addr)
}

// GetPendingMintReward is a free data retrieval call binding the contract method 0x8a56d306.
//
// Solidity: function getPendingMintReward(address account) view returns(uint256 reward)
func (_DappHorc *DappHorcCaller) GetPendingMintReward(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getPendingMintReward", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingMintReward is a free data retrieval call binding the contract method 0x8a56d306.
//
// Solidity: function getPendingMintReward(address account) view returns(uint256 reward)
func (_DappHorc *DappHorcSession) GetPendingMintReward(account common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetPendingMintReward(&_DappHorc.CallOpts, account)
}

// GetPendingMintReward is a free data retrieval call binding the contract method 0x8a56d306.
//
// Solidity: function getPendingMintReward(address account) view returns(uint256 reward)
func (_DappHorc *DappHorcCallerSession) GetPendingMintReward(account common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetPendingMintReward(&_DappHorc.CallOpts, account)
}

// GetSelfRate is a free data retrieval call binding the contract method 0x24fd8bcb.
//
// Solidity: function getSelfRate(uint256 userAmount, uint256 amount) view returns(uint256 pow)
func (_DappHorc *DappHorcCaller) GetSelfRate(opts *bind.CallOpts, userAmount *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getSelfRate", userAmount, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfRate is a free data retrieval call binding the contract method 0x24fd8bcb.
//
// Solidity: function getSelfRate(uint256 userAmount, uint256 amount) view returns(uint256 pow)
func (_DappHorc *DappHorcSession) GetSelfRate(userAmount *big.Int, amount *big.Int) (*big.Int, error) {
	return _DappHorc.Contract.GetSelfRate(&_DappHorc.CallOpts, userAmount, amount)
}

// GetSelfRate is a free data retrieval call binding the contract method 0x24fd8bcb.
//
// Solidity: function getSelfRate(uint256 userAmount, uint256 amount) view returns(uint256 pow)
func (_DappHorc *DappHorcCallerSession) GetSelfRate(userAmount *big.Int, amount *big.Int) (*big.Int, error) {
	return _DappHorc.Contract.GetSelfRate(&_DappHorc.CallOpts, userAmount, amount)
}

// GetSellRate is a free data retrieval call binding the contract method 0x226fd9eb.
//
// Solidity: function getSellRate() view returns(uint256 sellRate)
func (_DappHorc *DappHorcCaller) GetSellRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getSellRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellRate is a free data retrieval call binding the contract method 0x226fd9eb.
//
// Solidity: function getSellRate() view returns(uint256 sellRate)
func (_DappHorc *DappHorcSession) GetSellRate() (*big.Int, error) {
	return _DappHorc.Contract.GetSellRate(&_DappHorc.CallOpts)
}

// GetSellRate is a free data retrieval call binding the contract method 0x226fd9eb.
//
// Solidity: function getSellRate() view returns(uint256 sellRate)
func (_DappHorc *DappHorcCallerSession) GetSellRate() (*big.Int, error) {
	return _DappHorc.Contract.GetSellRate(&_DappHorc.CallOpts)
}

// GetTokenUsdtPrice is a free data retrieval call binding the contract method 0x18836b9d.
//
// Solidity: function getTokenUsdtPrice(address token) view returns(uint256 price)
func (_DappHorc *DappHorcCaller) GetTokenUsdtPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getTokenUsdtPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenUsdtPrice is a free data retrieval call binding the contract method 0x18836b9d.
//
// Solidity: function getTokenUsdtPrice(address token) view returns(uint256 price)
func (_DappHorc *DappHorcSession) GetTokenUsdtPrice(token common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetTokenUsdtPrice(&_DappHorc.CallOpts, token)
}

// GetTokenUsdtPrice is a free data retrieval call binding the contract method 0x18836b9d.
//
// Solidity: function getTokenUsdtPrice(address token) view returns(uint256 price)
func (_DappHorc *DappHorcCallerSession) GetTokenUsdtPrice(token common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetTokenUsdtPrice(&_DappHorc.CallOpts, token)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account) view returns(uint256 amount, uint256 pendingMintReward, uint256 inviteAmount, uint256[] joinTokenBalances, uint256[] joinTokenAllowances, uint256 teamNum)
func (_DappHorc *DappHorcCaller) GetUserInfo(opts *bind.CallOpts, account common.Address) (struct {
	Amount              *big.Int
	PendingMintReward   *big.Int
	InviteAmount        *big.Int
	JoinTokenBalances   []*big.Int
	JoinTokenAllowances []*big.Int
	TeamNum             *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getUserInfo", account)

	outstruct := new(struct {
		Amount              *big.Int
		PendingMintReward   *big.Int
		InviteAmount        *big.Int
		JoinTokenBalances   []*big.Int
		JoinTokenAllowances []*big.Int
		TeamNum             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingMintReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InviteAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.JoinTokenBalances = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.JoinTokenAllowances = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)
	outstruct.TeamNum = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account) view returns(uint256 amount, uint256 pendingMintReward, uint256 inviteAmount, uint256[] joinTokenBalances, uint256[] joinTokenAllowances, uint256 teamNum)
func (_DappHorc *DappHorcSession) GetUserInfo(account common.Address) (struct {
	Amount              *big.Int
	PendingMintReward   *big.Int
	InviteAmount        *big.Int
	JoinTokenBalances   []*big.Int
	JoinTokenAllowances []*big.Int
	TeamNum             *big.Int
}, error) {
	return _DappHorc.Contract.GetUserInfo(&_DappHorc.CallOpts, account)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account) view returns(uint256 amount, uint256 pendingMintReward, uint256 inviteAmount, uint256[] joinTokenBalances, uint256[] joinTokenAllowances, uint256 teamNum)
func (_DappHorc *DappHorcCallerSession) GetUserInfo(account common.Address) (struct {
	Amount              *big.Int
	PendingMintReward   *big.Int
	InviteAmount        *big.Int
	JoinTokenBalances   []*big.Int
	JoinTokenAllowances []*big.Int
	TeamNum             *big.Int
}, error) {
	return _DappHorc.Contract.GetUserInfo(&_DappHorc.CallOpts, account)
}

// LimitUsdt is a free data retrieval call binding the contract method 0xd52e9a9d.
//
// Solidity: function limitUsdt() view returns(uint256)
func (_DappHorc *DappHorcCaller) LimitUsdt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "limitUsdt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUsdt is a free data retrieval call binding the contract method 0xd52e9a9d.
//
// Solidity: function limitUsdt() view returns(uint256)
func (_DappHorc *DappHorcSession) LimitUsdt() (*big.Int, error) {
	return _DappHorc.Contract.LimitUsdt(&_DappHorc.CallOpts)
}

// LimitUsdt is a free data retrieval call binding the contract method 0xd52e9a9d.
//
// Solidity: function limitUsdt() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) LimitUsdt() (*big.Int, error) {
	return _DappHorc.Contract.LimitUsdt(&_DappHorc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappHorc *DappHorcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappHorc *DappHorcSession) Owner() (common.Address, error) {
	return _DappHorc.Contract.Owner(&_DappHorc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappHorc *DappHorcCallerSession) Owner() (common.Address, error) {
	return _DappHorc.Contract.Owner(&_DappHorc.CallOpts)
}

// PowX is a free data retrieval call binding the contract method 0x7f7902de.
//
// Solidity: function powX() view returns(uint256)
func (_DappHorc *DappHorcCaller) PowX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "powX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PowX is a free data retrieval call binding the contract method 0x7f7902de.
//
// Solidity: function powX() view returns(uint256)
func (_DappHorc *DappHorcSession) PowX() (*big.Int, error) {
	return _DappHorc.Contract.PowX(&_DappHorc.CallOpts)
}

// PowX is a free data retrieval call binding the contract method 0x7f7902de.
//
// Solidity: function powX() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) PowX() (*big.Int, error) {
	return _DappHorc.Contract.PowX(&_DappHorc.CallOpts)
}

// SelfLimit1 is a free data retrieval call binding the contract method 0x05effb48.
//
// Solidity: function selfLimit1() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfLimit1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfLimit1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfLimit1 is a free data retrieval call binding the contract method 0x05effb48.
//
// Solidity: function selfLimit1() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfLimit1() (*big.Int, error) {
	return _DappHorc.Contract.SelfLimit1(&_DappHorc.CallOpts)
}

// SelfLimit1 is a free data retrieval call binding the contract method 0x05effb48.
//
// Solidity: function selfLimit1() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfLimit1() (*big.Int, error) {
	return _DappHorc.Contract.SelfLimit1(&_DappHorc.CallOpts)
}

// SelfLimit2 is a free data retrieval call binding the contract method 0x665ad62d.
//
// Solidity: function selfLimit2() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfLimit2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfLimit2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfLimit2 is a free data retrieval call binding the contract method 0x665ad62d.
//
// Solidity: function selfLimit2() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfLimit2() (*big.Int, error) {
	return _DappHorc.Contract.SelfLimit2(&_DappHorc.CallOpts)
}

// SelfLimit2 is a free data retrieval call binding the contract method 0x665ad62d.
//
// Solidity: function selfLimit2() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfLimit2() (*big.Int, error) {
	return _DappHorc.Contract.SelfLimit2(&_DappHorc.CallOpts)
}

// SelfLimit3 is a free data retrieval call binding the contract method 0x45be87e2.
//
// Solidity: function selfLimit3() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfLimit3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfLimit3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfLimit3 is a free data retrieval call binding the contract method 0x45be87e2.
//
// Solidity: function selfLimit3() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfLimit3() (*big.Int, error) {
	return _DappHorc.Contract.SelfLimit3(&_DappHorc.CallOpts)
}

// SelfLimit3 is a free data retrieval call binding the contract method 0x45be87e2.
//
// Solidity: function selfLimit3() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfLimit3() (*big.Int, error) {
	return _DappHorc.Contract.SelfLimit3(&_DappHorc.CallOpts)
}

// SelfRate1 is a free data retrieval call binding the contract method 0xf2323a39.
//
// Solidity: function selfRate1() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfRate1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfRate1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfRate1 is a free data retrieval call binding the contract method 0xf2323a39.
//
// Solidity: function selfRate1() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfRate1() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate1(&_DappHorc.CallOpts)
}

// SelfRate1 is a free data retrieval call binding the contract method 0xf2323a39.
//
// Solidity: function selfRate1() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfRate1() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate1(&_DappHorc.CallOpts)
}

// SelfRate2 is a free data retrieval call binding the contract method 0xfbda34db.
//
// Solidity: function selfRate2() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfRate2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfRate2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfRate2 is a free data retrieval call binding the contract method 0xfbda34db.
//
// Solidity: function selfRate2() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfRate2() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate2(&_DappHorc.CallOpts)
}

// SelfRate2 is a free data retrieval call binding the contract method 0xfbda34db.
//
// Solidity: function selfRate2() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfRate2() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate2(&_DappHorc.CallOpts)
}

// SelfRate3 is a free data retrieval call binding the contract method 0x5e0620fa.
//
// Solidity: function selfRate3() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfRate3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfRate3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfRate3 is a free data retrieval call binding the contract method 0x5e0620fa.
//
// Solidity: function selfRate3() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfRate3() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate3(&_DappHorc.CallOpts)
}

// SelfRate3 is a free data retrieval call binding the contract method 0x5e0620fa.
//
// Solidity: function selfRate3() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfRate3() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate3(&_DappHorc.CallOpts)
}

// SelfRate4 is a free data retrieval call binding the contract method 0x568469ed.
//
// Solidity: function selfRate4() view returns(uint256)
func (_DappHorc *DappHorcCaller) SelfRate4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "selfRate4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfRate4 is a free data retrieval call binding the contract method 0x568469ed.
//
// Solidity: function selfRate4() view returns(uint256)
func (_DappHorc *DappHorcSession) SelfRate4() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate4(&_DappHorc.CallOpts)
}

// SelfRate4 is a free data retrieval call binding the contract method 0x568469ed.
//
// Solidity: function selfRate4() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) SelfRate4() (*big.Int, error) {
	return _DappHorc.Contract.SelfRate4(&_DappHorc.CallOpts)
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() view returns(uint256)
func (_DappHorc *DappHorcCaller) Starttime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "starttime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() view returns(uint256)
func (_DappHorc *DappHorcSession) Starttime() (*big.Int, error) {
	return _DappHorc.Contract.Starttime(&_DappHorc.CallOpts)
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() view returns(uint256)
func (_DappHorc *DappHorcCallerSession) Starttime() (*big.Int, error) {
	return _DappHorc.Contract.Starttime(&_DappHorc.CallOpts)
}

// ViewPoolInfo is a free data retrieval call binding the contract method 0x277c5e5f.
//
// Solidity: function viewPoolInfo() view returns(uint256 totalAmount, uint256 accMintPerShare, uint256 accMintReward, uint256 mintPerSec, uint256 lastMintTime, uint256 totalMintReward)
func (_DappHorc *DappHorcCaller) ViewPoolInfo(opts *bind.CallOpts) (struct {
	TotalAmount     *big.Int
	AccMintPerShare *big.Int
	AccMintReward   *big.Int
	MintPerSec      *big.Int
	LastMintTime    *big.Int
	TotalMintReward *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "viewPoolInfo")

	outstruct := new(struct {
		TotalAmount     *big.Int
		AccMintPerShare *big.Int
		AccMintReward   *big.Int
		MintPerSec      *big.Int
		LastMintTime    *big.Int
		TotalMintReward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccMintPerShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccMintReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MintPerSec = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastMintTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalMintReward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ViewPoolInfo is a free data retrieval call binding the contract method 0x277c5e5f.
//
// Solidity: function viewPoolInfo() view returns(uint256 totalAmount, uint256 accMintPerShare, uint256 accMintReward, uint256 mintPerSec, uint256 lastMintTime, uint256 totalMintReward)
func (_DappHorc *DappHorcSession) ViewPoolInfo() (struct {
	TotalAmount     *big.Int
	AccMintPerShare *big.Int
	AccMintReward   *big.Int
	MintPerSec      *big.Int
	LastMintTime    *big.Int
	TotalMintReward *big.Int
}, error) {
	return _DappHorc.Contract.ViewPoolInfo(&_DappHorc.CallOpts)
}

// ViewPoolInfo is a free data retrieval call binding the contract method 0x277c5e5f.
//
// Solidity: function viewPoolInfo() view returns(uint256 totalAmount, uint256 accMintPerShare, uint256 accMintReward, uint256 mintPerSec, uint256 lastMintTime, uint256 totalMintReward)
func (_DappHorc *DappHorcCallerSession) ViewPoolInfo() (struct {
	TotalAmount     *big.Int
	AccMintPerShare *big.Int
	AccMintReward   *big.Int
	MintPerSec      *big.Int
	LastMintTime    *big.Int
	TotalMintReward *big.Int
}, error) {
	return _DappHorc.Contract.ViewPoolInfo(&_DappHorc.CallOpts)
}

// ViewUserInfo is a free data retrieval call binding the contract method 0xa37d9850.
//
// Solidity: function viewUserInfo(address account) view returns(bool isActive, uint256 amount, uint256 amountUsdt, uint256 calMintReward, uint256 rewardMintDebt, uint256 level, uint256 inviteNum, uint256 inviteAmount, uint256 invite2Amount, uint256 nodeNum, uint256 superNum, uint256 teamPow)
func (_DappHorc *DappHorcCaller) ViewUserInfo(opts *bind.CallOpts, account common.Address) (struct {
	IsActive       bool
	Amount         *big.Int
	AmountUsdt     *big.Int
	CalMintReward  *big.Int
	RewardMintDebt *big.Int
	Level          *big.Int
	InviteNum      *big.Int
	InviteAmount   *big.Int
	Invite2Amount  *big.Int
	NodeNum        *big.Int
	SuperNum       *big.Int
	TeamPow        *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "viewUserInfo", account)

	outstruct := new(struct {
		IsActive       bool
		Amount         *big.Int
		AmountUsdt     *big.Int
		CalMintReward  *big.Int
		RewardMintDebt *big.Int
		Level          *big.Int
		InviteNum      *big.Int
		InviteAmount   *big.Int
		Invite2Amount  *big.Int
		NodeNum        *big.Int
		SuperNum       *big.Int
		TeamPow        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountUsdt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CalMintReward = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RewardMintDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.InviteNum = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.InviteAmount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Invite2Amount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.NodeNum = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.SuperNum = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TeamPow = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ViewUserInfo is a free data retrieval call binding the contract method 0xa37d9850.
//
// Solidity: function viewUserInfo(address account) view returns(bool isActive, uint256 amount, uint256 amountUsdt, uint256 calMintReward, uint256 rewardMintDebt, uint256 level, uint256 inviteNum, uint256 inviteAmount, uint256 invite2Amount, uint256 nodeNum, uint256 superNum, uint256 teamPow)
func (_DappHorc *DappHorcSession) ViewUserInfo(account common.Address) (struct {
	IsActive       bool
	Amount         *big.Int
	AmountUsdt     *big.Int
	CalMintReward  *big.Int
	RewardMintDebt *big.Int
	Level          *big.Int
	InviteNum      *big.Int
	InviteAmount   *big.Int
	Invite2Amount  *big.Int
	NodeNum        *big.Int
	SuperNum       *big.Int
	TeamPow        *big.Int
}, error) {
	return _DappHorc.Contract.ViewUserInfo(&_DappHorc.CallOpts, account)
}

// ViewUserInfo is a free data retrieval call binding the contract method 0xa37d9850.
//
// Solidity: function viewUserInfo(address account) view returns(bool isActive, uint256 amount, uint256 amountUsdt, uint256 calMintReward, uint256 rewardMintDebt, uint256 level, uint256 inviteNum, uint256 inviteAmount, uint256 invite2Amount, uint256 nodeNum, uint256 superNum, uint256 teamPow)
func (_DappHorc *DappHorcCallerSession) ViewUserInfo(account common.Address) (struct {
	IsActive       bool
	Amount         *big.Int
	AmountUsdt     *big.Int
	CalMintReward  *big.Int
	RewardMintDebt *big.Int
	Level          *big.Int
	InviteNum      *big.Int
	InviteAmount   *big.Int
	Invite2Amount  *big.Int
	NodeNum        *big.Int
	SuperNum       *big.Int
	TeamPow        *big.Int
}, error) {
	return _DappHorc.Contract.ViewUserInfo(&_DappHorc.CallOpts, account)
}

// UpdateDailyUpRate is a paid mutator transaction binding the contract method 0xcc603eee.
//
// Solidity: function _updateDailyUpRate() returns()
func (_DappHorc *DappHorcTransactor) UpdateDailyUpRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "_updateDailyUpRate")
}

// UpdateDailyUpRate is a paid mutator transaction binding the contract method 0xcc603eee.
//
// Solidity: function _updateDailyUpRate() returns()
func (_DappHorc *DappHorcSession) UpdateDailyUpRate() (*types.Transaction, error) {
	return _DappHorc.Contract.UpdateDailyUpRate(&_DappHorc.TransactOpts)
}

// UpdateDailyUpRate is a paid mutator transaction binding the contract method 0xcc603eee.
//
// Solidity: function _updateDailyUpRate() returns()
func (_DappHorc *DappHorcTransactorSession) UpdateDailyUpRate() (*types.Transaction, error) {
	return _DappHorc.Contract.UpdateDailyUpRate(&_DappHorc.TransactOpts)
}

// AddExtendList is a paid mutator transaction binding the contract method 0x2b7f7307.
//
// Solidity: function addExtendList(address addr) returns()
func (_DappHorc *DappHorcTransactor) AddExtendList(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "addExtendList", addr)
}

// AddExtendList is a paid mutator transaction binding the contract method 0x2b7f7307.
//
// Solidity: function addExtendList(address addr) returns()
func (_DappHorc *DappHorcSession) AddExtendList(addr common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.AddExtendList(&_DappHorc.TransactOpts, addr)
}

// AddExtendList is a paid mutator transaction binding the contract method 0x2b7f7307.
//
// Solidity: function addExtendList(address addr) returns()
func (_DappHorc *DappHorcTransactorSession) AddExtendList(addr common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.AddExtendList(&_DappHorc.TransactOpts, addr)
}

// AddJoinToken is a paid mutator transaction binding the contract method 0xe049f010.
//
// Solidity: function addJoinToken(address token, bool isMain) returns()
func (_DappHorc *DappHorcTransactor) AddJoinToken(opts *bind.TransactOpts, token common.Address, isMain bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "addJoinToken", token, isMain)
}

// AddJoinToken is a paid mutator transaction binding the contract method 0xe049f010.
//
// Solidity: function addJoinToken(address token, bool isMain) returns()
func (_DappHorc *DappHorcSession) AddJoinToken(token common.Address, isMain bool) (*types.Transaction, error) {
	return _DappHorc.Contract.AddJoinToken(&_DappHorc.TransactOpts, token, isMain)
}

// AddJoinToken is a paid mutator transaction binding the contract method 0xe049f010.
//
// Solidity: function addJoinToken(address token, bool isMain) returns()
func (_DappHorc *DappHorcTransactorSession) AddJoinToken(token common.Address, isMain bool) (*types.Transaction, error) {
	return _DappHorc.Contract.AddJoinToken(&_DappHorc.TransactOpts, token, isMain)
}

// AddMintAmount is a paid mutator transaction binding the contract method 0xe9f22169.
//
// Solidity: function addMintAmount(address account, uint256 amount) returns()
func (_DappHorc *DappHorcTransactor) AddMintAmount(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "addMintAmount", account, amount)
}

// AddMintAmount is a paid mutator transaction binding the contract method 0xe9f22169.
//
// Solidity: function addMintAmount(address account, uint256 amount) returns()
func (_DappHorc *DappHorcSession) AddMintAmount(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.AddMintAmount(&_DappHorc.TransactOpts, account, amount)
}

// AddMintAmount is a paid mutator transaction binding the contract method 0xe9f22169.
//
// Solidity: function addMintAmount(address account, uint256 amount) returns()
func (_DappHorc *DappHorcTransactorSession) AddMintAmount(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.AddMintAmount(&_DappHorc.TransactOpts, account, amount)
}

// AddUserAmount is a paid mutator transaction binding the contract method 0x0600ffc1.
//
// Solidity: function addUserAmount(address account, uint256 amount, bool calInvite) returns()
func (_DappHorc *DappHorcTransactor) AddUserAmount(opts *bind.TransactOpts, account common.Address, amount *big.Int, calInvite bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "addUserAmount", account, amount, calInvite)
}

// AddUserAmount is a paid mutator transaction binding the contract method 0x0600ffc1.
//
// Solidity: function addUserAmount(address account, uint256 amount, bool calInvite) returns()
func (_DappHorc *DappHorcSession) AddUserAmount(account common.Address, amount *big.Int, calInvite bool) (*types.Transaction, error) {
	return _DappHorc.Contract.AddUserAmount(&_DappHorc.TransactOpts, account, amount, calInvite)
}

// AddUserAmount is a paid mutator transaction binding the contract method 0x0600ffc1.
//
// Solidity: function addUserAmount(address account, uint256 amount, bool calInvite) returns()
func (_DappHorc *DappHorcTransactorSession) AddUserAmount(account common.Address, amount *big.Int, calInvite bool) (*types.Transaction, error) {
	return _DappHorc.Contract.AddUserAmount(&_DappHorc.TransactOpts, account, amount, calInvite)
}

// BindInvitor is a paid mutator transaction binding the contract method 0xf12e6eb7.
//
// Solidity: function bindInvitor(address account, address invitor) returns()
func (_DappHorc *DappHorcTransactor) BindInvitor(opts *bind.TransactOpts, account common.Address, invitor common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "bindInvitor", account, invitor)
}

// BindInvitor is a paid mutator transaction binding the contract method 0xf12e6eb7.
//
// Solidity: function bindInvitor(address account, address invitor) returns()
func (_DappHorc *DappHorcSession) BindInvitor(account common.Address, invitor common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.BindInvitor(&_DappHorc.TransactOpts, account, invitor)
}

// BindInvitor is a paid mutator transaction binding the contract method 0xf12e6eb7.
//
// Solidity: function bindInvitor(address account, address invitor) returns()
func (_DappHorc *DappHorcTransactorSession) BindInvitor(account common.Address, invitor common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.BindInvitor(&_DappHorc.TransactOpts, account, invitor)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_DappHorc *DappHorcTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_DappHorc *DappHorcSession) Claim() (*types.Transaction, error) {
	return _DappHorc.Contract.Claim(&_DappHorc.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_DappHorc *DappHorcTransactorSession) Claim() (*types.Transaction, error) {
	return _DappHorc.Contract.Claim(&_DappHorc.TransactOpts)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0xaa8b38d9.
//
// Solidity: function claimBalance(address to, uint256 amount) returns()
func (_DappHorc *DappHorcTransactor) ClaimBalance(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "claimBalance", to, amount)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0xaa8b38d9.
//
// Solidity: function claimBalance(address to, uint256 amount) returns()
func (_DappHorc *DappHorcSession) ClaimBalance(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ClaimBalance(&_DappHorc.TransactOpts, to, amount)
}

// ClaimBalance is a paid mutator transaction binding the contract method 0xaa8b38d9.
//
// Solidity: function claimBalance(address to, uint256 amount) returns()
func (_DappHorc *DappHorcTransactorSession) ClaimBalance(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ClaimBalance(&_DappHorc.TransactOpts, to, amount)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x125bfb66.
//
// Solidity: function claimToken(address token, address to, uint256 amount) returns()
func (_DappHorc *DappHorcTransactor) ClaimToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "claimToken", token, to, amount)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x125bfb66.
//
// Solidity: function claimToken(address token, address to, uint256 amount) returns()
func (_DappHorc *DappHorcSession) ClaimToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ClaimToken(&_DappHorc.TransactOpts, token, to, amount)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x125bfb66.
//
// Solidity: function claimToken(address token, address to, uint256 amount) returns()
func (_DappHorc *DappHorcTransactorSession) ClaimToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ClaimToken(&_DappHorc.TransactOpts, token, to, amount)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_DappHorc *DappHorcTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_DappHorc *DappHorcSession) Close() (*types.Transaction, error) {
	return _DappHorc.Contract.Close(&_DappHorc.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_DappHorc *DappHorcTransactorSession) Close() (*types.Transaction, error) {
	return _DappHorc.Contract.Close(&_DappHorc.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xfad3cc4b.
//
// Solidity: function deposit(uint256 i, uint256 amount, uint256 maxTokenAmount, address invitor) returns()
func (_DappHorc *DappHorcTransactor) Deposit(opts *bind.TransactOpts, i *big.Int, amount *big.Int, maxTokenAmount *big.Int, invitor common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "deposit", i, amount, maxTokenAmount, invitor)
}

// Deposit is a paid mutator transaction binding the contract method 0xfad3cc4b.
//
// Solidity: function deposit(uint256 i, uint256 amount, uint256 maxTokenAmount, address invitor) returns()
func (_DappHorc *DappHorcSession) Deposit(i *big.Int, amount *big.Int, maxTokenAmount *big.Int, invitor common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.Deposit(&_DappHorc.TransactOpts, i, amount, maxTokenAmount, invitor)
}

// Deposit is a paid mutator transaction binding the contract method 0xfad3cc4b.
//
// Solidity: function deposit(uint256 i, uint256 amount, uint256 maxTokenAmount, address invitor) returns()
func (_DappHorc *DappHorcTransactorSession) Deposit(i *big.Int, amount *big.Int, maxTokenAmount *big.Int, invitor common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.Deposit(&_DappHorc.TransactOpts, i, amount, maxTokenAmount, invitor)
}

// ModUser is a paid mutator transaction binding the contract method 0xab038827.
//
// Solidity: function modUser(address account, uint256 level, uint256 nodeNum, uint256 superNum, uint256 inviteAmount, uint256 invite2Amount, uint256 teamPow) returns()
func (_DappHorc *DappHorcTransactor) ModUser(opts *bind.TransactOpts, account common.Address, level *big.Int, nodeNum *big.Int, superNum *big.Int, inviteAmount *big.Int, invite2Amount *big.Int, teamPow *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "modUser", account, level, nodeNum, superNum, inviteAmount, invite2Amount, teamPow)
}

// ModUser is a paid mutator transaction binding the contract method 0xab038827.
//
// Solidity: function modUser(address account, uint256 level, uint256 nodeNum, uint256 superNum, uint256 inviteAmount, uint256 invite2Amount, uint256 teamPow) returns()
func (_DappHorc *DappHorcSession) ModUser(account common.Address, level *big.Int, nodeNum *big.Int, superNum *big.Int, inviteAmount *big.Int, invite2Amount *big.Int, teamPow *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ModUser(&_DappHorc.TransactOpts, account, level, nodeNum, superNum, inviteAmount, invite2Amount, teamPow)
}

// ModUser is a paid mutator transaction binding the contract method 0xab038827.
//
// Solidity: function modUser(address account, uint256 level, uint256 nodeNum, uint256 superNum, uint256 inviteAmount, uint256 invite2Amount, uint256 teamPow) returns()
func (_DappHorc *DappHorcTransactorSession) ModUser(account common.Address, level *big.Int, nodeNum *big.Int, superNum *big.Int, inviteAmount *big.Int, invite2Amount *big.Int, teamPow *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ModUser(&_DappHorc.TransactOpts, account, level, nodeNum, superNum, inviteAmount, invite2Amount, teamPow)
}

// ModselfLimit is a paid mutator transaction binding the contract method 0xeb87a540.
//
// Solidity: function modselfLimit(uint256 num1, uint256 num2, uint256 num3) returns()
func (_DappHorc *DappHorcTransactor) ModselfLimit(opts *bind.TransactOpts, num1 *big.Int, num2 *big.Int, num3 *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "modselfLimit", num1, num2, num3)
}

// ModselfLimit is a paid mutator transaction binding the contract method 0xeb87a540.
//
// Solidity: function modselfLimit(uint256 num1, uint256 num2, uint256 num3) returns()
func (_DappHorc *DappHorcSession) ModselfLimit(num1 *big.Int, num2 *big.Int, num3 *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ModselfLimit(&_DappHorc.TransactOpts, num1, num2, num3)
}

// ModselfLimit is a paid mutator transaction binding the contract method 0xeb87a540.
//
// Solidity: function modselfLimit(uint256 num1, uint256 num2, uint256 num3) returns()
func (_DappHorc *DappHorcTransactorSession) ModselfLimit(num1 *big.Int, num2 *big.Int, num3 *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ModselfLimit(&_DappHorc.TransactOpts, num1, num2, num3)
}

// ModselfRate is a paid mutator transaction binding the contract method 0x8723a2b2.
//
// Solidity: function modselfRate(uint256 num1, uint256 num2, uint256 num3, uint256 num4) returns()
func (_DappHorc *DappHorcTransactor) ModselfRate(opts *bind.TransactOpts, num1 *big.Int, num2 *big.Int, num3 *big.Int, num4 *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "modselfRate", num1, num2, num3, num4)
}

// ModselfRate is a paid mutator transaction binding the contract method 0x8723a2b2.
//
// Solidity: function modselfRate(uint256 num1, uint256 num2, uint256 num3, uint256 num4) returns()
func (_DappHorc *DappHorcSession) ModselfRate(num1 *big.Int, num2 *big.Int, num3 *big.Int, num4 *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ModselfRate(&_DappHorc.TransactOpts, num1, num2, num3, num4)
}

// ModselfRate is a paid mutator transaction binding the contract method 0x8723a2b2.
//
// Solidity: function modselfRate(uint256 num1, uint256 num2, uint256 num3, uint256 num4) returns()
func (_DappHorc *DappHorcTransactorSession) ModselfRate(num1 *big.Int, num2 *big.Int, num3 *big.Int, num4 *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.ModselfRate(&_DappHorc.TransactOpts, num1, num2, num3, num4)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_DappHorc *DappHorcTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_DappHorc *DappHorcSession) Open() (*types.Transaction, error) {
	return _DappHorc.Contract.Open(&_DappHorc.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_DappHorc *DappHorcTransactorSession) Open() (*types.Transaction, error) {
	return _DappHorc.Contract.Open(&_DappHorc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappHorc *DappHorcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappHorc *DappHorcSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappHorc.Contract.RenounceOwnership(&_DappHorc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappHorc *DappHorcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappHorc.Contract.RenounceOwnership(&_DappHorc.TransactOpts)
}

// SetAmountDailyUp is a paid mutator transaction binding the contract method 0xfd3cf885.
//
// Solidity: function setAmountDailyUp(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetAmountDailyUp(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setAmountDailyUp", r)
}

// SetAmountDailyUp is a paid mutator transaction binding the contract method 0xfd3cf885.
//
// Solidity: function setAmountDailyUp(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetAmountDailyUp(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetAmountDailyUp(&_DappHorc.TransactOpts, r)
}

// SetAmountDailyUp is a paid mutator transaction binding the contract method 0xfd3cf885.
//
// Solidity: function setAmountDailyUp(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetAmountDailyUp(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetAmountDailyUp(&_DappHorc.TransactOpts, r)
}

// SetAutoCalRewardPerDay is a paid mutator transaction binding the contract method 0x167ad857.
//
// Solidity: function setAutoCalRewardPerDay() returns()
func (_DappHorc *DappHorcTransactor) SetAutoCalRewardPerDay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setAutoCalRewardPerDay")
}

// SetAutoCalRewardPerDay is a paid mutator transaction binding the contract method 0x167ad857.
//
// Solidity: function setAutoCalRewardPerDay() returns()
func (_DappHorc *DappHorcSession) SetAutoCalRewardPerDay() (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalRewardPerDay(&_DappHorc.TransactOpts)
}

// SetAutoCalRewardPerDay is a paid mutator transaction binding the contract method 0x167ad857.
//
// Solidity: function setAutoCalRewardPerDay() returns()
func (_DappHorc *DappHorcTransactorSession) SetAutoCalRewardPerDay() (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalRewardPerDay(&_DappHorc.TransactOpts)
}

// SetAutoCalSellRate is a paid mutator transaction binding the contract method 0xc63250f7.
//
// Solidity: function setAutoCalSellRate() returns()
func (_DappHorc *DappHorcTransactor) SetAutoCalSellRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setAutoCalSellRate")
}

// SetAutoCalSellRate is a paid mutator transaction binding the contract method 0xc63250f7.
//
// Solidity: function setAutoCalSellRate() returns()
func (_DappHorc *DappHorcSession) SetAutoCalSellRate() (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalSellRate(&_DappHorc.TransactOpts)
}

// SetAutoCalSellRate is a paid mutator transaction binding the contract method 0xc63250f7.
//
// Solidity: function setAutoCalSellRate() returns()
func (_DappHorc *DappHorcTransactorSession) SetAutoCalSellRate() (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalSellRate(&_DappHorc.TransactOpts)
}

// SetBurnReceiver is a paid mutator transaction binding the contract method 0x66270efd.
//
// Solidity: function setBurnReceiver(address a) returns()
func (_DappHorc *DappHorcTransactor) SetBurnReceiver(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setBurnReceiver", a)
}

// SetBurnReceiver is a paid mutator transaction binding the contract method 0x66270efd.
//
// Solidity: function setBurnReceiver(address a) returns()
func (_DappHorc *DappHorcSession) SetBurnReceiver(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetBurnReceiver(&_DappHorc.TransactOpts, a)
}

// SetBurnReceiver is a paid mutator transaction binding the contract method 0x66270efd.
//
// Solidity: function setBurnReceiver(address a) returns()
func (_DappHorc *DappHorcTransactorSession) SetBurnReceiver(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetBurnReceiver(&_DappHorc.TransactOpts, a)
}

// SetDefaultInvitor is a paid mutator transaction binding the contract method 0xb75265d3.
//
// Solidity: function setDefaultInvitor(address adr) returns()
func (_DappHorc *DappHorcTransactor) SetDefaultInvitor(opts *bind.TransactOpts, adr common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setDefaultInvitor", adr)
}

// SetDefaultInvitor is a paid mutator transaction binding the contract method 0xb75265d3.
//
// Solidity: function setDefaultInvitor(address adr) returns()
func (_DappHorc *DappHorcSession) SetDefaultInvitor(adr common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetDefaultInvitor(&_DappHorc.TransactOpts, adr)
}

// SetDefaultInvitor is a paid mutator transaction binding the contract method 0xb75265d3.
//
// Solidity: function setDefaultInvitor(address adr) returns()
func (_DappHorc *DappHorcTransactorSession) SetDefaultInvitor(adr common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetDefaultInvitor(&_DappHorc.TransactOpts, adr)
}

// SetFundAddress is a paid mutator transaction binding the contract method 0x0e7a71fe.
//
// Solidity: function setFundAddress(address a, address a2, address a3) returns()
func (_DappHorc *DappHorcTransactor) SetFundAddress(opts *bind.TransactOpts, a common.Address, a2 common.Address, a3 common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setFundAddress", a, a2, a3)
}

// SetFundAddress is a paid mutator transaction binding the contract method 0x0e7a71fe.
//
// Solidity: function setFundAddress(address a, address a2, address a3) returns()
func (_DappHorc *DappHorcSession) SetFundAddress(a common.Address, a2 common.Address, a3 common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetFundAddress(&_DappHorc.TransactOpts, a, a2, a3)
}

// SetFundAddress is a paid mutator transaction binding the contract method 0x0e7a71fe.
//
// Solidity: function setFundAddress(address a, address a2, address a3) returns()
func (_DappHorc *DappHorcTransactorSession) SetFundAddress(a common.Address, a2 common.Address, a3 common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetFundAddress(&_DappHorc.TransactOpts, a, a2, a3)
}

// SetFundAddress0 is a paid mutator transaction binding the contract method 0x85dc3004.
//
// Solidity: function setFundAddress(address a3) returns()
func (_DappHorc *DappHorcTransactor) SetFundAddress0(opts *bind.TransactOpts, a3 common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setFundAddress0", a3)
}

// SetFundAddress0 is a paid mutator transaction binding the contract method 0x85dc3004.
//
// Solidity: function setFundAddress(address a3) returns()
func (_DappHorc *DappHorcSession) SetFundAddress0(a3 common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetFundAddress0(&_DappHorc.TransactOpts, a3)
}

// SetFundAddress0 is a paid mutator transaction binding the contract method 0x85dc3004.
//
// Solidity: function setFundAddress(address a3) returns()
func (_DappHorc *DappHorcTransactorSession) SetFundAddress0(a3 common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetFundAddress0(&_DappHorc.TransactOpts, a3)
}

// SetHighPrice is a paid mutator transaction binding the contract method 0xc03d7c94.
//
// Solidity: function setHighPrice(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetHighPrice(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setHighPrice", r)
}

// SetHighPrice is a paid mutator transaction binding the contract method 0xc03d7c94.
//
// Solidity: function setHighPrice(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetHighPrice(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetHighPrice(&_DappHorc.TransactOpts, r)
}

// SetHighPrice is a paid mutator transaction binding the contract method 0xc03d7c94.
//
// Solidity: function setHighPrice(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetHighPrice(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetHighPrice(&_DappHorc.TransactOpts, r)
}

// SetHighPriceSellRate is a paid mutator transaction binding the contract method 0x5be02c9a.
//
// Solidity: function setHighPriceSellRate(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetHighPriceSellRate(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setHighPriceSellRate", r)
}

// SetHighPriceSellRate is a paid mutator transaction binding the contract method 0x5be02c9a.
//
// Solidity: function setHighPriceSellRate(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetHighPriceSellRate(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetHighPriceSellRate(&_DappHorc.TransactOpts, r)
}

// SetHighPriceSellRate is a paid mutator transaction binding the contract method 0x5be02c9a.
//
// Solidity: function setHighPriceSellRate(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetHighPriceSellRate(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetHighPriceSellRate(&_DappHorc.TransactOpts, r)
}

// SetInviteFee is a paid mutator transaction binding the contract method 0x5ca8e3f2.
//
// Solidity: function setInviteFee(uint256 i, uint256 fee) returns()
func (_DappHorc *DappHorcTransactor) SetInviteFee(opts *bind.TransactOpts, i *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setInviteFee", i, fee)
}

// SetInviteFee is a paid mutator transaction binding the contract method 0x5ca8e3f2.
//
// Solidity: function setInviteFee(uint256 i, uint256 fee) returns()
func (_DappHorc *DappHorcSession) SetInviteFee(i *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetInviteFee(&_DappHorc.TransactOpts, i, fee)
}

// SetInviteFee is a paid mutator transaction binding the contract method 0x5ca8e3f2.
//
// Solidity: function setInviteFee(uint256 i, uint256 fee) returns()
func (_DappHorc *DappHorcTransactorSession) SetInviteFee(i *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetInviteFee(&_DappHorc.TransactOpts, i, fee)
}

// SetJoinToken is a paid mutator transaction binding the contract method 0x138c9ddd.
//
// Solidity: function setJoinToken(uint256 i, address addr) returns()
func (_DappHorc *DappHorcTransactor) SetJoinToken(opts *bind.TransactOpts, i *big.Int, addr common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setJoinToken", i, addr)
}

// SetJoinToken is a paid mutator transaction binding the contract method 0x138c9ddd.
//
// Solidity: function setJoinToken(uint256 i, address addr) returns()
func (_DappHorc *DappHorcSession) SetJoinToken(i *big.Int, addr common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinToken(&_DappHorc.TransactOpts, i, addr)
}

// SetJoinToken is a paid mutator transaction binding the contract method 0x138c9ddd.
//
// Solidity: function setJoinToken(uint256 i, address addr) returns()
func (_DappHorc *DappHorcTransactorSession) SetJoinToken(i *big.Int, addr common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinToken(&_DappHorc.TransactOpts, i, addr)
}

// SetJoinToken0 is a paid mutator transaction binding the contract method 0xa9f0fb42.
//
// Solidity: function setJoinToken(uint256 i, address token, bool isMain) returns()
func (_DappHorc *DappHorcTransactor) SetJoinToken0(opts *bind.TransactOpts, i *big.Int, token common.Address, isMain bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setJoinToken0", i, token, isMain)
}

// SetJoinToken0 is a paid mutator transaction binding the contract method 0xa9f0fb42.
//
// Solidity: function setJoinToken(uint256 i, address token, bool isMain) returns()
func (_DappHorc *DappHorcSession) SetJoinToken0(i *big.Int, token common.Address, isMain bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinToken0(&_DappHorc.TransactOpts, i, token, isMain)
}

// SetJoinToken0 is a paid mutator transaction binding the contract method 0xa9f0fb42.
//
// Solidity: function setJoinToken(uint256 i, address token, bool isMain) returns()
func (_DappHorc *DappHorcTransactorSession) SetJoinToken0(i *big.Int, token common.Address, isMain bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinToken0(&_DappHorc.TransactOpts, i, token, isMain)
}

// SetJoinTokens is a paid mutator transaction binding the contract method 0x5c64e897.
//
// Solidity: function setJoinTokens(address[] tokens, bool[] isMain) returns()
func (_DappHorc *DappHorcTransactor) SetJoinTokens(opts *bind.TransactOpts, tokens []common.Address, isMain []bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setJoinTokens", tokens, isMain)
}

// SetJoinTokens is a paid mutator transaction binding the contract method 0x5c64e897.
//
// Solidity: function setJoinTokens(address[] tokens, bool[] isMain) returns()
func (_DappHorc *DappHorcSession) SetJoinTokens(tokens []common.Address, isMain []bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinTokens(&_DappHorc.TransactOpts, tokens, isMain)
}

// SetJoinTokens is a paid mutator transaction binding the contract method 0x5c64e897.
//
// Solidity: function setJoinTokens(address[] tokens, bool[] isMain) returns()
func (_DappHorc *DappHorcTransactorSession) SetJoinTokens(tokens []common.Address, isMain []bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinTokens(&_DappHorc.TransactOpts, tokens, isMain)
}

// SetLastAmountRate is a paid mutator transaction binding the contract method 0x2e0dd9c7.
//
// Solidity: function setLastAmountRate(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetLastAmountRate(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setLastAmountRate", r)
}

// SetLastAmountRate is a paid mutator transaction binding the contract method 0x2e0dd9c7.
//
// Solidity: function setLastAmountRate(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetLastAmountRate(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetLastAmountRate(&_DappHorc.TransactOpts, r)
}

// SetLastAmountRate is a paid mutator transaction binding the contract method 0x2e0dd9c7.
//
// Solidity: function setLastAmountRate(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetLastAmountRate(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetLastAmountRate(&_DappHorc.TransactOpts, r)
}

// SetLastDailyUpTime is a paid mutator transaction binding the contract method 0x029a0c9d.
//
// Solidity: function setLastDailyUpTime(uint256 t) returns()
func (_DappHorc *DappHorcTransactor) SetLastDailyUpTime(opts *bind.TransactOpts, t *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setLastDailyUpTime", t)
}

// SetLastDailyUpTime is a paid mutator transaction binding the contract method 0x029a0c9d.
//
// Solidity: function setLastDailyUpTime(uint256 t) returns()
func (_DappHorc *DappHorcSession) SetLastDailyUpTime(t *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetLastDailyUpTime(&_DappHorc.TransactOpts, t)
}

// SetLastDailyUpTime is a paid mutator transaction binding the contract method 0x029a0c9d.
//
// Solidity: function setLastDailyUpTime(uint256 t) returns()
func (_DappHorc *DappHorcTransactorSession) SetLastDailyUpTime(t *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetLastDailyUpTime(&_DappHorc.TransactOpts, t)
}

// SetNormalRewardPerDay is a paid mutator transaction binding the contract method 0x4b3f51c4.
//
// Solidity: function setNormalRewardPerDay(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetNormalRewardPerDay(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setNormalRewardPerDay", r)
}

// SetNormalRewardPerDay is a paid mutator transaction binding the contract method 0x4b3f51c4.
//
// Solidity: function setNormalRewardPerDay(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetNormalRewardPerDay(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetNormalRewardPerDay(&_DappHorc.TransactOpts, r)
}

// SetNormalRewardPerDay is a paid mutator transaction binding the contract method 0x4b3f51c4.
//
// Solidity: function setNormalRewardPerDay(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetNormalRewardPerDay(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetNormalRewardPerDay(&_DappHorc.TransactOpts, r)
}

// SetRewardPerDay is a paid mutator transaction binding the contract method 0x90c35a93.
//
// Solidity: function setRewardPerDay(uint256 reward) returns()
func (_DappHorc *DappHorcTransactor) SetRewardPerDay(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setRewardPerDay", reward)
}

// SetRewardPerDay is a paid mutator transaction binding the contract method 0x90c35a93.
//
// Solidity: function setRewardPerDay(uint256 reward) returns()
func (_DappHorc *DappHorcSession) SetRewardPerDay(reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetRewardPerDay(&_DappHorc.TransactOpts, reward)
}

// SetRewardPerDay is a paid mutator transaction binding the contract method 0x90c35a93.
//
// Solidity: function setRewardPerDay(uint256 reward) returns()
func (_DappHorc *DappHorcTransactorSession) SetRewardPerDay(reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetRewardPerDay(&_DappHorc.TransactOpts, reward)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address a) returns()
func (_DappHorc *DappHorcTransactor) SetRewardToken(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setRewardToken", a)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address a) returns()
func (_DappHorc *DappHorcSession) SetRewardToken(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetRewardToken(&_DappHorc.TransactOpts, a)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address a) returns()
func (_DappHorc *DappHorcTransactorSession) SetRewardToken(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetRewardToken(&_DappHorc.TransactOpts, a)
}

// SetSellRate is a paid mutator transaction binding the contract method 0x8e0b017d.
//
// Solidity: function setSellRate(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetSellRate(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setSellRate", r)
}

// SetSellRate is a paid mutator transaction binding the contract method 0x8e0b017d.
//
// Solidity: function setSellRate(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetSellRate(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetSellRate(&_DappHorc.TransactOpts, r)
}

// SetSellRate is a paid mutator transaction binding the contract method 0x8e0b017d.
//
// Solidity: function setSellRate(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetSellRate(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetSellRate(&_DappHorc.TransactOpts, r)
}

// SetTotalMintReward is a paid mutator transaction binding the contract method 0x02f6b29e.
//
// Solidity: function setTotalMintReward(uint256 reward) returns()
func (_DappHorc *DappHorcTransactor) SetTotalMintReward(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setTotalMintReward", reward)
}

// SetTotalMintReward is a paid mutator transaction binding the contract method 0x02f6b29e.
//
// Solidity: function setTotalMintReward(uint256 reward) returns()
func (_DappHorc *DappHorcSession) SetTotalMintReward(reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetTotalMintReward(&_DappHorc.TransactOpts, reward)
}

// SetTotalMintReward is a paid mutator transaction binding the contract method 0x02f6b29e.
//
// Solidity: function setTotalMintReward(uint256 reward) returns()
func (_DappHorc *DappHorcTransactorSession) SetTotalMintReward(reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetTotalMintReward(&_DappHorc.TransactOpts, reward)
}

// SetUpLevelAddr is a paid mutator transaction binding the contract method 0x15509725.
//
// Solidity: function setUpLevelAddr(address a) returns()
func (_DappHorc *DappHorcTransactor) SetUpLevelAddr(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setUpLevelAddr", a)
}

// SetUpLevelAddr is a paid mutator transaction binding the contract method 0x15509725.
//
// Solidity: function setUpLevelAddr(address a) returns()
func (_DappHorc *DappHorcSession) SetUpLevelAddr(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetUpLevelAddr(&_DappHorc.TransactOpts, a)
}

// SetUpLevelAddr is a paid mutator transaction binding the contract method 0x15509725.
//
// Solidity: function setUpLevelAddr(address a) returns()
func (_DappHorc *DappHorcTransactorSession) SetUpLevelAddr(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetUpLevelAddr(&_DappHorc.TransactOpts, a)
}

// SetDailyDuration is a paid mutator transaction binding the contract method 0x0c4e2550.
//
// Solidity: function set_dailyDuration(uint256 r) returns()
func (_DappHorc *DappHorcTransactor) SetDailyDuration(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "set_dailyDuration", r)
}

// SetDailyDuration is a paid mutator transaction binding the contract method 0x0c4e2550.
//
// Solidity: function set_dailyDuration(uint256 r) returns()
func (_DappHorc *DappHorcSession) SetDailyDuration(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetDailyDuration(&_DappHorc.TransactOpts, r)
}

// SetDailyDuration is a paid mutator transaction binding the contract method 0x0c4e2550.
//
// Solidity: function set_dailyDuration(uint256 r) returns()
func (_DappHorc *DappHorcTransactorSession) SetDailyDuration(r *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetDailyDuration(&_DappHorc.TransactOpts, r)
}

// SetExtendCurrent is a paid mutator transaction binding the contract method 0xfd327605.
//
// Solidity: function set_extendCurrent(uint256 num) returns()
func (_DappHorc *DappHorcTransactor) SetExtendCurrent(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "set_extendCurrent", num)
}

// SetExtendCurrent is a paid mutator transaction binding the contract method 0xfd327605.
//
// Solidity: function set_extendCurrent(uint256 num) returns()
func (_DappHorc *DappHorcSession) SetExtendCurrent(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetExtendCurrent(&_DappHorc.TransactOpts, num)
}

// SetExtendCurrent is a paid mutator transaction binding the contract method 0xfd327605.
//
// Solidity: function set_extendCurrent(uint256 num) returns()
func (_DappHorc *DappHorcTransactorSession) SetExtendCurrent(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetExtendCurrent(&_DappHorc.TransactOpts, num)
}

// SetExtendList is a paid mutator transaction binding the contract method 0xe2ac3aca.
//
// Solidity: function set_extendList(address[] addr) returns()
func (_DappHorc *DappHorcTransactor) SetExtendList(opts *bind.TransactOpts, addr []common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "set_extendList", addr)
}

// SetExtendList is a paid mutator transaction binding the contract method 0xe2ac3aca.
//
// Solidity: function set_extendList(address[] addr) returns()
func (_DappHorc *DappHorcSession) SetExtendList(addr []common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetExtendList(&_DappHorc.TransactOpts, addr)
}

// SetExtendList is a paid mutator transaction binding the contract method 0xe2ac3aca.
//
// Solidity: function set_extendList(address[] addr) returns()
func (_DappHorc *DappHorcTransactorSession) SetExtendList(addr []common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetExtendList(&_DappHorc.TransactOpts, addr)
}

// SetExtendRate is a paid mutator transaction binding the contract method 0x6781ae16.
//
// Solidity: function set_extendRate(uint256 num) returns()
func (_DappHorc *DappHorcTransactor) SetExtendRate(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "set_extendRate", num)
}

// SetExtendRate is a paid mutator transaction binding the contract method 0x6781ae16.
//
// Solidity: function set_extendRate(uint256 num) returns()
func (_DappHorc *DappHorcSession) SetExtendRate(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetExtendRate(&_DappHorc.TransactOpts, num)
}

// SetExtendRate is a paid mutator transaction binding the contract method 0x6781ae16.
//
// Solidity: function set_extendRate(uint256 num) returns()
func (_DappHorc *DappHorcTransactorSession) SetExtendRate(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetExtendRate(&_DappHorc.TransactOpts, num)
}

// SetIsExtend is a paid mutator transaction binding the contract method 0x3771746b.
//
// Solidity: function set_isExtend() returns()
func (_DappHorc *DappHorcTransactor) SetIsExtend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "set_isExtend")
}

// SetIsExtend is a paid mutator transaction binding the contract method 0x3771746b.
//
// Solidity: function set_isExtend() returns()
func (_DappHorc *DappHorcSession) SetIsExtend() (*types.Transaction, error) {
	return _DappHorc.Contract.SetIsExtend(&_DappHorc.TransactOpts)
}

// SetIsExtend is a paid mutator transaction binding the contract method 0x3771746b.
//
// Solidity: function set_isExtend() returns()
func (_DappHorc *DappHorcTransactorSession) SetIsExtend() (*types.Transaction, error) {
	return _DappHorc.Contract.SetIsExtend(&_DappHorc.TransactOpts)
}

// Setcanclaim is a paid mutator transaction binding the contract method 0x9b187bfa.
//
// Solidity: function setcanclaim() returns()
func (_DappHorc *DappHorcTransactor) Setcanclaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setcanclaim")
}

// Setcanclaim is a paid mutator transaction binding the contract method 0x9b187bfa.
//
// Solidity: function setcanclaim() returns()
func (_DappHorc *DappHorcSession) Setcanclaim() (*types.Transaction, error) {
	return _DappHorc.Contract.Setcanclaim(&_DappHorc.TransactOpts)
}

// Setcanclaim is a paid mutator transaction binding the contract method 0x9b187bfa.
//
// Solidity: function setcanclaim() returns()
func (_DappHorc *DappHorcTransactorSession) Setcanclaim() (*types.Transaction, error) {
	return _DappHorc.Contract.Setcanclaim(&_DappHorc.TransactOpts)
}

// SetlimitUsdt is a paid mutator transaction binding the contract method 0x2f02d135.
//
// Solidity: function setlimitUsdt(uint256 num) returns()
func (_DappHorc *DappHorcTransactor) SetlimitUsdt(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setlimitUsdt", num)
}

// SetlimitUsdt is a paid mutator transaction binding the contract method 0x2f02d135.
//
// Solidity: function setlimitUsdt(uint256 num) returns()
func (_DappHorc *DappHorcSession) SetlimitUsdt(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetlimitUsdt(&_DappHorc.TransactOpts, num)
}

// SetlimitUsdt is a paid mutator transaction binding the contract method 0x2f02d135.
//
// Solidity: function setlimitUsdt(uint256 num) returns()
func (_DappHorc *DappHorcTransactorSession) SetlimitUsdt(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetlimitUsdt(&_DappHorc.TransactOpts, num)
}

// SetminAmount is a paid mutator transaction binding the contract method 0xd7dc6458.
//
// Solidity: function setminAmount(uint256 num) returns()
func (_DappHorc *DappHorcTransactor) SetminAmount(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setminAmount", num)
}

// SetminAmount is a paid mutator transaction binding the contract method 0xd7dc6458.
//
// Solidity: function setminAmount(uint256 num) returns()
func (_DappHorc *DappHorcSession) SetminAmount(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetminAmount(&_DappHorc.TransactOpts, num)
}

// SetminAmount is a paid mutator transaction binding the contract method 0xd7dc6458.
//
// Solidity: function setminAmount(uint256 num) returns()
func (_DappHorc *DappHorcTransactorSession) SetminAmount(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetminAmount(&_DappHorc.TransactOpts, num)
}

// SetpowX is a paid mutator transaction binding the contract method 0x3c6e9793.
//
// Solidity: function setpowX(uint256 num) returns()
func (_DappHorc *DappHorcTransactor) SetpowX(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setpowX", num)
}

// SetpowX is a paid mutator transaction binding the contract method 0x3c6e9793.
//
// Solidity: function setpowX(uint256 num) returns()
func (_DappHorc *DappHorcSession) SetpowX(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetpowX(&_DappHorc.TransactOpts, num)
}

// SetpowX is a paid mutator transaction binding the contract method 0x3c6e9793.
//
// Solidity: function setpowX(uint256 num) returns()
func (_DappHorc *DappHorcTransactorSession) SetpowX(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetpowX(&_DappHorc.TransactOpts, num)
}

// Setstarttime is a paid mutator transaction binding the contract method 0xce2a0ba6.
//
// Solidity: function setstarttime(uint256 num) returns()
func (_DappHorc *DappHorcTransactor) Setstarttime(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setstarttime", num)
}

// Setstarttime is a paid mutator transaction binding the contract method 0xce2a0ba6.
//
// Solidity: function setstarttime(uint256 num) returns()
func (_DappHorc *DappHorcSession) Setstarttime(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.Setstarttime(&_DappHorc.TransactOpts, num)
}

// Setstarttime is a paid mutator transaction binding the contract method 0xce2a0ba6.
//
// Solidity: function setstarttime(uint256 num) returns()
func (_DappHorc *DappHorcTransactorSession) Setstarttime(num *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.Setstarttime(&_DappHorc.TransactOpts, num)
}

// Setstop is a paid mutator transaction binding the contract method 0xa084c854.
//
// Solidity: function setstop(address addr, bool succ) returns()
func (_DappHorc *DappHorcTransactor) Setstop(opts *bind.TransactOpts, addr common.Address, succ bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setstop", addr, succ)
}

// Setstop is a paid mutator transaction binding the contract method 0xa084c854.
//
// Solidity: function setstop(address addr, bool succ) returns()
func (_DappHorc *DappHorcSession) Setstop(addr common.Address, succ bool) (*types.Transaction, error) {
	return _DappHorc.Contract.Setstop(&_DappHorc.TransactOpts, addr, succ)
}

// Setstop is a paid mutator transaction binding the contract method 0xa084c854.
//
// Solidity: function setstop(address addr, bool succ) returns()
func (_DappHorc *DappHorcTransactorSession) Setstop(addr common.Address, succ bool) (*types.Transaction, error) {
	return _DappHorc.Contract.Setstop(&_DappHorc.TransactOpts, addr, succ)
}

// Setuserinfo is a paid mutator transaction binding the contract method 0x26d3c6e8.
//
// Solidity: function setuserinfo(address account, bool isActive, uint256 amount, uint256 amountUsdt, uint256 rewardMintDebt, uint256 calMintReward, uint256 level, uint256 inviteAmount, uint256 invite2Amount, uint256 nodeNum, uint256 superNum, uint256 teamPow) returns()
func (_DappHorc *DappHorcTransactor) Setuserinfo(opts *bind.TransactOpts, account common.Address, isActive bool, amount *big.Int, amountUsdt *big.Int, rewardMintDebt *big.Int, calMintReward *big.Int, level *big.Int, inviteAmount *big.Int, invite2Amount *big.Int, nodeNum *big.Int, superNum *big.Int, teamPow *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setuserinfo", account, isActive, amount, amountUsdt, rewardMintDebt, calMintReward, level, inviteAmount, invite2Amount, nodeNum, superNum, teamPow)
}

// Setuserinfo is a paid mutator transaction binding the contract method 0x26d3c6e8.
//
// Solidity: function setuserinfo(address account, bool isActive, uint256 amount, uint256 amountUsdt, uint256 rewardMintDebt, uint256 calMintReward, uint256 level, uint256 inviteAmount, uint256 invite2Amount, uint256 nodeNum, uint256 superNum, uint256 teamPow) returns()
func (_DappHorc *DappHorcSession) Setuserinfo(account common.Address, isActive bool, amount *big.Int, amountUsdt *big.Int, rewardMintDebt *big.Int, calMintReward *big.Int, level *big.Int, inviteAmount *big.Int, invite2Amount *big.Int, nodeNum *big.Int, superNum *big.Int, teamPow *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.Setuserinfo(&_DappHorc.TransactOpts, account, isActive, amount, amountUsdt, rewardMintDebt, calMintReward, level, inviteAmount, invite2Amount, nodeNum, superNum, teamPow)
}

// Setuserinfo is a paid mutator transaction binding the contract method 0x26d3c6e8.
//
// Solidity: function setuserinfo(address account, bool isActive, uint256 amount, uint256 amountUsdt, uint256 rewardMintDebt, uint256 calMintReward, uint256 level, uint256 inviteAmount, uint256 invite2Amount, uint256 nodeNum, uint256 superNum, uint256 teamPow) returns()
func (_DappHorc *DappHorcTransactorSession) Setuserinfo(account common.Address, isActive bool, amount *big.Int, amountUsdt *big.Int, rewardMintDebt *big.Int, calMintReward *big.Int, level *big.Int, inviteAmount *big.Int, invite2Amount *big.Int, nodeNum *big.Int, superNum *big.Int, teamPow *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.Setuserinfo(&_DappHorc.TransactOpts, account, isActive, amount, amountUsdt, rewardMintDebt, calMintReward, level, inviteAmount, invite2Amount, nodeNum, superNum, teamPow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappHorc *DappHorcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappHorc *DappHorcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.TransferOwnership(&_DappHorc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappHorc *DappHorcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.TransferOwnership(&_DappHorc.TransactOpts, newOwner)
}

// UpdateRewardPerDay is a paid mutator transaction binding the contract method 0x0a9b19de.
//
// Solidity: function updateRewardPerDay() returns()
func (_DappHorc *DappHorcTransactor) UpdateRewardPerDay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "updateRewardPerDay")
}

// UpdateRewardPerDay is a paid mutator transaction binding the contract method 0x0a9b19de.
//
// Solidity: function updateRewardPerDay() returns()
func (_DappHorc *DappHorcSession) UpdateRewardPerDay() (*types.Transaction, error) {
	return _DappHorc.Contract.UpdateRewardPerDay(&_DappHorc.TransactOpts)
}

// UpdateRewardPerDay is a paid mutator transaction binding the contract method 0x0a9b19de.
//
// Solidity: function updateRewardPerDay() returns()
func (_DappHorc *DappHorcTransactorSession) UpdateRewardPerDay() (*types.Transaction, error) {
	return _DappHorc.Contract.UpdateRewardPerDay(&_DappHorc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappHorc *DappHorcTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappHorc *DappHorcSession) Receive() (*types.Transaction, error) {
	return _DappHorc.Contract.Receive(&_DappHorc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappHorc *DappHorcTransactorSession) Receive() (*types.Transaction, error) {
	return _DappHorc.Contract.Receive(&_DappHorc.TransactOpts)
}

// DappHorcADDPOWIterator is returned from FilterADDPOW and is used to iterate over the raw logs and unpacked data for ADDPOW events raised by the DappHorc contract.
type DappHorcADDPOWIterator struct {
	Event *DappHorcADDPOW // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappHorcADDPOWIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappHorcADDPOW)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappHorcADDPOW)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappHorcADDPOWIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappHorcADDPOWIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappHorcADDPOW represents a ADDPOW event raised by the DappHorc contract.
type DappHorcADDPOW struct {
	Addr  common.Address
	Pow   *big.Int
	Ttype *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterADDPOW is a free log retrieval operation binding the contract event 0xb6683455a303b44ee8c2b35e1d3988ba05daedaa41becbc19196d60cf6f09bd9.
//
// Solidity: event ADDPOW(address addr, uint256 pow, uint256 ttype)
func (_DappHorc *DappHorcFilterer) FilterADDPOW(opts *bind.FilterOpts) (*DappHorcADDPOWIterator, error) {

	logs, sub, err := _DappHorc.contract.FilterLogs(opts, "ADDPOW")
	if err != nil {
		return nil, err
	}
	return &DappHorcADDPOWIterator{contract: _DappHorc.contract, event: "ADDPOW", logs: logs, sub: sub}, nil
}

// WatchADDPOW is a free log subscription operation binding the contract event 0xb6683455a303b44ee8c2b35e1d3988ba05daedaa41becbc19196d60cf6f09bd9.
//
// Solidity: event ADDPOW(address addr, uint256 pow, uint256 ttype)
func (_DappHorc *DappHorcFilterer) WatchADDPOW(opts *bind.WatchOpts, sink chan<- *DappHorcADDPOW) (event.Subscription, error) {

	logs, sub, err := _DappHorc.contract.WatchLogs(opts, "ADDPOW")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappHorcADDPOW)
				if err := _DappHorc.contract.UnpackLog(event, "ADDPOW", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseADDPOW is a log parse operation binding the contract event 0xb6683455a303b44ee8c2b35e1d3988ba05daedaa41becbc19196d60cf6f09bd9.
//
// Solidity: event ADDPOW(address addr, uint256 pow, uint256 ttype)
func (_DappHorc *DappHorcFilterer) ParseADDPOW(log types.Log) (*DappHorcADDPOW, error) {
	event := new(DappHorcADDPOW)
	if err := _DappHorc.contract.UnpackLog(event, "ADDPOW", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappHorcCLAIMEVENTIterator is returned from FilterCLAIMEVENT and is used to iterate over the raw logs and unpacked data for CLAIMEVENT events raised by the DappHorc contract.
type DappHorcCLAIMEVENTIterator struct {
	Event *DappHorcCLAIMEVENT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappHorcCLAIMEVENTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappHorcCLAIMEVENT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappHorcCLAIMEVENT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappHorcCLAIMEVENTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappHorcCLAIMEVENTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappHorcCLAIMEVENT represents a CLAIMEVENT event raised by the DappHorc contract.
type DappHorcCLAIMEVENT struct {
	Addr  common.Address
	Num   *big.Int
	Ttype *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCLAIMEVENT is a free log retrieval operation binding the contract event 0x6482570a7245374fe3789f9c38cc58514166d451234dec215ce6cda72c306c51.
//
// Solidity: event CLAIMEVENT(address addr, uint256 num, uint256 ttype)
func (_DappHorc *DappHorcFilterer) FilterCLAIMEVENT(opts *bind.FilterOpts) (*DappHorcCLAIMEVENTIterator, error) {

	logs, sub, err := _DappHorc.contract.FilterLogs(opts, "CLAIMEVENT")
	if err != nil {
		return nil, err
	}
	return &DappHorcCLAIMEVENTIterator{contract: _DappHorc.contract, event: "CLAIMEVENT", logs: logs, sub: sub}, nil
}

// WatchCLAIMEVENT is a free log subscription operation binding the contract event 0x6482570a7245374fe3789f9c38cc58514166d451234dec215ce6cda72c306c51.
//
// Solidity: event CLAIMEVENT(address addr, uint256 num, uint256 ttype)
func (_DappHorc *DappHorcFilterer) WatchCLAIMEVENT(opts *bind.WatchOpts, sink chan<- *DappHorcCLAIMEVENT) (event.Subscription, error) {

	logs, sub, err := _DappHorc.contract.WatchLogs(opts, "CLAIMEVENT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappHorcCLAIMEVENT)
				if err := _DappHorc.contract.UnpackLog(event, "CLAIMEVENT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCLAIMEVENT is a log parse operation binding the contract event 0x6482570a7245374fe3789f9c38cc58514166d451234dec215ce6cda72c306c51.
//
// Solidity: event CLAIMEVENT(address addr, uint256 num, uint256 ttype)
func (_DappHorc *DappHorcFilterer) ParseCLAIMEVENT(log types.Log) (*DappHorcCLAIMEVENT, error) {
	event := new(DappHorcCLAIMEVENT)
	if err := _DappHorc.contract.UnpackLog(event, "CLAIMEVENT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappHorcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DappHorc contract.
type DappHorcOwnershipTransferredIterator struct {
	Event *DappHorcOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappHorcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappHorcOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappHorcOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappHorcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappHorcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappHorcOwnershipTransferred represents a OwnershipTransferred event raised by the DappHorc contract.
type DappHorcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappHorc *DappHorcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DappHorcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappHorc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DappHorcOwnershipTransferredIterator{contract: _DappHorc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappHorc *DappHorcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DappHorcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappHorc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappHorcOwnershipTransferred)
				if err := _DappHorc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappHorc *DappHorcFilterer) ParseOwnershipTransferred(log types.Log) (*DappHorcOwnershipTransferred, error) {
	event := new(DappHorcOwnershipTransferred)
	if err := _DappHorc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
