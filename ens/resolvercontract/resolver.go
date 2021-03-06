// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolvercontract

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

// ResolverContractABI is the input ABI used to generate the binding from.
const ResolverContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"}]"

// ResolverContract is an auto generated Go binding around an Ethereum contract.
type ResolverContract struct {
	ResolverContractCaller     // Read-only binding to the contract
	ResolverContractTransactor // Write-only binding to the contract
	ResolverContractFilterer   // Log filterer for contract events
}

// ResolverContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverContractSession struct {
	Contract     *ResolverContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverContractCallerSession struct {
	Contract *ResolverContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ResolverContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverContractTransactorSession struct {
	Contract     *ResolverContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ResolverContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverContractRaw struct {
	Contract *ResolverContract // Generic contract binding to access the raw methods on
}

// ResolverContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverContractCallerRaw struct {
	Contract *ResolverContractCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverContractTransactorRaw struct {
	Contract *ResolverContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverContract creates a new instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContract(address common.Address, backend bind.ContractBackend) (*ResolverContract, error) {
	contract, err := bindResolverContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverContract{ResolverContractCaller: ResolverContractCaller{contract: contract}, ResolverContractTransactor: ResolverContractTransactor{contract: contract}, ResolverContractFilterer: ResolverContractFilterer{contract: contract}}, nil
}

// NewResolverContractCaller creates a new read-only instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContractCaller(address common.Address, caller bind.ContractCaller) (*ResolverContractCaller, error) {
	contract, err := bindResolverContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverContractCaller{contract: contract}, nil
}

// NewResolverContractTransactor creates a new write-only instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverContractTransactor, error) {
	contract, err := bindResolverContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverContractTransactor{contract: contract}, nil
}

// NewResolverContractFilterer creates a new log filterer instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverContractFilterer, error) {
	contract, err := bindResolverContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverContractFilterer{contract: contract}, nil
}

// bindResolverContract binds a generic wrapper to an already deployed contract.
func bindResolverContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverContract *ResolverContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverContract.Contract.ResolverContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverContract *ResolverContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverContract.Contract.ResolverContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverContract *ResolverContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverContract.Contract.ResolverContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverContract *ResolverContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverContract *ResolverContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverContract *ResolverContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverContract.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_ResolverContract *ResolverContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _ResolverContract.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_ResolverContract *ResolverContractSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _ResolverContract.Contract.ABI(&_ResolverContract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_ResolverContract *ResolverContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _ResolverContract.Contract.ABI(&_ResolverContract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_ResolverContract *ResolverContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_ResolverContract *ResolverContractSession) Addr(node [32]byte) (common.Address, error) {
	return _ResolverContract.Contract.Addr(&_ResolverContract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_ResolverContract *ResolverContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _ResolverContract.Contract.Addr(&_ResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_ResolverContract *ResolverContractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_ResolverContract *ResolverContractSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverContract.Contract.Content(&_ResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_ResolverContract *ResolverContractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverContract.Contract.Content(&_ResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_ResolverContract *ResolverContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_ResolverContract *ResolverContractSession) Name(node [32]byte) (string, error) {
	return _ResolverContract.Contract.Name(&_ResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_ResolverContract *ResolverContractCallerSession) Name(node [32]byte) (string, error) {
	return _ResolverContract.Contract.Name(&_ResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_ResolverContract *ResolverContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _ResolverContract.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_ResolverContract *ResolverContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverContract.Contract.Pubkey(&_ResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_ResolverContract *ResolverContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverContract.Contract.Pubkey(&_ResolverContract.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_ResolverContract *ResolverContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_ResolverContract *ResolverContractSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverContract.Contract.SupportsInterface(&_ResolverContract.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_ResolverContract *ResolverContractCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverContract.Contract.SupportsInterface(&_ResolverContract.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_ResolverContract *ResolverContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_ResolverContract *ResolverContractSession) Text(node [32]byte, key string) (string, error) {
	return _ResolverContract.Contract.Text(&_ResolverContract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_ResolverContract *ResolverContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _ResolverContract.Contract.Text(&_ResolverContract.CallOpts, node, key)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_ResolverContract *ResolverContractTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_ResolverContract *ResolverContractSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetABI(&_ResolverContract.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetABI(&_ResolverContract.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_ResolverContract *ResolverContractTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_ResolverContract *ResolverContractSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetAddr(&_ResolverContract.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetAddr(&_ResolverContract.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_ResolverContract *ResolverContractTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_ResolverContract *ResolverContractSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetContent(&_ResolverContract.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetContent(&_ResolverContract.TransactOpts, node, hash)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_ResolverContract *ResolverContractTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_ResolverContract *ResolverContractSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetName(&_ResolverContract.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetName(&_ResolverContract.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_ResolverContract *ResolverContractTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_ResolverContract *ResolverContractSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetPubkey(&_ResolverContract.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetPubkey(&_ResolverContract.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_ResolverContract *ResolverContractTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_ResolverContract *ResolverContractSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetText(&_ResolverContract.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetText(&_ResolverContract.TransactOpts, node, key, value)
}

// ResolverContractABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the ResolverContract contract.
type ResolverContractABIChangedIterator struct {
	Event *ResolverContractABIChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContractABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContractABIChanged)
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
		it.Event = new(ResolverContractABIChanged)
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
func (it *ResolverContractABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContractABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContractABIChanged represents a ABIChanged event raised by the ResolverContract contract.
type ResolverContractABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: e ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_ResolverContract *ResolverContractFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*ResolverContractABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ResolverContract.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContractABIChangedIterator{contract: _ResolverContract.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: e ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_ResolverContract *ResolverContractFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *ResolverContractABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ResolverContract.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContractABIChanged)
				if err := _ResolverContract.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ResolverContractAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the ResolverContract contract.
type ResolverContractAddrChangedIterator struct {
	Event *ResolverContractAddrChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContractAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContractAddrChanged)
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
		it.Event = new(ResolverContractAddrChanged)
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
func (it *ResolverContractAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContractAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContractAddrChanged represents a AddrChanged event raised by the ResolverContract contract.
type ResolverContractAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_ResolverContract *ResolverContractFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverContractAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContractAddrChangedIterator{contract: _ResolverContract.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_ResolverContract *ResolverContractFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ResolverContractAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContractAddrChanged)
				if err := _ResolverContract.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ResolverContractContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the ResolverContract contract.
type ResolverContractContentChangedIterator struct {
	Event *ResolverContractContentChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContractContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContractContentChanged)
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
		it.Event = new(ResolverContractContentChanged)
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
func (it *ResolverContractContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContractContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContractContentChanged represents a ContentChanged event raised by the ResolverContract contract.
type ResolverContractContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_ResolverContract *ResolverContractFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverContractContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContractContentChangedIterator{contract: _ResolverContract.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_ResolverContract *ResolverContractFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *ResolverContractContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContractContentChanged)
				if err := _ResolverContract.contract.UnpackLog(event, "ContentChanged", log); err != nil {
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

// ResolverContractNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the ResolverContract contract.
type ResolverContractNameChangedIterator struct {
	Event *ResolverContractNameChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContractNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContractNameChanged)
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
		it.Event = new(ResolverContractNameChanged)
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
func (it *ResolverContractNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContractNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContractNameChanged represents a NameChanged event raised by the ResolverContract contract.
type ResolverContractNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: e NameChanged(node indexed bytes32, name string)
func (_ResolverContract *ResolverContractFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverContractNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContractNameChangedIterator{contract: _ResolverContract.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: e NameChanged(node indexed bytes32, name string)
func (_ResolverContract *ResolverContractFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ResolverContractNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContractNameChanged)
				if err := _ResolverContract.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ResolverContractPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the ResolverContract contract.
type ResolverContractPubkeyChangedIterator struct {
	Event *ResolverContractPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContractPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContractPubkeyChanged)
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
		it.Event = new(ResolverContractPubkeyChanged)
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
func (it *ResolverContractPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContractPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContractPubkeyChanged represents a PubkeyChanged event raised by the ResolverContract contract.
type ResolverContractPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: e PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_ResolverContract *ResolverContractFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverContractPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContractPubkeyChangedIterator{contract: _ResolverContract.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: e PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_ResolverContract *ResolverContractFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ResolverContractPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverContract.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContractPubkeyChanged)
				if err := _ResolverContract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ResolverContractTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the ResolverContract contract.
type ResolverContractTextChangedIterator struct {
	Event *ResolverContractTextChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContractTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContractTextChanged)
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
		it.Event = new(ResolverContractTextChanged)
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
func (it *ResolverContractTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContractTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContractTextChanged represents a TextChanged event raised by the ResolverContract contract.
type ResolverContractTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: e TextChanged(node indexed bytes32, indexedKey indexed string, key string)
func (_ResolverContract *ResolverContractFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverContractTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ResolverContract.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContractTextChangedIterator{contract: _ResolverContract.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: e TextChanged(node indexed bytes32, indexedKey indexed string, key string)
func (_ResolverContract *ResolverContractFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ResolverContractTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ResolverContract.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContractTextChanged)
				if err := _ResolverContract.contract.UnpackLog(event, "TextChanged", log); err != nil {
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
