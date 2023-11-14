// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DataAvailabilityChallengeMetaData contains all meta data concerning the DataAvailabilityChallenge contract.
var DataAvailabilityChallengeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"BondTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengeExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengeNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengeWindowNotOpen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidInputData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"challengedBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumChallengeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ChallengeStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeWindow\",\"type\":\"uint256\"}],\"name\":\"RequiredBondSizeChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengedHash\",\"type\":\"bytes32\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolvedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengedHash\",\"type\":\"bytes32\"}],\"name\":\"getChallengeStatus\",\"outputs\":[{\"internalType\":\"enumChallengeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_challengeWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_resolveWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bondSize\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"preImage\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolveWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bondSize\",\"type\":\"uint256\"}],\"name\":\"setBondSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengedHash\",\"type\":\"bytes32\"}],\"name\":\"unlockBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611173806100206000396000f3fe6080604052600436106101125760003560e01c8063715018a6116100a5578063c459f80911610074578063d0e30db011610059578063d0e30db014610398578063d7d04e54146103a0578063f2fde38b146103c057600080fd5b8063c459f809146102dd578063c4ee20d4146102fd57600080fd5b8063715018a614610250578063861a1412146102655780638da5cb5b1461027b578063b740a2db146102b057600080fd5b80633ccfd60b116100e15780633ccfd60b146101af5780634ec81af1146101c457806354fd4d50146101e45780637099c5811461023a57600080fd5b806302b2f7c7146101265780630b1a73f41461013957806321cf39ee1461015957806327e235e31461018257600080fd5b366101215761011f6103e0565b005b600080fd5b61011f610134366004610ed6565b61044e565b34801561014557600080fd5b5061011f610154366004610ef8565b610656565b34801561016557600080fd5b5061016f60665481565b6040519081526020015b60405180910390f35b34801561018e57600080fd5b5061016f61019d366004610fa1565b60686020526000908152604090205481565b3480156101bb57600080fd5b5061011f61077c565b3480156101d057600080fd5b5061011f6101df366004610fc3565b6107da565b3480156101f057600080fd5b5061022d6040518060400160405280600581526020017f302e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101799190610ffc565b34801561024657600080fd5b5061016f60675481565b34801561025c57600080fd5b5061011f61098c565b34801561027157600080fd5b5061016f60655481565b34801561028757600080fd5b5060335460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610179565b3480156102bc57600080fd5b506102d06102cb366004610ed6565b6109a0565b604051610179919061109e565b3480156102e957600080fd5b5061011f6102f8366004610ed6565b610a48565b34801561030957600080fd5b50610361610318366004610ed6565b6069602090815260009283526040808420909152908252902080546001820154600283015460039093015473ffffffffffffffffffffffffffffffffffffffff90921692909184565b6040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152608001610179565b61011f6103e0565b3480156103ac57600080fd5b5061011f6103bb3660046110df565b610b5d565b3480156103cc57600080fd5b5061011f6103db366004610fa1565b610ba0565b33600090815260686020526040812080543492906103ff908490611127565b909155505033600081815260686020908152604091829020548251938452908301527fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a1565b6104566103e0565b6067543360009081526068602052604090205410156104c75733600090815260686020526040908190205460675491517e0155b50000000000000000000000000000000000000000000000000000000081526104be9290600401918252602082015260400190565b60405180910390fd5b60006104d383836109a0565b60038111156104e4576104e461106f565b1461051b576040517f9bb6c64e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61052482610c57565b61055a576040517ff9e0d1f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606754336000908152606860205260408120805490919061057c90849061113f565b90915550506040805160808101825233815260675460208083019182524383850190815260006060850181815288825260698452868220888352909352859020935184547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178455915160018085019190915591516002840155516003909201919091559051839183917f73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f89161064a9161109e565b60405180910390a35050565b8181604051610666929190611156565b604051809103902083146106c4578181604051610684929190611156565b6040519081900381207f3b7d73720000000000000000000000000000000000000000000000000000000082526004820152602481018490526044016104be565b60016106d085856109a0565b60038111156106e1576106e161106f565b14610718576040517fbeb11d3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848152606960209081526040808320868452909152908190204360038201559051859085907f73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f89061076d9060029061109e565b60405180910390a35050505050565b3360008181526068602052604081208054908290559161079d905a84610c79565b9050806107d6576040517f27fcd9d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b600054610100900460ff16158080156107fa5750600054600160ff909116105b806108145750303b158015610814575060005460ff166001145b6108a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104be565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108fe57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610906610c8f565b6065849055606683905561091982610b5d565b61092285610d2e565b801561098557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b610994610da5565b61099e6000610d2e565b565b600082815260696020908152604080832084845282528083208151608081018352815473ffffffffffffffffffffffffffffffffffffffff1680825260018301549482019490945260028201549281019290925260030154606082015290610a0c576000915050610a42565b606081015115610a20576002915050610a42565b610a2d8160400151610e26565b15610a3c576001915050610a42565b60039150505b92915050565b6003610a5483836109a0565b6003811115610a6557610a6561106f565b14610a9c576040517f151f07fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260696020908152604080832084845282528083206001810154815473ffffffffffffffffffffffffffffffffffffffff1685526068909352908320805491939091610aed908490611127565b9091555050600060018201819055815473ffffffffffffffffffffffffffffffffffffffff1680825260686020908152604092839020548351928352908201527fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a1505050565b610b65610da5565b60678190556040518181527f4468d695a0389e5f9e8ef0c9aee6d84e74cc0d0e0a28c8413badb54697d1bbae9060200160405180910390a150565b610ba8610da5565b73ffffffffffffffffffffffffffffffffffffffff8116610c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104be565b610c5481610d2e565b50565b60008143118015610a425750606554610c709083611127565b43111592915050565b600080600080600080868989f195945050505050565b600054610100900460ff16610d26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104be565b61099e610e36565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461099e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104be565b600060665482610c709190611127565b600054610100900460ff16610ecd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104be565b61099e33610d2e565b60008060408385031215610ee957600080fd5b50508035926020909101359150565b60008060008060608587031215610f0e57600080fd5b8435935060208501359250604085013567ffffffffffffffff80821115610f3457600080fd5b818701915087601f830112610f4857600080fd5b813581811115610f5757600080fd5b886020828501011115610f6957600080fd5b95989497505060200194505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f9c57600080fd5b919050565b600060208284031215610fb357600080fd5b610fbc82610f78565b9392505050565b60008060008060808587031215610fd957600080fd5b610fe285610f78565b966020860135965060408601359560600135945092505050565b600060208083528351808285015260005b818110156110295785810183015185820160400152820161100d565b8181111561103b576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600483106110d9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000602082840312156110f157600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561113a5761113a6110f8565b500190565b600082821015611151576111516110f8565b500390565b818382376000910190815291905056fea164736f6c634300080f000a",
}

// DataAvailabilityChallengeABI is the input ABI used to generate the binding from.
// Deprecated: Use DataAvailabilityChallengeMetaData.ABI instead.
var DataAvailabilityChallengeABI = DataAvailabilityChallengeMetaData.ABI

// DataAvailabilityChallengeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataAvailabilityChallengeMetaData.Bin instead.
var DataAvailabilityChallengeBin = DataAvailabilityChallengeMetaData.Bin

// DeployDataAvailabilityChallenge deploys a new Ethereum contract, binding an instance of DataAvailabilityChallenge to it.
func DeployDataAvailabilityChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataAvailabilityChallenge, error) {
	parsed, err := DataAvailabilityChallengeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataAvailabilityChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataAvailabilityChallenge{DataAvailabilityChallengeCaller: DataAvailabilityChallengeCaller{contract: contract}, DataAvailabilityChallengeTransactor: DataAvailabilityChallengeTransactor{contract: contract}, DataAvailabilityChallengeFilterer: DataAvailabilityChallengeFilterer{contract: contract}}, nil
}

// DataAvailabilityChallenge is an auto generated Go binding around an Ethereum contract.
type DataAvailabilityChallenge struct {
	DataAvailabilityChallengeCaller     // Read-only binding to the contract
	DataAvailabilityChallengeTransactor // Write-only binding to the contract
	DataAvailabilityChallengeFilterer   // Log filterer for contract events
}

// DataAvailabilityChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataAvailabilityChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataAvailabilityChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataAvailabilityChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataAvailabilityChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataAvailabilityChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataAvailabilityChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataAvailabilityChallengeSession struct {
	Contract     *DataAvailabilityChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DataAvailabilityChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataAvailabilityChallengeCallerSession struct {
	Contract *DataAvailabilityChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// DataAvailabilityChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataAvailabilityChallengeTransactorSession struct {
	Contract     *DataAvailabilityChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// DataAvailabilityChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataAvailabilityChallengeRaw struct {
	Contract *DataAvailabilityChallenge // Generic contract binding to access the raw methods on
}

// DataAvailabilityChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataAvailabilityChallengeCallerRaw struct {
	Contract *DataAvailabilityChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// DataAvailabilityChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataAvailabilityChallengeTransactorRaw struct {
	Contract *DataAvailabilityChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataAvailabilityChallenge creates a new instance of DataAvailabilityChallenge, bound to a specific deployed contract.
func NewDataAvailabilityChallenge(address common.Address, backend bind.ContractBackend) (*DataAvailabilityChallenge, error) {
	contract, err := bindDataAvailabilityChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallenge{DataAvailabilityChallengeCaller: DataAvailabilityChallengeCaller{contract: contract}, DataAvailabilityChallengeTransactor: DataAvailabilityChallengeTransactor{contract: contract}, DataAvailabilityChallengeFilterer: DataAvailabilityChallengeFilterer{contract: contract}}, nil
}

// NewDataAvailabilityChallengeCaller creates a new read-only instance of DataAvailabilityChallenge, bound to a specific deployed contract.
func NewDataAvailabilityChallengeCaller(address common.Address, caller bind.ContractCaller) (*DataAvailabilityChallengeCaller, error) {
	contract, err := bindDataAvailabilityChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeCaller{contract: contract}, nil
}

// NewDataAvailabilityChallengeTransactor creates a new write-only instance of DataAvailabilityChallenge, bound to a specific deployed contract.
func NewDataAvailabilityChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*DataAvailabilityChallengeTransactor, error) {
	contract, err := bindDataAvailabilityChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeTransactor{contract: contract}, nil
}

// NewDataAvailabilityChallengeFilterer creates a new log filterer instance of DataAvailabilityChallenge, bound to a specific deployed contract.
func NewDataAvailabilityChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*DataAvailabilityChallengeFilterer, error) {
	contract, err := bindDataAvailabilityChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeFilterer{contract: contract}, nil
}

// bindDataAvailabilityChallenge binds a generic wrapper to an already deployed contract.
func bindDataAvailabilityChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataAvailabilityChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataAvailabilityChallenge *DataAvailabilityChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataAvailabilityChallenge.Contract.DataAvailabilityChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataAvailabilityChallenge *DataAvailabilityChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.DataAvailabilityChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataAvailabilityChallenge *DataAvailabilityChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.DataAvailabilityChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataAvailabilityChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.Balances(&_DataAvailabilityChallenge.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.Balances(&_DataAvailabilityChallenge.CallOpts, arg0)
}

// BondSize is a free data retrieval call binding the contract method 0x7099c581.
//
// Solidity: function bondSize() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) BondSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "bondSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondSize is a free data retrieval call binding the contract method 0x7099c581.
//
// Solidity: function bondSize() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) BondSize() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.BondSize(&_DataAvailabilityChallenge.CallOpts)
}

// BondSize is a free data retrieval call binding the contract method 0x7099c581.
//
// Solidity: function bondSize() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) BondSize() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.BondSize(&_DataAvailabilityChallenge.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) ChallengeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "challengeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) ChallengeWindow() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.ChallengeWindow(&_DataAvailabilityChallenge.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) ChallengeWindow() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.ChallengeWindow(&_DataAvailabilityChallenge.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0xc4ee20d4.
//
// Solidity: function challenges(uint256 , bytes32 ) view returns(address challenger, uint256 lockedBond, uint256 startBlock, uint256 resolvedBlock)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Challenger    common.Address
	LockedBond    *big.Int
	StartBlock    *big.Int
	ResolvedBlock *big.Int
}, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "challenges", arg0, arg1)

	outstruct := new(struct {
		Challenger    common.Address
		LockedBond    *big.Int
		StartBlock    *big.Int
		ResolvedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenger = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LockedBond = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ResolvedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0xc4ee20d4.
//
// Solidity: function challenges(uint256 , bytes32 ) view returns(address challenger, uint256 lockedBond, uint256 startBlock, uint256 resolvedBlock)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Challenges(arg0 *big.Int, arg1 [32]byte) (struct {
	Challenger    common.Address
	LockedBond    *big.Int
	StartBlock    *big.Int
	ResolvedBlock *big.Int
}, error) {
	return _DataAvailabilityChallenge.Contract.Challenges(&_DataAvailabilityChallenge.CallOpts, arg0, arg1)
}

// Challenges is a free data retrieval call binding the contract method 0xc4ee20d4.
//
// Solidity: function challenges(uint256 , bytes32 ) view returns(address challenger, uint256 lockedBond, uint256 startBlock, uint256 resolvedBlock)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) Challenges(arg0 *big.Int, arg1 [32]byte) (struct {
	Challenger    common.Address
	LockedBond    *big.Int
	StartBlock    *big.Int
	ResolvedBlock *big.Int
}, error) {
	return _DataAvailabilityChallenge.Contract.Challenges(&_DataAvailabilityChallenge.CallOpts, arg0, arg1)
}

// GetChallengeStatus is a free data retrieval call binding the contract method 0xb740a2db.
//
// Solidity: function getChallengeStatus(uint256 challengedBlockNumber, bytes32 challengedHash) view returns(uint8)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) GetChallengeStatus(opts *bind.CallOpts, challengedBlockNumber *big.Int, challengedHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "getChallengeStatus", challengedBlockNumber, challengedHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetChallengeStatus is a free data retrieval call binding the contract method 0xb740a2db.
//
// Solidity: function getChallengeStatus(uint256 challengedBlockNumber, bytes32 challengedHash) view returns(uint8)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) GetChallengeStatus(challengedBlockNumber *big.Int, challengedHash [32]byte) (uint8, error) {
	return _DataAvailabilityChallenge.Contract.GetChallengeStatus(&_DataAvailabilityChallenge.CallOpts, challengedBlockNumber, challengedHash)
}

// GetChallengeStatus is a free data retrieval call binding the contract method 0xb740a2db.
//
// Solidity: function getChallengeStatus(uint256 challengedBlockNumber, bytes32 challengedHash) view returns(uint8)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) GetChallengeStatus(challengedBlockNumber *big.Int, challengedHash [32]byte) (uint8, error) {
	return _DataAvailabilityChallenge.Contract.GetChallengeStatus(&_DataAvailabilityChallenge.CallOpts, challengedBlockNumber, challengedHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Owner() (common.Address, error) {
	return _DataAvailabilityChallenge.Contract.Owner(&_DataAvailabilityChallenge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) Owner() (common.Address, error) {
	return _DataAvailabilityChallenge.Contract.Owner(&_DataAvailabilityChallenge.CallOpts)
}

// ResolveWindow is a free data retrieval call binding the contract method 0x21cf39ee.
//
// Solidity: function resolveWindow() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) ResolveWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "resolveWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResolveWindow is a free data retrieval call binding the contract method 0x21cf39ee.
//
// Solidity: function resolveWindow() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) ResolveWindow() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.ResolveWindow(&_DataAvailabilityChallenge.CallOpts)
}

// ResolveWindow is a free data retrieval call binding the contract method 0x21cf39ee.
//
// Solidity: function resolveWindow() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) ResolveWindow() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.ResolveWindow(&_DataAvailabilityChallenge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Version() (string, error) {
	return _DataAvailabilityChallenge.Contract.Version(&_DataAvailabilityChallenge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) Version() (string, error) {
	return _DataAvailabilityChallenge.Contract.Version(&_DataAvailabilityChallenge.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x02b2f7c7.
//
// Solidity: function challenge(uint256 challengedBlockNumber, bytes32 challengedHash) payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Challenge(opts *bind.TransactOpts, challengedBlockNumber *big.Int, challengedHash [32]byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "challenge", challengedBlockNumber, challengedHash)
}

// Challenge is a paid mutator transaction binding the contract method 0x02b2f7c7.
//
// Solidity: function challenge(uint256 challengedBlockNumber, bytes32 challengedHash) payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Challenge(challengedBlockNumber *big.Int, challengedHash [32]byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Challenge(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedHash)
}

// Challenge is a paid mutator transaction binding the contract method 0x02b2f7c7.
//
// Solidity: function challenge(uint256 challengedBlockNumber, bytes32 challengedHash) payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Challenge(challengedBlockNumber *big.Int, challengedHash [32]byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Challenge(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedHash)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Deposit() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Deposit(&_DataAvailabilityChallenge.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Deposit() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Deposit(&_DataAvailabilityChallenge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _owner, uint256 _challengeWindow, uint256 _resolveWindow, uint256 _bondSize) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _challengeWindow *big.Int, _resolveWindow *big.Int, _bondSize *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "initialize", _owner, _challengeWindow, _resolveWindow, _bondSize)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _owner, uint256 _challengeWindow, uint256 _resolveWindow, uint256 _bondSize) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Initialize(_owner common.Address, _challengeWindow *big.Int, _resolveWindow *big.Int, _bondSize *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Initialize(&_DataAvailabilityChallenge.TransactOpts, _owner, _challengeWindow, _resolveWindow, _bondSize)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _owner, uint256 _challengeWindow, uint256 _resolveWindow, uint256 _bondSize) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Initialize(_owner common.Address, _challengeWindow *big.Int, _resolveWindow *big.Int, _bondSize *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Initialize(&_DataAvailabilityChallenge.TransactOpts, _owner, _challengeWindow, _resolveWindow, _bondSize)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.RenounceOwnership(&_DataAvailabilityChallenge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.RenounceOwnership(&_DataAvailabilityChallenge.TransactOpts)
}

// Resolve is a paid mutator transaction binding the contract method 0x0b1a73f4.
//
// Solidity: function resolve(uint256 challengedBlockNumber, bytes32 challengedHash, bytes preImage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Resolve(opts *bind.TransactOpts, challengedBlockNumber *big.Int, challengedHash [32]byte, preImage []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "resolve", challengedBlockNumber, challengedHash, preImage)
}

// Resolve is a paid mutator transaction binding the contract method 0x0b1a73f4.
//
// Solidity: function resolve(uint256 challengedBlockNumber, bytes32 challengedHash, bytes preImage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Resolve(challengedBlockNumber *big.Int, challengedHash [32]byte, preImage []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Resolve(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedHash, preImage)
}

// Resolve is a paid mutator transaction binding the contract method 0x0b1a73f4.
//
// Solidity: function resolve(uint256 challengedBlockNumber, bytes32 challengedHash, bytes preImage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Resolve(challengedBlockNumber *big.Int, challengedHash [32]byte, preImage []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Resolve(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedHash, preImage)
}

// SetBondSize is a paid mutator transaction binding the contract method 0xd7d04e54.
//
// Solidity: function setBondSize(uint256 _bondSize) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) SetBondSize(opts *bind.TransactOpts, _bondSize *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "setBondSize", _bondSize)
}

// SetBondSize is a paid mutator transaction binding the contract method 0xd7d04e54.
//
// Solidity: function setBondSize(uint256 _bondSize) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) SetBondSize(_bondSize *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.SetBondSize(&_DataAvailabilityChallenge.TransactOpts, _bondSize)
}

// SetBondSize is a paid mutator transaction binding the contract method 0xd7d04e54.
//
// Solidity: function setBondSize(uint256 _bondSize) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) SetBondSize(_bondSize *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.SetBondSize(&_DataAvailabilityChallenge.TransactOpts, _bondSize)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.TransferOwnership(&_DataAvailabilityChallenge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.TransferOwnership(&_DataAvailabilityChallenge.TransactOpts, newOwner)
}

// UnlockBond is a paid mutator transaction binding the contract method 0xc459f809.
//
// Solidity: function unlockBond(uint256 challengedBlockNumber, bytes32 challengedHash) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) UnlockBond(opts *bind.TransactOpts, challengedBlockNumber *big.Int, challengedHash [32]byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "unlockBond", challengedBlockNumber, challengedHash)
}

// UnlockBond is a paid mutator transaction binding the contract method 0xc459f809.
//
// Solidity: function unlockBond(uint256 challengedBlockNumber, bytes32 challengedHash) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) UnlockBond(challengedBlockNumber *big.Int, challengedHash [32]byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.UnlockBond(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedHash)
}

// UnlockBond is a paid mutator transaction binding the contract method 0xc459f809.
//
// Solidity: function unlockBond(uint256 challengedBlockNumber, bytes32 challengedHash) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) UnlockBond(challengedBlockNumber *big.Int, challengedHash [32]byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.UnlockBond(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Withdraw() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Withdraw(&_DataAvailabilityChallenge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Withdraw(&_DataAvailabilityChallenge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Receive() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Receive(&_DataAvailabilityChallenge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Receive() (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Receive(&_DataAvailabilityChallenge.TransactOpts)
}

// DataAvailabilityChallengeBalanceChangedIterator is returned from FilterBalanceChanged and is used to iterate over the raw logs and unpacked data for BalanceChanged events raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeBalanceChangedIterator struct {
	Event *DataAvailabilityChallengeBalanceChanged // Event containing the contract specifics and raw log

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
func (it *DataAvailabilityChallengeBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataAvailabilityChallengeBalanceChanged)
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
		it.Event = new(DataAvailabilityChallengeBalanceChanged)
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
func (it *DataAvailabilityChallengeBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataAvailabilityChallengeBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataAvailabilityChallengeBalanceChanged represents a BalanceChanged event raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeBalanceChanged struct {
	Account common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceChanged is a free log retrieval operation binding the contract event 0xa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05.
//
// Solidity: event BalanceChanged(address account, uint256 balance)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterBalanceChanged(opts *bind.FilterOpts) (*DataAvailabilityChallengeBalanceChangedIterator, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeBalanceChangedIterator{contract: _DataAvailabilityChallenge.contract, event: "BalanceChanged", logs: logs, sub: sub}, nil
}

// WatchBalanceChanged is a free log subscription operation binding the contract event 0xa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05.
//
// Solidity: event BalanceChanged(address account, uint256 balance)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchBalanceChanged(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeBalanceChanged) (event.Subscription, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataAvailabilityChallengeBalanceChanged)
				if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
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

// ParseBalanceChanged is a log parse operation binding the contract event 0xa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05.
//
// Solidity: event BalanceChanged(address account, uint256 balance)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) ParseBalanceChanged(log types.Log) (*DataAvailabilityChallengeBalanceChanged, error) {
	event := new(DataAvailabilityChallengeBalanceChanged)
	if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataAvailabilityChallengeChallengeStatusChangedIterator is returned from FilterChallengeStatusChanged and is used to iterate over the raw logs and unpacked data for ChallengeStatusChanged events raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeChallengeStatusChangedIterator struct {
	Event *DataAvailabilityChallengeChallengeStatusChanged // Event containing the contract specifics and raw log

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
func (it *DataAvailabilityChallengeChallengeStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataAvailabilityChallengeChallengeStatusChanged)
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
		it.Event = new(DataAvailabilityChallengeChallengeStatusChanged)
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
func (it *DataAvailabilityChallengeChallengeStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataAvailabilityChallengeChallengeStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataAvailabilityChallengeChallengeStatusChanged represents a ChallengeStatusChanged event raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeChallengeStatusChanged struct {
	ChallengedHash        [32]byte
	ChallengedBlockNumber *big.Int
	Status                uint8
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChallengeStatusChanged is a free log retrieval operation binding the contract event 0x73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f8.
//
// Solidity: event ChallengeStatusChanged(bytes32 indexed challengedHash, uint256 indexed challengedBlockNumber, uint8 status)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterChallengeStatusChanged(opts *bind.FilterOpts, challengedHash [][32]byte, challengedBlockNumber []*big.Int) (*DataAvailabilityChallengeChallengeStatusChangedIterator, error) {

	var challengedHashRule []interface{}
	for _, challengedHashItem := range challengedHash {
		challengedHashRule = append(challengedHashRule, challengedHashItem)
	}
	var challengedBlockNumberRule []interface{}
	for _, challengedBlockNumberItem := range challengedBlockNumber {
		challengedBlockNumberRule = append(challengedBlockNumberRule, challengedBlockNumberItem)
	}

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "ChallengeStatusChanged", challengedHashRule, challengedBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeChallengeStatusChangedIterator{contract: _DataAvailabilityChallenge.contract, event: "ChallengeStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChallengeStatusChanged is a free log subscription operation binding the contract event 0x73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f8.
//
// Solidity: event ChallengeStatusChanged(bytes32 indexed challengedHash, uint256 indexed challengedBlockNumber, uint8 status)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchChallengeStatusChanged(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeChallengeStatusChanged, challengedHash [][32]byte, challengedBlockNumber []*big.Int) (event.Subscription, error) {

	var challengedHashRule []interface{}
	for _, challengedHashItem := range challengedHash {
		challengedHashRule = append(challengedHashRule, challengedHashItem)
	}
	var challengedBlockNumberRule []interface{}
	for _, challengedBlockNumberItem := range challengedBlockNumber {
		challengedBlockNumberRule = append(challengedBlockNumberRule, challengedBlockNumberItem)
	}

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "ChallengeStatusChanged", challengedHashRule, challengedBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataAvailabilityChallengeChallengeStatusChanged)
				if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "ChallengeStatusChanged", log); err != nil {
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

// ParseChallengeStatusChanged is a log parse operation binding the contract event 0x73b78891d84bab8633915b22168a5ed8a2f0b86fbaf9733698fbacea9a2b11f8.
//
// Solidity: event ChallengeStatusChanged(bytes32 indexed challengedHash, uint256 indexed challengedBlockNumber, uint8 status)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) ParseChallengeStatusChanged(log types.Log) (*DataAvailabilityChallengeChallengeStatusChanged, error) {
	event := new(DataAvailabilityChallengeChallengeStatusChanged)
	if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "ChallengeStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataAvailabilityChallengeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeInitializedIterator struct {
	Event *DataAvailabilityChallengeInitialized // Event containing the contract specifics and raw log

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
func (it *DataAvailabilityChallengeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataAvailabilityChallengeInitialized)
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
		it.Event = new(DataAvailabilityChallengeInitialized)
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
func (it *DataAvailabilityChallengeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataAvailabilityChallengeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataAvailabilityChallengeInitialized represents a Initialized event raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterInitialized(opts *bind.FilterOpts) (*DataAvailabilityChallengeInitializedIterator, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeInitializedIterator{contract: _DataAvailabilityChallenge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeInitialized) (event.Subscription, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataAvailabilityChallengeInitialized)
				if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) ParseInitialized(log types.Log) (*DataAvailabilityChallengeInitialized, error) {
	event := new(DataAvailabilityChallengeInitialized)
	if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataAvailabilityChallengeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeOwnershipTransferredIterator struct {
	Event *DataAvailabilityChallengeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DataAvailabilityChallengeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataAvailabilityChallengeOwnershipTransferred)
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
		it.Event = new(DataAvailabilityChallengeOwnershipTransferred)
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
func (it *DataAvailabilityChallengeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataAvailabilityChallengeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataAvailabilityChallengeOwnershipTransferred represents a OwnershipTransferred event raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DataAvailabilityChallengeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeOwnershipTransferredIterator{contract: _DataAvailabilityChallenge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataAvailabilityChallengeOwnershipTransferred)
				if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) ParseOwnershipTransferred(log types.Log) (*DataAvailabilityChallengeOwnershipTransferred, error) {
	event := new(DataAvailabilityChallengeOwnershipTransferred)
	if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataAvailabilityChallengeRequiredBondSizeChangedIterator is returned from FilterRequiredBondSizeChanged and is used to iterate over the raw logs and unpacked data for RequiredBondSizeChanged events raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeRequiredBondSizeChangedIterator struct {
	Event *DataAvailabilityChallengeRequiredBondSizeChanged // Event containing the contract specifics and raw log

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
func (it *DataAvailabilityChallengeRequiredBondSizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataAvailabilityChallengeRequiredBondSizeChanged)
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
		it.Event = new(DataAvailabilityChallengeRequiredBondSizeChanged)
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
func (it *DataAvailabilityChallengeRequiredBondSizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataAvailabilityChallengeRequiredBondSizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataAvailabilityChallengeRequiredBondSizeChanged represents a RequiredBondSizeChanged event raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeRequiredBondSizeChanged struct {
	ChallengeWindow *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRequiredBondSizeChanged is a free log retrieval operation binding the contract event 0x4468d695a0389e5f9e8ef0c9aee6d84e74cc0d0e0a28c8413badb54697d1bbae.
//
// Solidity: event RequiredBondSizeChanged(uint256 challengeWindow)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterRequiredBondSizeChanged(opts *bind.FilterOpts) (*DataAvailabilityChallengeRequiredBondSizeChangedIterator, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "RequiredBondSizeChanged")
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeRequiredBondSizeChangedIterator{contract: _DataAvailabilityChallenge.contract, event: "RequiredBondSizeChanged", logs: logs, sub: sub}, nil
}

// WatchRequiredBondSizeChanged is a free log subscription operation binding the contract event 0x4468d695a0389e5f9e8ef0c9aee6d84e74cc0d0e0a28c8413badb54697d1bbae.
//
// Solidity: event RequiredBondSizeChanged(uint256 challengeWindow)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchRequiredBondSizeChanged(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeRequiredBondSizeChanged) (event.Subscription, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "RequiredBondSizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataAvailabilityChallengeRequiredBondSizeChanged)
				if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "RequiredBondSizeChanged", log); err != nil {
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

// ParseRequiredBondSizeChanged is a log parse operation binding the contract event 0x4468d695a0389e5f9e8ef0c9aee6d84e74cc0d0e0a28c8413badb54697d1bbae.
//
// Solidity: event RequiredBondSizeChanged(uint256 challengeWindow)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) ParseRequiredBondSizeChanged(log types.Log) (*DataAvailabilityChallengeRequiredBondSizeChanged, error) {
	event := new(DataAvailabilityChallengeRequiredBondSizeChanged)
	if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "RequiredBondSizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
