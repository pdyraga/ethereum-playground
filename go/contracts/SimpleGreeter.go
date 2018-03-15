// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SimpleGreeterABI is the input ABI used to generate the binding from.
const SimpleGreeterABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_greeting\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SimpleGreeter is an auto generated Go binding around an Ethereum contract.
type SimpleGreeter struct {
	SimpleGreeterCaller     // Read-only binding to the contract
	SimpleGreeterTransactor // Write-only binding to the contract
	SimpleGreeterFilterer   // Log filterer for contract events
}

// SimpleGreeterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleGreeterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleGreeterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleGreeterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleGreeterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleGreeterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleGreeterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleGreeterSession struct {
	Contract     *SimpleGreeter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleGreeterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleGreeterCallerSession struct {
	Contract *SimpleGreeterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleGreeterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleGreeterTransactorSession struct {
	Contract     *SimpleGreeterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleGreeterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleGreeterRaw struct {
	Contract *SimpleGreeter // Generic contract binding to access the raw methods on
}

// SimpleGreeterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleGreeterCallerRaw struct {
	Contract *SimpleGreeterCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleGreeterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleGreeterTransactorRaw struct {
	Contract *SimpleGreeterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleGreeter creates a new instance of SimpleGreeter, bound to a specific deployed contract.
func NewSimpleGreeter(address common.Address, backend bind.ContractBackend) (*SimpleGreeter, error) {
	contract, err := bindSimpleGreeter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleGreeter{SimpleGreeterCaller: SimpleGreeterCaller{contract: contract}, SimpleGreeterTransactor: SimpleGreeterTransactor{contract: contract}, SimpleGreeterFilterer: SimpleGreeterFilterer{contract: contract}}, nil
}

// NewSimpleGreeterCaller creates a new read-only instance of SimpleGreeter, bound to a specific deployed contract.
func NewSimpleGreeterCaller(address common.Address, caller bind.ContractCaller) (*SimpleGreeterCaller, error) {
	contract, err := bindSimpleGreeter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleGreeterCaller{contract: contract}, nil
}

// NewSimpleGreeterTransactor creates a new write-only instance of SimpleGreeter, bound to a specific deployed contract.
func NewSimpleGreeterTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleGreeterTransactor, error) {
	contract, err := bindSimpleGreeter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleGreeterTransactor{contract: contract}, nil
}

// NewSimpleGreeterFilterer creates a new log filterer instance of SimpleGreeter, bound to a specific deployed contract.
func NewSimpleGreeterFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleGreeterFilterer, error) {
	contract, err := bindSimpleGreeter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleGreeterFilterer{contract: contract}, nil
}

// bindSimpleGreeter binds a generic wrapper to an already deployed contract.
func bindSimpleGreeter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleGreeterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleGreeter *SimpleGreeterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleGreeter.Contract.SimpleGreeterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleGreeter *SimpleGreeterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleGreeter.Contract.SimpleGreeterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleGreeter *SimpleGreeterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleGreeter.Contract.SimpleGreeterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleGreeter *SimpleGreeterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleGreeter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleGreeter *SimpleGreeterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleGreeter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleGreeter *SimpleGreeterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleGreeter.Contract.contract.Transact(opts, method, params...)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() constant returns(string)
func (_SimpleGreeter *SimpleGreeterCaller) Greet(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleGreeter.contract.Call(opts, out, "greet")
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() constant returns(string)
func (_SimpleGreeter *SimpleGreeterSession) Greet() (string, error) {
	return _SimpleGreeter.Contract.Greet(&_SimpleGreeter.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() constant returns(string)
func (_SimpleGreeter *SimpleGreeterCallerSession) Greet() (string, error) {
	return _SimpleGreeter.Contract.Greet(&_SimpleGreeter.CallOpts)
}
