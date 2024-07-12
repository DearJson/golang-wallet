// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappDlc

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

// DappDlcMetaData contains all meta data concerning the DappDlc contract.
var DappDlcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_wlist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"}],\"name\":\"amountTokenToAmountUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUSDT\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUSDT\",\"type\":\"uint256\"}],\"name\":\"amountUSDTToAmountToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"cmain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenaddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ctoken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotal\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"amountToken\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"amountUSDT\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"total\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAdd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAddLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBuy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRemove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair_USDT\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Pair\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rRate3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setIsBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setRRate3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setRate3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setisAddLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setisSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setremove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"succ\",\"type\":\"bool\"}],\"name\":\"setwlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DappDlcABI is the input ABI used to generate the binding from.
// Deprecated: Use DappDlcMetaData.ABI instead.
var DappDlcABI = DappDlcMetaData.ABI

// DappDlc is an auto generated Go binding around an Ethereum contract.
type DappDlc struct {
	DappDlcCaller     // Read-only binding to the contract
	DappDlcTransactor // Write-only binding to the contract
	DappDlcFilterer   // Log filterer for contract events
}

// DappDlcCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappDlcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappDlcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappDlcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappDlcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappDlcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappDlcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappDlcSession struct {
	Contract     *DappDlc          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappDlcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappDlcCallerSession struct {
	Contract *DappDlcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DappDlcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappDlcTransactorSession struct {
	Contract     *DappDlcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DappDlcRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappDlcRaw struct {
	Contract *DappDlc // Generic contract binding to access the raw methods on
}

// DappDlcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappDlcCallerRaw struct {
	Contract *DappDlcCaller // Generic read-only contract binding to access the raw methods on
}

// DappDlcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappDlcTransactorRaw struct {
	Contract *DappDlcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappDlc creates a new instance of DappDlc, bound to a specific deployed contract.
func NewDappDlc(address common.Address, backend bind.ContractBackend) (*DappDlc, error) {
	contract, err := bindDappDlc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappDlc{DappDlcCaller: DappDlcCaller{contract: contract}, DappDlcTransactor: DappDlcTransactor{contract: contract}, DappDlcFilterer: DappDlcFilterer{contract: contract}}, nil
}

// NewDappDlcCaller creates a new read-only instance of DappDlc, bound to a specific deployed contract.
func NewDappDlcCaller(address common.Address, caller bind.ContractCaller) (*DappDlcCaller, error) {
	contract, err := bindDappDlc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappDlcCaller{contract: contract}, nil
}

// NewDappDlcTransactor creates a new write-only instance of DappDlc, bound to a specific deployed contract.
func NewDappDlcTransactor(address common.Address, transactor bind.ContractTransactor) (*DappDlcTransactor, error) {
	contract, err := bindDappDlc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappDlcTransactor{contract: contract}, nil
}

// NewDappDlcFilterer creates a new log filterer instance of DappDlc, bound to a specific deployed contract.
func NewDappDlcFilterer(address common.Address, filterer bind.ContractFilterer) (*DappDlcFilterer, error) {
	contract, err := bindDappDlc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappDlcFilterer{contract: contract}, nil
}

// bindDappDlc binds a generic wrapper to an already deployed contract.
func bindDappDlc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappDlcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappDlc *DappDlcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappDlc.Contract.DappDlcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappDlc *DappDlcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.Contract.DappDlcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappDlc *DappDlcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappDlc.Contract.DappDlcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappDlc *DappDlcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappDlc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappDlc *DappDlcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappDlc *DappDlcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappDlc.Contract.contract.Transact(opts, method, params...)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_DappDlc *DappDlcCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_DappDlc *DappDlcSession) TOKEN() (common.Address, error) {
	return _DappDlc.Contract.TOKEN(&_DappDlc.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_DappDlc *DappDlcCallerSession) TOKEN() (common.Address, error) {
	return _DappDlc.Contract.TOKEN(&_DappDlc.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_DappDlc *DappDlcCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_DappDlc *DappDlcSession) USDT() (common.Address, error) {
	return _DappDlc.Contract.USDT(&_DappDlc.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_DappDlc *DappDlcCallerSession) USDT() (common.Address, error) {
	return _DappDlc.Contract.USDT(&_DappDlc.CallOpts)
}

// Wlist is a free data retrieval call binding the contract method 0xe49008b4.
//
// Solidity: function _wlist(address ) view returns(bool)
func (_DappDlc *DappDlcCaller) Wlist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "_wlist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Wlist is a free data retrieval call binding the contract method 0xe49008b4.
//
// Solidity: function _wlist(address ) view returns(bool)
func (_DappDlc *DappDlcSession) Wlist(arg0 common.Address) (bool, error) {
	return _DappDlc.Contract.Wlist(&_DappDlc.CallOpts, arg0)
}

// Wlist is a free data retrieval call binding the contract method 0xe49008b4.
//
// Solidity: function _wlist(address ) view returns(bool)
func (_DappDlc *DappDlcCallerSession) Wlist(arg0 common.Address) (bool, error) {
	return _DappDlc.Contract.Wlist(&_DappDlc.CallOpts, arg0)
}

// AmountTokenToAmountUSDT is a free data retrieval call binding the contract method 0x23efb18d.
//
// Solidity: function amountTokenToAmountUSDT(uint256 amountToken) view returns(uint256 amountUSDT)
func (_DappDlc *DappDlcCaller) AmountTokenToAmountUSDT(opts *bind.CallOpts, amountToken *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "amountTokenToAmountUSDT", amountToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountTokenToAmountUSDT is a free data retrieval call binding the contract method 0x23efb18d.
//
// Solidity: function amountTokenToAmountUSDT(uint256 amountToken) view returns(uint256 amountUSDT)
func (_DappDlc *DappDlcSession) AmountTokenToAmountUSDT(amountToken *big.Int) (*big.Int, error) {
	return _DappDlc.Contract.AmountTokenToAmountUSDT(&_DappDlc.CallOpts, amountToken)
}

// AmountTokenToAmountUSDT is a free data retrieval call binding the contract method 0x23efb18d.
//
// Solidity: function amountTokenToAmountUSDT(uint256 amountToken) view returns(uint256 amountUSDT)
func (_DappDlc *DappDlcCallerSession) AmountTokenToAmountUSDT(amountToken *big.Int) (*big.Int, error) {
	return _DappDlc.Contract.AmountTokenToAmountUSDT(&_DappDlc.CallOpts, amountToken)
}

// AmountUSDTToAmountToken is a free data retrieval call binding the contract method 0xe9239937.
//
// Solidity: function amountUSDTToAmountToken(uint256 amountUSDT) view returns(uint256 amountToken)
func (_DappDlc *DappDlcCaller) AmountUSDTToAmountToken(opts *bind.CallOpts, amountUSDT *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "amountUSDTToAmountToken", amountUSDT)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountUSDTToAmountToken is a free data retrieval call binding the contract method 0xe9239937.
//
// Solidity: function amountUSDTToAmountToken(uint256 amountUSDT) view returns(uint256 amountToken)
func (_DappDlc *DappDlcSession) AmountUSDTToAmountToken(amountUSDT *big.Int) (*big.Int, error) {
	return _DappDlc.Contract.AmountUSDTToAmountToken(&_DappDlc.CallOpts, amountUSDT)
}

// AmountUSDTToAmountToken is a free data retrieval call binding the contract method 0xe9239937.
//
// Solidity: function amountUSDTToAmountToken(uint256 amountUSDT) view returns(uint256 amountToken)
func (_DappDlc *DappDlcCallerSession) AmountUSDTToAmountToken(amountUSDT *big.Int) (*big.Int, error) {
	return _DappDlc.Contract.AmountUSDTToAmountToken(&_DappDlc.CallOpts, amountUSDT)
}

// GetPricR is a free data retrieval call binding the contract method 0x0198581e.
//
// Solidity: function getPricR() view returns(uint256)
func (_DappDlc *DappDlcCaller) GetPricR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "getPricR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricR is a free data retrieval call binding the contract method 0x0198581e.
//
// Solidity: function getPricR() view returns(uint256)
func (_DappDlc *DappDlcSession) GetPricR() (*big.Int, error) {
	return _DappDlc.Contract.GetPricR(&_DappDlc.CallOpts)
}

// GetPricR is a free data retrieval call binding the contract method 0x0198581e.
//
// Solidity: function getPricR() view returns(uint256)
func (_DappDlc *DappDlcCallerSession) GetPricR() (*big.Int, error) {
	return _DappDlc.Contract.GetPricR(&_DappDlc.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_DappDlc *DappDlcCaller) GetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_DappDlc *DappDlcSession) GetPrice() (*big.Int, error) {
	return _DappDlc.Contract.GetPrice(&_DappDlc.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_DappDlc *DappDlcCallerSession) GetPrice() (*big.Int, error) {
	return _DappDlc.Contract.GetPrice(&_DappDlc.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1)
func (_DappDlc *DappDlcCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "getReserves")

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
func (_DappDlc *DappDlcSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _DappDlc.Contract.GetReserves(&_DappDlc.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1)
func (_DappDlc *DappDlcCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _DappDlc.Contract.GetReserves(&_DappDlc.CallOpts)
}

// GetTotal is a free data retrieval call binding the contract method 0x775a25e3.
//
// Solidity: function getTotal() view returns(uint112 amountToken, uint112 amountUSDT, uint112 total)
func (_DappDlc *DappDlcCaller) GetTotal(opts *bind.CallOpts) (struct {
	AmountToken *big.Int
	AmountUSDT  *big.Int
	Total       *big.Int
}, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "getTotal")

	outstruct := new(struct {
		AmountToken *big.Int
		AmountUSDT  *big.Int
		Total       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountToken = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountUSDT = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Total = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTotal is a free data retrieval call binding the contract method 0x775a25e3.
//
// Solidity: function getTotal() view returns(uint112 amountToken, uint112 amountUSDT, uint112 total)
func (_DappDlc *DappDlcSession) GetTotal() (struct {
	AmountToken *big.Int
	AmountUSDT  *big.Int
	Total       *big.Int
}, error) {
	return _DappDlc.Contract.GetTotal(&_DappDlc.CallOpts)
}

// GetTotal is a free data retrieval call binding the contract method 0x775a25e3.
//
// Solidity: function getTotal() view returns(uint112 amountToken, uint112 amountUSDT, uint112 total)
func (_DappDlc *DappDlcCallerSession) GetTotal() (struct {
	AmountToken *big.Int
	AmountUSDT  *big.Int
	Total       *big.Int
}, error) {
	return _DappDlc.Contract.GetTotal(&_DappDlc.CallOpts)
}

// IsAdd is a free data retrieval call binding the contract method 0x2151a9f0.
//
// Solidity: function isAdd() view returns(bool)
func (_DappDlc *DappDlcCaller) IsAdd(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "isAdd")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdd is a free data retrieval call binding the contract method 0x2151a9f0.
//
// Solidity: function isAdd() view returns(bool)
func (_DappDlc *DappDlcSession) IsAdd() (bool, error) {
	return _DappDlc.Contract.IsAdd(&_DappDlc.CallOpts)
}

// IsAdd is a free data retrieval call binding the contract method 0x2151a9f0.
//
// Solidity: function isAdd() view returns(bool)
func (_DappDlc *DappDlcCallerSession) IsAdd() (bool, error) {
	return _DappDlc.Contract.IsAdd(&_DappDlc.CallOpts)
}

// IsAddLiquidity is a free data retrieval call binding the contract method 0x7515dfc6.
//
// Solidity: function isAddLiquidity() view returns(bool)
func (_DappDlc *DappDlcCaller) IsAddLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "isAddLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAddLiquidity is a free data retrieval call binding the contract method 0x7515dfc6.
//
// Solidity: function isAddLiquidity() view returns(bool)
func (_DappDlc *DappDlcSession) IsAddLiquidity() (bool, error) {
	return _DappDlc.Contract.IsAddLiquidity(&_DappDlc.CallOpts)
}

// IsAddLiquidity is a free data retrieval call binding the contract method 0x7515dfc6.
//
// Solidity: function isAddLiquidity() view returns(bool)
func (_DappDlc *DappDlcCallerSession) IsAddLiquidity() (bool, error) {
	return _DappDlc.Contract.IsAddLiquidity(&_DappDlc.CallOpts)
}

// IsBuy is a free data retrieval call binding the contract method 0x4fb90848.
//
// Solidity: function isBuy() view returns(bool)
func (_DappDlc *DappDlcCaller) IsBuy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "isBuy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuy is a free data retrieval call binding the contract method 0x4fb90848.
//
// Solidity: function isBuy() view returns(bool)
func (_DappDlc *DappDlcSession) IsBuy() (bool, error) {
	return _DappDlc.Contract.IsBuy(&_DappDlc.CallOpts)
}

// IsBuy is a free data retrieval call binding the contract method 0x4fb90848.
//
// Solidity: function isBuy() view returns(bool)
func (_DappDlc *DappDlcCallerSession) IsBuy() (bool, error) {
	return _DappDlc.Contract.IsBuy(&_DappDlc.CallOpts)
}

// IsRemove is a free data retrieval call binding the contract method 0x667d9565.
//
// Solidity: function isRemove() view returns(bool)
func (_DappDlc *DappDlcCaller) IsRemove(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "isRemove")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRemove is a free data retrieval call binding the contract method 0x667d9565.
//
// Solidity: function isRemove() view returns(bool)
func (_DappDlc *DappDlcSession) IsRemove() (bool, error) {
	return _DappDlc.Contract.IsRemove(&_DappDlc.CallOpts)
}

// IsRemove is a free data retrieval call binding the contract method 0x667d9565.
//
// Solidity: function isRemove() view returns(bool)
func (_DappDlc *DappDlcCallerSession) IsRemove() (bool, error) {
	return _DappDlc.Contract.IsRemove(&_DappDlc.CallOpts)
}

// IsSell is a free data retrieval call binding the contract method 0x0c570c86.
//
// Solidity: function isSell() view returns(bool)
func (_DappDlc *DappDlcCaller) IsSell(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "isSell")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSell is a free data retrieval call binding the contract method 0x0c570c86.
//
// Solidity: function isSell() view returns(bool)
func (_DappDlc *DappDlcSession) IsSell() (bool, error) {
	return _DappDlc.Contract.IsSell(&_DappDlc.CallOpts)
}

// IsSell is a free data retrieval call binding the contract method 0x0c570c86.
//
// Solidity: function isSell() view returns(bool)
func (_DappDlc *DappDlcCallerSession) IsSell() (bool, error) {
	return _DappDlc.Contract.IsSell(&_DappDlc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappDlc *DappDlcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappDlc *DappDlcSession) Owner() (common.Address, error) {
	return _DappDlc.Contract.Owner(&_DappDlc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappDlc *DappDlcCallerSession) Owner() (common.Address, error) {
	return _DappDlc.Contract.Owner(&_DappDlc.CallOpts)
}

// PairUSDT is a free data retrieval call binding the contract method 0xd466c295.
//
// Solidity: function pair_USDT() view returns(address)
func (_DappDlc *DappDlcCaller) PairUSDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "pair_USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairUSDT is a free data retrieval call binding the contract method 0xd466c295.
//
// Solidity: function pair_USDT() view returns(address)
func (_DappDlc *DappDlcSession) PairUSDT() (common.Address, error) {
	return _DappDlc.Contract.PairUSDT(&_DappDlc.CallOpts)
}

// PairUSDT is a free data retrieval call binding the contract method 0xd466c295.
//
// Solidity: function pair_USDT() view returns(address)
func (_DappDlc *DappDlcCallerSession) PairUSDT() (common.Address, error) {
	return _DappDlc.Contract.PairUSDT(&_DappDlc.CallOpts)
}

// RRate3 is a free data retrieval call binding the contract method 0xb3401328.
//
// Solidity: function rRate3() view returns(uint256)
func (_DappDlc *DappDlcCaller) RRate3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "rRate3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RRate3 is a free data retrieval call binding the contract method 0xb3401328.
//
// Solidity: function rRate3() view returns(uint256)
func (_DappDlc *DappDlcSession) RRate3() (*big.Int, error) {
	return _DappDlc.Contract.RRate3(&_DappDlc.CallOpts)
}

// RRate3 is a free data retrieval call binding the contract method 0xb3401328.
//
// Solidity: function rRate3() view returns(uint256)
func (_DappDlc *DappDlcCallerSession) RRate3() (*big.Int, error) {
	return _DappDlc.Contract.RRate3(&_DappDlc.CallOpts)
}

// Rate3 is a free data retrieval call binding the contract method 0xaff1f15f.
//
// Solidity: function rate3() view returns(uint256)
func (_DappDlc *DappDlcCaller) Rate3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "rate3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate3 is a free data retrieval call binding the contract method 0xaff1f15f.
//
// Solidity: function rate3() view returns(uint256)
func (_DappDlc *DappDlcSession) Rate3() (*big.Int, error) {
	return _DappDlc.Contract.Rate3(&_DappDlc.CallOpts)
}

// Rate3 is a free data retrieval call binding the contract method 0xaff1f15f.
//
// Solidity: function rate3() view returns(uint256)
func (_DappDlc *DappDlcCallerSession) Rate3() (*big.Int, error) {
	return _DappDlc.Contract.Rate3(&_DappDlc.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DappDlc *DappDlcCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DappDlc *DappDlcSession) StartTime() (*big.Int, error) {
	return _DappDlc.Contract.StartTime(&_DappDlc.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DappDlc *DappDlcCallerSession) StartTime() (*big.Int, error) {
	return _DappDlc.Contract.StartTime(&_DappDlc.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_DappDlc *DappDlcCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_DappDlc *DappDlcSession) UniswapV2Pair() (common.Address, error) {
	return _DappDlc.Contract.UniswapV2Pair(&_DappDlc.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_DappDlc *DappDlcCallerSession) UniswapV2Pair() (common.Address, error) {
	return _DappDlc.Contract.UniswapV2Pair(&_DappDlc.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_DappDlc *DappDlcCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappDlc.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_DappDlc *DappDlcSession) UniswapV2Router() (common.Address, error) {
	return _DappDlc.Contract.UniswapV2Router(&_DappDlc.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_DappDlc *DappDlcCallerSession) UniswapV2Router() (common.Address, error) {
	return _DappDlc.Contract.UniswapV2Router(&_DappDlc.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tokenAmount, uint256 usdtAmount) returns()
func (_DappDlc *DappDlcTransactor) AddLiquidity(opts *bind.TransactOpts, tokenAmount *big.Int, usdtAmount *big.Int) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "addLiquidity", tokenAmount, usdtAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tokenAmount, uint256 usdtAmount) returns()
func (_DappDlc *DappDlcSession) AddLiquidity(tokenAmount *big.Int, usdtAmount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.AddLiquidity(&_DappDlc.TransactOpts, tokenAmount, usdtAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 tokenAmount, uint256 usdtAmount) returns()
func (_DappDlc *DappDlcTransactorSession) AddLiquidity(tokenAmount *big.Int, usdtAmount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.AddLiquidity(&_DappDlc.TransactOpts, tokenAmount, usdtAmount)
}

// Buy is a paid mutator transaction binding the contract method 0x7deb6025.
//
// Solidity: function buy(uint256 usdtAmount, address to) returns()
func (_DappDlc *DappDlcTransactor) Buy(opts *bind.TransactOpts, usdtAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "buy", usdtAmount, to)
}

// Buy is a paid mutator transaction binding the contract method 0x7deb6025.
//
// Solidity: function buy(uint256 usdtAmount, address to) returns()
func (_DappDlc *DappDlcSession) Buy(usdtAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.Buy(&_DappDlc.TransactOpts, usdtAmount, to)
}

// Buy is a paid mutator transaction binding the contract method 0x7deb6025.
//
// Solidity: function buy(uint256 usdtAmount, address to) returns()
func (_DappDlc *DappDlcTransactorSession) Buy(usdtAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.Buy(&_DappDlc.TransactOpts, usdtAmount, to)
}

// Cmain is a paid mutator transaction binding the contract method 0x4fe8bd39.
//
// Solidity: function cmain(address _addr) returns()
func (_DappDlc *DappDlcTransactor) Cmain(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "cmain", _addr)
}

// Cmain is a paid mutator transaction binding the contract method 0x4fe8bd39.
//
// Solidity: function cmain(address _addr) returns()
func (_DappDlc *DappDlcSession) Cmain(_addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.Cmain(&_DappDlc.TransactOpts, _addr)
}

// Cmain is a paid mutator transaction binding the contract method 0x4fe8bd39.
//
// Solidity: function cmain(address _addr) returns()
func (_DappDlc *DappDlcTransactorSession) Cmain(_addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.Cmain(&_DappDlc.TransactOpts, _addr)
}

// Ctoken is a paid mutator transaction binding the contract method 0xaf804de1.
//
// Solidity: function ctoken(address _tokenaddr, address _addr, uint256 _amount) returns()
func (_DappDlc *DappDlcTransactor) Ctoken(opts *bind.TransactOpts, _tokenaddr common.Address, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "ctoken", _tokenaddr, _addr, _amount)
}

// Ctoken is a paid mutator transaction binding the contract method 0xaf804de1.
//
// Solidity: function ctoken(address _tokenaddr, address _addr, uint256 _amount) returns()
func (_DappDlc *DappDlcSession) Ctoken(_tokenaddr common.Address, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.Ctoken(&_DappDlc.TransactOpts, _tokenaddr, _addr, _amount)
}

// Ctoken is a paid mutator transaction binding the contract method 0xaf804de1.
//
// Solidity: function ctoken(address _tokenaddr, address _addr, uint256 _amount) returns()
func (_DappDlc *DappDlcTransactorSession) Ctoken(_tokenaddr common.Address, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.Ctoken(&_DappDlc.TransactOpts, _tokenaddr, _addr, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpAmount) returns()
func (_DappDlc *DappDlcTransactor) RemoveLiquidity(opts *bind.TransactOpts, lpAmount *big.Int) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "removeLiquidity", lpAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpAmount) returns()
func (_DappDlc *DappDlcSession) RemoveLiquidity(lpAmount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.RemoveLiquidity(&_DappDlc.TransactOpts, lpAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpAmount) returns()
func (_DappDlc *DappDlcTransactorSession) RemoveLiquidity(lpAmount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.RemoveLiquidity(&_DappDlc.TransactOpts, lpAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappDlc *DappDlcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappDlc *DappDlcSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappDlc.Contract.RenounceOwnership(&_DappDlc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappDlc *DappDlcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappDlc.Contract.RenounceOwnership(&_DappDlc.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x4189a68e.
//
// Solidity: function sell(uint256 tokenAmount, address to) returns()
func (_DappDlc *DappDlcTransactor) Sell(opts *bind.TransactOpts, tokenAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "sell", tokenAmount, to)
}

// Sell is a paid mutator transaction binding the contract method 0x4189a68e.
//
// Solidity: function sell(uint256 tokenAmount, address to) returns()
func (_DappDlc *DappDlcSession) Sell(tokenAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.Sell(&_DappDlc.TransactOpts, tokenAmount, to)
}

// Sell is a paid mutator transaction binding the contract method 0x4189a68e.
//
// Solidity: function sell(uint256 tokenAmount, address to) returns()
func (_DappDlc *DappDlcTransactorSession) Sell(tokenAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.Sell(&_DappDlc.TransactOpts, tokenAmount, to)
}

// SetAdd is a paid mutator transaction binding the contract method 0x22cc65c6.
//
// Solidity: function setAdd() returns()
func (_DappDlc *DappDlcTransactor) SetAdd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setAdd")
}

// SetAdd is a paid mutator transaction binding the contract method 0x22cc65c6.
//
// Solidity: function setAdd() returns()
func (_DappDlc *DappDlcSession) SetAdd() (*types.Transaction, error) {
	return _DappDlc.Contract.SetAdd(&_DappDlc.TransactOpts)
}

// SetAdd is a paid mutator transaction binding the contract method 0x22cc65c6.
//
// Solidity: function setAdd() returns()
func (_DappDlc *DappDlcTransactorSession) SetAdd() (*types.Transaction, error) {
	return _DappDlc.Contract.SetAdd(&_DappDlc.TransactOpts)
}

// SetIsBuy is a paid mutator transaction binding the contract method 0x71f5773a.
//
// Solidity: function setIsBuy() returns()
func (_DappDlc *DappDlcTransactor) SetIsBuy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setIsBuy")
}

// SetIsBuy is a paid mutator transaction binding the contract method 0x71f5773a.
//
// Solidity: function setIsBuy() returns()
func (_DappDlc *DappDlcSession) SetIsBuy() (*types.Transaction, error) {
	return _DappDlc.Contract.SetIsBuy(&_DappDlc.TransactOpts)
}

// SetIsBuy is a paid mutator transaction binding the contract method 0x71f5773a.
//
// Solidity: function setIsBuy() returns()
func (_DappDlc *DappDlcTransactorSession) SetIsBuy() (*types.Transaction, error) {
	return _DappDlc.Contract.SetIsBuy(&_DappDlc.TransactOpts)
}

// SetPair is a paid mutator transaction binding the contract method 0x8187f516.
//
// Solidity: function setPair(address addr) returns()
func (_DappDlc *DappDlcTransactor) SetPair(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setPair", addr)
}

// SetPair is a paid mutator transaction binding the contract method 0x8187f516.
//
// Solidity: function setPair(address addr) returns()
func (_DappDlc *DappDlcSession) SetPair(addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.SetPair(&_DappDlc.TransactOpts, addr)
}

// SetPair is a paid mutator transaction binding the contract method 0x8187f516.
//
// Solidity: function setPair(address addr) returns()
func (_DappDlc *DappDlcTransactorSession) SetPair(addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.SetPair(&_DappDlc.TransactOpts, addr)
}

// SetRRate3 is a paid mutator transaction binding the contract method 0x409128d7.
//
// Solidity: function setRRate3(uint256 _amount) returns()
func (_DappDlc *DappDlcTransactor) SetRRate3(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setRRate3", _amount)
}

// SetRRate3 is a paid mutator transaction binding the contract method 0x409128d7.
//
// Solidity: function setRRate3(uint256 _amount) returns()
func (_DappDlc *DappDlcSession) SetRRate3(_amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.SetRRate3(&_DappDlc.TransactOpts, _amount)
}

// SetRRate3 is a paid mutator transaction binding the contract method 0x409128d7.
//
// Solidity: function setRRate3(uint256 _amount) returns()
func (_DappDlc *DappDlcTransactorSession) SetRRate3(_amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.SetRRate3(&_DappDlc.TransactOpts, _amount)
}

// SetRate3 is a paid mutator transaction binding the contract method 0xf3728b67.
//
// Solidity: function setRate3(uint256 _amount) returns()
func (_DappDlc *DappDlcTransactor) SetRate3(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setRate3", _amount)
}

// SetRate3 is a paid mutator transaction binding the contract method 0xf3728b67.
//
// Solidity: function setRate3(uint256 _amount) returns()
func (_DappDlc *DappDlcSession) SetRate3(_amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.SetRate3(&_DappDlc.TransactOpts, _amount)
}

// SetRate3 is a paid mutator transaction binding the contract method 0xf3728b67.
//
// Solidity: function setRate3(uint256 _amount) returns()
func (_DappDlc *DappDlcTransactorSession) SetRate3(_amount *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.SetRate3(&_DappDlc.TransactOpts, _amount)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns()
func (_DappDlc *DappDlcTransactor) SetStartTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setStartTime", _time)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns()
func (_DappDlc *DappDlcSession) SetStartTime(_time *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.SetStartTime(&_DappDlc.TransactOpts, _time)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns()
func (_DappDlc *DappDlcTransactorSession) SetStartTime(_time *big.Int) (*types.Transaction, error) {
	return _DappDlc.Contract.SetStartTime(&_DappDlc.TransactOpts, _time)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address addr) returns()
func (_DappDlc *DappDlcTransactor) SetToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setToken", addr)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address addr) returns()
func (_DappDlc *DappDlcSession) SetToken(addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.SetToken(&_DappDlc.TransactOpts, addr)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address addr) returns()
func (_DappDlc *DappDlcTransactorSession) SetToken(addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.SetToken(&_DappDlc.TransactOpts, addr)
}

// SetUsdt is a paid mutator transaction binding the contract method 0x58979bfe.
//
// Solidity: function setUsdt(address addr) returns()
func (_DappDlc *DappDlcTransactor) SetUsdt(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setUsdt", addr)
}

// SetUsdt is a paid mutator transaction binding the contract method 0x58979bfe.
//
// Solidity: function setUsdt(address addr) returns()
func (_DappDlc *DappDlcSession) SetUsdt(addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.SetUsdt(&_DappDlc.TransactOpts, addr)
}

// SetUsdt is a paid mutator transaction binding the contract method 0x58979bfe.
//
// Solidity: function setUsdt(address addr) returns()
func (_DappDlc *DappDlcTransactorSession) SetUsdt(addr common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.SetUsdt(&_DappDlc.TransactOpts, addr)
}

// SetisAddLiquidity is a paid mutator transaction binding the contract method 0x63bf2c75.
//
// Solidity: function setisAddLiquidity() returns()
func (_DappDlc *DappDlcTransactor) SetisAddLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setisAddLiquidity")
}

// SetisAddLiquidity is a paid mutator transaction binding the contract method 0x63bf2c75.
//
// Solidity: function setisAddLiquidity() returns()
func (_DappDlc *DappDlcSession) SetisAddLiquidity() (*types.Transaction, error) {
	return _DappDlc.Contract.SetisAddLiquidity(&_DappDlc.TransactOpts)
}

// SetisAddLiquidity is a paid mutator transaction binding the contract method 0x63bf2c75.
//
// Solidity: function setisAddLiquidity() returns()
func (_DappDlc *DappDlcTransactorSession) SetisAddLiquidity() (*types.Transaction, error) {
	return _DappDlc.Contract.SetisAddLiquidity(&_DappDlc.TransactOpts)
}

// SetisSell is a paid mutator transaction binding the contract method 0x17916c53.
//
// Solidity: function setisSell() returns()
func (_DappDlc *DappDlcTransactor) SetisSell(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setisSell")
}

// SetisSell is a paid mutator transaction binding the contract method 0x17916c53.
//
// Solidity: function setisSell() returns()
func (_DappDlc *DappDlcSession) SetisSell() (*types.Transaction, error) {
	return _DappDlc.Contract.SetisSell(&_DappDlc.TransactOpts)
}

// SetisSell is a paid mutator transaction binding the contract method 0x17916c53.
//
// Solidity: function setisSell() returns()
func (_DappDlc *DappDlcTransactorSession) SetisSell() (*types.Transaction, error) {
	return _DappDlc.Contract.SetisSell(&_DappDlc.TransactOpts)
}

// Setremove is a paid mutator transaction binding the contract method 0x2acfc878.
//
// Solidity: function setremove() returns()
func (_DappDlc *DappDlcTransactor) Setremove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setremove")
}

// Setremove is a paid mutator transaction binding the contract method 0x2acfc878.
//
// Solidity: function setremove() returns()
func (_DappDlc *DappDlcSession) Setremove() (*types.Transaction, error) {
	return _DappDlc.Contract.Setremove(&_DappDlc.TransactOpts)
}

// Setremove is a paid mutator transaction binding the contract method 0x2acfc878.
//
// Solidity: function setremove() returns()
func (_DappDlc *DappDlcTransactorSession) Setremove() (*types.Transaction, error) {
	return _DappDlc.Contract.Setremove(&_DappDlc.TransactOpts)
}

// Setwlist is a paid mutator transaction binding the contract method 0xf2031267.
//
// Solidity: function setwlist(address account, bool succ) returns()
func (_DappDlc *DappDlcTransactor) Setwlist(opts *bind.TransactOpts, account common.Address, succ bool) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "setwlist", account, succ)
}

// Setwlist is a paid mutator transaction binding the contract method 0xf2031267.
//
// Solidity: function setwlist(address account, bool succ) returns()
func (_DappDlc *DappDlcSession) Setwlist(account common.Address, succ bool) (*types.Transaction, error) {
	return _DappDlc.Contract.Setwlist(&_DappDlc.TransactOpts, account, succ)
}

// Setwlist is a paid mutator transaction binding the contract method 0xf2031267.
//
// Solidity: function setwlist(address account, bool succ) returns()
func (_DappDlc *DappDlcTransactorSession) Setwlist(account common.Address, succ bool) (*types.Transaction, error) {
	return _DappDlc.Contract.Setwlist(&_DappDlc.TransactOpts, account, succ)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappDlc *DappDlcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappDlc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappDlc *DappDlcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.TransferOwnership(&_DappDlc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappDlc *DappDlcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappDlc.Contract.TransferOwnership(&_DappDlc.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappDlc *DappDlcTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappDlc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappDlc *DappDlcSession) Receive() (*types.Transaction, error) {
	return _DappDlc.Contract.Receive(&_DappDlc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DappDlc *DappDlcTransactorSession) Receive() (*types.Transaction, error) {
	return _DappDlc.Contract.Receive(&_DappDlc.TransactOpts)
}

// DappDlcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DappDlc contract.
type DappDlcOwnershipTransferredIterator struct {
	Event *DappDlcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DappDlcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappDlcOwnershipTransferred)
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
		it.Event = new(DappDlcOwnershipTransferred)
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
func (it *DappDlcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappDlcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappDlcOwnershipTransferred represents a OwnershipTransferred event raised by the DappDlc contract.
type DappDlcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappDlc *DappDlcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DappDlcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappDlc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DappDlcOwnershipTransferredIterator{contract: _DappDlc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappDlc *DappDlcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DappDlcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappDlc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappDlcOwnershipTransferred)
				if err := _DappDlc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DappDlc *DappDlcFilterer) ParseOwnershipTransferred(log types.Log) (*DappDlcOwnershipTransferred, error) {
	event := new(DappDlcOwnershipTransferred)
	if err := _DappDlc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
