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

// Challenge is an auto generated low-level Go binding around an user-defined struct.
type Challenge struct {
	Challenger    common.Address
	LockedBond    *big.Int
	StartBlock    *big.Int
	ResolvedBlock *big.Int
}

// DataAvailabilityChallengeMetaData contains all meta data concerning the DataAvailabilityChallenge contract.
var DataAvailabilityChallengeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[{\"name\":\"challengedBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"challengeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fixedResolutionCost\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallenge\",\"inputs\":[{\"name\":\"challengedBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structChallenge\",\"components\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockedBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolvedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengeStatus\",\"inputs\":[{\"name\":\"challengedBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChallengeStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_challengeWindow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_resolveWindow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_bondSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_resolverRefundPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"challengedBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"resolveData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolverRefundPercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBondSize\",\"inputs\":[{\"name\":\"_bondSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setResolverRefundPercentage\",\"inputs\":[{\"name\":\"_resolverRefundPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockBond\",\"inputs\":[{\"name\":\"challengedBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateCommitment\",\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"variableResolutionCost\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"variableResolutionCostPrecision\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BalanceChanged\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeStatusChanged\",\"inputs\":[{\"name\":\"challengedBlockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"challengedCommitment\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChallengeStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequiredBondSizeChanged\",\"inputs\":[{\"name\":\"challengeWindow\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResolverRefundPercentageChanged\",\"inputs\":[{\"name\":\"resolverRefundPercentage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BondTooLow\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChallengeExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowNotOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitmentLength\",\"inputs\":[{\"name\":\"commitmentType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"expectedLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInputData\",\"inputs\":[{\"name\":\"providedDataCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expectedCommitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidResolverRefundPercentage\",\"inputs\":[{\"name\":\"invalidResolverRefundPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnknownCommitmentType\",\"inputs\":[{\"name\":\"commitmentType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100de565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156100dc576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611de6806100ed6000396000f3fe60806040526004361061018f5760003560e01c8063848afb3d116100d6578063956118521161007f578063d7d04e5411610059578063d7d04e5414610484578063f2fde38b146104a4578063f92ad219146104c457600080fd5b80639561185214610452578063a03aafbf14610469578063d0e30db01461047c57600080fd5b80638ecb85e1116100b05780638ecb85e1146103fc578063939882331461041257806393fb19441461043257600080fd5b8063848afb3d14610348578063861a1412146103b15780638da5cb5b146103c757600080fd5b806354fd4d501161013857806379e8a8b31161011257806379e8a8b3146102db5780637ae929d91461030857806382fead4d1461032857600080fd5b806354fd4d501461025a5780637099c581146102b0578063715018a6146102c657600080fd5b8063336409fd11610169578063336409fd1461020f5780633ccfd60b1461022f5780634ebaf3ce1461024457600080fd5b806321cf39ee146101a357806323c30f59146101cc57806327e235e3146101e257600080fd5b3661019e5761019c6104e4565b005b600080fd5b3480156101af57600080fd5b506101b960665481565b6040519081526020015b60405180910390f35b3480156101d857600080fd5b506101b961410081565b3480156101ee57600080fd5b506101b96101fd366004611869565b60696020526000908152604090205481565b34801561021b57600080fd5b5061019c61022a366004611884565b610552565b34801561023b57600080fd5b5061019c6105a2565b34801561025057600080fd5b506101b96103e881565b34801561026657600080fd5b506102a36040518060400160405280600c81526020017f312e302e312d626574612e35000000000000000000000000000000000000000081525081565b6040516101c39190611908565b3480156102bc57600080fd5b506101b960675481565b3480156102d257600080fd5b5061019c61063c565b3480156102e757600080fd5b506102fb6102f6366004611964565b610650565b6040516101c39190611a1a565b34801561031457600080fd5b5061019c610323366004611a28565b610713565b34801561033457600080fd5b5061019c610343366004611aa2565b610802565b34801561035457600080fd5b50610368610363366004611964565b61094f565b6040516101c39190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015190820152604080830151908201526060918201519181019190915260800190565b3480156103bd57600080fd5b506101b960655481565b3480156103d357600080fd5b5060335460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c3565b34801561040857600080fd5b506101b960685481565b34801561041e57600080fd5b5061019c61042d366004611964565b610a0b565b34801561043e57600080fd5b5061019c61044d366004611b0e565b610b3d565b34801561045e57600080fd5b506101b962011cdd81565b61019c610477366004611964565b610cad565b61019c6104e4565b34801561049057600080fd5b5061019c61049f366004611884565b610ef6565b3480156104b057600080fd5b5061019c6104bf366004611869565b610f39565b3480156104d057600080fd5b5061019c6104df366004611b50565b610ff0565b3360009081526069602052604081208054349290610503908490611bc1565b909155505033600081815260696020908152604091829020548251938452908301527fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a1565b61055a6111ab565b606481111561059d576040517f16aa4e80000000000000000000000000000000000000000000000000000000008152600481018290526024015b60405180910390fd5b606855565b336000818152606960209081526040808320805490849055815194855291840192909252917fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a160006105ff335a8461122c565b905080610638576040517f27fcd9d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6106446111ab565b61064e6000611240565b565b6000838152606a602052604080822090518291906106719086908690611bd9565b9081526040805160209281900383018120608082018352805473ffffffffffffffffffffffffffffffffffffffff16808352600182015494830194909452600281015492820192909252600390910154606082015291506106d657600091505061070c565b6060810151156106ea57600291505061070c565b6106f781604001516112b7565b1561070657600191505061070c565b60039150505b9392505050565b61071d8484610b3d565b600161072a868686610650565b600381111561073b5761073b6119b0565b14610772576040517fbeb11d3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61077e84848484610802565b6000858152606a6020526040808220905161079c9087908790611bd9565b908152604051908190036020018120436003820155915086907fc5d8c630ba2fdacb1db24c4599df78c7fb8cf97b5aecde34939597f6697bb1ad906107e79088908890600290611c32565b60405180910390a26107fa8183336112d0565b505050505050565b600061080e85856114f0565b905060606000826001811115610826576108266119b0565b036108715761086a84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061151992505050565b90506108f2565b6001826001811115610885576108856119b0565b036108f25760006108968787611552565b905060008180156108a9576108a96119b0565b036108f0576108ed85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061160492505050565b91505b505b8585604051610902929190611bd9565b60405180910390208180519060200120146107fa578086866040517f1a0bbf9f00000000000000000000000000000000000000000000000000000000815260040161059493929190611c5d565b6109906040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6000848152606a60205260409081902090516109af9085908590611bd9565b908152604080519182900360209081018320608084018352805473ffffffffffffffffffffffffffffffffffffffff16845260018101549184019190915260028101549183019190915260030154606082015290509392505050565b6003610a18848484610650565b6003811115610a2957610a296119b0565b14610a60576040517f151f07fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838152606a60205260408082209051610a7e9085908590611bd9565b90815260408051602092819003830190206001810154815473ffffffffffffffffffffffffffffffffffffffff16600090815260699094529183208054919450919290610acc908490611bc1565b9091555050600060018201819055815473ffffffffffffffffffffffffffffffffffffffff1680825260696020908152604092839020548351928352908201527fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a150505050565b6000610b4983836114f0565b90506000816001811115610b5f57610b5f6119b0565b03610bb45760218214610baf576040517ffd9a7e5b000000000000000000000000000000000000000000000000000000008152600060048201526021602482015260448101839052606401610594565b505050565b6001816001811115610bc857610bc86119b0565b03610c63576000610bd98484611552565b90506000818015610bec57610bec6119b0565b03610c5d576022831080610c1557506020610c08600285611c8d565b610c129190611cd3565b15155b15610c5d576040517ffd9a7e5b000000000000000000000000000000000000000000000000000000008152600160048201526022602482015260448101849052606401610594565b50505050565b806001811115610c7557610c756119b0565b6040517f81ff071300000000000000000000000000000000000000000000000000000000815260ff9091166004820152602401610594565b610cb78282610b3d565b610cbf6104e4565b606754336000908152606960205260409020541015610d275733600090815260696020526040908190205460675491517e0155b50000000000000000000000000000000000000000000000000000000081526105949290600401918252602082015260400190565b6000610d34848484610650565b6003811115610d4557610d456119b0565b14610d7c576040517f9bb6c64e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d8583611683565b610dbb576040517ff9e0d1f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6067543360009081526069602052604081208054909190610ddd908490611c8d565b9250508190555060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200160675481526020014381526020016000815250606a60008581526020019081526020016000208383604051610e41929190611bd9565b9081526040805160209281900383018120845181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161781559284015160018085019190915591840151600284015560609093015160039092019190915584917fc5d8c630ba2fdacb1db24c4599df78c7fb8cf97b5aecde34939597f6697bb1ad91610ee99186918691611c32565b60405180910390a2505050565b610efe6111ab565b60678190556040518181527f4468d695a0389e5f9e8ef0c9aee6d84e74cc0d0e0a28c8413badb54697d1bbae9060200160405180910390a150565b610f416111ab565b73ffffffffffffffffffffffffffffffffffffffff8116610fe4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610594565b610fed81611240565b50565b600054610100900460ff16158080156110105750600054600160ff909116105b8061102a5750303b15801561102a575060005460ff166001145b6110b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610594565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561111457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61111c61169d565b6065859055606684905561112f83610ef6565b61113882610552565b61114186611240565b80156107fa57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461064e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610594565b6000806000806000858888f1949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000606654826112c79190611bc1565b43111592915050565b6001830154835473ffffffffffffffffffffffffffffffffffffffff166000486103e86112ff61410088611ce7565b6113099190611d24565b6113169062011cdd611bc1565b6113209190611ce7565b9050808311156113d2576113348184611c8d565b73ffffffffffffffffffffffffffffffffffffffff831660009081526069602052604081208054909190611369908490611bc1565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000818152606960209081526040918290205482519384529083015291935083917fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a15b60006064606854836113e49190611ce7565b6113ee9190611d24565b9050838111156113fb5750825b80156114a85773ffffffffffffffffffffffffffffffffffffffff851660009081526069602052604081208054839290611436908490611bc1565b9091555061144690508185611c8d565b73ffffffffffffffffffffffffffffffffffffffff8616600081815260696020908152604091829020548251938452908301529195507fa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05910160405180910390a15b83156114dd5760405160009085156108fc0290869083818181858288f193505050501580156114db573d6000803e3d6000fd5b505b6000876001018190555050505050505050565b60006114fc838361173c565b60ff166001811115611510576115106119b0565b90505b92915050565b805160208083019190912060405160009281019290925260218201526060906041015b6040516020818303038152906040529050919050565b60008061156a6115658460018188611d38565b61173c565b90506115766000611752565b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168160f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036115cd576000915050611513565b6040517f1f2387d700000000000000000000000000000000000000000000000000000000815260ff82166004820152602401610594565b60607f01000000000000000000000000000000000000000000000000000000000000006116316000611752565b838051906020012060405160200161153c939291907fff000000000000000000000000000000000000000000000000000000000000009384168152919092166001820152600281019190915260220190565b600081431015801561151357506065546112c79083611bc1565b600054610100900460ff16611734576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610594565b61064e6117a0565b60006117488284611d62565b60f81c9392505050565b600080828015611764576117646119b0565b0361179057507fff00000000000000000000000000000000000000000000000000000000000000919050565b611798611daa565b506000919050565b600054610100900460ff16611837576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610594565b61064e33611240565b803573ffffffffffffffffffffffffffffffffffffffff8116811461186457600080fd5b919050565b60006020828403121561187b57600080fd5b61151082611840565b60006020828403121561189657600080fd5b5035919050565b6000815180845260005b818110156118c3576020818501810151868301820152016118a7565b818111156118d5576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611510602083018461189d565b60008083601f84011261192d57600080fd5b50813567ffffffffffffffff81111561194557600080fd5b60208301915083602082850101111561195d57600080fd5b9250929050565b60008060006040848603121561197957600080fd5b83359250602084013567ffffffffffffffff81111561199757600080fd5b6119a38682870161191b565b9497909650939450505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110611a16577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b6020810161151382846119df565b600080600080600060608688031215611a4057600080fd5b85359450602086013567ffffffffffffffff80821115611a5f57600080fd5b611a6b89838a0161191b565b90965094506040880135915080821115611a8457600080fd5b50611a918882890161191b565b969995985093965092949392505050565b60008060008060408587031215611ab857600080fd5b843567ffffffffffffffff80821115611ad057600080fd5b611adc8883890161191b565b90965094506020870135915080821115611af557600080fd5b50611b028782880161191b565b95989497509550505050565b60008060208385031215611b2157600080fd5b823567ffffffffffffffff811115611b3857600080fd5b611b448582860161191b565b90969095509350505050565b600080600080600060a08688031215611b6857600080fd5b611b7186611840565b97602087013597506040870135966060810135965060800135945092505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611bd457611bd4611b92565b500190565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000611c46604083018587611be9565b9050611c5560208301846119df565b949350505050565b604081526000611c70604083018661189d565b8281036020840152611c83818587611be9565b9695505050505050565b600082821015611c9f57611c9f611b92565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611ce257611ce2611ca4565b500690565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d1f57611d1f611b92565b500290565b600082611d3357611d33611ca4565b500490565b60008085851115611d4857600080fd5b83861115611d5557600080fd5b5050820193919092039150565b7fff000000000000000000000000000000000000000000000000000000000000008135818116916001851015611da25780818660010360031b1b83161692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea164736f6c634300080f000a",
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

// FixedResolutionCost is a free data retrieval call binding the contract method 0x95611852.
//
// Solidity: function fixedResolutionCost() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) FixedResolutionCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "fixedResolutionCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FixedResolutionCost is a free data retrieval call binding the contract method 0x95611852.
//
// Solidity: function fixedResolutionCost() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) FixedResolutionCost() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.FixedResolutionCost(&_DataAvailabilityChallenge.CallOpts)
}

// FixedResolutionCost is a free data retrieval call binding the contract method 0x95611852.
//
// Solidity: function fixedResolutionCost() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) FixedResolutionCost() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.FixedResolutionCost(&_DataAvailabilityChallenge.CallOpts)
}

// GetChallenge is a free data retrieval call binding the contract method 0x848afb3d.
//
// Solidity: function getChallenge(uint256 challengedBlockNumber, bytes challengedCommitment) view returns((address,uint256,uint256,uint256))
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) GetChallenge(opts *bind.CallOpts, challengedBlockNumber *big.Int, challengedCommitment []byte) (Challenge, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "getChallenge", challengedBlockNumber, challengedCommitment)

	if err != nil {
		return *new(Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new(Challenge)).(*Challenge)

	return out0, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x848afb3d.
//
// Solidity: function getChallenge(uint256 challengedBlockNumber, bytes challengedCommitment) view returns((address,uint256,uint256,uint256))
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) GetChallenge(challengedBlockNumber *big.Int, challengedCommitment []byte) (Challenge, error) {
	return _DataAvailabilityChallenge.Contract.GetChallenge(&_DataAvailabilityChallenge.CallOpts, challengedBlockNumber, challengedCommitment)
}

// GetChallenge is a free data retrieval call binding the contract method 0x848afb3d.
//
// Solidity: function getChallenge(uint256 challengedBlockNumber, bytes challengedCommitment) view returns((address,uint256,uint256,uint256))
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) GetChallenge(challengedBlockNumber *big.Int, challengedCommitment []byte) (Challenge, error) {
	return _DataAvailabilityChallenge.Contract.GetChallenge(&_DataAvailabilityChallenge.CallOpts, challengedBlockNumber, challengedCommitment)
}

// GetChallengeStatus is a free data retrieval call binding the contract method 0x79e8a8b3.
//
// Solidity: function getChallengeStatus(uint256 challengedBlockNumber, bytes challengedCommitment) view returns(uint8)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) GetChallengeStatus(opts *bind.CallOpts, challengedBlockNumber *big.Int, challengedCommitment []byte) (uint8, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "getChallengeStatus", challengedBlockNumber, challengedCommitment)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetChallengeStatus is a free data retrieval call binding the contract method 0x79e8a8b3.
//
// Solidity: function getChallengeStatus(uint256 challengedBlockNumber, bytes challengedCommitment) view returns(uint8)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) GetChallengeStatus(challengedBlockNumber *big.Int, challengedCommitment []byte) (uint8, error) {
	return _DataAvailabilityChallenge.Contract.GetChallengeStatus(&_DataAvailabilityChallenge.CallOpts, challengedBlockNumber, challengedCommitment)
}

// GetChallengeStatus is a free data retrieval call binding the contract method 0x79e8a8b3.
//
// Solidity: function getChallengeStatus(uint256 challengedBlockNumber, bytes challengedCommitment) view returns(uint8)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) GetChallengeStatus(challengedBlockNumber *big.Int, challengedCommitment []byte) (uint8, error) {
	return _DataAvailabilityChallenge.Contract.GetChallengeStatus(&_DataAvailabilityChallenge.CallOpts, challengedBlockNumber, challengedCommitment)
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

// ResolverRefundPercentage is a free data retrieval call binding the contract method 0x8ecb85e1.
//
// Solidity: function resolverRefundPercentage() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) ResolverRefundPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "resolverRefundPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResolverRefundPercentage is a free data retrieval call binding the contract method 0x8ecb85e1.
//
// Solidity: function resolverRefundPercentage() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) ResolverRefundPercentage() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.ResolverRefundPercentage(&_DataAvailabilityChallenge.CallOpts)
}

// ResolverRefundPercentage is a free data retrieval call binding the contract method 0x8ecb85e1.
//
// Solidity: function resolverRefundPercentage() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) ResolverRefundPercentage() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.ResolverRefundPercentage(&_DataAvailabilityChallenge.CallOpts)
}

// ValidateCommitment is a free data retrieval call binding the contract method 0x93fb1944.
//
// Solidity: function validateCommitment(bytes commitment) pure returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) ValidateCommitment(opts *bind.CallOpts, commitment []byte) error {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "validateCommitment", commitment)

	if err != nil {
		return err
	}

	return err

}

// ValidateCommitment is a free data retrieval call binding the contract method 0x93fb1944.
//
// Solidity: function validateCommitment(bytes commitment) pure returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) ValidateCommitment(commitment []byte) error {
	return _DataAvailabilityChallenge.Contract.ValidateCommitment(&_DataAvailabilityChallenge.CallOpts, commitment)
}

// ValidateCommitment is a free data retrieval call binding the contract method 0x93fb1944.
//
// Solidity: function validateCommitment(bytes commitment) pure returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) ValidateCommitment(commitment []byte) error {
	return _DataAvailabilityChallenge.Contract.ValidateCommitment(&_DataAvailabilityChallenge.CallOpts, commitment)
}

// VariableResolutionCost is a free data retrieval call binding the contract method 0x23c30f59.
//
// Solidity: function variableResolutionCost() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) VariableResolutionCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "variableResolutionCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VariableResolutionCost is a free data retrieval call binding the contract method 0x23c30f59.
//
// Solidity: function variableResolutionCost() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) VariableResolutionCost() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.VariableResolutionCost(&_DataAvailabilityChallenge.CallOpts)
}

// VariableResolutionCost is a free data retrieval call binding the contract method 0x23c30f59.
//
// Solidity: function variableResolutionCost() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) VariableResolutionCost() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.VariableResolutionCost(&_DataAvailabilityChallenge.CallOpts)
}

// VariableResolutionCostPrecision is a free data retrieval call binding the contract method 0x4ebaf3ce.
//
// Solidity: function variableResolutionCostPrecision() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCaller) VariableResolutionCostPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataAvailabilityChallenge.contract.Call(opts, &out, "variableResolutionCostPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VariableResolutionCostPrecision is a free data retrieval call binding the contract method 0x4ebaf3ce.
//
// Solidity: function variableResolutionCostPrecision() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) VariableResolutionCostPrecision() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.VariableResolutionCostPrecision(&_DataAvailabilityChallenge.CallOpts)
}

// VariableResolutionCostPrecision is a free data retrieval call binding the contract method 0x4ebaf3ce.
//
// Solidity: function variableResolutionCostPrecision() view returns(uint256)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession) VariableResolutionCostPrecision() (*big.Int, error) {
	return _DataAvailabilityChallenge.Contract.VariableResolutionCostPrecision(&_DataAvailabilityChallenge.CallOpts)
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

// Challenge is a paid mutator transaction binding the contract method 0xa03aafbf.
//
// Solidity: function challenge(uint256 challengedBlockNumber, bytes challengedCommitment) payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Challenge(opts *bind.TransactOpts, challengedBlockNumber *big.Int, challengedCommitment []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "challenge", challengedBlockNumber, challengedCommitment)
}

// Challenge is a paid mutator transaction binding the contract method 0xa03aafbf.
//
// Solidity: function challenge(uint256 challengedBlockNumber, bytes challengedCommitment) payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Challenge(challengedBlockNumber *big.Int, challengedCommitment []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Challenge(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedCommitment)
}

// Challenge is a paid mutator transaction binding the contract method 0xa03aafbf.
//
// Solidity: function challenge(uint256 challengedBlockNumber, bytes challengedCommitment) payable returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Challenge(challengedBlockNumber *big.Int, challengedCommitment []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Challenge(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedCommitment)
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

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address _owner, uint256 _challengeWindow, uint256 _resolveWindow, uint256 _bondSize, uint256 _resolverRefundPercentage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _challengeWindow *big.Int, _resolveWindow *big.Int, _bondSize *big.Int, _resolverRefundPercentage *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "initialize", _owner, _challengeWindow, _resolveWindow, _bondSize, _resolverRefundPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address _owner, uint256 _challengeWindow, uint256 _resolveWindow, uint256 _bondSize, uint256 _resolverRefundPercentage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Initialize(_owner common.Address, _challengeWindow *big.Int, _resolveWindow *big.Int, _bondSize *big.Int, _resolverRefundPercentage *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Initialize(&_DataAvailabilityChallenge.TransactOpts, _owner, _challengeWindow, _resolveWindow, _bondSize, _resolverRefundPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address _owner, uint256 _challengeWindow, uint256 _resolveWindow, uint256 _bondSize, uint256 _resolverRefundPercentage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Initialize(_owner common.Address, _challengeWindow *big.Int, _resolveWindow *big.Int, _bondSize *big.Int, _resolverRefundPercentage *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Initialize(&_DataAvailabilityChallenge.TransactOpts, _owner, _challengeWindow, _resolveWindow, _bondSize, _resolverRefundPercentage)
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

// Resolve is a paid mutator transaction binding the contract method 0x7ae929d9.
//
// Solidity: function resolve(uint256 challengedBlockNumber, bytes challengedCommitment, bytes resolveData) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) Resolve(opts *bind.TransactOpts, challengedBlockNumber *big.Int, challengedCommitment []byte, resolveData []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "resolve", challengedBlockNumber, challengedCommitment, resolveData)
}

// Resolve is a paid mutator transaction binding the contract method 0x7ae929d9.
//
// Solidity: function resolve(uint256 challengedBlockNumber, bytes challengedCommitment, bytes resolveData) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) Resolve(challengedBlockNumber *big.Int, challengedCommitment []byte, resolveData []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Resolve(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedCommitment, resolveData)
}

// Resolve is a paid mutator transaction binding the contract method 0x7ae929d9.
//
// Solidity: function resolve(uint256 challengedBlockNumber, bytes challengedCommitment, bytes resolveData) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) Resolve(challengedBlockNumber *big.Int, challengedCommitment []byte, resolveData []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.Resolve(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedCommitment, resolveData)
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

// SetResolverRefundPercentage is a paid mutator transaction binding the contract method 0x336409fd.
//
// Solidity: function setResolverRefundPercentage(uint256 _resolverRefundPercentage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) SetResolverRefundPercentage(opts *bind.TransactOpts, _resolverRefundPercentage *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "setResolverRefundPercentage", _resolverRefundPercentage)
}

// SetResolverRefundPercentage is a paid mutator transaction binding the contract method 0x336409fd.
//
// Solidity: function setResolverRefundPercentage(uint256 _resolverRefundPercentage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) SetResolverRefundPercentage(_resolverRefundPercentage *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.SetResolverRefundPercentage(&_DataAvailabilityChallenge.TransactOpts, _resolverRefundPercentage)
}

// SetResolverRefundPercentage is a paid mutator transaction binding the contract method 0x336409fd.
//
// Solidity: function setResolverRefundPercentage(uint256 _resolverRefundPercentage) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) SetResolverRefundPercentage(_resolverRefundPercentage *big.Int) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.SetResolverRefundPercentage(&_DataAvailabilityChallenge.TransactOpts, _resolverRefundPercentage)
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

// UnlockBond is a paid mutator transaction binding the contract method 0x93988233.
//
// Solidity: function unlockBond(uint256 challengedBlockNumber, bytes challengedCommitment) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactor) UnlockBond(opts *bind.TransactOpts, challengedBlockNumber *big.Int, challengedCommitment []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.contract.Transact(opts, "unlockBond", challengedBlockNumber, challengedCommitment)
}

// UnlockBond is a paid mutator transaction binding the contract method 0x93988233.
//
// Solidity: function unlockBond(uint256 challengedBlockNumber, bytes challengedCommitment) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeSession) UnlockBond(challengedBlockNumber *big.Int, challengedCommitment []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.UnlockBond(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedCommitment)
}

// UnlockBond is a paid mutator transaction binding the contract method 0x93988233.
//
// Solidity: function unlockBond(uint256 challengedBlockNumber, bytes challengedCommitment) returns()
func (_DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession) UnlockBond(challengedBlockNumber *big.Int, challengedCommitment []byte) (*types.Transaction, error) {
	return _DataAvailabilityChallenge.Contract.UnlockBond(&_DataAvailabilityChallenge.TransactOpts, challengedBlockNumber, challengedCommitment)
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
	ChallengedBlockNumber *big.Int
	ChallengedCommitment  []byte
	Status                uint8
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChallengeStatusChanged is a free log retrieval operation binding the contract event 0xc5d8c630ba2fdacb1db24c4599df78c7fb8cf97b5aecde34939597f6697bb1ad.
//
// Solidity: event ChallengeStatusChanged(uint256 indexed challengedBlockNumber, bytes challengedCommitment, uint8 status)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterChallengeStatusChanged(opts *bind.FilterOpts, challengedBlockNumber []*big.Int) (*DataAvailabilityChallengeChallengeStatusChangedIterator, error) {

	var challengedBlockNumberRule []interface{}
	for _, challengedBlockNumberItem := range challengedBlockNumber {
		challengedBlockNumberRule = append(challengedBlockNumberRule, challengedBlockNumberItem)
	}

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "ChallengeStatusChanged", challengedBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeChallengeStatusChangedIterator{contract: _DataAvailabilityChallenge.contract, event: "ChallengeStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChallengeStatusChanged is a free log subscription operation binding the contract event 0xc5d8c630ba2fdacb1db24c4599df78c7fb8cf97b5aecde34939597f6697bb1ad.
//
// Solidity: event ChallengeStatusChanged(uint256 indexed challengedBlockNumber, bytes challengedCommitment, uint8 status)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchChallengeStatusChanged(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeChallengeStatusChanged, challengedBlockNumber []*big.Int) (event.Subscription, error) {

	var challengedBlockNumberRule []interface{}
	for _, challengedBlockNumberItem := range challengedBlockNumber {
		challengedBlockNumberRule = append(challengedBlockNumberRule, challengedBlockNumberItem)
	}

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "ChallengeStatusChanged", challengedBlockNumberRule)
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

// ParseChallengeStatusChanged is a log parse operation binding the contract event 0xc5d8c630ba2fdacb1db24c4599df78c7fb8cf97b5aecde34939597f6697bb1ad.
//
// Solidity: event ChallengeStatusChanged(uint256 indexed challengedBlockNumber, bytes challengedCommitment, uint8 status)
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

// DataAvailabilityChallengeResolverRefundPercentageChangedIterator is returned from FilterResolverRefundPercentageChanged and is used to iterate over the raw logs and unpacked data for ResolverRefundPercentageChanged events raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeResolverRefundPercentageChangedIterator struct {
	Event *DataAvailabilityChallengeResolverRefundPercentageChanged // Event containing the contract specifics and raw log

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
func (it *DataAvailabilityChallengeResolverRefundPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataAvailabilityChallengeResolverRefundPercentageChanged)
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
		it.Event = new(DataAvailabilityChallengeResolverRefundPercentageChanged)
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
func (it *DataAvailabilityChallengeResolverRefundPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataAvailabilityChallengeResolverRefundPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataAvailabilityChallengeResolverRefundPercentageChanged represents a ResolverRefundPercentageChanged event raised by the DataAvailabilityChallenge contract.
type DataAvailabilityChallengeResolverRefundPercentageChanged struct {
	ResolverRefundPercentage *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterResolverRefundPercentageChanged is a free log retrieval operation binding the contract event 0xbbd8605de8f773fedb355c2fecd4b6b2e16a10a44e6676a63b375ac8f693758d.
//
// Solidity: event ResolverRefundPercentageChanged(uint256 resolverRefundPercentage)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) FilterResolverRefundPercentageChanged(opts *bind.FilterOpts) (*DataAvailabilityChallengeResolverRefundPercentageChangedIterator, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.FilterLogs(opts, "ResolverRefundPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &DataAvailabilityChallengeResolverRefundPercentageChangedIterator{contract: _DataAvailabilityChallenge.contract, event: "ResolverRefundPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchResolverRefundPercentageChanged is a free log subscription operation binding the contract event 0xbbd8605de8f773fedb355c2fecd4b6b2e16a10a44e6676a63b375ac8f693758d.
//
// Solidity: event ResolverRefundPercentageChanged(uint256 resolverRefundPercentage)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) WatchResolverRefundPercentageChanged(opts *bind.WatchOpts, sink chan<- *DataAvailabilityChallengeResolverRefundPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _DataAvailabilityChallenge.contract.WatchLogs(opts, "ResolverRefundPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataAvailabilityChallengeResolverRefundPercentageChanged)
				if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "ResolverRefundPercentageChanged", log); err != nil {
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

// ParseResolverRefundPercentageChanged is a log parse operation binding the contract event 0xbbd8605de8f773fedb355c2fecd4b6b2e16a10a44e6676a63b375ac8f693758d.
//
// Solidity: event ResolverRefundPercentageChanged(uint256 resolverRefundPercentage)
func (_DataAvailabilityChallenge *DataAvailabilityChallengeFilterer) ParseResolverRefundPercentageChanged(log types.Log) (*DataAvailabilityChallengeResolverRefundPercentageChanged, error) {
	event := new(DataAvailabilityChallengeResolverRefundPercentageChanged)
	if err := _DataAvailabilityChallenge.contract.UnpackLog(event, "ResolverRefundPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
