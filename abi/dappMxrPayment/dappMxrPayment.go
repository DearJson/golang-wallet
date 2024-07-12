// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dappMxrPayment

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

// DappMxrPaymentMetaData contains all meta data concerning the DappMxrPayment contract.
var DappMxrPaymentMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToCoinAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_shoubiAddress\",\"type\":\"address\"}],\"name\":\"setShoubi\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whiteAddress\",\"type\":\"address\"}],\"name\":\"setWhite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DappMxrPaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use DappMxrPaymentMetaData.ABI instead.
var DappMxrPaymentABI = DappMxrPaymentMetaData.ABI

// DappMxrPayment is an auto generated Go binding around an Ethereum contract.
type DappMxrPayment struct {
	DappMxrPaymentCaller     // Read-only binding to the contract
	DappMxrPaymentTransactor // Write-only binding to the contract
	DappMxrPaymentFilterer   // Log filterer for contract events
}

// DappMxrPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappMxrPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappMxrPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappMxrPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappMxrPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappMxrPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappMxrPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappMxrPaymentSession struct {
	Contract     *DappMxrPayment   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappMxrPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappMxrPaymentCallerSession struct {
	Contract *DappMxrPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DappMxrPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappMxrPaymentTransactorSession struct {
	Contract     *DappMxrPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DappMxrPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappMxrPaymentRaw struct {
	Contract *DappMxrPayment // Generic contract binding to access the raw methods on
}

// DappMxrPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappMxrPaymentCallerRaw struct {
	Contract *DappMxrPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// DappMxrPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappMxrPaymentTransactorRaw struct {
	Contract *DappMxrPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappMxrPayment creates a new instance of DappMxrPayment, bound to a specific deployed contract.
func NewDappMxrPayment(address common.Address, backend bind.ContractBackend) (*DappMxrPayment, error) {
	contract, err := bindDappMxrPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappMxrPayment{DappMxrPaymentCaller: DappMxrPaymentCaller{contract: contract}, DappMxrPaymentTransactor: DappMxrPaymentTransactor{contract: contract}, DappMxrPaymentFilterer: DappMxrPaymentFilterer{contract: contract}}, nil
}

// NewDappMxrPaymentCaller creates a new read-only instance of DappMxrPayment, bound to a specific deployed contract.
func NewDappMxrPaymentCaller(address common.Address, caller bind.ContractCaller) (*DappMxrPaymentCaller, error) {
	contract, err := bindDappMxrPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappMxrPaymentCaller{contract: contract}, nil
}

// NewDappMxrPaymentTransactor creates a new write-only instance of DappMxrPayment, bound to a specific deployed contract.
func NewDappMxrPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*DappMxrPaymentTransactor, error) {
	contract, err := bindDappMxrPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappMxrPaymentTransactor{contract: contract}, nil
}

// NewDappMxrPaymentFilterer creates a new log filterer instance of DappMxrPayment, bound to a specific deployed contract.
func NewDappMxrPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*DappMxrPaymentFilterer, error) {
	contract, err := bindDappMxrPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappMxrPaymentFilterer{contract: contract}, nil
}

// bindDappMxrPayment binds a generic wrapper to an already deployed contract.
func bindDappMxrPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappMxrPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappMxrPayment *DappMxrPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappMxrPayment.Contract.DappMxrPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappMxrPayment *DappMxrPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.DappMxrPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappMxrPayment *DappMxrPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.DappMxrPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappMxrPayment *DappMxrPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappMxrPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappMxrPayment *DappMxrPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappMxrPayment *DappMxrPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappMxrPayment *DappMxrPaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappMxrPayment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappMxrPayment *DappMxrPaymentSession) Owner() (common.Address, error) {
	return _DappMxrPayment.Contract.Owner(&_DappMxrPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappMxrPayment *DappMxrPaymentCallerSession) Owner() (common.Address, error) {
	return _DappMxrPayment.Contract.Owner(&_DappMxrPayment.CallOpts)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_DappMxrPayment *DappMxrPaymentCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DappMxrPayment.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_DappMxrPayment *DappMxrPaymentSession) WhiteList(arg0 common.Address) (bool, error) {
	return _DappMxrPayment.Contract.WhiteList(&_DappMxrPayment.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_DappMxrPayment *DappMxrPaymentCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _DappMxrPayment.Contract.WhiteList(&_DappMxrPayment.CallOpts, arg0)
}

// SetShoubi is a paid mutator transaction binding the contract method 0xb52d55e1.
//
// Solidity: function setShoubi(address _shoubiAddress) returns()
func (_DappMxrPayment *DappMxrPaymentTransactor) SetShoubi(opts *bind.TransactOpts, _shoubiAddress common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.contract.Transact(opts, "setShoubi", _shoubiAddress)
}

// SetShoubi is a paid mutator transaction binding the contract method 0xb52d55e1.
//
// Solidity: function setShoubi(address _shoubiAddress) returns()
func (_DappMxrPayment *DappMxrPaymentSession) SetShoubi(_shoubiAddress common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.SetShoubi(&_DappMxrPayment.TransactOpts, _shoubiAddress)
}

// SetShoubi is a paid mutator transaction binding the contract method 0xb52d55e1.
//
// Solidity: function setShoubi(address _shoubiAddress) returns()
func (_DappMxrPayment *DappMxrPaymentTransactorSession) SetShoubi(_shoubiAddress common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.SetShoubi(&_DappMxrPayment.TransactOpts, _shoubiAddress)
}

// SetWhite is a paid mutator transaction binding the contract method 0xc03646ba.
//
// Solidity: function setWhite(address _whiteAddress) returns()
func (_DappMxrPayment *DappMxrPaymentTransactor) SetWhite(opts *bind.TransactOpts, _whiteAddress common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.contract.Transact(opts, "setWhite", _whiteAddress)
}

// SetWhite is a paid mutator transaction binding the contract method 0xc03646ba.
//
// Solidity: function setWhite(address _whiteAddress) returns()
func (_DappMxrPayment *DappMxrPaymentSession) SetWhite(_whiteAddress common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.SetWhite(&_DappMxrPayment.TransactOpts, _whiteAddress)
}

// SetWhite is a paid mutator transaction binding the contract method 0xc03646ba.
//
// Solidity: function setWhite(address _whiteAddress) returns()
func (_DappMxrPayment *DappMxrPaymentTransactorSession) SetWhite(_whiteAddress common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.SetWhite(&_DappMxrPayment.TransactOpts, _whiteAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappMxrPayment *DappMxrPaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappMxrPayment *DappMxrPaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.TransferOwnership(&_DappMxrPayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappMxrPayment *DappMxrPaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.TransferOwnership(&_DappMxrPayment.TransactOpts, newOwner)
}

// WithdrawToCoinAddress is a paid mutator transaction binding the contract method 0xa3af9845.
//
// Solidity: function withdrawToCoinAddress(address _tokenAddress, uint256 _amount) returns()
func (_DappMxrPayment *DappMxrPaymentTransactor) WithdrawToCoinAddress(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DappMxrPayment.contract.Transact(opts, "withdrawToCoinAddress", _tokenAddress, _amount)
}

// WithdrawToCoinAddress is a paid mutator transaction binding the contract method 0xa3af9845.
//
// Solidity: function withdrawToCoinAddress(address _tokenAddress, uint256 _amount) returns()
func (_DappMxrPayment *DappMxrPaymentSession) WithdrawToCoinAddress(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.WithdrawToCoinAddress(&_DappMxrPayment.TransactOpts, _tokenAddress, _amount)
}

// WithdrawToCoinAddress is a paid mutator transaction binding the contract method 0xa3af9845.
//
// Solidity: function withdrawToCoinAddress(address _tokenAddress, uint256 _amount) returns()
func (_DappMxrPayment *DappMxrPaymentTransactorSession) WithdrawToCoinAddress(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DappMxrPayment.Contract.WithdrawToCoinAddress(&_DappMxrPayment.TransactOpts, _tokenAddress, _amount)
}
