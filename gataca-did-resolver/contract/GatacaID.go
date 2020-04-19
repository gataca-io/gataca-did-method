// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// GatacaIDABI is the input ABI used to generate the binding from.
const GatacaIDABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"getEntity\",\"outputs\":[{\"name\":\"status\",\"type\":\"string\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"suspendEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"activateEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_did\",\"type\":\"string\"},{\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"createEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"revokeEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// GatacaIDBin is the compiled bytecode used for deploying new contracts.
var GatacaIDBin = "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061180b806100616000396000f3fe608060405234801561001057600080fd5b5060043610610074576000357c010000000000000000000000000000000000000000000000000000000090048063070c412b14610079578063698e3683146101d7578063741cecc71461029257806397c5f12f1461034d578063b46e43931461049f575b600080fd5b6100f06004803603602081101561008f57600080fd5b81019080803590602001906401000000008111156100ac57600080fd5b8201836020820111156100be57600080fd5b803590602001918460018302840111640100000000831117156100e057600080fd5b909192939192939050505061055a565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610134578082015181840152602081019050610119565b50505050905090810190601f1680156101615780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561019a57808201518184015260208101905061017f565b50505050905090810190601f1680156101c75780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610290600480360360208110156101ed57600080fd5b810190808035906020019064010000000081111561020a57600080fd5b82018360208201111561021c57600080fd5b8035906020019184600183028401116401000000008311171561023e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506109e1565b005b61034b600480360360208110156102a857600080fd5b81019080803590602001906401000000008111156102c557600080fd5b8201836020820111156102d757600080fd5b803590602001918460018302840111640100000000831117156102f957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d31565b005b61049d6004803603604081101561036357600080fd5b810190808035906020019064010000000081111561038057600080fd5b82018360208201111561039257600080fd5b803590602001918460018302840111640100000000831117156103b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561041757600080fd5b82018360208201111561042957600080fd5b8035906020019184600183028401116401000000008311171561044b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611081565b005b610558600480360360208110156104b557600080fd5b81019080803590602001906401000000008111156104d257600080fd5b8201836020820111156104e457600080fd5b8035906020019184600183028401116401000000008311171561050657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061148a565b005b6060806105656116a0565b600085856040518083838082843780830192505050925050509081526020016040518091039020604080519081016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106305780601f1061060557610100808354040283529160200191610630565b820191906000526020600020905b81548152906001019060200180831161061357829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106d25780601f106106a7576101008083540402835291602001916106d2565b820191906000526020600020905b8154815290600101906020018083116106b557829003601f168201915b505050505081525050905060008160000151511415151561075b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f696e76616c69642064696400000000000000000000000000000000000000000081525060200191505060405180910390fd5b60405160200180807f7265766f6b65640000000000000000000000000000000000000000000000000081525060070190506040516020818303038152906040528051906020012081600001516040516020018082805190602001908083835b6020831015156107df57805182526020820191506020810190506020830392506107ba565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014151515610890576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f646964207265766f6b656400000000000000000000000000000000000000000081525060200191505060405180910390fd5b60405160200180807f73757370656e646564000000000000000000000000000000000000000000000081525060090190506040516020818303038152906040528051906020012081600001516040516020018082805190602001908083835b60208310151561091457805182526020820191506020810190506020830392506108ef565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120141515156109c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f6469642073757370656e6465640000000000000000000000000000000000000081525060200191505060405180910390fd5b8060000151816020015181915080905092509250509250929050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610aa6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642075736572000000000000000000000000000000000000000081525060200191505060405180910390fd5b600080826040518082805190602001908083835b602083101515610adf5780518252602082019150602081019050602083039250610aba565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020905060008160000180546001816001161561010002031660029004905014151515610ba5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f696e76616c69642064696400000000000000000000000000000000000000000081525060200191505060405180910390fd5b60405160200180807f7265766f6b656400000000000000000000000000000000000000000000000000815250600701905060405160208183030381529060405280519060200120816000016040516020018082805460018160011615610100020316600290048015610c4e5780601f10610c2c576101008083540402835291820191610c4e565b820191906000526020600020905b815481529060010190602001808311610c3a575b50509150506040516020818303038152906040528051906020012014151515610cdf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f646964207265766f6b656400000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040805190810160405280600981526020017f73757370656e6465640000000000000000000000000000000000000000000000815250816000019080519060200190610d2c9291906116ba565b505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610df6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642075736572000000000000000000000000000000000000000081525060200191505060405180910390fd5b600080826040518082805190602001908083835b602083101515610e2f5780518252602082019150602081019050602083039250610e0a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020905060008160000180546001816001161561010002031660029004905014151515610ef5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f696e76616c69642064696400000000000000000000000000000000000000000081525060200191505060405180910390fd5b60405160200180807f7265766f6b656400000000000000000000000000000000000000000000000000815250600701905060405160208183030381529060405280519060200120816000016040516020018082805460018160011615610100020316600290048015610f9e5780601f10610f7c576101008083540402835291820191610f9e565b820191906000526020600020905b815481529060010190602001808311610f8a575b5050915050604051602081830303815290604052805190602001201415151561102f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f646964207265766f6b656400000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040805190810160405280600681526020017f616374697665000000000000000000000000000000000000000000000000000081525081600001908051906020019061107c9291906116ba565b505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611146576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642075736572000000000000000000000000000000000000000081525060200191505060405180910390fd5b61114e6116a0565b6000836040518082805190602001908083835b6020831015156111865780518252602082019150602081019050602083039250611161565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020604080519081016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561125d5780601f106112325761010080835404028352916020019161125d565b820191906000526020600020905b81548152906001019060200180831161124057829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112ff5780601f106112d4576101008083540402835291602001916112ff565b820191906000526020600020905b8154815290600101906020018083116112e257829003601f168201915b50505050508152505090506000816000015151141515611387576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f64696420616c726561647920696e20757365000000000000000000000000000081525060200191505060405180910390fd5b61138f6116a0565b60408051908101604052806040805190810160405280600681526020017f61637469766500000000000000000000000000000000000000000000000000008152508152602001848152509050806000856040518082805190602001908083835b60208310151561141457805182526020820191506020810190506020830392506113ef565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082015181600001908051906020019061146392919061173a565b50602082015181600101908051906020019061148092919061173a565b5090505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561154f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642075736572000000000000000000000000000000000000000081525060200191505060405180910390fd5b600080826040518082805190602001908083835b6020831015156115885780518252602082019150602081019050602083039250611563565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090506000816000018054600181600116156101000203166002900490501415151561164e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f696e76616c69642064696400000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040805190810160405280600781526020017f7265766f6b65640000000000000000000000000000000000000000000000000081525081600001908051906020019061169b9291906116ba565b505050565b604080519081016040528060608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116fb57805160ff1916838001178555611729565b82800160010185558215611729579182015b8281111561172857825182559160200191906001019061170d565b5b50905061173691906117ba565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061177b57805160ff19168380011785556117a9565b828001600101855582156117a9579182015b828111156117a857825182559160200191906001019061178d565b5b5090506117b691906117ba565b5090565b6117dc91905b808211156117d85760008160009055506001016117c0565b5090565b9056fea165627a7a723058207d774860eb49ee2fc2e002bfaeacf581448712772d18dcd49d972c137040e8ba0029"

// DeployGatacaID deploys a new Ethereum contract, binding an instance of GatacaID to it.
func DeployGatacaID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatacaID, error) {
	parsed, err := abi.JSON(strings.NewReader(GatacaIDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GatacaIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatacaID{GatacaIDCaller: GatacaIDCaller{contract: contract}, GatacaIDTransactor: GatacaIDTransactor{contract: contract}, GatacaIDFilterer: GatacaIDFilterer{contract: contract}}, nil
}

// GatacaID is an auto generated Go binding around an Ethereum contract.
type GatacaID struct {
	GatacaIDCaller     // Read-only binding to the contract
	GatacaIDTransactor // Write-only binding to the contract
	GatacaIDFilterer   // Log filterer for contract events
}

// GatacaIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatacaIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatacaIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatacaIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatacaIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatacaIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatacaIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatacaIDSession struct {
	Contract     *GatacaID         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatacaIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatacaIDCallerSession struct {
	Contract *GatacaIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GatacaIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatacaIDTransactorSession struct {
	Contract     *GatacaIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GatacaIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatacaIDRaw struct {
	Contract *GatacaID // Generic contract binding to access the raw methods on
}

// GatacaIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatacaIDCallerRaw struct {
	Contract *GatacaIDCaller // Generic read-only contract binding to access the raw methods on
}

// GatacaIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatacaIDTransactorRaw struct {
	Contract *GatacaIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatacaID creates a new instance of GatacaID, bound to a specific deployed contract.
func NewGatacaID(address common.Address, backend bind.ContractBackend) (*GatacaID, error) {
	contract, err := bindGatacaID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatacaID{GatacaIDCaller: GatacaIDCaller{contract: contract}, GatacaIDTransactor: GatacaIDTransactor{contract: contract}, GatacaIDFilterer: GatacaIDFilterer{contract: contract}}, nil
}

// NewGatacaIDCaller creates a new read-only instance of GatacaID, bound to a specific deployed contract.
func NewGatacaIDCaller(address common.Address, caller bind.ContractCaller) (*GatacaIDCaller, error) {
	contract, err := bindGatacaID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatacaIDCaller{contract: contract}, nil
}

// NewGatacaIDTransactor creates a new write-only instance of GatacaID, bound to a specific deployed contract.
func NewGatacaIDTransactor(address common.Address, transactor bind.ContractTransactor) (*GatacaIDTransactor, error) {
	contract, err := bindGatacaID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatacaIDTransactor{contract: contract}, nil
}

// NewGatacaIDFilterer creates a new log filterer instance of GatacaID, bound to a specific deployed contract.
func NewGatacaIDFilterer(address common.Address, filterer bind.ContractFilterer) (*GatacaIDFilterer, error) {
	contract, err := bindGatacaID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatacaIDFilterer{contract: contract}, nil
}

// bindGatacaID binds a generic wrapper to an already deployed contract.
func bindGatacaID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatacaIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatacaID *GatacaIDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GatacaID.Contract.GatacaIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatacaID *GatacaIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatacaID.Contract.GatacaIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatacaID *GatacaIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatacaID.Contract.GatacaIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatacaID *GatacaIDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GatacaID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatacaID *GatacaIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatacaID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatacaID *GatacaIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatacaID.Contract.contract.Transact(opts, method, params...)
}

// GetEntity is a free data retrieval call binding the contract method 0x070c412b.
//
// Solidity: function getEntity(string _did) constant returns(string status, string pubKey)
func (_GatacaID *GatacaIDCaller) GetEntity(opts *bind.CallOpts, _did string) (struct {
	Status string
	PubKey string
}, error) {
	ret := new(struct {
		Status string
		PubKey string
	})
	out := ret
	err := _GatacaID.contract.Call(opts, out, "getEntity", _did)
	return *ret, err
}

// GetEntity is a free data retrieval call binding the contract method 0x070c412b.
//
// Solidity: function getEntity(string _did) constant returns(string status, string pubKey)
func (_GatacaID *GatacaIDSession) GetEntity(_did string) (struct {
	Status string
	PubKey string
}, error) {
	return _GatacaID.Contract.GetEntity(&_GatacaID.CallOpts, _did)
}

// GetEntity is a free data retrieval call binding the contract method 0x070c412b.
//
// Solidity: function getEntity(string _did) constant returns(string status, string pubKey)
func (_GatacaID *GatacaIDCallerSession) GetEntity(_did string) (struct {
	Status string
	PubKey string
}, error) {
	return _GatacaID.Contract.GetEntity(&_GatacaID.CallOpts, _did)
}

// ActivateEntity is a paid mutator transaction binding the contract method 0x741cecc7.
//
// Solidity: function activateEntity(string _did) returns()
func (_GatacaID *GatacaIDTransactor) ActivateEntity(opts *bind.TransactOpts, _did string) (*types.Transaction, error) {
	return _GatacaID.contract.Transact(opts, "activateEntity", _did)
}

// ActivateEntity is a paid mutator transaction binding the contract method 0x741cecc7.
//
// Solidity: function activateEntity(string _did) returns()
func (_GatacaID *GatacaIDSession) ActivateEntity(_did string) (*types.Transaction, error) {
	return _GatacaID.Contract.ActivateEntity(&_GatacaID.TransactOpts, _did)
}

// ActivateEntity is a paid mutator transaction binding the contract method 0x741cecc7.
//
// Solidity: function activateEntity(string _did) returns()
func (_GatacaID *GatacaIDTransactorSession) ActivateEntity(_did string) (*types.Transaction, error) {
	return _GatacaID.Contract.ActivateEntity(&_GatacaID.TransactOpts, _did)
}

// CreateEntity is a paid mutator transaction binding the contract method 0x97c5f12f.
//
// Solidity: function createEntity(string _did, string _pubkey) returns()
func (_GatacaID *GatacaIDTransactor) CreateEntity(opts *bind.TransactOpts, _did string, _pubkey string) (*types.Transaction, error) {
	return _GatacaID.contract.Transact(opts, "createEntity", _did, _pubkey)
}

// CreateEntity is a paid mutator transaction binding the contract method 0x97c5f12f.
//
// Solidity: function createEntity(string _did, string _pubkey) returns()
func (_GatacaID *GatacaIDSession) CreateEntity(_did string, _pubkey string) (*types.Transaction, error) {
	return _GatacaID.Contract.CreateEntity(&_GatacaID.TransactOpts, _did, _pubkey)
}

// CreateEntity is a paid mutator transaction binding the contract method 0x97c5f12f.
//
// Solidity: function createEntity(string _did, string _pubkey) returns()
func (_GatacaID *GatacaIDTransactorSession) CreateEntity(_did string, _pubkey string) (*types.Transaction, error) {
	return _GatacaID.Contract.CreateEntity(&_GatacaID.TransactOpts, _did, _pubkey)
}

// RevokeEntity is a paid mutator transaction binding the contract method 0xb46e4393.
//
// Solidity: function revokeEntity(string _did) returns()
func (_GatacaID *GatacaIDTransactor) RevokeEntity(opts *bind.TransactOpts, _did string) (*types.Transaction, error) {
	return _GatacaID.contract.Transact(opts, "revokeEntity", _did)
}

// RevokeEntity is a paid mutator transaction binding the contract method 0xb46e4393.
//
// Solidity: function revokeEntity(string _did) returns()
func (_GatacaID *GatacaIDSession) RevokeEntity(_did string) (*types.Transaction, error) {
	return _GatacaID.Contract.RevokeEntity(&_GatacaID.TransactOpts, _did)
}

// RevokeEntity is a paid mutator transaction binding the contract method 0xb46e4393.
//
// Solidity: function revokeEntity(string _did) returns()
func (_GatacaID *GatacaIDTransactorSession) RevokeEntity(_did string) (*types.Transaction, error) {
	return _GatacaID.Contract.RevokeEntity(&_GatacaID.TransactOpts, _did)
}

// SuspendEntity is a paid mutator transaction binding the contract method 0x698e3683.
//
// Solidity: function suspendEntity(string _did) returns()
func (_GatacaID *GatacaIDTransactor) SuspendEntity(opts *bind.TransactOpts, _did string) (*types.Transaction, error) {
	return _GatacaID.contract.Transact(opts, "suspendEntity", _did)
}

// SuspendEntity is a paid mutator transaction binding the contract method 0x698e3683.
//
// Solidity: function suspendEntity(string _did) returns()
func (_GatacaID *GatacaIDSession) SuspendEntity(_did string) (*types.Transaction, error) {
	return _GatacaID.Contract.SuspendEntity(&_GatacaID.TransactOpts, _did)
}

// SuspendEntity is a paid mutator transaction binding the contract method 0x698e3683.
//
// Solidity: function suspendEntity(string _did) returns()
func (_GatacaID *GatacaIDTransactorSession) SuspendEntity(_did string) (*types.Transaction, error) {
	return _GatacaID.Contract.SuspendEntity(&_GatacaID.TransactOpts, _did)
}
