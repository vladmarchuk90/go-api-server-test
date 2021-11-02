// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SmartcontractMetaData contains all meta data concerning the Smartcontract contract.
var SmartcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGroupIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_groupId\",\"type\":\"uint256\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_indexId\",\"type\":\"uint256\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ethPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPriceInCents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdCapitalization\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"percentageChange\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SmartcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartcontractMetaData.ABI instead.
var SmartcontractABI = SmartcontractMetaData.ABI

// Smartcontract is an auto generated Go binding around an Ethereum contract.
type Smartcontract struct {
	SmartcontractCaller     // Read-only binding to the contract
	SmartcontractTransactor // Write-only binding to the contract
	SmartcontractFilterer   // Log filterer for contract events
}

// SmartcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartcontractSession struct {
	Contract     *Smartcontract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartcontractCallerSession struct {
	Contract *SmartcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SmartcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartcontractTransactorSession struct {
	Contract     *SmartcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SmartcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartcontractRaw struct {
	Contract *Smartcontract // Generic contract binding to access the raw methods on
}

// SmartcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartcontractCallerRaw struct {
	Contract *SmartcontractCaller // Generic read-only contract binding to access the raw methods on
}

// SmartcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartcontractTransactorRaw struct {
	Contract *SmartcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartcontract creates a new instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontract(address common.Address, backend bind.ContractBackend) (*Smartcontract, error) {
	contract, err := bindSmartcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartcontract{SmartcontractCaller: SmartcontractCaller{contract: contract}, SmartcontractTransactor: SmartcontractTransactor{contract: contract}, SmartcontractFilterer: SmartcontractFilterer{contract: contract}}, nil
}

// NewSmartcontractCaller creates a new read-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractCaller(address common.Address, caller bind.ContractCaller) (*SmartcontractCaller, error) {
	contract, err := bindSmartcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractCaller{contract: contract}, nil
}

// NewSmartcontractTransactor creates a new write-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartcontractTransactor, error) {
	contract, err := bindSmartcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractTransactor{contract: contract}, nil
}

// NewSmartcontractFilterer creates a new log filterer instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartcontractFilterer, error) {
	contract, err := bindSmartcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartcontractFilterer{contract: contract}, nil
}

// bindSmartcontract binds a generic wrapper to an already deployed contract.
func bindSmartcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.SmartcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transact(opts, method, params...)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Smartcontract *SmartcontractCaller) GetGroup(opts *bind.CallOpts, _groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "getGroup", _groupId)

	outstruct := new(struct {
		Name    string
		Indexes []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Indexes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Smartcontract *SmartcontractSession) GetGroup(_groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	return _Smartcontract.Contract.GetGroup(&_Smartcontract.CallOpts, _groupId)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Smartcontract *SmartcontractCallerSession) GetGroup(_groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	return _Smartcontract.Contract.GetGroup(&_Smartcontract.CallOpts, _groupId)
}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Smartcontract *SmartcontractCaller) GetGroupIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "getGroupIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Smartcontract *SmartcontractSession) GetGroupIds() ([]*big.Int, error) {
	return _Smartcontract.Contract.GetGroupIds(&_Smartcontract.CallOpts)
}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Smartcontract *SmartcontractCallerSession) GetGroupIds() ([]*big.Int, error) {
	return _Smartcontract.Contract.GetGroupIds(&_Smartcontract.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Smartcontract *SmartcontractCaller) GetIndex(opts *bind.CallOpts, _indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "getIndex", _indexId)

	outstruct := new(struct {
		Name              string
		EthPriceInWei     *big.Int
		UsdPriceInCents   *big.Int
		UsdCapitalization *big.Int
		PercentageChange  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EthPriceInWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UsdPriceInCents = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UsdCapitalization = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PercentageChange = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Smartcontract *SmartcontractSession) GetIndex(_indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	return _Smartcontract.Contract.GetIndex(&_Smartcontract.CallOpts, _indexId)
}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Smartcontract *SmartcontractCallerSession) GetIndex(_indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	return _Smartcontract.Contract.GetIndex(&_Smartcontract.CallOpts, _indexId)
}
