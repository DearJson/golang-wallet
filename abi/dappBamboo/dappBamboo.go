// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dapp_bamboo

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

// DappBambooMetaData contains all meta data concerning the DappBamboo contract.
var DappBambooMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIntoLiqudity\",\"type\":\"uint256\"}],\"name\":\"SwapAndLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SwapAndLiquifyEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LPFeefenhong\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRADE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_autoSwap\",\"outputs\":[{\"internalType\":\"contractAutoSwap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_destroyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fundFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_inviterFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lpFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_rOwned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_taxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_updated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"includeInFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviterNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpenBuy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpensell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRoute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquifyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensSellToAddToLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setDes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setFun\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setFun2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setLiquifyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"setMinAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setOpen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"setStop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSwapAndLiquifyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setTRADE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setisOpenBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setisOpensell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shareholderIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shareholders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapAndLiquifyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DappBambooABI is the input ABI used to generate the binding from.
// Deprecated: Use DappBambooMetaData.ABI instead.
var DappBambooABI = DappBambooMetaData.ABI

// DappBamboo is an auto generated Go binding around an Ethereum contract.
type DappBamboo struct {
	DappBambooCaller     // Read-only binding to the contract
	DappBambooTransactor // Write-only binding to the contract
	DappBambooFilterer   // Log filterer for contract events
}

// DappBambooCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappBambooCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappBambooTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappBambooTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappBambooFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappBambooFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappBambooSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappBambooSession struct {
	Contract     *DappBamboo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappBambooCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappBambooCallerSession struct {
	Contract *DappBambooCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DappBambooTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappBambooTransactorSession struct {
	Contract     *DappBambooTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DappBambooRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappBambooRaw struct {
	Contract *DappBamboo // Generic contract binding to access the raw methods on
}

// DappBambooCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappBambooCallerRaw struct {
	Contract *DappBambooCaller // Generic read-only contract binding to access the raw methods on
}

// DappBambooTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappBambooTransactorRaw struct {
	Contract *DappBambooTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappBamboo creates a new instance of DappBamboo, bound to a specific deployed contract.
func NewDappBamboo(address common.Address, backend bind.ContractBackend) (*DappBamboo, error) {
	contract, err := bindDappBamboo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappBamboo{DappBambooCaller: DappBambooCaller{contract: contract}, DappBambooTransactor: DappBambooTransactor{contract: contract}, DappBambooFilterer: DappBambooFilterer{contract: contract}}, nil
}

// NewDappBambooCaller creates a new read-only instance of DappBamboo, bound to a specific deployed contract.
func NewDappBambooCaller(address common.Address, caller bind.ContractCaller) (*DappBambooCaller, error) {
	contract, err := bindDappBamboo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappBambooCaller{contract: contract}, nil
}

// NewDappBambooTransactor creates a new write-only instance of DappBamboo, bound to a specific deployed contract.
func NewDappBambooTransactor(address common.Address, transactor bind.ContractTransactor) (*DappBambooTransactor, error) {
	contract, err := bindDappBamboo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappBambooTransactor{contract: contract}, nil
}

// NewDappBambooFilterer creates a new log filterer instance of DappBamboo, bound to a specific deployed contract.
func NewDappBambooFilterer(address common.Address, filterer bind.ContractFilterer) (*DappBambooFilterer, error) {
	contract, err := bindDappBamboo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappBambooFilterer{contract: contract}, nil
}

// bindDappBamboo binds a generic wrapper to an already deployed contract.
func bindDappBamboo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappBambooABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappBamboo *DappBambooRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappBamboo.Contract.DappBambooCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappBamboo *DappBambooRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.Contract.DappBambooTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappBamboo *DappBambooRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappBamboo.Contract.DappBambooTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappBamboo *DappBambooCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappBamboo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappBamboo *DappBambooTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappBamboo *DappBambooTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappBamboo.Contract.contract.Transact(opts, method, params...)
}

// LPFeefenhong is a free data retrieval call binding the contract method 0x0b1d406b.
//
// Solidity: function LPFeefenhong() view returns(uint256)
func (_DappBamboo *DappBambooCaller) LPFeefenhong(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "LPFeefenhong")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPFeefenhong is a free data retrieval call binding the contract method 0x0b1d406b.
//
// Solidity: function LPFeefenhong() view returns(uint256)
func (_DappBamboo *DappBambooSession) LPFeefenhong() (*big.Int, error) {
	return _DappBamboo.Contract.LPFeefenhong(&_DappBamboo.CallOpts)
}

// LPFeefenhong is a free data retrieval call binding the contract method 0x0b1d406b.
//
// Solidity: function LPFeefenhong() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) LPFeefenhong() (*big.Int, error) {
	return _DappBamboo.Contract.LPFeefenhong(&_DappBamboo.CallOpts)
}

// TRADE is a free data retrieval call binding the contract method 0xc08ed94a.
//
// Solidity: function TRADE() view returns(address)
func (_DappBamboo *DappBambooCaller) TRADE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "TRADE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TRADE is a free data retrieval call binding the contract method 0xc08ed94a.
//
// Solidity: function TRADE() view returns(address)
func (_DappBamboo *DappBambooSession) TRADE() (common.Address, error) {
	return _DappBamboo.Contract.TRADE(&_DappBamboo.CallOpts)
}

// TRADE is a free data retrieval call binding the contract method 0xc08ed94a.
//
// Solidity: function TRADE() view returns(address)
func (_DappBamboo *DappBambooCallerSession) TRADE() (common.Address, error) {
	return _DappBamboo.Contract.TRADE(&_DappBamboo.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_DappBamboo *DappBambooCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_DappBamboo *DappBambooSession) USDT() (common.Address, error) {
	return _DappBamboo.Contract.USDT(&_DappBamboo.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_DappBamboo *DappBambooCallerSession) USDT() (common.Address, error) {
	return _DappBamboo.Contract.USDT(&_DappBamboo.CallOpts)
}

// AutoSwap is a free data retrieval call binding the contract method 0xf6e1c1c4.
//
// Solidity: function _autoSwap() view returns(address)
func (_DappBamboo *DappBambooCaller) AutoSwap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_autoSwap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AutoSwap is a free data retrieval call binding the contract method 0xf6e1c1c4.
//
// Solidity: function _autoSwap() view returns(address)
func (_DappBamboo *DappBambooSession) AutoSwap() (common.Address, error) {
	return _DappBamboo.Contract.AutoSwap(&_DappBamboo.CallOpts)
}

// AutoSwap is a free data retrieval call binding the contract method 0xf6e1c1c4.
//
// Solidity: function _autoSwap() view returns(address)
func (_DappBamboo *DappBambooCallerSession) AutoSwap() (common.Address, error) {
	return _DappBamboo.Contract.AutoSwap(&_DappBamboo.CallOpts)
}

// DestroyFee is a free data retrieval call binding the contract method 0x964e45f5.
//
// Solidity: function _destroyFee() view returns(uint256)
func (_DappBamboo *DappBambooCaller) DestroyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_destroyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestroyFee is a free data retrieval call binding the contract method 0x964e45f5.
//
// Solidity: function _destroyFee() view returns(uint256)
func (_DappBamboo *DappBambooSession) DestroyFee() (*big.Int, error) {
	return _DappBamboo.Contract.DestroyFee(&_DappBamboo.CallOpts)
}

// DestroyFee is a free data retrieval call binding the contract method 0x964e45f5.
//
// Solidity: function _destroyFee() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) DestroyFee() (*big.Int, error) {
	return _DappBamboo.Contract.DestroyFee(&_DappBamboo.CallOpts)
}

// FundFee is a free data retrieval call binding the contract method 0xd91eb81b.
//
// Solidity: function _fundFee() view returns(uint256)
func (_DappBamboo *DappBambooCaller) FundFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_fundFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundFee is a free data retrieval call binding the contract method 0xd91eb81b.
//
// Solidity: function _fundFee() view returns(uint256)
func (_DappBamboo *DappBambooSession) FundFee() (*big.Int, error) {
	return _DappBamboo.Contract.FundFee(&_DappBamboo.CallOpts)
}

// FundFee is a free data retrieval call binding the contract method 0xd91eb81b.
//
// Solidity: function _fundFee() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) FundFee() (*big.Int, error) {
	return _DappBamboo.Contract.FundFee(&_DappBamboo.CallOpts)
}

// InviterFee is a free data retrieval call binding the contract method 0xa41a08fb.
//
// Solidity: function _inviterFee() view returns(uint256)
func (_DappBamboo *DappBambooCaller) InviterFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_inviterFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviterFee is a free data retrieval call binding the contract method 0xa41a08fb.
//
// Solidity: function _inviterFee() view returns(uint256)
func (_DappBamboo *DappBambooSession) InviterFee() (*big.Int, error) {
	return _DappBamboo.Contract.InviterFee(&_DappBamboo.CallOpts)
}

// InviterFee is a free data retrieval call binding the contract method 0xa41a08fb.
//
// Solidity: function _inviterFee() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) InviterFee() (*big.Int, error) {
	return _DappBamboo.Contract.InviterFee(&_DappBamboo.CallOpts)
}

// LpFee is a free data retrieval call binding the contract method 0x15c93a7d.
//
// Solidity: function _lpFee() view returns(uint256)
func (_DappBamboo *DappBambooCaller) LpFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_lpFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpFee is a free data retrieval call binding the contract method 0x15c93a7d.
//
// Solidity: function _lpFee() view returns(uint256)
func (_DappBamboo *DappBambooSession) LpFee() (*big.Int, error) {
	return _DappBamboo.Contract.LpFee(&_DappBamboo.CallOpts)
}

// LpFee is a free data retrieval call binding the contract method 0x15c93a7d.
//
// Solidity: function _lpFee() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) LpFee() (*big.Int, error) {
	return _DappBamboo.Contract.LpFee(&_DappBamboo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_DappBamboo *DappBambooCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_DappBamboo *DappBambooSession) Owner() (common.Address, error) {
	return _DappBamboo.Contract.Owner(&_DappBamboo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_DappBamboo *DappBambooCallerSession) Owner() (common.Address, error) {
	return _DappBamboo.Contract.Owner(&_DappBamboo.CallOpts)
}

// ROwned is a free data retrieval call binding the contract method 0x0cfc15f9.
//
// Solidity: function _rOwned(address ) view returns(uint256)
func (_DappBamboo *DappBambooCaller) ROwned(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_rOwned", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROwned is a free data retrieval call binding the contract method 0x0cfc15f9.
//
// Solidity: function _rOwned(address ) view returns(uint256)
func (_DappBamboo *DappBambooSession) ROwned(arg0 common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.ROwned(&_DappBamboo.CallOpts, arg0)
}

// ROwned is a free data retrieval call binding the contract method 0x0cfc15f9.
//
// Solidity: function _rOwned(address ) view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) ROwned(arg0 common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.ROwned(&_DappBamboo.CallOpts, arg0)
}

// TaxFee is a free data retrieval call binding the contract method 0x3b124fe7.
//
// Solidity: function _taxFee() view returns(uint256)
func (_DappBamboo *DappBambooCaller) TaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_taxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxFee is a free data retrieval call binding the contract method 0x3b124fe7.
//
// Solidity: function _taxFee() view returns(uint256)
func (_DappBamboo *DappBambooSession) TaxFee() (*big.Int, error) {
	return _DappBamboo.Contract.TaxFee(&_DappBamboo.CallOpts)
}

// TaxFee is a free data retrieval call binding the contract method 0x3b124fe7.
//
// Solidity: function _taxFee() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) TaxFee() (*big.Int, error) {
	return _DappBamboo.Contract.TaxFee(&_DappBamboo.CallOpts)
}

// Updated is a free data retrieval call binding the contract method 0x98d0062c.
//
// Solidity: function _updated(address ) view returns(bool)
func (_DappBamboo *DappBambooCaller) Updated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "_updated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Updated is a free data retrieval call binding the contract method 0x98d0062c.
//
// Solidity: function _updated(address ) view returns(bool)
func (_DappBamboo *DappBambooSession) Updated(arg0 common.Address) (bool, error) {
	return _DappBamboo.Contract.Updated(&_DappBamboo.CallOpts, arg0)
}

// Updated is a free data retrieval call binding the contract method 0x98d0062c.
//
// Solidity: function _updated(address ) view returns(bool)
func (_DappBamboo *DappBambooCallerSession) Updated(arg0 common.Address) (bool, error) {
	return _DappBamboo.Contract.Updated(&_DappBamboo.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DappBamboo *DappBambooCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DappBamboo *DappBambooSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.Allowance(&_DappBamboo.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.Allowance(&_DappBamboo.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DappBamboo *DappBambooCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DappBamboo *DappBambooSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.BalanceOf(&_DappBamboo.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.BalanceOf(&_DappBamboo.CallOpts, account)
}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint256)
func (_DappBamboo *DappBambooCaller) CurrentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "currentIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint256)
func (_DappBamboo *DappBambooSession) CurrentIndex() (*big.Int, error) {
	return _DappBamboo.Contract.CurrentIndex(&_DappBamboo.CallOpts)
}

// CurrentIndex is a free data retrieval call binding the contract method 0x26987b60.
//
// Solidity: function currentIndex() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) CurrentIndex() (*big.Int, error) {
	return _DappBamboo.Contract.CurrentIndex(&_DappBamboo.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_DappBamboo *DappBambooCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_DappBamboo *DappBambooSession) Decimals() (*big.Int, error) {
	return _DappBamboo.Contract.Decimals(&_DappBamboo.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) Decimals() (*big.Int, error) {
	return _DappBamboo.Contract.Decimals(&_DappBamboo.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1)
func (_DappBamboo *DappBambooCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1)
func (_DappBamboo *DappBambooSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _DappBamboo.Contract.GetReserves(&_DappBamboo.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1)
func (_DappBamboo *DappBambooCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _DappBamboo.Contract.GetReserves(&_DappBamboo.CallOpts)
}

// Inviter is a free data retrieval call binding the contract method 0xee8f0b7a.
//
// Solidity: function inviter(address ) view returns(address)
func (_DappBamboo *DappBambooCaller) Inviter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "inviter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inviter is a free data retrieval call binding the contract method 0xee8f0b7a.
//
// Solidity: function inviter(address ) view returns(address)
func (_DappBamboo *DappBambooSession) Inviter(arg0 common.Address) (common.Address, error) {
	return _DappBamboo.Contract.Inviter(&_DappBamboo.CallOpts, arg0)
}

// Inviter is a free data retrieval call binding the contract method 0xee8f0b7a.
//
// Solidity: function inviter(address ) view returns(address)
func (_DappBamboo *DappBambooCallerSession) Inviter(arg0 common.Address) (common.Address, error) {
	return _DappBamboo.Contract.Inviter(&_DappBamboo.CallOpts, arg0)
}

// InviterNum is a free data retrieval call binding the contract method 0xbfa268dd.
//
// Solidity: function inviterNum(address ) view returns(uint256)
func (_DappBamboo *DappBambooCaller) InviterNum(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "inviterNum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviterNum is a free data retrieval call binding the contract method 0xbfa268dd.
//
// Solidity: function inviterNum(address ) view returns(uint256)
func (_DappBamboo *DappBambooSession) InviterNum(arg0 common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.InviterNum(&_DappBamboo.CallOpts, arg0)
}

// InviterNum is a free data retrieval call binding the contract method 0xbfa268dd.
//
// Solidity: function inviterNum(address ) view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) InviterNum(arg0 common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.InviterNum(&_DappBamboo.CallOpts, arg0)
}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_DappBamboo *DappBambooCaller) IsExcludedFromFee(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "isExcludedFromFee", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_DappBamboo *DappBambooSession) IsExcludedFromFee(account common.Address) (bool, error) {
	return _DappBamboo.Contract.IsExcludedFromFee(&_DappBamboo.CallOpts, account)
}

// IsExcludedFromFee is a free data retrieval call binding the contract method 0x5342acb4.
//
// Solidity: function isExcludedFromFee(address account) view returns(bool)
func (_DappBamboo *DappBambooCallerSession) IsExcludedFromFee(account common.Address) (bool, error) {
	return _DappBamboo.Contract.IsExcludedFromFee(&_DappBamboo.CallOpts, account)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_DappBamboo *DappBambooCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_DappBamboo *DappBambooSession) IsOpen() (bool, error) {
	return _DappBamboo.Contract.IsOpen(&_DappBamboo.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_DappBamboo *DappBambooCallerSession) IsOpen() (bool, error) {
	return _DappBamboo.Contract.IsOpen(&_DappBamboo.CallOpts)
}

// IsOpenBuy is a free data retrieval call binding the contract method 0x75fd3758.
//
// Solidity: function isOpenBuy() view returns(bool)
func (_DappBamboo *DappBambooCaller) IsOpenBuy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "isOpenBuy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpenBuy is a free data retrieval call binding the contract method 0x75fd3758.
//
// Solidity: function isOpenBuy() view returns(bool)
func (_DappBamboo *DappBambooSession) IsOpenBuy() (bool, error) {
	return _DappBamboo.Contract.IsOpenBuy(&_DappBamboo.CallOpts)
}

// IsOpenBuy is a free data retrieval call binding the contract method 0x75fd3758.
//
// Solidity: function isOpenBuy() view returns(bool)
func (_DappBamboo *DappBambooCallerSession) IsOpenBuy() (bool, error) {
	return _DappBamboo.Contract.IsOpenBuy(&_DappBamboo.CallOpts)
}

// IsOpensell is a free data retrieval call binding the contract method 0x4bb441d6.
//
// Solidity: function isOpensell() view returns(bool)
func (_DappBamboo *DappBambooCaller) IsOpensell(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "isOpensell")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpensell is a free data retrieval call binding the contract method 0x4bb441d6.
//
// Solidity: function isOpensell() view returns(bool)
func (_DappBamboo *DappBambooSession) IsOpensell() (bool, error) {
	return _DappBamboo.Contract.IsOpensell(&_DappBamboo.CallOpts)
}

// IsOpensell is a free data retrieval call binding the contract method 0x4bb441d6.
//
// Solidity: function isOpensell() view returns(bool)
func (_DappBamboo *DappBambooCallerSession) IsOpensell() (bool, error) {
	return _DappBamboo.Contract.IsOpensell(&_DappBamboo.CallOpts)
}

// IsRoute is a free data retrieval call binding the contract method 0x1702cf66.
//
// Solidity: function isRoute(address ) view returns(bool)
func (_DappBamboo *DappBambooCaller) IsRoute(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "isRoute", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoute is a free data retrieval call binding the contract method 0x1702cf66.
//
// Solidity: function isRoute(address ) view returns(bool)
func (_DappBamboo *DappBambooSession) IsRoute(arg0 common.Address) (bool, error) {
	return _DappBamboo.Contract.IsRoute(&_DappBamboo.CallOpts, arg0)
}

// IsRoute is a free data retrieval call binding the contract method 0x1702cf66.
//
// Solidity: function isRoute(address ) view returns(bool)
func (_DappBamboo *DappBambooCallerSession) IsRoute(arg0 common.Address) (bool, error) {
	return _DappBamboo.Contract.IsRoute(&_DappBamboo.CallOpts, arg0)
}

// IsStop is a free data retrieval call binding the contract method 0xcf359165.
//
// Solidity: function isStop(address ) view returns(bool)
func (_DappBamboo *DappBambooCaller) IsStop(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "isStop", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStop is a free data retrieval call binding the contract method 0xcf359165.
//
// Solidity: function isStop(address ) view returns(bool)
func (_DappBamboo *DappBambooSession) IsStop(arg0 common.Address) (bool, error) {
	return _DappBamboo.Contract.IsStop(&_DappBamboo.CallOpts, arg0)
}

// IsStop is a free data retrieval call binding the contract method 0xcf359165.
//
// Solidity: function isStop(address ) view returns(bool)
func (_DappBamboo *DappBambooCallerSession) IsStop(arg0 common.Address) (bool, error) {
	return _DappBamboo.Contract.IsStop(&_DappBamboo.CallOpts, arg0)
}

// LiquifyEnabled is a free data retrieval call binding the contract method 0xed5792d7.
//
// Solidity: function liquifyEnabled() view returns(bool)
func (_DappBamboo *DappBambooCaller) LiquifyEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "liquifyEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LiquifyEnabled is a free data retrieval call binding the contract method 0xed5792d7.
//
// Solidity: function liquifyEnabled() view returns(bool)
func (_DappBamboo *DappBambooSession) LiquifyEnabled() (bool, error) {
	return _DappBamboo.Contract.LiquifyEnabled(&_DappBamboo.CallOpts)
}

// LiquifyEnabled is a free data retrieval call binding the contract method 0xed5792d7.
//
// Solidity: function liquifyEnabled() view returns(bool)
func (_DappBamboo *DappBambooCallerSession) LiquifyEnabled() (bool, error) {
	return _DappBamboo.Contract.LiquifyEnabled(&_DappBamboo.CallOpts)
}

// MinPeriod is a free data retrieval call binding the contract method 0xffd49c84.
//
// Solidity: function minPeriod() view returns(uint256)
func (_DappBamboo *DappBambooCaller) MinPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "minPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPeriod is a free data retrieval call binding the contract method 0xffd49c84.
//
// Solidity: function minPeriod() view returns(uint256)
func (_DappBamboo *DappBambooSession) MinPeriod() (*big.Int, error) {
	return _DappBamboo.Contract.MinPeriod(&_DappBamboo.CallOpts)
}

// MinPeriod is a free data retrieval call binding the contract method 0xffd49c84.
//
// Solidity: function minPeriod() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) MinPeriod() (*big.Int, error) {
	return _DappBamboo.Contract.MinPeriod(&_DappBamboo.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DappBamboo *DappBambooCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DappBamboo *DappBambooSession) Name() (string, error) {
	return _DappBamboo.Contract.Name(&_DappBamboo.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DappBamboo *DappBambooCallerSession) Name() (string, error) {
	return _DappBamboo.Contract.Name(&_DappBamboo.CallOpts)
}

// NumTokensSellToAddToLiquidity is a free data retrieval call binding the contract method 0xd12a7688.
//
// Solidity: function numTokensSellToAddToLiquidity() view returns(uint256)
func (_DappBamboo *DappBambooCaller) NumTokensSellToAddToLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "numTokensSellToAddToLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensSellToAddToLiquidity is a free data retrieval call binding the contract method 0xd12a7688.
//
// Solidity: function numTokensSellToAddToLiquidity() view returns(uint256)
func (_DappBamboo *DappBambooSession) NumTokensSellToAddToLiquidity() (*big.Int, error) {
	return _DappBamboo.Contract.NumTokensSellToAddToLiquidity(&_DappBamboo.CallOpts)
}

// NumTokensSellToAddToLiquidity is a free data retrieval call binding the contract method 0xd12a7688.
//
// Solidity: function numTokensSellToAddToLiquidity() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) NumTokensSellToAddToLiquidity() (*big.Int, error) {
	return _DappBamboo.Contract.NumTokensSellToAddToLiquidity(&_DappBamboo.CallOpts)
}

// Owner1 is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappBamboo *DappBambooCaller) Owner1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner1 is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappBamboo *DappBambooSession) Owner1() (common.Address, error) {
	return _DappBamboo.Contract.Owner1(&_DappBamboo.CallOpts)
}

// Owner1 is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappBamboo *DappBambooCallerSession) Owner1() (common.Address, error) {
	return _DappBamboo.Contract.Owner1(&_DappBamboo.CallOpts)
}

// ShareholderIndexes is a free data retrieval call binding the contract method 0xd4fda1f2.
//
// Solidity: function shareholderIndexes(address ) view returns(uint256)
func (_DappBamboo *DappBambooCaller) ShareholderIndexes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "shareholderIndexes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareholderIndexes is a free data retrieval call binding the contract method 0xd4fda1f2.
//
// Solidity: function shareholderIndexes(address ) view returns(uint256)
func (_DappBamboo *DappBambooSession) ShareholderIndexes(arg0 common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.ShareholderIndexes(&_DappBamboo.CallOpts, arg0)
}

// ShareholderIndexes is a free data retrieval call binding the contract method 0xd4fda1f2.
//
// Solidity: function shareholderIndexes(address ) view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) ShareholderIndexes(arg0 common.Address) (*big.Int, error) {
	return _DappBamboo.Contract.ShareholderIndexes(&_DappBamboo.CallOpts, arg0)
}

// Shareholders is a free data retrieval call binding the contract method 0xab377daa.
//
// Solidity: function shareholders(uint256 ) view returns(address)
func (_DappBamboo *DappBambooCaller) Shareholders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "shareholders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Shareholders is a free data retrieval call binding the contract method 0xab377daa.
//
// Solidity: function shareholders(uint256 ) view returns(address)
func (_DappBamboo *DappBambooSession) Shareholders(arg0 *big.Int) (common.Address, error) {
	return _DappBamboo.Contract.Shareholders(&_DappBamboo.CallOpts, arg0)
}

// Shareholders is a free data retrieval call binding the contract method 0xab377daa.
//
// Solidity: function shareholders(uint256 ) view returns(address)
func (_DappBamboo *DappBambooCallerSession) Shareholders(arg0 *big.Int) (common.Address, error) {
	return _DappBamboo.Contract.Shareholders(&_DappBamboo.CallOpts, arg0)
}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_DappBamboo *DappBambooCaller) SwapAndLiquifyEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "swapAndLiquifyEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_DappBamboo *DappBambooSession) SwapAndLiquifyEnabled() (bool, error) {
	return _DappBamboo.Contract.SwapAndLiquifyEnabled(&_DappBamboo.CallOpts)
}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_DappBamboo *DappBambooCallerSession) SwapAndLiquifyEnabled() (bool, error) {
	return _DappBamboo.Contract.SwapAndLiquifyEnabled(&_DappBamboo.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DappBamboo *DappBambooCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DappBamboo *DappBambooSession) Symbol() (string, error) {
	return _DappBamboo.Contract.Symbol(&_DappBamboo.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DappBamboo *DappBambooCallerSession) Symbol() (string, error) {
	return _DappBamboo.Contract.Symbol(&_DappBamboo.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_DappBamboo *DappBambooCaller) TotalFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "totalFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_DappBamboo *DappBambooSession) TotalFees() (*big.Int, error) {
	return _DappBamboo.Contract.TotalFees(&_DappBamboo.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) TotalFees() (*big.Int, error) {
	return _DappBamboo.Contract.TotalFees(&_DappBamboo.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DappBamboo *DappBambooCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DappBamboo *DappBambooSession) TotalSupply() (*big.Int, error) {
	return _DappBamboo.Contract.TotalSupply(&_DappBamboo.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DappBamboo *DappBambooCallerSession) TotalSupply() (*big.Int, error) {
	return _DappBamboo.Contract.TotalSupply(&_DappBamboo.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_DappBamboo *DappBambooCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_DappBamboo *DappBambooSession) UniswapV2Pair() (common.Address, error) {
	return _DappBamboo.Contract.UniswapV2Pair(&_DappBamboo.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_DappBamboo *DappBambooCallerSession) UniswapV2Pair() (common.Address, error) {
	return _DappBamboo.Contract.UniswapV2Pair(&_DappBamboo.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_DappBamboo *DappBambooCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappBamboo.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_DappBamboo *DappBambooSession) UniswapV2Router() (common.Address, error) {
	return _DappBamboo.Contract.UniswapV2Router(&_DappBamboo.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_DappBamboo *DappBambooCallerSession) UniswapV2Router() (common.Address, error) {
	return _DappBamboo.Contract.UniswapV2Router(&_DappBamboo.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.Approve(&_DappBamboo.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.Approve(&_DappBamboo.TransactOpts, spender, amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_DappBamboo *DappBambooTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_DappBamboo *DappBambooSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.ChangeOwner(&_DappBamboo.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_DappBamboo *DappBambooTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.ChangeOwner(&_DappBamboo.TransactOpts, newOwner)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_DappBamboo *DappBambooTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_DappBamboo *DappBambooSession) Claim() (*types.Transaction, error) {
	return _DappBamboo.Contract.Claim(&_DappBamboo.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_DappBamboo *DappBambooTransactorSession) Claim() (*types.Transaction, error) {
	return _DappBamboo.Contract.Claim(&_DappBamboo.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x9fc314c8.
//
// Solidity: function claimTokens(address token, address to, uint256 amount) returns()
func (_DappBamboo *DappBambooTransactor) ClaimTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "claimTokens", token, to, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x9fc314c8.
//
// Solidity: function claimTokens(address token, address to, uint256 amount) returns()
func (_DappBamboo *DappBambooSession) ClaimTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.ClaimTokens(&_DappBamboo.TransactOpts, token, to, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x9fc314c8.
//
// Solidity: function claimTokens(address token, address to, uint256 amount) returns()
func (_DappBamboo *DappBambooTransactorSession) ClaimTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.ClaimTokens(&_DappBamboo.TransactOpts, token, to, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DappBamboo *DappBambooTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DappBamboo *DappBambooSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.DecreaseAllowance(&_DappBamboo.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DappBamboo *DappBambooTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.DecreaseAllowance(&_DappBamboo.TransactOpts, spender, subtractedValue)
}

// ExcludeFromFee is a paid mutator transaction binding the contract method 0x437823ec.
//
// Solidity: function excludeFromFee(address account) returns()
func (_DappBamboo *DappBambooTransactor) ExcludeFromFee(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "excludeFromFee", account)
}

// ExcludeFromFee is a paid mutator transaction binding the contract method 0x437823ec.
//
// Solidity: function excludeFromFee(address account) returns()
func (_DappBamboo *DappBambooSession) ExcludeFromFee(account common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.ExcludeFromFee(&_DappBamboo.TransactOpts, account)
}

// ExcludeFromFee is a paid mutator transaction binding the contract method 0x437823ec.
//
// Solidity: function excludeFromFee(address account) returns()
func (_DappBamboo *DappBambooTransactorSession) ExcludeFromFee(account common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.ExcludeFromFee(&_DappBamboo.TransactOpts, account)
}

// IncludeInFee is a paid mutator transaction binding the contract method 0xea2f0b37.
//
// Solidity: function includeInFee(address account) returns()
func (_DappBamboo *DappBambooTransactor) IncludeInFee(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "includeInFee", account)
}

// IncludeInFee is a paid mutator transaction binding the contract method 0xea2f0b37.
//
// Solidity: function includeInFee(address account) returns()
func (_DappBamboo *DappBambooSession) IncludeInFee(account common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.IncludeInFee(&_DappBamboo.TransactOpts, account)
}

// IncludeInFee is a paid mutator transaction binding the contract method 0xea2f0b37.
//
// Solidity: function includeInFee(address account) returns()
func (_DappBamboo *DappBambooTransactorSession) IncludeInFee(account common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.IncludeInFee(&_DappBamboo.TransactOpts, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DappBamboo *DappBambooTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DappBamboo *DappBambooSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.IncreaseAllowance(&_DappBamboo.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DappBamboo *DappBambooTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.IncreaseAllowance(&_DappBamboo.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_DappBamboo *DappBambooTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_DappBamboo *DappBambooSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.Mint(&_DappBamboo.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_DappBamboo *DappBambooTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.Mint(&_DappBamboo.TransactOpts, addr, amount)
}

// SetDes is a paid mutator transaction binding the contract method 0x629e5306.
//
// Solidity: function setDes() returns()
func (_DappBamboo *DappBambooTransactor) SetDes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setDes")
}

// SetDes is a paid mutator transaction binding the contract method 0x629e5306.
//
// Solidity: function setDes() returns()
func (_DappBamboo *DappBambooSession) SetDes() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetDes(&_DappBamboo.TransactOpts)
}

// SetDes is a paid mutator transaction binding the contract method 0x629e5306.
//
// Solidity: function setDes() returns()
func (_DappBamboo *DappBambooTransactorSession) SetDes() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetDes(&_DappBamboo.TransactOpts)
}

// SetFun is a paid mutator transaction binding the contract method 0xca88f11d.
//
// Solidity: function setFun(address addr) returns()
func (_DappBamboo *DappBambooTransactor) SetFun(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setFun", addr)
}

// SetFun is a paid mutator transaction binding the contract method 0xca88f11d.
//
// Solidity: function setFun(address addr) returns()
func (_DappBamboo *DappBambooSession) SetFun(addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetFun(&_DappBamboo.TransactOpts, addr)
}

// SetFun is a paid mutator transaction binding the contract method 0xca88f11d.
//
// Solidity: function setFun(address addr) returns()
func (_DappBamboo *DappBambooTransactorSession) SetFun(addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetFun(&_DappBamboo.TransactOpts, addr)
}

// SetFun2 is a paid mutator transaction binding the contract method 0x98454cbb.
//
// Solidity: function setFun2(address addr) returns()
func (_DappBamboo *DappBambooTransactor) SetFun2(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setFun2", addr)
}

// SetFun2 is a paid mutator transaction binding the contract method 0x98454cbb.
//
// Solidity: function setFun2(address addr) returns()
func (_DappBamboo *DappBambooSession) SetFun2(addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetFun2(&_DappBamboo.TransactOpts, addr)
}

// SetFun2 is a paid mutator transaction binding the contract method 0x98454cbb.
//
// Solidity: function setFun2(address addr) returns()
func (_DappBamboo *DappBambooTransactorSession) SetFun2(addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetFun2(&_DappBamboo.TransactOpts, addr)
}

// SetLiquifyEnabled is a paid mutator transaction binding the contract method 0xda1cdbe4.
//
// Solidity: function setLiquifyEnabled(bool _enabled) returns()
func (_DappBamboo *DappBambooTransactor) SetLiquifyEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setLiquifyEnabled", _enabled)
}

// SetLiquifyEnabled is a paid mutator transaction binding the contract method 0xda1cdbe4.
//
// Solidity: function setLiquifyEnabled(bool _enabled) returns()
func (_DappBamboo *DappBambooSession) SetLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetLiquifyEnabled(&_DappBamboo.TransactOpts, _enabled)
}

// SetLiquifyEnabled is a paid mutator transaction binding the contract method 0xda1cdbe4.
//
// Solidity: function setLiquifyEnabled(bool _enabled) returns()
func (_DappBamboo *DappBambooTransactorSession) SetLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetLiquifyEnabled(&_DappBamboo.TransactOpts, _enabled)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0x7916a5ee.
//
// Solidity: function setMinAdd(uint256 _num) returns()
func (_DappBamboo *DappBambooTransactor) SetMinAdd(opts *bind.TransactOpts, _num *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setMinAdd", _num)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0x7916a5ee.
//
// Solidity: function setMinAdd(uint256 _num) returns()
func (_DappBamboo *DappBambooSession) SetMinAdd(_num *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetMinAdd(&_DappBamboo.TransactOpts, _num)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0x7916a5ee.
//
// Solidity: function setMinAdd(uint256 _num) returns()
func (_DappBamboo *DappBambooTransactorSession) SetMinAdd(_num *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetMinAdd(&_DappBamboo.TransactOpts, _num)
}

// SetOpen is a paid mutator transaction binding the contract method 0x712b7b14.
//
// Solidity: function setOpen() returns()
func (_DappBamboo *DappBambooTransactor) SetOpen(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setOpen")
}

// SetOpen is a paid mutator transaction binding the contract method 0x712b7b14.
//
// Solidity: function setOpen() returns()
func (_DappBamboo *DappBambooSession) SetOpen() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetOpen(&_DappBamboo.TransactOpts)
}

// SetOpen is a paid mutator transaction binding the contract method 0x712b7b14.
//
// Solidity: function setOpen() returns()
func (_DappBamboo *DappBambooTransactorSession) SetOpen() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetOpen(&_DappBamboo.TransactOpts)
}

// SetStop is a paid mutator transaction binding the contract method 0x957138a4.
//
// Solidity: function setStop(address _addr, bool _bool) returns()
func (_DappBamboo *DappBambooTransactor) SetStop(opts *bind.TransactOpts, _addr common.Address, _bool bool) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setStop", _addr, _bool)
}

// SetStop is a paid mutator transaction binding the contract method 0x957138a4.
//
// Solidity: function setStop(address _addr, bool _bool) returns()
func (_DappBamboo *DappBambooSession) SetStop(_addr common.Address, _bool bool) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetStop(&_DappBamboo.TransactOpts, _addr, _bool)
}

// SetStop is a paid mutator transaction binding the contract method 0x957138a4.
//
// Solidity: function setStop(address _addr, bool _bool) returns()
func (_DappBamboo *DappBambooTransactorSession) SetStop(_addr common.Address, _bool bool) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetStop(&_DappBamboo.TransactOpts, _addr, _bool)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_DappBamboo *DappBambooTransactor) SetSwapAndLiquifyEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setSwapAndLiquifyEnabled", _enabled)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_DappBamboo *DappBambooSession) SetSwapAndLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetSwapAndLiquifyEnabled(&_DappBamboo.TransactOpts, _enabled)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_DappBamboo *DappBambooTransactorSession) SetSwapAndLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetSwapAndLiquifyEnabled(&_DappBamboo.TransactOpts, _enabled)
}

// SetTRADE is a paid mutator transaction binding the contract method 0x76f878b0.
//
// Solidity: function setTRADE(address _addr) returns()
func (_DappBamboo *DappBambooTransactor) SetTRADE(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setTRADE", _addr)
}

// SetTRADE is a paid mutator transaction binding the contract method 0x76f878b0.
//
// Solidity: function setTRADE(address _addr) returns()
func (_DappBamboo *DappBambooSession) SetTRADE(_addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetTRADE(&_DappBamboo.TransactOpts, _addr)
}

// SetTRADE is a paid mutator transaction binding the contract method 0x76f878b0.
//
// Solidity: function setTRADE(address _addr) returns()
func (_DappBamboo *DappBambooTransactorSession) SetTRADE(_addr common.Address) (*types.Transaction, error) {
	return _DappBamboo.Contract.SetTRADE(&_DappBamboo.TransactOpts, _addr)
}

// SetisOpenBuy is a paid mutator transaction binding the contract method 0x4a4d9a29.
//
// Solidity: function setisOpenBuy() returns()
func (_DappBamboo *DappBambooTransactor) SetisOpenBuy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setisOpenBuy")
}

// SetisOpenBuy is a paid mutator transaction binding the contract method 0x4a4d9a29.
//
// Solidity: function setisOpenBuy() returns()
func (_DappBamboo *DappBambooSession) SetisOpenBuy() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetisOpenBuy(&_DappBamboo.TransactOpts)
}

// SetisOpenBuy is a paid mutator transaction binding the contract method 0x4a4d9a29.
//
// Solidity: function setisOpenBuy() returns()
func (_DappBamboo *DappBambooTransactorSession) SetisOpenBuy() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetisOpenBuy(&_DappBamboo.TransactOpts)
}

// SetisOpensell is a paid mutator transaction binding the contract method 0xf97eef54.
//
// Solidity: function setisOpensell() returns()
func (_DappBamboo *DappBambooTransactor) SetisOpensell(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "setisOpensell")
}

// SetisOpensell is a paid mutator transaction binding the contract method 0xf97eef54.
//
// Solidity: function setisOpensell() returns()
func (_DappBamboo *DappBambooSession) SetisOpensell() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetisOpensell(&_DappBamboo.TransactOpts)
}

// SetisOpensell is a paid mutator transaction binding the contract method 0xf97eef54.
//
// Solidity: function setisOpensell() returns()
func (_DappBamboo *DappBambooTransactorSession) SetisOpensell() (*types.Transaction, error) {
	return _DappBamboo.Contract.SetisOpensell(&_DappBamboo.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.Transfer(&_DappBamboo.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.Transfer(&_DappBamboo.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.TransferFrom(&_DappBamboo.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DappBamboo *DappBambooTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DappBamboo.Contract.TransferFrom(&_DappBamboo.TransactOpts, sender, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappBamboo *DappBambooTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappBamboo.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappBamboo *DappBambooSession) Receive() (*types.Transaction, error) {
	return _DappBamboo.Contract.Receive(&_DappBamboo.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappBamboo *DappBambooTransactorSession) Receive() (*types.Transaction, error) {
	return _DappBamboo.Contract.Receive(&_DappBamboo.TransactOpts)
}

// DappBambooApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DappBamboo contract.
type DappBambooApprovalIterator struct {
	Event *DappBambooApproval // Event containing the contract specifics and raw log

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
func (it *DappBambooApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappBambooApproval)
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
		it.Event = new(DappBambooApproval)
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
func (it *DappBambooApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappBambooApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappBambooApproval represents a Approval event raised by the DappBamboo contract.
type DappBambooApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DappBamboo *DappBambooFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DappBambooApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DappBamboo.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DappBambooApprovalIterator{contract: _DappBamboo.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DappBamboo *DappBambooFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DappBambooApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DappBamboo.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappBambooApproval)
				if err := _DappBamboo.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DappBamboo *DappBambooFilterer) ParseApproval(log types.Log) (*DappBambooApproval, error) {
	event := new(DappBambooApproval)
	if err := _DappBamboo.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappBambooSwapAndLiquifyIterator is returned from FilterSwapAndLiquify and is used to iterate over the raw logs and unpacked data for SwapAndLiquify events raised by the DappBamboo contract.
type DappBambooSwapAndLiquifyIterator struct {
	Event *DappBambooSwapAndLiquify // Event containing the contract specifics and raw log

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
func (it *DappBambooSwapAndLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappBambooSwapAndLiquify)
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
		it.Event = new(DappBambooSwapAndLiquify)
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
func (it *DappBambooSwapAndLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappBambooSwapAndLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappBambooSwapAndLiquify represents a SwapAndLiquify event raised by the DappBamboo contract.
type DappBambooSwapAndLiquify struct {
	TokensSwapped      *big.Int
	EthReceived        *big.Int
	TokensIntoLiqudity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquify is a free log retrieval operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_DappBamboo *DappBambooFilterer) FilterSwapAndLiquify(opts *bind.FilterOpts) (*DappBambooSwapAndLiquifyIterator, error) {

	logs, sub, err := _DappBamboo.contract.FilterLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return &DappBambooSwapAndLiquifyIterator{contract: _DappBamboo.contract, event: "SwapAndLiquify", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquify is a free log subscription operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_DappBamboo *DappBambooFilterer) WatchSwapAndLiquify(opts *bind.WatchOpts, sink chan<- *DappBambooSwapAndLiquify) (event.Subscription, error) {

	logs, sub, err := _DappBamboo.contract.WatchLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappBambooSwapAndLiquify)
				if err := _DappBamboo.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
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

// ParseSwapAndLiquify is a log parse operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_DappBamboo *DappBambooFilterer) ParseSwapAndLiquify(log types.Log) (*DappBambooSwapAndLiquify, error) {
	event := new(DappBambooSwapAndLiquify)
	if err := _DappBamboo.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappBambooSwapAndLiquifyEnabledUpdatedIterator is returned from FilterSwapAndLiquifyEnabledUpdated and is used to iterate over the raw logs and unpacked data for SwapAndLiquifyEnabledUpdated events raised by the DappBamboo contract.
type DappBambooSwapAndLiquifyEnabledUpdatedIterator struct {
	Event *DappBambooSwapAndLiquifyEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *DappBambooSwapAndLiquifyEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappBambooSwapAndLiquifyEnabledUpdated)
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
		it.Event = new(DappBambooSwapAndLiquifyEnabledUpdated)
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
func (it *DappBambooSwapAndLiquifyEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappBambooSwapAndLiquifyEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappBambooSwapAndLiquifyEnabledUpdated represents a SwapAndLiquifyEnabledUpdated event raised by the DappBamboo contract.
type DappBambooSwapAndLiquifyEnabledUpdated struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquifyEnabledUpdated is a free log retrieval operation binding the contract event 0x53726dfcaf90650aa7eb35524f4d3220f07413c8d6cb404cc8c18bf5591bc159.
//
// Solidity: event SwapAndLiquifyEnabledUpdated(bool enabled)
func (_DappBamboo *DappBambooFilterer) FilterSwapAndLiquifyEnabledUpdated(opts *bind.FilterOpts) (*DappBambooSwapAndLiquifyEnabledUpdatedIterator, error) {

	logs, sub, err := _DappBamboo.contract.FilterLogs(opts, "SwapAndLiquifyEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &DappBambooSwapAndLiquifyEnabledUpdatedIterator{contract: _DappBamboo.contract, event: "SwapAndLiquifyEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquifyEnabledUpdated is a free log subscription operation binding the contract event 0x53726dfcaf90650aa7eb35524f4d3220f07413c8d6cb404cc8c18bf5591bc159.
//
// Solidity: event SwapAndLiquifyEnabledUpdated(bool enabled)
func (_DappBamboo *DappBambooFilterer) WatchSwapAndLiquifyEnabledUpdated(opts *bind.WatchOpts, sink chan<- *DappBambooSwapAndLiquifyEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _DappBamboo.contract.WatchLogs(opts, "SwapAndLiquifyEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappBambooSwapAndLiquifyEnabledUpdated)
				if err := _DappBamboo.contract.UnpackLog(event, "SwapAndLiquifyEnabledUpdated", log); err != nil {
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

// ParseSwapAndLiquifyEnabledUpdated is a log parse operation binding the contract event 0x53726dfcaf90650aa7eb35524f4d3220f07413c8d6cb404cc8c18bf5591bc159.
//
// Solidity: event SwapAndLiquifyEnabledUpdated(bool enabled)
func (_DappBamboo *DappBambooFilterer) ParseSwapAndLiquifyEnabledUpdated(log types.Log) (*DappBambooSwapAndLiquifyEnabledUpdated, error) {
	event := new(DappBambooSwapAndLiquifyEnabledUpdated)
	if err := _DappBamboo.contract.UnpackLog(event, "SwapAndLiquifyEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappBambooTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DappBamboo contract.
type DappBambooTransferIterator struct {
	Event *DappBambooTransfer // Event containing the contract specifics and raw log

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
func (it *DappBambooTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappBambooTransfer)
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
		it.Event = new(DappBambooTransfer)
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
func (it *DappBambooTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappBambooTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappBambooTransfer represents a Transfer event raised by the DappBamboo contract.
type DappBambooTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DappBamboo *DappBambooFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DappBambooTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DappBamboo.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DappBambooTransferIterator{contract: _DappBamboo.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DappBamboo *DappBambooFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DappBambooTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DappBamboo.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappBambooTransfer)
				if err := _DappBamboo.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DappBamboo *DappBambooFilterer) ParseTransfer(log types.Log) (*DappBambooTransfer, error) {
	event := new(DappBambooTransfer)
	if err := _DappBamboo.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
