// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappaco_recharge

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

// DappacoMetaData contains all meta data concerning the Dappaco contract.
var DappacoMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_coinAddress\",\"type\":\"address\"},{\"name\":\"_userAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"contractTransfer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenOneTransfers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DappacoABI is the input ABI used to generate the binding from.
// Deprecated: Use DappacoMetaData.ABI instead.
var DappacoABI = DappacoMetaData.ABI

// Dappaco is an auto generated Go binding around an Ethereum contract.
type Dappaco struct {
	DappacoCaller     // Read-only binding to the contract
	DappacoTransactor // Write-only binding to the contract
	DappacoFilterer   // Log filterer for contract events
}

// DappacoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappacoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappacoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappacoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappacoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappacoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappacoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappacoSession struct {
	Contract     *Dappaco          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappacoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappacoCallerSession struct {
	Contract *DappacoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DappacoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappacoTransactorSession struct {
	Contract     *DappacoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DappacoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappacoRaw struct {
	Contract *Dappaco // Generic contract binding to access the raw methods on
}

// DappacoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappacoCallerRaw struct {
	Contract *DappacoCaller // Generic read-only contract binding to access the raw methods on
}

// DappacoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappacoTransactorRaw struct {
	Contract *DappacoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappaco creates a new instance of Dappaco, bound to a specific deployed contract.
func NewDappaco(address common.Address, backend bind.ContractBackend) (*Dappaco, error) {
	contract, err := bindDappaco(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dappaco{DappacoCaller: DappacoCaller{contract: contract}, DappacoTransactor: DappacoTransactor{contract: contract}, DappacoFilterer: DappacoFilterer{contract: contract}}, nil
}

// NewDappacoCaller creates a new read-only instance of Dappaco, bound to a specific deployed contract.
func NewDappacoCaller(address common.Address, caller bind.ContractCaller) (*DappacoCaller, error) {
	contract, err := bindDappaco(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappacoCaller{contract: contract}, nil
}

// NewDappacoTransactor creates a new write-only instance of Dappaco, bound to a specific deployed contract.
func NewDappacoTransactor(address common.Address, transactor bind.ContractTransactor) (*DappacoTransactor, error) {
	contract, err := bindDappaco(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappacoTransactor{contract: contract}, nil
}

// NewDappacoFilterer creates a new log filterer instance of Dappaco, bound to a specific deployed contract.
func NewDappacoFilterer(address common.Address, filterer bind.ContractFilterer) (*DappacoFilterer, error) {
	contract, err := bindDappaco(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappacoFilterer{contract: contract}, nil
}

// bindDappaco binds a generic wrapper to an already deployed contract.
func bindDappaco(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappacoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dappaco *DappacoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dappaco.Contract.DappacoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dappaco *DappacoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dappaco.Contract.DappacoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dappaco *DappacoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dappaco.Contract.DappacoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dappaco *DappacoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dappaco.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dappaco *DappacoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dappaco.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dappaco *DappacoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dappaco.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dappaco *DappacoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappaco.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dappaco *DappacoSession) Owner() (common.Address, error) {
	return _Dappaco.Contract.Owner(&_Dappaco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dappaco *DappacoCallerSession) Owner() (common.Address, error) {
	return _Dappaco.Contract.Owner(&_Dappaco.CallOpts)
}

// TokenOneTransfers is a free data retrieval call binding the contract method 0x327a274e.
//
// Solidity: function tokenOneTransfers() view returns(address)
func (_Dappaco *DappacoCaller) TokenOneTransfers(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dappaco.contract.Call(opts, &out, "tokenOneTransfers")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOneTransfers is a free data retrieval call binding the contract method 0x327a274e.
//
// Solidity: function tokenOneTransfers() view returns(address)
func (_Dappaco *DappacoSession) TokenOneTransfers() (common.Address, error) {
	return _Dappaco.Contract.TokenOneTransfers(&_Dappaco.CallOpts)
}

// TokenOneTransfers is a free data retrieval call binding the contract method 0x327a274e.
//
// Solidity: function tokenOneTransfers() view returns(address)
func (_Dappaco *DappacoCallerSession) TokenOneTransfers() (common.Address, error) {
	return _Dappaco.Contract.TokenOneTransfers(&_Dappaco.CallOpts)
}

// ContractTransfer is a paid mutator transaction binding the contract method 0xf7ad906f.
//
// Solidity: function contractTransfer(address _coinAddress, address _userAddress, uint256 amount) payable returns()
func (_Dappaco *DappacoTransactor) ContractTransfer(opts *bind.TransactOpts, _coinAddress common.Address, _userAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dappaco.contract.Transact(opts, "contractTransfer", _coinAddress, _userAddress, amount)
}

// ContractTransfer is a paid mutator transaction binding the contract method 0xf7ad906f.
//
// Solidity: function contractTransfer(address _coinAddress, address _userAddress, uint256 amount) payable returns()
func (_Dappaco *DappacoSession) ContractTransfer(_coinAddress common.Address, _userAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dappaco.Contract.ContractTransfer(&_Dappaco.TransactOpts, _coinAddress, _userAddress, amount)
}

// ContractTransfer is a paid mutator transaction binding the contract method 0xf7ad906f.
//
// Solidity: function contractTransfer(address _coinAddress, address _userAddress, uint256 amount) payable returns()
func (_Dappaco *DappacoTransactorSession) ContractTransfer(_coinAddress common.Address, _userAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dappaco.Contract.ContractTransfer(&_Dappaco.TransactOpts, _coinAddress, _userAddress, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dappaco *DappacoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dappaco.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dappaco *DappacoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dappaco.Contract.TransferOwnership(&_Dappaco.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dappaco *DappacoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dappaco.Contract.TransferOwnership(&_Dappaco.TransactOpts, newOwner)
}
