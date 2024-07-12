// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3

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

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// Uniswapv3MetaData contains all meta data concerning the Uniswapv3 contract.
var Uniswapv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Uniswapv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv3MetaData.ABI instead.
var Uniswapv3ABI = Uniswapv3MetaData.ABI

// Uniswapv3 is an auto generated Go binding around an Ethereum contract.
type Uniswapv3 struct {
	Uniswapv3Caller     // Read-only binding to the contract
	Uniswapv3Transactor // Write-only binding to the contract
	Uniswapv3Filterer   // Log filterer for contract events
}

// Uniswapv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3Session struct {
	Contract     *Uniswapv3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3CallerSession struct {
	Contract *Uniswapv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Uniswapv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3TransactorSession struct {
	Contract     *Uniswapv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Uniswapv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3Raw struct {
	Contract *Uniswapv3 // Generic contract binding to access the raw methods on
}

// Uniswapv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3CallerRaw struct {
	Contract *Uniswapv3Caller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3TransactorRaw struct {
	Contract *Uniswapv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3 creates a new instance of Uniswapv3, bound to a specific deployed contract.
func NewUniswapv3(address common.Address, backend bind.ContractBackend) (*Uniswapv3, error) {
	contract, err := bindUniswapv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3{Uniswapv3Caller: Uniswapv3Caller{contract: contract}, Uniswapv3Transactor: Uniswapv3Transactor{contract: contract}, Uniswapv3Filterer: Uniswapv3Filterer{contract: contract}}, nil
}

// NewUniswapv3Caller creates a new read-only instance of Uniswapv3, bound to a specific deployed contract.
func NewUniswapv3Caller(address common.Address, caller bind.ContractCaller) (*Uniswapv3Caller, error) {
	contract, err := bindUniswapv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3Caller{contract: contract}, nil
}

// NewUniswapv3Transactor creates a new write-only instance of Uniswapv3, bound to a specific deployed contract.
func NewUniswapv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3Transactor, error) {
	contract, err := bindUniswapv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3Transactor{contract: contract}, nil
}

// NewUniswapv3Filterer creates a new log filterer instance of Uniswapv3, bound to a specific deployed contract.
func NewUniswapv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3Filterer, error) {
	contract, err := bindUniswapv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3Filterer{contract: contract}, nil
}

// bindUniswapv3 binds a generic wrapper to an already deployed contract.
func bindUniswapv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswapv3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3 *Uniswapv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3.Contract.Uniswapv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3 *Uniswapv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3.Contract.Uniswapv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3 *Uniswapv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3.Contract.Uniswapv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3 *Uniswapv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3 *Uniswapv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3 *Uniswapv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3 *Uniswapv3Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3 *Uniswapv3Session) WETH9() (common.Address, error) {
	return _Uniswapv3.Contract.WETH9(&_Uniswapv3.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3 *Uniswapv3CallerSession) WETH9() (common.Address, error) {
	return _Uniswapv3.Contract.WETH9(&_Uniswapv3.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Uniswapv3 *Uniswapv3Caller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Uniswapv3 *Uniswapv3Session) Deployer() (common.Address, error) {
	return _Uniswapv3.Contract.Deployer(&_Uniswapv3.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Uniswapv3 *Uniswapv3CallerSession) Deployer() (common.Address, error) {
	return _Uniswapv3.Contract.Deployer(&_Uniswapv3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3 *Uniswapv3Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3 *Uniswapv3Session) Factory() (common.Address, error) {
	return _Uniswapv3.Contract.Factory(&_Uniswapv3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3 *Uniswapv3CallerSession) Factory() (common.Address, error) {
	return _Uniswapv3.Contract.Factory(&_Uniswapv3.CallOpts)
}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Uniswapv3 *Uniswapv3Caller) PancakeV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _Uniswapv3.contract.Call(opts, &out, "pancakeV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Uniswapv3 *Uniswapv3Session) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _Uniswapv3.Contract.PancakeV3SwapCallback(&_Uniswapv3.CallOpts, amount0Delta, amount1Delta, path)
}

// PancakeV3SwapCallback is a free data retrieval call binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_Uniswapv3 *Uniswapv3CallerSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _Uniswapv3.Contract.PancakeV3SwapCallback(&_Uniswapv3.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Transactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Uniswapv3.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Session) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactInput(&_Uniswapv3.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3TransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactInput(&_Uniswapv3.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Transactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Session) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactInputSingle(&_Uniswapv3.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3TransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactInputSingle(&_Uniswapv3.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Transactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Uniswapv3.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Session) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactOutput(&_Uniswapv3.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3TransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactOutput(&_Uniswapv3.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Transactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3Session) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactOutputSingle(&_Uniswapv3.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_Uniswapv3 *Uniswapv3TransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3.Contract.QuoteExactOutputSingle(&_Uniswapv3.TransactOpts, params)
}
