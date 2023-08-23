// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pandasia

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PandasiaAirdrop is an auto generated low-level Go binding around an user-defined struct.
type PandasiaAirdrop struct {
	Owner   common.Address
	Erc20   common.Address
	Balance *big.Int
	Amount  *big.Int
	Expires *big.Int
	Memo    [32]byte
}

// PandasiaMetaData contains all meta data concerning the Pandasia contract.
var PandasiaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AddressNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AirdropOutOfFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PAddrNotInValidatorMerkleTree\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"airdropCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"airdropIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"airdrops\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"memo\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"c2p\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"airdropId\",\"type\":\"uint256\"}],\"name\":\"claimAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"airdropId\",\"type\":\"uint256\"}],\"name\":\"claimAirdropWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"airdropId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmt\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"airdropId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"airdropId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"fundAirdropWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getAirdropIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAirdrops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"memo\",\"type\":\"bytes32\"}],\"internalType\":\"structPandasia.Airdrop[]\",\"name\":\"pageOfAirdrops\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"hashChecksummedMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isRegisteredValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"memo\",\"type\":\"bytes32\"}],\"name\":\"newAirdrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"pubKeyBytesToAvaAddressBytes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"registerPChainAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"setRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"airdropId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmt\",\"type\":\"uint256\"}],\"name\":\"withdrawFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PandasiaABI is the input ABI used to generate the binding from.
// Deprecated: Use PandasiaMetaData.ABI instead.
var PandasiaABI = PandasiaMetaData.ABI

// Pandasia is an auto generated Go binding around an Ethereum contract.
type Pandasia struct {
	PandasiaCaller     // Read-only binding to the contract
	PandasiaTransactor // Write-only binding to the contract
	PandasiaFilterer   // Log filterer for contract events
}

// PandasiaCaller is an auto generated read-only Go binding around an Ethereum contract.
type PandasiaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PandasiaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PandasiaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PandasiaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PandasiaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PandasiaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PandasiaSession struct {
	Contract     *Pandasia         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PandasiaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PandasiaCallerSession struct {
	Contract *PandasiaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PandasiaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PandasiaTransactorSession struct {
	Contract     *PandasiaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PandasiaRaw is an auto generated low-level Go binding around an Ethereum contract.
type PandasiaRaw struct {
	Contract *Pandasia // Generic contract binding to access the raw methods on
}

// PandasiaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PandasiaCallerRaw struct {
	Contract *PandasiaCaller // Generic read-only contract binding to access the raw methods on
}

// PandasiaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PandasiaTransactorRaw struct {
	Contract *PandasiaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPandasia creates a new instance of Pandasia, bound to a specific deployed contract.
func NewPandasia(address common.Address, backend bind.ContractBackend) (*Pandasia, error) {
	contract, err := bindPandasia(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pandasia{PandasiaCaller: PandasiaCaller{contract: contract}, PandasiaTransactor: PandasiaTransactor{contract: contract}, PandasiaFilterer: PandasiaFilterer{contract: contract}}, nil
}

// NewPandasiaCaller creates a new read-only instance of Pandasia, bound to a specific deployed contract.
func NewPandasiaCaller(address common.Address, caller bind.ContractCaller) (*PandasiaCaller, error) {
	contract, err := bindPandasia(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PandasiaCaller{contract: contract}, nil
}

// NewPandasiaTransactor creates a new write-only instance of Pandasia, bound to a specific deployed contract.
func NewPandasiaTransactor(address common.Address, transactor bind.ContractTransactor) (*PandasiaTransactor, error) {
	contract, err := bindPandasia(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PandasiaTransactor{contract: contract}, nil
}

// NewPandasiaFilterer creates a new log filterer instance of Pandasia, bound to a specific deployed contract.
func NewPandasiaFilterer(address common.Address, filterer bind.ContractFilterer) (*PandasiaFilterer, error) {
	contract, err := bindPandasia(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PandasiaFilterer{contract: contract}, nil
}

// bindPandasia binds a generic wrapper to an already deployed contract.
func bindPandasia(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PandasiaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pandasia *PandasiaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pandasia.Contract.PandasiaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pandasia *PandasiaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pandasia.Contract.PandasiaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pandasia *PandasiaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pandasia.Contract.PandasiaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pandasia *PandasiaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pandasia.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pandasia *PandasiaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pandasia.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pandasia *PandasiaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pandasia.Contract.contract.Transact(opts, method, params...)
}

// AirdropCount is a free data retrieval call binding the contract method 0xe952f74f.
//
// Solidity: function airdropCount() view returns(uint256)
func (_Pandasia *PandasiaCaller) AirdropCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "airdropCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirdropCount is a free data retrieval call binding the contract method 0xe952f74f.
//
// Solidity: function airdropCount() view returns(uint256)
func (_Pandasia *PandasiaSession) AirdropCount() (*big.Int, error) {
	return _Pandasia.Contract.AirdropCount(&_Pandasia.CallOpts)
}

// AirdropCount is a free data retrieval call binding the contract method 0xe952f74f.
//
// Solidity: function airdropCount() view returns(uint256)
func (_Pandasia *PandasiaCallerSession) AirdropCount() (*big.Int, error) {
	return _Pandasia.Contract.AirdropCount(&_Pandasia.CallOpts)
}

// AirdropIds is a free data retrieval call binding the contract method 0x566851e9.
//
// Solidity: function airdropIds(address , uint256 ) view returns(uint256)
func (_Pandasia *PandasiaCaller) AirdropIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "airdropIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirdropIds is a free data retrieval call binding the contract method 0x566851e9.
//
// Solidity: function airdropIds(address , uint256 ) view returns(uint256)
func (_Pandasia *PandasiaSession) AirdropIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Pandasia.Contract.AirdropIds(&_Pandasia.CallOpts, arg0, arg1)
}

// AirdropIds is a free data retrieval call binding the contract method 0x566851e9.
//
// Solidity: function airdropIds(address , uint256 ) view returns(uint256)
func (_Pandasia *PandasiaCallerSession) AirdropIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Pandasia.Contract.AirdropIds(&_Pandasia.CallOpts, arg0, arg1)
}

// Airdrops is a free data retrieval call binding the contract method 0x60db5082.
//
// Solidity: function airdrops(uint256 ) view returns(address owner, address erc20, uint256 balance, uint256 amount, uint256 expires, bytes32 memo)
func (_Pandasia *PandasiaCaller) Airdrops(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner   common.Address
	Erc20   common.Address
	Balance *big.Int
	Amount  *big.Int
	Expires *big.Int
	Memo    [32]byte
}, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "airdrops", arg0)

	outstruct := new(struct {
		Owner   common.Address
		Erc20   common.Address
		Balance *big.Int
		Amount  *big.Int
		Expires *big.Int
		Memo    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Erc20 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Balance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Expires = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Memo = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Airdrops is a free data retrieval call binding the contract method 0x60db5082.
//
// Solidity: function airdrops(uint256 ) view returns(address owner, address erc20, uint256 balance, uint256 amount, uint256 expires, bytes32 memo)
func (_Pandasia *PandasiaSession) Airdrops(arg0 *big.Int) (struct {
	Owner   common.Address
	Erc20   common.Address
	Balance *big.Int
	Amount  *big.Int
	Expires *big.Int
	Memo    [32]byte
}, error) {
	return _Pandasia.Contract.Airdrops(&_Pandasia.CallOpts, arg0)
}

// Airdrops is a free data retrieval call binding the contract method 0x60db5082.
//
// Solidity: function airdrops(uint256 ) view returns(address owner, address erc20, uint256 balance, uint256 amount, uint256 expires, bytes32 memo)
func (_Pandasia *PandasiaCallerSession) Airdrops(arg0 *big.Int) (struct {
	Owner   common.Address
	Erc20   common.Address
	Balance *big.Int
	Amount  *big.Int
	Expires *big.Int
	Memo    [32]byte
}, error) {
	return _Pandasia.Contract.Airdrops(&_Pandasia.CallOpts, arg0)
}

// C2p is a free data retrieval call binding the contract method 0xc421e5dc.
//
// Solidity: function c2p(address ) view returns(address)
func (_Pandasia *PandasiaCaller) C2p(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "c2p", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// C2p is a free data retrieval call binding the contract method 0xc421e5dc.
//
// Solidity: function c2p(address ) view returns(address)
func (_Pandasia *PandasiaSession) C2p(arg0 common.Address) (common.Address, error) {
	return _Pandasia.Contract.C2p(&_Pandasia.CallOpts, arg0)
}

// C2p is a free data retrieval call binding the contract method 0xc421e5dc.
//
// Solidity: function c2p(address ) view returns(address)
func (_Pandasia *PandasiaCallerSession) C2p(arg0 common.Address) (common.Address, error) {
	return _Pandasia.Contract.C2p(&_Pandasia.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Pandasia *PandasiaCaller) Claimed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "claimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Pandasia *PandasiaSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Pandasia.Contract.Claimed(&_Pandasia.CallOpts, arg0, arg1)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Pandasia *PandasiaCallerSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Pandasia.Contract.Claimed(&_Pandasia.CallOpts, arg0, arg1)
}

// GetAirdropIds is a free data retrieval call binding the contract method 0x58dd78c5.
//
// Solidity: function getAirdropIds(address owner) view returns(uint256[])
func (_Pandasia *PandasiaCaller) GetAirdropIds(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getAirdropIds", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAirdropIds is a free data retrieval call binding the contract method 0x58dd78c5.
//
// Solidity: function getAirdropIds(address owner) view returns(uint256[])
func (_Pandasia *PandasiaSession) GetAirdropIds(owner common.Address) ([]*big.Int, error) {
	return _Pandasia.Contract.GetAirdropIds(&_Pandasia.CallOpts, owner)
}

// GetAirdropIds is a free data retrieval call binding the contract method 0x58dd78c5.
//
// Solidity: function getAirdropIds(address owner) view returns(uint256[])
func (_Pandasia *PandasiaCallerSession) GetAirdropIds(owner common.Address) ([]*big.Int, error) {
	return _Pandasia.Contract.GetAirdropIds(&_Pandasia.CallOpts, owner)
}

// GetAirdrops is a free data retrieval call binding the contract method 0x6ae1a80e.
//
// Solidity: function getAirdrops(uint256 offset, uint256 limit) view returns((address,address,uint256,uint256,uint256,bytes32)[] pageOfAirdrops)
func (_Pandasia *PandasiaCaller) GetAirdrops(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]PandasiaAirdrop, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getAirdrops", offset, limit)

	if err != nil {
		return *new([]PandasiaAirdrop), err
	}

	out0 := *abi.ConvertType(out[0], new([]PandasiaAirdrop)).(*[]PandasiaAirdrop)

	return out0, err

}

// GetAirdrops is a free data retrieval call binding the contract method 0x6ae1a80e.
//
// Solidity: function getAirdrops(uint256 offset, uint256 limit) view returns((address,address,uint256,uint256,uint256,bytes32)[] pageOfAirdrops)
func (_Pandasia *PandasiaSession) GetAirdrops(offset *big.Int, limit *big.Int) ([]PandasiaAirdrop, error) {
	return _Pandasia.Contract.GetAirdrops(&_Pandasia.CallOpts, offset, limit)
}

// GetAirdrops is a free data retrieval call binding the contract method 0x6ae1a80e.
//
// Solidity: function getAirdrops(uint256 offset, uint256 limit) view returns((address,address,uint256,uint256,uint256,bytes32)[] pageOfAirdrops)
func (_Pandasia *PandasiaCallerSession) GetAirdrops(offset *big.Int, limit *big.Int) ([]PandasiaAirdrop, error) {
	return _Pandasia.Contract.GetAirdrops(&_Pandasia.CallOpts, offset, limit)
}

// HashChecksummedMessage is a free data retrieval call binding the contract method 0xf7bdccb6.
//
// Solidity: function hashChecksummedMessage(address addr) pure returns(bytes32)
func (_Pandasia *PandasiaCaller) HashChecksummedMessage(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "hashChecksummedMessage", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashChecksummedMessage is a free data retrieval call binding the contract method 0xf7bdccb6.
//
// Solidity: function hashChecksummedMessage(address addr) pure returns(bytes32)
func (_Pandasia *PandasiaSession) HashChecksummedMessage(addr common.Address) ([32]byte, error) {
	return _Pandasia.Contract.HashChecksummedMessage(&_Pandasia.CallOpts, addr)
}

// HashChecksummedMessage is a free data retrieval call binding the contract method 0xf7bdccb6.
//
// Solidity: function hashChecksummedMessage(address addr) pure returns(bytes32)
func (_Pandasia *PandasiaCallerSession) HashChecksummedMessage(addr common.Address) ([32]byte, error) {
	return _Pandasia.Contract.HashChecksummedMessage(&_Pandasia.CallOpts, addr)
}

// IsRegisteredValidator is a free data retrieval call binding the contract method 0x6ed14c27.
//
// Solidity: function isRegisteredValidator(address addr) view returns(bool)
func (_Pandasia *PandasiaCaller) IsRegisteredValidator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "isRegisteredValidator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredValidator is a free data retrieval call binding the contract method 0x6ed14c27.
//
// Solidity: function isRegisteredValidator(address addr) view returns(bool)
func (_Pandasia *PandasiaSession) IsRegisteredValidator(addr common.Address) (bool, error) {
	return _Pandasia.Contract.IsRegisteredValidator(&_Pandasia.CallOpts, addr)
}

// IsRegisteredValidator is a free data retrieval call binding the contract method 0x6ed14c27.
//
// Solidity: function isRegisteredValidator(address addr) view returns(bool)
func (_Pandasia *PandasiaCallerSession) IsRegisteredValidator(addr common.Address) (bool, error) {
	return _Pandasia.Contract.IsRegisteredValidator(&_Pandasia.CallOpts, addr)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Pandasia *PandasiaCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Pandasia *PandasiaSession) MerkleRoot() ([32]byte, error) {
	return _Pandasia.Contract.MerkleRoot(&_Pandasia.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Pandasia *PandasiaCallerSession) MerkleRoot() ([32]byte, error) {
	return _Pandasia.Contract.MerkleRoot(&_Pandasia.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pandasia *PandasiaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pandasia *PandasiaSession) Owner() (common.Address, error) {
	return _Pandasia.Contract.Owner(&_Pandasia.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pandasia *PandasiaCallerSession) Owner() (common.Address, error) {
	return _Pandasia.Contract.Owner(&_Pandasia.CallOpts)
}

// PubKeyBytesToAvaAddressBytes is a free data retrieval call binding the contract method 0x42da5791.
//
// Solidity: function pubKeyBytesToAvaAddressBytes(uint256 x, uint256 y) pure returns(address)
func (_Pandasia *PandasiaCaller) PubKeyBytesToAvaAddressBytes(opts *bind.CallOpts, x *big.Int, y *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "pubKeyBytesToAvaAddressBytes", x, y)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubKeyBytesToAvaAddressBytes is a free data retrieval call binding the contract method 0x42da5791.
//
// Solidity: function pubKeyBytesToAvaAddressBytes(uint256 x, uint256 y) pure returns(address)
func (_Pandasia *PandasiaSession) PubKeyBytesToAvaAddressBytes(x *big.Int, y *big.Int) (common.Address, error) {
	return _Pandasia.Contract.PubKeyBytesToAvaAddressBytes(&_Pandasia.CallOpts, x, y)
}

// PubKeyBytesToAvaAddressBytes is a free data retrieval call binding the contract method 0x42da5791.
//
// Solidity: function pubKeyBytesToAvaAddressBytes(uint256 x, uint256 y) pure returns(address)
func (_Pandasia *PandasiaCallerSession) PubKeyBytesToAvaAddressBytes(x *big.Int, y *big.Int) (common.Address, error) {
	return _Pandasia.Contract.PubKeyBytesToAvaAddressBytes(&_Pandasia.CallOpts, x, y)
}

// Verify is a free data retrieval call binding the contract method 0x9a99fcf3.
//
// Solidity: function verify(bytes32 root, address account, bytes32[] proof) pure returns(bool)
func (_Pandasia *PandasiaCaller) Verify(opts *bind.CallOpts, root [32]byte, account common.Address, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "verify", root, account, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x9a99fcf3.
//
// Solidity: function verify(bytes32 root, address account, bytes32[] proof) pure returns(bool)
func (_Pandasia *PandasiaSession) Verify(root [32]byte, account common.Address, proof [][32]byte) (bool, error) {
	return _Pandasia.Contract.Verify(&_Pandasia.CallOpts, root, account, proof)
}

// Verify is a free data retrieval call binding the contract method 0x9a99fcf3.
//
// Solidity: function verify(bytes32 root, address account, bytes32[] proof) pure returns(bool)
func (_Pandasia *PandasiaCallerSession) Verify(root [32]byte, account common.Address, proof [][32]byte) (bool, error) {
	return _Pandasia.Contract.Verify(&_Pandasia.CallOpts, root, account, proof)
}

// ClaimAirdrop is a paid mutator transaction binding the contract method 0xe30d4440.
//
// Solidity: function claimAirdrop(uint256 airdropId) returns()
func (_Pandasia *PandasiaTransactor) ClaimAirdrop(opts *bind.TransactOpts, airdropId *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "claimAirdrop", airdropId)
}

// ClaimAirdrop is a paid mutator transaction binding the contract method 0xe30d4440.
//
// Solidity: function claimAirdrop(uint256 airdropId) returns()
func (_Pandasia *PandasiaSession) ClaimAirdrop(airdropId *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.ClaimAirdrop(&_Pandasia.TransactOpts, airdropId)
}

// ClaimAirdrop is a paid mutator transaction binding the contract method 0xe30d4440.
//
// Solidity: function claimAirdrop(uint256 airdropId) returns()
func (_Pandasia *PandasiaTransactorSession) ClaimAirdrop(airdropId *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.ClaimAirdrop(&_Pandasia.TransactOpts, airdropId)
}

// ClaimAirdropWithPermit is a paid mutator transaction binding the contract method 0xfdc6b24c.
//
// Solidity: function claimAirdropWithPermit(uint256 airdropId) returns()
func (_Pandasia *PandasiaTransactor) ClaimAirdropWithPermit(opts *bind.TransactOpts, airdropId *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "claimAirdropWithPermit", airdropId)
}

// ClaimAirdropWithPermit is a paid mutator transaction binding the contract method 0xfdc6b24c.
//
// Solidity: function claimAirdropWithPermit(uint256 airdropId) returns()
func (_Pandasia *PandasiaSession) ClaimAirdropWithPermit(airdropId *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.ClaimAirdropWithPermit(&_Pandasia.TransactOpts, airdropId)
}

// ClaimAirdropWithPermit is a paid mutator transaction binding the contract method 0xfdc6b24c.
//
// Solidity: function claimAirdropWithPermit(uint256 airdropId) returns()
func (_Pandasia *PandasiaTransactorSession) ClaimAirdropWithPermit(airdropId *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.ClaimAirdropWithPermit(&_Pandasia.TransactOpts, airdropId)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x1f276b6e.
//
// Solidity: function emergencyWithdraw(uint256 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactor) EmergencyWithdraw(opts *bind.TransactOpts, airdropId *big.Int, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "emergencyWithdraw", airdropId, withdrawAmt)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x1f276b6e.
//
// Solidity: function emergencyWithdraw(uint256 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaSession) EmergencyWithdraw(airdropId *big.Int, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.EmergencyWithdraw(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x1f276b6e.
//
// Solidity: function emergencyWithdraw(uint256 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactorSession) EmergencyWithdraw(airdropId *big.Int, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.EmergencyWithdraw(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// FundAirdrop is a paid mutator transaction binding the contract method 0xcb0d3561.
//
// Solidity: function fundAirdrop(uint256 airdropId, uint256 amount) returns()
func (_Pandasia *PandasiaTransactor) FundAirdrop(opts *bind.TransactOpts, airdropId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "fundAirdrop", airdropId, amount)
}

// FundAirdrop is a paid mutator transaction binding the contract method 0xcb0d3561.
//
// Solidity: function fundAirdrop(uint256 airdropId, uint256 amount) returns()
func (_Pandasia *PandasiaSession) FundAirdrop(airdropId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.FundAirdrop(&_Pandasia.TransactOpts, airdropId, amount)
}

// FundAirdrop is a paid mutator transaction binding the contract method 0xcb0d3561.
//
// Solidity: function fundAirdrop(uint256 airdropId, uint256 amount) returns()
func (_Pandasia *PandasiaTransactorSession) FundAirdrop(airdropId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.FundAirdrop(&_Pandasia.TransactOpts, airdropId, amount)
}

// FundAirdropWithPermit is a paid mutator transaction binding the contract method 0x002c21c6.
//
// Solidity: function fundAirdropWithPermit(uint256 airdropId, uint256 amount, address funder, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pandasia *PandasiaTransactor) FundAirdropWithPermit(opts *bind.TransactOpts, airdropId *big.Int, amount *big.Int, funder common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "fundAirdropWithPermit", airdropId, amount, funder, deadline, v, r, s)
}

// FundAirdropWithPermit is a paid mutator transaction binding the contract method 0x002c21c6.
//
// Solidity: function fundAirdropWithPermit(uint256 airdropId, uint256 amount, address funder, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pandasia *PandasiaSession) FundAirdropWithPermit(airdropId *big.Int, amount *big.Int, funder common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.FundAirdropWithPermit(&_Pandasia.TransactOpts, airdropId, amount, funder, deadline, v, r, s)
}

// FundAirdropWithPermit is a paid mutator transaction binding the contract method 0x002c21c6.
//
// Solidity: function fundAirdropWithPermit(uint256 airdropId, uint256 amount, address funder, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pandasia *PandasiaTransactorSession) FundAirdropWithPermit(airdropId *big.Int, amount *big.Int, funder common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.FundAirdropWithPermit(&_Pandasia.TransactOpts, airdropId, amount, funder, deadline, v, r, s)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0xe534f12f.
//
// Solidity: function newAirdrop(address erc20, uint256 amount, uint256 expires, bytes32 memo) returns(uint256)
func (_Pandasia *PandasiaTransactor) NewAirdrop(opts *bind.TransactOpts, erc20 common.Address, amount *big.Int, expires *big.Int, memo [32]byte) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "newAirdrop", erc20, amount, expires, memo)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0xe534f12f.
//
// Solidity: function newAirdrop(address erc20, uint256 amount, uint256 expires, bytes32 memo) returns(uint256)
func (_Pandasia *PandasiaSession) NewAirdrop(erc20 common.Address, amount *big.Int, expires *big.Int, memo [32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.NewAirdrop(&_Pandasia.TransactOpts, erc20, amount, expires, memo)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0xe534f12f.
//
// Solidity: function newAirdrop(address erc20, uint256 amount, uint256 expires, bytes32 memo) returns(uint256)
func (_Pandasia *PandasiaTransactorSession) NewAirdrop(erc20 common.Address, amount *big.Int, expires *big.Int, memo [32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.NewAirdrop(&_Pandasia.TransactOpts, erc20, amount, expires, memo)
}

// RegisterPChainAddr is a paid mutator transaction binding the contract method 0x52f9b6f3.
//
// Solidity: function registerPChainAddr(uint8 v, bytes32 r, bytes32 s, bytes32[] proof) returns()
func (_Pandasia *PandasiaTransactor) RegisterPChainAddr(opts *bind.TransactOpts, v uint8, r [32]byte, s [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "registerPChainAddr", v, r, s, proof)
}

// RegisterPChainAddr is a paid mutator transaction binding the contract method 0x52f9b6f3.
//
// Solidity: function registerPChainAddr(uint8 v, bytes32 r, bytes32 s, bytes32[] proof) returns()
func (_Pandasia *PandasiaSession) RegisterPChainAddr(v uint8, r [32]byte, s [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.RegisterPChainAddr(&_Pandasia.TransactOpts, v, r, s, proof)
}

// RegisterPChainAddr is a paid mutator transaction binding the contract method 0x52f9b6f3.
//
// Solidity: function registerPChainAddr(uint8 v, bytes32 r, bytes32 s, bytes32[] proof) returns()
func (_Pandasia *PandasiaTransactorSession) RegisterPChainAddr(v uint8, r [32]byte, s [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.RegisterPChainAddr(&_Pandasia.TransactOpts, v, r, s, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pandasia *PandasiaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pandasia *PandasiaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pandasia.Contract.RenounceOwnership(&_Pandasia.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pandasia *PandasiaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pandasia.Contract.RenounceOwnership(&_Pandasia.TransactOpts)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 root) returns()
func (_Pandasia *PandasiaTransactor) SetRoot(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "setRoot", root)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 root) returns()
func (_Pandasia *PandasiaSession) SetRoot(root [32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.SetRoot(&_Pandasia.TransactOpts, root)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 root) returns()
func (_Pandasia *PandasiaTransactorSession) SetRoot(root [32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.SetRoot(&_Pandasia.TransactOpts, root)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pandasia *PandasiaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pandasia *PandasiaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.TransferOwnership(&_Pandasia.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pandasia *PandasiaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.TransferOwnership(&_Pandasia.TransactOpts, newOwner)
}

// WithdrawFunding is a paid mutator transaction binding the contract method 0x030040b0.
//
// Solidity: function withdrawFunding(uint256 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactor) WithdrawFunding(opts *bind.TransactOpts, airdropId *big.Int, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "withdrawFunding", airdropId, withdrawAmt)
}

// WithdrawFunding is a paid mutator transaction binding the contract method 0x030040b0.
//
// Solidity: function withdrawFunding(uint256 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaSession) WithdrawFunding(airdropId *big.Int, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.WithdrawFunding(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// WithdrawFunding is a paid mutator transaction binding the contract method 0x030040b0.
//
// Solidity: function withdrawFunding(uint256 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactorSession) WithdrawFunding(airdropId *big.Int, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.WithdrawFunding(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// PandasiaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pandasia contract.
type PandasiaOwnershipTransferredIterator struct {
	Event *PandasiaOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PandasiaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaOwnershipTransferred)
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
		it.Event = new(PandasiaOwnershipTransferred)
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
func (it *PandasiaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaOwnershipTransferred represents a OwnershipTransferred event raised by the Pandasia contract.
type PandasiaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pandasia *PandasiaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PandasiaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PandasiaOwnershipTransferredIterator{contract: _Pandasia.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pandasia *PandasiaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PandasiaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaOwnershipTransferred)
				if err := _Pandasia.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pandasia *PandasiaFilterer) ParseOwnershipTransferred(log types.Log) (*PandasiaOwnershipTransferred, error) {
	event := new(PandasiaOwnershipTransferred)
	if err := _Pandasia.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
