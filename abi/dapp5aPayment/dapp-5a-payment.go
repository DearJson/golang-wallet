// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dapp5aPayment

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

// Dapp5aPaymentMetaData contains all meta data concerning the Dapp5aPayment contract.
var Dapp5aPaymentMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_coinAddress\",\"type\":\"address\"}],\"name\":\"setCoinAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToCoinAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CoinAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Dapp5aPaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use Dapp5aPaymentMetaData.ABI instead.
var Dapp5aPaymentABI = Dapp5aPaymentMetaData.ABI

// Dapp5aPayment is an auto generated Go binding around an Ethereum contract.
type Dapp5aPayment struct {
	Dapp5aPaymentCaller     // Read-only binding to the contract
	Dapp5aPaymentTransactor // Write-only binding to the contract
	Dapp5aPaymentFilterer   // Log filterer for contract events
}

// Dapp5aPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type Dapp5aPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dapp5aPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Dapp5aPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dapp5aPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Dapp5aPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dapp5aPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Dapp5aPaymentSession struct {
	Contract     *Dapp5aPayment    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Dapp5aPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Dapp5aPaymentCallerSession struct {
	Contract *Dapp5aPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Dapp5aPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Dapp5aPaymentTransactorSession struct {
	Contract     *Dapp5aPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Dapp5aPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type Dapp5aPaymentRaw struct {
	Contract *Dapp5aPayment // Generic contract binding to access the raw methods on
}

// Dapp5aPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Dapp5aPaymentCallerRaw struct {
	Contract *Dapp5aPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// Dapp5aPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Dapp5aPaymentTransactorRaw struct {
	Contract *Dapp5aPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDapp5aPayment creates a new instance of Dapp5aPayment, bound to a specific deployed contract.
func NewDapp5aPayment(address common.Address, backend bind.ContractBackend) (*Dapp5aPayment, error) {
	contract, err := bindDapp5aPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dapp5aPayment{Dapp5aPaymentCaller: Dapp5aPaymentCaller{contract: contract}, Dapp5aPaymentTransactor: Dapp5aPaymentTransactor{contract: contract}, Dapp5aPaymentFilterer: Dapp5aPaymentFilterer{contract: contract}}, nil
}

// NewDapp5aPaymentCaller creates a new read-only instance of Dapp5aPayment, bound to a specific deployed contract.
func NewDapp5aPaymentCaller(address common.Address, caller bind.ContractCaller) (*Dapp5aPaymentCaller, error) {
	contract, err := bindDapp5aPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Dapp5aPaymentCaller{contract: contract}, nil
}

// NewDapp5aPaymentTransactor creates a new write-only instance of Dapp5aPayment, bound to a specific deployed contract.
func NewDapp5aPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*Dapp5aPaymentTransactor, error) {
	contract, err := bindDapp5aPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Dapp5aPaymentTransactor{contract: contract}, nil
}

// NewDapp5aPaymentFilterer creates a new log filterer instance of Dapp5aPayment, bound to a specific deployed contract.
func NewDapp5aPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*Dapp5aPaymentFilterer, error) {
	contract, err := bindDapp5aPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Dapp5aPaymentFilterer{contract: contract}, nil
}

// bindDapp5aPayment binds a generic wrapper to an already deployed contract.
func bindDapp5aPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Dapp5aPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dapp5aPayment *Dapp5aPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dapp5aPayment.Contract.Dapp5aPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dapp5aPayment *Dapp5aPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.Dapp5aPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dapp5aPayment *Dapp5aPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.Dapp5aPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dapp5aPayment *Dapp5aPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dapp5aPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dapp5aPayment *Dapp5aPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dapp5aPayment *Dapp5aPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.contract.Transact(opts, method, params...)
}

// CoinAddress is a free data retrieval call binding the contract method 0x3b23b084.
//
// Solidity: function CoinAddress() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentCaller) CoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dapp5aPayment.contract.Call(opts, &out, "CoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CoinAddress is a free data retrieval call binding the contract method 0x3b23b084.
//
// Solidity: function CoinAddress() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentSession) CoinAddress() (common.Address, error) {
	return _Dapp5aPayment.Contract.CoinAddress(&_Dapp5aPayment.CallOpts)
}

// CoinAddress is a free data retrieval call binding the contract method 0x3b23b084.
//
// Solidity: function CoinAddress() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentCallerSession) CoinAddress() (common.Address, error) {
	return _Dapp5aPayment.Contract.CoinAddress(&_Dapp5aPayment.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x66d11c1a.
//
// Solidity: function ContractAddress() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentCaller) ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dapp5aPayment.contract.Call(opts, &out, "ContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddress is a free data retrieval call binding the contract method 0x66d11c1a.
//
// Solidity: function ContractAddress() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentSession) ContractAddress() (common.Address, error) {
	return _Dapp5aPayment.Contract.ContractAddress(&_Dapp5aPayment.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x66d11c1a.
//
// Solidity: function ContractAddress() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentCallerSession) ContractAddress() (common.Address, error) {
	return _Dapp5aPayment.Contract.ContractAddress(&_Dapp5aPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dapp5aPayment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentSession) Owner() (common.Address, error) {
	return _Dapp5aPayment.Contract.Owner(&_Dapp5aPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dapp5aPayment *Dapp5aPaymentCallerSession) Owner() (common.Address, error) {
	return _Dapp5aPayment.Contract.Owner(&_Dapp5aPayment.CallOpts)
}

// SetCoinAddress is a paid mutator transaction binding the contract method 0xb0bdacc6.
//
// Solidity: function setCoinAddress(address _coinAddress) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactor) SetCoinAddress(opts *bind.TransactOpts, _coinAddress common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.contract.Transact(opts, "setCoinAddress", _coinAddress)
}

// SetCoinAddress is a paid mutator transaction binding the contract method 0xb0bdacc6.
//
// Solidity: function setCoinAddress(address _coinAddress) returns()
func (_Dapp5aPayment *Dapp5aPaymentSession) SetCoinAddress(_coinAddress common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.SetCoinAddress(&_Dapp5aPayment.TransactOpts, _coinAddress)
}

// SetCoinAddress is a paid mutator transaction binding the contract method 0xb0bdacc6.
//
// Solidity: function setCoinAddress(address _coinAddress) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactorSession) SetCoinAddress(_coinAddress common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.SetCoinAddress(&_Dapp5aPayment.TransactOpts, _coinAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0x477bddaa.
//
// Solidity: function setContractAddress(address _contractAddress) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactor) SetContractAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.contract.Transact(opts, "setContractAddress", _contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0x477bddaa.
//
// Solidity: function setContractAddress(address _contractAddress) returns()
func (_Dapp5aPayment *Dapp5aPaymentSession) SetContractAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.SetContractAddress(&_Dapp5aPayment.TransactOpts, _contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0x477bddaa.
//
// Solidity: function setContractAddress(address _contractAddress) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactorSession) SetContractAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.SetContractAddress(&_Dapp5aPayment.TransactOpts, _contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dapp5aPayment *Dapp5aPaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.TransferOwnership(&_Dapp5aPayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.TransferOwnership(&_Dapp5aPayment.TransactOpts, newOwner)
}

// WithdrawToCoinAddress is a paid mutator transaction binding the contract method 0x7a888662.
//
// Solidity: function withdrawToCoinAddress(uint256 _amount) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactor) WithdrawToCoinAddress(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Dapp5aPayment.contract.Transact(opts, "withdrawToCoinAddress", _amount)
}

// WithdrawToCoinAddress is a paid mutator transaction binding the contract method 0x7a888662.
//
// Solidity: function withdrawToCoinAddress(uint256 _amount) returns()
func (_Dapp5aPayment *Dapp5aPaymentSession) WithdrawToCoinAddress(_amount *big.Int) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.WithdrawToCoinAddress(&_Dapp5aPayment.TransactOpts, _amount)
}

// WithdrawToCoinAddress is a paid mutator transaction binding the contract method 0x7a888662.
//
// Solidity: function withdrawToCoinAddress(uint256 _amount) returns()
func (_Dapp5aPayment *Dapp5aPaymentTransactorSession) WithdrawToCoinAddress(_amount *big.Int) (*types.Transaction, error) {
	return _Dapp5aPayment.Contract.WithdrawToCoinAddress(&_Dapp5aPayment.TransactOpts, _amount)
}
