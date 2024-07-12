// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package new_recharge

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

// NewRechargeMetaData contains all meta data concerning the NewRecharge contract.
var NewRechargeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_users\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_remark\",\"type\":\"string\"}],\"name\":\"saveOrder\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCustomer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainCoinAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftcoinAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"rechargeCard\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"rechargeOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokenAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_remarks\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_customeCoin\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_customeUser\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_customeAmount\",\"type\":\"uint256[]\"}],\"name\":\"rechargeSpecifyNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_customer\",\"type\":\"bool\"}],\"name\":\"setCustomer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainAddress\",\"type\":\"address\"}],\"name\":\"setMainAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftcoinAddress\",\"type\":\"address\"}],\"name\":\"setNftcoinAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subcoinAddress\",\"type\":\"address\"}],\"name\":\"setSubcoinAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subcoinAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"fromAddress\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"withdrawalManage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NewRechargeABI is the input ABI used to generate the binding from.
// Deprecated: Use NewRechargeMetaData.ABI instead.
var NewRechargeABI = NewRechargeMetaData.ABI

// NewRecharge is an auto generated Go binding around an Ethereum contract.
type NewRecharge struct {
	NewRechargeCaller     // Read-only binding to the contract
	NewRechargeTransactor // Write-only binding to the contract
	NewRechargeFilterer   // Log filterer for contract events
}

// NewRechargeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NewRechargeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewRechargeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NewRechargeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewRechargeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NewRechargeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewRechargeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NewRechargeSession struct {
	Contract     *NewRecharge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewRechargeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NewRechargeCallerSession struct {
	Contract *NewRechargeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NewRechargeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NewRechargeTransactorSession struct {
	Contract     *NewRechargeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NewRechargeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NewRechargeRaw struct {
	Contract *NewRecharge // Generic contract binding to access the raw methods on
}

// NewRechargeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NewRechargeCallerRaw struct {
	Contract *NewRechargeCaller // Generic read-only contract binding to access the raw methods on
}

// NewRechargeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NewRechargeTransactorRaw struct {
	Contract *NewRechargeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNewRecharge creates a new instance of NewRecharge, bound to a specific deployed contract.
func NewNewRecharge(address common.Address, backend bind.ContractBackend) (*NewRecharge, error) {
	contract, err := bindNewRecharge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NewRecharge{NewRechargeCaller: NewRechargeCaller{contract: contract}, NewRechargeTransactor: NewRechargeTransactor{contract: contract}, NewRechargeFilterer: NewRechargeFilterer{contract: contract}}, nil
}

// NewNewRechargeCaller creates a new read-only instance of NewRecharge, bound to a specific deployed contract.
func NewNewRechargeCaller(address common.Address, caller bind.ContractCaller) (*NewRechargeCaller, error) {
	contract, err := bindNewRecharge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NewRechargeCaller{contract: contract}, nil
}

// NewNewRechargeTransactor creates a new write-only instance of NewRecharge, bound to a specific deployed contract.
func NewNewRechargeTransactor(address common.Address, transactor bind.ContractTransactor) (*NewRechargeTransactor, error) {
	contract, err := bindNewRecharge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NewRechargeTransactor{contract: contract}, nil
}

// NewNewRechargeFilterer creates a new log filterer instance of NewRecharge, bound to a specific deployed contract.
func NewNewRechargeFilterer(address common.Address, filterer bind.ContractFilterer) (*NewRechargeFilterer, error) {
	contract, err := bindNewRecharge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NewRechargeFilterer{contract: contract}, nil
}

// bindNewRecharge binds a generic wrapper to an already deployed contract.
func bindNewRecharge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NewRechargeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewRecharge *NewRechargeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NewRecharge.Contract.NewRechargeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewRecharge *NewRechargeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewRecharge.Contract.NewRechargeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewRecharge *NewRechargeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewRecharge.Contract.NewRechargeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewRecharge *NewRechargeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NewRecharge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewRecharge *NewRechargeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewRecharge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewRecharge *NewRechargeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewRecharge.Contract.contract.Transact(opts, method, params...)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_NewRecharge *NewRechargeCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_NewRecharge *NewRechargeSession) USDT() (common.Address, error) {
	return _NewRecharge.Contract.USDT(&_NewRecharge.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_NewRecharge *NewRechargeCallerSession) USDT() (common.Address, error) {
	return _NewRecharge.Contract.USDT(&_NewRecharge.CallOpts)
}

// ExchangeAddress is a free data retrieval call binding the contract method 0x9cd01605.
//
// Solidity: function exchangeAddress() view returns(address)
func (_NewRecharge *NewRechargeCaller) ExchangeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "exchangeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeAddress is a free data retrieval call binding the contract method 0x9cd01605.
//
// Solidity: function exchangeAddress() view returns(address)
func (_NewRecharge *NewRechargeSession) ExchangeAddress() (common.Address, error) {
	return _NewRecharge.Contract.ExchangeAddress(&_NewRecharge.CallOpts)
}

// ExchangeAddress is a free data retrieval call binding the contract method 0x9cd01605.
//
// Solidity: function exchangeAddress() view returns(address)
func (_NewRecharge *NewRechargeCallerSession) ExchangeAddress() (common.Address, error) {
	return _NewRecharge.Contract.ExchangeAddress(&_NewRecharge.CallOpts)
}

// IsCustomer is a free data retrieval call binding the contract method 0x71c3418e.
//
// Solidity: function isCustomer() view returns(bool)
func (_NewRecharge *NewRechargeCaller) IsCustomer(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "isCustomer")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCustomer is a free data retrieval call binding the contract method 0x71c3418e.
//
// Solidity: function isCustomer() view returns(bool)
func (_NewRecharge *NewRechargeSession) IsCustomer() (bool, error) {
	return _NewRecharge.Contract.IsCustomer(&_NewRecharge.CallOpts)
}

// IsCustomer is a free data retrieval call binding the contract method 0x71c3418e.
//
// Solidity: function isCustomer() view returns(bool)
func (_NewRecharge *NewRechargeCallerSession) IsCustomer() (bool, error) {
	return _NewRecharge.Contract.IsCustomer(&_NewRecharge.CallOpts)
}

// MainCoinAddress is a free data retrieval call binding the contract method 0x2d6469d7.
//
// Solidity: function mainCoinAddress() view returns(address)
func (_NewRecharge *NewRechargeCaller) MainCoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "mainCoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainCoinAddress is a free data retrieval call binding the contract method 0x2d6469d7.
//
// Solidity: function mainCoinAddress() view returns(address)
func (_NewRecharge *NewRechargeSession) MainCoinAddress() (common.Address, error) {
	return _NewRecharge.Contract.MainCoinAddress(&_NewRecharge.CallOpts)
}

// MainCoinAddress is a free data retrieval call binding the contract method 0x2d6469d7.
//
// Solidity: function mainCoinAddress() view returns(address)
func (_NewRecharge *NewRechargeCallerSession) MainCoinAddress() (common.Address, error) {
	return _NewRecharge.Contract.MainCoinAddress(&_NewRecharge.CallOpts)
}

// NftcoinAddress is a free data retrieval call binding the contract method 0xbea3cb5c.
//
// Solidity: function nftcoinAddress() view returns(address)
func (_NewRecharge *NewRechargeCaller) NftcoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "nftcoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftcoinAddress is a free data retrieval call binding the contract method 0xbea3cb5c.
//
// Solidity: function nftcoinAddress() view returns(address)
func (_NewRecharge *NewRechargeSession) NftcoinAddress() (common.Address, error) {
	return _NewRecharge.Contract.NftcoinAddress(&_NewRecharge.CallOpts)
}

// NftcoinAddress is a free data retrieval call binding the contract method 0xbea3cb5c.
//
// Solidity: function nftcoinAddress() view returns(address)
func (_NewRecharge *NewRechargeCallerSession) NftcoinAddress() (common.Address, error) {
	return _NewRecharge.Contract.NftcoinAddress(&_NewRecharge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NewRecharge *NewRechargeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NewRecharge *NewRechargeSession) Owner() (common.Address, error) {
	return _NewRecharge.Contract.Owner(&_NewRecharge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NewRecharge *NewRechargeCallerSession) Owner() (common.Address, error) {
	return _NewRecharge.Contract.Owner(&_NewRecharge.CallOpts)
}

// SubcoinAddress is a free data retrieval call binding the contract method 0xd1fb3b81.
//
// Solidity: function subcoinAddress() view returns(address)
func (_NewRecharge *NewRechargeCaller) SubcoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "subcoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubcoinAddress is a free data retrieval call binding the contract method 0xd1fb3b81.
//
// Solidity: function subcoinAddress() view returns(address)
func (_NewRecharge *NewRechargeSession) SubcoinAddress() (common.Address, error) {
	return _NewRecharge.Contract.SubcoinAddress(&_NewRecharge.CallOpts)
}

// SubcoinAddress is a free data retrieval call binding the contract method 0xd1fb3b81.
//
// Solidity: function subcoinAddress() view returns(address)
func (_NewRecharge *NewRechargeCallerSession) SubcoinAddress() (common.Address, error) {
	return _NewRecharge.Contract.SubcoinAddress(&_NewRecharge.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(uint256)
func (_NewRecharge *NewRechargeCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NewRecharge.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(uint256)
func (_NewRecharge *NewRechargeSession) Whitelist(arg0 common.Address) (*big.Int, error) {
	return _NewRecharge.Contract.Whitelist(&_NewRecharge.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(uint256)
func (_NewRecharge *NewRechargeCallerSession) Whitelist(arg0 common.Address) (*big.Int, error) {
	return _NewRecharge.Contract.Whitelist(&_NewRecharge.CallOpts, arg0)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_NewRecharge *NewRechargeTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_NewRecharge *NewRechargeSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.ChangeOwner(&_NewRecharge.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_NewRecharge *NewRechargeTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.ChangeOwner(&_NewRecharge.TransactOpts, newOwner)
}

// RechargeCard is a paid mutator transaction binding the contract method 0xf5125fa4.
//
// Solidity: function rechargeCard(address _tokenAddress, uint256 _tokenId, uint256 _num, string remarks) payable returns()
func (_NewRecharge *NewRechargeTransactor) RechargeCard(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenId *big.Int, _num *big.Int, remarks string) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "rechargeCard", _tokenAddress, _tokenId, _num, remarks)
}

// RechargeCard is a paid mutator transaction binding the contract method 0xf5125fa4.
//
// Solidity: function rechargeCard(address _tokenAddress, uint256 _tokenId, uint256 _num, string remarks) payable returns()
func (_NewRecharge *NewRechargeSession) RechargeCard(_tokenAddress common.Address, _tokenId *big.Int, _num *big.Int, remarks string) (*types.Transaction, error) {
	return _NewRecharge.Contract.RechargeCard(&_NewRecharge.TransactOpts, _tokenAddress, _tokenId, _num, remarks)
}

// RechargeCard is a paid mutator transaction binding the contract method 0xf5125fa4.
//
// Solidity: function rechargeCard(address _tokenAddress, uint256 _tokenId, uint256 _num, string remarks) payable returns()
func (_NewRecharge *NewRechargeTransactorSession) RechargeCard(_tokenAddress common.Address, _tokenId *big.Int, _num *big.Int, remarks string) (*types.Transaction, error) {
	return _NewRecharge.Contract.RechargeCard(&_NewRecharge.TransactOpts, _tokenAddress, _tokenId, _num, remarks)
}

// RechargeOne is a paid mutator transaction binding the contract method 0x68ca0399.
//
// Solidity: function rechargeOne(address _tokenAddress, uint256 amount, string remarks) payable returns()
func (_NewRecharge *NewRechargeTransactor) RechargeOne(opts *bind.TransactOpts, _tokenAddress common.Address, amount *big.Int, remarks string) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "rechargeOne", _tokenAddress, amount, remarks)
}

// RechargeOne is a paid mutator transaction binding the contract method 0x68ca0399.
//
// Solidity: function rechargeOne(address _tokenAddress, uint256 amount, string remarks) payable returns()
func (_NewRecharge *NewRechargeSession) RechargeOne(_tokenAddress common.Address, amount *big.Int, remarks string) (*types.Transaction, error) {
	return _NewRecharge.Contract.RechargeOne(&_NewRecharge.TransactOpts, _tokenAddress, amount, remarks)
}

// RechargeOne is a paid mutator transaction binding the contract method 0x68ca0399.
//
// Solidity: function rechargeOne(address _tokenAddress, uint256 amount, string remarks) payable returns()
func (_NewRecharge *NewRechargeTransactorSession) RechargeOne(_tokenAddress common.Address, amount *big.Int, remarks string) (*types.Transaction, error) {
	return _NewRecharge.Contract.RechargeOne(&_NewRecharge.TransactOpts, _tokenAddress, amount, remarks)
}

// RechargeSpecifyNew is a paid mutator transaction binding the contract method 0x2e886a1d.
//
// Solidity: function rechargeSpecifyNew(address[] _tokenAddress, uint256[] _amount, string _remarks, address[] _customeCoin, address[] _customeUser, uint256[] _customeAmount) payable returns()
func (_NewRecharge *NewRechargeTransactor) RechargeSpecifyNew(opts *bind.TransactOpts, _tokenAddress []common.Address, _amount []*big.Int, _remarks string, _customeCoin []common.Address, _customeUser []common.Address, _customeAmount []*big.Int) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "rechargeSpecifyNew", _tokenAddress, _amount, _remarks, _customeCoin, _customeUser, _customeAmount)
}

// RechargeSpecifyNew is a paid mutator transaction binding the contract method 0x2e886a1d.
//
// Solidity: function rechargeSpecifyNew(address[] _tokenAddress, uint256[] _amount, string _remarks, address[] _customeCoin, address[] _customeUser, uint256[] _customeAmount) payable returns()
func (_NewRecharge *NewRechargeSession) RechargeSpecifyNew(_tokenAddress []common.Address, _amount []*big.Int, _remarks string, _customeCoin []common.Address, _customeUser []common.Address, _customeAmount []*big.Int) (*types.Transaction, error) {
	return _NewRecharge.Contract.RechargeSpecifyNew(&_NewRecharge.TransactOpts, _tokenAddress, _amount, _remarks, _customeCoin, _customeUser, _customeAmount)
}

// RechargeSpecifyNew is a paid mutator transaction binding the contract method 0x2e886a1d.
//
// Solidity: function rechargeSpecifyNew(address[] _tokenAddress, uint256[] _amount, string _remarks, address[] _customeCoin, address[] _customeUser, uint256[] _customeAmount) payable returns()
func (_NewRecharge *NewRechargeTransactorSession) RechargeSpecifyNew(_tokenAddress []common.Address, _amount []*big.Int, _remarks string, _customeCoin []common.Address, _customeUser []common.Address, _customeAmount []*big.Int) (*types.Transaction, error) {
	return _NewRecharge.Contract.RechargeSpecifyNew(&_NewRecharge.TransactOpts, _tokenAddress, _amount, _remarks, _customeCoin, _customeUser, _customeAmount)
}

// SetCustomer is a paid mutator transaction binding the contract method 0x9bb93ff3.
//
// Solidity: function setCustomer(bool _customer) returns()
func (_NewRecharge *NewRechargeTransactor) SetCustomer(opts *bind.TransactOpts, _customer bool) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "setCustomer", _customer)
}

// SetCustomer is a paid mutator transaction binding the contract method 0x9bb93ff3.
//
// Solidity: function setCustomer(bool _customer) returns()
func (_NewRecharge *NewRechargeSession) SetCustomer(_customer bool) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetCustomer(&_NewRecharge.TransactOpts, _customer)
}

// SetCustomer is a paid mutator transaction binding the contract method 0x9bb93ff3.
//
// Solidity: function setCustomer(bool _customer) returns()
func (_NewRecharge *NewRechargeTransactorSession) SetCustomer(_customer bool) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetCustomer(&_NewRecharge.TransactOpts, _customer)
}

// SetMainAddress is a paid mutator transaction binding the contract method 0xdb9771f5.
//
// Solidity: function setMainAddress(address _mainAddress) returns()
func (_NewRecharge *NewRechargeTransactor) SetMainAddress(opts *bind.TransactOpts, _mainAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "setMainAddress", _mainAddress)
}

// SetMainAddress is a paid mutator transaction binding the contract method 0xdb9771f5.
//
// Solidity: function setMainAddress(address _mainAddress) returns()
func (_NewRecharge *NewRechargeSession) SetMainAddress(_mainAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetMainAddress(&_NewRecharge.TransactOpts, _mainAddress)
}

// SetMainAddress is a paid mutator transaction binding the contract method 0xdb9771f5.
//
// Solidity: function setMainAddress(address _mainAddress) returns()
func (_NewRecharge *NewRechargeTransactorSession) SetMainAddress(_mainAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetMainAddress(&_NewRecharge.TransactOpts, _mainAddress)
}

// SetNftcoinAddress is a paid mutator transaction binding the contract method 0x73d5a913.
//
// Solidity: function setNftcoinAddress(address _nftcoinAddress) returns()
func (_NewRecharge *NewRechargeTransactor) SetNftcoinAddress(opts *bind.TransactOpts, _nftcoinAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "setNftcoinAddress", _nftcoinAddress)
}

// SetNftcoinAddress is a paid mutator transaction binding the contract method 0x73d5a913.
//
// Solidity: function setNftcoinAddress(address _nftcoinAddress) returns()
func (_NewRecharge *NewRechargeSession) SetNftcoinAddress(_nftcoinAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetNftcoinAddress(&_NewRecharge.TransactOpts, _nftcoinAddress)
}

// SetNftcoinAddress is a paid mutator transaction binding the contract method 0x73d5a913.
//
// Solidity: function setNftcoinAddress(address _nftcoinAddress) returns()
func (_NewRecharge *NewRechargeTransactorSession) SetNftcoinAddress(_nftcoinAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetNftcoinAddress(&_NewRecharge.TransactOpts, _nftcoinAddress)
}

// SetSubcoinAddress is a paid mutator transaction binding the contract method 0xc12bf719.
//
// Solidity: function setSubcoinAddress(address _subcoinAddress) returns()
func (_NewRecharge *NewRechargeTransactor) SetSubcoinAddress(opts *bind.TransactOpts, _subcoinAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "setSubcoinAddress", _subcoinAddress)
}

// SetSubcoinAddress is a paid mutator transaction binding the contract method 0xc12bf719.
//
// Solidity: function setSubcoinAddress(address _subcoinAddress) returns()
func (_NewRecharge *NewRechargeSession) SetSubcoinAddress(_subcoinAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetSubcoinAddress(&_NewRecharge.TransactOpts, _subcoinAddress)
}

// SetSubcoinAddress is a paid mutator transaction binding the contract method 0xc12bf719.
//
// Solidity: function setSubcoinAddress(address _subcoinAddress) returns()
func (_NewRecharge *NewRechargeTransactorSession) SetSubcoinAddress(_subcoinAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetSubcoinAddress(&_NewRecharge.TransactOpts, _subcoinAddress)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _userAddress) returns()
func (_NewRecharge *NewRechargeTransactor) SetWhitelist(opts *bind.TransactOpts, _userAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "setWhitelist", _userAddress)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _userAddress) returns()
func (_NewRecharge *NewRechargeSession) SetWhitelist(_userAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetWhitelist(&_NewRecharge.TransactOpts, _userAddress)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _userAddress) returns()
func (_NewRecharge *NewRechargeTransactorSession) SetWhitelist(_userAddress common.Address) (*types.Transaction, error) {
	return _NewRecharge.Contract.SetWhitelist(&_NewRecharge.TransactOpts, _userAddress)
}

// WithdrawalManage is a paid mutator transaction binding the contract method 0xd3c6e6ff.
//
// Solidity: function withdrawalManage(address token, address[] fromAddress, address toAddress, uint256[] amount) payable returns()
func (_NewRecharge *NewRechargeTransactor) WithdrawalManage(opts *bind.TransactOpts, token common.Address, fromAddress []common.Address, toAddress common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _NewRecharge.contract.Transact(opts, "withdrawalManage", token, fromAddress, toAddress, amount)
}

// WithdrawalManage is a paid mutator transaction binding the contract method 0xd3c6e6ff.
//
// Solidity: function withdrawalManage(address token, address[] fromAddress, address toAddress, uint256[] amount) payable returns()
func (_NewRecharge *NewRechargeSession) WithdrawalManage(token common.Address, fromAddress []common.Address, toAddress common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _NewRecharge.Contract.WithdrawalManage(&_NewRecharge.TransactOpts, token, fromAddress, toAddress, amount)
}

// WithdrawalManage is a paid mutator transaction binding the contract method 0xd3c6e6ff.
//
// Solidity: function withdrawalManage(address token, address[] fromAddress, address toAddress, uint256[] amount) payable returns()
func (_NewRecharge *NewRechargeTransactorSession) WithdrawalManage(token common.Address, fromAddress []common.Address, toAddress common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _NewRecharge.Contract.WithdrawalManage(&_NewRecharge.TransactOpts, token, fromAddress, toAddress, amount)
}

// NewRechargeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NewRecharge contract.
type NewRechargeOwnershipTransferredIterator struct {
	Event *NewRechargeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NewRechargeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewRechargeOwnershipTransferred)
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
		it.Event = new(NewRechargeOwnershipTransferred)
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
func (it *NewRechargeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewRechargeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewRechargeOwnershipTransferred represents a OwnershipTransferred event raised by the NewRecharge contract.
type NewRechargeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NewRecharge *NewRechargeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NewRechargeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NewRecharge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NewRechargeOwnershipTransferredIterator{contract: _NewRecharge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NewRecharge *NewRechargeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NewRechargeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NewRecharge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewRechargeOwnershipTransferred)
				if err := _NewRecharge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NewRecharge *NewRechargeFilterer) ParseOwnershipTransferred(log types.Log) (*NewRechargeOwnershipTransferred, error) {
	event := new(NewRechargeOwnershipTransferred)
	if err := _NewRecharge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewRechargeSaveOrderIterator is returned from FilterSaveOrder and is used to iterate over the raw logs and unpacked data for SaveOrder events raised by the NewRecharge contract.
type NewRechargeSaveOrderIterator struct {
	Event *NewRechargeSaveOrder // Event containing the contract specifics and raw log

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
func (it *NewRechargeSaveOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewRechargeSaveOrder)
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
		it.Event = new(NewRechargeSaveOrder)
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
func (it *NewRechargeSaveOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewRechargeSaveOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewRechargeSaveOrder represents a SaveOrder event raised by the NewRecharge contract.
type NewRechargeSaveOrder struct {
	Users        common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Remark       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSaveOrder is a free log retrieval operation binding the contract event 0x48e57f9ad43d2dc8d5a986194fd61f3b269afbf2c510b2e9364cd42a38218e19.
//
// Solidity: event saveOrder(address _users, address _tokenAddress, uint256 _amount, string _remark)
func (_NewRecharge *NewRechargeFilterer) FilterSaveOrder(opts *bind.FilterOpts) (*NewRechargeSaveOrderIterator, error) {

	logs, sub, err := _NewRecharge.contract.FilterLogs(opts, "saveOrder")
	if err != nil {
		return nil, err
	}
	return &NewRechargeSaveOrderIterator{contract: _NewRecharge.contract, event: "saveOrder", logs: logs, sub: sub}, nil
}

// WatchSaveOrder is a free log subscription operation binding the contract event 0x48e57f9ad43d2dc8d5a986194fd61f3b269afbf2c510b2e9364cd42a38218e19.
//
// Solidity: event saveOrder(address _users, address _tokenAddress, uint256 _amount, string _remark)
func (_NewRecharge *NewRechargeFilterer) WatchSaveOrder(opts *bind.WatchOpts, sink chan<- *NewRechargeSaveOrder) (event.Subscription, error) {

	logs, sub, err := _NewRecharge.contract.WatchLogs(opts, "saveOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewRechargeSaveOrder)
				if err := _NewRecharge.contract.UnpackLog(event, "saveOrder", log); err != nil {
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

// ParseSaveOrder is a log parse operation binding the contract event 0x48e57f9ad43d2dc8d5a986194fd61f3b269afbf2c510b2e9364cd42a38218e19.
//
// Solidity: event saveOrder(address _users, address _tokenAddress, uint256 _amount, string _remark)
func (_NewRecharge *NewRechargeFilterer) ParseSaveOrder(log types.Log) (*NewRechargeSaveOrder, error) {
	event := new(NewRechargeSaveOrder)
	if err := _NewRecharge.contract.UnpackLog(event, "saveOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
