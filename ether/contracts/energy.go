// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EnergyABI is the input ABI used to generate the binding from.
const EnergyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"finishSup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consumer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishExp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_supplier\",\"type\":\"address\"},{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// EnergyBin is the compiled bytecode used for deploying new contracts.
const EnergyBin = `0x60806040526040516060806103068339810180604052606081101561002357600080fd5b508051602082015160409092015160018054600160a060020a0319908116331790915560008054600160a060020a03909416939091169290921790915560029190915560035561028e806100786000396000f3fe608060405234801561001057600080fd5b506004361061007e577c010000000000000000000000000000000000000000000000000000000060003504630fba8a1581146100835780631b3012a3146100b75780633197cbb6146100d157806336da4468146100d9578063b4fd7296146100fd578063bba3815d14610105575b600080fd5b6100b56004803603608081101561009957600080fd5b5080359060ff602082013516906040810135906060013561010d565b005b6100bf61021b565b60408051918252519081900360200190f35b6100bf610221565b6100e1610227565b60408051600160a060020a039092168252519081900360200190f35b6100e1610236565b6100b5610245565b600054600160a060020a0316331461012457600080fd5b6001805460035460408051602080820193909352808201899052815180820383018152606082018084528151918501919091206000909152608082018084525260ff881660a082015260c0810187905260e081018690529051600160a060020a0390931693926101008281019392601f198301929081900390910190855afa1580156101b4573d6000803e3d6000fd5b50505060206040510351600160a060020a03161415156101d357600080fd5b60008054604051600160a060020a039091169186156108fc02918791818181858888f1935050505015801561020c573d6000803e3d6000fd5b50600154600160a060020a0316ff5b60035481565b60025481565b600054600160a060020a031681565b600154600160a060020a031681565b60025442101561025457600080fd5b600154600160a060020a0316fffea165627a7a72305820786995ab02e274c577d8baddda9d63a64912462adcee6ae0cd677752c6f2c7ad0029`

// DeployEnergy deploys a new Ethereum contract, binding an instance of Energy to it.
func DeployEnergy(auth *bind.TransactOpts, backend bind.ContractBackend, _supplier common.Address, _endTime *big.Int, _dataHash [32]byte) (common.Address, *types.Transaction, *Energy, error) {
	parsed, err := abi.JSON(strings.NewReader(EnergyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnergyBin), backend, _supplier, _endTime, _dataHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Energy{EnergyCaller: EnergyCaller{contract: contract}, EnergyTransactor: EnergyTransactor{contract: contract}, EnergyFilterer: EnergyFilterer{contract: contract}}, nil
}

// Energy is an auto generated Go binding around an Ethereum contract.
type Energy struct {
	EnergyCaller     // Read-only binding to the contract
	EnergyTransactor // Write-only binding to the contract
	EnergyFilterer   // Log filterer for contract events
}

// EnergyCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnergyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnergyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnergyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnergySession struct {
	Contract     *Energy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnergyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnergyCallerSession struct {
	Contract *EnergyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EnergyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnergyTransactorSession struct {
	Contract     *EnergyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnergyRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnergyRaw struct {
	Contract *Energy // Generic contract binding to access the raw methods on
}

// EnergyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnergyCallerRaw struct {
	Contract *EnergyCaller // Generic read-only contract binding to access the raw methods on
}

// EnergyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnergyTransactorRaw struct {
	Contract *EnergyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnergy creates a new instance of Energy, bound to a specific deployed contract.
func NewEnergy(address common.Address, backend bind.ContractBackend) (*Energy, error) {
	contract, err := bindEnergy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Energy{EnergyCaller: EnergyCaller{contract: contract}, EnergyTransactor: EnergyTransactor{contract: contract}, EnergyFilterer: EnergyFilterer{contract: contract}}, nil
}

// NewEnergyCaller creates a new read-only instance of Energy, bound to a specific deployed contract.
func NewEnergyCaller(address common.Address, caller bind.ContractCaller) (*EnergyCaller, error) {
	contract, err := bindEnergy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnergyCaller{contract: contract}, nil
}

// NewEnergyTransactor creates a new write-only instance of Energy, bound to a specific deployed contract.
func NewEnergyTransactor(address common.Address, transactor bind.ContractTransactor) (*EnergyTransactor, error) {
	contract, err := bindEnergy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnergyTransactor{contract: contract}, nil
}

// NewEnergyFilterer creates a new log filterer instance of Energy, bound to a specific deployed contract.
func NewEnergyFilterer(address common.Address, filterer bind.ContractFilterer) (*EnergyFilterer, error) {
	contract, err := bindEnergy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnergyFilterer{contract: contract}, nil
}

// bindEnergy binds a generic wrapper to an already deployed contract.
func bindEnergy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnergyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Energy *EnergyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Energy.Contract.EnergyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Energy *EnergyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energy.Contract.EnergyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Energy *EnergyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Energy.Contract.EnergyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Energy *EnergyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Energy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Energy *EnergyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Energy *EnergyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Energy.Contract.contract.Transact(opts, method, params...)
}

// Consumer is a free data retrieval call binding the contract method 0xb4fd7296.
//
// Solidity: function consumer() constant returns(address)
func (_Energy *EnergyCaller) Consumer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Energy.contract.Call(opts, out, "consumer")
	return *ret0, err
}

// Consumer is a free data retrieval call binding the contract method 0xb4fd7296.
//
// Solidity: function consumer() constant returns(address)
func (_Energy *EnergySession) Consumer() (common.Address, error) {
	return _Energy.Contract.Consumer(&_Energy.CallOpts)
}

// Consumer is a free data retrieval call binding the contract method 0xb4fd7296.
//
// Solidity: function consumer() constant returns(address)
func (_Energy *EnergyCallerSession) Consumer() (common.Address, error) {
	return _Energy.Contract.Consumer(&_Energy.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(bytes32)
func (_Energy *EnergyCaller) DataHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Energy.contract.Call(opts, out, "dataHash")
	return *ret0, err
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(bytes32)
func (_Energy *EnergySession) DataHash() ([32]byte, error) {
	return _Energy.Contract.DataHash(&_Energy.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(bytes32)
func (_Energy *EnergyCallerSession) DataHash() ([32]byte, error) {
	return _Energy.Contract.DataHash(&_Energy.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Energy *EnergyCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Energy.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Energy *EnergySession) EndTime() (*big.Int, error) {
	return _Energy.Contract.EndTime(&_Energy.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_Energy *EnergyCallerSession) EndTime() (*big.Int, error) {
	return _Energy.Contract.EndTime(&_Energy.CallOpts)
}

// Supplier is a free data retrieval call binding the contract method 0x36da4468.
//
// Solidity: function supplier() constant returns(address)
func (_Energy *EnergyCaller) Supplier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Energy.contract.Call(opts, out, "supplier")
	return *ret0, err
}

// Supplier is a free data retrieval call binding the contract method 0x36da4468.
//
// Solidity: function supplier() constant returns(address)
func (_Energy *EnergySession) Supplier() (common.Address, error) {
	return _Energy.Contract.Supplier(&_Energy.CallOpts)
}

// Supplier is a free data retrieval call binding the contract method 0x36da4468.
//
// Solidity: function supplier() constant returns(address)
func (_Energy *EnergyCallerSession) Supplier() (common.Address, error) {
	return _Energy.Contract.Supplier(&_Energy.CallOpts)
}

// FinishExp is a paid mutator transaction binding the contract method 0xbba3815d.
//
// Solidity: function finishExp() returns()
func (_Energy *EnergyTransactor) FinishExp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energy.contract.Transact(opts, "finishExp")
}

// FinishExp is a paid mutator transaction binding the contract method 0xbba3815d.
//
// Solidity: function finishExp() returns()
func (_Energy *EnergySession) FinishExp() (*types.Transaction, error) {
	return _Energy.Contract.FinishExp(&_Energy.TransactOpts)
}

// FinishExp is a paid mutator transaction binding the contract method 0xbba3815d.
//
// Solidity: function finishExp() returns()
func (_Energy *EnergyTransactorSession) FinishExp() (*types.Transaction, error) {
	return _Energy.Contract.FinishExp(&_Energy.TransactOpts)
}

// FinishSup is a paid mutator transaction binding the contract method 0x0fba8a15.
//
// Solidity: function finishSup(amount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Energy *EnergyTransactor) FinishSup(opts *bind.TransactOpts, amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Energy.contract.Transact(opts, "finishSup", amount, v, r, s)
}

// FinishSup is a paid mutator transaction binding the contract method 0x0fba8a15.
//
// Solidity: function finishSup(amount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Energy *EnergySession) FinishSup(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Energy.Contract.FinishSup(&_Energy.TransactOpts, amount, v, r, s)
}

// FinishSup is a paid mutator transaction binding the contract method 0x0fba8a15.
//
// Solidity: function finishSup(amount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Energy *EnergyTransactorSession) FinishSup(amount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Energy.Contract.FinishSup(&_Energy.TransactOpts, amount, v, r, s)
}
