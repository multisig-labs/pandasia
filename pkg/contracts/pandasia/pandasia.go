// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pandasia

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

// PandasiaAirdrop is an auto generated low-level Go binding around an user-defined struct.
type PandasiaAirdrop struct {
	Id          uint64
	Owner       common.Address
	Erc20       common.Address
	Balance     *big.Int
	CustomRoot  [32]byte
	ClaimAmount *big.Int
	StartsAt    uint64
	ExpiresAt   uint64
}

// PandasiaUser is an auto generated low-level Go binding around an user-defined struct.
type PandasiaUser struct {
	CChainAddr common.Address
	PChainAddr common.Address
}

// PandasiaMetaData contains all meta data concerning the Pandasia contract.
var PandasiaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AirdropExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AirdropNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AirdropOutOfFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AirdropStillActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PAddrAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PAddrNotInMerkleTree\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"}],\"name\":\"AirdropClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"AirdropCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_UPDATER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"airdropCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"airdropIds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"airdrops\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"customRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"c2p\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cChainAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cChainAddrsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cChainAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"canClaimAirdrop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"claimAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmt\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePct\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fundAmount\",\"type\":\"uint256\"}],\"name\":\"fundAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"}],\"name\":\"getAirdrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"customRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"internalType\":\"structPandasia.Airdrop\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getAirdropIds\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"getAirdrops\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"customRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"internalType\":\"structPandasia.Airdrop[]\",\"name\":\"pageOfAirdrops\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getRegisteredUsers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cChainAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pChainAddr\",\"type\":\"address\"}],\"internalType\":\"structPandasia.User[]\",\"name\":\"users\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"hasClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"hashChecksummedMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isMinipoolOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isRegisteredValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"customRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"name\":\"newAirdrop\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2c\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"pubKeyBytesToAvaAddressBytes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"recoverMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"registerPChainAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"fee\",\"type\":\"uint32\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setStorageContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterPChainAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"airdropId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmt\",\"type\":\"uint256\"}],\"name\":\"withdrawFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := PandasiaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pandasia *PandasiaCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pandasia *PandasiaSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pandasia.Contract.DEFAULTADMINROLE(&_Pandasia.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pandasia *PandasiaCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pandasia.Contract.DEFAULTADMINROLE(&_Pandasia.CallOpts)
}

// ROOTUPDATER is a free data retrieval call binding the contract method 0xe990dbe8.
//
// Solidity: function ROOT_UPDATER() view returns(bytes32)
func (_Pandasia *PandasiaCaller) ROOTUPDATER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "ROOT_UPDATER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOTUPDATER is a free data retrieval call binding the contract method 0xe990dbe8.
//
// Solidity: function ROOT_UPDATER() view returns(bytes32)
func (_Pandasia *PandasiaSession) ROOTUPDATER() ([32]byte, error) {
	return _Pandasia.Contract.ROOTUPDATER(&_Pandasia.CallOpts)
}

// ROOTUPDATER is a free data retrieval call binding the contract method 0xe990dbe8.
//
// Solidity: function ROOT_UPDATER() view returns(bytes32)
func (_Pandasia *PandasiaCallerSession) ROOTUPDATER() ([32]byte, error) {
	return _Pandasia.Contract.ROOTUPDATER(&_Pandasia.CallOpts)
}

// AirdropCount is a free data retrieval call binding the contract method 0xe952f74f.
//
// Solidity: function airdropCount() view returns(uint64)
func (_Pandasia *PandasiaCaller) AirdropCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "airdropCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AirdropCount is a free data retrieval call binding the contract method 0xe952f74f.
//
// Solidity: function airdropCount() view returns(uint64)
func (_Pandasia *PandasiaSession) AirdropCount() (uint64, error) {
	return _Pandasia.Contract.AirdropCount(&_Pandasia.CallOpts)
}

// AirdropCount is a free data retrieval call binding the contract method 0xe952f74f.
//
// Solidity: function airdropCount() view returns(uint64)
func (_Pandasia *PandasiaCallerSession) AirdropCount() (uint64, error) {
	return _Pandasia.Contract.AirdropCount(&_Pandasia.CallOpts)
}

// AirdropIds is a free data retrieval call binding the contract method 0x566851e9.
//
// Solidity: function airdropIds(address , uint256 ) view returns(uint64)
func (_Pandasia *PandasiaCaller) AirdropIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint64, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "airdropIds", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AirdropIds is a free data retrieval call binding the contract method 0x566851e9.
//
// Solidity: function airdropIds(address , uint256 ) view returns(uint64)
func (_Pandasia *PandasiaSession) AirdropIds(arg0 common.Address, arg1 *big.Int) (uint64, error) {
	return _Pandasia.Contract.AirdropIds(&_Pandasia.CallOpts, arg0, arg1)
}

// AirdropIds is a free data retrieval call binding the contract method 0x566851e9.
//
// Solidity: function airdropIds(address , uint256 ) view returns(uint64)
func (_Pandasia *PandasiaCallerSession) AirdropIds(arg0 common.Address, arg1 *big.Int) (uint64, error) {
	return _Pandasia.Contract.AirdropIds(&_Pandasia.CallOpts, arg0, arg1)
}

// Airdrops is a free data retrieval call binding the contract method 0x0821774f.
//
// Solidity: function airdrops(uint64 ) view returns(uint64 id, address owner, address erc20, uint256 balance, bytes32 customRoot, uint256 claimAmount, uint64 startsAt, uint64 expiresAt)
func (_Pandasia *PandasiaCaller) Airdrops(opts *bind.CallOpts, arg0 uint64) (struct {
	Id          uint64
	Owner       common.Address
	Erc20       common.Address
	Balance     *big.Int
	CustomRoot  [32]byte
	ClaimAmount *big.Int
	StartsAt    uint64
	ExpiresAt   uint64
}, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "airdrops", arg0)

	outstruct := new(struct {
		Id          uint64
		Owner       common.Address
		Erc20       common.Address
		Balance     *big.Int
		CustomRoot  [32]byte
		ClaimAmount *big.Int
		StartsAt    uint64
		ExpiresAt   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Erc20 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Balance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CustomRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.ClaimAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StartsAt = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.ExpiresAt = *abi.ConvertType(out[7], new(uint64)).(*uint64)

	return *outstruct, err

}

// Airdrops is a free data retrieval call binding the contract method 0x0821774f.
//
// Solidity: function airdrops(uint64 ) view returns(uint64 id, address owner, address erc20, uint256 balance, bytes32 customRoot, uint256 claimAmount, uint64 startsAt, uint64 expiresAt)
func (_Pandasia *PandasiaSession) Airdrops(arg0 uint64) (struct {
	Id          uint64
	Owner       common.Address
	Erc20       common.Address
	Balance     *big.Int
	CustomRoot  [32]byte
	ClaimAmount *big.Int
	StartsAt    uint64
	ExpiresAt   uint64
}, error) {
	return _Pandasia.Contract.Airdrops(&_Pandasia.CallOpts, arg0)
}

// Airdrops is a free data retrieval call binding the contract method 0x0821774f.
//
// Solidity: function airdrops(uint64 ) view returns(uint64 id, address owner, address erc20, uint256 balance, bytes32 customRoot, uint256 claimAmount, uint64 startsAt, uint64 expiresAt)
func (_Pandasia *PandasiaCallerSession) Airdrops(arg0 uint64) (struct {
	Id          uint64
	Owner       common.Address
	Erc20       common.Address
	Balance     *big.Int
	CustomRoot  [32]byte
	ClaimAmount *big.Int
	StartsAt    uint64
	ExpiresAt   uint64
}, error) {
	return _Pandasia.Contract.Airdrops(&_Pandasia.CallOpts, arg0)
}

// BlockHeight is a free data retrieval call binding the contract method 0xf44ff712.
//
// Solidity: function blockHeight() view returns(uint64)
func (_Pandasia *PandasiaCaller) BlockHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "blockHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockHeight is a free data retrieval call binding the contract method 0xf44ff712.
//
// Solidity: function blockHeight() view returns(uint64)
func (_Pandasia *PandasiaSession) BlockHeight() (uint64, error) {
	return _Pandasia.Contract.BlockHeight(&_Pandasia.CallOpts)
}

// BlockHeight is a free data retrieval call binding the contract method 0xf44ff712.
//
// Solidity: function blockHeight() view returns(uint64)
func (_Pandasia *PandasiaCallerSession) BlockHeight() (uint64, error) {
	return _Pandasia.Contract.BlockHeight(&_Pandasia.CallOpts)
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

// CChainAddrs is a free data retrieval call binding the contract method 0xb8b1f86e.
//
// Solidity: function cChainAddrs(uint256 ) view returns(address)
func (_Pandasia *PandasiaCaller) CChainAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "cChainAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CChainAddrs is a free data retrieval call binding the contract method 0xb8b1f86e.
//
// Solidity: function cChainAddrs(uint256 ) view returns(address)
func (_Pandasia *PandasiaSession) CChainAddrs(arg0 *big.Int) (common.Address, error) {
	return _Pandasia.Contract.CChainAddrs(&_Pandasia.CallOpts, arg0)
}

// CChainAddrs is a free data retrieval call binding the contract method 0xb8b1f86e.
//
// Solidity: function cChainAddrs(uint256 ) view returns(address)
func (_Pandasia *PandasiaCallerSession) CChainAddrs(arg0 *big.Int) (common.Address, error) {
	return _Pandasia.Contract.CChainAddrs(&_Pandasia.CallOpts, arg0)
}

// CChainAddrsCount is a free data retrieval call binding the contract method 0x88cb71d6.
//
// Solidity: function cChainAddrsCount() view returns(uint256)
func (_Pandasia *PandasiaCaller) CChainAddrsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "cChainAddrsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CChainAddrsCount is a free data retrieval call binding the contract method 0x88cb71d6.
//
// Solidity: function cChainAddrsCount() view returns(uint256)
func (_Pandasia *PandasiaSession) CChainAddrsCount() (*big.Int, error) {
	return _Pandasia.Contract.CChainAddrsCount(&_Pandasia.CallOpts)
}

// CChainAddrsCount is a free data retrieval call binding the contract method 0x88cb71d6.
//
// Solidity: function cChainAddrsCount() view returns(uint256)
func (_Pandasia *PandasiaCallerSession) CChainAddrsCount() (*big.Int, error) {
	return _Pandasia.Contract.CChainAddrsCount(&_Pandasia.CallOpts)
}

// CanClaimAirdrop is a free data retrieval call binding the contract method 0xeb38c3a3.
//
// Solidity: function canClaimAirdrop(address cChainAddr, uint64 airdropId, bytes32[] proof) view returns(bool)
func (_Pandasia *PandasiaCaller) CanClaimAirdrop(opts *bind.CallOpts, cChainAddr common.Address, airdropId uint64, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "canClaimAirdrop", cChainAddr, airdropId, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimAirdrop is a free data retrieval call binding the contract method 0xeb38c3a3.
//
// Solidity: function canClaimAirdrop(address cChainAddr, uint64 airdropId, bytes32[] proof) view returns(bool)
func (_Pandasia *PandasiaSession) CanClaimAirdrop(cChainAddr common.Address, airdropId uint64, proof [][32]byte) (bool, error) {
	return _Pandasia.Contract.CanClaimAirdrop(&_Pandasia.CallOpts, cChainAddr, airdropId, proof)
}

// CanClaimAirdrop is a free data retrieval call binding the contract method 0xeb38c3a3.
//
// Solidity: function canClaimAirdrop(address cChainAddr, uint64 airdropId, bytes32[] proof) view returns(bool)
func (_Pandasia *PandasiaCallerSession) CanClaimAirdrop(cChainAddr common.Address, airdropId uint64, proof [][32]byte) (bool, error) {
	return _Pandasia.Contract.CanClaimAirdrop(&_Pandasia.CallOpts, cChainAddr, airdropId, proof)
}

// Claimed is a free data retrieval call binding the contract method 0xa60628e2.
//
// Solidity: function claimed(uint64 , address ) view returns(bool)
func (_Pandasia *PandasiaCaller) Claimed(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "claimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xa60628e2.
//
// Solidity: function claimed(uint64 , address ) view returns(bool)
func (_Pandasia *PandasiaSession) Claimed(arg0 uint64, arg1 common.Address) (bool, error) {
	return _Pandasia.Contract.Claimed(&_Pandasia.CallOpts, arg0, arg1)
}

// Claimed is a free data retrieval call binding the contract method 0xa60628e2.
//
// Solidity: function claimed(uint64 , address ) view returns(bool)
func (_Pandasia *PandasiaCallerSession) Claimed(arg0 uint64, arg1 common.Address) (bool, error) {
	return _Pandasia.Contract.Claimed(&_Pandasia.CallOpts, arg0, arg1)
}

// FeePct is a free data retrieval call binding the contract method 0xa02cf937.
//
// Solidity: function feePct() view returns(uint32)
func (_Pandasia *PandasiaCaller) FeePct(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "feePct")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeePct is a free data retrieval call binding the contract method 0xa02cf937.
//
// Solidity: function feePct() view returns(uint32)
func (_Pandasia *PandasiaSession) FeePct() (uint32, error) {
	return _Pandasia.Contract.FeePct(&_Pandasia.CallOpts)
}

// FeePct is a free data retrieval call binding the contract method 0xa02cf937.
//
// Solidity: function feePct() view returns(uint32)
func (_Pandasia *PandasiaCallerSession) FeePct() (uint32, error) {
	return _Pandasia.Contract.FeePct(&_Pandasia.CallOpts)
}

// GetAirdrop is a free data retrieval call binding the contract method 0x83caf3e2.
//
// Solidity: function getAirdrop(uint64 airdropId) view returns((uint64,address,address,uint256,bytes32,uint256,uint64,uint64))
func (_Pandasia *PandasiaCaller) GetAirdrop(opts *bind.CallOpts, airdropId uint64) (PandasiaAirdrop, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getAirdrop", airdropId)

	if err != nil {
		return *new(PandasiaAirdrop), err
	}

	out0 := *abi.ConvertType(out[0], new(PandasiaAirdrop)).(*PandasiaAirdrop)

	return out0, err

}

// GetAirdrop is a free data retrieval call binding the contract method 0x83caf3e2.
//
// Solidity: function getAirdrop(uint64 airdropId) view returns((uint64,address,address,uint256,bytes32,uint256,uint64,uint64))
func (_Pandasia *PandasiaSession) GetAirdrop(airdropId uint64) (PandasiaAirdrop, error) {
	return _Pandasia.Contract.GetAirdrop(&_Pandasia.CallOpts, airdropId)
}

// GetAirdrop is a free data retrieval call binding the contract method 0x83caf3e2.
//
// Solidity: function getAirdrop(uint64 airdropId) view returns((uint64,address,address,uint256,bytes32,uint256,uint64,uint64))
func (_Pandasia *PandasiaCallerSession) GetAirdrop(airdropId uint64) (PandasiaAirdrop, error) {
	return _Pandasia.Contract.GetAirdrop(&_Pandasia.CallOpts, airdropId)
}

// GetAirdropIds is a free data retrieval call binding the contract method 0x58dd78c5.
//
// Solidity: function getAirdropIds(address owner) view returns(uint64[])
func (_Pandasia *PandasiaCaller) GetAirdropIds(opts *bind.CallOpts, owner common.Address) ([]uint64, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getAirdropIds", owner)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetAirdropIds is a free data retrieval call binding the contract method 0x58dd78c5.
//
// Solidity: function getAirdropIds(address owner) view returns(uint64[])
func (_Pandasia *PandasiaSession) GetAirdropIds(owner common.Address) ([]uint64, error) {
	return _Pandasia.Contract.GetAirdropIds(&_Pandasia.CallOpts, owner)
}

// GetAirdropIds is a free data retrieval call binding the contract method 0x58dd78c5.
//
// Solidity: function getAirdropIds(address owner) view returns(uint64[])
func (_Pandasia *PandasiaCallerSession) GetAirdropIds(owner common.Address) ([]uint64, error) {
	return _Pandasia.Contract.GetAirdropIds(&_Pandasia.CallOpts, owner)
}

// GetAirdrops is a free data retrieval call binding the contract method 0x74ee3db2.
//
// Solidity: function getAirdrops(uint64 offset, uint64 limit) view returns((uint64,address,address,uint256,bytes32,uint256,uint64,uint64)[] pageOfAirdrops)
func (_Pandasia *PandasiaCaller) GetAirdrops(opts *bind.CallOpts, offset uint64, limit uint64) ([]PandasiaAirdrop, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getAirdrops", offset, limit)

	if err != nil {
		return *new([]PandasiaAirdrop), err
	}

	out0 := *abi.ConvertType(out[0], new([]PandasiaAirdrop)).(*[]PandasiaAirdrop)

	return out0, err

}

// GetAirdrops is a free data retrieval call binding the contract method 0x74ee3db2.
//
// Solidity: function getAirdrops(uint64 offset, uint64 limit) view returns((uint64,address,address,uint256,bytes32,uint256,uint64,uint64)[] pageOfAirdrops)
func (_Pandasia *PandasiaSession) GetAirdrops(offset uint64, limit uint64) ([]PandasiaAirdrop, error) {
	return _Pandasia.Contract.GetAirdrops(&_Pandasia.CallOpts, offset, limit)
}

// GetAirdrops is a free data retrieval call binding the contract method 0x74ee3db2.
//
// Solidity: function getAirdrops(uint64 offset, uint64 limit) view returns((uint64,address,address,uint256,bytes32,uint256,uint64,uint64)[] pageOfAirdrops)
func (_Pandasia *PandasiaCallerSession) GetAirdrops(offset uint64, limit uint64) ([]PandasiaAirdrop, error) {
	return _Pandasia.Contract.GetAirdrops(&_Pandasia.CallOpts, offset, limit)
}

// GetRegisteredUsers is a free data retrieval call binding the contract method 0xd91f2fa5.
//
// Solidity: function getRegisteredUsers(uint256 offset, uint256 limit) view returns((address,address)[] users)
func (_Pandasia *PandasiaCaller) GetRegisteredUsers(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]PandasiaUser, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getRegisteredUsers", offset, limit)

	if err != nil {
		return *new([]PandasiaUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]PandasiaUser)).(*[]PandasiaUser)

	return out0, err

}

// GetRegisteredUsers is a free data retrieval call binding the contract method 0xd91f2fa5.
//
// Solidity: function getRegisteredUsers(uint256 offset, uint256 limit) view returns((address,address)[] users)
func (_Pandasia *PandasiaSession) GetRegisteredUsers(offset *big.Int, limit *big.Int) ([]PandasiaUser, error) {
	return _Pandasia.Contract.GetRegisteredUsers(&_Pandasia.CallOpts, offset, limit)
}

// GetRegisteredUsers is a free data retrieval call binding the contract method 0xd91f2fa5.
//
// Solidity: function getRegisteredUsers(uint256 offset, uint256 limit) view returns((address,address)[] users)
func (_Pandasia *PandasiaCallerSession) GetRegisteredUsers(offset *big.Int, limit *big.Int) ([]PandasiaUser, error) {
	return _Pandasia.Contract.GetRegisteredUsers(&_Pandasia.CallOpts, offset, limit)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pandasia *PandasiaCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pandasia *PandasiaSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pandasia.Contract.GetRoleAdmin(&_Pandasia.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pandasia *PandasiaCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pandasia.Contract.GetRoleAdmin(&_Pandasia.CallOpts, role)
}

// HasClaimed is a free data retrieval call binding the contract method 0xd7a8a2b7.
//
// Solidity: function hasClaimed(uint64 airdropId, address addr) view returns(bool)
func (_Pandasia *PandasiaCaller) HasClaimed(opts *bind.CallOpts, airdropId uint64, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "hasClaimed", airdropId, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClaimed is a free data retrieval call binding the contract method 0xd7a8a2b7.
//
// Solidity: function hasClaimed(uint64 airdropId, address addr) view returns(bool)
func (_Pandasia *PandasiaSession) HasClaimed(airdropId uint64, addr common.Address) (bool, error) {
	return _Pandasia.Contract.HasClaimed(&_Pandasia.CallOpts, airdropId, addr)
}

// HasClaimed is a free data retrieval call binding the contract method 0xd7a8a2b7.
//
// Solidity: function hasClaimed(uint64 airdropId, address addr) view returns(bool)
func (_Pandasia *PandasiaCallerSession) HasClaimed(airdropId uint64, addr common.Address) (bool, error) {
	return _Pandasia.Contract.HasClaimed(&_Pandasia.CallOpts, airdropId, addr)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pandasia *PandasiaCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pandasia *PandasiaSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pandasia.Contract.HasRole(&_Pandasia.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pandasia *PandasiaCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pandasia.Contract.HasRole(&_Pandasia.CallOpts, role, account)
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

// IsMinipoolOperator is a free data retrieval call binding the contract method 0xbf4379f6.
//
// Solidity: function isMinipoolOperator(address addr) view returns(bool)
func (_Pandasia *PandasiaCaller) IsMinipoolOperator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "isMinipoolOperator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinipoolOperator is a free data retrieval call binding the contract method 0xbf4379f6.
//
// Solidity: function isMinipoolOperator(address addr) view returns(bool)
func (_Pandasia *PandasiaSession) IsMinipoolOperator(addr common.Address) (bool, error) {
	return _Pandasia.Contract.IsMinipoolOperator(&_Pandasia.CallOpts, addr)
}

// IsMinipoolOperator is a free data retrieval call binding the contract method 0xbf4379f6.
//
// Solidity: function isMinipoolOperator(address addr) view returns(bool)
func (_Pandasia *PandasiaCallerSession) IsMinipoolOperator(addr common.Address) (bool, error) {
	return _Pandasia.Contract.IsMinipoolOperator(&_Pandasia.CallOpts, addr)
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

// P2c is a free data retrieval call binding the contract method 0x5dd1cf00.
//
// Solidity: function p2c(address ) view returns(address)
func (_Pandasia *PandasiaCaller) P2c(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "p2c", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// P2c is a free data retrieval call binding the contract method 0x5dd1cf00.
//
// Solidity: function p2c(address ) view returns(address)
func (_Pandasia *PandasiaSession) P2c(arg0 common.Address) (common.Address, error) {
	return _Pandasia.Contract.P2c(&_Pandasia.CallOpts, arg0)
}

// P2c is a free data retrieval call binding the contract method 0x5dd1cf00.
//
// Solidity: function p2c(address ) view returns(address)
func (_Pandasia *PandasiaCallerSession) P2c(arg0 common.Address) (common.Address, error) {
	return _Pandasia.Contract.P2c(&_Pandasia.CallOpts, arg0)
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

// RecoverMessage is a free data retrieval call binding the contract method 0x2bd7c773.
//
// Solidity: function recoverMessage(uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Pandasia *PandasiaCaller) RecoverMessage(opts *bind.CallOpts, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "recoverMessage", v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverMessage is a free data retrieval call binding the contract method 0x2bd7c773.
//
// Solidity: function recoverMessage(uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Pandasia *PandasiaSession) RecoverMessage(v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Pandasia.Contract.RecoverMessage(&_Pandasia.CallOpts, v, r, s)
}

// RecoverMessage is a free data retrieval call binding the contract method 0x2bd7c773.
//
// Solidity: function recoverMessage(uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Pandasia *PandasiaCallerSession) RecoverMessage(v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Pandasia.Contract.RecoverMessage(&_Pandasia.CallOpts, v, r, s)
}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_Pandasia *PandasiaCaller) StorageContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "storageContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_Pandasia *PandasiaSession) StorageContract() (common.Address, error) {
	return _Pandasia.Contract.StorageContract(&_Pandasia.CallOpts)
}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_Pandasia *PandasiaCallerSession) StorageContract() (common.Address, error) {
	return _Pandasia.Contract.StorageContract(&_Pandasia.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pandasia *PandasiaCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Pandasia.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pandasia *PandasiaSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pandasia.Contract.SupportsInterface(&_Pandasia.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pandasia *PandasiaCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pandasia.Contract.SupportsInterface(&_Pandasia.CallOpts, interfaceId)
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

// ClaimAirdrop is a paid mutator transaction binding the contract method 0x744e00a0.
//
// Solidity: function claimAirdrop(uint64 airdropId, bytes32[] proof) returns()
func (_Pandasia *PandasiaTransactor) ClaimAirdrop(opts *bind.TransactOpts, airdropId uint64, proof [][32]byte) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "claimAirdrop", airdropId, proof)
}

// ClaimAirdrop is a paid mutator transaction binding the contract method 0x744e00a0.
//
// Solidity: function claimAirdrop(uint64 airdropId, bytes32[] proof) returns()
func (_Pandasia *PandasiaSession) ClaimAirdrop(airdropId uint64, proof [][32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.ClaimAirdrop(&_Pandasia.TransactOpts, airdropId, proof)
}

// ClaimAirdrop is a paid mutator transaction binding the contract method 0x744e00a0.
//
// Solidity: function claimAirdrop(uint64 airdropId, bytes32[] proof) returns()
func (_Pandasia *PandasiaTransactorSession) ClaimAirdrop(airdropId uint64, proof [][32]byte) (*types.Transaction, error) {
	return _Pandasia.Contract.ClaimAirdrop(&_Pandasia.TransactOpts, airdropId, proof)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x30e178cb.
//
// Solidity: function emergencyWithdraw(uint64 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactor) EmergencyWithdraw(opts *bind.TransactOpts, airdropId uint64, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "emergencyWithdraw", airdropId, withdrawAmt)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x30e178cb.
//
// Solidity: function emergencyWithdraw(uint64 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaSession) EmergencyWithdraw(airdropId uint64, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.EmergencyWithdraw(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x30e178cb.
//
// Solidity: function emergencyWithdraw(uint64 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactorSession) EmergencyWithdraw(airdropId uint64, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.EmergencyWithdraw(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// FundAirdrop is a paid mutator transaction binding the contract method 0x92ac15f5.
//
// Solidity: function fundAirdrop(uint64 airdropId, uint256 fundAmount) returns()
func (_Pandasia *PandasiaTransactor) FundAirdrop(opts *bind.TransactOpts, airdropId uint64, fundAmount *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "fundAirdrop", airdropId, fundAmount)
}

// FundAirdrop is a paid mutator transaction binding the contract method 0x92ac15f5.
//
// Solidity: function fundAirdrop(uint64 airdropId, uint256 fundAmount) returns()
func (_Pandasia *PandasiaSession) FundAirdrop(airdropId uint64, fundAmount *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.FundAirdrop(&_Pandasia.TransactOpts, airdropId, fundAmount)
}

// FundAirdrop is a paid mutator transaction binding the contract method 0x92ac15f5.
//
// Solidity: function fundAirdrop(uint64 airdropId, uint256 fundAmount) returns()
func (_Pandasia *PandasiaTransactorSession) FundAirdrop(airdropId uint64, fundAmount *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.FundAirdrop(&_Pandasia.TransactOpts, airdropId, fundAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pandasia *PandasiaTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pandasia *PandasiaSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.GrantRole(&_Pandasia.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pandasia *PandasiaTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.GrantRole(&_Pandasia.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pandasia *PandasiaTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pandasia *PandasiaSession) Initialize() (*types.Transaction, error) {
	return _Pandasia.Contract.Initialize(&_Pandasia.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pandasia *PandasiaTransactorSession) Initialize() (*types.Transaction, error) {
	return _Pandasia.Contract.Initialize(&_Pandasia.TransactOpts)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0xe58681a6.
//
// Solidity: function newAirdrop(bytes32 customRoot, address erc20, uint256 claimAmount, uint64 startsAt, uint64 expiresAt) returns(uint64)
func (_Pandasia *PandasiaTransactor) NewAirdrop(opts *bind.TransactOpts, customRoot [32]byte, erc20 common.Address, claimAmount *big.Int, startsAt uint64, expiresAt uint64) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "newAirdrop", customRoot, erc20, claimAmount, startsAt, expiresAt)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0xe58681a6.
//
// Solidity: function newAirdrop(bytes32 customRoot, address erc20, uint256 claimAmount, uint64 startsAt, uint64 expiresAt) returns(uint64)
func (_Pandasia *PandasiaSession) NewAirdrop(customRoot [32]byte, erc20 common.Address, claimAmount *big.Int, startsAt uint64, expiresAt uint64) (*types.Transaction, error) {
	return _Pandasia.Contract.NewAirdrop(&_Pandasia.TransactOpts, customRoot, erc20, claimAmount, startsAt, expiresAt)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0xe58681a6.
//
// Solidity: function newAirdrop(bytes32 customRoot, address erc20, uint256 claimAmount, uint64 startsAt, uint64 expiresAt) returns(uint64)
func (_Pandasia *PandasiaTransactorSession) NewAirdrop(customRoot [32]byte, erc20 common.Address, claimAmount *big.Int, startsAt uint64, expiresAt uint64) (*types.Transaction, error) {
	return _Pandasia.Contract.NewAirdrop(&_Pandasia.TransactOpts, customRoot, erc20, claimAmount, startsAt, expiresAt)
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

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Pandasia *PandasiaTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Pandasia *PandasiaSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.RenounceRole(&_Pandasia.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Pandasia *PandasiaTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.RenounceRole(&_Pandasia.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pandasia *PandasiaTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pandasia *PandasiaSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.RevokeRole(&_Pandasia.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pandasia *PandasiaTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.RevokeRole(&_Pandasia.TransactOpts, role, account)
}

// SetFee is a paid mutator transaction binding the contract method 0x1ab971ab.
//
// Solidity: function setFee(uint32 fee) returns()
func (_Pandasia *PandasiaTransactor) SetFee(opts *bind.TransactOpts, fee uint32) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "setFee", fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x1ab971ab.
//
// Solidity: function setFee(uint32 fee) returns()
func (_Pandasia *PandasiaSession) SetFee(fee uint32) (*types.Transaction, error) {
	return _Pandasia.Contract.SetFee(&_Pandasia.TransactOpts, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x1ab971ab.
//
// Solidity: function setFee(uint32 fee) returns()
func (_Pandasia *PandasiaTransactorSession) SetFee(fee uint32) (*types.Transaction, error) {
	return _Pandasia.Contract.SetFee(&_Pandasia.TransactOpts, fee)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x3c4440cf.
//
// Solidity: function setMerkleRoot(bytes32 root, uint64 height) returns()
func (_Pandasia *PandasiaTransactor) SetMerkleRoot(opts *bind.TransactOpts, root [32]byte, height uint64) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "setMerkleRoot", root, height)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x3c4440cf.
//
// Solidity: function setMerkleRoot(bytes32 root, uint64 height) returns()
func (_Pandasia *PandasiaSession) SetMerkleRoot(root [32]byte, height uint64) (*types.Transaction, error) {
	return _Pandasia.Contract.SetMerkleRoot(&_Pandasia.TransactOpts, root, height)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x3c4440cf.
//
// Solidity: function setMerkleRoot(bytes32 root, uint64 height) returns()
func (_Pandasia *PandasiaTransactorSession) SetMerkleRoot(root [32]byte, height uint64) (*types.Transaction, error) {
	return _Pandasia.Contract.SetMerkleRoot(&_Pandasia.TransactOpts, root, height)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address addr) returns()
func (_Pandasia *PandasiaTransactor) SetStorageContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "setStorageContract", addr)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address addr) returns()
func (_Pandasia *PandasiaSession) SetStorageContract(addr common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.SetStorageContract(&_Pandasia.TransactOpts, addr)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address addr) returns()
func (_Pandasia *PandasiaTransactorSession) SetStorageContract(addr common.Address) (*types.Transaction, error) {
	return _Pandasia.Contract.SetStorageContract(&_Pandasia.TransactOpts, addr)
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

// UnregisterPChainAddr is a paid mutator transaction binding the contract method 0xcbadade8.
//
// Solidity: function unregisterPChainAddr() returns()
func (_Pandasia *PandasiaTransactor) UnregisterPChainAddr(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "unregisterPChainAddr")
}

// UnregisterPChainAddr is a paid mutator transaction binding the contract method 0xcbadade8.
//
// Solidity: function unregisterPChainAddr() returns()
func (_Pandasia *PandasiaSession) UnregisterPChainAddr() (*types.Transaction, error) {
	return _Pandasia.Contract.UnregisterPChainAddr(&_Pandasia.TransactOpts)
}

// UnregisterPChainAddr is a paid mutator transaction binding the contract method 0xcbadade8.
//
// Solidity: function unregisterPChainAddr() returns()
func (_Pandasia *PandasiaTransactorSession) UnregisterPChainAddr() (*types.Transaction, error) {
	return _Pandasia.Contract.UnregisterPChainAddr(&_Pandasia.TransactOpts)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xb4950bfe.
//
// Solidity: function withdrawFees(uint64 airdropId) returns()
func (_Pandasia *PandasiaTransactor) WithdrawFees(opts *bind.TransactOpts, airdropId uint64) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "withdrawFees", airdropId)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xb4950bfe.
//
// Solidity: function withdrawFees(uint64 airdropId) returns()
func (_Pandasia *PandasiaSession) WithdrawFees(airdropId uint64) (*types.Transaction, error) {
	return _Pandasia.Contract.WithdrawFees(&_Pandasia.TransactOpts, airdropId)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xb4950bfe.
//
// Solidity: function withdrawFees(uint64 airdropId) returns()
func (_Pandasia *PandasiaTransactorSession) WithdrawFees(airdropId uint64) (*types.Transaction, error) {
	return _Pandasia.Contract.WithdrawFees(&_Pandasia.TransactOpts, airdropId)
}

// WithdrawFunding is a paid mutator transaction binding the contract method 0xbe6c5346.
//
// Solidity: function withdrawFunding(uint64 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactor) WithdrawFunding(opts *bind.TransactOpts, airdropId uint64, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.contract.Transact(opts, "withdrawFunding", airdropId, withdrawAmt)
}

// WithdrawFunding is a paid mutator transaction binding the contract method 0xbe6c5346.
//
// Solidity: function withdrawFunding(uint64 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaSession) WithdrawFunding(airdropId uint64, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.WithdrawFunding(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// WithdrawFunding is a paid mutator transaction binding the contract method 0xbe6c5346.
//
// Solidity: function withdrawFunding(uint64 airdropId, uint256 withdrawAmt) returns()
func (_Pandasia *PandasiaTransactorSession) WithdrawFunding(airdropId uint64, withdrawAmt *big.Int) (*types.Transaction, error) {
	return _Pandasia.Contract.WithdrawFunding(&_Pandasia.TransactOpts, airdropId, withdrawAmt)
}

// PandasiaAirdropClaimedIterator is returned from FilterAirdropClaimed and is used to iterate over the raw logs and unpacked data for AirdropClaimed events raised by the Pandasia contract.
type PandasiaAirdropClaimedIterator struct {
	Event *PandasiaAirdropClaimed // Event containing the contract specifics and raw log

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
func (it *PandasiaAirdropClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaAirdropClaimed)
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
		it.Event = new(PandasiaAirdropClaimed)
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
func (it *PandasiaAirdropClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaAirdropClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaAirdropClaimed represents a AirdropClaimed event raised by the Pandasia contract.
type PandasiaAirdropClaimed struct {
	Id       uint64
	Claimant common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAirdropClaimed is a free log retrieval operation binding the contract event 0xf178a0585b458461acedd6bb1c991cda9d7223b0b94b5775d82a5c24166ebf99.
//
// Solidity: event AirdropClaimed(uint64 indexed id, address indexed claimant)
func (_Pandasia *PandasiaFilterer) FilterAirdropClaimed(opts *bind.FilterOpts, id []uint64, claimant []common.Address) (*PandasiaAirdropClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "AirdropClaimed", idRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return &PandasiaAirdropClaimedIterator{contract: _Pandasia.contract, event: "AirdropClaimed", logs: logs, sub: sub}, nil
}

// WatchAirdropClaimed is a free log subscription operation binding the contract event 0xf178a0585b458461acedd6bb1c991cda9d7223b0b94b5775d82a5c24166ebf99.
//
// Solidity: event AirdropClaimed(uint64 indexed id, address indexed claimant)
func (_Pandasia *PandasiaFilterer) WatchAirdropClaimed(opts *bind.WatchOpts, sink chan<- *PandasiaAirdropClaimed, id []uint64, claimant []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "AirdropClaimed", idRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaAirdropClaimed)
				if err := _Pandasia.contract.UnpackLog(event, "AirdropClaimed", log); err != nil {
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

// ParseAirdropClaimed is a log parse operation binding the contract event 0xf178a0585b458461acedd6bb1c991cda9d7223b0b94b5775d82a5c24166ebf99.
//
// Solidity: event AirdropClaimed(uint64 indexed id, address indexed claimant)
func (_Pandasia *PandasiaFilterer) ParseAirdropClaimed(log types.Log) (*PandasiaAirdropClaimed, error) {
	event := new(PandasiaAirdropClaimed)
	if err := _Pandasia.contract.UnpackLog(event, "AirdropClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PandasiaAirdropCreatedIterator is returned from FilterAirdropCreated and is used to iterate over the raw logs and unpacked data for AirdropCreated events raised by the Pandasia contract.
type PandasiaAirdropCreatedIterator struct {
	Event *PandasiaAirdropCreated // Event containing the contract specifics and raw log

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
func (it *PandasiaAirdropCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaAirdropCreated)
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
		it.Event = new(PandasiaAirdropCreated)
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
func (it *PandasiaAirdropCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaAirdropCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaAirdropCreated represents a AirdropCreated event raised by the Pandasia contract.
type PandasiaAirdropCreated struct {
	Id  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAirdropCreated is a free log retrieval operation binding the contract event 0x3aeb3bde18346c94bee94f4e93b995b047f057a0597450940a982121d645a8ab.
//
// Solidity: event AirdropCreated(uint64 indexed id)
func (_Pandasia *PandasiaFilterer) FilterAirdropCreated(opts *bind.FilterOpts, id []uint64) (*PandasiaAirdropCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "AirdropCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &PandasiaAirdropCreatedIterator{contract: _Pandasia.contract, event: "AirdropCreated", logs: logs, sub: sub}, nil
}

// WatchAirdropCreated is a free log subscription operation binding the contract event 0x3aeb3bde18346c94bee94f4e93b995b047f057a0597450940a982121d645a8ab.
//
// Solidity: event AirdropCreated(uint64 indexed id)
func (_Pandasia *PandasiaFilterer) WatchAirdropCreated(opts *bind.WatchOpts, sink chan<- *PandasiaAirdropCreated, id []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "AirdropCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaAirdropCreated)
				if err := _Pandasia.contract.UnpackLog(event, "AirdropCreated", log); err != nil {
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

// ParseAirdropCreated is a log parse operation binding the contract event 0x3aeb3bde18346c94bee94f4e93b995b047f057a0597450940a982121d645a8ab.
//
// Solidity: event AirdropCreated(uint64 indexed id)
func (_Pandasia *PandasiaFilterer) ParseAirdropCreated(log types.Log) (*PandasiaAirdropCreated, error) {
	event := new(PandasiaAirdropCreated)
	if err := _Pandasia.contract.UnpackLog(event, "AirdropCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PandasiaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pandasia contract.
type PandasiaInitializedIterator struct {
	Event *PandasiaInitialized // Event containing the contract specifics and raw log

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
func (it *PandasiaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaInitialized)
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
		it.Event = new(PandasiaInitialized)
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
func (it *PandasiaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaInitialized represents a Initialized event raised by the Pandasia contract.
type PandasiaInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pandasia *PandasiaFilterer) FilterInitialized(opts *bind.FilterOpts) (*PandasiaInitializedIterator, error) {

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PandasiaInitializedIterator{contract: _Pandasia.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pandasia *PandasiaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PandasiaInitialized) (event.Subscription, error) {

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaInitialized)
				if err := _Pandasia.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pandasia *PandasiaFilterer) ParseInitialized(log types.Log) (*PandasiaInitialized, error) {
	event := new(PandasiaInitialized)
	if err := _Pandasia.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PandasiaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pandasia contract.
type PandasiaOwnershipTransferredIterator struct {
	Event *PandasiaOwnershipTransferred // Event containing the contract specifics and raw log

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

// PandasiaRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pandasia contract.
type PandasiaRoleAdminChangedIterator struct {
	Event *PandasiaRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PandasiaRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaRoleAdminChanged)
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
		it.Event = new(PandasiaRoleAdminChanged)
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
func (it *PandasiaRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaRoleAdminChanged represents a RoleAdminChanged event raised by the Pandasia contract.
type PandasiaRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pandasia *PandasiaFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PandasiaRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PandasiaRoleAdminChangedIterator{contract: _Pandasia.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pandasia *PandasiaFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PandasiaRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaRoleAdminChanged)
				if err := _Pandasia.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pandasia *PandasiaFilterer) ParseRoleAdminChanged(log types.Log) (*PandasiaRoleAdminChanged, error) {
	event := new(PandasiaRoleAdminChanged)
	if err := _Pandasia.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PandasiaRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pandasia contract.
type PandasiaRoleGrantedIterator struct {
	Event *PandasiaRoleGranted // Event containing the contract specifics and raw log

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
func (it *PandasiaRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaRoleGranted)
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
		it.Event = new(PandasiaRoleGranted)
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
func (it *PandasiaRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaRoleGranted represents a RoleGranted event raised by the Pandasia contract.
type PandasiaRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pandasia *PandasiaFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PandasiaRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PandasiaRoleGrantedIterator{contract: _Pandasia.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pandasia *PandasiaFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PandasiaRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaRoleGranted)
				if err := _Pandasia.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pandasia *PandasiaFilterer) ParseRoleGranted(log types.Log) (*PandasiaRoleGranted, error) {
	event := new(PandasiaRoleGranted)
	if err := _Pandasia.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PandasiaRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pandasia contract.
type PandasiaRoleRevokedIterator struct {
	Event *PandasiaRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PandasiaRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PandasiaRoleRevoked)
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
		it.Event = new(PandasiaRoleRevoked)
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
func (it *PandasiaRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PandasiaRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PandasiaRoleRevoked represents a RoleRevoked event raised by the Pandasia contract.
type PandasiaRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pandasia *PandasiaFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PandasiaRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pandasia.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PandasiaRoleRevokedIterator{contract: _Pandasia.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pandasia *PandasiaFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PandasiaRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pandasia.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PandasiaRoleRevoked)
				if err := _Pandasia.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pandasia *PandasiaFilterer) ParseRoleRevoked(log types.Log) (*PandasiaRoleRevoked, error) {
	event := new(PandasiaRoleRevoked)
	if err := _Pandasia.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
