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

// EchoGreeterABI is the input ABI used to generate the binding from.
const EchoGreeterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"greeting\",\"type\":\"string\"}],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// EchoGreeter is an auto generated Go binding around an Ethereum contract.
type EchoGreeter struct {
	EchoGreeterCaller     // Read-only binding to the contract
	EchoGreeterTransactor // Write-only binding to the contract
	EchoGreeterFilterer   // Log filterer for contract events
}

// EchoGreeterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EchoGreeterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EchoGreeterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EchoGreeterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EchoGreeterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EchoGreeterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EchoGreeterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EchoGreeterSession struct {
	Contract     *EchoGreeter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EchoGreeterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EchoGreeterCallerSession struct {
	Contract *EchoGreeterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EchoGreeterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EchoGreeterTransactorSession struct {
	Contract     *EchoGreeterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EchoGreeterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EchoGreeterRaw struct {
	Contract *EchoGreeter // Generic contract binding to access the raw methods on
}

// EchoGreeterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EchoGreeterCallerRaw struct {
	Contract *EchoGreeterCaller // Generic read-only contract binding to access the raw methods on
}

// EchoGreeterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EchoGreeterTransactorRaw struct {
	Contract *EchoGreeterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEchoGreeter creates a new instance of EchoGreeter, bound to a specific deployed contract.
func NewEchoGreeter(address common.Address, backend bind.ContractBackend) (*EchoGreeter, error) {
	contract, err := bindEchoGreeter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EchoGreeter{EchoGreeterCaller: EchoGreeterCaller{contract: contract}, EchoGreeterTransactor: EchoGreeterTransactor{contract: contract}, EchoGreeterFilterer: EchoGreeterFilterer{contract: contract}}, nil
}

// NewEchoGreeterCaller creates a new read-only instance of EchoGreeter, bound to a specific deployed contract.
func NewEchoGreeterCaller(address common.Address, caller bind.ContractCaller) (*EchoGreeterCaller, error) {
	contract, err := bindEchoGreeter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EchoGreeterCaller{contract: contract}, nil
}

// NewEchoGreeterTransactor creates a new write-only instance of EchoGreeter, bound to a specific deployed contract.
func NewEchoGreeterTransactor(address common.Address, transactor bind.ContractTransactor) (*EchoGreeterTransactor, error) {
	contract, err := bindEchoGreeter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EchoGreeterTransactor{contract: contract}, nil
}

// NewEchoGreeterFilterer creates a new log filterer instance of EchoGreeter, bound to a specific deployed contract.
func NewEchoGreeterFilterer(address common.Address, filterer bind.ContractFilterer) (*EchoGreeterFilterer, error) {
	contract, err := bindEchoGreeter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EchoGreeterFilterer{contract: contract}, nil
}

// bindEchoGreeter binds a generic wrapper to an already deployed contract.
func bindEchoGreeter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EchoGreeterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EchoGreeter *EchoGreeterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EchoGreeter.Contract.EchoGreeterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EchoGreeter *EchoGreeterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EchoGreeter.Contract.EchoGreeterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EchoGreeter *EchoGreeterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EchoGreeter.Contract.EchoGreeterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EchoGreeter *EchoGreeterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EchoGreeter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EchoGreeter *EchoGreeterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EchoGreeter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EchoGreeter *EchoGreeterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EchoGreeter.Contract.contract.Transact(opts, method, params...)
}

// Greet is a free data retrieval call binding the contract method 0xead710c4.
//
// Solidity: function greet(greeting string) constant returns(string)
func (_EchoGreeter *EchoGreeterCaller) Greet(opts *bind.CallOpts, greeting string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EchoGreeter.contract.Call(opts, out, "greet", greeting)
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xead710c4.
//
// Solidity: function greet(greeting string) constant returns(string)
func (_EchoGreeter *EchoGreeterSession) Greet(greeting string) (string, error) {
	return _EchoGreeter.Contract.Greet(&_EchoGreeter.CallOpts, greeting)
}

// Greet is a free data retrieval call binding the contract method 0xead710c4.
//
// Solidity: function greet(greeting string) constant returns(string)
func (_EchoGreeter *EchoGreeterCallerSession) Greet(greeting string) (string, error) {
	return _EchoGreeter.Contract.Greet(&_EchoGreeter.CallOpts, greeting)
}
