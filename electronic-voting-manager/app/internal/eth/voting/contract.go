// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_votingTitle\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_votingEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_votingOptionName\",\"type\":\"string\"}],\"name\":\"addVotingOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeVotingOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNameVotingOption\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberRegisteredVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNumberVotesVotingOption\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptionsCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingOptionsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voterAddress\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasRightToVote\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votedFor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votingOptions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000cb638038062000cb683398101604081905262000034916200008f565b804211156200004257600080fd5b6000620000508382620001f9565b50600060015560025550600380546001600160a81b0319163360ff60a01b1916179055620002c5565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215620000a357600080fd5b82516001600160401b0380821115620000bb57600080fd5b818501915085601f830112620000d057600080fd5b815181811115620000e557620000e562000079565b604051601f8201601f19908116603f0116810190838211818310171562000110576200011062000079565b816040528281526020935088848487010111156200012d57600080fd5b600091505b8282101562000151578482018401518183018501529083019062000132565b6000928101840192909252509401519395939450505050565b600181811c908216806200017f57607f821691505b602082108103620001a057634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001f457600081815260208120601f850160051c81016020861015620001cf5750805b601f850160051c820191505b81811015620001f057828155600101620001db565b5050505b505050565b81516001600160401b0381111562000215576200021562000079565b6200022d816200022684546200016a565b84620001a6565b602080601f8311600181146200026557600084156200024c5750858301515b600019600386901b1c1916600185901b178555620001f0565b600085815260208120601f198616915b82811015620002965788860151825594840194600190910190840162000275565b5085821015620002b55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6109e180620002d56000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806393d8bc051161008c578063b1bf87b011610066578063b1bf87b0146101be578063d03c85ec146101c6578063dcf7b628146101e3578063dff67f6b146101eb57600080fd5b806393d8bc05146101525780639e7b8d611461015a578063a3ec138d1461016d57600080fd5b80630121b93f146100d45780631cfe7e5a146100e95780631da3596d146100fc5780634ea221ac1461012557806368765e581461012d57806385c3f85d1461013f575b600080fd5b6100e76100e236600461067f565b61020c565b005b6100e76100f73660046106ae565b6102f9565b61010f61010a36600461067f565b6103a0565b60405161011c91906107a5565b60405180910390f35b6100e7610456565b6001545b60405190815260200161011c565b61013161014d36600461067f565b61049e565b61010f6104cc565b6100e76101683660046107bf565b61055e565b6101a161017b3660046107bf565b6005602052600090815260409020805460019091015460ff808316926101009004169083565b60408051931515845291151560208401529082015260600161011c565b600254610131565b600354600160a01b900460ff16604051901515815260200161011c565b600454610131565b6101fe6101f936600461067f565b6105c3565b60405161011c9291906107e8565b6002544211806102265750600354600160a01b900460ff16155b1561023057600080fd5b3360009081526005602052604081208054909160ff9091161515900361025557600080fd5b8054610100900460ff1615156001036102a757600160048260010154815481106102815761028161080a565b906000526020600020906002020160010160008282546102a19190610836565b90915550505b805461ff001916610100178155600181810183905560048054849081106102d0576102d061080a565b906000526020600020906002020160010160008282546102f0919061084f565b90915550505050565b6003546001600160a01b0316331461031057600080fd5b60025442118061032e5750600354600160a01b900460ff1615156001145b1561033857600080fd5b60408051808201909152818152600060208201819052600480546001810182559152815160029091027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0190819061039090826108eb565b5060208201518160010155505050565b6060600482815481106103b5576103b561080a565b906000526020600020906002020160000180546103d190610862565b80601f01602080910402602001604051908101604052809291908181526020018280546103fd90610862565b801561044a5780601f1061041f5761010080835404028352916020019161044a565b820191906000526020600020905b81548152906001019060200180831161042d57829003601f168201915b50505050509050919050565b6003546001600160a01b0316331461046d57600080fd5b60025442118061047f57506004546002115b1561048957600080fd5b6003805460ff60a01b1916600160a01b179055565b6000600482815481106104b3576104b361080a565b9060005260206000209060020201600101549050919050565b6060600080546104db90610862565b80601f016020809104026020016040519081016040528092919081815260200182805461050790610862565b80156105545780601f1061052957610100808354040283529160200191610554565b820191906000526020600020905b81548152906001019060200180831161053757829003601f168201915b5050505050905090565b6003546001600160a01b0316331461057557600080fd5b60025442111561058457600080fd5b6001600160a01b0381166000908152600560205260408120805460ff191660019081179091558054909182916105bb90839061084f565b909155505050565b600481815481106105d357600080fd5b90600052602060002090600202016000915090508060000180546105f690610862565b80601f016020809104026020016040519081016040528092919081815260200182805461062290610862565b801561066f5780601f106106445761010080835404028352916020019161066f565b820191906000526020600020905b81548152906001019060200180831161065257829003601f168201915b5050505050908060010154905082565b60006020828403121561069157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156106c057600080fd5b813567ffffffffffffffff808211156106d857600080fd5b818401915084601f8301126106ec57600080fd5b8135818111156106fe576106fe610698565b604051601f8201601f19908116603f0116810190838211818310171561072657610726610698565b8160405282815287602084870101111561073f57600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000815180845260005b8181101561078557602081850181015186830182015201610769565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006107b8602083018461075f565b9392505050565b6000602082840312156107d157600080fd5b81356001600160a01b03811681146107b857600080fd5b6040815260006107fb604083018561075f565b90508260208301529392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561084957610849610820565b92915050565b8082018082111561084957610849610820565b600181811c9082168061087657607f821691505b60208210810361089657634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156108e657600081815260208120601f850160051c810160208610156108c35750805b601f850160051c820191505b818110156108e2578281556001016108cf565b5050505b505050565b815167ffffffffffffffff81111561090557610905610698565b610919816109138454610862565b8461089c565b602080601f83116001811461094e57600084156109365750858301515b600019600386901b1c1916600185901b1785556108e2565b600085815260208120601f198616915b8281101561097d5788860151825594840194600190910190840161095e565b508582101561099b5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220aa15b195bcb4c92166f5c4fcc24eee073c9a7cc030c161858d6067e0602b731364736f6c63430008130033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _votingTitle string, _votingEndTime *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _votingTitle, _votingEndTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetNameVotingOption is a free data retrieval call binding the contract method 0x1da3596d.
//
// Solidity: function getNameVotingOption(uint256 idx) view returns(string)
func (_Contract *ContractCaller) GetNameVotingOption(opts *bind.CallOpts, idx *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNameVotingOption", idx)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNameVotingOption is a free data retrieval call binding the contract method 0x1da3596d.
//
// Solidity: function getNameVotingOption(uint256 idx) view returns(string)
func (_Contract *ContractSession) GetNameVotingOption(idx *big.Int) (string, error) {
	return _Contract.Contract.GetNameVotingOption(&_Contract.CallOpts, idx)
}

// GetNameVotingOption is a free data retrieval call binding the contract method 0x1da3596d.
//
// Solidity: function getNameVotingOption(uint256 idx) view returns(string)
func (_Contract *ContractCallerSession) GetNameVotingOption(idx *big.Int) (string, error) {
	return _Contract.Contract.GetNameVotingOption(&_Contract.CallOpts, idx)
}

// GetNumberRegisteredVoters is a free data retrieval call binding the contract method 0x68765e58.
//
// Solidity: function getNumberRegisteredVoters() view returns(uint256)
func (_Contract *ContractCaller) GetNumberRegisteredVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNumberRegisteredVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberRegisteredVoters is a free data retrieval call binding the contract method 0x68765e58.
//
// Solidity: function getNumberRegisteredVoters() view returns(uint256)
func (_Contract *ContractSession) GetNumberRegisteredVoters() (*big.Int, error) {
	return _Contract.Contract.GetNumberRegisteredVoters(&_Contract.CallOpts)
}

// GetNumberRegisteredVoters is a free data retrieval call binding the contract method 0x68765e58.
//
// Solidity: function getNumberRegisteredVoters() view returns(uint256)
func (_Contract *ContractCallerSession) GetNumberRegisteredVoters() (*big.Int, error) {
	return _Contract.Contract.GetNumberRegisteredVoters(&_Contract.CallOpts)
}

// GetNumberVotesVotingOption is a free data retrieval call binding the contract method 0x85c3f85d.
//
// Solidity: function getNumberVotesVotingOption(uint256 idx) view returns(uint256)
func (_Contract *ContractCaller) GetNumberVotesVotingOption(opts *bind.CallOpts, idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNumberVotesVotingOption", idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberVotesVotingOption is a free data retrieval call binding the contract method 0x85c3f85d.
//
// Solidity: function getNumberVotesVotingOption(uint256 idx) view returns(uint256)
func (_Contract *ContractSession) GetNumberVotesVotingOption(idx *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetNumberVotesVotingOption(&_Contract.CallOpts, idx)
}

// GetNumberVotesVotingOption is a free data retrieval call binding the contract method 0x85c3f85d.
//
// Solidity: function getNumberVotesVotingOption(uint256 idx) view returns(uint256)
func (_Contract *ContractCallerSession) GetNumberVotesVotingOption(idx *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetNumberVotesVotingOption(&_Contract.CallOpts, idx)
}

// GetOptionsCompleted is a free data retrieval call binding the contract method 0xd03c85ec.
//
// Solidity: function getOptionsCompleted() view returns(bool)
func (_Contract *ContractCaller) GetOptionsCompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOptionsCompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOptionsCompleted is a free data retrieval call binding the contract method 0xd03c85ec.
//
// Solidity: function getOptionsCompleted() view returns(bool)
func (_Contract *ContractSession) GetOptionsCompleted() (bool, error) {
	return _Contract.Contract.GetOptionsCompleted(&_Contract.CallOpts)
}

// GetOptionsCompleted is a free data retrieval call binding the contract method 0xd03c85ec.
//
// Solidity: function getOptionsCompleted() view returns(bool)
func (_Contract *ContractCallerSession) GetOptionsCompleted() (bool, error) {
	return _Contract.Contract.GetOptionsCompleted(&_Contract.CallOpts)
}

// GetVotingEndTime is a free data retrieval call binding the contract method 0xb1bf87b0.
//
// Solidity: function getVotingEndTime() view returns(uint256)
func (_Contract *ContractCaller) GetVotingEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVotingEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingEndTime is a free data retrieval call binding the contract method 0xb1bf87b0.
//
// Solidity: function getVotingEndTime() view returns(uint256)
func (_Contract *ContractSession) GetVotingEndTime() (*big.Int, error) {
	return _Contract.Contract.GetVotingEndTime(&_Contract.CallOpts)
}

// GetVotingEndTime is a free data retrieval call binding the contract method 0xb1bf87b0.
//
// Solidity: function getVotingEndTime() view returns(uint256)
func (_Contract *ContractCallerSession) GetVotingEndTime() (*big.Int, error) {
	return _Contract.Contract.GetVotingEndTime(&_Contract.CallOpts)
}

// GetVotingOptionsLength is a free data retrieval call binding the contract method 0xdcf7b628.
//
// Solidity: function getVotingOptionsLength() view returns(uint256)
func (_Contract *ContractCaller) GetVotingOptionsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVotingOptionsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingOptionsLength is a free data retrieval call binding the contract method 0xdcf7b628.
//
// Solidity: function getVotingOptionsLength() view returns(uint256)
func (_Contract *ContractSession) GetVotingOptionsLength() (*big.Int, error) {
	return _Contract.Contract.GetVotingOptionsLength(&_Contract.CallOpts)
}

// GetVotingOptionsLength is a free data retrieval call binding the contract method 0xdcf7b628.
//
// Solidity: function getVotingOptionsLength() view returns(uint256)
func (_Contract *ContractCallerSession) GetVotingOptionsLength() (*big.Int, error) {
	return _Contract.Contract.GetVotingOptionsLength(&_Contract.CallOpts)
}

// GetVotingTitle is a free data retrieval call binding the contract method 0x93d8bc05.
//
// Solidity: function getVotingTitle() view returns(string)
func (_Contract *ContractCaller) GetVotingTitle(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVotingTitle")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVotingTitle is a free data retrieval call binding the contract method 0x93d8bc05.
//
// Solidity: function getVotingTitle() view returns(string)
func (_Contract *ContractSession) GetVotingTitle() (string, error) {
	return _Contract.Contract.GetVotingTitle(&_Contract.CallOpts)
}

// GetVotingTitle is a free data retrieval call binding the contract method 0x93d8bc05.
//
// Solidity: function getVotingTitle() view returns(string)
func (_Contract *ContractCallerSession) GetVotingTitle() (string, error) {
	return _Contract.Contract.GetVotingTitle(&_Contract.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool hasRightToVote, bool voted, uint256 votedFor)
func (_Contract *ContractCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	HasRightToVote bool
	Voted          bool
	VotedFor       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		HasRightToVote bool
		Voted          bool
		VotedFor       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HasRightToVote = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.VotedFor = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool hasRightToVote, bool voted, uint256 votedFor)
func (_Contract *ContractSession) Voters(arg0 common.Address) (struct {
	HasRightToVote bool
	Voted          bool
	VotedFor       *big.Int
}, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool hasRightToVote, bool voted, uint256 votedFor)
func (_Contract *ContractCallerSession) Voters(arg0 common.Address) (struct {
	HasRightToVote bool
	Voted          bool
	VotedFor       *big.Int
}, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// VotingOptions is a free data retrieval call binding the contract method 0xdff67f6b.
//
// Solidity: function votingOptions(uint256 ) view returns(string name, uint256 numberVotes)
func (_Contract *ContractCaller) VotingOptions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        string
	NumberVotes *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "votingOptions", arg0)

	outstruct := new(struct {
		Name        string
		NumberVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.NumberVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VotingOptions is a free data retrieval call binding the contract method 0xdff67f6b.
//
// Solidity: function votingOptions(uint256 ) view returns(string name, uint256 numberVotes)
func (_Contract *ContractSession) VotingOptions(arg0 *big.Int) (struct {
	Name        string
	NumberVotes *big.Int
}, error) {
	return _Contract.Contract.VotingOptions(&_Contract.CallOpts, arg0)
}

// VotingOptions is a free data retrieval call binding the contract method 0xdff67f6b.
//
// Solidity: function votingOptions(uint256 ) view returns(string name, uint256 numberVotes)
func (_Contract *ContractCallerSession) VotingOptions(arg0 *big.Int) (struct {
	Name        string
	NumberVotes *big.Int
}, error) {
	return _Contract.Contract.VotingOptions(&_Contract.CallOpts, arg0)
}

// AddVotingOption is a paid mutator transaction binding the contract method 0x1cfe7e5a.
//
// Solidity: function addVotingOption(string _votingOptionName) returns()
func (_Contract *ContractTransactor) AddVotingOption(opts *bind.TransactOpts, _votingOptionName string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addVotingOption", _votingOptionName)
}

// AddVotingOption is a paid mutator transaction binding the contract method 0x1cfe7e5a.
//
// Solidity: function addVotingOption(string _votingOptionName) returns()
func (_Contract *ContractSession) AddVotingOption(_votingOptionName string) (*types.Transaction, error) {
	return _Contract.Contract.AddVotingOption(&_Contract.TransactOpts, _votingOptionName)
}

// AddVotingOption is a paid mutator transaction binding the contract method 0x1cfe7e5a.
//
// Solidity: function addVotingOption(string _votingOptionName) returns()
func (_Contract *ContractTransactorSession) AddVotingOption(_votingOptionName string) (*types.Transaction, error) {
	return _Contract.Contract.AddVotingOption(&_Contract.TransactOpts, _votingOptionName)
}

// CompleteVotingOptions is a paid mutator transaction binding the contract method 0x4ea221ac.
//
// Solidity: function completeVotingOptions() returns()
func (_Contract *ContractTransactor) CompleteVotingOptions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "completeVotingOptions")
}

// CompleteVotingOptions is a paid mutator transaction binding the contract method 0x4ea221ac.
//
// Solidity: function completeVotingOptions() returns()
func (_Contract *ContractSession) CompleteVotingOptions() (*types.Transaction, error) {
	return _Contract.Contract.CompleteVotingOptions(&_Contract.TransactOpts)
}

// CompleteVotingOptions is a paid mutator transaction binding the contract method 0x4ea221ac.
//
// Solidity: function completeVotingOptions() returns()
func (_Contract *ContractTransactorSession) CompleteVotingOptions() (*types.Transaction, error) {
	return _Contract.Contract.CompleteVotingOptions(&_Contract.TransactOpts)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address _voterAddress) returns()
func (_Contract *ContractTransactor) GiveRightToVote(opts *bind.TransactOpts, _voterAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "giveRightToVote", _voterAddress)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address _voterAddress) returns()
func (_Contract *ContractSession) GiveRightToVote(_voterAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GiveRightToVote(&_Contract.TransactOpts, _voterAddress)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address _voterAddress) returns()
func (_Contract *ContractTransactorSession) GiveRightToVote(_voterAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GiveRightToVote(&_Contract.TransactOpts, _voterAddress)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 idx) returns()
func (_Contract *ContractTransactor) Vote(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "vote", idx)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 idx) returns()
func (_Contract *ContractSession) Vote(idx *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, idx)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 idx) returns()
func (_Contract *ContractTransactorSession) Vote(idx *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, idx)
}
