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
)

// DappHorcMetaData contains all meta data concerning the DappHorc contract.
var DappHorcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"__getTokenETHReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_amountDailyUp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_autoCalRewardPerDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_autoCalSellRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_binder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_burnReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fundAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_highPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_highPriceSellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_inProject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_inviteFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_invitor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastAmountRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastDailyUpTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_normalRewardPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_updateDailyUpRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"addConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addJoinToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addMintAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"calInvite\",\"type\":\"bool\"}],\"name\":\"addUserAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"invitor\",\"type\":\"address\"}],\"name\":\"bindInvitor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"invitor\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAutoRewardConfigs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"uppeds\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"droppeds\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdtDecimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintRewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintRewardTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDailyReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dailyAmountRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"defaultInvitor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pauseJoin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBinderLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"}],\"name\":\"getJoinTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getJoinTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenDecimals\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"tokenSymbols\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"poolUsdts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"poolTokens\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPendingMintReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSellRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sellRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenETHPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"joinTokenBalances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"joinTokenAllowances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"teamNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setAmountDailyUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"e\",\"type\":\"bool\"}],\"name\":\"setAutoCalRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"e\",\"type\":\"bool\"}],\"name\":\"setAutoCalSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setBurnReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"e\",\"type\":\"bool\"}],\"name\":\"setConfigDropped\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setConfigPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerDay\",\"type\":\"uint256\"}],\"name\":\"setConfigRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"e\",\"type\":\"bool\"}],\"name\":\"setConfigUpped\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"}],\"name\":\"setDefaultInvitor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setFundAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setHighPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setHighPriceSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setInProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setInviteFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setJoinToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"setJoinTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setLastAmountRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"setLastDailyUpTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setNormalRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"setSellRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setTotalMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewardPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewPoolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMintPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastMintTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMintReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"viewUserInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calMintReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardMintDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(DappHorcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetTokenETHReserves is a free data retrieval call binding the contract method 0x0a849d61.
//
// Solidity: function __getTokenETHReserves(address token) view returns(uint256 rEth, uint256 rToken)
func (_DappHorc *DappHorcCaller) GetTokenETHReserves(opts *bind.CallOpts, token common.Address) (struct {
	REth   *big.Int
	RToken *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "__getTokenETHReserves", token)

	outstruct := new(struct {
		REth   *big.Int
		RToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.REth = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenETHReserves is a free data retrieval call binding the contract method 0x0a849d61.
//
// Solidity: function __getTokenETHReserves(address token) view returns(uint256 rEth, uint256 rToken)
func (_DappHorc *DappHorcSession) GetTokenETHReserves(token common.Address) (struct {
	REth   *big.Int
	RToken *big.Int
}, error) {
	return _DappHorc.Contract.GetTokenETHReserves(&_DappHorc.CallOpts, token)
}

// GetTokenETHReserves is a free data retrieval call binding the contract method 0x0a849d61.
//
// Solidity: function __getTokenETHReserves(address token) view returns(uint256 rEth, uint256 rToken)
func (_DappHorc *DappHorcCallerSession) GetTokenETHReserves(token common.Address) (struct {
	REth   *big.Int
	RToken *big.Int
}, error) {
	return _DappHorc.Contract.GetTokenETHReserves(&_DappHorc.CallOpts, token)
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

// InProject is a free data retrieval call binding the contract method 0xe5d65633.
//
// Solidity: function _inProject(address ) view returns(bool)
func (_DappHorc *DappHorcCaller) InProject(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "_inProject", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InProject is a free data retrieval call binding the contract method 0xe5d65633.
//
// Solidity: function _inProject(address ) view returns(bool)
func (_DappHorc *DappHorcSession) InProject(arg0 common.Address) (bool, error) {
	return _DappHorc.Contract.InProject(&_DappHorc.CallOpts, arg0)
}

// InProject is a free data retrieval call binding the contract method 0xe5d65633.
//
// Solidity: function _inProject(address ) view returns(bool)
func (_DappHorc *DappHorcCallerSession) InProject(arg0 common.Address) (bool, error) {
	return _DappHorc.Contract.InProject(&_DappHorc.CallOpts, arg0)
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

// GetAutoRewardConfigs is a free data retrieval call binding the contract method 0x7e6a584c.
//
// Solidity: function getAutoRewardConfigs() view returns(uint256[] prices, uint256[] rewards, bool[] uppeds, bool[] droppeds)
func (_DappHorc *DappHorcCaller) GetAutoRewardConfigs(opts *bind.CallOpts) (struct {
	Prices   []*big.Int
	Rewards  []*big.Int
	Uppeds   []bool
	Droppeds []bool
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getAutoRewardConfigs")

	outstruct := new(struct {
		Prices   []*big.Int
		Rewards  []*big.Int
		Uppeds   []bool
		Droppeds []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Prices = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Rewards = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Uppeds = *abi.ConvertType(out[2], new([]bool)).(*[]bool)
	outstruct.Droppeds = *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetAutoRewardConfigs is a free data retrieval call binding the contract method 0x7e6a584c.
//
// Solidity: function getAutoRewardConfigs() view returns(uint256[] prices, uint256[] rewards, bool[] uppeds, bool[] droppeds)
func (_DappHorc *DappHorcSession) GetAutoRewardConfigs() (struct {
	Prices   []*big.Int
	Rewards  []*big.Int
	Uppeds   []bool
	Droppeds []bool
}, error) {
	return _DappHorc.Contract.GetAutoRewardConfigs(&_DappHorc.CallOpts)
}

// GetAutoRewardConfigs is a free data retrieval call binding the contract method 0x7e6a584c.
//
// Solidity: function getAutoRewardConfigs() view returns(uint256[] prices, uint256[] rewards, bool[] uppeds, bool[] droppeds)
func (_DappHorc *DappHorcCallerSession) GetAutoRewardConfigs() (struct {
	Prices   []*big.Int
	Rewards  []*big.Int
	Uppeds   []bool
	Droppeds []bool
}, error) {
	return _DappHorc.Contract.GetAutoRewardConfigs(&_DappHorc.CallOpts)
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

// GetTokenETHPrice is a free data retrieval call binding the contract method 0xe88c1997.
//
// Solidity: function getTokenETHPrice(address token) view returns(uint256 price)
func (_DappHorc *DappHorcCaller) GetTokenETHPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "getTokenETHPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenETHPrice is a free data retrieval call binding the contract method 0xe88c1997.
//
// Solidity: function getTokenETHPrice(address token) view returns(uint256 price)
func (_DappHorc *DappHorcSession) GetTokenETHPrice(token common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetTokenETHPrice(&_DappHorc.CallOpts, token)
}

// GetTokenETHPrice is a free data retrieval call binding the contract method 0xe88c1997.
//
// Solidity: function getTokenETHPrice(address token) view returns(uint256 price)
func (_DappHorc *DappHorcCallerSession) GetTokenETHPrice(token common.Address) (*big.Int, error) {
	return _DappHorc.Contract.GetTokenETHPrice(&_DappHorc.CallOpts, token)
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
// Solidity: function viewUserInfo(address account) view returns(bool isActive, uint256 amount, uint256 calMintReward, uint256 rewardMintDebt)
func (_DappHorc *DappHorcCaller) ViewUserInfo(opts *bind.CallOpts, account common.Address) (struct {
	IsActive       bool
	Amount         *big.Int
	CalMintReward  *big.Int
	RewardMintDebt *big.Int
}, error) {
	var out []interface{}
	err := _DappHorc.contract.Call(opts, &out, "viewUserInfo", account)

	outstruct := new(struct {
		IsActive       bool
		Amount         *big.Int
		CalMintReward  *big.Int
		RewardMintDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CalMintReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardMintDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ViewUserInfo is a free data retrieval call binding the contract method 0xa37d9850.
//
// Solidity: function viewUserInfo(address account) view returns(bool isActive, uint256 amount, uint256 calMintReward, uint256 rewardMintDebt)
func (_DappHorc *DappHorcSession) ViewUserInfo(account common.Address) (struct {
	IsActive       bool
	Amount         *big.Int
	CalMintReward  *big.Int
	RewardMintDebt *big.Int
}, error) {
	return _DappHorc.Contract.ViewUserInfo(&_DappHorc.CallOpts, account)
}

// ViewUserInfo is a free data retrieval call binding the contract method 0xa37d9850.
//
// Solidity: function viewUserInfo(address account) view returns(bool isActive, uint256 amount, uint256 calMintReward, uint256 rewardMintDebt)
func (_DappHorc *DappHorcCallerSession) ViewUserInfo(account common.Address) (struct {
	IsActive       bool
	Amount         *big.Int
	CalMintReward  *big.Int
	RewardMintDebt *big.Int
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

// AddConfig is a paid mutator transaction binding the contract method 0x157e4089.
//
// Solidity: function addConfig(uint256 price, uint256 reward) returns()
func (_DappHorc *DappHorcTransactor) AddConfig(opts *bind.TransactOpts, price *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "addConfig", price, reward)
}

// AddConfig is a paid mutator transaction binding the contract method 0x157e4089.
//
// Solidity: function addConfig(uint256 price, uint256 reward) returns()
func (_DappHorc *DappHorcSession) AddConfig(price *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.AddConfig(&_DappHorc.TransactOpts, price, reward)
}

// AddConfig is a paid mutator transaction binding the contract method 0x157e4089.
//
// Solidity: function addConfig(uint256 price, uint256 reward) returns()
func (_DappHorc *DappHorcTransactorSession) AddConfig(price *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.AddConfig(&_DappHorc.TransactOpts, price, reward)
}

// AddJoinToken is a paid mutator transaction binding the contract method 0x656cbb86.
//
// Solidity: function addJoinToken(address token) returns()
func (_DappHorc *DappHorcTransactor) AddJoinToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "addJoinToken", token)
}

// AddJoinToken is a paid mutator transaction binding the contract method 0x656cbb86.
//
// Solidity: function addJoinToken(address token) returns()
func (_DappHorc *DappHorcSession) AddJoinToken(token common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.AddJoinToken(&_DappHorc.TransactOpts, token)
}

// AddJoinToken is a paid mutator transaction binding the contract method 0x656cbb86.
//
// Solidity: function addJoinToken(address token) returns()
func (_DappHorc *DappHorcTransactorSession) AddJoinToken(token common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.AddJoinToken(&_DappHorc.TransactOpts, token)
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

// ClearConfig is a paid mutator transaction binding the contract method 0xacb62d7c.
//
// Solidity: function clearConfig() returns()
func (_DappHorc *DappHorcTransactor) ClearConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "clearConfig")
}

// ClearConfig is a paid mutator transaction binding the contract method 0xacb62d7c.
//
// Solidity: function clearConfig() returns()
func (_DappHorc *DappHorcSession) ClearConfig() (*types.Transaction, error) {
	return _DappHorc.Contract.ClearConfig(&_DappHorc.TransactOpts)
}

// ClearConfig is a paid mutator transaction binding the contract method 0xacb62d7c.
//
// Solidity: function clearConfig() returns()
func (_DappHorc *DappHorcTransactorSession) ClearConfig() (*types.Transaction, error) {
	return _DappHorc.Contract.ClearConfig(&_DappHorc.TransactOpts)
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

// SetAutoCalRewardPerDay is a paid mutator transaction binding the contract method 0x2c410d60.
//
// Solidity: function setAutoCalRewardPerDay(bool e) returns()
func (_DappHorc *DappHorcTransactor) SetAutoCalRewardPerDay(opts *bind.TransactOpts, e bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setAutoCalRewardPerDay", e)
}

// SetAutoCalRewardPerDay is a paid mutator transaction binding the contract method 0x2c410d60.
//
// Solidity: function setAutoCalRewardPerDay(bool e) returns()
func (_DappHorc *DappHorcSession) SetAutoCalRewardPerDay(e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalRewardPerDay(&_DappHorc.TransactOpts, e)
}

// SetAutoCalRewardPerDay is a paid mutator transaction binding the contract method 0x2c410d60.
//
// Solidity: function setAutoCalRewardPerDay(bool e) returns()
func (_DappHorc *DappHorcTransactorSession) SetAutoCalRewardPerDay(e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalRewardPerDay(&_DappHorc.TransactOpts, e)
}

// SetAutoCalSellRate is a paid mutator transaction binding the contract method 0x058106be.
//
// Solidity: function setAutoCalSellRate(bool e) returns()
func (_DappHorc *DappHorcTransactor) SetAutoCalSellRate(opts *bind.TransactOpts, e bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setAutoCalSellRate", e)
}

// SetAutoCalSellRate is a paid mutator transaction binding the contract method 0x058106be.
//
// Solidity: function setAutoCalSellRate(bool e) returns()
func (_DappHorc *DappHorcSession) SetAutoCalSellRate(e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalSellRate(&_DappHorc.TransactOpts, e)
}

// SetAutoCalSellRate is a paid mutator transaction binding the contract method 0x058106be.
//
// Solidity: function setAutoCalSellRate(bool e) returns()
func (_DappHorc *DappHorcTransactorSession) SetAutoCalSellRate(e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetAutoCalSellRate(&_DappHorc.TransactOpts, e)
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

// SetConfigDropped is a paid mutator transaction binding the contract method 0x9e0dfb46.
//
// Solidity: function setConfigDropped(uint256 i, bool e) returns()
func (_DappHorc *DappHorcTransactor) SetConfigDropped(opts *bind.TransactOpts, i *big.Int, e bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setConfigDropped", i, e)
}

// SetConfigDropped is a paid mutator transaction binding the contract method 0x9e0dfb46.
//
// Solidity: function setConfigDropped(uint256 i, bool e) returns()
func (_DappHorc *DappHorcSession) SetConfigDropped(i *big.Int, e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigDropped(&_DappHorc.TransactOpts, i, e)
}

// SetConfigDropped is a paid mutator transaction binding the contract method 0x9e0dfb46.
//
// Solidity: function setConfigDropped(uint256 i, bool e) returns()
func (_DappHorc *DappHorcTransactorSession) SetConfigDropped(i *big.Int, e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigDropped(&_DappHorc.TransactOpts, i, e)
}

// SetConfigPrice is a paid mutator transaction binding the contract method 0x19632b25.
//
// Solidity: function setConfigPrice(uint256 i, uint256 price) returns()
func (_DappHorc *DappHorcTransactor) SetConfigPrice(opts *bind.TransactOpts, i *big.Int, price *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setConfigPrice", i, price)
}

// SetConfigPrice is a paid mutator transaction binding the contract method 0x19632b25.
//
// Solidity: function setConfigPrice(uint256 i, uint256 price) returns()
func (_DappHorc *DappHorcSession) SetConfigPrice(i *big.Int, price *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigPrice(&_DappHorc.TransactOpts, i, price)
}

// SetConfigPrice is a paid mutator transaction binding the contract method 0x19632b25.
//
// Solidity: function setConfigPrice(uint256 i, uint256 price) returns()
func (_DappHorc *DappHorcTransactorSession) SetConfigPrice(i *big.Int, price *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigPrice(&_DappHorc.TransactOpts, i, price)
}

// SetConfigRewardPerDay is a paid mutator transaction binding the contract method 0x30b1c2e2.
//
// Solidity: function setConfigRewardPerDay(uint256 i, uint256 rewardPerDay) returns()
func (_DappHorc *DappHorcTransactor) SetConfigRewardPerDay(opts *bind.TransactOpts, i *big.Int, rewardPerDay *big.Int) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setConfigRewardPerDay", i, rewardPerDay)
}

// SetConfigRewardPerDay is a paid mutator transaction binding the contract method 0x30b1c2e2.
//
// Solidity: function setConfigRewardPerDay(uint256 i, uint256 rewardPerDay) returns()
func (_DappHorc *DappHorcSession) SetConfigRewardPerDay(i *big.Int, rewardPerDay *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigRewardPerDay(&_DappHorc.TransactOpts, i, rewardPerDay)
}

// SetConfigRewardPerDay is a paid mutator transaction binding the contract method 0x30b1c2e2.
//
// Solidity: function setConfigRewardPerDay(uint256 i, uint256 rewardPerDay) returns()
func (_DappHorc *DappHorcTransactorSession) SetConfigRewardPerDay(i *big.Int, rewardPerDay *big.Int) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigRewardPerDay(&_DappHorc.TransactOpts, i, rewardPerDay)
}

// SetConfigUpped is a paid mutator transaction binding the contract method 0x1512ac99.
//
// Solidity: function setConfigUpped(uint256 i, bool e) returns()
func (_DappHorc *DappHorcTransactor) SetConfigUpped(opts *bind.TransactOpts, i *big.Int, e bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setConfigUpped", i, e)
}

// SetConfigUpped is a paid mutator transaction binding the contract method 0x1512ac99.
//
// Solidity: function setConfigUpped(uint256 i, bool e) returns()
func (_DappHorc *DappHorcSession) SetConfigUpped(i *big.Int, e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigUpped(&_DappHorc.TransactOpts, i, e)
}

// SetConfigUpped is a paid mutator transaction binding the contract method 0x1512ac99.
//
// Solidity: function setConfigUpped(uint256 i, bool e) returns()
func (_DappHorc *DappHorcTransactorSession) SetConfigUpped(i *big.Int, e bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetConfigUpped(&_DappHorc.TransactOpts, i, e)
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

// SetFundAddress is a paid mutator transaction binding the contract method 0x85dc3004.
//
// Solidity: function setFundAddress(address a) returns()
func (_DappHorc *DappHorcTransactor) SetFundAddress(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setFundAddress", a)
}

// SetFundAddress is a paid mutator transaction binding the contract method 0x85dc3004.
//
// Solidity: function setFundAddress(address a) returns()
func (_DappHorc *DappHorcSession) SetFundAddress(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetFundAddress(&_DappHorc.TransactOpts, a)
}

// SetFundAddress is a paid mutator transaction binding the contract method 0x85dc3004.
//
// Solidity: function setFundAddress(address a) returns()
func (_DappHorc *DappHorcTransactorSession) SetFundAddress(a common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetFundAddress(&_DappHorc.TransactOpts, a)
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

// SetInProject is a paid mutator transaction binding the contract method 0x692c7bf7.
//
// Solidity: function setInProject(address adr, bool enable) returns()
func (_DappHorc *DappHorcTransactor) SetInProject(opts *bind.TransactOpts, adr common.Address, enable bool) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setInProject", adr, enable)
}

// SetInProject is a paid mutator transaction binding the contract method 0x692c7bf7.
//
// Solidity: function setInProject(address adr, bool enable) returns()
func (_DappHorc *DappHorcSession) SetInProject(adr common.Address, enable bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetInProject(&_DappHorc.TransactOpts, adr, enable)
}

// SetInProject is a paid mutator transaction binding the contract method 0x692c7bf7.
//
// Solidity: function setInProject(address adr, bool enable) returns()
func (_DappHorc *DappHorcTransactorSession) SetInProject(adr common.Address, enable bool) (*types.Transaction, error) {
	return _DappHorc.Contract.SetInProject(&_DappHorc.TransactOpts, adr, enable)
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
// Solidity: function setJoinToken(uint256 i, address token) returns()
func (_DappHorc *DappHorcTransactor) SetJoinToken(opts *bind.TransactOpts, i *big.Int, token common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setJoinToken", i, token)
}

// SetJoinToken is a paid mutator transaction binding the contract method 0x138c9ddd.
//
// Solidity: function setJoinToken(uint256 i, address token) returns()
func (_DappHorc *DappHorcSession) SetJoinToken(i *big.Int, token common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinToken(&_DappHorc.TransactOpts, i, token)
}

// SetJoinToken is a paid mutator transaction binding the contract method 0x138c9ddd.
//
// Solidity: function setJoinToken(uint256 i, address token) returns()
func (_DappHorc *DappHorcTransactorSession) SetJoinToken(i *big.Int, token common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinToken(&_DappHorc.TransactOpts, i, token)
}

// SetJoinTokens is a paid mutator transaction binding the contract method 0x8caf1528.
//
// Solidity: function setJoinTokens(address[] tokens) returns()
func (_DappHorc *DappHorcTransactor) SetJoinTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _DappHorc.contract.Transact(opts, "setJoinTokens", tokens)
}

// SetJoinTokens is a paid mutator transaction binding the contract method 0x8caf1528.
//
// Solidity: function setJoinTokens(address[] tokens) returns()
func (_DappHorc *DappHorcSession) SetJoinTokens(tokens []common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinTokens(&_DappHorc.TransactOpts, tokens)
}

// SetJoinTokens is a paid mutator transaction binding the contract method 0x8caf1528.
//
// Solidity: function setJoinTokens(address[] tokens) returns()
func (_DappHorc *DappHorcTransactorSession) SetJoinTokens(tokens []common.Address) (*types.Transaction, error) {
	return _DappHorc.Contract.SetJoinTokens(&_DappHorc.TransactOpts, tokens)
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
