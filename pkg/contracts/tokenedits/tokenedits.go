// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenedits

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
	_ = abi.ConvertType
)

// ITokenEditsTokenEdit is an auto generated low-level Go binding around an user-defined struct.
type ITokenEditsTokenEdit struct {
	Token   common.Address
	EditIds []*big.Int
	Updates [][]MetadataInput
}

// MetadataInput is an auto generated low-level Go binding around an user-defined struct.
type MetadataInput struct {
	Field string
	Value string
}

// TokenEditsMetaData contains all meta data concerning the TokenEdits contract.
var TokenEditsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_tokentroller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenMetadata\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptEdit\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"editId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"edits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"field\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEditCount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenEdits\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"editIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"updates\",\"type\":\"tuple[][]\",\"internalType\":\"structMetadataInput[][]\",\"components\":[{\"name\":\"field\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokensWithEditsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listEdits\",\"inputs\":[{\"name\":\"initialIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenEdits\",\"type\":\"tuple[]\",\"internalType\":\"structITokenEdits.TokenEdit[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"editIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"updates\",\"type\":\"tuple[][]\",\"internalType\":\"structMetadataInput[][]\",\"components\":[{\"name\":\"field\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]}]},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeEdit\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"tuple[]\",\"internalType\":\"structMetadataInput[]\",\"components\":[{\"name\":\"field\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectEdit\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"editId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenMetadata\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokentroller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateTokentroller\",\"inputs\":[{\"name\":\"newTokentroller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EditAccepted\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"editId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EditProposed\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"editId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structMetadataInput[]\",\"components\":[{\"name\":\"field\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EditRejected\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"editId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokentrollerUpdated\",\"inputs\":[{\"name\":\"newTokentroller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// TokenEditsABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenEditsMetaData.ABI instead.
var TokenEditsABI = TokenEditsMetaData.ABI

// TokenEdits is an auto generated Go binding around an Ethereum contract.
type TokenEdits struct {
	TokenEditsCaller     // Read-only binding to the contract
	TokenEditsTransactor // Write-only binding to the contract
	TokenEditsFilterer   // Log filterer for contract events
}

// TokenEditsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenEditsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenEditsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenEditsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenEditsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenEditsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenEditsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenEditsSession struct {
	Contract     *TokenEdits       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenEditsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenEditsCallerSession struct {
	Contract *TokenEditsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenEditsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenEditsTransactorSession struct {
	Contract     *TokenEditsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenEditsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenEditsRaw struct {
	Contract *TokenEdits // Generic contract binding to access the raw methods on
}

// TokenEditsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenEditsCallerRaw struct {
	Contract *TokenEditsCaller // Generic read-only contract binding to access the raw methods on
}

// TokenEditsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenEditsTransactorRaw struct {
	Contract *TokenEditsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenEdits creates a new instance of TokenEdits, bound to a specific deployed contract.
func NewTokenEdits(address common.Address, backend bind.ContractBackend) (*TokenEdits, error) {
	contract, err := bindTokenEdits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenEdits{TokenEditsCaller: TokenEditsCaller{contract: contract}, TokenEditsTransactor: TokenEditsTransactor{contract: contract}, TokenEditsFilterer: TokenEditsFilterer{contract: contract}}, nil
}

// NewTokenEditsCaller creates a new read-only instance of TokenEdits, bound to a specific deployed contract.
func NewTokenEditsCaller(address common.Address, caller bind.ContractCaller) (*TokenEditsCaller, error) {
	contract, err := bindTokenEdits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenEditsCaller{contract: contract}, nil
}

// NewTokenEditsTransactor creates a new write-only instance of TokenEdits, bound to a specific deployed contract.
func NewTokenEditsTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenEditsTransactor, error) {
	contract, err := bindTokenEdits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenEditsTransactor{contract: contract}, nil
}

// NewTokenEditsFilterer creates a new log filterer instance of TokenEdits, bound to a specific deployed contract.
func NewTokenEditsFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenEditsFilterer, error) {
	contract, err := bindTokenEdits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenEditsFilterer{contract: contract}, nil
}

// bindTokenEdits binds a generic wrapper to an already deployed contract.
func bindTokenEdits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenEditsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenEdits *TokenEditsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenEdits.Contract.TokenEditsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenEdits *TokenEditsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenEdits.Contract.TokenEditsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenEdits *TokenEditsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenEdits.Contract.TokenEditsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenEdits *TokenEditsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenEdits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenEdits *TokenEditsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenEdits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenEdits *TokenEditsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenEdits.Contract.contract.Transact(opts, method, params...)
}

// Edits is a free data retrieval call binding the contract method 0x18fb1ce6.
//
// Solidity: function edits(address , uint256 , uint256 ) view returns(string field, string value)
func (_TokenEdits *TokenEditsCaller) Edits(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Field string
	Value string
}, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "edits", arg0, arg1, arg2)

	outstruct := new(struct {
		Field string
		Value string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Field = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Edits is a free data retrieval call binding the contract method 0x18fb1ce6.
//
// Solidity: function edits(address , uint256 , uint256 ) view returns(string field, string value)
func (_TokenEdits *TokenEditsSession) Edits(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Field string
	Value string
}, error) {
	return _TokenEdits.Contract.Edits(&_TokenEdits.CallOpts, arg0, arg1, arg2)
}

// Edits is a free data retrieval call binding the contract method 0x18fb1ce6.
//
// Solidity: function edits(address , uint256 , uint256 ) view returns(string field, string value)
func (_TokenEdits *TokenEditsCallerSession) Edits(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Field string
	Value string
}, error) {
	return _TokenEdits.Contract.Edits(&_TokenEdits.CallOpts, arg0, arg1, arg2)
}

// GetEditCount is a free data retrieval call binding the contract method 0x64906c90.
//
// Solidity: function getEditCount(address token) view returns(uint256)
func (_TokenEdits *TokenEditsCaller) GetEditCount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "getEditCount", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEditCount is a free data retrieval call binding the contract method 0x64906c90.
//
// Solidity: function getEditCount(address token) view returns(uint256)
func (_TokenEdits *TokenEditsSession) GetEditCount(token common.Address) (*big.Int, error) {
	return _TokenEdits.Contract.GetEditCount(&_TokenEdits.CallOpts, token)
}

// GetEditCount is a free data retrieval call binding the contract method 0x64906c90.
//
// Solidity: function getEditCount(address token) view returns(uint256)
func (_TokenEdits *TokenEditsCallerSession) GetEditCount(token common.Address) (*big.Int, error) {
	return _TokenEdits.Contract.GetEditCount(&_TokenEdits.CallOpts, token)
}

// GetTokenEdits is a free data retrieval call binding the contract method 0xddbc7979.
//
// Solidity: function getTokenEdits(address token) view returns(uint256[] editIds, (string,string)[][] updates)
func (_TokenEdits *TokenEditsCaller) GetTokenEdits(opts *bind.CallOpts, token common.Address) (struct {
	EditIds []*big.Int
	Updates [][]MetadataInput
}, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "getTokenEdits", token)

	outstruct := new(struct {
		EditIds []*big.Int
		Updates [][]MetadataInput
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EditIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Updates = *abi.ConvertType(out[1], new([][]MetadataInput)).(*[][]MetadataInput)

	return *outstruct, err

}

// GetTokenEdits is a free data retrieval call binding the contract method 0xddbc7979.
//
// Solidity: function getTokenEdits(address token) view returns(uint256[] editIds, (string,string)[][] updates)
func (_TokenEdits *TokenEditsSession) GetTokenEdits(token common.Address) (struct {
	EditIds []*big.Int
	Updates [][]MetadataInput
}, error) {
	return _TokenEdits.Contract.GetTokenEdits(&_TokenEdits.CallOpts, token)
}

// GetTokenEdits is a free data retrieval call binding the contract method 0xddbc7979.
//
// Solidity: function getTokenEdits(address token) view returns(uint256[] editIds, (string,string)[][] updates)
func (_TokenEdits *TokenEditsCallerSession) GetTokenEdits(token common.Address) (struct {
	EditIds []*big.Int
	Updates [][]MetadataInput
}, error) {
	return _TokenEdits.Contract.GetTokenEdits(&_TokenEdits.CallOpts, token)
}

// GetTokensWithEditsCount is a free data retrieval call binding the contract method 0x2bc30466.
//
// Solidity: function getTokensWithEditsCount() view returns(uint256)
func (_TokenEdits *TokenEditsCaller) GetTokensWithEditsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "getTokensWithEditsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokensWithEditsCount is a free data retrieval call binding the contract method 0x2bc30466.
//
// Solidity: function getTokensWithEditsCount() view returns(uint256)
func (_TokenEdits *TokenEditsSession) GetTokensWithEditsCount() (*big.Int, error) {
	return _TokenEdits.Contract.GetTokensWithEditsCount(&_TokenEdits.CallOpts)
}

// GetTokensWithEditsCount is a free data retrieval call binding the contract method 0x2bc30466.
//
// Solidity: function getTokensWithEditsCount() view returns(uint256)
func (_TokenEdits *TokenEditsCallerSession) GetTokensWithEditsCount() (*big.Int, error) {
	return _TokenEdits.Contract.GetTokensWithEditsCount(&_TokenEdits.CallOpts)
}

// ListEdits is a free data retrieval call binding the contract method 0x7c2a62f4.
//
// Solidity: function listEdits(uint256 initialIndex, uint256 size) view returns((address,uint256[],(string,string)[][])[] tokenEdits, uint256 total)
func (_TokenEdits *TokenEditsCaller) ListEdits(opts *bind.CallOpts, initialIndex *big.Int, size *big.Int) (struct {
	TokenEdits []ITokenEditsTokenEdit
	Total      *big.Int
}, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "listEdits", initialIndex, size)

	outstruct := new(struct {
		TokenEdits []ITokenEditsTokenEdit
		Total      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenEdits = *abi.ConvertType(out[0], new([]ITokenEditsTokenEdit)).(*[]ITokenEditsTokenEdit)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ListEdits is a free data retrieval call binding the contract method 0x7c2a62f4.
//
// Solidity: function listEdits(uint256 initialIndex, uint256 size) view returns((address,uint256[],(string,string)[][])[] tokenEdits, uint256 total)
func (_TokenEdits *TokenEditsSession) ListEdits(initialIndex *big.Int, size *big.Int) (struct {
	TokenEdits []ITokenEditsTokenEdit
	Total      *big.Int
}, error) {
	return _TokenEdits.Contract.ListEdits(&_TokenEdits.CallOpts, initialIndex, size)
}

// ListEdits is a free data retrieval call binding the contract method 0x7c2a62f4.
//
// Solidity: function listEdits(uint256 initialIndex, uint256 size) view returns((address,uint256[],(string,string)[][])[] tokenEdits, uint256 total)
func (_TokenEdits *TokenEditsCallerSession) ListEdits(initialIndex *big.Int, size *big.Int) (struct {
	TokenEdits []ITokenEditsTokenEdit
	Total      *big.Int
}, error) {
	return _TokenEdits.Contract.ListEdits(&_TokenEdits.CallOpts, initialIndex, size)
}

// TokenMetadata is a free data retrieval call binding the contract method 0xf5b2cad7.
//
// Solidity: function tokenMetadata() view returns(address)
func (_TokenEdits *TokenEditsCaller) TokenMetadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "tokenMetadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0xf5b2cad7.
//
// Solidity: function tokenMetadata() view returns(address)
func (_TokenEdits *TokenEditsSession) TokenMetadata() (common.Address, error) {
	return _TokenEdits.Contract.TokenMetadata(&_TokenEdits.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0xf5b2cad7.
//
// Solidity: function tokenMetadata() view returns(address)
func (_TokenEdits *TokenEditsCallerSession) TokenMetadata() (common.Address, error) {
	return _TokenEdits.Contract.TokenMetadata(&_TokenEdits.CallOpts)
}

// Tokentroller is a free data retrieval call binding the contract method 0x53403983.
//
// Solidity: function tokentroller() view returns(address)
func (_TokenEdits *TokenEditsCaller) Tokentroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenEdits.contract.Call(opts, &out, "tokentroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokentroller is a free data retrieval call binding the contract method 0x53403983.
//
// Solidity: function tokentroller() view returns(address)
func (_TokenEdits *TokenEditsSession) Tokentroller() (common.Address, error) {
	return _TokenEdits.Contract.Tokentroller(&_TokenEdits.CallOpts)
}

// Tokentroller is a free data retrieval call binding the contract method 0x53403983.
//
// Solidity: function tokentroller() view returns(address)
func (_TokenEdits *TokenEditsCallerSession) Tokentroller() (common.Address, error) {
	return _TokenEdits.Contract.Tokentroller(&_TokenEdits.CallOpts)
}

// AcceptEdit is a paid mutator transaction binding the contract method 0x7b92a17f.
//
// Solidity: function acceptEdit(address contractAddress, uint256 editId) returns()
func (_TokenEdits *TokenEditsTransactor) AcceptEdit(opts *bind.TransactOpts, contractAddress common.Address, editId *big.Int) (*types.Transaction, error) {
	return _TokenEdits.contract.Transact(opts, "acceptEdit", contractAddress, editId)
}

// AcceptEdit is a paid mutator transaction binding the contract method 0x7b92a17f.
//
// Solidity: function acceptEdit(address contractAddress, uint256 editId) returns()
func (_TokenEdits *TokenEditsSession) AcceptEdit(contractAddress common.Address, editId *big.Int) (*types.Transaction, error) {
	return _TokenEdits.Contract.AcceptEdit(&_TokenEdits.TransactOpts, contractAddress, editId)
}

// AcceptEdit is a paid mutator transaction binding the contract method 0x7b92a17f.
//
// Solidity: function acceptEdit(address contractAddress, uint256 editId) returns()
func (_TokenEdits *TokenEditsTransactorSession) AcceptEdit(contractAddress common.Address, editId *big.Int) (*types.Transaction, error) {
	return _TokenEdits.Contract.AcceptEdit(&_TokenEdits.TransactOpts, contractAddress, editId)
}

// ProposeEdit is a paid mutator transaction binding the contract method 0xbdaa83e3.
//
// Solidity: function proposeEdit(address contractAddress, (string,string)[] metadata) returns(uint256)
func (_TokenEdits *TokenEditsTransactor) ProposeEdit(opts *bind.TransactOpts, contractAddress common.Address, metadata []MetadataInput) (*types.Transaction, error) {
	return _TokenEdits.contract.Transact(opts, "proposeEdit", contractAddress, metadata)
}

// ProposeEdit is a paid mutator transaction binding the contract method 0xbdaa83e3.
//
// Solidity: function proposeEdit(address contractAddress, (string,string)[] metadata) returns(uint256)
func (_TokenEdits *TokenEditsSession) ProposeEdit(contractAddress common.Address, metadata []MetadataInput) (*types.Transaction, error) {
	return _TokenEdits.Contract.ProposeEdit(&_TokenEdits.TransactOpts, contractAddress, metadata)
}

// ProposeEdit is a paid mutator transaction binding the contract method 0xbdaa83e3.
//
// Solidity: function proposeEdit(address contractAddress, (string,string)[] metadata) returns(uint256)
func (_TokenEdits *TokenEditsTransactorSession) ProposeEdit(contractAddress common.Address, metadata []MetadataInput) (*types.Transaction, error) {
	return _TokenEdits.Contract.ProposeEdit(&_TokenEdits.TransactOpts, contractAddress, metadata)
}

// RejectEdit is a paid mutator transaction binding the contract method 0x3024d8da.
//
// Solidity: function rejectEdit(address contractAddress, uint256 editId, string reason) returns()
func (_TokenEdits *TokenEditsTransactor) RejectEdit(opts *bind.TransactOpts, contractAddress common.Address, editId *big.Int, reason string) (*types.Transaction, error) {
	return _TokenEdits.contract.Transact(opts, "rejectEdit", contractAddress, editId, reason)
}

// RejectEdit is a paid mutator transaction binding the contract method 0x3024d8da.
//
// Solidity: function rejectEdit(address contractAddress, uint256 editId, string reason) returns()
func (_TokenEdits *TokenEditsSession) RejectEdit(contractAddress common.Address, editId *big.Int, reason string) (*types.Transaction, error) {
	return _TokenEdits.Contract.RejectEdit(&_TokenEdits.TransactOpts, contractAddress, editId, reason)
}

// RejectEdit is a paid mutator transaction binding the contract method 0x3024d8da.
//
// Solidity: function rejectEdit(address contractAddress, uint256 editId, string reason) returns()
func (_TokenEdits *TokenEditsTransactorSession) RejectEdit(contractAddress common.Address, editId *big.Int, reason string) (*types.Transaction, error) {
	return _TokenEdits.Contract.RejectEdit(&_TokenEdits.TransactOpts, contractAddress, editId, reason)
}

// UpdateTokentroller is a paid mutator transaction binding the contract method 0xf177a13a.
//
// Solidity: function updateTokentroller(address newTokentroller) returns()
func (_TokenEdits *TokenEditsTransactor) UpdateTokentroller(opts *bind.TransactOpts, newTokentroller common.Address) (*types.Transaction, error) {
	return _TokenEdits.contract.Transact(opts, "updateTokentroller", newTokentroller)
}

// UpdateTokentroller is a paid mutator transaction binding the contract method 0xf177a13a.
//
// Solidity: function updateTokentroller(address newTokentroller) returns()
func (_TokenEdits *TokenEditsSession) UpdateTokentroller(newTokentroller common.Address) (*types.Transaction, error) {
	return _TokenEdits.Contract.UpdateTokentroller(&_TokenEdits.TransactOpts, newTokentroller)
}

// UpdateTokentroller is a paid mutator transaction binding the contract method 0xf177a13a.
//
// Solidity: function updateTokentroller(address newTokentroller) returns()
func (_TokenEdits *TokenEditsTransactorSession) UpdateTokentroller(newTokentroller common.Address) (*types.Transaction, error) {
	return _TokenEdits.Contract.UpdateTokentroller(&_TokenEdits.TransactOpts, newTokentroller)
}

// TokenEditsEditAcceptedIterator is returned from FilterEditAccepted and is used to iterate over the raw logs and unpacked data for EditAccepted events raised by the TokenEdits contract.
type TokenEditsEditAcceptedIterator struct {
	Event *TokenEditsEditAccepted // Event containing the contract specifics and raw log

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
func (it *TokenEditsEditAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEditsEditAccepted)
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
		it.Event = new(TokenEditsEditAccepted)
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
func (it *TokenEditsEditAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEditsEditAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEditsEditAccepted represents a EditAccepted event raised by the TokenEdits contract.
type TokenEditsEditAccepted struct {
	ContractAddress common.Address
	EditId          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEditAccepted is a free log retrieval operation binding the contract event 0x74adbe7ee9d46944947571be53e196f1fdfddd8ad202711c309d9081495d96f3.
//
// Solidity: event EditAccepted(address indexed contractAddress, uint256 editId)
func (_TokenEdits *TokenEditsFilterer) FilterEditAccepted(opts *bind.FilterOpts, contractAddress []common.Address) (*TokenEditsEditAcceptedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenEdits.contract.FilterLogs(opts, "EditAccepted", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenEditsEditAcceptedIterator{contract: _TokenEdits.contract, event: "EditAccepted", logs: logs, sub: sub}, nil
}

// WatchEditAccepted is a free log subscription operation binding the contract event 0x74adbe7ee9d46944947571be53e196f1fdfddd8ad202711c309d9081495d96f3.
//
// Solidity: event EditAccepted(address indexed contractAddress, uint256 editId)
func (_TokenEdits *TokenEditsFilterer) WatchEditAccepted(opts *bind.WatchOpts, sink chan<- *TokenEditsEditAccepted, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenEdits.contract.WatchLogs(opts, "EditAccepted", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEditsEditAccepted)
				if err := _TokenEdits.contract.UnpackLog(event, "EditAccepted", log); err != nil {
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

// ParseEditAccepted is a log parse operation binding the contract event 0x74adbe7ee9d46944947571be53e196f1fdfddd8ad202711c309d9081495d96f3.
//
// Solidity: event EditAccepted(address indexed contractAddress, uint256 editId)
func (_TokenEdits *TokenEditsFilterer) ParseEditAccepted(log types.Log) (*TokenEditsEditAccepted, error) {
	event := new(TokenEditsEditAccepted)
	if err := _TokenEdits.contract.UnpackLog(event, "EditAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEditsEditProposedIterator is returned from FilterEditProposed and is used to iterate over the raw logs and unpacked data for EditProposed events raised by the TokenEdits contract.
type TokenEditsEditProposedIterator struct {
	Event *TokenEditsEditProposed // Event containing the contract specifics and raw log

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
func (it *TokenEditsEditProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEditsEditProposed)
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
		it.Event = new(TokenEditsEditProposed)
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
func (it *TokenEditsEditProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEditsEditProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEditsEditProposed represents a EditProposed event raised by the TokenEdits contract.
type TokenEditsEditProposed struct {
	ContractAddress common.Address
	EditId          *big.Int
	Submitter       common.Address
	Metadata        []MetadataInput
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEditProposed is a free log retrieval operation binding the contract event 0x3b9f5b46962168b86fb803e1e5f0b760bb51cf8075acdd8278a9cb9a9afa40fc.
//
// Solidity: event EditProposed(address indexed contractAddress, uint256 editId, address indexed submitter, (string,string)[] metadata)
func (_TokenEdits *TokenEditsFilterer) FilterEditProposed(opts *bind.FilterOpts, contractAddress []common.Address, submitter []common.Address) (*TokenEditsEditProposedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _TokenEdits.contract.FilterLogs(opts, "EditProposed", contractAddressRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &TokenEditsEditProposedIterator{contract: _TokenEdits.contract, event: "EditProposed", logs: logs, sub: sub}, nil
}

// WatchEditProposed is a free log subscription operation binding the contract event 0x3b9f5b46962168b86fb803e1e5f0b760bb51cf8075acdd8278a9cb9a9afa40fc.
//
// Solidity: event EditProposed(address indexed contractAddress, uint256 editId, address indexed submitter, (string,string)[] metadata)
func (_TokenEdits *TokenEditsFilterer) WatchEditProposed(opts *bind.WatchOpts, sink chan<- *TokenEditsEditProposed, contractAddress []common.Address, submitter []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _TokenEdits.contract.WatchLogs(opts, "EditProposed", contractAddressRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEditsEditProposed)
				if err := _TokenEdits.contract.UnpackLog(event, "EditProposed", log); err != nil {
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

// ParseEditProposed is a log parse operation binding the contract event 0x3b9f5b46962168b86fb803e1e5f0b760bb51cf8075acdd8278a9cb9a9afa40fc.
//
// Solidity: event EditProposed(address indexed contractAddress, uint256 editId, address indexed submitter, (string,string)[] metadata)
func (_TokenEdits *TokenEditsFilterer) ParseEditProposed(log types.Log) (*TokenEditsEditProposed, error) {
	event := new(TokenEditsEditProposed)
	if err := _TokenEdits.contract.UnpackLog(event, "EditProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEditsEditRejectedIterator is returned from FilterEditRejected and is used to iterate over the raw logs and unpacked data for EditRejected events raised by the TokenEdits contract.
type TokenEditsEditRejectedIterator struct {
	Event *TokenEditsEditRejected // Event containing the contract specifics and raw log

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
func (it *TokenEditsEditRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEditsEditRejected)
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
		it.Event = new(TokenEditsEditRejected)
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
func (it *TokenEditsEditRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEditsEditRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEditsEditRejected represents a EditRejected event raised by the TokenEdits contract.
type TokenEditsEditRejected struct {
	ContractAddress common.Address
	EditId          *big.Int
	Reason          string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEditRejected is a free log retrieval operation binding the contract event 0x7c93114d227f9093d9e61216f371446b57663abb7044743032b2579310fe4078.
//
// Solidity: event EditRejected(address indexed contractAddress, uint256 editId, string reason)
func (_TokenEdits *TokenEditsFilterer) FilterEditRejected(opts *bind.FilterOpts, contractAddress []common.Address) (*TokenEditsEditRejectedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenEdits.contract.FilterLogs(opts, "EditRejected", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenEditsEditRejectedIterator{contract: _TokenEdits.contract, event: "EditRejected", logs: logs, sub: sub}, nil
}

// WatchEditRejected is a free log subscription operation binding the contract event 0x7c93114d227f9093d9e61216f371446b57663abb7044743032b2579310fe4078.
//
// Solidity: event EditRejected(address indexed contractAddress, uint256 editId, string reason)
func (_TokenEdits *TokenEditsFilterer) WatchEditRejected(opts *bind.WatchOpts, sink chan<- *TokenEditsEditRejected, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenEdits.contract.WatchLogs(opts, "EditRejected", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEditsEditRejected)
				if err := _TokenEdits.contract.UnpackLog(event, "EditRejected", log); err != nil {
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

// ParseEditRejected is a log parse operation binding the contract event 0x7c93114d227f9093d9e61216f371446b57663abb7044743032b2579310fe4078.
//
// Solidity: event EditRejected(address indexed contractAddress, uint256 editId, string reason)
func (_TokenEdits *TokenEditsFilterer) ParseEditRejected(log types.Log) (*TokenEditsEditRejected, error) {
	event := new(TokenEditsEditRejected)
	if err := _TokenEdits.contract.UnpackLog(event, "EditRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEditsTokentrollerUpdatedIterator is returned from FilterTokentrollerUpdated and is used to iterate over the raw logs and unpacked data for TokentrollerUpdated events raised by the TokenEdits contract.
type TokenEditsTokentrollerUpdatedIterator struct {
	Event *TokenEditsTokentrollerUpdated // Event containing the contract specifics and raw log

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
func (it *TokenEditsTokentrollerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEditsTokentrollerUpdated)
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
		it.Event = new(TokenEditsTokentrollerUpdated)
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
func (it *TokenEditsTokentrollerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEditsTokentrollerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEditsTokentrollerUpdated represents a TokentrollerUpdated event raised by the TokenEdits contract.
type TokenEditsTokentrollerUpdated struct {
	NewTokentroller common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokentrollerUpdated is a free log retrieval operation binding the contract event 0x24628d813f923f48992318b49019f3219233a50975fe9694bf526acffe4330ca.
//
// Solidity: event TokentrollerUpdated(address indexed newTokentroller)
func (_TokenEdits *TokenEditsFilterer) FilterTokentrollerUpdated(opts *bind.FilterOpts, newTokentroller []common.Address) (*TokenEditsTokentrollerUpdatedIterator, error) {

	var newTokentrollerRule []interface{}
	for _, newTokentrollerItem := range newTokentroller {
		newTokentrollerRule = append(newTokentrollerRule, newTokentrollerItem)
	}

	logs, sub, err := _TokenEdits.contract.FilterLogs(opts, "TokentrollerUpdated", newTokentrollerRule)
	if err != nil {
		return nil, err
	}
	return &TokenEditsTokentrollerUpdatedIterator{contract: _TokenEdits.contract, event: "TokentrollerUpdated", logs: logs, sub: sub}, nil
}

// WatchTokentrollerUpdated is a free log subscription operation binding the contract event 0x24628d813f923f48992318b49019f3219233a50975fe9694bf526acffe4330ca.
//
// Solidity: event TokentrollerUpdated(address indexed newTokentroller)
func (_TokenEdits *TokenEditsFilterer) WatchTokentrollerUpdated(opts *bind.WatchOpts, sink chan<- *TokenEditsTokentrollerUpdated, newTokentroller []common.Address) (event.Subscription, error) {

	var newTokentrollerRule []interface{}
	for _, newTokentrollerItem := range newTokentroller {
		newTokentrollerRule = append(newTokentrollerRule, newTokentrollerItem)
	}

	logs, sub, err := _TokenEdits.contract.WatchLogs(opts, "TokentrollerUpdated", newTokentrollerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEditsTokentrollerUpdated)
				if err := _TokenEdits.contract.UnpackLog(event, "TokentrollerUpdated", log); err != nil {
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

// ParseTokentrollerUpdated is a log parse operation binding the contract event 0x24628d813f923f48992318b49019f3219233a50975fe9694bf526acffe4330ca.
//
// Solidity: event TokentrollerUpdated(address indexed newTokentroller)
func (_TokenEdits *TokenEditsFilterer) ParseTokentrollerUpdated(log types.Log) (*TokenEditsTokentrollerUpdated, error) {
	event := new(TokenEditsTokentrollerUpdated)
	if err := _TokenEdits.contract.UnpackLog(event, "TokentrollerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
