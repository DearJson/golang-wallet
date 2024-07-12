// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recharge

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

// RechargeMetaData contains all meta data concerning the Recharge contract.
var RechargeMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"tokenTwoTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mainAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenOneTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"rechargeCard\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastOrderId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"rechargeOne\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawalManage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenCard\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress1\",\"type\":\"address\"},{\"name\":\"_amount1\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress2\",\"type\":\"address\"},{\"name\":\"_amount2\",\"type\":\"uint256\"},{\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"rechargeTwo\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderInfos\",\"outputs\":[{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"tokenOneAddress\",\"type\":\"address\"},{\"name\":\"amountOne\",\"type\":\"uint256\"},{\"name\":\"tokenTwoAddress\",\"type\":\"address\"},{\"name\":\"amountTwo\",\"type\":\"uint256\"},{\"name\":\"functionName\",\"type\":\"string\"},{\"name\":\"remarks\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lastOrderId\",\"type\":\"uint256\"}],\"name\":\"setLastOrderId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_subcoinAddress\",\"type\":\"address\"}],\"name\":\"setSubcoinAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"saveFunction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subcoinAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mainAddress\",\"type\":\"address\"}],\"name\":\"setMainAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_subcoinAddress\",\"type\":\"address\"}],\"name\":\"getList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RechargeABI is the input ABI used to generate the binding from.
// Deprecated: Use RechargeMetaData.ABI instead.
var RechargeABI = RechargeMetaData.ABI

// Recharge is an auto generated Go binding around an Ethereum contract.
type Recharge struct {
	RechargeCaller     // Read-only binding to the contract
	RechargeTransactor // Write-only binding to the contract
	RechargeFilterer   // Log filterer for contract events
}

// RechargeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RechargeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RechargeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RechargeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RechargeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RechargeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RechargeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RechargeSession struct {
	Contract     *Recharge         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RechargeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RechargeCallerSession struct {
	Contract *RechargeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RechargeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RechargeTransactorSession struct {
	Contract     *RechargeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RechargeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RechargeRaw struct {
	Contract *Recharge // Generic contract binding to access the raw methods on
}

// RechargeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RechargeCallerRaw struct {
	Contract *RechargeCaller // Generic read-only contract binding to access the raw methods on
}

// RechargeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RechargeTransactorRaw struct {
	Contract *RechargeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecharge creates a new instance of Recharge, bound to a specific deployed contract.
func NewRecharge(address common.Address, backend bind.ContractBackend) (*Recharge, error) {
	contract, err := bindRecharge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recharge{RechargeCaller: RechargeCaller{contract: contract}, RechargeTransactor: RechargeTransactor{contract: contract}, RechargeFilterer: RechargeFilterer{contract: contract}}, nil
}

// NewRechargeCaller creates a new read-only instance of Recharge, bound to a specific deployed contract.
func NewRechargeCaller(address common.Address, caller bind.ContractCaller) (*RechargeCaller, error) {
	contract, err := bindRecharge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RechargeCaller{contract: contract}, nil
}

// NewRechargeTransactor creates a new write-only instance of Recharge, bound to a specific deployed contract.
func NewRechargeTransactor(address common.Address, transactor bind.ContractTransactor) (*RechargeTransactor, error) {
	contract, err := bindRecharge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RechargeTransactor{contract: contract}, nil
}

// NewRechargeFilterer creates a new log filterer instance of Recharge, bound to a specific deployed contract.
func NewRechargeFilterer(address common.Address, filterer bind.ContractFilterer) (*RechargeFilterer, error) {
	contract, err := bindRecharge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RechargeFilterer{contract: contract}, nil
}

// bindRecharge binds a generic wrapper to an already deployed contract.
func bindRecharge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RechargeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recharge *RechargeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recharge.Contract.RechargeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recharge *RechargeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recharge *RechargeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recharge *RechargeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recharge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recharge *RechargeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recharge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recharge *RechargeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recharge.Contract.contract.Transact(opts, method, params...)
}

// LastOrderId is a free data retrieval call binding the contract method 0x5662ecc7.
//
// Solidity: function lastOrderId() view returns(uint256)
func (_Recharge *RechargeCaller) LastOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "lastOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOrderId is a free data retrieval call binding the contract method 0x5662ecc7.
//
// Solidity: function lastOrderId() view returns(uint256)
func (_Recharge *RechargeSession) LastOrderId() (*big.Int, error) {
	return _Recharge.Contract.LastOrderId(&_Recharge.CallOpts)
}

// LastOrderId is a free data retrieval call binding the contract method 0x5662ecc7.
//
// Solidity: function lastOrderId() view returns(uint256)
func (_Recharge *RechargeCallerSession) LastOrderId() (*big.Int, error) {
	return _Recharge.Contract.LastOrderId(&_Recharge.CallOpts)
}

// MainAddress is a free data retrieval call binding the contract method 0x0cdd4234.
//
// Solidity: function mainAddress() view returns(address)
func (_Recharge *RechargeCaller) MainAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "mainAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainAddress is a free data retrieval call binding the contract method 0x0cdd4234.
//
// Solidity: function mainAddress() view returns(address)
func (_Recharge *RechargeSession) MainAddress() (common.Address, error) {
	return _Recharge.Contract.MainAddress(&_Recharge.CallOpts)
}

// MainAddress is a free data retrieval call binding the contract method 0x0cdd4234.
//
// Solidity: function mainAddress() view returns(address)
func (_Recharge *RechargeCallerSession) MainAddress() (common.Address, error) {
	return _Recharge.Contract.MainAddress(&_Recharge.CallOpts)
}

// OrderInfos is a free data retrieval call binding the contract method 0xb4a1753a.
//
// Solidity: function orderInfos(uint256 ) view returns(address userAddress, address tokenOneAddress, uint256 amountOne, address tokenTwoAddress, uint256 amountTwo, string functionName, string remarks)
func (_Recharge *RechargeCaller) OrderInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserAddress     common.Address
	TokenOneAddress common.Address
	AmountOne       *big.Int
	TokenTwoAddress common.Address
	AmountTwo       *big.Int
	FunctionName    string
	Remarks         string
}, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "orderInfos", arg0)

	outstruct := new(struct {
		UserAddress     common.Address
		TokenOneAddress common.Address
		AmountOne       *big.Int
		TokenTwoAddress common.Address
		AmountTwo       *big.Int
		FunctionName    string
		Remarks         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenOneAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AmountOne = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenTwoAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.AmountTwo = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FunctionName = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Remarks = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// OrderInfos is a free data retrieval call binding the contract method 0xb4a1753a.
//
// Solidity: function orderInfos(uint256 ) view returns(address userAddress, address tokenOneAddress, uint256 amountOne, address tokenTwoAddress, uint256 amountTwo, string functionName, string remarks)
func (_Recharge *RechargeSession) OrderInfos(arg0 *big.Int) (struct {
	UserAddress     common.Address
	TokenOneAddress common.Address
	AmountOne       *big.Int
	TokenTwoAddress common.Address
	AmountTwo       *big.Int
	FunctionName    string
	Remarks         string
}, error) {
	return _Recharge.Contract.OrderInfos(&_Recharge.CallOpts, arg0)
}

// OrderInfos is a free data retrieval call binding the contract method 0xb4a1753a.
//
// Solidity: function orderInfos(uint256 ) view returns(address userAddress, address tokenOneAddress, uint256 amountOne, address tokenTwoAddress, uint256 amountTwo, string functionName, string remarks)
func (_Recharge *RechargeCallerSession) OrderInfos(arg0 *big.Int) (struct {
	UserAddress     common.Address
	TokenOneAddress common.Address
	AmountOne       *big.Int
	TokenTwoAddress common.Address
	AmountTwo       *big.Int
	FunctionName    string
	Remarks         string
}, error) {
	return _Recharge.Contract.OrderInfos(&_Recharge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recharge *RechargeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recharge *RechargeSession) Owner() (common.Address, error) {
	return _Recharge.Contract.Owner(&_Recharge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recharge *RechargeCallerSession) Owner() (common.Address, error) {
	return _Recharge.Contract.Owner(&_Recharge.CallOpts)
}

// SubcoinAddress is a free data retrieval call binding the contract method 0xd1fb3b81.
//
// Solidity: function subcoinAddress() view returns(address)
func (_Recharge *RechargeCaller) SubcoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "subcoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubcoinAddress is a free data retrieval call binding the contract method 0xd1fb3b81.
//
// Solidity: function subcoinAddress() view returns(address)
func (_Recharge *RechargeSession) SubcoinAddress() (common.Address, error) {
	return _Recharge.Contract.SubcoinAddress(&_Recharge.CallOpts)
}

// SubcoinAddress is a free data retrieval call binding the contract method 0xd1fb3b81.
//
// Solidity: function subcoinAddress() view returns(address)
func (_Recharge *RechargeCallerSession) SubcoinAddress() (common.Address, error) {
	return _Recharge.Contract.SubcoinAddress(&_Recharge.CallOpts)
}

// TokenCard is a free data retrieval call binding the contract method 0xa39b5ac3.
//
// Solidity: function tokenCard() view returns(address)
func (_Recharge *RechargeCaller) TokenCard(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "tokenCard")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenCard is a free data retrieval call binding the contract method 0xa39b5ac3.
//
// Solidity: function tokenCard() view returns(address)
func (_Recharge *RechargeSession) TokenCard() (common.Address, error) {
	return _Recharge.Contract.TokenCard(&_Recharge.CallOpts)
}

// TokenCard is a free data retrieval call binding the contract method 0xa39b5ac3.
//
// Solidity: function tokenCard() view returns(address)
func (_Recharge *RechargeCallerSession) TokenCard() (common.Address, error) {
	return _Recharge.Contract.TokenCard(&_Recharge.CallOpts)
}

// TokenOneTransfers is a free data retrieval call binding the contract method 0x327a274e.
//
// Solidity: function tokenOneTransfers() view returns(address)
func (_Recharge *RechargeCaller) TokenOneTransfers(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "tokenOneTransfers")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOneTransfers is a free data retrieval call binding the contract method 0x327a274e.
//
// Solidity: function tokenOneTransfers() view returns(address)
func (_Recharge *RechargeSession) TokenOneTransfers() (common.Address, error) {
	return _Recharge.Contract.TokenOneTransfers(&_Recharge.CallOpts)
}

// TokenOneTransfers is a free data retrieval call binding the contract method 0x327a274e.
//
// Solidity: function tokenOneTransfers() view returns(address)
func (_Recharge *RechargeCallerSession) TokenOneTransfers() (common.Address, error) {
	return _Recharge.Contract.TokenOneTransfers(&_Recharge.CallOpts)
}

// TokenTwoTransfers is a free data retrieval call binding the contract method 0x05817772.
//
// Solidity: function tokenTwoTransfers() view returns(address)
func (_Recharge *RechargeCaller) TokenTwoTransfers(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recharge.contract.Call(opts, &out, "tokenTwoTransfers")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenTwoTransfers is a free data retrieval call binding the contract method 0x05817772.
//
// Solidity: function tokenTwoTransfers() view returns(address)
func (_Recharge *RechargeSession) TokenTwoTransfers() (common.Address, error) {
	return _Recharge.Contract.TokenTwoTransfers(&_Recharge.CallOpts)
}

// TokenTwoTransfers is a free data retrieval call binding the contract method 0x05817772.
//
// Solidity: function tokenTwoTransfers() view returns(address)
func (_Recharge *RechargeCallerSession) TokenTwoTransfers() (common.Address, error) {
	return _Recharge.Contract.TokenTwoTransfers(&_Recharge.CallOpts)
}

// GetList is a paid mutator transaction binding the contract method 0xef00e37a.
//
// Solidity: function getList(address _subcoinAddress) returns()
func (_Recharge *RechargeTransactor) GetList(opts *bind.TransactOpts, _subcoinAddress common.Address) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "getList", _subcoinAddress)
}

// GetList is a paid mutator transaction binding the contract method 0xef00e37a.
//
// Solidity: function getList(address _subcoinAddress) returns()
func (_Recharge *RechargeSession) GetList(_subcoinAddress common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.GetList(&_Recharge.TransactOpts, _subcoinAddress)
}

// GetList is a paid mutator transaction binding the contract method 0xef00e37a.
//
// Solidity: function getList(address _subcoinAddress) returns()
func (_Recharge *RechargeTransactorSession) GetList(_subcoinAddress common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.GetList(&_Recharge.TransactOpts, _subcoinAddress)
}

// RechargeCard is a paid mutator transaction binding the contract method 0x34fe59b1.
//
// Solidity: function rechargeCard(address _tokenAddress, uint256 _tokenId, string remarks) payable returns()
func (_Recharge *RechargeTransactor) RechargeCard(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenId *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "rechargeCard", _tokenAddress, _tokenId, remarks)
}

// RechargeCard is a paid mutator transaction binding the contract method 0x34fe59b1.
//
// Solidity: function rechargeCard(address _tokenAddress, uint256 _tokenId, string remarks) payable returns()
func (_Recharge *RechargeSession) RechargeCard(_tokenAddress common.Address, _tokenId *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeCard(&_Recharge.TransactOpts, _tokenAddress, _tokenId, remarks)
}

// RechargeCard is a paid mutator transaction binding the contract method 0x34fe59b1.
//
// Solidity: function rechargeCard(address _tokenAddress, uint256 _tokenId, string remarks) payable returns()
func (_Recharge *RechargeTransactorSession) RechargeCard(_tokenAddress common.Address, _tokenId *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeCard(&_Recharge.TransactOpts, _tokenAddress, _tokenId, remarks)
}

// RechargeOne is a paid mutator transaction binding the contract method 0x68ca0399.
//
// Solidity: function rechargeOne(address _tokenAddress, uint256 amount, string remarks) payable returns()
func (_Recharge *RechargeTransactor) RechargeOne(opts *bind.TransactOpts, _tokenAddress common.Address, amount *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "rechargeOne", _tokenAddress, amount, remarks)
}

// RechargeOne is a paid mutator transaction binding the contract method 0x68ca0399.
//
// Solidity: function rechargeOne(address _tokenAddress, uint256 amount, string remarks) payable returns()
func (_Recharge *RechargeSession) RechargeOne(_tokenAddress common.Address, amount *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeOne(&_Recharge.TransactOpts, _tokenAddress, amount, remarks)
}

// RechargeOne is a paid mutator transaction binding the contract method 0x68ca0399.
//
// Solidity: function rechargeOne(address _tokenAddress, uint256 amount, string remarks) payable returns()
func (_Recharge *RechargeTransactorSession) RechargeOne(_tokenAddress common.Address, amount *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeOne(&_Recharge.TransactOpts, _tokenAddress, amount, remarks)
}

// RechargeTwo is a paid mutator transaction binding the contract method 0xa6dba280.
//
// Solidity: function rechargeTwo(address _tokenAddress1, uint256 _amount1, address _tokenAddress2, uint256 _amount2, string remarks) payable returns()
func (_Recharge *RechargeTransactor) RechargeTwo(opts *bind.TransactOpts, _tokenAddress1 common.Address, _amount1 *big.Int, _tokenAddress2 common.Address, _amount2 *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "rechargeTwo", _tokenAddress1, _amount1, _tokenAddress2, _amount2, remarks)
}

// RechargeTwo is a paid mutator transaction binding the contract method 0xa6dba280.
//
// Solidity: function rechargeTwo(address _tokenAddress1, uint256 _amount1, address _tokenAddress2, uint256 _amount2, string remarks) payable returns()
func (_Recharge *RechargeSession) RechargeTwo(_tokenAddress1 common.Address, _amount1 *big.Int, _tokenAddress2 common.Address, _amount2 *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeTwo(&_Recharge.TransactOpts, _tokenAddress1, _amount1, _tokenAddress2, _amount2, remarks)
}

// RechargeTwo is a paid mutator transaction binding the contract method 0xa6dba280.
//
// Solidity: function rechargeTwo(address _tokenAddress1, uint256 _amount1, address _tokenAddress2, uint256 _amount2, string remarks) payable returns()
func (_Recharge *RechargeTransactorSession) RechargeTwo(_tokenAddress1 common.Address, _amount1 *big.Int, _tokenAddress2 common.Address, _amount2 *big.Int, remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.RechargeTwo(&_Recharge.TransactOpts, _tokenAddress1, _amount1, _tokenAddress2, _amount2, remarks)
}

// SaveFunction is a paid mutator transaction binding the contract method 0xc6487348.
//
// Solidity: function saveFunction(string remarks) returns()
func (_Recharge *RechargeTransactor) SaveFunction(opts *bind.TransactOpts, remarks string) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "saveFunction", remarks)
}

// SaveFunction is a paid mutator transaction binding the contract method 0xc6487348.
//
// Solidity: function saveFunction(string remarks) returns()
func (_Recharge *RechargeSession) SaveFunction(remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.SaveFunction(&_Recharge.TransactOpts, remarks)
}

// SaveFunction is a paid mutator transaction binding the contract method 0xc6487348.
//
// Solidity: function saveFunction(string remarks) returns()
func (_Recharge *RechargeTransactorSession) SaveFunction(remarks string) (*types.Transaction, error) {
	return _Recharge.Contract.SaveFunction(&_Recharge.TransactOpts, remarks)
}

// SetLastOrderId is a paid mutator transaction binding the contract method 0xb83b845f.
//
// Solidity: function setLastOrderId(uint256 _lastOrderId) returns()
func (_Recharge *RechargeTransactor) SetLastOrderId(opts *bind.TransactOpts, _lastOrderId *big.Int) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "setLastOrderId", _lastOrderId)
}

// SetLastOrderId is a paid mutator transaction binding the contract method 0xb83b845f.
//
// Solidity: function setLastOrderId(uint256 _lastOrderId) returns()
func (_Recharge *RechargeSession) SetLastOrderId(_lastOrderId *big.Int) (*types.Transaction, error) {
	return _Recharge.Contract.SetLastOrderId(&_Recharge.TransactOpts, _lastOrderId)
}

// SetLastOrderId is a paid mutator transaction binding the contract method 0xb83b845f.
//
// Solidity: function setLastOrderId(uint256 _lastOrderId) returns()
func (_Recharge *RechargeTransactorSession) SetLastOrderId(_lastOrderId *big.Int) (*types.Transaction, error) {
	return _Recharge.Contract.SetLastOrderId(&_Recharge.TransactOpts, _lastOrderId)
}

// SetMainAddress is a paid mutator transaction binding the contract method 0xdb9771f5.
//
// Solidity: function setMainAddress(address _mainAddress) returns()
func (_Recharge *RechargeTransactor) SetMainAddress(opts *bind.TransactOpts, _mainAddress common.Address) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "setMainAddress", _mainAddress)
}

// SetMainAddress is a paid mutator transaction binding the contract method 0xdb9771f5.
//
// Solidity: function setMainAddress(address _mainAddress) returns()
func (_Recharge *RechargeSession) SetMainAddress(_mainAddress common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.SetMainAddress(&_Recharge.TransactOpts, _mainAddress)
}

// SetMainAddress is a paid mutator transaction binding the contract method 0xdb9771f5.
//
// Solidity: function setMainAddress(address _mainAddress) returns()
func (_Recharge *RechargeTransactorSession) SetMainAddress(_mainAddress common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.SetMainAddress(&_Recharge.TransactOpts, _mainAddress)
}

// SetSubcoinAddress is a paid mutator transaction binding the contract method 0xc12bf719.
//
// Solidity: function setSubcoinAddress(address _subcoinAddress) returns()
func (_Recharge *RechargeTransactor) SetSubcoinAddress(opts *bind.TransactOpts, _subcoinAddress common.Address) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "setSubcoinAddress", _subcoinAddress)
}

// SetSubcoinAddress is a paid mutator transaction binding the contract method 0xc12bf719.
//
// Solidity: function setSubcoinAddress(address _subcoinAddress) returns()
func (_Recharge *RechargeSession) SetSubcoinAddress(_subcoinAddress common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.SetSubcoinAddress(&_Recharge.TransactOpts, _subcoinAddress)
}

// SetSubcoinAddress is a paid mutator transaction binding the contract method 0xc12bf719.
//
// Solidity: function setSubcoinAddress(address _subcoinAddress) returns()
func (_Recharge *RechargeTransactorSession) SetSubcoinAddress(_subcoinAddress common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.SetSubcoinAddress(&_Recharge.TransactOpts, _subcoinAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Recharge *RechargeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Recharge *RechargeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.TransferOwnership(&_Recharge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Recharge *RechargeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Recharge.Contract.TransferOwnership(&_Recharge.TransactOpts, newOwner)
}

// WithdrawalManage is a paid mutator transaction binding the contract method 0x6c43c7ed.
//
// Solidity: function withdrawalManage(address userAddress, address tokenAddress, uint256 amount) returns()
func (_Recharge *RechargeTransactor) WithdrawalManage(opts *bind.TransactOpts, userAddress common.Address, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Recharge.contract.Transact(opts, "withdrawalManage", userAddress, tokenAddress, amount)
}

// WithdrawalManage is a paid mutator transaction binding the contract method 0x6c43c7ed.
//
// Solidity: function withdrawalManage(address userAddress, address tokenAddress, uint256 amount) returns()
func (_Recharge *RechargeSession) WithdrawalManage(userAddress common.Address, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Recharge.Contract.WithdrawalManage(&_Recharge.TransactOpts, userAddress, tokenAddress, amount)
}

// WithdrawalManage is a paid mutator transaction binding the contract method 0x6c43c7ed.
//
// Solidity: function withdrawalManage(address userAddress, address tokenAddress, uint256 amount) returns()
func (_Recharge *RechargeTransactorSession) WithdrawalManage(userAddress common.Address, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Recharge.Contract.WithdrawalManage(&_Recharge.TransactOpts, userAddress, tokenAddress, amount)
}
