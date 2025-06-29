// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"FeeReceiverChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TreasuryDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TreasuryWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isInstant\",\"type\":\"bool\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bossToken\",\"outputs\":[{\"internalType\":\"contractBossToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"calculateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimBossToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelayed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToTreasury\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bossToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delayDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStakeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxStakePerUser\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakePerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delayDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStakeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxStakePerUser\",\"type\":\"uint256\"}],\"name\":\"setStakingParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Abi *AbiCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Abi *AbiSession) MAXFEE() (*big.Int, error) {
	return _Abi.Contract.MAXFEE(&_Abi.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Abi *AbiCallerSession) MAXFEE() (*big.Int, error) {
	return _Abi.Contract.MAXFEE(&_Abi.CallOpts)
}

// BossToken is a free data retrieval call binding the contract method 0xa60e40e5.
//
// Solidity: function bossToken() view returns(address)
func (_Abi *AbiCaller) BossToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "bossToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BossToken is a free data retrieval call binding the contract method 0xa60e40e5.
//
// Solidity: function bossToken() view returns(address)
func (_Abi *AbiSession) BossToken() (common.Address, error) {
	return _Abi.Contract.BossToken(&_Abi.CallOpts)
}

// BossToken is a free data retrieval call binding the contract method 0xa60e40e5.
//
// Solidity: function bossToken() view returns(address)
func (_Abi *AbiCallerSession) BossToken() (common.Address, error) {
	return _Abi.Contract.BossToken(&_Abi.CallOpts)
}

// CalculateReward is a free data retrieval call binding the contract method 0xd82e3962.
//
// Solidity: function calculateReward(address user) view returns(uint256)
func (_Abi *AbiCaller) CalculateReward(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "calculateReward", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0xd82e3962.
//
// Solidity: function calculateReward(address user) view returns(uint256)
func (_Abi *AbiSession) CalculateReward(user common.Address) (*big.Int, error) {
	return _Abi.Contract.CalculateReward(&_Abi.CallOpts, user)
}

// CalculateReward is a free data retrieval call binding the contract method 0xd82e3962.
//
// Solidity: function calculateReward(address user) view returns(uint256)
func (_Abi *AbiCallerSession) CalculateReward(user common.Address) (*big.Int, error) {
	return _Abi.Contract.CalculateReward(&_Abi.CallOpts, user)
}

// DelayDuration is a free data retrieval call binding the contract method 0xb5984a36.
//
// Solidity: function delayDuration() view returns(uint256)
func (_Abi *AbiCaller) DelayDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "delayDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayDuration is a free data retrieval call binding the contract method 0xb5984a36.
//
// Solidity: function delayDuration() view returns(uint256)
func (_Abi *AbiSession) DelayDuration() (*big.Int, error) {
	return _Abi.Contract.DelayDuration(&_Abi.CallOpts)
}

// DelayDuration is a free data retrieval call binding the contract method 0xb5984a36.
//
// Solidity: function delayDuration() view returns(uint256)
func (_Abi *AbiCallerSession) DelayDuration() (*big.Int, error) {
	return _Abi.Contract.DelayDuration(&_Abi.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Abi *AbiCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Abi *AbiSession) FeeRate() (*big.Int, error) {
	return _Abi.Contract.FeeRate(&_Abi.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Abi *AbiCallerSession) FeeRate() (*big.Int, error) {
	return _Abi.Contract.FeeRate(&_Abi.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Abi *AbiCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Abi *AbiSession) FeeReceiver() (common.Address, error) {
	return _Abi.Contract.FeeReceiver(&_Abi.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Abi *AbiCallerSession) FeeReceiver() (common.Address, error) {
	return _Abi.Contract.FeeReceiver(&_Abi.CallOpts)
}

// GetTotalFees is a free data retrieval call binding the contract method 0x626e1ae7.
//
// Solidity: function getTotalFees() view returns(uint256)
func (_Abi *AbiCaller) GetTotalFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getTotalFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalFees is a free data retrieval call binding the contract method 0x626e1ae7.
//
// Solidity: function getTotalFees() view returns(uint256)
func (_Abi *AbiSession) GetTotalFees() (*big.Int, error) {
	return _Abi.Contract.GetTotalFees(&_Abi.CallOpts)
}

// GetTotalFees is a free data retrieval call binding the contract method 0x626e1ae7.
//
// Solidity: function getTotalFees() view returns(uint256)
func (_Abi *AbiCallerSession) GetTotalFees() (*big.Int, error) {
	return _Abi.Contract.GetTotalFees(&_Abi.CallOpts)
}

// MaxStakePerUser is a free data retrieval call binding the contract method 0x0717d032.
//
// Solidity: function maxStakePerUser() view returns(uint256)
func (_Abi *AbiCaller) MaxStakePerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "maxStakePerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakePerUser is a free data retrieval call binding the contract method 0x0717d032.
//
// Solidity: function maxStakePerUser() view returns(uint256)
func (_Abi *AbiSession) MaxStakePerUser() (*big.Int, error) {
	return _Abi.Contract.MaxStakePerUser(&_Abi.CallOpts)
}

// MaxStakePerUser is a free data retrieval call binding the contract method 0x0717d032.
//
// Solidity: function maxStakePerUser() view returns(uint256)
func (_Abi *AbiCallerSession) MaxStakePerUser() (*big.Int, error) {
	return _Abi.Contract.MaxStakePerUser(&_Abi.CallOpts)
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() view returns(uint256)
func (_Abi *AbiCaller) MinStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "minStakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() view returns(uint256)
func (_Abi *AbiSession) MinStakeDuration() (*big.Int, error) {
	return _Abi.Contract.MinStakeDuration(&_Abi.CallOpts)
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() view returns(uint256)
func (_Abi *AbiCallerSession) MinStakeDuration() (*big.Int, error) {
	return _Abi.Contract.MinStakeDuration(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCallerSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiSession) Paused() (bool, error) {
	return _Abi.Contract.Paused(&_Abi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiCallerSession) Paused() (bool, error) {
	return _Abi.Contract.Paused(&_Abi.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_Abi *AbiCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_Abi *AbiSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.PendingWithdrawals(&_Abi.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_Abi *AbiCallerSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.PendingWithdrawals(&_Abi.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Abi *AbiCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Abi *AbiSession) ProxiableUUID() ([32]byte, error) {
	return _Abi.Contract.ProxiableUUID(&_Abi.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Abi *AbiCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Abi.Contract.ProxiableUUID(&_Abi.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 lastRewardTime, uint256 stakeTime)
func (_Abi *AbiCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount         *big.Int
	LastRewardTime *big.Int
	StakeTime      *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Amount         *big.Int
		LastRewardTime *big.Int
		StakeTime      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 lastRewardTime, uint256 stakeTime)
func (_Abi *AbiSession) Stakes(arg0 common.Address) (struct {
	Amount         *big.Int
	LastRewardTime *big.Int
	StakeTime      *big.Int
}, error) {
	return _Abi.Contract.Stakes(&_Abi.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 lastRewardTime, uint256 stakeTime)
func (_Abi *AbiCallerSession) Stakes(arg0 common.Address) (struct {
	Amount         *big.Int
	LastRewardTime *big.Int
	StakeTime      *big.Int
}, error) {
	return _Abi.Contract.Stakes(&_Abi.CallOpts, arg0)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Abi *AbiCaller) TotalFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "totalFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Abi *AbiSession) TotalFees() (*big.Int, error) {
	return _Abi.Contract.TotalFees(&_Abi.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0x13114a9d.
//
// Solidity: function totalFees() view returns(uint256)
func (_Abi *AbiCallerSession) TotalFees() (*big.Int, error) {
	return _Abi.Contract.TotalFees(&_Abi.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Abi *AbiCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Abi *AbiSession) TotalStaked() (*big.Int, error) {
	return _Abi.Contract.TotalStaked(&_Abi.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Abi *AbiCallerSession) TotalStaked() (*big.Int, error) {
	return _Abi.Contract.TotalStaked(&_Abi.CallOpts)
}

// UnlockTimes is a free data retrieval call binding the contract method 0x6bf807fc.
//
// Solidity: function unlockTimes(address ) view returns(uint256)
func (_Abi *AbiCaller) UnlockTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "unlockTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockTimes is a free data retrieval call binding the contract method 0x6bf807fc.
//
// Solidity: function unlockTimes(address ) view returns(uint256)
func (_Abi *AbiSession) UnlockTimes(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.UnlockTimes(&_Abi.CallOpts, arg0)
}

// UnlockTimes is a free data retrieval call binding the contract method 0x6bf807fc.
//
// Solidity: function unlockTimes(address ) view returns(uint256)
func (_Abi *AbiCallerSession) UnlockTimes(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.UnlockTimes(&_Abi.CallOpts, arg0)
}

// ClaimBossToken is a paid mutator transaction binding the contract method 0x160a47ad.
//
// Solidity: function claimBossToken() returns()
func (_Abi *AbiTransactor) ClaimBossToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "claimBossToken")
}

// ClaimBossToken is a paid mutator transaction binding the contract method 0x160a47ad.
//
// Solidity: function claimBossToken() returns()
func (_Abi *AbiSession) ClaimBossToken() (*types.Transaction, error) {
	return _Abi.Contract.ClaimBossToken(&_Abi.TransactOpts)
}

// ClaimBossToken is a paid mutator transaction binding the contract method 0x160a47ad.
//
// Solidity: function claimBossToken() returns()
func (_Abi *AbiTransactorSession) ClaimBossToken() (*types.Transaction, error) {
	return _Abi.Contract.ClaimBossToken(&_Abi.TransactOpts)
}

// ClaimDelayed is a paid mutator transaction binding the contract method 0x62d8327f.
//
// Solidity: function claimDelayed() returns()
func (_Abi *AbiTransactor) ClaimDelayed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "claimDelayed")
}

// ClaimDelayed is a paid mutator transaction binding the contract method 0x62d8327f.
//
// Solidity: function claimDelayed() returns()
func (_Abi *AbiSession) ClaimDelayed() (*types.Transaction, error) {
	return _Abi.Contract.ClaimDelayed(&_Abi.TransactOpts)
}

// ClaimDelayed is a paid mutator transaction binding the contract method 0x62d8327f.
//
// Solidity: function claimDelayed() returns()
func (_Abi *AbiTransactorSession) ClaimDelayed() (*types.Transaction, error) {
	return _Abi.Contract.ClaimDelayed(&_Abi.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Abi *AbiTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Abi *AbiSession) ClaimFees() (*types.Transaction, error) {
	return _Abi.Contract.ClaimFees(&_Abi.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Abi *AbiTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Abi.Contract.ClaimFees(&_Abi.TransactOpts)
}

// DepositToTreasury is a paid mutator transaction binding the contract method 0x718b2058.
//
// Solidity: function depositToTreasury() payable returns()
func (_Abi *AbiTransactor) DepositToTreasury(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "depositToTreasury")
}

// DepositToTreasury is a paid mutator transaction binding the contract method 0x718b2058.
//
// Solidity: function depositToTreasury() payable returns()
func (_Abi *AbiSession) DepositToTreasury() (*types.Transaction, error) {
	return _Abi.Contract.DepositToTreasury(&_Abi.TransactOpts)
}

// DepositToTreasury is a paid mutator transaction binding the contract method 0x718b2058.
//
// Solidity: function depositToTreasury() payable returns()
func (_Abi *AbiTransactorSession) DepositToTreasury() (*types.Transaction, error) {
	return _Abi.Contract.DepositToTreasury(&_Abi.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _bossToken, address _feeReceiver, uint256 _feeRate, uint256 _delayDuration, uint256 _minStakeDuration, uint256 _maxStakePerUser) returns()
func (_Abi *AbiTransactor) Initialize(opts *bind.TransactOpts, _bossToken common.Address, _feeReceiver common.Address, _feeRate *big.Int, _delayDuration *big.Int, _minStakeDuration *big.Int, _maxStakePerUser *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "initialize", _bossToken, _feeReceiver, _feeRate, _delayDuration, _minStakeDuration, _maxStakePerUser)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _bossToken, address _feeReceiver, uint256 _feeRate, uint256 _delayDuration, uint256 _minStakeDuration, uint256 _maxStakePerUser) returns()
func (_Abi *AbiSession) Initialize(_bossToken common.Address, _feeReceiver common.Address, _feeRate *big.Int, _delayDuration *big.Int, _minStakeDuration *big.Int, _maxStakePerUser *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts, _bossToken, _feeReceiver, _feeRate, _delayDuration, _minStakeDuration, _maxStakePerUser)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _bossToken, address _feeReceiver, uint256 _feeRate, uint256 _delayDuration, uint256 _minStakeDuration, uint256 _maxStakePerUser) returns()
func (_Abi *AbiTransactorSession) Initialize(_bossToken common.Address, _feeReceiver common.Address, _feeRate *big.Int, _delayDuration *big.Int, _minStakeDuration *big.Int, _maxStakePerUser *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts, _bossToken, _feeReceiver, _feeRate, _delayDuration, _minStakeDuration, _maxStakePerUser)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactorSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newReceiver) returns()
func (_Abi *AbiTransactor) SetFeeReceiver(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setFeeReceiver", newReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newReceiver) returns()
func (_Abi *AbiSession) SetFeeReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetFeeReceiver(&_Abi.TransactOpts, newReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newReceiver) returns()
func (_Abi *AbiTransactorSession) SetFeeReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetFeeReceiver(&_Abi.TransactOpts, newReceiver)
}

// SetStakingParameters is a paid mutator transaction binding the contract method 0x4b50c244.
//
// Solidity: function setStakingParameters(uint256 _delayDuration, uint256 _minStakeDuration, uint256 _maxStakePerUser) returns()
func (_Abi *AbiTransactor) SetStakingParameters(opts *bind.TransactOpts, _delayDuration *big.Int, _minStakeDuration *big.Int, _maxStakePerUser *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setStakingParameters", _delayDuration, _minStakeDuration, _maxStakePerUser)
}

// SetStakingParameters is a paid mutator transaction binding the contract method 0x4b50c244.
//
// Solidity: function setStakingParameters(uint256 _delayDuration, uint256 _minStakeDuration, uint256 _maxStakePerUser) returns()
func (_Abi *AbiSession) SetStakingParameters(_delayDuration *big.Int, _minStakeDuration *big.Int, _maxStakePerUser *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetStakingParameters(&_Abi.TransactOpts, _delayDuration, _minStakeDuration, _maxStakePerUser)
}

// SetStakingParameters is a paid mutator transaction binding the contract method 0x4b50c244.
//
// Solidity: function setStakingParameters(uint256 _delayDuration, uint256 _minStakeDuration, uint256 _maxStakePerUser) returns()
func (_Abi *AbiTransactorSession) SetStakingParameters(_delayDuration *big.Int, _minStakeDuration *big.Int, _maxStakePerUser *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetStakingParameters(&_Abi.TransactOpts, _delayDuration, _minStakeDuration, _maxStakePerUser)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Abi *AbiTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Abi *AbiSession) Stake() (*types.Transaction, error) {
	return _Abi.Contract.Stake(&_Abi.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Abi *AbiTransactorSession) Stake() (*types.Transaction, error) {
	return _Abi.Contract.Stake(&_Abi.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactorSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Abi *AbiTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Abi *AbiSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Abi.Contract.UpgradeTo(&_Abi.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Abi *AbiTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Abi.Contract.UpgradeTo(&_Abi.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Abi *AbiTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Abi *AbiSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.UpgradeToAndCall(&_Abi.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Abi *AbiTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.UpgradeToAndCall(&_Abi.TransactOpts, newImplementation, data)
}

// WithdrawFromTreasury is a paid mutator transaction binding the contract method 0x63555624.
//
// Solidity: function withdrawFromTreasury(uint256 amount) returns()
func (_Abi *AbiTransactor) WithdrawFromTreasury(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdrawFromTreasury", amount)
}

// WithdrawFromTreasury is a paid mutator transaction binding the contract method 0x63555624.
//
// Solidity: function withdrawFromTreasury(uint256 amount) returns()
func (_Abi *AbiSession) WithdrawFromTreasury(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawFromTreasury(&_Abi.TransactOpts, amount)
}

// WithdrawFromTreasury is a paid mutator transaction binding the contract method 0x63555624.
//
// Solidity: function withdrawFromTreasury(uint256 amount) returns()
func (_Abi *AbiTransactorSession) WithdrawFromTreasury(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawFromTreasury(&_Abi.TransactOpts, amount)
}

// WithdrawLater is a paid mutator transaction binding the contract method 0x36543a5d.
//
// Solidity: function withdrawLater(uint256 amount) returns()
func (_Abi *AbiTransactor) WithdrawLater(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdrawLater", amount)
}

// WithdrawLater is a paid mutator transaction binding the contract method 0x36543a5d.
//
// Solidity: function withdrawLater(uint256 amount) returns()
func (_Abi *AbiSession) WithdrawLater(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawLater(&_Abi.TransactOpts, amount)
}

// WithdrawLater is a paid mutator transaction binding the contract method 0x36543a5d.
//
// Solidity: function withdrawLater(uint256 amount) returns()
func (_Abi *AbiTransactorSession) WithdrawLater(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawLater(&_Abi.TransactOpts, amount)
}

// WithdrawNow is a paid mutator transaction binding the contract method 0xa5469f85.
//
// Solidity: function withdrawNow(uint256 amount) returns()
func (_Abi *AbiTransactor) WithdrawNow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdrawNow", amount)
}

// WithdrawNow is a paid mutator transaction binding the contract method 0xa5469f85.
//
// Solidity: function withdrawNow(uint256 amount) returns()
func (_Abi *AbiSession) WithdrawNow(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawNow(&_Abi.TransactOpts, amount)
}

// WithdrawNow is a paid mutator transaction binding the contract method 0xa5469f85.
//
// Solidity: function withdrawNow(uint256 amount) returns()
func (_Abi *AbiTransactorSession) WithdrawNow(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawNow(&_Abi.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiSession) Receive() (*types.Transaction, error) {
	return _Abi.Contract.Receive(&_Abi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiTransactorSession) Receive() (*types.Transaction, error) {
	return _Abi.Contract.Receive(&_Abi.TransactOpts)
}

// AbiAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Abi contract.
type AbiAdminChangedIterator struct {
	Event *AbiAdminChanged // Event containing the contract specifics and raw log

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
func (it *AbiAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiAdminChanged)
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
		it.Event = new(AbiAdminChanged)
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
func (it *AbiAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiAdminChanged represents a AdminChanged event raised by the Abi contract.
type AbiAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Abi *AbiFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AbiAdminChangedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AbiAdminChangedIterator{contract: _Abi.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Abi *AbiFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AbiAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiAdminChanged)
				if err := _Abi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Abi *AbiFilterer) ParseAdminChanged(log types.Log) (*AbiAdminChanged, error) {
	event := new(AbiAdminChanged)
	if err := _Abi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Abi contract.
type AbiBeaconUpgradedIterator struct {
	Event *AbiBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AbiBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiBeaconUpgraded)
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
		it.Event = new(AbiBeaconUpgraded)
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
func (it *AbiBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiBeaconUpgraded represents a BeaconUpgraded event raised by the Abi contract.
type AbiBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Abi *AbiFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AbiBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AbiBeaconUpgradedIterator{contract: _Abi.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Abi *AbiFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AbiBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiBeaconUpgraded)
				if err := _Abi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Abi *AbiFilterer) ParseBeaconUpgraded(log types.Log) (*AbiBeaconUpgraded, error) {
	event := new(AbiBeaconUpgraded)
	if err := _Abi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiFeeReceiverChangedIterator is returned from FilterFeeReceiverChanged and is used to iterate over the raw logs and unpacked data for FeeReceiverChanged events raised by the Abi contract.
type AbiFeeReceiverChangedIterator struct {
	Event *AbiFeeReceiverChanged // Event containing the contract specifics and raw log

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
func (it *AbiFeeReceiverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiFeeReceiverChanged)
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
		it.Event = new(AbiFeeReceiverChanged)
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
func (it *AbiFeeReceiverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiFeeReceiverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiFeeReceiverChanged represents a FeeReceiverChanged event raised by the Abi contract.
type AbiFeeReceiverChanged struct {
	PreviousReceiver common.Address
	NewReceiver      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeeReceiverChanged is a free log retrieval operation binding the contract event 0xa4b009cc442411b602eaf94bc0579b6abdb8fd90b4ef5b9426e270038906bd03.
//
// Solidity: event FeeReceiverChanged(address indexed previousReceiver, address indexed newReceiver)
func (_Abi *AbiFilterer) FilterFeeReceiverChanged(opts *bind.FilterOpts, previousReceiver []common.Address, newReceiver []common.Address) (*AbiFeeReceiverChangedIterator, error) {

	var previousReceiverRule []interface{}
	for _, previousReceiverItem := range previousReceiver {
		previousReceiverRule = append(previousReceiverRule, previousReceiverItem)
	}
	var newReceiverRule []interface{}
	for _, newReceiverItem := range newReceiver {
		newReceiverRule = append(newReceiverRule, newReceiverItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "FeeReceiverChanged", previousReceiverRule, newReceiverRule)
	if err != nil {
		return nil, err
	}
	return &AbiFeeReceiverChangedIterator{contract: _Abi.contract, event: "FeeReceiverChanged", logs: logs, sub: sub}, nil
}

// WatchFeeReceiverChanged is a free log subscription operation binding the contract event 0xa4b009cc442411b602eaf94bc0579b6abdb8fd90b4ef5b9426e270038906bd03.
//
// Solidity: event FeeReceiverChanged(address indexed previousReceiver, address indexed newReceiver)
func (_Abi *AbiFilterer) WatchFeeReceiverChanged(opts *bind.WatchOpts, sink chan<- *AbiFeeReceiverChanged, previousReceiver []common.Address, newReceiver []common.Address) (event.Subscription, error) {

	var previousReceiverRule []interface{}
	for _, previousReceiverItem := range previousReceiver {
		previousReceiverRule = append(previousReceiverRule, previousReceiverItem)
	}
	var newReceiverRule []interface{}
	for _, newReceiverItem := range newReceiver {
		newReceiverRule = append(newReceiverRule, newReceiverItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "FeeReceiverChanged", previousReceiverRule, newReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiFeeReceiverChanged)
				if err := _Abi.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
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

// ParseFeeReceiverChanged is a log parse operation binding the contract event 0xa4b009cc442411b602eaf94bc0579b6abdb8fd90b4ef5b9426e270038906bd03.
//
// Solidity: event FeeReceiverChanged(address indexed previousReceiver, address indexed newReceiver)
func (_Abi *AbiFilterer) ParseFeeReceiverChanged(log types.Log) (*AbiFeeReceiverChanged, error) {
	event := new(AbiFeeReceiverChanged)
	if err := _Abi.contract.UnpackLog(event, "FeeReceiverChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the Abi contract.
type AbiFeesClaimedIterator struct {
	Event *AbiFeesClaimed // Event containing the contract specifics and raw log

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
func (it *AbiFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiFeesClaimed)
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
		it.Event = new(AbiFeesClaimed)
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
func (it *AbiFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiFeesClaimed represents a FeesClaimed event raised by the Abi contract.
type AbiFeesClaimed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed receiver, uint256 amount)
func (_Abi *AbiFilterer) FilterFeesClaimed(opts *bind.FilterOpts, receiver []common.Address) (*AbiFeesClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "FeesClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &AbiFeesClaimedIterator{contract: _Abi.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed receiver, uint256 amount)
func (_Abi *AbiFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *AbiFeesClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "FeesClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiFeesClaimed)
				if err := _Abi.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed receiver, uint256 amount)
func (_Abi *AbiFilterer) ParseFeesClaimed(log types.Log) (*AbiFeesClaimed, error) {
	event := new(AbiFeesClaimed)
	if err := _Abi.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Abi contract.
type AbiInitializedIterator struct {
	Event *AbiInitialized // Event containing the contract specifics and raw log

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
func (it *AbiInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInitialized)
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
		it.Event = new(AbiInitialized)
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
func (it *AbiInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInitialized represents a Initialized event raised by the Abi contract.
type AbiInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Abi *AbiFilterer) FilterInitialized(opts *bind.FilterOpts) (*AbiInitializedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AbiInitializedIterator{contract: _Abi.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Abi *AbiFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AbiInitialized) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInitialized)
				if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Abi *AbiFilterer) ParseInitialized(log types.Log) (*AbiInitialized, error) {
	event := new(AbiInitialized)
	if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abi contract.
type AbiOwnershipTransferredIterator struct {
	Event *AbiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOwnershipTransferred)
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
		it.Event = new(AbiOwnershipTransferred)
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
func (it *AbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOwnershipTransferred represents a OwnershipTransferred event raised by the Abi contract.
type AbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbiOwnershipTransferredIterator{contract: _Abi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOwnershipTransferred)
				if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) ParseOwnershipTransferred(log types.Log) (*AbiOwnershipTransferred, error) {
	event := new(AbiOwnershipTransferred)
	if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Abi contract.
type AbiPausedIterator struct {
	Event *AbiPaused // Event containing the contract specifics and raw log

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
func (it *AbiPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPaused)
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
		it.Event = new(AbiPaused)
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
func (it *AbiPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPaused represents a Paused event raised by the Abi contract.
type AbiPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) FilterPaused(opts *bind.FilterOpts) (*AbiPausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AbiPausedIterator{contract: _Abi.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AbiPaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPaused)
				if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) ParsePaused(log types.Log) (*AbiPaused, error) {
	event := new(AbiPaused)
	if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Abi contract.
type AbiRewardClaimedIterator struct {
	Event *AbiRewardClaimed // Event containing the contract specifics and raw log

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
func (it *AbiRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRewardClaimed)
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
		it.Event = new(AbiRewardClaimed)
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
func (it *AbiRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRewardClaimed represents a RewardClaimed event raised by the Abi contract.
type AbiRewardClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 amount)
func (_Abi *AbiFilterer) FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*AbiRewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &AbiRewardClaimedIterator{contract: _Abi.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 amount)
func (_Abi *AbiFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *AbiRewardClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRewardClaimed)
				if err := _Abi.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 amount)
func (_Abi *AbiFilterer) ParseRewardClaimed(log types.Log) (*AbiRewardClaimed, error) {
	event := new(AbiRewardClaimed)
	if err := _Abi.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Abi contract.
type AbiStakedIterator struct {
	Event *AbiStaked // Event containing the contract specifics and raw log

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
func (it *AbiStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiStaked)
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
		it.Event = new(AbiStaked)
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
func (it *AbiStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiStaked represents a Staked event raised by the Abi contract.
type AbiStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Abi *AbiFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*AbiStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &AbiStakedIterator{contract: _Abi.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Abi *AbiFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AbiStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiStaked)
				if err := _Abi.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Abi *AbiFilterer) ParseStaked(log types.Log) (*AbiStaked, error) {
	event := new(AbiStaked)
	if err := _Abi.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTreasuryDepositIterator is returned from FilterTreasuryDeposit and is used to iterate over the raw logs and unpacked data for TreasuryDeposit events raised by the Abi contract.
type AbiTreasuryDepositIterator struct {
	Event *AbiTreasuryDeposit // Event containing the contract specifics and raw log

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
func (it *AbiTreasuryDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTreasuryDeposit)
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
		it.Event = new(AbiTreasuryDeposit)
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
func (it *AbiTreasuryDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTreasuryDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTreasuryDeposit represents a TreasuryDeposit event raised by the Abi contract.
type AbiTreasuryDeposit struct {
	Owner     common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTreasuryDeposit is a free log retrieval operation binding the contract event 0xb68e8301ad7468a283824a6592dd711e09b95af56caa96757e636a4a8588bf29.
//
// Solidity: event TreasuryDeposit(address indexed owner, uint256 amount, uint256 timestamp)
func (_Abi *AbiFilterer) FilterTreasuryDeposit(opts *bind.FilterOpts, owner []common.Address) (*AbiTreasuryDepositIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TreasuryDeposit", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AbiTreasuryDepositIterator{contract: _Abi.contract, event: "TreasuryDeposit", logs: logs, sub: sub}, nil
}

// WatchTreasuryDeposit is a free log subscription operation binding the contract event 0xb68e8301ad7468a283824a6592dd711e09b95af56caa96757e636a4a8588bf29.
//
// Solidity: event TreasuryDeposit(address indexed owner, uint256 amount, uint256 timestamp)
func (_Abi *AbiFilterer) WatchTreasuryDeposit(opts *bind.WatchOpts, sink chan<- *AbiTreasuryDeposit, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TreasuryDeposit", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTreasuryDeposit)
				if err := _Abi.contract.UnpackLog(event, "TreasuryDeposit", log); err != nil {
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

// ParseTreasuryDeposit is a log parse operation binding the contract event 0xb68e8301ad7468a283824a6592dd711e09b95af56caa96757e636a4a8588bf29.
//
// Solidity: event TreasuryDeposit(address indexed owner, uint256 amount, uint256 timestamp)
func (_Abi *AbiFilterer) ParseTreasuryDeposit(log types.Log) (*AbiTreasuryDeposit, error) {
	event := new(AbiTreasuryDeposit)
	if err := _Abi.contract.UnpackLog(event, "TreasuryDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTreasuryWithdrawIterator is returned from FilterTreasuryWithdraw and is used to iterate over the raw logs and unpacked data for TreasuryWithdraw events raised by the Abi contract.
type AbiTreasuryWithdrawIterator struct {
	Event *AbiTreasuryWithdraw // Event containing the contract specifics and raw log

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
func (it *AbiTreasuryWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTreasuryWithdraw)
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
		it.Event = new(AbiTreasuryWithdraw)
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
func (it *AbiTreasuryWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTreasuryWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTreasuryWithdraw represents a TreasuryWithdraw event raised by the Abi contract.
type AbiTreasuryWithdraw struct {
	Owner     common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTreasuryWithdraw is a free log retrieval operation binding the contract event 0x9134665a0d95d8d6af858b05eae0e9fd650ad93e9fa6620f27e35d011a358afb.
//
// Solidity: event TreasuryWithdraw(address indexed owner, uint256 amount, uint256 timestamp)
func (_Abi *AbiFilterer) FilterTreasuryWithdraw(opts *bind.FilterOpts, owner []common.Address) (*AbiTreasuryWithdrawIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TreasuryWithdraw", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AbiTreasuryWithdrawIterator{contract: _Abi.contract, event: "TreasuryWithdraw", logs: logs, sub: sub}, nil
}

// WatchTreasuryWithdraw is a free log subscription operation binding the contract event 0x9134665a0d95d8d6af858b05eae0e9fd650ad93e9fa6620f27e35d011a358afb.
//
// Solidity: event TreasuryWithdraw(address indexed owner, uint256 amount, uint256 timestamp)
func (_Abi *AbiFilterer) WatchTreasuryWithdraw(opts *bind.WatchOpts, sink chan<- *AbiTreasuryWithdraw, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TreasuryWithdraw", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTreasuryWithdraw)
				if err := _Abi.contract.UnpackLog(event, "TreasuryWithdraw", log); err != nil {
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

// ParseTreasuryWithdraw is a log parse operation binding the contract event 0x9134665a0d95d8d6af858b05eae0e9fd650ad93e9fa6620f27e35d011a358afb.
//
// Solidity: event TreasuryWithdraw(address indexed owner, uint256 amount, uint256 timestamp)
func (_Abi *AbiFilterer) ParseTreasuryWithdraw(log types.Log) (*AbiTreasuryWithdraw, error) {
	event := new(AbiTreasuryWithdraw)
	if err := _Abi.contract.UnpackLog(event, "TreasuryWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Abi contract.
type AbiUnpausedIterator struct {
	Event *AbiUnpaused // Event containing the contract specifics and raw log

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
func (it *AbiUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUnpaused)
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
		it.Event = new(AbiUnpaused)
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
func (it *AbiUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUnpaused represents a Unpaused event raised by the Abi contract.
type AbiUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AbiUnpausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AbiUnpausedIterator{contract: _Abi.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AbiUnpaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUnpaused)
				if err := _Abi.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) ParseUnpaused(log types.Log) (*AbiUnpaused, error) {
	event := new(AbiUnpaused)
	if err := _Abi.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Abi contract.
type AbiUpgradedIterator struct {
	Event *AbiUpgraded // Event containing the contract specifics and raw log

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
func (it *AbiUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUpgraded)
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
		it.Event = new(AbiUpgraded)
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
func (it *AbiUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUpgraded represents a Upgraded event raised by the Abi contract.
type AbiUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AbiUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AbiUpgradedIterator{contract: _Abi.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AbiUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUpgraded)
				if err := _Abi.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) ParseUpgraded(log types.Log) (*AbiUpgraded, error) {
	event := new(AbiUpgraded)
	if err := _Abi.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Abi contract.
type AbiWithdrawnIterator struct {
	Event *AbiWithdrawn // Event containing the contract specifics and raw log

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
func (it *AbiWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiWithdrawn)
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
		it.Event = new(AbiWithdrawn)
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
func (it *AbiWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiWithdrawn represents a Withdrawn event raised by the Abi contract.
type AbiWithdrawn struct {
	User      common.Address
	Amount    *big.Int
	IsInstant bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, bool isInstant)
func (_Abi *AbiFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*AbiWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &AbiWithdrawnIterator{contract: _Abi.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, bool isInstant)
func (_Abi *AbiFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AbiWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiWithdrawn)
				if err := _Abi.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, bool isInstant)
func (_Abi *AbiFilterer) ParseWithdrawn(log types.Log) (*AbiWithdrawn, error) {
	event := new(AbiWithdrawn)
	if err := _Abi.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
