// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test_helpersgen

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

// BridgeTesterMetaData contains all meta data concerning the BridgeTester contract.
var BridgeTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotDelayedInbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOutbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"rollup_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161149a61002d6000396000505061149a6000f3fe60806040526004361061016d5760003560e01c80639e5d4c49116100cb578063cee3d7281161007f578063e77145f411610059578063e77145f414610229578063eca067ad14610439578063ee35f3271461044e57600080fd5b8063cee3d728146103d9578063d5719dc2146103f9578063e76f5c8d1461041957600080fd5b8063ae60bd13116100b0578063ae60bd131461035c578063c4d66de814610399578063cb23bcb5146103b957600080fd5b80639e5d4c4914610319578063ab5d89431461034757600080fd5b80635fca4a161161012257806386598a561161010757806386598a56146102845780638db5993b146102ce578063945e1147146102e157600080fd5b80635fca4a161461024b5780637a88b1071461026157600080fd5b8063413b35bd11610153578063413b35bd146101bd57806347fb24c5146102095780634f61f8501461022b57600080fd5b806284120c1461017957806316bf55791461019d57600080fd5b3661017457005b600080fd5b34801561018557600080fd5b506009545b6040519081526020015b60405180910390f35b3480156101a957600080fd5b5061018a6101b836600461119b565b61046e565b3480156101c957600080fd5b506101f96101d83660046111cc565b6001600160a01b031660009081526002602052604090206001015460ff1690565b6040519015158152602001610194565b34801561021557600080fd5b506102296102243660046111f0565b61048f565b005b34801561023757600080fd5b506102296102463660046111cc565b610795565b34801561025757600080fd5b5061018a600a5481565b34801561026d57600080fd5b5061018a61027c36600461122e565b600092915050565b34801561029057600080fd5b506102ae61029f36600461125a565b50600093849350839250829150565b604080519485526020850193909352918301526060820152608001610194565b61018a6102dc36600461128c565b6108c0565b3480156102ed57600080fd5b506103016102fc36600461119b565b610924565b6040516001600160a01b039091168152602001610194565b34801561032557600080fd5b506103396103343660046112d3565b61094e565b60405161019492919061135c565b34801561035357600080fd5b50610301610afb565b34801561036857600080fd5b506101f96103773660046111cc565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b3480156103a557600080fd5b506102296103b43660046111cc565b610b46565b3480156103c557600080fd5b50600654610301906001600160a01b031681565b3480156103e557600080fd5b506102296103f43660046111f0565b610cb2565b34801561040557600080fd5b5061018a61041436600461119b565b610fb3565b34801561042557600080fd5b5061030161043436600461119b565b610fc3565b34801561044557600080fd5b5060085461018a565b34801561045a57600080fd5b50600754610301906001600160a01b031681565b6009818154811061047e57600080fd5b600091825260209091200154905081565b6006546001600160a01b0316331461055e5760065460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa1580156104eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050f91906113b4565b9050336001600160a01b0382161461055c57600654604051630739600760e01b81523360048201526001600160a01b03918216602482015290821660448201526064015b60405180910390fd5b505b6001600160a01b0382166000818152600160208181526040928390209182015492518515158152919360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a2821515811515036105ca5750505050565b821561066557604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916909117905561078e565b60038054610675906001906113d1565b81548110610685576106856113f2565b6000918252602090912001548254600380546001600160a01b039093169290919081106106b4576106b46113f2565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610702576107026113f2565b60009182526020808320909101546001600160a01b03168352820192909252604001902055600380548061073857610738611408565b600082815260208082208301600019908101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff191690555b50505b5050565b6006546001600160a01b0316331461085f5760065460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa1580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081591906113b4565b9050336001600160a01b0382161461085d57600654604051630739600760e01b81523360048201526001600160a01b0391821660248201529082166044820152606401610553565b505b6007805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a9060200160405180910390a150565b3360009081526001602081905260408220015460ff1661090e576040517fb6c60ea3000000000000000000000000000000000000000000000000000000008152336004820152602401610553565b61091c848443424887610fd3565b949350505050565b6004818154811061093457600080fd5b6000918252602090912001546001600160a01b0316905081565b3360009081526002602052604081206001015460609060ff1661099f576040517f32ea82ab000000000000000000000000000000000000000000000000000000008152336004820152602401610553565b82158015906109b657506001600160a01b0386163b155b156109f8576040517fb5cf5b8f0000000000000000000000000000000000000000000000000000000081526001600160a01b0387166004820152602401610553565b6005805473ffffffffffffffffffffffffffffffffffffffff19811633179091556040516001600160a01b03918216918816908790610a3a908890889061141e565b60006040518083038185875af1925050503d8060008114610a77576040519150601f19603f3d011682016040523d82523d6000602084013e610a7c565b606091505b506005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038581169190911790915560405192955090935088169033907f2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d46690610ae9908a908a908a9061142e565b60405180910390a35094509492505050565b6005546000906001600160a01b03167fffffffffffffffffffffffff000000000000000000000000000000000000000101610b365750600090565b506005546001600160a01b031690565b600054610100900460ff1615808015610b665750600054600160ff909116105b80610b805750303b158015610b80575060005460ff166001145b610c0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610553565b6000805460ff191660011790558015610c2f576000805461ff0019166101001790555b600580546001600160a01b0373ffffffffffffffffffffffffffffffffffffffff199182168117909255600680549091169184169190911790558015610791576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6006546001600160a01b03163314610d7c5760065460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b9160048083019260209291908290030181865afa158015610d0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3291906113b4565b9050336001600160a01b03821614610d7a57600654604051630739600760e01b81523360048201526001600160a01b0391821660248201529082166044820152606401610553565b505b6001600160a01b038216600081815260026020908152604091829020600181015492518515158152909360ff90931692917f49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa910160405180910390a282151581151503610de95750505050565b8215610e8557604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916909117905561078e565b60048054610e95906001906113d1565b81548110610ea557610ea56113f2565b6000918252602090912001548254600480546001600160a01b03909316929091908110610ed457610ed46113f2565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600260006004856000015481548110610f2257610f226113f2565b60009182526020808320909101546001600160a01b031683528201929092526040019020556004805480610f5857610f58611408565b600082815260208082208301600019908101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff1916905550505050565b6008818154811061047e57600080fd5b6003818154811061093457600080fd5b600854604080517fff0000000000000000000000000000000000000000000000000000000000000060f88a901b166020808301919091527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608a901b1660218301527fffffffffffffffff00000000000000000000000000000000000000000000000060c089811b8216603585015288901b16603d830152604582018490526065820186905260858083018690528351808403909101815260a5909201909252805191012060009190600082156110d05760086110b26001856113d1565b815481106110c2576110c26113f2565b906000526020600020015490505b60086110dc828461116c565b8154600181018355600092835260209283902001556040805133815260ff8d16928101929092526001600160a01b038b1682820152606082018790526080820188905267ffffffffffffffff891660a083015251829185917f5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe19181900360c00190a3509098975050505050505050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b6000602082840312156111ad57600080fd5b5035919050565b6001600160a01b03811681146111c957600080fd5b50565b6000602082840312156111de57600080fd5b81356111e9816111b4565b9392505050565b6000806040838503121561120357600080fd5b823561120e816111b4565b91506020830135801515811461122357600080fd5b809150509250929050565b6000806040838503121561124157600080fd5b823561124c816111b4565b946020939093013593505050565b6000806000806080858703121561127057600080fd5b5050823594602084013594506040840135936060013592509050565b6000806000606084860312156112a157600080fd5b833560ff811681146112b257600080fd5b925060208401356112c2816111b4565b929592945050506040919091013590565b600080600080606085870312156112e957600080fd5b84356112f4816111b4565b935060208501359250604085013567ffffffffffffffff8082111561131857600080fd5b818701915087601f83011261132c57600080fd5b81358181111561133b57600080fd5b88602082850101111561134d57600080fd5b95989497505060200194505050565b821515815260006020604081840152835180604085015260005b8181101561139257858101830151858201606001528201611376565b506000606082860101526060601f19601f830116850101925050509392505050565b6000602082840312156113c657600080fd5b81516111e9816111b4565b8181038181111561119557634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b8183823760009101908152919050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea26469706673582212202cf26e3163ea7f3e4ce47312088d88338374d1aabb615e30b835fd852fb7048564736f6c63430008110033",
}

// BridgeTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTesterMetaData.ABI instead.
var BridgeTesterABI = BridgeTesterMetaData.ABI

// BridgeTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTesterMetaData.Bin instead.
var BridgeTesterBin = BridgeTesterMetaData.Bin

// DeployBridgeTester deploys a new Ethereum contract, binding an instance of BridgeTester to it.
func DeployBridgeTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeTester, error) {
	parsed, err := BridgeTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTester{BridgeTesterCaller: BridgeTesterCaller{contract: contract}, BridgeTesterTransactor: BridgeTesterTransactor{contract: contract}, BridgeTesterFilterer: BridgeTesterFilterer{contract: contract}}, nil
}

// BridgeTester is an auto generated Go binding around an Ethereum contract.
type BridgeTester struct {
	BridgeTesterCaller     // Read-only binding to the contract
	BridgeTesterTransactor // Write-only binding to the contract
	BridgeTesterFilterer   // Log filterer for contract events
}

// BridgeTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTesterSession struct {
	Contract     *BridgeTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTesterCallerSession struct {
	Contract *BridgeTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTesterTransactorSession struct {
	Contract     *BridgeTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTesterRaw struct {
	Contract *BridgeTester // Generic contract binding to access the raw methods on
}

// BridgeTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTesterCallerRaw struct {
	Contract *BridgeTesterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTesterTransactorRaw struct {
	Contract *BridgeTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTester creates a new instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTester(address common.Address, backend bind.ContractBackend) (*BridgeTester, error) {
	contract, err := bindBridgeTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTester{BridgeTesterCaller: BridgeTesterCaller{contract: contract}, BridgeTesterTransactor: BridgeTesterTransactor{contract: contract}, BridgeTesterFilterer: BridgeTesterFilterer{contract: contract}}, nil
}

// NewBridgeTesterCaller creates a new read-only instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterCaller(address common.Address, caller bind.ContractCaller) (*BridgeTesterCaller, error) {
	contract, err := bindBridgeTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterCaller{contract: contract}, nil
}

// NewBridgeTesterTransactor creates a new write-only instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTesterTransactor, error) {
	contract, err := bindBridgeTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterTransactor{contract: contract}, nil
}

// NewBridgeTesterFilterer creates a new log filterer instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTesterFilterer, error) {
	contract, err := bindBridgeTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterFilterer{contract: contract}, nil
}

// bindBridgeTester binds a generic wrapper to an already deployed contract.
func bindBridgeTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTester *BridgeTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTester.Contract.BridgeTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTester *BridgeTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.Contract.BridgeTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTester *BridgeTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTester.Contract.BridgeTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTester *BridgeTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTester *BridgeTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTester *BridgeTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTester.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterSession) ActiveOutbox() (common.Address, error) {
	return _BridgeTester.Contract.ActiveOutbox(&_BridgeTester.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeTester.Contract.ActiveOutbox(&_BridgeTester.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxes(&_BridgeTester.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxes(&_BridgeTester.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedOutboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedOutboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedOutboxes(&_BridgeTester.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedOutboxes(&_BridgeTester.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.DelayedInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.DelayedInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.DelayedMessageCount(&_BridgeTester.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.DelayedMessageCount(&_BridgeTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterSession) Rollup() (common.Address, error) {
	return _BridgeTester.Contract.Rollup(&_BridgeTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) Rollup() (common.Address, error) {
	return _BridgeTester.Contract.Rollup(&_BridgeTester.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterSession) SequencerInbox() (common.Address, error) {
	return _BridgeTester.Contract.SequencerInbox(&_BridgeTester.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeTester.Contract.SequencerInbox(&_BridgeTester.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.SequencerInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.SequencerInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerMessageCount(&_BridgeTester.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerMessageCount(&_BridgeTester.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerReportedSubMessageCount(&_BridgeTester.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerReportedSubMessageCount(&_BridgeTester.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeTester.Contract.AcceptFundsFromOldBridge(&_BridgeTester.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeTester.Contract.AcceptFundsFromOldBridge(&_BridgeTester.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueDelayedMessage(&_BridgeTester.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueDelayedMessage(&_BridgeTester.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueSequencerMessage(&_BridgeTester.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueSequencerMessage(&_BridgeTester.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.ExecuteCall(&_BridgeTester.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.ExecuteCall(&_BridgeTester.TransactOpts, to, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterTransactor) Initialize(opts *bind.TransactOpts, rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "initialize", rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.Initialize(&_BridgeTester.TransactOpts, rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterTransactorSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.Initialize(&_BridgeTester.TransactOpts, rollup_)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetDelayedInbox(&_BridgeTester.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetDelayedInbox(&_BridgeTester.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetOutbox(&_BridgeTester.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetOutbox(&_BridgeTester.TransactOpts, outbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetSequencerInbox(&_BridgeTester.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetSequencerInbox(&_BridgeTester.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.SubmitBatchSpendingReport(&_BridgeTester.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.SubmitBatchSpendingReport(&_BridgeTester.TransactOpts, batchPoster, dataHash)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterSession) Receive() (*types.Transaction, error) {
	return _BridgeTester.Contract.Receive(&_BridgeTester.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeTester.Contract.Receive(&_BridgeTester.TransactOpts)
}

// BridgeTesterBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeTester contract.
type BridgeTesterBridgeCallTriggeredIterator struct {
	Event *BridgeTesterBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeTesterBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterBridgeCallTriggered)
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
		it.Event = new(BridgeTesterBridgeCallTriggered)
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
func (it *BridgeTesterBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeTester contract.
type BridgeTesterBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeTesterBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterBridgeCallTriggeredIterator{contract: _BridgeTester.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeTesterBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterBridgeCallTriggered)
				if err := _BridgeTester.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeTesterBridgeCallTriggered, error) {
	event := new(BridgeTesterBridgeCallTriggered)
	if err := _BridgeTester.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeTester contract.
type BridgeTesterInboxToggleIterator struct {
	Event *BridgeTesterInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeTesterInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterInboxToggle)
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
		it.Event = new(BridgeTesterInboxToggle)
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
func (it *BridgeTesterInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterInboxToggle represents a InboxToggle event raised by the BridgeTester contract.
type BridgeTesterInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeTesterInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterInboxToggleIterator{contract: _BridgeTester.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeTesterInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterInboxToggle)
				if err := _BridgeTester.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) ParseInboxToggle(log types.Log) (*BridgeTesterInboxToggle, error) {
	event := new(BridgeTesterInboxToggle)
	if err := _BridgeTester.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeTester contract.
type BridgeTesterInitializedIterator struct {
	Event *BridgeTesterInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeTesterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterInitialized)
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
		it.Event = new(BridgeTesterInitialized)
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
func (it *BridgeTesterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterInitialized represents a Initialized event raised by the BridgeTester contract.
type BridgeTesterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTester *BridgeTesterFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeTesterInitializedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterInitializedIterator{contract: _BridgeTester.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTester *BridgeTesterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeTesterInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterInitialized)
				if err := _BridgeTester.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeTester *BridgeTesterFilterer) ParseInitialized(log types.Log) (*BridgeTesterInitialized, error) {
	event := new(BridgeTesterInitialized)
	if err := _BridgeTester.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeTester contract.
type BridgeTesterMessageDeliveredIterator struct {
	Event *BridgeTesterMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeTesterMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterMessageDelivered)
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
		it.Event = new(BridgeTesterMessageDelivered)
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
func (it *BridgeTesterMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterMessageDelivered represents a MessageDelivered event raised by the BridgeTester contract.
type BridgeTesterMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeTesterMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterMessageDeliveredIterator{contract: _BridgeTester.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeTesterMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterMessageDelivered)
				if err := _BridgeTester.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) ParseMessageDelivered(log types.Log) (*BridgeTesterMessageDelivered, error) {
	event := new(BridgeTesterMessageDelivered)
	if err := _BridgeTester.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeTester contract.
type BridgeTesterOutboxToggleIterator struct {
	Event *BridgeTesterOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeTesterOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterOutboxToggle)
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
		it.Event = new(BridgeTesterOutboxToggle)
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
func (it *BridgeTesterOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterOutboxToggle represents a OutboxToggle event raised by the BridgeTester contract.
type BridgeTesterOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeTesterOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterOutboxToggleIterator{contract: _BridgeTester.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeTesterOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterOutboxToggle)
				if err := _BridgeTester.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) ParseOutboxToggle(log types.Log) (*BridgeTesterOutboxToggle, error) {
	event := new(BridgeTesterOutboxToggle)
	if err := _BridgeTester.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeTester contract.
type BridgeTesterSequencerInboxUpdatedIterator struct {
	Event *BridgeTesterSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTesterSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterSequencerInboxUpdated)
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
		it.Event = new(BridgeTesterSequencerInboxUpdated)
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
func (it *BridgeTesterSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeTester contract.
type BridgeTesterSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeTesterSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterSequencerInboxUpdatedIterator{contract: _BridgeTester.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTesterSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterSequencerInboxUpdated)
				if err := _BridgeTester.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeTesterSequencerInboxUpdated, error) {
	event := new(BridgeTesterSequencerInboxUpdated)
	if err := _BridgeTester.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptographyPrimitivesTesterMetaData contains all meta data concerning the CryptographyPrimitivesTester contract.
var CryptographyPrimitivesTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"input\",\"type\":\"uint256[25]\"}],\"name\":\"keccakF\",\"outputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"\",\"type\":\"uint256[25]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"inputChunk\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"hashState\",\"type\":\"bytes32\"}],\"name\":\"sha256Block\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61173b61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063ac90ed4614610045578063e479f5321461006e575b600080fd5b610058610053366004611555565b61008f565b60405161006591906115e2565b60405180910390f35b61008161007c366004611614565b6100a6565b604051908152602001610065565b6100976114da565b6100a0826100e9565b92915050565b60006100e26040518060400160405280856000600281106100c9576100c96116a7565b6020908102919091015182528681015191015283610d35565b9392505050565b6100f16114da565b6100f96114f9565b6101016114f9565b6101096114da565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015610d2a576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e01511818181892880183905267ffffffffffffffff6002820216678000000000000000918290041790921886525104856002602002015160020267ffffffffffffffff16178560006020020151188460016020020152678000000000000000856003602002015181610364576103646116bd565b04856003602002015160020267ffffffffffffffff161785600160200201511884600260200201526780000000000000008560046020020151816103aa576103aa6116bd565b04856004602002015160020267ffffffffffffffff1617856002600581106103d4576103d46116a7565b6020020151186060850152845167800000000000000090865160608089015193909204600290910267ffffffffffffffff1617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a018051909118905290845251631000000090602089015167ffffffffffffffff6410000000009091021691900417610100840152604087015167200000000000000090604089015167ffffffffffffffff6008909102169190041761016084015260608701516280000090606089015167ffffffffffffffff65020000000000909102169190041761026084015260808701516540000000000090608089015167ffffffffffffffff6204000090910216919004176102c084015260a08701516780000000000000009004876005602002015160020267ffffffffffffffff161783600260198110610656576106566116a7565b602002015260c08701516210000081046510000000000090910267ffffffffffffffff9081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166740000000000000009091041761030085015261014088015161014089015167ffffffffffffffff674000000000000000909102169190041760808401526101608701516704000000000000009061016089015167ffffffffffffffff6040909102169190041760e0840152610180870151622000009061018089015167ffffffffffffffff6508000000000090910216919004176101408401526101a08701516602000000000000906101a089015167ffffffffffffffff61800090910216919004176102408401526101c08701516008906101c089015167ffffffffffffffff67200000000000000090910216919004176102a08401526101e0870151641000000000906101e089015167ffffffffffffffff6310000000909102169190041760208401526102008088015161020089015167ffffffffffffffff668000000000000090910216919004176101208401526102208701516480000000009061022089015167ffffffffffffffff63020000009091021691900417610180840152610240870151650800000000009061024089015167ffffffffffffffff6220000090910216919004176101e08401526102608701516101009061026089015167ffffffffffffffff67010000000000000090910216919004176102e08401526102808701516420000000009061028089015167ffffffffffffffff6308000000909102169190041760608401526102a087015165100000000000906102a089015167ffffffffffffffff62100000909102169190041760c08401526102c08701516302000000906102c089015167ffffffffffffffff64800000000090910216919004176101c08401526102e0870151670100000000000000906102e089015167ffffffffffffffff61010090910216919004176102208401526103008701516604000000000000900487601860200201516140000267ffffffffffffffff1617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f6020020152826001602002015183601560200201511916836010602002015118876010602002015282600260200201518360166020020151191683601160200201511887601160200201528260036020020151836017602002015119168360126020020151188760126020020152826004602002015183601860200201511916836013602002015118876013602002015282600560200201518360006020020151191683601460200201511887601460200201528260066020020151836001602002015119168360156020020151188760156020020152826007602002015183600260200201511916836016602002015118876016602002015282600860200201518360036020020151191683601760200201511887601760200201528260096020020151836004602002015119168360186020020151188760186020020152818160188110610d1857610d186116a7565b6020020151875118875260010161022f565b509495945050505050565b604080516108008101825263428a2f9881526371374491602082015263b5c0fbcf9181019190915263e9b5dba56060820152633956c25b60808201526359f111f160a082015263923f82a460c082015263ab1c5ed560e082015263d807aa986101008201526312835b0161012082015263243185be61014082015263550c7dc36101608201526372be5d746101808201526380deb1fe6101a0820152639bdc06a76101c082015263c19bf1746101e082015263e49b69c161020082015263efbe4786610220820152630fc19dc661024082015263240ca1cc610260820152632de92c6f610280820152634a7484aa6102a0820152635cb0a9dc6102c08201526376f988da6102e082015263983e515261030082015263a831c66d61032082015263b00327c861034082015263bf597fc761036082015263c6e00bf361038082015263d5a791476103a08201526306ca63516103c082015263142929676103e08201526327b70a85610400820152632e1b2138610420820152634d2c6dfc6104408201526353380d1361046082015263650a735461048082015263766a0abb6104a08201526381c2c92e6104c08201526392722c856104e082015263a2bfe8a161050082015263a81a664b61052082015263c24b8b7061054082015263c76c51a361056082015263d192e81961058082015263d69906246105a082015263f40e35856105c082015263106aa0706105e08201526319a4c116610600820152631e376c08610620820152632748774c6106408201526334b0bcb561066082015263391c0cb3610680820152634ed8aa4a6106a0820152635b9cca4f6106c082015263682e6ff36106e082015263748f82ee6107008201526378a5636f6107208201526384c87814610740820152638cc702086107608201526390befffa61078082015263a4506ceb6107a082015263bef9a3f76107c082015263c67178f26107e0820152600090611002611517565b60005b60088163ffffffff16101561109b5763ffffffff6020820260e003168660006020020151901c828263ffffffff1660408110611043576110436116a7565b63ffffffff92831660209182029290920191909152820260e003168660016020020151901c828260080163ffffffff1660408110611083576110836116a7565b63ffffffff9092166020929092020152600101611005565b5060106000805b60408363ffffffff16101561122d57600384600f850363ffffffff16604081106110ce576110ce6116a7565b602002015163ffffffff16901c61110585600f860363ffffffff16604081106110f9576110f96116a7565b602002015160126114a5565b61112f86600f870363ffffffff1660408110611123576111236116a7565b602002015160076114a5565b18189150600a846002850363ffffffff1660408110611150576111506116a7565b602002015163ffffffff16901c611187856002860363ffffffff166040811061117b5761117b6116a7565b602002015160136114a5565b6111b1866002870363ffffffff16604081106111a5576111a56116a7565b602002015160116114a5565b1818905080846007850363ffffffff16604081106111d1576111d16116a7565b602002015183866010870363ffffffff16604081106111f2576111f26116a7565b6020020151010101848463ffffffff1660408110611212576112126116a7565b63ffffffff90921660209290920201526001909201916110a2565b611235611536565b600093505b60088463ffffffff16101561128c578360200260e00363ffffffff1688901c818563ffffffff1660088110611271576112716116a7565b63ffffffff909216602092909202015260019093019261123a565b60008060008096505b60408763ffffffff1610156113e85760808401516112b49060196114a5565b60808501516112c490600b6114a5565b60808601516112d49060066114a5565b18189450878763ffffffff16604081106112f0576112f06116a7565b6020020151898863ffffffff166040811061130d5761130d6116a7565b6020020151608086015160a087015160c0880151821916911618878760076020020151010101019250611348846000602002015160166114a5565b845161135590600d6114a5565b85516113629060026114a5565b6040870180516020890180518a5160c08c01805163ffffffff90811660e08f015260a08e018051821690925260808e018051821690925260608e0180518e01821690925280861690915280831690955284811690925280831891909116911618929091189290921881810186810190931687526001999099019897509092509050611295565b600096505b60088763ffffffff161015611442578660200260e00363ffffffff168b901c848863ffffffff1660088110611424576114246116a7565b60200201805163ffffffff92019190911690526001909601956113ed565b60008097505b60088863ffffffff161015611495578760200260e00363ffffffff16858963ffffffff166008811061147c5761147c6116a7565b602002015160019099019863ffffffff16901b17611448565b9c9b505050505050505050505050565b60006114b28260206116d3565b63ffffffff168363ffffffff16901b8263ffffffff168463ffffffff16901c17905092915050565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b6040518061080001604052806040906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b600061032080838503121561156957600080fd5b83601f84011261157857600080fd5b60405181810181811067ffffffffffffffff821117156115a857634e487b7160e01b600052604160045260246000fd5b6040529083019080858311156115bd57600080fd5b845b838110156115d75780358252602091820191016115bf565b509095945050505050565b6103208101818360005b601981101561160b5781518352602092830192909101906001016115ec565b50505092915050565b6000806060838503121561162757600080fd5b83601f84011261163657600080fd5b6040516040810181811067ffffffffffffffff8211171561166757634e487b7160e01b600052604160045260246000fd5b806040525080604085018681111561167e57600080fd5b855b81811015611698578035835260209283019201611680565b50919691359550909350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601260045260246000fd5b63ffffffff8281168282160390808211156116fe57634e487b7160e01b600052601160045260246000fd5b509291505056fea2646970667358221220e85f8f1ad8770e74b68e54d7ae7cd46ac4d12065a1497c2f907303e52876d55464736f6c63430008110033",
}

// CryptographyPrimitivesTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use CryptographyPrimitivesTesterMetaData.ABI instead.
var CryptographyPrimitivesTesterABI = CryptographyPrimitivesTesterMetaData.ABI

// CryptographyPrimitivesTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CryptographyPrimitivesTesterMetaData.Bin instead.
var CryptographyPrimitivesTesterBin = CryptographyPrimitivesTesterMetaData.Bin

// DeployCryptographyPrimitivesTester deploys a new Ethereum contract, binding an instance of CryptographyPrimitivesTester to it.
func DeployCryptographyPrimitivesTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptographyPrimitivesTester, error) {
	parsed, err := CryptographyPrimitivesTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CryptographyPrimitivesTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptographyPrimitivesTester{CryptographyPrimitivesTesterCaller: CryptographyPrimitivesTesterCaller{contract: contract}, CryptographyPrimitivesTesterTransactor: CryptographyPrimitivesTesterTransactor{contract: contract}, CryptographyPrimitivesTesterFilterer: CryptographyPrimitivesTesterFilterer{contract: contract}}, nil
}

// CryptographyPrimitivesTester is an auto generated Go binding around an Ethereum contract.
type CryptographyPrimitivesTester struct {
	CryptographyPrimitivesTesterCaller     // Read-only binding to the contract
	CryptographyPrimitivesTesterTransactor // Write-only binding to the contract
	CryptographyPrimitivesTesterFilterer   // Log filterer for contract events
}

// CryptographyPrimitivesTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptographyPrimitivesTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptographyPrimitivesTesterSession struct {
	Contract     *CryptographyPrimitivesTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CryptographyPrimitivesTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptographyPrimitivesTesterCallerSession struct {
	Contract *CryptographyPrimitivesTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CryptographyPrimitivesTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptographyPrimitivesTesterTransactorSession struct {
	Contract     *CryptographyPrimitivesTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CryptographyPrimitivesTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterRaw struct {
	Contract *CryptographyPrimitivesTester // Generic contract binding to access the raw methods on
}

// CryptographyPrimitivesTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterCallerRaw struct {
	Contract *CryptographyPrimitivesTesterCaller // Generic read-only contract binding to access the raw methods on
}

// CryptographyPrimitivesTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterTransactorRaw struct {
	Contract *CryptographyPrimitivesTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptographyPrimitivesTester creates a new instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTester(address common.Address, backend bind.ContractBackend) (*CryptographyPrimitivesTester, error) {
	contract, err := bindCryptographyPrimitivesTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTester{CryptographyPrimitivesTesterCaller: CryptographyPrimitivesTesterCaller{contract: contract}, CryptographyPrimitivesTesterTransactor: CryptographyPrimitivesTesterTransactor{contract: contract}, CryptographyPrimitivesTesterFilterer: CryptographyPrimitivesTesterFilterer{contract: contract}}, nil
}

// NewCryptographyPrimitivesTesterCaller creates a new read-only instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterCaller(address common.Address, caller bind.ContractCaller) (*CryptographyPrimitivesTesterCaller, error) {
	contract, err := bindCryptographyPrimitivesTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterCaller{contract: contract}, nil
}

// NewCryptographyPrimitivesTesterTransactor creates a new write-only instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptographyPrimitivesTesterTransactor, error) {
	contract, err := bindCryptographyPrimitivesTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterTransactor{contract: contract}, nil
}

// NewCryptographyPrimitivesTesterFilterer creates a new log filterer instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptographyPrimitivesTesterFilterer, error) {
	contract, err := bindCryptographyPrimitivesTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterFilterer{contract: contract}, nil
}

// bindCryptographyPrimitivesTester binds a generic wrapper to an already deployed contract.
func bindCryptographyPrimitivesTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CryptographyPrimitivesTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptographyPrimitivesTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.contract.Transact(opts, method, params...)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCaller) KeccakF(opts *bind.CallOpts, input [25]*big.Int) ([25]*big.Int, error) {
	var out []interface{}
	err := _CryptographyPrimitivesTester.contract.Call(opts, &out, "keccakF", input)

	if err != nil {
		return *new([25]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([25]*big.Int)).(*[25]*big.Int)

	return out0, err

}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _CryptographyPrimitivesTester.Contract.KeccakF(&_CryptographyPrimitivesTester.CallOpts, input)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _CryptographyPrimitivesTester.Contract.KeccakF(&_CryptographyPrimitivesTester.CallOpts, input)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCaller) Sha256Block(opts *bind.CallOpts, inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CryptographyPrimitivesTester.contract.Call(opts, &out, "sha256Block", inputChunk, hashState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _CryptographyPrimitivesTester.Contract.Sha256Block(&_CryptographyPrimitivesTester.CallOpts, inputChunk, hashState)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _CryptographyPrimitivesTester.Contract.Sha256Block(&_CryptographyPrimitivesTester.CallOpts, inputChunk, hashState)
}

// MessageTesterMetaData contains all meta data concerning the MessageTester contract.
var MessageTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"accumulateInboxMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610267806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638f3c79c01461003b578063bf00905214610087575b600080fd5b610075610049366004610160565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b60405190815260200160405180910390f35b61007561009536600461019f565b6040805160f89890981b7fff00000000000000000000000000000000000000000000000000000000000000166020808a019190915260609790971b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602189015260c095861b7fffffffffffffffff00000000000000000000000000000000000000000000000090811660358a01529490951b909316603d870152604586019190915260658501526085808501919091528151808503909101815260a59093019052815191012090565b6000806040838503121561017357600080fd5b50508035926020909101359150565b803567ffffffffffffffff8116811461019a57600080fd5b919050565b600080600080600080600060e0888a0312156101ba57600080fd5b873560ff811681146101cb57600080fd5b9650602088013573ffffffffffffffffffffffffffffffffffffffff811681146101f457600080fd5b955061020260408901610182565b945061021060608901610182565b9699959850939660808101359560a0820135955060c090910135935091505056fea26469706673582212201d0e022d2f84790bf2a803422ec0c01010eb3da921305b374b53ba5ddcc5ca7b64736f6c63430008110033",
}

// MessageTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageTesterMetaData.ABI instead.
var MessageTesterABI = MessageTesterMetaData.ABI

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageTesterMetaData.Bin instead.
var MessageTesterBin = MessageTesterMetaData.Bin

// DeployMessageTester deploys a new Ethereum contract, binding an instance of MessageTester to it.
func DeployMessageTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageTester, error) {
	parsed, err := MessageTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// MessageTester is an auto generated Go binding around an Ethereum contract.
type MessageTester struct {
	MessageTesterCaller     // Read-only binding to the contract
	MessageTesterTransactor // Write-only binding to the contract
	MessageTesterFilterer   // Log filterer for contract events
}

// MessageTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTesterSession struct {
	Contract     *MessageTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTesterCallerSession struct {
	Contract *MessageTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessageTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTesterTransactorSession struct {
	Contract     *MessageTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessageTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTesterRaw struct {
	Contract *MessageTester // Generic contract binding to access the raw methods on
}

// MessageTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTesterCallerRaw struct {
	Contract *MessageTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTesterTransactorRaw struct {
	Contract *MessageTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTester creates a new instance of MessageTester, bound to a specific deployed contract.
func NewMessageTester(address common.Address, backend bind.ContractBackend) (*MessageTester, error) {
	contract, err := bindMessageTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// NewMessageTesterCaller creates a new read-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterCaller(address common.Address, caller bind.ContractCaller) (*MessageTesterCaller, error) {
	contract, err := bindMessageTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterCaller{contract: contract}, nil
}

// NewMessageTesterTransactor creates a new write-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTesterTransactor, error) {
	contract, err := bindMessageTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterTransactor{contract: contract}, nil
}

// NewMessageTesterFilterer creates a new log filterer instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTesterFilterer, error) {
	contract, err := bindMessageTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTesterFilterer{contract: contract}, nil
}

// bindMessageTester binds a generic wrapper to an already deployed contract.
func bindMessageTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.MessageTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transact(opts, method, params...)
}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) AccumulateInboxMessage(opts *bind.CallOpts, inbox [32]byte, message [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "accumulateInboxMessage", inbox, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) AccumulateInboxMessage(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AccumulateInboxMessage(&_MessageTester.CallOpts, inbox, message)
}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AccumulateInboxMessage(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AccumulateInboxMessage(&_MessageTester.CallOpts, inbox, message)
}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) MessageHash(opts *bind.CallOpts, messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "messageHash", messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) MessageHash(messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)
}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) MessageHash(messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)
}

// OutboxWithoutOptTesterMetaData contains all meta data concerning the OutboxWithoutOptTester contract.
var OutboxWithoutOptTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlreadySpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofLength\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"name\":\"PathNotMinimal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proofLength\",\"type\":\"uint256\"}],\"name\":\"ProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"UnknownRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"zero\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"SendRootUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OUTBOX_VERSION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeTransactionSimulation\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1BatchNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1OutputId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"updateSendRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161131161002d600039600050506113116000f3fe608060405234801561001057600080fd5b50600436106101355760003560e01c80639f0c04bf116100b2578063c4d66de811610081578063cb23bcb511610066578063cb23bcb5146102af578063d5b5cc23146102c2578063e78cea92146102d557600080fd5b8063c4d66de81461027c578063c75184df1461028f57600080fd5b80639f0c04bf14610225578063a04cee6014610238578063ae6dead71461024b578063b0f305371461026b57600080fd5b8063465477901161010957806372f2a8c7116100ee57806372f2a8c7146101d357806380648b02146101db5780638515bc6a1461020057600080fd5b8063465477901461018f5780635a129efe146101a057600080fd5b80627436d31461013a57806308635a95146101605780631198527114610175578063288e5b101461017c575b600080fd5b61014d610148366004610c71565b6102e8565b6040519081526020015b60405180910390f35b61017361016e366004610d98565b610325565b005b600061014d565b61017361018a366004610e8d565b610699565b6004546001600160801b031661014d565b6101c36101ae366004610f29565b60026020526000908152604090205460ff1681565b6040519015158152602001610157565b60065461014d565b6007546001600160a01b03165b6040516001600160a01b039091168152602001610157565b60045470010000000000000000000000000000000090046001600160801b031661014d565b61014d610233366004610f42565b610700565b610173610246366004610fd1565b610745565b61014d610259366004610f29565b60036020526000908152604090205481565b6005546001600160801b031661014d565b61017361028a366004610ff3565b610784565b610297600281565b6040516001600160801b039091168152602001610157565b6000546101e8906001600160a01b031681565b61014d6102d0366004610f29565b61089a565b6001546101e8906001600160a01b031681565b600061031d84848460405160200161030291815260200190565b604051602081830303815290604052805190602001206108ff565b949350505050565b6000806103388a8a8a8a8a8a8a8a610700565b905061037a8d8d808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508f92508591506109ba9050565b915060008a6001600160a01b03168a6001600160a01b03167f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab189648e6040516103c391815260200190565b60405180910390a450600060046040518060a00160405290816000820160009054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016000820160109054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016001820160009054906101000a90046001600160801b03166001600160801b03166001600160801b03168152602001600282015481526020016003820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152505090506040518060a00160405280896001600160801b03168152602001886001600160801b03168152602001876001600160801b031681526020018381526020018b6001600160a01b0316815250600460008201518160000160006101000a8154816001600160801b0302191690836001600160801b0316021790555060208201518160000160106101000a8154816001600160801b0302191690836001600160801b0316021790555060408201518160010160006101000a8154816001600160801b0302191690836001600160801b031602179055506060820151816002015560808201518160030160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506105f2898686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b3f92505050565b805160208201516001600160801b0391821670010000000000000000000000000000000091831691909102176004556040820151600580547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016919092161790556060810151600655608001516007805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055505050505050505050505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4e6f7420696d706c656d656e746564000000000000000000000000000000000060448201526064015b60405180910390fd5b60008888888888888888604051602001610721989796959493929190611017565b60405160208183030381529060405280519060200120905098975050505050505050565b60008281526003602052604080822083905551829184917fb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f67489190a35050565b6001546001600160a01b0316156107c7576040517fef34ca5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038316908117909155604080517fcb23bcb5000000000000000000000000000000000000000000000000000000008152905163cb23bcb5916004808201926020929091908290030181865afa158015610846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086a9190611083565b6000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905550565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f494d504c454d4554454400000000000000000000000000000000000060448201526000906064016106f7565b825160009061010081111561094b576040517ffdac331e0000000000000000000000000000000000000000000000000000000081526004810182905261010060248201526044016106f7565b8260005b828110156109b057600087828151811061096b5761096b6110a0565b60200260200101519050816001901b8716600003610997578260005280602052604060002092506109a7565b8060005282602052604060002092505b5060010161094f565b5095945050505050565b60006101008451106109fd5783516040517fab6a06830000000000000000000000000000000000000000000000000000000081526004016106f791815260200190565b8351610a0a9060026111b2565b8310610a5a578284516002610a1f91906111b2565b6040517f0b8a724b000000000000000000000000000000000000000000000000000000008152600481019290925260248201526044016106f7565b6000610a678585856102e8565b600081815260036020526040902054909150610ab2576040517f8730d7c8000000000000000000000000000000000000000000000000000000008152600481018290526024016106f7565b60008481526002602052604090205460ff1615610afe576040517f9715b8d3000000000000000000000000000000000000000000000000000000008152600481018590526024016106f7565b5050600082815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055819392505050565b6001546040517f9e5d4c4900000000000000000000000000000000000000000000000000000000815260009182916001600160a01b0390911690639e5d4c4990610b91908890889088906004016111e2565b6000604051808303816000875af1158015610bb0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610bd8919081019061122c565b9150915081610c2357805115610bf15780518082602001fd5b6040517f376fb55a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610c6957610c69610c2a565b604052919050565b600080600060608486031215610c8657600080fd5b833567ffffffffffffffff80821115610c9e57600080fd5b818601915086601f830112610cb257600080fd5b8135602082821115610cc657610cc6610c2a565b8160051b9250610cd7818401610c40565b828152928401810192818101908a851115610cf157600080fd5b948201945b84861015610d0f57853582529482019490820190610cf6565b9a918901359950506040909701359695505050505050565b6001600160a01b0381168114610d3c57600080fd5b50565b8035610d4a81610d27565b919050565b60008083601f840112610d6157600080fd5b50813567ffffffffffffffff811115610d7957600080fd5b602083019150836020828501011115610d9157600080fd5b9250929050565b60008060008060008060008060008060006101208c8e031215610dba57600080fd5b8b3567ffffffffffffffff80821115610dd257600080fd5b818e0191508e601f830112610de657600080fd5b813581811115610df557600080fd5b8f60208260051b8501011115610e0a57600080fd5b60208381019e50909c508e01359a50610e2560408f01610d3f565b9950610e3360608f01610d3f565b985060808e0135975060a08e0135965060c08e0135955060e08e013594506101008e0135915080821115610e6657600080fd5b50610e738e828f01610d4f565b915080935050809150509295989b509295989b9093969950565b60008060008060008060008060006101008a8c031215610eac57600080fd5b8935985060208a0135610ebe81610d27565b975060408a0135610ece81610d27565b965060608a0135955060808a0135945060a08a0135935060c08a0135925060e08a013567ffffffffffffffff811115610f0657600080fd5b610f128c828d01610d4f565b915080935050809150509295985092959850929598565b600060208284031215610f3b57600080fd5b5035919050565b60008060008060008060008060e0898b031215610f5e57600080fd5b8835610f6981610d27565b97506020890135610f7981610d27565b965060408901359550606089013594506080890135935060a0890135925060c089013567ffffffffffffffff811115610fb157600080fd5b610fbd8b828c01610d4f565b999c989b5096995094979396929594505050565b60008060408385031215610fe457600080fd5b50508035926020909101359150565b60006020828403121561100557600080fd5b813561101081610d27565b9392505050565b60007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000808b60601b168352808a60601b16601484015250876028830152866048830152856068830152846088830152828460a8840137506000910160a801908152979650505050505050565b60006020828403121561109557600080fd5b815161101081610d27565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600181815b808511156111075781600019048211156110ed576110ed6110b6565b808516156110fa57918102915b93841c93908002906110d1565b509250929050565b60008261111e575060016111ac565b8161112b575060006111ac565b8160018114611141576002811461114b57611167565b60019150506111ac565b60ff84111561115c5761115c6110b6565b50506001821b6111ac565b5060208310610133831016604e8410600b841016171561118a575081810a6111ac565b61119483836110cc565b80600019048211156111a8576111a86110b6565b0290505b92915050565b6000611010838361110f565b60005b838110156111d95781810151838201526020016111c1565b50506000910152565b6001600160a01b038416815282602082015260606040820152600082518060608401526112168160808501602087016111be565b601f01601f191691909101608001949350505050565b6000806040838503121561123f57600080fd5b8251801515811461124f57600080fd5b602084015190925067ffffffffffffffff8082111561126d57600080fd5b818501915085601f83011261128157600080fd5b81518181111561129357611293610c2a565b6112a66020601f19601f84011601610c40565b91508082528660208285010111156112bd57600080fd5b6112ce8160208401602086016111be565b508092505050925092905056fea2646970667358221220dff53adaa809b0a95c4259a18be6ff2e18eeba627f70bad96d93adf2b907aa9e64736f6c63430008110033",
}

// OutboxWithoutOptTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use OutboxWithoutOptTesterMetaData.ABI instead.
var OutboxWithoutOptTesterABI = OutboxWithoutOptTesterMetaData.ABI

// OutboxWithoutOptTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutboxWithoutOptTesterMetaData.Bin instead.
var OutboxWithoutOptTesterBin = OutboxWithoutOptTesterMetaData.Bin

// DeployOutboxWithoutOptTester deploys a new Ethereum contract, binding an instance of OutboxWithoutOptTester to it.
func DeployOutboxWithoutOptTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxWithoutOptTester, error) {
	parsed, err := OutboxWithoutOptTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutboxWithoutOptTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxWithoutOptTester{OutboxWithoutOptTesterCaller: OutboxWithoutOptTesterCaller{contract: contract}, OutboxWithoutOptTesterTransactor: OutboxWithoutOptTesterTransactor{contract: contract}, OutboxWithoutOptTesterFilterer: OutboxWithoutOptTesterFilterer{contract: contract}}, nil
}

// OutboxWithoutOptTester is an auto generated Go binding around an Ethereum contract.
type OutboxWithoutOptTester struct {
	OutboxWithoutOptTesterCaller     // Read-only binding to the contract
	OutboxWithoutOptTesterTransactor // Write-only binding to the contract
	OutboxWithoutOptTesterFilterer   // Log filterer for contract events
}

// OutboxWithoutOptTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxWithoutOptTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxWithoutOptTesterSession struct {
	Contract     *OutboxWithoutOptTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OutboxWithoutOptTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxWithoutOptTesterCallerSession struct {
	Contract *OutboxWithoutOptTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// OutboxWithoutOptTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxWithoutOptTesterTransactorSession struct {
	Contract     *OutboxWithoutOptTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// OutboxWithoutOptTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxWithoutOptTesterRaw struct {
	Contract *OutboxWithoutOptTester // Generic contract binding to access the raw methods on
}

// OutboxWithoutOptTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterCallerRaw struct {
	Contract *OutboxWithoutOptTesterCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxWithoutOptTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterTransactorRaw struct {
	Contract *OutboxWithoutOptTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxWithoutOptTester creates a new instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTester(address common.Address, backend bind.ContractBackend) (*OutboxWithoutOptTester, error) {
	contract, err := bindOutboxWithoutOptTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTester{OutboxWithoutOptTesterCaller: OutboxWithoutOptTesterCaller{contract: contract}, OutboxWithoutOptTesterTransactor: OutboxWithoutOptTesterTransactor{contract: contract}, OutboxWithoutOptTesterFilterer: OutboxWithoutOptTesterFilterer{contract: contract}}, nil
}

// NewOutboxWithoutOptTesterCaller creates a new read-only instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterCaller(address common.Address, caller bind.ContractCaller) (*OutboxWithoutOptTesterCaller, error) {
	contract, err := bindOutboxWithoutOptTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterCaller{contract: contract}, nil
}

// NewOutboxWithoutOptTesterTransactor creates a new write-only instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxWithoutOptTesterTransactor, error) {
	contract, err := bindOutboxWithoutOptTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterTransactor{contract: contract}, nil
}

// NewOutboxWithoutOptTesterFilterer creates a new log filterer instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxWithoutOptTesterFilterer, error) {
	contract, err := bindOutboxWithoutOptTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterFilterer{contract: contract}, nil
}

// bindOutboxWithoutOptTester binds a generic wrapper to an already deployed contract.
func bindOutboxWithoutOptTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OutboxWithoutOptTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxWithoutOptTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.contract.Transact(opts, method, params...)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) OUTBOXVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "OUTBOX_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) OUTBOXVERSION() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.OUTBOXVERSION(&_OutboxWithoutOptTester.CallOpts)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) OUTBOXVERSION() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.OUTBOXVERSION(&_OutboxWithoutOptTester.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Bridge() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Bridge(&_OutboxWithoutOptTester.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Bridge() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Bridge(&_OutboxWithoutOptTester.CallOpts)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "calculateItemHash", l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateItemHash(&_OutboxWithoutOptTester.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateItemHash(&_OutboxWithoutOptTester.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateMerkleRoot(&_OutboxWithoutOptTester.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateMerkleRoot(&_OutboxWithoutOptTester.CallOpts, proof, path, item)
}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) ExecuteTransactionSimulation(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "executeTransactionSimulation", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	if err != nil {
		return err
	}

	return err

}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) ExecuteTransactionSimulation(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	return _OutboxWithoutOptTester.Contract.ExecuteTransactionSimulation(&_OutboxWithoutOptTester.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) ExecuteTransactionSimulation(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	return _OutboxWithoutOptTester.Contract.ExecuteTransactionSimulation(&_OutboxWithoutOptTester.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) IsSpent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "isSpent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _OutboxWithoutOptTester.Contract.IsSpent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _OutboxWithoutOptTester.Contract.IsSpent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1BatchNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1BatchNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1BatchNum() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1BatchNum(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1BatchNum() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1BatchNum(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Block() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Block(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Block() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Block(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1EthBlock() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1EthBlock(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1EthBlock(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1OutputId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1OutputId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1OutputId() ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1OutputId(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1OutputId() ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1OutputId(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Sender() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Sender(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Sender() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Sender(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Timestamp() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Timestamp(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Timestamp(&_OutboxWithoutOptTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Rollup() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Rollup(&_OutboxWithoutOptTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Rollup() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Rollup(&_OutboxWithoutOptTester.CallOpts)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Roots(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Roots(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Roots(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Spent(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "spent", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Spent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Spent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) ExecuteTransaction(opts *bind.TransactOpts, proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "executeTransaction", proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.ExecuteTransaction(&_OutboxWithoutOptTester.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.ExecuteTransaction(&_OutboxWithoutOptTester.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.Initialize(&_OutboxWithoutOptTester.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.Initialize(&_OutboxWithoutOptTester.TransactOpts, _bridge)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) UpdateSendRoot(opts *bind.TransactOpts, root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "updateSendRoot", root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateSendRoot(&_OutboxWithoutOptTester.TransactOpts, root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateSendRoot(&_OutboxWithoutOptTester.TransactOpts, root, l2BlockHash)
}

// OutboxWithoutOptTesterOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterOutBoxTransactionExecutedIterator struct {
	Event *OutboxWithoutOptTesterOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
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
		it.Event = new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
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
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxWithoutOptTesterOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterOutBoxTransactionExecuted struct {
	To               common.Address
	L2Sender         common.Address
	Zero             *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, to []common.Address, l2Sender []common.Address, zero []*big.Int) (*OutboxWithoutOptTesterOutBoxTransactionExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.FilterLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterOutBoxTransactionExecutedIterator{contract: _OutboxWithoutOptTester.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *OutboxWithoutOptTesterOutBoxTransactionExecuted, to []common.Address, l2Sender []common.Address, zero []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.WatchLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
				if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*OutboxWithoutOptTesterOutBoxTransactionExecuted, error) {
	event := new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
	if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxWithoutOptTesterSendRootUpdatedIterator is returned from FilterSendRootUpdated and is used to iterate over the raw logs and unpacked data for SendRootUpdated events raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterSendRootUpdatedIterator struct {
	Event *OutboxWithoutOptTesterSendRootUpdated // Event containing the contract specifics and raw log

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
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxWithoutOptTesterSendRootUpdated)
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
		it.Event = new(OutboxWithoutOptTesterSendRootUpdated)
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
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxWithoutOptTesterSendRootUpdated represents a SendRootUpdated event raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterSendRootUpdated struct {
	OutputRoot  [32]byte
	L2BlockHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSendRootUpdated is a free log retrieval operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) FilterSendRootUpdated(opts *bind.FilterOpts, outputRoot [][32]byte, l2BlockHash [][32]byte) (*OutboxWithoutOptTesterSendRootUpdatedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockHashRule []interface{}
	for _, l2BlockHashItem := range l2BlockHash {
		l2BlockHashRule = append(l2BlockHashRule, l2BlockHashItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.FilterLogs(opts, "SendRootUpdated", outputRootRule, l2BlockHashRule)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterSendRootUpdatedIterator{contract: _OutboxWithoutOptTester.contract, event: "SendRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSendRootUpdated is a free log subscription operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) WatchSendRootUpdated(opts *bind.WatchOpts, sink chan<- *OutboxWithoutOptTesterSendRootUpdated, outputRoot [][32]byte, l2BlockHash [][32]byte) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockHashRule []interface{}
	for _, l2BlockHashItem := range l2BlockHash {
		l2BlockHashRule = append(l2BlockHashRule, l2BlockHashItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.WatchLogs(opts, "SendRootUpdated", outputRootRule, l2BlockHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxWithoutOptTesterSendRootUpdated)
				if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
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

// ParseSendRootUpdated is a log parse operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) ParseSendRootUpdated(log types.Log) (*OutboxWithoutOptTesterSendRootUpdated, error) {
	event := new(OutboxWithoutOptTesterSendRootUpdated)
	if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMockMetaData contains all meta data concerning the RollupMock contract.
var RollupMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"WithdrawTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ZombieTriggered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060ac8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636137391914602d575b600080fd5b60336045565b60405190815260200160405180910390f35b6040516000907f1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f908290a15060009056fea26469706673582212204f76b9b2686fad490865cc82482a095a8f63ef5154cebe1af4a93a76e604d0f864736f6c63430008110033",
}

// RollupMockABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMockMetaData.ABI instead.
var RollupMockABI = RollupMockMetaData.ABI

// RollupMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMockMetaData.Bin instead.
var RollupMockBin = RollupMockMetaData.Bin

// DeployRollupMock deploys a new Ethereum contract, binding an instance of RollupMock to it.
func DeployRollupMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupMock, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// RollupMock is an auto generated Go binding around an Ethereum contract.
type RollupMock struct {
	RollupMockCaller     // Read-only binding to the contract
	RollupMockTransactor // Write-only binding to the contract
	RollupMockFilterer   // Log filterer for contract events
}

// RollupMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupMockSession struct {
	Contract     *RollupMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupMockCallerSession struct {
	Contract *RollupMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupMockTransactorSession struct {
	Contract     *RollupMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupMockRaw struct {
	Contract *RollupMock // Generic contract binding to access the raw methods on
}

// RollupMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupMockCallerRaw struct {
	Contract *RollupMockCaller // Generic read-only contract binding to access the raw methods on
}

// RollupMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupMockTransactorRaw struct {
	Contract *RollupMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupMock creates a new instance of RollupMock, bound to a specific deployed contract.
func NewRollupMock(address common.Address, backend bind.ContractBackend) (*RollupMock, error) {
	contract, err := bindRollupMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// NewRollupMockCaller creates a new read-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockCaller(address common.Address, caller bind.ContractCaller) (*RollupMockCaller, error) {
	contract, err := bindRollupMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockCaller{contract: contract}, nil
}

// NewRollupMockTransactor creates a new write-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupMockTransactor, error) {
	contract, err := bindRollupMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockTransactor{contract: contract}, nil
}

// NewRollupMockFilterer creates a new log filterer instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupMockFilterer, error) {
	contract, err := bindRollupMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupMockFilterer{contract: contract}, nil
}

// bindRollupMock binds a generic wrapper to an already deployed contract.
func bindRollupMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.RollupMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transact(opts, method, params...)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockTransactor) WithdrawStakerFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.contract.Transact(opts, "withdrawStakerFunds")
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockSession) WithdrawStakerFunds() (*types.Transaction, error) {
	return _RollupMock.Contract.WithdrawStakerFunds(&_RollupMock.TransactOpts)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockTransactorSession) WithdrawStakerFunds() (*types.Transaction, error) {
	return _RollupMock.Contract.WithdrawStakerFunds(&_RollupMock.TransactOpts)
}

// RollupMockWithdrawTriggeredIterator is returned from FilterWithdrawTriggered and is used to iterate over the raw logs and unpacked data for WithdrawTriggered events raised by the RollupMock contract.
type RollupMockWithdrawTriggeredIterator struct {
	Event *RollupMockWithdrawTriggered // Event containing the contract specifics and raw log

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
func (it *RollupMockWithdrawTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMockWithdrawTriggered)
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
		it.Event = new(RollupMockWithdrawTriggered)
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
func (it *RollupMockWithdrawTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMockWithdrawTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMockWithdrawTriggered represents a WithdrawTriggered event raised by the RollupMock contract.
type RollupMockWithdrawTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTriggered is a free log retrieval operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) FilterWithdrawTriggered(opts *bind.FilterOpts) (*RollupMockWithdrawTriggeredIterator, error) {

	logs, sub, err := _RollupMock.contract.FilterLogs(opts, "WithdrawTriggered")
	if err != nil {
		return nil, err
	}
	return &RollupMockWithdrawTriggeredIterator{contract: _RollupMock.contract, event: "WithdrawTriggered", logs: logs, sub: sub}, nil
}

// WatchWithdrawTriggered is a free log subscription operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) WatchWithdrawTriggered(opts *bind.WatchOpts, sink chan<- *RollupMockWithdrawTriggered) (event.Subscription, error) {

	logs, sub, err := _RollupMock.contract.WatchLogs(opts, "WithdrawTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMockWithdrawTriggered)
				if err := _RollupMock.contract.UnpackLog(event, "WithdrawTriggered", log); err != nil {
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

// ParseWithdrawTriggered is a log parse operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) ParseWithdrawTriggered(log types.Log) (*RollupMockWithdrawTriggered, error) {
	event := new(RollupMockWithdrawTriggered)
	if err := _RollupMock.contract.UnpackLog(event, "WithdrawTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMockZombieTriggeredIterator is returned from FilterZombieTriggered and is used to iterate over the raw logs and unpacked data for ZombieTriggered events raised by the RollupMock contract.
type RollupMockZombieTriggeredIterator struct {
	Event *RollupMockZombieTriggered // Event containing the contract specifics and raw log

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
func (it *RollupMockZombieTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMockZombieTriggered)
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
		it.Event = new(RollupMockZombieTriggered)
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
func (it *RollupMockZombieTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMockZombieTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMockZombieTriggered represents a ZombieTriggered event raised by the RollupMock contract.
type RollupMockZombieTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterZombieTriggered is a free log retrieval operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) FilterZombieTriggered(opts *bind.FilterOpts) (*RollupMockZombieTriggeredIterator, error) {

	logs, sub, err := _RollupMock.contract.FilterLogs(opts, "ZombieTriggered")
	if err != nil {
		return nil, err
	}
	return &RollupMockZombieTriggeredIterator{contract: _RollupMock.contract, event: "ZombieTriggered", logs: logs, sub: sub}, nil
}

// WatchZombieTriggered is a free log subscription operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) WatchZombieTriggered(opts *bind.WatchOpts, sink chan<- *RollupMockZombieTriggered) (event.Subscription, error) {

	logs, sub, err := _RollupMock.contract.WatchLogs(opts, "ZombieTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMockZombieTriggered)
				if err := _RollupMock.contract.UnpackLog(event, "ZombieTriggered", log); err != nil {
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

// ParseZombieTriggered is a log parse operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) ParseZombieTriggered(log types.Log) (*RollupMockZombieTriggered, error) {
	event := new(RollupMockZombieTriggered)
	if err := _RollupMock.contract.UnpackLog(event, "ZombieTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueArrayTesterMetaData contains all meta data concerning the ValueArrayTester contract.
var ValueArrayTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061080a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f8a8fd6d14610030575b600080fd5b61003861003a565b005b6040805160026020820181815260808301845260009383929083015b604080518082019091526000808252602082015281526020019060019003908161005657505090528051519091506002146100d85760405162461bcd60e51b815260206004820152600960248201527f53544152545f4c454e000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b60408051808201825260008082526020918201819052825180840190935280835260019183019190915261010f915b8391906104b9565b604080518082018252600080825260209182018190528251808401909352825260029082015261014190600190610107565b60408051808201825260008082526020918201819052825180840190935282526003908201526101729082906104e0565b8051516003146101c45760405162461bcd60e51b815260206004820152600860248201527f505553485f4c454e00000000000000000000000000000000000000000000000060448201526064016100cf565b60005b8151518110156102b15760006101dd83836105d4565b90506000815160068111156101f4576101f4610734565b146102415760405162461bcd60e51b815260206004820152600d60248201527f505553485f56414c5f545950450000000000000000000000000000000000000060448201526064016100cf565b61024c826001610760565b81602001511461029e5760405162461bcd60e51b815260206004820152601160248201527f505553485f56414c5f434f4e54454e545300000000000000000000000000000060448201526064016100cf565b50806102a981610773565b9150506101c7565b5060006102bd8261060d565b90506000815160068111156102d4576102d4610734565b146103215760405162461bcd60e51b815260206004820152600c60248201527f504f505f5245545f54595045000000000000000000000000000000000000000060448201526064016100cf565b80602001516003146103755760405162461bcd60e51b815260206004820152601060248201527f504f505f5245545f434f4e54454e54530000000000000000000000000000000060448201526064016100cf565b8151516002146103c75760405162461bcd60e51b815260206004820152600760248201527f504f505f4c454e0000000000000000000000000000000000000000000000000060448201526064016100cf565b60005b8251518110156104b45760006103e084836105d4565b90506000815160068111156103f7576103f7610734565b146104445760405162461bcd60e51b815260206004820152600c60248201527f504f505f56414c5f54595045000000000000000000000000000000000000000060448201526064016100cf565b61044f826001610760565b8160200151146104a15760405162461bcd60e51b815260206004820152601060248201527f504f505f56414c5f434f4e54454e54530000000000000000000000000000000060448201526064016100cf565b50806104ac81610773565b9150506103ca565b505050565b80836000015183815181106104d0576104d06107ab565b6020026020010181905250505050565b8151516000906104f1906001610760565b67ffffffffffffffff8111156105095761050961071e565b60405190808252806020026020018201604052801561054e57816020015b60408051808201909152600080825260208201528152602001906001900390816105275790505b50905060005b8351518110156105aa578351805182908110610572576105726107ab565b602002602001015182828151811061058c5761058c6107ab565b602002602001018190525080806105a290610773565b915050610554565b508181846000015151815181106105c3576105c36107ab565b602090810291909101015290915250565b604080518082019091526000808252602082015282518051839081106105fc576105fc6107ab565b602002602001015190505b92915050565b604080518082019091526000808252602082015281518051610631906001906107c1565b81518110610641576106416107ab565b602002602001015190506000600183600001515161065f91906107c1565b67ffffffffffffffff8111156106775761067761071e565b6040519080825280602002602001820160405280156106bc57816020015b60408051808201909152600080825260208201528152602001906001900390816106955790505b50905060005b81518110156107175783518051829081106106df576106df6107ab565b60200260200101518282815181106106f9576106f96107ab565b6020026020010181905250808061070f90610773565b9150506106c2565b5090915290565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156106075761060761074a565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107a4576107a461074a565b5060010190565b634e487b7160e01b600052603260045260246000fd5b818103818111156106075761060761074a56fea2646970667358221220998f8916386b0d930286fafab271d8338de849670f2ee895dbcbef104a61d0cd64736f6c63430008110033",
}

// ValueArrayTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use ValueArrayTesterMetaData.ABI instead.
var ValueArrayTesterABI = ValueArrayTesterMetaData.ABI

// ValueArrayTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValueArrayTesterMetaData.Bin instead.
var ValueArrayTesterBin = ValueArrayTesterMetaData.Bin

// DeployValueArrayTester deploys a new Ethereum contract, binding an instance of ValueArrayTester to it.
func DeployValueArrayTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueArrayTester, error) {
	parsed, err := ValueArrayTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValueArrayTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueArrayTester{ValueArrayTesterCaller: ValueArrayTesterCaller{contract: contract}, ValueArrayTesterTransactor: ValueArrayTesterTransactor{contract: contract}, ValueArrayTesterFilterer: ValueArrayTesterFilterer{contract: contract}}, nil
}

// ValueArrayTester is an auto generated Go binding around an Ethereum contract.
type ValueArrayTester struct {
	ValueArrayTesterCaller     // Read-only binding to the contract
	ValueArrayTesterTransactor // Write-only binding to the contract
	ValueArrayTesterFilterer   // Log filterer for contract events
}

// ValueArrayTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueArrayTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueArrayTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueArrayTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueArrayTesterSession struct {
	Contract     *ValueArrayTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueArrayTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueArrayTesterCallerSession struct {
	Contract *ValueArrayTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValueArrayTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueArrayTesterTransactorSession struct {
	Contract     *ValueArrayTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValueArrayTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueArrayTesterRaw struct {
	Contract *ValueArrayTester // Generic contract binding to access the raw methods on
}

// ValueArrayTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueArrayTesterCallerRaw struct {
	Contract *ValueArrayTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ValueArrayTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueArrayTesterTransactorRaw struct {
	Contract *ValueArrayTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueArrayTester creates a new instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTester(address common.Address, backend bind.ContractBackend) (*ValueArrayTester, error) {
	contract, err := bindValueArrayTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTester{ValueArrayTesterCaller: ValueArrayTesterCaller{contract: contract}, ValueArrayTesterTransactor: ValueArrayTesterTransactor{contract: contract}, ValueArrayTesterFilterer: ValueArrayTesterFilterer{contract: contract}}, nil
}

// NewValueArrayTesterCaller creates a new read-only instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterCaller(address common.Address, caller bind.ContractCaller) (*ValueArrayTesterCaller, error) {
	contract, err := bindValueArrayTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterCaller{contract: contract}, nil
}

// NewValueArrayTesterTransactor creates a new write-only instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueArrayTesterTransactor, error) {
	contract, err := bindValueArrayTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterTransactor{contract: contract}, nil
}

// NewValueArrayTesterFilterer creates a new log filterer instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueArrayTesterFilterer, error) {
	contract, err := bindValueArrayTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterFilterer{contract: contract}, nil
}

// bindValueArrayTester binds a generic wrapper to an already deployed contract.
func bindValueArrayTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValueArrayTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueArrayTester *ValueArrayTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueArrayTester.Contract.ValueArrayTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueArrayTester *ValueArrayTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.ValueArrayTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueArrayTester *ValueArrayTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.ValueArrayTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueArrayTester *ValueArrayTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueArrayTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueArrayTester *ValueArrayTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueArrayTester *ValueArrayTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.contract.Transact(opts, method, params...)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterCaller) Test(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ValueArrayTester.contract.Call(opts, &out, "test")

	if err != nil {
		return err
	}

	return err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterSession) Test() error {
	return _ValueArrayTester.Contract.Test(&_ValueArrayTester.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterCallerSession) Test() error {
	return _ValueArrayTester.Contract.Test(&_ValueArrayTester.CallOpts)
}
