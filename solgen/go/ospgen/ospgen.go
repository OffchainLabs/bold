// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ospgen

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

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead  *big.Int
	Bridge                common.Address
	InitialWasmModuleRoot [32]byte
}

// ExecutionState is an auto generated low-level Go binding around an user-defined struct.
type ExecutionState struct {
	GlobalState   GlobalState
	MachineStatus uint8
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// Instruction is an auto generated low-level Go binding around an user-defined struct.
type Instruction struct {
	Opcode       uint16
	ArgumentData *big.Int
}

// Machine is an auto generated low-level Go binding around an user-defined struct.
type Machine struct {
	Status          uint8
	ValueStack      ValueStack
	ValueMultiStack MultiStack
	InternalStack   ValueStack
	FrameStack      StackFrameWindow
	FrameMultiStack MultiStack
	GlobalStateHash [32]byte
	ModuleIdx       uint32
	FunctionIdx     uint32
	FunctionPc      uint32
	RecoveryPc      [32]byte
	ModulesRoot     [32]byte
}

// Module is an auto generated low-level Go binding around an user-defined struct.
type Module struct {
	GlobalsMerkleRoot   [32]byte
	ModuleMemory        ModuleMemory
	TablesMerkleRoot    [32]byte
	FunctionsMerkleRoot [32]byte
	ExtraHash           [32]byte
	InternalsOffset     uint32
}

// ModuleMemory is an auto generated low-level Go binding around an user-defined struct.
type ModuleMemory struct {
	Size       uint64
	MaxSize    uint64
	MerkleRoot [32]byte
}

// MultiStack is an auto generated low-level Go binding around an user-defined struct.
type MultiStack struct {
	InactiveStackHash [32]byte
	RemainingHash     [32]byte
}

// StackFrame is an auto generated low-level Go binding around an user-defined struct.
type StackFrame struct {
	ReturnPc              Value
	LocalsMerkleRoot      [32]byte
	CallerModule          uint32
	CallerModuleInternals uint32
}

// StackFrameWindow is an auto generated low-level Go binding around an user-defined struct.
type StackFrameWindow struct {
	Proved        []StackFrame
	RemainingHash [32]byte
}

// Value is an auto generated low-level Go binding around an user-defined struct.
type Value struct {
	ValueType uint8
	Contents  *big.Int
}

// ValueArray is an auto generated low-level Go binding around an user-defined struct.
type ValueArray struct {
	Inner []Value
}

// ValueStack is an auto generated low-level Go binding around an user-defined struct.
type ValueStack struct {
	Proved        ValueArray
	RemainingHash [32]byte
}

// HashProofHelperMetaData contains all meta data concerning the HashProofHelper contract.
var HashProofHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"NotProven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"}],\"name\":\"PreimagePartProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearSplitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"getPreimagePart\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keccakStates\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"proveWithFullPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"proveWithSplitPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e58806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063ae364ac211610050578063ae364ac2146100b6578063b7465799146100c0578063d4e5dd2b146100e257600080fd5b8063740085d71461006c57806379754cba14610095575b600080fd5b61007f61007a36600461194f565b6100f5565b60405161008c91906119c1565b60405180910390f35b6100a86100a3366004611a24565b610204565b60405190815260200161008c565b6100be610767565b005b6100d36100ce366004611a80565b6107af565b60405161008c93929190611ab6565b6100a86100f0366004611ae9565b610866565b60008281526020818152604080832067ffffffffffffffff85168452909152902080546060919060ff1661016d576040517f139647920000000000000000000000000000000000000000000000000000000081526004810185905267ffffffffffffffff841660248201526044015b60405180910390fd5b80600101805461017c90611b3d565b80601f01602080910402602001604051908101604052809291908181526020018280546101a890611b3d565b80156101f55780601f106101ca576101008083540402835291602001916101f5565b820191906000526020600020905b8154815290600101906020018083116101d857829003601f168201915b50505050509150505b92915050565b6000600182161515600283161561025c573360009081526001602081905260408220805467ffffffffffffffff19168155919061024390830182611890565b6102516002830160006118cd565b600982016000905550505b8080610270575061026e608886611b8d565b155b6102d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f4e4f545f424c4f434b5f414c49474e45440000000000000000000000000000006044820152606401610164565b3360009081526001602052604081206009810154909181900361031357815467ffffffffffffffff191667ffffffffffffffff871617825561038a565b815467ffffffffffffffff87811691161461038a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f444946465f4f46465345540000000000000000000000000000000000000000006044820152606401610164565b610396828989866109bd565b806103ac602067ffffffffffffffff8916611bb7565b1180156103c6575081600901548667ffffffffffffffff16105b156104fa576000818767ffffffffffffffff1611156103f6576103f38267ffffffffffffffff8916611bca565b90505b60008261040e602067ffffffffffffffff8b16611bb7565b6104189190611bca565b9050888111156104255750875b815b818110156104f657846001018b8b8381811061044557610445611bdd565b9050013560f81c60f81b908080548061045d90611b3d565b80601f810361047c5783600052602060002060ff1984168155603f9350505b506002919091019091558154600116156104a55790600052602060002090602091828204019190065b909190919091601f036101000a81548160ff021916907f01000000000000000000000000000000000000000000000000000000000000008404021790555080806104ee90611bf3565b915050610427565b5050505b8261050c57506000925061075f915050565b60005b60208110156105dd576000610525600883611c0d565b9050610532600582611c0d565b61053d600583611b8d565b610548906005611c21565b6105529190611bb7565b90506000610561600884611b8d565b61056c906008611c21565b85600201836019811061058157610581611bdd565b600481049091015467ffffffffffffffff6008600390931683026101000a9091041690911c91506105b3908490611c21565b6105be9060f8611bca565b60ff909116901b959095179450806105d581611bf3565b91505061050f565b50604051806040016040528060011515815260200183600101805461060190611b3d565b80601f016020809104026020016040519081016040528092919081815260200182805461062d90611b3d565b801561067a5780601f1061064f5761010080835404028352916020019161067a565b820191906000526020600020905b81548152906001019060200180831161065d57829003601f168201915b505050919092525050600085815260208181526040808320865467ffffffffffffffff16845282529091208251815460ff19169015151781559082015160018201906106c69082611c9d565b5050825460405167ffffffffffffffff909116915085907ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c9061070d906001870190611d5d565b60405180910390a33360009081526001602081905260408220805467ffffffffffffffff19168155919061074390830182611890565b6107516002830160006118cd565b600982016000905550505050505b949350505050565b3360009081526001602081905260408220805467ffffffffffffffff19168155919061079590830182611890565b6107a36002830160006118cd565b60098201600090555050565b60016020819052600091825260409091208054918101805467ffffffffffffffff909316926107dd90611b3d565b80601f016020809104026020016040519081016040528092919081815260200182805461080990611b3d565b80156108565780601f1061082b57610100808354040283529160200191610856565b820191906000526020600020905b81548152906001019060200180831161083957829003601f168201915b5050505050908060090154905083565b60008383604051610878929190611de8565b6040519081900390209050606067ffffffffffffffff83168411156109195760006108ad67ffffffffffffffff851686611bca565b905060208111156108bc575060205b8567ffffffffffffffff8516866108d38483611bb7565b926108e093929190611df8565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929450505050505b60408051808201825260018082526020808301858152600087815280835285812067ffffffffffffffff8a1682529092529390208251815460ff191690151517815592519192919082019061096e9082611c9d565b509050508267ffffffffffffffff16827ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c836040516109ad91906119c1565b60405180910390a3509392505050565b828290508460090160008282546109d49190611bb7565b90915550505b811580156109e6575080155b610c3e5760005b6088811015610b1257600083821015610a2357848483818110610a1257610a12611bdd565b919091013560f81c9150610a449050565b838203610a2e576001175b610a3a60016088611bca565b8203610a44576080175b6000610a51600884611c0d565b9050610a5e600582611c0d565b610a69600583611b8d565b610a74906005611c21565b610a7e9190611bb7565b9050610a8b600884611b8d565b610a96906008611c21565b67ffffffffffffffff168260ff1667ffffffffffffffff16901b876002018260198110610ac557610ac5611bdd565b60048104909101805467ffffffffffffffff60086003909416939093026101000a808204841690941883168402929093021990921617905550819050610b0a81611bf3565b9150506109ed565b50610b1b6118dc565b60005b6019811015610b8f57856002018160198110610b3c57610b3c611bdd565b600491828204019190066008029054906101000a900467ffffffffffffffff1667ffffffffffffffff16828260198110610b7857610b78611bdd565b602002015280610b8781611bf3565b915050610b1e565b50610b9981610c44565b905060005b6019811015610c1757818160198110610bb957610bb9611bdd565b6020020151866002018260198110610bd357610bd3611bdd565b600491828204019190066008026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508080610c0f90611bf3565b915050610b9e565b506088831015610c275750610c3e565b610c348360888187611df8565b93509350506109da565b50505050565b610c4c6118dc565b610c546118fb565b610c5c6118fb565b610c646118dc565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015611885576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e01511818181892880183905267ffffffffffffffff6002820216678000000000000000918290041790921886525104856002602002015160020267ffffffffffffffff16178560006020020151188460016020020152678000000000000000856003602002015181610ebf57610ebf611b77565b04856003602002015160020267ffffffffffffffff16178560016020020151188460026020020152678000000000000000856004602002015181610f0557610f05611b77565b04856004602002015160020267ffffffffffffffff161785600260058110610f2f57610f2f611bdd565b6020020151186060850152845167800000000000000090865160608089015193909204600290910267ffffffffffffffff1617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a018051909118905290845251631000000090602089015167ffffffffffffffff6410000000009091021691900417610100840152604087015167200000000000000090604089015167ffffffffffffffff6008909102169190041761016084015260608701516280000090606089015167ffffffffffffffff65020000000000909102169190041761026084015260808701516540000000000090608089015167ffffffffffffffff6204000090910216919004176102c084015260a08701516780000000000000009004876005602002015160020267ffffffffffffffff1617836002601981106111b1576111b1611bdd565b602002015260c08701516210000081046510000000000090910267ffffffffffffffff9081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166740000000000000009091041761030085015261014088015161014089015167ffffffffffffffff674000000000000000909102169190041760808401526101608701516704000000000000009061016089015167ffffffffffffffff6040909102169190041760e0840152610180870151622000009061018089015167ffffffffffffffff6508000000000090910216919004176101408401526101a08701516602000000000000906101a089015167ffffffffffffffff61800090910216919004176102408401526101c08701516008906101c089015167ffffffffffffffff67200000000000000090910216919004176102a08401526101e0870151641000000000906101e089015167ffffffffffffffff6310000000909102169190041760208401526102008088015161020089015167ffffffffffffffff668000000000000090910216919004176101208401526102208701516480000000009061022089015167ffffffffffffffff63020000009091021691900417610180840152610240870151650800000000009061024089015167ffffffffffffffff6220000090910216919004176101e08401526102608701516101009061026089015167ffffffffffffffff67010000000000000090910216919004176102e08401526102808701516420000000009061028089015167ffffffffffffffff6308000000909102169190041760608401526102a087015165100000000000906102a089015167ffffffffffffffff62100000909102169190041760c08401526102c08701516302000000906102c089015167ffffffffffffffff64800000000090910216919004176101c08401526102e0870151670100000000000000906102e089015167ffffffffffffffff61010090910216919004176102208401526103008701516604000000000000900487601860200201516140000267ffffffffffffffff1617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f602002015282600160200201518360156020020151191683601060200201511887601060200201528260026020020151836016602002015119168360116020020151188760116020020152826003602002015183601760200201511916836012602002015118876012602002015282600460200201518360186020020151191683601360200201511887601360200201528260056020020151836000602002015119168360146020020151188760146020020152826006602002015183600160200201511916836015602002015118876015602002015282600760200201518360026020020151191683601660200201511887601660200201528260086020020151836003602002015119168360176020020151188760176020020152826009602002015183600460200201511916836018602002015118876018602002015281816018811061187357611873611bdd565b60200201518751188752600101610d8a565b509495945050505050565b50805461189c90611b3d565b6000825580601f106118ac575050565b601f0160209004906000526020600020908101906118ca9190611919565b50565b506118ca906007810190611919565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b5b8082111561192e576000815560010161191a565b5090565b803567ffffffffffffffff8116811461194a57600080fd5b919050565b6000806040838503121561196257600080fd5b8235915061197260208401611932565b90509250929050565b6000815180845260005b818110156119a157602081850181015186830182015201611985565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006119d4602083018461197b565b9392505050565b60008083601f8401126119ed57600080fd5b50813567ffffffffffffffff811115611a0557600080fd5b602083019150836020828501011115611a1d57600080fd5b9250929050565b60008060008060608587031215611a3a57600080fd5b843567ffffffffffffffff811115611a5157600080fd5b611a5d878288016119db565b9095509350611a70905060208601611932565b9396929550929360400135925050565b600060208284031215611a9257600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146119d457600080fd5b67ffffffffffffffff84168152606060208201526000611ad9606083018561197b565b9050826040830152949350505050565b600080600060408486031215611afe57600080fd5b833567ffffffffffffffff811115611b1557600080fd5b611b21868287016119db565b9094509250611b34905060208501611932565b90509250925092565b600181811c90821680611b5157607f821691505b602082108103611b7157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601260045260246000fd5b600082611b9c57611b9c611b77565b500690565b634e487b7160e01b600052601160045260246000fd5b808201808211156101fe576101fe611ba1565b818103818111156101fe576101fe611ba1565b634e487b7160e01b600052603260045260246000fd5b60006000198203611c0657611c06611ba1565b5060010190565b600082611c1c57611c1c611b77565b500490565b80820281158282048414176101fe576101fe611ba1565b634e487b7160e01b600052604160045260246000fd5b601f821115611c9857600081815260208120601f850160051c81016020861015611c755750805b601f850160051c820191505b81811015611c9457828155600101611c81565b5050505b505050565b815167ffffffffffffffff811115611cb757611cb7611c38565b611ccb81611cc58454611b3d565b84611c4e565b602080601f831160018114611d005760008415611ce85750858301515b600019600386901b1c1916600185901b178555611c94565b600085815260208120601f198616915b82811015611d2f57888601518255948401946001909101908401611d10565b5085821015611d4d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602080835260008454611d7181611b3d565b80848701526040600180841660008114611d925760018114611dac57611dda565b60ff198516838a01528284151560051b8a01019550611dda565b896000528660002060005b85811015611dd25781548b8201860152908301908801611db7565b8a0184019650505b509398975050505050505050565b8183823760009101908152919050565b60008085851115611e0857600080fd5b83861115611e1557600080fd5b505082019391909203915056fea26469706673582212206f97ea9f601c498d1c41e40bccd525fb2b012712e46a4bbe5007cb7d6976986864736f6c63430008110033",
}

// HashProofHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use HashProofHelperMetaData.ABI instead.
var HashProofHelperABI = HashProofHelperMetaData.ABI

// HashProofHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashProofHelperMetaData.Bin instead.
var HashProofHelperBin = HashProofHelperMetaData.Bin

// DeployHashProofHelper deploys a new Ethereum contract, binding an instance of HashProofHelper to it.
func DeployHashProofHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashProofHelper, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashProofHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// HashProofHelper is an auto generated Go binding around an Ethereum contract.
type HashProofHelper struct {
	HashProofHelperCaller     // Read-only binding to the contract
	HashProofHelperTransactor // Write-only binding to the contract
	HashProofHelperFilterer   // Log filterer for contract events
}

// HashProofHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashProofHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashProofHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashProofHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashProofHelperSession struct {
	Contract     *HashProofHelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashProofHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashProofHelperCallerSession struct {
	Contract *HashProofHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HashProofHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashProofHelperTransactorSession struct {
	Contract     *HashProofHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HashProofHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashProofHelperRaw struct {
	Contract *HashProofHelper // Generic contract binding to access the raw methods on
}

// HashProofHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashProofHelperCallerRaw struct {
	Contract *HashProofHelperCaller // Generic read-only contract binding to access the raw methods on
}

// HashProofHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashProofHelperTransactorRaw struct {
	Contract *HashProofHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashProofHelper creates a new instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelper(address common.Address, backend bind.ContractBackend) (*HashProofHelper, error) {
	contract, err := bindHashProofHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// NewHashProofHelperCaller creates a new read-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperCaller(address common.Address, caller bind.ContractCaller) (*HashProofHelperCaller, error) {
	contract, err := bindHashProofHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperCaller{contract: contract}, nil
}

// NewHashProofHelperTransactor creates a new write-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HashProofHelperTransactor, error) {
	contract, err := bindHashProofHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperTransactor{contract: contract}, nil
}

// NewHashProofHelperFilterer creates a new log filterer instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HashProofHelperFilterer, error) {
	contract, err := bindHashProofHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperFilterer{contract: contract}, nil
}

// bindHashProofHelper binds a generic wrapper to an already deployed contract.
func bindHashProofHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.HashProofHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transact(opts, method, params...)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCaller) GetPreimagePart(opts *bind.CallOpts, fullHash [32]byte, offset uint64) ([]byte, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "getPreimagePart", fullHash, offset)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCallerSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCaller) KeccakStates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "keccakStates", arg0)

	outstruct := new(struct {
		Offset uint64
		Part   []byte
		Length *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Offset = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Part = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Length = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCallerSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactor) ClearSplitProof(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "clearSplitProof")
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactorSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithFullPreimage(opts *bind.TransactOpts, data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithFullPreimage", data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithSplitPreimage(opts *bind.TransactOpts, data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithSplitPreimage", data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// HashProofHelperPreimagePartProvenIterator is returned from FilterPreimagePartProven and is used to iterate over the raw logs and unpacked data for PreimagePartProven events raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProvenIterator struct {
	Event *HashProofHelperPreimagePartProven // Event containing the contract specifics and raw log

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
func (it *HashProofHelperPreimagePartProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashProofHelperPreimagePartProven)
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
		it.Event = new(HashProofHelperPreimagePartProven)
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
func (it *HashProofHelperPreimagePartProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashProofHelperPreimagePartProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashProofHelperPreimagePartProven represents a PreimagePartProven event raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProven struct {
	FullHash [32]byte
	Offset   uint64
	Part     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreimagePartProven is a free log retrieval operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) FilterPreimagePartProven(opts *bind.FilterOpts, fullHash [][32]byte, offset []uint64) (*HashProofHelperPreimagePartProvenIterator, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.FilterLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperPreimagePartProvenIterator{contract: _HashProofHelper.contract, event: "PreimagePartProven", logs: logs, sub: sub}, nil
}

// WatchPreimagePartProven is a free log subscription operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) WatchPreimagePartProven(opts *bind.WatchOpts, sink chan<- *HashProofHelperPreimagePartProven, fullHash [][32]byte, offset []uint64) (event.Subscription, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.WatchLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashProofHelperPreimagePartProven)
				if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
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

// ParsePreimagePartProven is a log parse operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) ParsePreimagePartProven(log types.Log) (*HashProofHelperPreimagePartProven, error) {
	event := new(HashProofHelperPreimagePartProven)
	if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICustomDAProofValidatorMetaData contains all meta data concerning the ICustomDAProofValidator contract.
var ICustomDAProofValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateReadPreimage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"preimageChunk\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICustomDAProofValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ICustomDAProofValidatorMetaData.ABI instead.
var ICustomDAProofValidatorABI = ICustomDAProofValidatorMetaData.ABI

// ICustomDAProofValidator is an auto generated Go binding around an Ethereum contract.
type ICustomDAProofValidator struct {
	ICustomDAProofValidatorCaller     // Read-only binding to the contract
	ICustomDAProofValidatorTransactor // Write-only binding to the contract
	ICustomDAProofValidatorFilterer   // Log filterer for contract events
}

// ICustomDAProofValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICustomDAProofValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICustomDAProofValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICustomDAProofValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICustomDAProofValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICustomDAProofValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICustomDAProofValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICustomDAProofValidatorSession struct {
	Contract     *ICustomDAProofValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ICustomDAProofValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICustomDAProofValidatorCallerSession struct {
	Contract *ICustomDAProofValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ICustomDAProofValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICustomDAProofValidatorTransactorSession struct {
	Contract     *ICustomDAProofValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ICustomDAProofValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICustomDAProofValidatorRaw struct {
	Contract *ICustomDAProofValidator // Generic contract binding to access the raw methods on
}

// ICustomDAProofValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICustomDAProofValidatorCallerRaw struct {
	Contract *ICustomDAProofValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ICustomDAProofValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICustomDAProofValidatorTransactorRaw struct {
	Contract *ICustomDAProofValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICustomDAProofValidator creates a new instance of ICustomDAProofValidator, bound to a specific deployed contract.
func NewICustomDAProofValidator(address common.Address, backend bind.ContractBackend) (*ICustomDAProofValidator, error) {
	contract, err := bindICustomDAProofValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICustomDAProofValidator{ICustomDAProofValidatorCaller: ICustomDAProofValidatorCaller{contract: contract}, ICustomDAProofValidatorTransactor: ICustomDAProofValidatorTransactor{contract: contract}, ICustomDAProofValidatorFilterer: ICustomDAProofValidatorFilterer{contract: contract}}, nil
}

// NewICustomDAProofValidatorCaller creates a new read-only instance of ICustomDAProofValidator, bound to a specific deployed contract.
func NewICustomDAProofValidatorCaller(address common.Address, caller bind.ContractCaller) (*ICustomDAProofValidatorCaller, error) {
	contract, err := bindICustomDAProofValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICustomDAProofValidatorCaller{contract: contract}, nil
}

// NewICustomDAProofValidatorTransactor creates a new write-only instance of ICustomDAProofValidator, bound to a specific deployed contract.
func NewICustomDAProofValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ICustomDAProofValidatorTransactor, error) {
	contract, err := bindICustomDAProofValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICustomDAProofValidatorTransactor{contract: contract}, nil
}

// NewICustomDAProofValidatorFilterer creates a new log filterer instance of ICustomDAProofValidator, bound to a specific deployed contract.
func NewICustomDAProofValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ICustomDAProofValidatorFilterer, error) {
	contract, err := bindICustomDAProofValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICustomDAProofValidatorFilterer{contract: contract}, nil
}

// bindICustomDAProofValidator binds a generic wrapper to an already deployed contract.
func bindICustomDAProofValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICustomDAProofValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICustomDAProofValidator *ICustomDAProofValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICustomDAProofValidator.Contract.ICustomDAProofValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICustomDAProofValidator *ICustomDAProofValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICustomDAProofValidator.Contract.ICustomDAProofValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICustomDAProofValidator *ICustomDAProofValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICustomDAProofValidator.Contract.ICustomDAProofValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICustomDAProofValidator *ICustomDAProofValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICustomDAProofValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICustomDAProofValidator *ICustomDAProofValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICustomDAProofValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICustomDAProofValidator *ICustomDAProofValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICustomDAProofValidator.Contract.contract.Transact(opts, method, params...)
}

// ValidateReadPreimage is a free data retrieval call binding the contract method 0xd2f26f52.
//
// Solidity: function validateReadPreimage(bytes proof) view returns(bytes preimageChunk)
func (_ICustomDAProofValidator *ICustomDAProofValidatorCaller) ValidateReadPreimage(opts *bind.CallOpts, proof []byte) ([]byte, error) {
	var out []interface{}
	err := _ICustomDAProofValidator.contract.Call(opts, &out, "validateReadPreimage", proof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateReadPreimage is a free data retrieval call binding the contract method 0xd2f26f52.
//
// Solidity: function validateReadPreimage(bytes proof) view returns(bytes preimageChunk)
func (_ICustomDAProofValidator *ICustomDAProofValidatorSession) ValidateReadPreimage(proof []byte) ([]byte, error) {
	return _ICustomDAProofValidator.Contract.ValidateReadPreimage(&_ICustomDAProofValidator.CallOpts, proof)
}

// ValidateReadPreimage is a free data retrieval call binding the contract method 0xd2f26f52.
//
// Solidity: function validateReadPreimage(bytes proof) view returns(bytes preimageChunk)
func (_ICustomDAProofValidator *ICustomDAProofValidatorCallerSession) ValidateReadPreimage(proof []byte) ([]byte, error) {
	return _ICustomDAProofValidator.Contract.ValidateReadPreimage(&_ICustomDAProofValidator.CallOpts, proof)
}

// IOneStepProofEntryMetaData contains all meta data concerning the IOneStepProofEntry contract.
var IOneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"execState\",\"type\":\"tuple\"}],\"name\":\"getMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProofEntryMetaData.ABI instead.
var IOneStepProofEntryABI = IOneStepProofEntryMetaData.ABI

// IOneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type IOneStepProofEntry struct {
	IOneStepProofEntryCaller     // Read-only binding to the contract
	IOneStepProofEntryTransactor // Write-only binding to the contract
	IOneStepProofEntryFilterer   // Log filterer for contract events
}

// IOneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofEntrySession struct {
	Contract     *IOneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofEntryCallerSession struct {
	Contract *IOneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IOneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofEntryTransactorSession struct {
	Contract     *IOneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IOneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofEntryRaw struct {
	Contract *IOneStepProofEntry // Generic contract binding to access the raw methods on
}

// IOneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCallerRaw struct {
	Contract *IOneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactorRaw struct {
	Contract *IOneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProofEntry creates a new instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*IOneStepProofEntry, error) {
	contract, err := bindIOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntry{IOneStepProofEntryCaller: IOneStepProofEntryCaller{contract: contract}, IOneStepProofEntryTransactor: IOneStepProofEntryTransactor{contract: contract}, IOneStepProofEntryFilterer: IOneStepProofEntryFilterer{contract: contract}}, nil
}

// NewIOneStepProofEntryCaller creates a new read-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofEntryCaller, error) {
	contract, err := bindIOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryCaller{contract: contract}, nil
}

// NewIOneStepProofEntryTransactor creates a new write-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofEntryTransactor, error) {
	contract, err := bindIOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryTransactor{contract: contract}, nil
}

// NewIOneStepProofEntryFilterer creates a new log filterer instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofEntryFilterer, error) {
	contract, err := bindIOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryFilterer{contract: contract}, nil
}

// bindIOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindIOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetMachineHash(opts *bind.CallOpts, execState ExecutionState) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getMachineHash", execState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetMachineHash(&_IOneStepProofEntry.CallOpts, execState)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetMachineHash(&_IOneStepProofEntry.CallOpts, execState)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetStartMachineHash(&_IOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetStartMachineHash(&_IOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// IOneStepProverMetaData contains all meta data concerning the IOneStepProver contract.
var IOneStepProverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"instruction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"result\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"resultMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProverABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProverMetaData.ABI instead.
var IOneStepProverABI = IOneStepProverMetaData.ABI

// IOneStepProver is an auto generated Go binding around an Ethereum contract.
type IOneStepProver struct {
	IOneStepProverCaller     // Read-only binding to the contract
	IOneStepProverTransactor // Write-only binding to the contract
	IOneStepProverFilterer   // Log filterer for contract events
}

// IOneStepProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProverSession struct {
	Contract     *IOneStepProver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProverCallerSession struct {
	Contract *IOneStepProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOneStepProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProverTransactorSession struct {
	Contract     *IOneStepProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOneStepProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProverRaw struct {
	Contract *IOneStepProver // Generic contract binding to access the raw methods on
}

// IOneStepProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProverCallerRaw struct {
	Contract *IOneStepProverCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProverTransactorRaw struct {
	Contract *IOneStepProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProver creates a new instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProver(address common.Address, backend bind.ContractBackend) (*IOneStepProver, error) {
	contract, err := bindIOneStepProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProver{IOneStepProverCaller: IOneStepProverCaller{contract: contract}, IOneStepProverTransactor: IOneStepProverTransactor{contract: contract}, IOneStepProverFilterer: IOneStepProverFilterer{contract: contract}}, nil
}

// NewIOneStepProverCaller creates a new read-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProverCaller, error) {
	contract, err := bindIOneStepProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverCaller{contract: contract}, nil
}

// NewIOneStepProverTransactor creates a new write-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProverTransactor, error) {
	contract, err := bindIOneStepProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverTransactor{contract: contract}, nil
}

// NewIOneStepProverFilterer creates a new log filterer instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProverFilterer, error) {
	contract, err := bindIOneStepProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverFilterer{contract: contract}, nil
}

// bindIOneStepProver binds a generic wrapper to an already deployed contract.
func bindIOneStepProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.IOneStepProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	var out []interface{}
	err := _IOneStepProver.contract.Call(opts, &out, "executeOneStep", execCtx, mach, mod, instruction, proof)

	outstruct := new(struct {
		Result    Machine
		ResultMod Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.ResultMod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCallerSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// OneStepProofEntryMetaData contains all meta data concerning the OneStepProofEntry contract.
var OneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"prover0_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMem_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMath_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverHostIo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"execState\",\"type\":\"tuple\"}],\"name\":\"getMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prover0\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverHostIo\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMath\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMem\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003595380380620035958339810160408190526200003491620000a5565b600080546001600160a01b039586166001600160a01b031991821617909155600180549486169482169490941790935560028054928516928416929092179091556003805491909316911617905562000102565b80516001600160a01b0381168114620000a057600080fd5b919050565b60008060008060808587031215620000bc57600080fd5b620000c78562000088565b9350620000d76020860162000088565b9250620000e76040860162000088565b9150620000f76060860162000088565b905092959194509250565b61348380620001126000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635f52fd7c1161005b5780635f52fd7c146100e657806366e5d9c3146100f9578063b5112fd21461010c578063c39619c41461011f57600080fd5b806304997be4146100825780631f128bc0146100a857806330a5509f146100d3575b600080fd5b610095610090366004612794565b610132565b6040519081526020015b60405180910390f35b6001546100bb906001600160a01b031681565b6040516001600160a01b03909116815260200161009f565b6000546100bb906001600160a01b031681565b6003546100bb906001600160a01b031681565b6002546100bb906001600160a01b031681565b61009561011a3660046127b6565b61034e565b61009561012d366004612852565b610b11565b60408051600380825260808201909252600091829190816020015b604080518082019091526000808252602082015281526020019060019003908161014d5750506040805180820182526000808252602091820181905282518084019093526004835290820152909150816000815181106101af576101af61287a565b60200260200101819052506101f26000604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b816001815181106102055761020561287a565b60200260200101819052506102486000604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b8160028151811061025b5761025b61287a565b60209081029190910181019190915260408051808301825283815281518083019092528082526000928201929092526102ab60408051606080820183529181019182529081526000602082015290565b604080518082018252606081526000602080830182905283518085019094528301526000198252906040805161018081018252600080825260208201879052918101839052606081018590526080810184905260a0810183905260c081018b905260e081018290526101008101829052610120810191909152600019610140820152610160810189905261033e81610c64565b9750505050505050505b92915050565b6000610358612670565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a0810191909152604080516020810190915260608152604080518082019091526000808252602082015260006103d2888883610eb4565b9095509050886103e186610c64565b146104335760405162461bcd60e51b815260206004820152601360248201527f4d414348494e455f4245464f52455f484153480000000000000000000000000060448201526064015b60405180910390fd5b60008551600281111561044857610448612890565b1461052f57610455612751565b61046089898461112c565b60c088015190935090915061047482611208565b146104c15760405162461bcd60e51b815260206004820152601060248201527f4241445f474c4f42414c5f535441544500000000000000000000000000000000604482015260640161042a565b6001865160028111156104d6576104d6612890565b1480156104e157508a155b801561050257508b356104f682602001515190565b67ffffffffffffffff16105b15610526576105198660c001518d60400135610132565b9650505050505050610b08565b61051986610c64565b650800000000006105418b60016128bc565b0361055f576002855261055385610c64565b95505050505050610b08565b61056a888883611298565b909450905061057a88888361139e565b80925081945050508461016001516105a78660e0015163ffffffff1686866114799092919063ffffffff16565b146105f45760405162461bcd60e51b815260206004820152600c60248201527f4d4f44554c45535f524f4f540000000000000000000000000000000000000000604482015260640161042a565b606061060c6040518060200160405280606081525090565b6040805160208101909152606081526106268b8b866114ce565b945092506106358b8b8661139e565b945091506106448b8b8661139e565b8095508192505050600061067a60408a610120015161066391906128e5565b63ffffffff1685856115ce9092919063ffffffff16565b9050600061069e8a610100015163ffffffff1683856116199092919063ffffffff16565b9050886060015181146106f35760405162461bcd60e51b815260206004820152601260248201527f4241445f46554e4354494f4e535f524f4f540000000000000000000000000000604482015260640161042a565b8460408b61012001516107069190612908565b63ffffffff168151811061071c5761071c61287a565b60200260200101519650505050505087878290809261073d9392919061292b565b975097505060008460e0015163ffffffff169050600185610120018181516107659190612955565b63ffffffff1690525081516000602861ffff83161080159061078c5750603561ffff831611155b806107ac5750603661ffff8316108015906107ac5750603e61ffff831611155b806107bb575061ffff8216603f145b806107ca575061ffff82166040145b156107e157506001546001600160a01b03166109f8565b61ffff8216604514806107f8575061ffff82166050145b806108265750604661ffff831610801590610826575061081a60096046612979565b61ffff168261ffff1611155b806108545750606761ffff831610801590610854575061084860026067612979565b61ffff168261ffff1611155b806108745750606a61ffff8316108015906108745750607861ffff831611155b806108a25750605161ffff8316108015906108a2575061089660096051612979565b61ffff168261ffff1611155b806108d05750607961ffff8316108015906108d057506108c460026079612979565b61ffff168261ffff1611155b806108f05750607c61ffff8316108015906108f05750608a61ffff831611155b806108ff575061ffff821660a7145b8061091c575061ffff821660ac148061091c575061ffff821660ad145b8061093c575060c061ffff83161080159061093c575060c461ffff831611155b8061095c575060bc61ffff83161080159061095c575060bf61ffff831611155b1561097357506002546001600160a01b03166109f8565b61801061ffff83161080159061098f575061801361ffff831611155b806109b1575061802061ffff8316108015906109b1575061802461ffff831611155b806109d3575061803061ffff8316108015906109d3575061803261ffff831611155b156109ea57506003546001600160a01b03166109f8565b506000546001600160a01b03165b806001600160a01b031663a92cb5018e8989888f8f6040518763ffffffff1660e01b8152600401610a2e96959493929190612ad8565b600060405180830381865afa158015610a4b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a739190810190613140565b9097509550600061ffff83166180231480610a93575061ffff8316618024145b1590508015610aae57610aa7868589611479565b6101608901525b600288516002811115610ac357610ac3612890565b148015610ad7575061014088015160001914155b15610af457610ae5886116aa565b610aee8861172e565b50600088525b610afd88610c64565b985050505050505050505b95945050505050565b60006001610b2560a084016080850161329b565b6002811115610b3657610b36612890565b03610ba457610b52610b4d368490038401846132b8565b611208565b6040517f4d616368696e652066696e69736865643a000000000000000000000000000000602082015260318101919091526051015b604051602081830303815290604052805190602001209050919050565b6002610bb660a084016080850161329b565b6002811115610bc757610bc7612890565b03610c1757610bde610b4d368490038401846132b8565b6040517f4d616368696e65206572726f7265643a0000000000000000000000000000000060208201526030810191909152605001610b87565b60405162461bcd60e51b815260206004820152601260248201527f4241445f4d414348494e455f5354415455530000000000000000000000000000604482015260640161042a565b919050565b60008082516002811115610c7a57610c7a612890565b03610dbc576000610ca8610c91846020015161175e565b6101408501516040860151919060001914156117f4565b90506000610cd3610cbc856080015161195e565b61014086015160a0870151919060001914156117f4565b9050600082610ce5866060015161175e565b60c087015160e0808901516101008a01516101208b01516101408c01516101608d01516040517f4d616368696e652072756e6e696e673a00000000000000000000000000000000602082015260308101999099526050890197909752607088018a905260908801959095527fffffffff0000000000000000000000000000000000000000000000000000000092841b831660b088015290831b821660b487015290911b1660b884015260bc83015260dc82015260fc0160408051601f19818403018152919052805160209091012095945050505050565b600182516002811115610dd157610dd1612890565b03610e145760c08201516040517f4d616368696e652066696e69736865643a00000000000000000000000000000060208201526031810191909152605101610b87565b600282516002811115610e2957610e29612890565b03610e6c5760c08201516040517f4d616368696e65206572726f7265643a0000000000000000000000000000000060208201526030810191909152605001610b87565b60405162461bcd60e51b815260206004820152600f60248201527f4241445f4d4143485f5354415455530000000000000000000000000000000000604482015260640161042a565b610ebc612670565b81600080610ecb878785611a02565b9350905060ff8116600003610ee35760009150610f53565b8060ff16600103610ef75760019150610f53565b8060ff16600203610f0b5760029150610f53565b60405162461bcd60e51b815260206004820152601360248201527f554e4b4e4f574e5f4d4143485f53544154555300000000000000000000000000604482015260640161042a565b5060408051606080820183529181019182529081526000602082015260408051606080820183529181019182529081526000602082015260408051808201909152600080825260208201526040805180820190915260608152600060208201526040805180820190915260008082526020820152610fd28b8b89611a38565b97509450610fe18b8b89611b4b565b97509250610ff08b8b89611a38565b97509350610fff8b8b89611ba1565b9750915061100e8b8b89611b4b565b809850819250505060405180610180016040528087600281111561103457611034612890565b8152602081019690965260408601939093526060850193909352608084015260a0830191909152600060c0830181905260e0830181905261010083018190526101208301819052610140830181905261016090920191909152925061109c9050858583611d2d565b60c084019190915290506110b1858583611d49565b63ffffffff90911660e084015290506110cb858583611d49565b63ffffffff90911661010084015290506110e6858583611d49565b63ffffffff9091166101208401529050611101858583611d2d565b6101408401919091529050611117858583611d2d565b61016084019190915291959194509092505050565b611134612751565b8161113d612776565b611145612776565b60005b600260ff821610156111905761115f888886611d2d565b848360ff16600281106111745761117461287a565b60200201919091529350806111888161337a565b915050611148565b5060005b600260ff821610156111eb576111ab888886611dad565b838360ff16600281106111c0576111c061287a565b67ffffffffffffffff90931660209390930201919091529350806111e38161337a565b915050611194565b506040805180820190915291825260208201529590945092505050565b80518051602091820151828401518051908401516040517f476c6f62616c2073746174653a0000000000000000000000000000000000000095810195909552602d850193909352604d8401919091527fffffffffffffffff00000000000000000000000000000000000000000000000060c091821b8116606d85015291901b166075820152600090607d01610b87565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a0810191909152604080516060810182526000808252602082018190529181018290528391906000806000806113128b8b89611d2d565b975095506113218b8b89611e0c565b975094506113308b8b89611d2d565b9750935061133f8b8b89611d2d565b9750925061134e8b8b89611d2d565b9750915061135d8b8b89611d49565b6040805160c081018252988952602089019790975295870194909452506060850191909152608084015263ffffffff1660a083015290969095509350505050565b6040805160208101909152606081528160006113bb868684611a02565b92509050600060ff821667ffffffffffffffff8111156113dd576113dd612864565b604051908082528060200260200182016040528015611406578160200160208202803683370190505b50905060005b8260ff168160ff16101561145d57611425888886611d2d565b838360ff168151811061143a5761143a61287a565b6020026020010181965082815250505080806114559061337a565b91505061140c565b5060405180602001604052808281525093505050935093915050565b60006114c4848461148985611e88565b6040518060400160405280601381526020017f4d6f64756c65206d65726b6c6520747265653a00000000000000000000000000815250611f32565b90505b9392505050565b60608160006114de868684611a02565b9250905060ff811667ffffffffffffffff8111156114fe576114fe612864565b60405190808252806020026020018201604052801561154357816020015b604080518082019091526000808252602082015281526020019060019003908161151c5790505b50925060005b8160ff168110156115c457600080611562898987612055565b955091506115718989876120ae565b809650819250505060405180604001604052808361ffff168152602001828152508684815181106115a4576115a461287a565b6020026020010181905250505080806115bc90613399565b915050611549565b5050935093915050565b60006114c484846115de85612103565b6040518060400160405280601881526020017f496e737472756374696f6e206d65726b6c6520747265653a0000000000000000815250611f32565b6040517f46756e6374696f6e3a00000000000000000000000000000000000000000000006020820152602981018290526000908190604901604051602081830303815290604052805190602001209050610b088585836040518060400160405280601581526020017f46756e6374696f6e206d65726b6c6520747265653a0000000000000000000000815250611f32565b60408101515160a0820151516000198114806116c7575060001982145b156116d457505060029052565b6116e1836080015161195e565b60a08401515260208301516116f59061175e565b60408401515260808301516117109082602082015260609052565b50602091820151808301919091526040805192830190526060825252565b60006117428283610140015160001c6122fb565b61174e57506000919050565b5060001961014090910152600190565b60208101518151515160005b818110156117ed57835161178790611782908361233d565b612375565b6040517f56616c756520737461636b3a00000000000000000000000000000000000000006020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806117e590613399565b91505061176a565b5050919050565b6000600183016118465760405162461bcd60e51b815260206004820152601960248201527f4d554c5449535441434b5f4e4f535441434b5f41435449564500000000000000604482015260640161042a565b811561190c57835160010161189d5760405162461bcd60e51b815260206004820152601760248201527f4d554c5449535441434b5f4e4f535441434b5f4d41494e000000000000000000604482015260640161042a565b83516020808601516040516118ef9392879291017f6d756c7469737461636b3a0000000000000000000000000000000000000000008152600b810193909352602b830191909152604b820152606b0190565b6040516020818303038152906040528051906020012090506114c7565b83516020808601516040517f6d756c7469737461636b3a00000000000000000000000000000000000000000092810192909252602b8201869052604b820192909252606b810191909152608b016118ef565b602081015160005b8251518110156119fc57611996836000015182815181106119895761198961287a565b6020026020010151612392565b6040517f537461636b206672616d6520737461636b3a0000000000000000000000000000602082015260328101919091526052810183905260720160405160208183030381529060405280519060200120915080806119f490613399565b915050611966565b50919050565b600081848482818110611a1757611a1761287a565b919091013560f81c9250819050611a2d81613399565b915050935093915050565b604080516060808201835291810191825290815260006020820152816000611a61868684611d2d565b925090506000611a728787856120ae565b9350905060008167ffffffffffffffff811115611a9157611a91612864565b604051908082528060200260200182016040528015611ad657816020015b6040805180820190915260008082526020820152815260200190600190039081611aaf5790505b50905060005b8151811015611b2457611af089898761242b565b838381518110611b0257611b0261287a565b6020026020010181975082905250508080611b1c90613399565b915050611adc565b50604080516060810182529081019182529081526020810192909252509590945092505050565b6040805180820190915260008082526020820152816000611b6d868684611d2d565b925090506000611b7e878785611d2d565b604080518082019091529384526020840191909152919791965090945050505050565b604080518082019091526060815260006020820152816000611bc4868684611d2d565b925090506060868684818110611bdc57611bdc61287a565b909101357fff0000000000000000000000000000000000000000000000000000000000000016159050611ca25782611c1381613399565b604080516001808252818301909252919550909150816020015b6040805160c08101825260006080820181815260a083018290528252602080830182905292820181905260608201528252600019909201910181611c2d579050509050611c7b878785612536565b82600081518110611c8e57611c8e61287a565b602002602001018195508290525050611d0c565b82611cac81613399565b60408051600080825260208201909252919550909150611d08565b6040805160c08101825260006080820181815260a083018290528252602080830182905292820181905260608201528252600019909201910181611cc75790505b5090505b60405180604001604052808281526020018381525093505050935093915050565b60008181611d3c8686846120ae565b9097909650945050505050565b600081815b6004811015611da45760088363ffffffff16901b9250858583818110611d7657611d7661287a565b919091013560f81c93909317925081611d8e81613399565b9250508080611d9c90613399565b915050611d4e565b50935093915050565b600081815b6008811015611da45760088367ffffffffffffffff16901b9250858583818110611dde57611dde61287a565b919091013560f81c93909317925081611df681613399565b9250508080611e0490613399565b915050611db2565b60408051606081018252600080825260208201819052918101919091528160008080611e39888886611dad565b94509250611e48888886611dad565b94509150611e57888886611d2d565b6040805160608101825267ffffffffffffffff96871681529490951660208501529383015250969095509350505050565b60008160000151611e9c83602001516125ee565b6040808501516060860151608087015160a08801519351610b87969594906020017f4d6f64756c653a0000000000000000000000000000000000000000000000000081526007810196909652602786019490945260478501929092526067840152608783015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660a782015260ab0190565b8160005b855151811015611ffe5784600116600003611f9a57828287600001518381518110611f6357611f6361287a565b6020026020010151604051602001611f7d939291906133b3565b604051602081830303815290604052805190602001209150611fe5565b8286600001518281518110611fb157611fb161287a565b602002602001015183604051602001611fcc939291906133b3565b6040516020818303038152906040528051906020012091505b60019490941c9380611ff681613399565b915050611f36565b50831561204d5760405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f52540000000000000000000000000000000000604482015260640161042a565b949350505050565b600081815b6002811015611da45760088361ffff16901b92508585838181106120805761208061287a565b919091013560f81c9390931792508161209881613399565b92505080806120a690613399565b91505061205a565b600081815b6020811015611da457600883901b92508585838181106120d5576120d561287a565b919091013560f81c939093179250816120ed81613399565b92505080806120fb90613399565b9150506120b3565b6000808251602261211491906133ea565b61211f90600e6128bc565b67ffffffffffffffff81111561213757612137612864565b6040519080825280601f01601f191660200182016040528015612161576020820181803683370190505b5090507f496e737472756374696f6e733a0000000000000000000000000000000000000060208201526000600d9050835160f81b8282815181106121a7576121a761287a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350806121e081613399565b91505060005b84518110156122eb5760008582815181106122035761220361287a565b602002602001015190506008816000015161ffff16901c60f81b84848151811061222f5761222f61287a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350805160f81b8461226f8560016128bc565b8151811061227f5761227f61287a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506122b96002846128bc565b60208083015186830182018190529194506122d490856128bc565b9350505080806122e390613399565b9150506121e6565b5050805160209091012092915050565b6000606082901c1561230f57506000610348565b5063ffffffff818116610120840152602082901c811661010084015260409190911c1660e090910152600190565b604080518082019091526000808252602082015282518051839081106123655761236561287a565b6020026020010151905092915050565b600081600001518260200151604051602001610b87929190613401565b60006123a18260000151612375565b602080840151604080860151606087015191517f537461636b206672616d653a000000000000000000000000000000000000000094810194909452602c840194909452604c8301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e093841b8116606c840152921b9091166070820152607401610b87565b60408051808201909152600080825260208201528160008585838181106124545761245461287a565b919091013560f81c915082905061246a81613399565b925050612475600690565b600681111561248657612486612890565b60ff168160ff1611156124db5760405162461bcd60e51b815260206004820152600e60248201527f4241445f56414c55455f54595045000000000000000000000000000000000000604482015260640161042a565b60006124e88787856120ae565b809450819250505060405180604001604052808360ff16600681111561251057612510612890565b600681111561252157612521612890565b81526020018281525093505050935093915050565b6040805160c08101825260006080820181815260a0830182905282526020808301829052828401829052606083018290528351808501909452818452830152908290600080600061258889898761242b565b95509350612597898987611d2d565b955092506125a6898987611d49565b955091506125b5898987611d49565b60408051608081018252968752602087019590955263ffffffff9384169486019490945290911660608401525090969095509350505050565b805160208083015160408085015190517f4d656d6f72793a00000000000000000000000000000000000000000000000000938101939093527fffffffffffffffff00000000000000000000000000000000000000000000000060c094851b811660278501529190931b16602f8201526037810191909152600090605701610b87565b60408051610180810190915280600081526020016126a560408051606080820183529181019182529081526000602082015290565b81526040805180820182526000808252602080830191909152830152016126e360408051606080820183529181019182529081526000602082015290565b8152602001612708604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b6040518060400160405280612764612776565b8152602001612771612776565b905290565b60405180604001604052806002906020820280368337509192915050565b600080604083850312156127a757600080fd5b50508035926020909101359150565b600080600080600085870360c08112156127cf57600080fd5b60608112156127dd57600080fd5b50859450606086013593506080860135925060a086013567ffffffffffffffff8082111561280a57600080fd5b818801915088601f83011261281e57600080fd5b81358181111561282d57600080fd5b89602082850101111561283f57600080fd5b9699959850939650602001949392505050565b600060a082840312156119fc57600080fd5b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610348576103486128a6565b634e487b7160e01b600052601260045260246000fd5b600063ffffffff808416806128fc576128fc6128cf565b92169190910492915050565b600063ffffffff8084168061291f5761291f6128cf565b92169190910692915050565b6000808585111561293b57600080fd5b8386111561294857600080fd5b5050820193919092039150565b63ffffffff818116838216019080821115612972576129726128a6565b5092915050565b61ffff818116838216019080821115612972576129726128a6565b600381106129a4576129a4612890565b9052565b8051600781106129ba576129ba612890565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015612a1757612a038286516129a8565b9382019360019390930192908501906129f0565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015612a96578451612a628582516129a8565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101612a4d565b509687015197909601969096525093949350505050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b60006101e08835835260208901356001600160a01b038116808214612afc57600080fd5b806020860152505060408901356040840152806060840152612b218184018951612994565b5060208701516101c080610200850152612b3f6103a08501836129c7565b60408a0151805161022087015260208101516102408701529092505060608901517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe208086850301610260870152612b9684836129c7565b935060808b01519150808685030161028087015250612bb58382612a2b565b92505060a0890151612bd56102a086018280518252602090810151910152565b5060c08901516102e085015260e089015163ffffffff81166103008601525061010089015163ffffffff81166103208601525061012089015163ffffffff811661034086015250610140890151610360850152610160890151610380850152612ca1608085018980518252602081015167ffffffffffffffff80825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b865161ffff1661018085015260208701516101a085015283820390840152612cca818587612aad565b9998505050505050505050565b6040805190810167ffffffffffffffff81118282101715612cfa57612cfa612864565b60405290565b6040516020810167ffffffffffffffff81118282101715612cfa57612cfa612864565b6040516080810167ffffffffffffffff81118282101715612cfa57612cfa612864565b60405160c0810167ffffffffffffffff81118282101715612cfa57612cfa612864565b6040516060810167ffffffffffffffff81118282101715612cfa57612cfa612864565b604051610180810167ffffffffffffffff81118282101715612cfa57612cfa612864565b604051601f8201601f1916810167ffffffffffffffff81118282101715612dd957612dd9612864565b604052919050565b60038110612dee57600080fd5b50565b8051610c5f81612de1565b600067ffffffffffffffff821115612e1657612e16612864565b5060051b60200190565b600060408284031215612e3257600080fd5b612e3a612cd7565b9050815160078110612e4b57600080fd5b808252506020820151602082015292915050565b60006040808385031215612e7257600080fd5b612e7a612cd7565b9150825167ffffffffffffffff80821115612e9457600080fd5b81850191506020808388031215612eaa57600080fd5b612eb2612d00565b835183811115612ec157600080fd5b80850194505087601f850112612ed657600080fd5b83519250612eeb612ee684612dfc565b612db0565b83815260069390931b84018201928281019089851115612f0a57600080fd5b948301945b84861015612f3057612f218a87612e20565b82529486019490830190612f0f565b8252508552948501519484019490945250909392505050565b600060408284031215612f5b57600080fd5b612f63612cd7565b9050815181526020820151602082015292915050565b805163ffffffff81168114610c5f57600080fd5b60006040808385031215612fa057600080fd5b612fa8612cd7565b9150825167ffffffffffffffff811115612fc157600080fd5b8301601f81018513612fd257600080fd5b80516020612fe2612ee683612dfc565b82815260a0928302840182019282820191908985111561300157600080fd5b948301945b8486101561306a5780868b03121561301e5760008081fd5b613026612d23565b6130308b88612e20565b815287870151858201526060613047818901612f79565b8983015261305760808901612f79565b9082015283529485019491830191613006565b50808752505080860151818601525050505092915050565b67ffffffffffffffff81168114612dee57600080fd5b60008183036101008112156130ac57600080fd5b6130b4612d46565b9150825182526060601f19820112156130cc57600080fd5b506130d5612d69565b60208301516130e381613082565b815260408301516130f381613082565b8060208301525060608301516040820152806020830152506080820151604082015260a0820151606082015260c0820151608082015261313560e08301612f79565b60a082015292915050565b60008061012080848603121561315557600080fd5b835167ffffffffffffffff8082111561316d57600080fd5b908501906101c0828803121561318257600080fd5b61318a612d8c565b61319383612df1565b81526020830151828111156131a757600080fd5b6131b389828601612e5f565b6020830152506131c68860408501612f49565b60408201526080830151828111156131dd57600080fd5b6131e989828601612e5f565b60608301525060a08301518281111561320157600080fd5b61320d89828601612f8d565b6080830152506132208860c08501612f49565b60a082015261010091508183015160c082015261323e848401612f79565b60e0820152610140613251818501612f79565b838301526101609250613265838501612f79565b8583015261018084015181830152506101a083015182820152809550505050506132928460208501613098565b90509250929050565b6000602082840312156132ad57600080fd5b81356114c781612de1565b6000608082840312156132ca57600080fd5b6132d2612cd7565b83601f8401126132e157600080fd5b6132e9612cd7565b8060408501868111156132fb57600080fd5b855b818110156133155780358452602093840193016132fd565b5081845286605f87011261332857600080fd5b613330612cd7565b9250829150608086018781111561334657600080fd5b8082101561336b57813561335981613082565b84526020938401939190910190613346565b50506020830152509392505050565b600060ff821660ff8103613390576133906128a6565b60010192915050565b600060001982036133ac576133ac6128a6565b5060010190565b6000845160005b818110156133d457602081880181015185830152016133ba565b5091909101928352506020820152604001919050565b8082028115828204841417610348576103486128a6565b7f56616c75653a0000000000000000000000000000000000000000000000000000815260006007841061343657613436612890565b5060f89290921b600683015260078201526027019056fea2646970667358221220534ee31bc0d4857954a59e241c4fc6f8d1b5f8c5f6e707ea2df2f0c067c6534464736f6c63430008110033",
}

// OneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryMetaData.ABI instead.
var OneStepProofEntryABI = OneStepProofEntryMetaData.ABI

// OneStepProofEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryMetaData.Bin instead.
var OneStepProofEntryBin = OneStepProofEntryMetaData.Bin

// DeployOneStepProofEntry deploys a new Ethereum contract, binding an instance of OneStepProofEntry to it.
func DeployOneStepProofEntry(auth *bind.TransactOpts, backend bind.ContractBackend, prover0_ common.Address, proverMem_ common.Address, proverMath_ common.Address, proverHostIo_ common.Address) (common.Address, *types.Transaction, *OneStepProofEntry, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryBin), backend, prover0_, proverMem_, proverMath_, proverHostIo_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// OneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntry struct {
	OneStepProofEntryCaller     // Read-only binding to the contract
	OneStepProofEntryTransactor // Write-only binding to the contract
	OneStepProofEntryFilterer   // Log filterer for contract events
}

// OneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntrySession struct {
	Contract     *OneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryCallerSession struct {
	Contract *OneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryTransactorSession struct {
	Contract     *OneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryRaw struct {
	Contract *OneStepProofEntry // Generic contract binding to access the raw methods on
}

// OneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryCallerRaw struct {
	Contract *OneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactorRaw struct {
	Contract *OneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntry creates a new instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*OneStepProofEntry, error) {
	contract, err := bindOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryCaller creates a new read-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryCaller, error) {
	contract, err := bindOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryCaller{contract: contract}, nil
}

// NewOneStepProofEntryTransactor creates a new write-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryTransactor, error) {
	contract, err := bindOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryTransactor{contract: contract}, nil
}

// NewOneStepProofEntryFilterer creates a new log filterer instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryFilterer, error) {
	contract, err := bindOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryFilterer{contract: contract}, nil
}

// bindOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.OneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetMachineHash(opts *bind.CallOpts, execState ExecutionState) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getMachineHash", execState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetMachineHash(&_OneStepProofEntry.CallOpts, execState)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetMachineHash(&_OneStepProofEntry.CallOpts, execState)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetStartMachineHash(&_OneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetStartMachineHash(&_OneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) Prover0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "prover0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverHostIo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverHostIo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMath(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMath")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// OneStepProofEntryLibMetaData contains all meta data concerning the OneStepProofEntryLib contract.
var OneStepProofEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220078c5b1f3558c0305a3ffc59d77d714be74dc68855d5f0bb1c7a94520d0d62ce64736f6c63430008110033",
}

// OneStepProofEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryLibMetaData.ABI instead.
var OneStepProofEntryLibABI = OneStepProofEntryLibMetaData.ABI

// OneStepProofEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryLibMetaData.Bin instead.
var OneStepProofEntryLibBin = OneStepProofEntryLibMetaData.Bin

// DeployOneStepProofEntryLib deploys a new Ethereum contract, binding an instance of OneStepProofEntryLib to it.
func DeployOneStepProofEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofEntryLib, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// OneStepProofEntryLib is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntryLib struct {
	OneStepProofEntryLibCaller     // Read-only binding to the contract
	OneStepProofEntryLibTransactor // Write-only binding to the contract
	OneStepProofEntryLibFilterer   // Log filterer for contract events
}

// OneStepProofEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntryLibSession struct {
	Contract     *OneStepProofEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryLibCallerSession struct {
	Contract *OneStepProofEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OneStepProofEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryLibTransactorSession struct {
	Contract     *OneStepProofEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryLibRaw struct {
	Contract *OneStepProofEntryLib // Generic contract binding to access the raw methods on
}

// OneStepProofEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCallerRaw struct {
	Contract *OneStepProofEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactorRaw struct {
	Contract *OneStepProofEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntryLib creates a new instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLib(address common.Address, backend bind.ContractBackend) (*OneStepProofEntryLib, error) {
	contract, err := bindOneStepProofEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryLibCaller creates a new read-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryLibCaller, error) {
	contract, err := bindOneStepProofEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibCaller{contract: contract}, nil
}

// NewOneStepProofEntryLibTransactor creates a new write-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryLibTransactor, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibTransactor{contract: contract}, nil
}

// NewOneStepProofEntryLibFilterer creates a new log filterer instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryLibFilterer, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibFilterer{contract: contract}, nil
}

// bindOneStepProofEntryLib binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProver0MetaData contains all meta data concerning the OneStepProver0 contract.
var OneStepProver0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50613458806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a92cb50114610030575b600080fd5b61004361003e366004612956565b61005a565b604051610051929190612b86565b60405180910390f35b610062612822565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526100b5876130d6565b91506100c636879003870187613214565b905060006100d760208701876132b6565b905061290361ffff82166100ee57506105146104f6565b60001961ffff831601610104575061051f6104f6565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff161ffff83160161013857506105266104f6565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff061ffff83160161016c575061054d6104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ff761ffff8316016101a057506106756104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ff561ffff8316016101d4575061077f6104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ff461ffff831601610208575061089b6104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ff661ffff83160161023c5750610aa06104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffef61ffff8316016102705750610bec6104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ffd61ffff8316016102a457506110866104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ffc61ffff8316016102d857506110f66104f6565b601f1961ffff8316016102ee57506111846104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdf61ffff83160161032257506111c66104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd61ffff831601610356575061120b6104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc61ffff83160161038a57506112336104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ffe61ffff8316016103be57506112636104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe661ffff8316016103f257506113006104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe561ffff831601610426575061130d6104f6565b604161ffff8316108015906104405750604461ffff831611155b1561044e575061137c6104f6565b61ffff82166180051480610467575061ffff8216618006145b1561047557506114ed6104f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7ff861ffff8316016104a957506115be6104f6565b60405162461bcd60e51b815260206004820152600e60248201527f494e56414c49445f4f50434f444500000000000000000000000000000000000060448201526064015b60405180910390fd5b61050784848989898663ffffffff16565b5050965096945050505050565b505060029092525050565b5050505050565b600061053586608001516115cd565b80519091506105459087906116ed565b505050505050565b61056461055986611765565b6020870151906117e7565b600061057386608001516117f3565b90506105be6105b38260400151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6020880151906117e7565b6105fc6105b38260600151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b602084013563ffffffff811681146106565760405162461bcd60e51b815260206004820152600d60248201527f4241445f43414c4c5f444154410000000000000000000000000000000000000060448201526064016104ed565b63ffffffff166101008701525050600061012090940193909352505050565b61068161055986611765565b6106bf6105598660e00151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6106fd6105598560a00151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6020808401359081901c604082901c156107595760405162461bcd60e51b815260206004820152601a60248201527f4241445f43524f53535f4d4f44554c455f43414c4c5f4441544100000000000060448201526064016104ed565b63ffffffff90811660e08801521661010086015250506000610120909301929092525050565b61078b61055986611765565b600061079a86608001516117f3565b90506107da6105b38260400151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6108186105b38260600151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6020808501359081901c604082901c156108745760405162461bcd60e51b815260206004820152601a60248201527f4241445f43524f53535f4d4f44554c455f43414c4c5f4441544100000000000060448201526064016104ed565b63ffffffff90811660e0890152166101008701525050600061012090940193909352505050565b60008360200135905060006108bb6108b68860200151611898565b6118b7565b90506109106040805160c0810182526000808252825160608101845281815260208181018390529381019190915290918201908152600060208201819052604082018190526060820181905260809091015290565b604080516020810190915260608152600061092c878783611974565b909350905061093c878783611a7a565b6101608c0151919350915061095c8363ffffffff808816908790611b5516565b146109cf5760405162461bcd60e51b815260206004820152602260248201527f43524f53535f4d4f44554c455f494e5445524e414c5f4d4f44554c45535f524f60448201527f4f5400000000000000000000000000000000000000000000000000000000000060648201526084016104ed565b6109e66109db8b611765565b60208c0151906117e7565b610a246109db8b60e00151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b610a626109db8a60a00151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b63ffffffff841660e08b015260a0830151610a7d90866132f0565b63ffffffff166101008b0152505060006101209098019790975250505050505050565b610aac61055986611765565b610aea6105598660e00151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b610b286105598560a00151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6000610b3786608001516117f3565b9050806060015163ffffffff16600003610b5557506002855261051f565b602084013563ffffffff81168114610baf5760405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f4441544100000060448201526064016104ed565b604082015163ffffffff1660e08801526060820151610bcf9082906132f0565b63ffffffff16610100880152505060006101208601525050505050565b600080610bff6108b68860200151611898565b9050600080600080806000610c206040518060200160405280606081525090565b610c2b8b8b87611baa565b95509350610c3a8b8b87611c12565b9096509450610c4a8b8b87611c2e565b95509250610c598b8b87611baa565b95509150610c688b8b87611c12565b9097509450610c788b8b87611a7a565b6040517f43616c6c20696e6469726563743a00000000000000000000000000000000000060208201527fffffffffffffffff00000000000000000000000000000000000000000000000060c088901b16602e8201526036810189905290965090915060009060560160408051601f19818403018152919052805160209182012091508d01358114610d4b5760405162461bcd60e51b815260206004820152601660248201527f4241445f43414c4c5f494e4449524543545f444154410000000000000000000060448201526064016104ed565b610d628267ffffffffffffffff871686868c611c64565b90508d604001518114610db75760405162461bcd60e51b815260206004820152600f60248201527f4241445f5441424c45535f524f4f54000000000000000000000000000000000060448201526064016104ed565b8267ffffffffffffffff168963ffffffff1610610de257505060028d525061051f9650505050505050565b50505050506000610e03604080518082019091526000808252602082015290565b604080516020810190915260608152610e1d8a8a86611c12565b94509250610e2c8a8a86611d58565b94509150610e3b8a8a86611a7a565b945090506000610e588263ffffffff808b169087908790611e6316565b9050868114610ea95760405162461bcd60e51b815260206004820152601160248201527f4241445f454c454d454e54535f524f4f5400000000000000000000000000000060448201526064016104ed565b858414610ed9578d60025b90816002811115610ec757610ec7612a57565b8152505050505050505050505061051f565b600483516006811115610eee57610eee612a57565b03610efb578d6002610eb4565b600583516006811115610f1057610f10612a57565b03610f76576020830151985063ffffffff89168914610f715760405162461bcd60e51b815260206004820152601560248201527f4241445f46554e435f5245465f434f4e54454e5453000000000000000000000060448201526064016104ed565b610fbe565b60405162461bcd60e51b815260206004820152600d60248201527f4241445f454c454d5f545950450000000000000000000000000000000000000060448201526064016104ed565b5050505050505050610fd26105b387611765565b6000610fe187608001516117f3565b905061102c6110218260400151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6020890151906117e7565b61106a6110218260600151604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b5063ffffffff1661010086015260006101208601525050505050565b602083013563ffffffff811681146110e05760405162461bcd60e51b815260206004820152600d60248201527f4241445f43414c4c5f444154410000000000000000000000000000000000000060448201526064016104ed565b63ffffffff166101209095019490945250505050565b60006111086108b68760200151611898565b905063ffffffff81161561054557602084013563ffffffff811681146111705760405162461bcd60e51b815260206004820152600d60248201527f4241445f43414c4c5f444154410000000000000000000000000000000000000060448201526064016104ed565b63ffffffff16610120870152505050505050565b600061119386608001516117f3565b905060006111ab826020015186602001358686611f0e565b60208801519091506111bd90826117e7565b50505050505050565b60006111d58660200151611898565b905060006111e687608001516117f3565b90506111fd81602001518660200135848787611fd6565b602090910152505050505050565b6000611221856000015185602001358585611f0e565b602087015190915061054590826117e7565b60006112428660200151611898565b905061125985600001518560200135838686611fd6565b9094525050505050565b60006112728660200151611898565b905060006112838760200151611898565b905060006112948860200151611898565b905060006040518060800160405280838152602001886020013560001b81526020016112bf856118b7565b63ffffffff1681526020016112d3866118b7565b63ffffffff1681525090506112f5818a608001516120a090919063ffffffff16565b505050505050505050565b6105458560200151611898565b600061131f6108b68760200151611898565b905060006113308760200151611898565b905060006113418860200151611898565b905063ffffffff83161561136357602088015161135e90826117e7565b611372565b602088015161137290836117e7565b5050505050505050565b600061138b60208501856132b6565b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbf61ffff8316016113c2575060006114a3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbe61ffff8316016113f5575060016114a3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbd61ffff831601611428575060026114a3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc61ffff83160161145b575060036114a3565b60405162461bcd60e51b815260206004820152601960248201527f434f4e53545f505553485f494e56414c49445f4f50434f44450000000000000060448201526064016104ed565b6111bd60405180604001604052808360068111156114c3576114c3612a57565b8152602001876020013567ffffffffffffffff1681525088602001516117e790919063ffffffff16565b604080518082019091526000808252602082015261800561151160208601866132b6565b61ffff160361153e576115278660200151611898565b606087015190915061153990826117e7565b610545565b61800661154e60208601866132b6565b61ffff1603611576576115648660600151611898565b602087015190915061153990826117e7565b60405162461bcd60e51b815260206004820152601c60248201527f4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f44450000000060448201526064016104ed565b600061122186602001516121ae565b6040805160c08101825260006080820181815260a0830182905282526020820181905291810182905260608101919091528151516001146116505760405162461bcd60e51b815260206004820152601160248201527f4241445f57494e444f575f4c454e47544800000000000000000000000000000060448201526064016104ed565b8151805160009061166357611663613314565b60200260200101519050600067ffffffffffffffff81111561168757611687612d1c565b6040519080825280602002602001820160405280156116e657816020015b6040805160c08101825260006080820181815260a0830182905282526020808301829052928201819052606082015282526000199092019101816116a55790505b5090915290565b60048151600681111561170257611702612a57565b03611725578160025b9081600281111561171e5761171e612a57565b9052505050565b60068151600681111561173a5761173a612a57565b146117475781600261170b565b6117558282602001516121dc565b6117615781600261170b565b5050565b60408051808201909152600080825260208201526117e18261012001518361010001518460e001516040805180820190915260008082526020820152506040805180820182526006815263ffffffff94909416602093841b67ffffffff00000000161791901b6bffffffff000000000000000016179082015290565b92915050565b8151611761908261221e565b6040805160c08101825260006080820181815260a0830182905282526020820181905291810182905260608101919091528151516001146118765760405162461bcd60e51b815260206004820152601160248201527f4241445f57494e444f575f4c454e47544800000000000000000000000000000060448201526064016104ed565b8151805160009061188957611889613314565b60200260200101519050919050565b604080518082019091526000808252602082015281516117e1906122e8565b602081015160009081835160068111156118d3576118d3612a57565b146119205760405162461bcd60e51b815260206004820152600760248201527f4e4f545f4933320000000000000000000000000000000000000000000000000060448201526064016104ed565b64010000000081106117e15760405162461bcd60e51b815260206004820152600760248201527f4241445f4933320000000000000000000000000000000000000000000000000060448201526064016104ed565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a0810191909152604080516060810182526000808252602082018190529181018290528391906000806000806119ee8b8b89611c12565b975095506119fd8b8b896123f2565b97509450611a0c8b8b89611c12565b97509350611a1b8b8b89611c12565b97509250611a2a8b8b89611c12565b97509150611a398b8b8961246e565b6040805160c081018252988952602089019790975295870194909452506060850191909152608084015263ffffffff1660a083015290969095509350505050565b604080516020810190915260608152816000611a97868684611c2e565b92509050600060ff821667ffffffffffffffff811115611ab957611ab9612d1c565b604051908082528060200260200182016040528015611ae2578160200160208202803683370190505b50905060005b8260ff168160ff161015611b3957611b01888886611c12565b838360ff1681518110611b1657611b16613314565b602002602001018196508281525050508080611b319061332a565b915050611ae8565b5060405180602001604052808281525093505050935093915050565b6000611ba08484611b65856124c9565b6040518060400160405280601381526020017f4d6f64756c65206d65726b6c6520747265653a00000000000000000000000000815250612590565b90505b9392505050565b600081815b6008811015611c095760088367ffffffffffffffff16901b9250858583818110611bdb57611bdb613314565b919091013560f81c93909317925081611bf381613349565b9250508080611c0190613349565b915050611baf565b50935093915050565b60008181611c218686846126ab565b9097909650945050505050565b600081848482818110611c4357611c43613314565b919091013560f81c9250819050611c5981613349565b915050935093915050565b6040517f5461626c653a000000000000000000000000000000000000000000000000000060208201527fff0000000000000000000000000000000000000000000000000000000000000060f885901b1660268201527fffffffffffffffff00000000000000000000000000000000000000000000000060c084901b166027820152602f81018290526000908190604f01604051602081830303815290604052805190602001209050611d4d8787836040518060400160405280601281526020017f5461626c65206d65726b6c6520747265653a0000000000000000000000000000815250612590565b979650505050505050565b6040805180820190915260008082526020820152816000858583818110611d8157611d81613314565b919091013560f81c9150829050611d9781613349565b925050611da2600690565b6006811115611db357611db3612a57565b60ff168160ff161115611e085760405162461bcd60e51b815260206004820152600e60248201527f4241445f56414c55455f5459504500000000000000000000000000000000000060448201526064016104ed565b6000611e158787856126ab565b809450819250505060405180604001604052808360ff166006811115611e3d57611e3d612a57565b6006811115611e4e57611e4e612a57565b81526020018281525093505050935093915050565b60008083611e7084612700565b6040517f5461626c6520656c656d656e743a0000000000000000000000000000000000006020820152602e810192909252604e820152606e01604051602081830303815290604052805190602001209050611f028686836040518060400160405280601a81526020017f5461626c6520656c656d656e74206d65726b6c6520747265653a000000000000815250612590565b9150505b949350505050565b60408051808201909152600080825260208201526000611f3e604080518082019091526000808252602082015290565b604080516020810190915260608152611f58868685611d58565b93509150611f67868685611a7a565b935090506000611f7882898561271d565b9050888114611fc95760405162461bcd60e51b815260206004820152601160248201527f57524f4e475f4d45524b4c455f524f4f5400000000000000000000000000000060448201526064016104ed565b5090979650505050505050565b6000611ff2604080518082019091526000808252602082015290565b600061200a6040518060200160405280606081525090565b612015868684611d58565b9093509150612025868684611a7a565b925090506000612036828a8661271d565b90508981146120875760405162461bcd60e51b815260206004820152601160248201527f57524f4e475f4d45524b4c455f524f4f5400000000000000000000000000000060448201526064016104ed565b612092828a8a61271d565b9a9950505050505050505050565b8151516000906120b1906001613363565b67ffffffffffffffff8111156120c9576120c9612d1c565b60405190808252806020026020018201604052801561212857816020015b6040805160c08101825260006080820181815260a0830182905282526020808301829052928201819052606082015282526000199092019101816120e75790505b50905060005b83515181101561218457835180518290811061214c5761214c613314565b602002602001015182828151811061216657612166613314565b6020026020010181905250808061217c90613349565b91505061212e565b5081818460000151518151811061219d5761219d613314565b602090810291909101015290915250565b604080518082019091526000808252602082015281515151611ba36121d4600183613376565b845190612768565b6000606082901c156121f0575060006117e1565b5063ffffffff818116610120840152602082901c811661010084015260409190911c1660e090910152600190565b81515160009061222f906001613363565b67ffffffffffffffff81111561224757612247612d1c565b60405190808252806020026020018201604052801561228c57816020015b60408051808201909152600080825260208201528152602001906001900390816122655790505b50905060005b8351518110156121845783518051829081106122b0576122b0613314565b60200260200101518282815181106122ca576122ca613314565b602002602001018190525080806122e090613349565b915050612292565b60408051808201909152600080825260208201528151805161230c90600190613376565b8151811061231c5761231c613314565b602002602001015190506000600183600001515161233a9190613376565b67ffffffffffffffff81111561235257612352612d1c565b60405190808252806020026020018201604052801561239757816020015b60408051808201909152600080825260208201528152602001906001900390816123705790505b50905060005b81518110156116e65783518051829081106123ba576123ba613314565b60200260200101518282815181106123d4576123d4613314565b602002602001018190525080806123ea90613349565b91505061239d565b6040805160608101825260008082526020820181905291810191909152816000808061241f888886611baa565b9450925061242e888886611baa565b9450915061243d888886611c12565b6040805160608101825267ffffffffffffffff96871681529490951660208501529383015250969095509350505050565b600081815b6004811015611c095760088363ffffffff16901b925085858381811061249b5761249b613314565b919091013560f81c939093179250816124b381613349565b92505080806124c190613349565b915050612473565b600081600001516124dd83602001516127a0565b6040808501516060860151608087015160a08801519351612573969594906020017f4d6f64756c653a0000000000000000000000000000000000000000000000000081526007810196909652602786019490945260478501929092526067840152608783015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660a782015260ab0190565b604051602081830303815290604052805190602001209050919050565b8160005b85515181101561265c57846001166000036125f8578282876000015183815181106125c1576125c1613314565b60200260200101516040516020016125db93929190613389565b604051602081830303815290604052805190602001209150612643565b828660000151828151811061260f5761260f613314565b60200260200101518360405160200161262a93929190613389565b6040516020818303038152906040528051906020012091505b60019490941c938061265481613349565b915050612594565b508315611f065760405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f5254000000000000000000000000000000000060448201526064016104ed565b600081815b6020811015611c0957600883901b92508585838181106126d2576126d2613314565b919091013560f81c939093179250816126ea81613349565b92505080806126f890613349565b9150506126b0565b6000816000015182602001516040516020016125739291906133c0565b6000611ba0848461272d85612700565b6040518060400160405280601281526020017f56616c7565206d65726b6c6520747265653a0000000000000000000000000000815250612590565b6040805180820190915260008082526020820152825180518390811061279057612790613314565b6020026020010151905092915050565b805160208083015160408085015190517f4d656d6f72793a00000000000000000000000000000000000000000000000000938101939093527fffffffffffffffff00000000000000000000000000000000000000000000000060c094851b811660278501529190931b16602f8201526037810191909152600090605701612573565b604080516101808101909152806000815260200161285760408051606080820183529181019182529081526000602082015290565b815260408051808201825260008082526020808301919091528301520161289560408051606080820183529181019182529081526000602082015290565b81526020016128ba604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b61290b61340c565b565b60008083601f84011261291f57600080fd5b50813567ffffffffffffffff81111561293757600080fd5b60208301915083602082850101111561294f57600080fd5b9250929050565b6000806000806000808688036101e081121561297157600080fd5b606081121561297f57600080fd5b879650606088013567ffffffffffffffff8082111561299d57600080fd5b818a0191506101c080838d0312156129b457600080fd5b8298506101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850112156129e857600080fd5b60808b01975060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8085011215612a1e57600080fd5b6101808b0196508a0135925080831115612a3757600080fd5b5050612a4589828a0161290d565b979a9699509497509295939492505050565b634e487b7160e01b600052602160045260246000fd5b60038110612a7d57612a7d612a57565b9052565b805160078110612a9357612a93612a57565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015612af057612adc828651612a81565b938201936001939093019290850190612ac9565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015612b6f578451612b3b858251612a81565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101612b26565b509687015197909601969096525093949350505050565b6000610120808352612b9b8184018651612a6d565b60208501516101c06101408181870152612bb96102e0870184612aa0565b92506040880151610160612bd98189018380518252602090810151910152565b60608a015191507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee080898703016101a08a0152612c168684612aa0565b955060808b015192508089870301858a015250612c338583612b04565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c0860152509150611ba39050602083018480518252602081015167ffffffffffffffff80825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715612d5557612d55612d1c565b60405290565b6040516020810167ffffffffffffffff81118282101715612d5557612d55612d1c565b6040516080810167ffffffffffffffff81118282101715612d5557612d55612d1c565b604051610180810167ffffffffffffffff81118282101715612d5557612d55612d1c565b60405160c0810167ffffffffffffffff81118282101715612d5557612d55612d1c565b6040516060810167ffffffffffffffff81118282101715612d5557612d55612d1c565b604051601f8201601f1916810167ffffffffffffffff81118282101715612e3457612e34612d1c565b604052919050565b803560038110612e4b57600080fd5b919050565b600067ffffffffffffffff821115612e6a57612e6a612d1c565b5060051b60200190565b600060408284031215612e8657600080fd5b612e8e612d32565b9050813560078110612e9f57600080fd5b808252506020820135602082015292915050565b60006040808385031215612ec657600080fd5b612ece612d32565b9150823567ffffffffffffffff80821115612ee857600080fd5b81850191506020808388031215612efe57600080fd5b612f06612d5b565b833583811115612f1557600080fd5b80850194505087601f850112612f2a57600080fd5b83359250612f3f612f3a84612e50565b612e0b565b83815260069390931b84018201928281019089851115612f5e57600080fd5b948301945b84861015612f8457612f758a87612e74565b82529486019490830190612f63565b8252508552948501359484019490945250909392505050565b600060408284031215612faf57600080fd5b612fb7612d32565b9050813581526020820135602082015292915050565b803563ffffffff81168114612e4b57600080fd5b60006040808385031215612ff457600080fd5b612ffc612d32565b9150823567ffffffffffffffff81111561301557600080fd5b8301601f8101851361302657600080fd5b80356020613036612f3a83612e50565b82815260a0928302840182019282820191908985111561305557600080fd5b948301945b848610156130be5780868b0312156130725760008081fd5b61307a612d7e565b6130848b88612e74565b81528787013585820152606061309b818901612fcd565b898301526130ab60808901612fcd565b908201528352948501949183019161305a565b50808752505080860135818601525050505092915050565b60006101c082360312156130e957600080fd5b6130f1612da1565b6130fa83612e3c565b8152602083013567ffffffffffffffff8082111561311757600080fd5b61312336838701612eb3565b60208401526131353660408701612f9d565b6040840152608085013591508082111561314e57600080fd5b61315a36838701612eb3565b606084015260a085013591508082111561317357600080fd5b5061318036828601612fe1565b6080830152506131933660c08501612f9d565b60a08201526101008084013560c08301526101206131b2818601612fcd565b60e08401526101406131c5818701612fcd565b8385015261016092506131d9838701612fcd565b91840191909152610180850135908301526101a090930135928101929092525090565b803567ffffffffffffffff81168114612e4b57600080fd5b600081830361010081121561322857600080fd5b613230612dc5565b833581526060601f198301121561324657600080fd5b61324e612de8565b915061325c602085016131fc565b825261326a604085016131fc565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c084013560808201526132a960e08501612fcd565b60a0820152949350505050565b6000602082840312156132c857600080fd5b813561ffff81168114611ba357600080fd5b634e487b7160e01b600052601160045260246000fd5b63ffffffff81811683821601908082111561330d5761330d6132da565b5092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff8103613340576133406132da565b60010192915050565b6000600019820361335c5761335c6132da565b5060010190565b808201808211156117e1576117e16132da565b818103818111156117e1576117e16132da565b6000845160005b818110156133aa5760208188018101518583015201613390565b5091909101928352506020820152604001919050565b7f56616c75653a000000000000000000000000000000000000000000000000000081526000600784106133f5576133f5612a57565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea2646970667358221220a2c0a36b8a04d3204c129f1b986f18ae1f521550993a17f1c40abb439323438264736f6c63430008110033",
}

// OneStepProver0ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProver0MetaData.ABI instead.
var OneStepProver0ABI = OneStepProver0MetaData.ABI

// OneStepProver0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProver0MetaData.Bin instead.
var OneStepProver0Bin = OneStepProver0MetaData.Bin

// DeployOneStepProver0 deploys a new Ethereum contract, binding an instance of OneStepProver0 to it.
func DeployOneStepProver0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProver0, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProver0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// OneStepProver0 is an auto generated Go binding around an Ethereum contract.
type OneStepProver0 struct {
	OneStepProver0Caller     // Read-only binding to the contract
	OneStepProver0Transactor // Write-only binding to the contract
	OneStepProver0Filterer   // Log filterer for contract events
}

// OneStepProver0Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProver0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProver0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProver0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProver0Session struct {
	Contract     *OneStepProver0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProver0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProver0CallerSession struct {
	Contract *OneStepProver0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OneStepProver0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProver0TransactorSession struct {
	Contract     *OneStepProver0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OneStepProver0Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProver0Raw struct {
	Contract *OneStepProver0 // Generic contract binding to access the raw methods on
}

// OneStepProver0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProver0CallerRaw struct {
	Contract *OneStepProver0Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProver0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProver0TransactorRaw struct {
	Contract *OneStepProver0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProver0 creates a new instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0(address common.Address, backend bind.ContractBackend) (*OneStepProver0, error) {
	contract, err := bindOneStepProver0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// NewOneStepProver0Caller creates a new read-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Caller(address common.Address, caller bind.ContractCaller) (*OneStepProver0Caller, error) {
	contract, err := bindOneStepProver0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Caller{contract: contract}, nil
}

// NewOneStepProver0Transactor creates a new write-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProver0Transactor, error) {
	contract, err := bindOneStepProver0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Transactor{contract: contract}, nil
}

// NewOneStepProver0Filterer creates a new log filterer instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProver0Filterer, error) {
	contract, err := bindOneStepProver0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Filterer{contract: contract}, nil
}

// bindOneStepProver0 binds a generic wrapper to an already deployed contract.
func bindOneStepProver0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.OneStepProver0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Caller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProver0.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Session) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0CallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverHostIoMetaData contains all meta data concerning the OneStepProverHostIo contract.
var OneStepProverHostIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_customDAValidator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"customDAValidator\",\"outputs\":[{\"internalType\":\"contractICustomDAProofValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620045b4380380620045b4833981016040819052620000349162000046565b6001600160a01b031660805262000078565b6000602082840312156200005957600080fd5b81516001600160a01b03811681146200007157600080fd5b9392505050565b608051614513620000a160003960008181606a01528181610c190152610d4801526145136000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a92cb5011461003b578063c3ea90ba14610065575b600080fd5b61004e61004936600461371d565b6100a4565b60405161005c92919061394d565b60405180910390f35b61008c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161005c565b6100ac6135a6565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526100ff87613e9d565b915061011036879003870187613fdb565b90506000610121602087018761407d565b905061368761801061ffff831610801590610142575061801361ffff831611155b15610150575061035c61033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fe061ffff83160161018457506104eb61033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fdf61ffff8316016101b85750610f1c61033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fde61ffff8316016101ec575061129c61033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fdd61ffff83160161022057506112a861033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fdc61ffff831601610254575061140661033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fd061ffff83160161028857506114b861033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fcf61ffff8316016102bc57506114ff61033d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fce61ffff8316016102f0575061155461033d565b60405162461bcd60e51b815260206004820152601560248201527f494e56414c49445f4d454d4f52595f4f50434f4445000000000000000000000060448201526064015b60405180910390fd5b61034f8a85858a8a8a8763ffffffff16565b5050965096945050505050565b600061036b602085018561407d565b9050610375613691565b60006103828585836115c7565b60c08a01519193509150610395836116a3565b146103e25760405162461bcd60e51b815260206004820152601060248201527f4241445f474c4f42414c5f5354415445000000000000000000000000000000006044820152606401610334565b61ffff831661801014806103fb575061ffff8316618011145b1561041d57610418888884896104138987818d6140a1565b61174c565b6104cf565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fee61ffff8416016104525761041888836118d5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fed61ffff841601610487576104188883611984565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f474c4f42414c53544154455f4f50434f44450000000000006044820152606401610334565b6104d8826116a3565b60c0909801979097525050505050505050565b60006105026104fd87602001516119fb565b611a20565b63ffffffff169050600061051c6104fd88602001516119fb565b63ffffffff16905061052f6020836140e1565b15158061055657506020808701515167ffffffffffffffff169061055490839061410b565b115b8061056a57506105676020826140e1565b15155b15610591578660025b908160028111156105865761058661381e565b815250505050610f14565b600061059e60208361411e565b90506000806105b96040518060200160405280606081525090565b60208a01516105cb90858a8a87611add565b9094509092509050606060008989868181106105e9576105e9614132565b919091013560f81c91508590506105ff81614148565b9550508a60200135600003610740578060ff166000036106f8573660006106288b88818f6140a1565b9150915085828260405161063d929190614162565b6040518091039020146106925760405162461bcd60e51b815260206004820152600c60248201527f4241445f505245494d41474500000000000000000000000000000000000000006044820152606401610334565b600061069f8b602061410b565b9050818111156106ac5750805b6106b8818c84866140a1565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929750610e7495505050505050565b60405162461bcd60e51b815260206004820152601660248201527f554e4b4e4f574e5f505245494d4147455f50524f4f46000000000000000000006044820152606401610334565b8a602001356001036108505760ff81161561079d5760405162461bcd60e51b815260206004820152601660248201527f554e4b4e4f574e5f505245494d4147455f50524f4f46000000000000000000006044820152606401610334565b3660006107ac8b88818f6140a1565b9150915085600283836040516107c3929190614162565b602060405180830381855afa1580156107e0573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906108039190614172565b146106925760405162461bcd60e51b815260206004820152600c60248201527f4241445f505245494d41474500000000000000000000000000000000000000006044820152606401610334565b8a60200135600203610c0b5760ff8116156108ad5760405162461bcd60e51b815260206004820152601660248201527f554e4b4e4f574e5f505245494d4147455f50524f4f46000000000000000000006044820152606401610334565b3660006108bc8b88818f6140a1565b9092509050856108d06020600084866140a1565b6108d99161418b565b146109265760405162461bcd60e51b815260206004820152601460248201527f4b5a475f50524f4f465f57524f4e475f484153480000000000000000000000006044820152606401610334565b600080600080600a6001600160a01b03168686604051610947929190614162565b600060405180830381855afa9150503d8060008114610982576040519150601f19603f3d011682016040523d82523d6000602084013e610987565b606091505b5091509150816109d95760405162461bcd60e51b815260206004820152601160248201527f494e56414c49445f4b5a475f50524f4f460000000000000000000000000000006044820152606401610334565b6000815111610a2a5760405162461bcd60e51b815260206004820152601660248201527f4b5a475f505245434f4d50494c455f4d495353494e47000000000000000000006044820152606401610334565b80806020019051810190610a3e91906141a9565b9094509250507f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff0000000182149050610ab55760405162461bcd60e51b815260206004820152601360248201527f554e4b4e4f574e5f424c535f4d4f44554c5553000000000000000000000000006044820152606401610334565b610ac08260206141cd565b8c1015610c0257600080610ad560208f61411e565b905060015b84811015610b0457600192831b928281169003610af8576001831792505b600191821c911b610ada565b506000610b168564010000000061411e565b9050610b2283826141cd565b90506000610b517f16a2a19edfe81f20d09b681922c813b4b63683508c2280b93829971f439f0d2b8387611b86565b905080610b62604060208a8c6140a1565b610b6b9161418b565b14610bb85760405162461bcd60e51b815260206004820152601160248201527f4b5a475f50524f4f465f57524f4e475f5a0000000000000000000000000000006044820152606401610334565b610bc660606040898b6140a1565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929c50505050505050505b50505050610e74565b8a60200135600303610e2c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610cb35760405162461bcd60e51b815260206004820152602160248201527f435553544f4d5f44415f56414c494441544f525f4e4f545f535550504f52544560448201527f44000000000000000000000000000000000000000000000000000000000000006064820152608401610334565b60ff811615610d045760405162461bcd60e51b815260206004820152601660248201527f554e4b4e4f574e5f505245494d4147455f50524f4f46000000000000000000006044820152606401610334565b366000610d138b88818f6140a1565b6040517fd2f26f5200000000000000000000000000000000000000000000000000000000815291935091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d2f26f5290610d7f90859085906004016141e4565b600060405180830381865afa158015610d9c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610dc49190810190614237565b935060008451118015610dd957506020845111155b610e255760405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f435553544f4d5f44415f524553504f4e53450000000000006044820152606401610334565b5050610e74565b60405162461bcd60e51b815260206004820152601560248201527f554e4b4e4f574e5f505245494d4147455f5459504500000000000000000000006044820152606401610334565b60005b8251811015610eb857610ea48582858481518110610e9757610e97614132565b016020015160f81c611cd4565b945080610eb081614148565b915050610e77565b50610ec4838786611d61565b60208d81015160409081019290925283518251808401845260008082529083018190528351808501909452835263ffffffff1690820152610f0b905b60208f015190611dfb565b50505050505050505b505050505050565b6000610f2e6104fd87602001516119fb565b63ffffffff1690506000610f486104fd88602001516119fb565b63ffffffff1690506000610f67610f6289602001516119fb565b611e0b565b67ffffffffffffffff1690506020860135158015610f86575088358110155b15610fae578760025b90816002811115610fa257610fa261381e565b81525050505050610f14565b6020808801515167ffffffffffffffff1690610fcb90849061410b565b1180610fe05750610fdd6020836140e1565b15155b15610fed57876002610f8f565b6000610ffa60208461411e565b90506000806110156040518060200160405280606081525090565b60208b015161102790858b8b87611add565b909450909250905088888481811061104157611041614132565b909101357fff00000000000000000000000000000000000000000000000000000000000000161590506110b65760405162461bcd60e51b815260206004820152601360248201527f554e4b4e4f574e5f494e424f585f50524f4f46000000000000000000000000006044820152606401610334565b826110c081614148565b93506136879050600060208c01356110dc57611ecd915061111b565b60018c60200135036110f257612203915061111b565b8d60025b908160028111156111095761110961381e565b81525050505050505050505050610f14565b61113b8f888d8d89908092611132939291906140a1565b8663ffffffff16565b90508061114a578d60026110f6565b50508288101561119c5760405162461bcd60e51b815260206004820152601160248201527f4241445f4d4553534147455f50524f4f460000000000000000000000000000006044820152606401610334565b60006111a8848a6142cb565b905060005b60208163ffffffff161080156111d15750816111cf63ffffffff83168b61410b565b105b1561122a576112168463ffffffff83168d8d826111ee8f8c61410b565b6111f8919061410b565b81811061120757611207614132565b919091013560f81c9050611cd4565b935080611222816142de565b9150506111ad565b611235838786611d61565b60208e01516040015261128b61127882604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b8f60200151611dfb90919063ffffffff16565b505050505050505050505050505050565b50506001909252505050565b60006040518060400160405280601381526020017f4d6f64756c65206d65726b6c6520747265653a0000000000000000000000000081525090506000866101600151905060006112fe6104fd89602001516119fb565b63ffffffff16905061131d8188602001516124d190919063ffffffff16565b61132957876002610f8f565b60008061134961133a60208561411e565b60208b01519089896000611add565b509150915060008061135d8c848b8b612507565b92505091506000611379836001611374919061410b565b612763565b905080156113a4576113998761139085600161410b565b8760008c612783565b6101608e01526113c2565b6113bb6113b284600161410b565b8390878b61282d565b6101608e01525b610f0b610f006113d385600161410b565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b60408051808201909152601381527f4d6f64756c65206d65726b6c6520747265653a00000000000000000000000000602082015260008061144988828787612507565b5091509150600061145983612763565b905080156114985781518051611471906001906142cb565b8151811061148157611481614132565b6020026020010151896101600181815250506114ac565b6114a5828460008761282d565b6101608a01525b50505050505050505050565b610140850151600019146114e5578460025b908160028111156114dd576114dd61381e565b905250610f14565b6114f28560a00151612948565b610f148560400151612948565b61014085015160001914611515578460026114ca565b60a08501515160010161152a578460026114ca565b611539856040015183836129bc565b60a0850151610f149061154f83604081876140a1565b6129bc565b60a085015151600101611569578460026114ca565b82602001356000036115985761014085015160010161158a578460026114ca565b6000196101408601526115be565b610140850151600019146115ae578460026114ca565b6115bc856020850135612b44565b505b610f1485612bb7565b6115cf613691565b816115d86136b6565b6115e06136b6565b60005b600260ff8216101561162b576115fa888886612c3b565b848360ff166002811061160f5761160f614132565b602002019190915293508061162381614301565b9150506115e3565b5060005b600260ff8216101561168657611646888886612c57565b838360ff166002811061165b5761165b614132565b67ffffffffffffffff909316602093909302019190915293508061167e81614301565b91505061162f565b506040805180820190915291825260208201529590945092505050565b80518051602091820151828401518051908401516040517f476c6f62616c2073746174653a0000000000000000000000000000000000000095810195909552602d850193909352604d8401919091527fffffffffffffffff00000000000000000000000000000000000000000000000060c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b600061175e6104fd88602001516119fb565b63ffffffff16905060006117786104fd89602001516119fb565b9050600263ffffffff82161061179057876002610573565b602087015161179f90836124d1565b6117ab57876002610573565b60006117b860208461411e565b90506000806117d36040518060200160405280606081525090565b60208b01516117e590858a8a87611add565b90945090925090506180106117fd60208b018b61407d565b61ffff160361184157611833848b600001518763ffffffff166002811061182657611826614132565b6020020151839190611d61565b60208c0151604001526118c7565b61801161185160208b018b61407d565b61ffff160361187f578951829063ffffffff87166002811061187557611875614132565b60200201526118c7565b60405162461bcd60e51b815260206004820152601760248201527f4241445f474c4f42414c5f53544154455f4f50434f44450000000000000000006044820152606401610334565b505050505050505050505050565b60006118e76104fd84602001516119fb565b9050600263ffffffff821610611916578260025b9081600281111561190e5761190e61381e565b905250505050565b61197f61197483602001518363ffffffff166002811061193857611938614132565b6020020151604080518082019091526000808252602082015250604080518082019091526001815267ffffffffffffffff909116602082015290565b602085015190611dfb565b505050565b6000611996610f6284602001516119fb565b905060006119aa6104fd85602001516119fb565b9050600263ffffffff8216106119c4575050600290915250565b8183602001518263ffffffff16600281106119e1576119e1614132565b67ffffffffffffffff909216602092909202015250505050565b60408051808201909152600080825260208201528151611a1a90612cbf565b92915050565b60208101516000908183516006811115611a3c57611a3c61381e565b14611a895760405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152606401610334565b6401000000008110611a1a5760405162461bcd60e51b815260206004820152600760248201527f4241445f493332000000000000000000000000000000000000000000000000006044820152606401610334565b600080611af66040518060200160405280606081525090565b839150611b04868684612c3b565b9093509150611b14868684612dd0565b925090506000611b25828986611d61565b905088604001518114611b7a5760405162461bcd60e51b815260206004820152600e60248201527f57524f4e475f4d454d5f524f4f540000000000000000000000000000000000006044820152606401610334565b50955095509592505050565b60408051602080820181905281830181905260608201526080810185905260a0810184905260c08082018490528251808303909101815260e09091019182905260009182908190600590611bdb908590614320565b600060405180830381855afa9150503d8060008114611c16576040519150601f19603f3d011682016040523d82523d6000602084013e611c1b565b606091505b509150915081611c6d5760405162461bcd60e51b815260206004820152600d60248201527f4d4f444558505f4641494c4544000000000000000000000000000000000000006044820152606401610334565b8051602014611cbe5760405162461bcd60e51b815260206004820152601360248201527f4d4f444558505f57524f4e475f4c454e475448000000000000000000000000006044820152606401610334565b611cc78161433c565b93505050505b9392505050565b600060208310611d265760405162461bcd60e51b815260206004820152601560248201527f4241445f5345545f4c4541465f425954455f49445800000000000000000000006044820152606401610334565b600083611d35600160206142cb565b611d3f91906142cb565b611d4a9060086141cd565b60ff848116821b911b198616179150509392505050565b6040517f4d656d6f7279206c6561663a00000000000000000000000000000000000000006020820152602c81018290526000908190604c01604051602081830303815290604052805190602001209050611df28585836040518060400160405280601381526020017f4d656d6f7279206d65726b6c6520747265653a0000000000000000000000000081525061282d565b95945050505050565b8151611e079082612eab565b5050565b6020810151600090600183516006811115611e2857611e2861381e565b14611e755760405162461bcd60e51b815260206004820152600760248201527f4e4f545f493634000000000000000000000000000000000000000000000000006044820152606401610334565b680100000000000000008110611a1a5760405162461bcd60e51b815260206004820152600760248201527f4241445f493634000000000000000000000000000000000000000000000000006044820152606401610334565b60006028821015611f205760405162461bcd60e51b815260206004820152601260248201527f4241445f534551494e424f585f50524f4f4600000000000000000000000000006044820152606401610334565b6000611f2e84846020612c57565b508091505060008484604051611f45929190614162565b604051908190039020905060008067ffffffffffffffff88161561201057611f7360408a0160208b01614360565b6001600160a01b03166316bf5579611f8c60018b614389565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa158015611fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061200d9190614172565b91505b67ffffffffffffffff8416156120cd5761203060408a0160208b01614360565b6001600160a01b031663d5719dc2612049600187614389565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa1580156120a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ca9190614172565b90505b6040805160208101849052908101849052606081018290526000906080016040516020818303038152906040528051906020012090508960200160208101906121169190614360565b6040517f16bf557900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b1660048201526001600160a01b0391909116906316bf557990602401602060405180830381865afa15801561217f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a39190614172565b81146121f15760405162461bcd60e51b815260206004820152601460248201527f4241445f534551494e424f585f4d4553534147450000000000000000000000006044820152606401610334565b6001955050505050505b949350505050565b600060718210156122565760405162461bcd60e51b815260206004820152601160248201527f4241445f44454c415945445f50524f4f460000000000000000000000000000006044820152606401610334565b600067ffffffffffffffff851615612315576122786040870160208801614360565b6001600160a01b031663d5719dc2612291600188614389565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa1580156122ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123129190614172565b90505b600061232484607181886140a1565b604051612332929190614162565b6040518091039020905060008585600081811061235157612351614132565b9050013560f81c60f81b9050600061236b87876001612f9f565b50905060008282612380607160218b8d6140a1565b876040516020016123959594939291906143b1565b60408051601f198184030181528282528051602091820120838201899052838301819052825180850384018152606090940190925282519201919091209091506123e560408c0160208d01614360565b6040517fd5719dc200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03919091169063d5719dc290602401602060405180830381865afa15801561244e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124729190614172565b81146124c05760405162461bcd60e51b815260206004820152601360248201527f4241445f44454c415945445f4d455353414745000000000000000000000000006044820152606401610334565b5060019a9950505050505050505050565b815160009067ffffffffffffffff166124eb83602061410b565b11158015611ccd57506124ff6020836140e1565b159392505050565b600061251f6040518060200160405280606081525090565b60408051602081019091526060815260408051808201909152601381527f4d6f64756c65206d65726b6c6520747265653a0000000000000000000000000060208201526101608801516125bb6040805160c0810182526000808252825160608101845281815260208181018390529381019190915290918201908152600060208201819052604082018190526060820181905260809091015290565b60006125c889898c612ff4565b9a5091506125d789898c6130fa565b9a5090506125e689898c612dd0565b9a5063ffffffff80831698509096506000906126089088908a90869061315516565b90508381146126595760405162461bcd60e51b815260206004820152601360248201527f57524f4e475f524f4f545f464f525f4c454146000000000000000000000000006044820152606401610334565b505050600061266e866001611374919061410b565b905080156126da5761268186600161410b565b8551516001901b146126d55760405162461bcd60e51b815260206004820152600a60248201527f57524f4e475f4c454146000000000000000000000000000000000000000000006044820152606401610334565b612756565b6126e588888b612dd0565b9950935060006127036126f988600161410b565b869060008761282d565b90508281146127545760405162461bcd60e51b815260206004820152601360248201527f57524f4e475f524f4f545f464f525f5a45524f000000000000000000000000006044820152606401610334565b505b5050509450945094915050565b60008115801590611a1a575061277a6001836142cb565b82161592915050565b600083855b60018111156127f5578382866040516020016127a69392919061441a565b6040516020818303038152906040528051906020012091508385866040516020016127d39392919061441a565b60408051601f198184030181529190528051602090910120945060011c612788565b83888360405160200161280a9392919061441a565b604051602081830303815290604052805190602001209250505095945050505050565b8160005b8551518110156128f957846001166000036128955782828760000151838151811061285e5761285e614132565b60200260200101516040516020016128789392919061441a565b6040516020818303038152906040528051906020012091506128e0565b82866000015182815181106128ac576128ac614132565b6020026020010151836040516020016128c79392919061441a565b6040516020818303038152906040528051906020012091505b60019490941c93806128f181614148565b915050612831565b5083156121fb5760405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f525400000000000000000000000000000000006044820152606401610334565b8051600019146129b65780516020808301516040516129999392017f636f7468726561643a000000000000000000000000000000000000000000000081526009810192909252602982015260490190565b60408051601f198184030181529190528051602091820120908201525b60009052565b600080806129cb858585612c3b565b935091506129da858585612c3b565b9350905060018201612a8b578015612a345760405162461bcd60e51b815260206004820152601460248201527f57524f4e475f434f5448524541445f454d5054590000000000000000000000006044820152606401610334565b602086015115612a865760405162461bcd60e51b815260206004820152601460248201527f57524f4e475f434f5448524541445f454d5054590000000000000000000000006044820152606401610334565b612b37565b85602001518282604051602001612ad49291907f636f7468726561643a000000000000000000000000000000000000000000000081526009810192909252602982015260490190565b6040516020818303038152906040528051906020012014612b375760405162461bcd60e51b815260206004820152601260248201527f57524f4e475f434f5448524541445f504f5000000000000000000000000000006044820152606401610334565b6020860152909352505050565b61014082015160009060001914612b5d57506000611a1a565b600060408460e0015163ffffffff16901b9050602084610100015163ffffffff16901b81179050600183856101200151612b979190614441565b612ba1919061445e565b63ffffffff161761014084015250600192915050565b60408101515160a082015151600019811480612bd4575060001982145b15612be1578260026118fb565b612bee83608001516131a0565b60a0840151526020830151612c0290613244565b6040840151526080830151612c1d9082602082015260609052565b50602091820151808301919091526040805192830190526060825252565b60008181612c4a868684612f9f565b9097909650945050505050565b600081815b6008811015612cb65760088367ffffffffffffffff16901b9250858583818110612c8857612c88614132565b919091013560f81c93909317925081612ca081614148565b9250508080612cae90614148565b915050612c5c565b50935093915050565b604080518082019091526000808252602082015281518051612ce3906001906142cb565b81518110612cf357612cf3614132565b6020026020010151905060006001836000015151612d1191906142cb565b67ffffffffffffffff811115612d2957612d29613ae3565b604051908082528060200260200182016040528015612d6e57816020015b6040805180820190915260008082526020820152815260200190600190039081612d475790505b50905060005b8151811015612dc9578351805182908110612d9157612d91614132565b6020026020010151828281518110612dab57612dab614132565b60200260200101819052508080612dc190614148565b915050612d74565b5090915290565b604080516020810190915260608152816000612ded8686846132da565b92509050600060ff821667ffffffffffffffff811115612e0f57612e0f613ae3565b604051908082528060200260200182016040528015612e38578160200160208202803683370190505b50905060005b8260ff168160ff161015612e8f57612e57888886612c3b565b838360ff1681518110612e6c57612e6c614132565b602002602001018196508281525050508080612e8790614301565b915050612e3e565b5060405180602001604052808281525093505050935093915050565b815151600090612ebc90600161410b565b67ffffffffffffffff811115612ed457612ed4613ae3565b604051908082528060200260200182016040528015612f1957816020015b6040805180820190915260008082526020820152815260200190600190039081612ef25790505b50905060005b835151811015612f75578351805182908110612f3d57612f3d614132565b6020026020010151828281518110612f5757612f57614132565b60200260200101819052508080612f6d90614148565b915050612f1f565b50818184600001515181518110612f8e57612f8e614132565b602090810291909101015290915250565b600081815b6020811015612cb657600883901b9250858583818110612fc657612fc6614132565b919091013560f81c93909317925081612fde81614148565b9250508080612fec90614148565b915050612fa4565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526040805160608101825260008082526020820181905291810182905283919060008060008061306e8b8b89612c3b565b9750955061307d8b8b89613310565b9750945061308c8b8b89612c3b565b9750935061309b8b8b89612c3b565b975092506130aa8b8b89612c3b565b975091506130b98b8b896130fa565b6040805160c081018252988952602089019790975295870194909452506060850191909152608084015263ffffffff1660a083015290969095509350505050565b600081815b6004811015612cb65760088363ffffffff16901b925085858381811061312757613127614132565b919091013560f81c9390931792508161313f81614148565b925050808061314d90614148565b9150506130ff565b60006121fb84846131658561338c565b6040518060400160405280601381526020017f4d6f64756c65206d65726b6c6520747265653a0000000000000000000000000081525061282d565b602081015160005b82515181101561323e576131d8836000015182815181106131cb576131cb614132565b6020026020010151613436565b6040517f537461636b206672616d6520737461636b3a00000000000000000000000000006020820152603281019190915260528101839052607201604051602081830303815290604052805190602001209150808061323690614148565b9150506131a8565b50919050565b60208101518151515160005b818110156132d357835161326d9061326890836134cf565b613507565b6040517f56616c756520737461636b3a00000000000000000000000000000000000000006020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806132cb90614148565b915050613250565b5050919050565b6000818484828181106132ef576132ef614132565b919091013560f81c925081905061330581614148565b915050935093915050565b6040805160608101825260008082526020820181905291810191909152816000808061333d888886612c57565b9450925061334c888886612c57565b9450915061335b888886612c3b565b6040805160608101825267ffffffffffffffff96871681529490951660208501529383015250969095509350505050565b600081600001516133a08360200151613524565b6040808501516060860151608087015160a0880151935161172f969594906020017f4d6f64756c653a0000000000000000000000000000000000000000000000000081526007810196909652602786019490945260478501929092526067840152608783015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660a782015260ab0190565b60006134458260000151613507565b602080840151604080860151606087015191517f537461636b206672616d653a000000000000000000000000000000000000000094810194909452602c840194909452604c8301919091527fffffffff0000000000000000000000000000000000000000000000000000000060e093841b8116606c840152921b909116607082015260740161172f565b604080518082019091526000808252602082015282518051839081106134f7576134f7614132565b6020026020010151905092915050565b60008160000151826020015160405160200161172f92919061447b565b805160208083015160408085015190517f4d656d6f72793a00000000000000000000000000000000000000000000000000938101939093527fffffffffffffffff00000000000000000000000000000000000000000000000060c094851b811660278501529190931b16602f820152603781019190915260009060570161172f565b60408051610180810190915280600081526020016135db60408051606080820183529181019182529081526000602082015290565b815260408051808201825260008082526020808301919091528301520161361960408051606080820183529181019182529081526000602082015290565b815260200161363e604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b61368f6144c7565b565b60405180604001604052806136a46136b6565b81526020016136b16136b6565b905290565b60405180604001604052806002906020820280368337509192915050565b60008083601f8401126136e657600080fd5b50813567ffffffffffffffff8111156136fe57600080fd5b60208301915083602082850101111561371657600080fd5b9250929050565b6000806000806000808688036101e081121561373857600080fd5b606081121561374657600080fd5b879650606088013567ffffffffffffffff8082111561376457600080fd5b818a0191506101c080838d03121561377b57600080fd5b8298506101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850112156137af57600080fd5b60808b01975060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe80850112156137e557600080fd5b6101808b0196508a01359250808311156137fe57600080fd5b505061380c89828a016136d4565b979a9699509497509295939492505050565b634e487b7160e01b600052602160045260246000fd5b600381106138445761384461381e565b9052565b80516007811061385a5761385a61381e565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b808410156138b7576138a3828651613848565b938201936001939093019290850190613890565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015613936578451613902858251613848565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a0909301926001016138ed565b509687015197909601969096525093949350505050565b60006101208083526139628184018651613834565b60208501516101c061014081818701526139806102e0870184613867565b925060408801516101606139a08189018380518252602090810151910152565b60608a015191507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee080898703016101a08a01526139dd8684613867565b955060808b015192508089870301858a0152506139fa85836138cb565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c0860152509150611ccd9050602083018480518252602081015167ffffffffffffffff80825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613b1c57613b1c613ae3565b60405290565b6040516020810167ffffffffffffffff81118282101715613b1c57613b1c613ae3565b6040516080810167ffffffffffffffff81118282101715613b1c57613b1c613ae3565b604051610180810167ffffffffffffffff81118282101715613b1c57613b1c613ae3565b60405160c0810167ffffffffffffffff81118282101715613b1c57613b1c613ae3565b6040516060810167ffffffffffffffff81118282101715613b1c57613b1c613ae3565b604051601f8201601f1916810167ffffffffffffffff81118282101715613bfb57613bfb613ae3565b604052919050565b803560038110613c1257600080fd5b919050565b600067ffffffffffffffff821115613c3157613c31613ae3565b5060051b60200190565b600060408284031215613c4d57600080fd5b613c55613af9565b9050813560078110613c6657600080fd5b808252506020820135602082015292915050565b60006040808385031215613c8d57600080fd5b613c95613af9565b9150823567ffffffffffffffff80821115613caf57600080fd5b81850191506020808388031215613cc557600080fd5b613ccd613b22565b833583811115613cdc57600080fd5b80850194505087601f850112613cf157600080fd5b83359250613d06613d0184613c17565b613bd2565b83815260069390931b84018201928281019089851115613d2557600080fd5b948301945b84861015613d4b57613d3c8a87613c3b565b82529486019490830190613d2a565b8252508552948501359484019490945250909392505050565b600060408284031215613d7657600080fd5b613d7e613af9565b9050813581526020820135602082015292915050565b803563ffffffff81168114613c1257600080fd5b60006040808385031215613dbb57600080fd5b613dc3613af9565b9150823567ffffffffffffffff811115613ddc57600080fd5b8301601f81018513613ded57600080fd5b80356020613dfd613d0183613c17565b82815260a09283028401820192828201919089851115613e1c57600080fd5b948301945b84861015613e855780868b031215613e395760008081fd5b613e41613b45565b613e4b8b88613c3b565b815287870135858201526060613e62818901613d94565b89830152613e7260808901613d94565b9082015283529485019491830191613e21565b50808752505080860135818601525050505092915050565b60006101c08236031215613eb057600080fd5b613eb8613b68565b613ec183613c03565b8152602083013567ffffffffffffffff80821115613ede57600080fd5b613eea36838701613c7a565b6020840152613efc3660408701613d64565b60408401526080850135915080821115613f1557600080fd5b613f2136838701613c7a565b606084015260a0850135915080821115613f3a57600080fd5b50613f4736828601613da8565b608083015250613f5a3660c08501613d64565b60a08201526101008084013560c0830152610120613f79818601613d94565b60e0840152610140613f8c818701613d94565b838501526101609250613fa0838701613d94565b91840191909152610180850135908301526101a090930135928101929092525090565b803567ffffffffffffffff81168114613c1257600080fd5b6000818303610100811215613fef57600080fd5b613ff7613b8c565b833581526060601f198301121561400d57600080fd5b614015613baf565b915061402360208501613fc3565b825261403160408501613fc3565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c0840135608082015261407060e08501613d94565b60a0820152949350505050565b60006020828403121561408f57600080fd5b813561ffff81168114611ccd57600080fd5b600080858511156140b157600080fd5b838611156140be57600080fd5b5050820193919092039150565b634e487b7160e01b600052601260045260246000fd5b6000826140f0576140f06140cb565b500690565b634e487b7160e01b600052601160045260246000fd5b80820180821115611a1a57611a1a6140f5565b60008261412d5761412d6140cb565b500490565b634e487b7160e01b600052603260045260246000fd5b6000600019820361415b5761415b6140f5565b5060010190565b8183823760009101908152919050565b60006020828403121561418457600080fd5b5051919050565b80356020831015611a1a57600019602084900360031b1b1692915050565b600080604083850312156141bc57600080fd5b505080516020909101519092909150565b8082028115828204841417611a1a57611a1a6140f5565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60005b8381101561422e578181015183820152602001614216565b50506000910152565b60006020828403121561424957600080fd5b815167ffffffffffffffff8082111561426157600080fd5b818401915084601f83011261427557600080fd5b81518181111561428757614287613ae3565b61429a6020601f19601f84011601613bd2565b91508082528560208285010111156142b157600080fd5b6142c2816020840160208601614213565b50949350505050565b81810381811115611a1a57611a1a6140f5565b600063ffffffff8083168181036142f7576142f76140f5565b6001019392505050565b600060ff821660ff8103614317576143176140f5565b60010192915050565b60008251614332818460208701614213565b9190910192915050565b8051602080830151919081101561323e5760001960209190910360031b1b16919050565b60006020828403121561437257600080fd5b81356001600160a01b0381168114611ccd57600080fd5b67ffffffffffffffff8281168282160390808211156143aa576143aa6140f5565b5092915050565b7fff00000000000000000000000000000000000000000000000000000000000000861681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b1660018201528284601583013760159201918201526035019392505050565b6000845161442c818460208901614213565b91909101928352506020820152604001919050565b63ffffffff8181168382160190808211156143aa576143aa6140f5565b63ffffffff8281168282160390808211156143aa576143aa6140f5565b7f56616c75653a000000000000000000000000000000000000000000000000000081526000600784106144b0576144b061381e565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea2646970667358221220d614804c551db2a39177c4bcd8b24ce3ac18a4b3ec38e85d474f99f648e1934364736f6c63430008110033",
}

// OneStepProverHostIoABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverHostIoMetaData.ABI instead.
var OneStepProverHostIoABI = OneStepProverHostIoMetaData.ABI

// OneStepProverHostIoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverHostIoMetaData.Bin instead.
var OneStepProverHostIoBin = OneStepProverHostIoMetaData.Bin

// DeployOneStepProverHostIo deploys a new Ethereum contract, binding an instance of OneStepProverHostIo to it.
func DeployOneStepProverHostIo(auth *bind.TransactOpts, backend bind.ContractBackend, _customDAValidator common.Address) (common.Address, *types.Transaction, *OneStepProverHostIo, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverHostIoBin), backend, _customDAValidator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// OneStepProverHostIo is an auto generated Go binding around an Ethereum contract.
type OneStepProverHostIo struct {
	OneStepProverHostIoCaller     // Read-only binding to the contract
	OneStepProverHostIoTransactor // Write-only binding to the contract
	OneStepProverHostIoFilterer   // Log filterer for contract events
}

// OneStepProverHostIoCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverHostIoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverHostIoSession struct {
	Contract     *OneStepProverHostIo // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverHostIoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverHostIoCallerSession struct {
	Contract *OneStepProverHostIoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverHostIoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverHostIoTransactorSession struct {
	Contract     *OneStepProverHostIoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverHostIoRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverHostIoRaw struct {
	Contract *OneStepProverHostIo // Generic contract binding to access the raw methods on
}

// OneStepProverHostIoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCallerRaw struct {
	Contract *OneStepProverHostIoCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverHostIoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactorRaw struct {
	Contract *OneStepProverHostIoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverHostIo creates a new instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIo(address common.Address, backend bind.ContractBackend) (*OneStepProverHostIo, error) {
	contract, err := bindOneStepProverHostIo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// NewOneStepProverHostIoCaller creates a new read-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverHostIoCaller, error) {
	contract, err := bindOneStepProverHostIo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoCaller{contract: contract}, nil
}

// NewOneStepProverHostIoTransactor creates a new write-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverHostIoTransactor, error) {
	contract, err := bindOneStepProverHostIo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoTransactor{contract: contract}, nil
}

// NewOneStepProverHostIoFilterer creates a new log filterer instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverHostIoFilterer, error) {
	contract, err := bindOneStepProverHostIo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoFilterer{contract: contract}, nil
}

// bindOneStepProverHostIo binds a generic wrapper to an already deployed contract.
func bindOneStepProverHostIo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transact(opts, method, params...)
}

// CustomDAValidator is a free data retrieval call binding the contract method 0xc3ea90ba.
//
// Solidity: function customDAValidator() view returns(address)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) CustomDAValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "customDAValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomDAValidator is a free data retrieval call binding the contract method 0xc3ea90ba.
//
// Solidity: function customDAValidator() view returns(address)
func (_OneStepProverHostIo *OneStepProverHostIoSession) CustomDAValidator() (common.Address, error) {
	return _OneStepProverHostIo.Contract.CustomDAValidator(&_OneStepProverHostIo.CallOpts)
}

// CustomDAValidator is a free data retrieval call binding the contract method 0xc3ea90ba.
//
// Solidity: function customDAValidator() view returns(address)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) CustomDAValidator() (common.Address, error) {
	return _OneStepProverHostIo.Contract.CustomDAValidator(&_OneStepProverHostIo.CallOpts)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "executeOneStep", execCtx, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// OneStepProverMathMetaData contains all meta data concerning the OneStepProverMath contract.
var OneStepProverMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506127fe806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a92cb50114610030575b600080fd5b61004361003e366004611c9f565b61005a565b604051610051929190611ecf565b60405180910390f35b610062611b6b565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526100b58761241a565b91506100c636879003870187612558565b905060006100d760208701876125fa565b9050611c4c61ffff8216604514806100f3575061ffff82166050145b15610101575061033d61031f565b604661ffff831610801590610129575061011d60096046612634565b61ffff168261ffff1611155b1561013757506104ed61031f565b606761ffff83161080159061015f575061015360026067612634565b61ffff168261ffff1611155b1561016d57506105d061031f565b606a61ffff8316108015906101875750607861ffff831611155b15610195575061065d61031f565b605161ffff8316108015906101bd57506101b160096051612634565b61ffff168261ffff1611155b156101cb575061088561031f565b607961ffff8316108015906101f357506101e760026079612634565b61ffff168261ffff1611155b1561020157506108ea61031f565b607c61ffff83161080159061021b5750608a61ffff831611155b15610229575061096461031f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5961ffff83160161025d5750610b6261031f565b61ffff821660ac1480610274575061ffff821660ad145b156102825750610bad61031f565b60c061ffff83161080159061029c575060c461ffff831611155b156102aa5750610c2c61031f565b60bc61ffff8316108015906102c4575060bf61ffff831611155b156102d25750610e4561031f565b60405162461bcd60e51b815260206004820152600e60248201527f494e56414c49445f4f50434f444500000000000000000000000000000000000060448201526064015b60405180910390fd5b61033084848989898663ffffffff16565b5050965096945050505050565b600061034c8660200151610fdb565b9050604561035d60208601866125fa565b61ffff16036103cd5760008151600681111561037b5761037b611da0565b146103c85760405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152606401610316565b61048f565b60506103dc60208601866125fa565b61ffff1603610447576001815160068111156103fa576103fa611da0565b146103c85760405162461bcd60e51b815260206004820152600760248201527f4e4f545f493634000000000000000000000000000000000000000000000000006044820152606401610316565b60405162461bcd60e51b815260206004820152600760248201527f4241445f45515a000000000000000000000000000000000000000000000000006044820152606401610316565b600081602001516000036104a5575060016104a9565b5060005b604080518082018252600080825260209182018190528251808401909352825263ffffffff8316908201526104e4905b602089015190611000565b50505050505050565b60006105046104ff8760200151610fdb565b611010565b905060006105186104ff8860200151610fdb565b90506000604661052b60208801886125fa565b6105359190612656565b905060008061ffff831660021480610551575061ffff83166004145b80610560575061ffff83166006145b8061056f575061ffff83166008145b1561058f5761057d846110cd565b9150610588856110cd565b905061059d565b505063ffffffff8083169084165b60006105aa8383866110f9565b90506105c36105b882611393565b60208d015190611000565b5050505050505050505050565b60006105e26104ff8760200151610fdb565b9050600060676105f560208701876125fa565b6105ff9190612656565b905060006106158363ffffffff16836020611407565b604080518082018252600080825260209182018190528251808401909352825263ffffffff831690820152909150610653905b60208a015190611000565b5050505050505050565b600061066f6104ff8760200151610fdb565b905060006106836104ff8860200151610fdb565b9050600080606a61069760208901896125fa565b6106a19190612656565b90508061ffff166003036107395763ffffffff841615806106f357508260030b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff800000001480156106f357508360030b600019145b1561071c578860025b9081600281111561070f5761070f611da0565b815250505050505061087e565b8360030b8360030b8161073157610731612671565b05915061083e565b8061ffff16600503610778578363ffffffff1660000361075b578860026106fc565b8360030b8360030b8161077057610770612671565b07915061083e565b8061ffff16600a036107975763ffffffff8316601f85161b915061083e565b8061ffff16600c036107b65763ffffffff8316601f85161c915061083e565b8061ffff16600b036107d357600383900b601f85161d915061083e565b8061ffff16600d036107f0576107e983856115e7565b915061083e565b8061ffff16600e03610806576107e98385611629565b6000806108208563ffffffff168763ffffffff168561166b565b91509150801561083a575050600289525061087e92505050565b5091505b604080518082018252600080825260209182018190528251808401909352825263ffffffff841690820152610879905b60208b015190611000565b505050505b5050505050565b600061089c6108978760200151610fdb565b611800565b905060006108b06108978860200151610fdb565b9050600060516108c360208801886125fa565b6108cd9190612656565b905060006108dc8385846110f9565b905061087961086e82611393565b60006108fc6108978760200151610fdb565b90506000607961090f60208701876125fa565b6109199190612656565b9050600061092983836040611407565b604080518082018252600080825260209182015281518083019092526001825263ffffffff9290921691810182905290915061065390610648565b60006109766108978760200151610fdb565b9050600061098a6108978860200151610fdb565b9050600080607c61099e60208901896125fa565b6109a89190612656565b90508061ffff16600303610a285767ffffffffffffffff841615806109fe57508260070b7fffffffffffffffffffffffffffffffffffffffffffffffff80000000000000001480156109fe57508360070b600019145b15610a0b578860026106fc565b8360070b8360070b81610a2057610a20612671565b059150610b2a565b8061ffff16600503610a6b578367ffffffffffffffff16600003610a4e578860026106fc565b8360070b8360070b81610a6357610a63612671565b079150610b2a565b8061ffff16600a03610a8e5767ffffffffffffffff8316603f85161b9150610b2a565b8061ffff16600c03610ab15767ffffffffffffffff8316603f85161c9150610b2a565b8061ffff16600b03610ace57600783900b603f85161d9150610b2a565b8061ffff16600d03610aeb57610ae483856118c2565b9150610b2a565b8061ffff16600e03610b0157610ae48385611914565b6000610b0e84868461166b565b90935090508015610b28575050600288525061087e915050565b505b604080518082018252600080825260209182015281518083019092526001825267ffffffffffffffff8416908201526108799061086e565b6000610b746108978760200151610fdb565b604080518082018252600080825260209182018190528251808401909352825263ffffffff83169082015290915081906104e4906104d9565b6000610bbf6104ff8760200151610fdb565b9050600060ac610bd260208701876125fa565b61ffff1603610beb57610be4826110cd565b9050610bf4565b5063ffffffff81165b604080518082018252600080825260209182015281518083019092526001825267ffffffffffffffff8316908201526104e4906104d9565b60008060c0610c3e60208701876125fa565b61ffff1603610c535750600090506008610d2b565b60c1610c6260208701876125fa565b61ffff1603610c775750600090506010610d2b565b60c2610c8660208701876125fa565b61ffff1603610c9b5750600190506008610d2b565b60c3610caa60208701876125fa565b61ffff1603610cbf5750600190506010610d2b565b60c4610cce60208701876125fa565b61ffff1603610ce35750600190506020610d2b565b60405162461bcd60e51b815260206004820152601860248201527f494e56414c49445f455854454e445f53414d455f5459504500000000000000006044820152606401610316565b600080836006811115610d4057610d40611da0565b03610d50575063ffffffff610d5b565b5067ffffffffffffffff5b6000610d6a8960200151610fdb565b9050836006811115610d7e57610d7e611da0565b81516006811115610d9157610d91611da0565b14610dde5760405162461bcd60e51b815260206004820152601960248201527f4241445f455854454e445f53414d455f545950455f54595045000000000000006044820152606401610316565b6000610df1600160ff861681901b612687565b602083018051821690529050610e0860018561269a565b60ff166001901b826020015116600014610e2a57602082018051821985161790525b60208a0151610e399083611000565b50505050505050505050565b60008060bc610e5760208701876125fa565b61ffff1603610e6c5750600090506002610f20565b60bd610e7b60208701876125fa565b61ffff1603610e905750600190506003610f20565b60be610e9f60208701876125fa565b61ffff1603610eb45750600290506000610f20565b60bf610ec360208701876125fa565b61ffff1603610ed85750600390506001610f20565b60405162461bcd60e51b815260206004820152601360248201527f494e56414c49445f5245494e54455250524554000000000000000000000000006044820152606401610316565b6000610f2f8860200151610fdb565b9050816006811115610f4357610f43611da0565b81516006811115610f5657610f56611da0565b14610fa35760405162461bcd60e51b815260206004820152601860248201527f494e56414c49445f5245494e544552505245545f5459504500000000000000006044820152606401610316565b80836006811115610fb657610fb6611da0565b90816006811115610fc957610fc9611da0565b90525060208801516106539082611000565b60408051808201909152600080825260208201528151610ffa90611966565b92915050565b815161100c9082611a77565b5050565b6020810151600090818351600681111561102c5761102c611da0565b146110795760405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152606401610316565b6401000000008110610ffa5760405162461bcd60e51b815260206004820152600760248201527f4241445f493332000000000000000000000000000000000000000000000000006044820152606401610316565b600063800000008216156110ef575063ffffffff1667ffffffff000000001790565b5063ffffffff1690565b600061ffff8216611122578267ffffffffffffffff168467ffffffffffffffff1614905061138c565b60001961ffff83160161114e578267ffffffffffffffff168467ffffffffffffffff161415905061138c565b60011961ffff83160161116b578260070b8460070b12905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd61ffff8316016111b4578267ffffffffffffffff168467ffffffffffffffff1610905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc61ffff8316016111ef578260070b8460070b13905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb61ffff831601611238578267ffffffffffffffff168467ffffffffffffffff1611905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa61ffff831601611274578260070b8460070b1315905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff961ffff8316016112be578267ffffffffffffffff168467ffffffffffffffff161115905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff861ffff8316016112fa578260070b8460070b1215905061138c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff761ffff831601611344578267ffffffffffffffff168467ffffffffffffffff161015905061138c565b60405162461bcd60e51b815260206004820152600a60248201527f424144204952454c4f50000000000000000000000000000000000000000000006044820152606401610316565b9392505050565b604080518082019091526000808252602082015281156113d8576040805180820182526000808252602091820181905282518084019093528252600190820152610ffa565b60408051808201825260008082526020918201819052825180840190935280835290820152610ffa565b919050565b60008161ffff166020148061142057508161ffff166040145b61146c5760405162461bcd60e51b815260206004820152601860248201527f57524f4e4720555345204f462067656e65726963556e4f7000000000000000006044820152606401610316565b61ffff83166114de5761ffff82165b60008163ffffffff161180156114b157506114976001826126b3565b63ffffffff166001901b8567ffffffffffffffff16166000145b156114c8576114c16001826126b3565b905061147b565b6114d68161ffff85166126b3565b91505061138c565b60001961ffff8416016115385760005b8261ffff168163ffffffff1610801561151a5750600163ffffffff82161b851667ffffffffffffffff16155b156115315761152a6001826126d0565b90506114ee565b905061138c565b60011961ffff84160161159f576000805b8361ffff168263ffffffff16101561159657600163ffffffff83161b861667ffffffffffffffff1615611584576115816001826126d0565b90505b8161158e816126ed565b925050611549565b915061138c9050565b60405162461bcd60e51b815260206004820152600960248201527f4241442049556e4f7000000000000000000000000000000000000000000000006044820152606401610316565b60006115f4602083612710565b91506116018260206126b3565b63ffffffff168363ffffffff16901c8263ffffffff168463ffffffff16901b17905092915050565b6000611636602083612710565b91506116438260206126b3565b63ffffffff168363ffffffff16901b8263ffffffff168463ffffffff16901c17905092915050565b6000808261ffff1660000361168657505082820160006117f8565b8261ffff1660010361169e57505081830360006117f8565b8261ffff166002036116b657505082820260006117f8565b8261ffff1660040361170f578367ffffffffffffffff166000036116e057506000905060016117f8565b8367ffffffffffffffff168567ffffffffffffffff168161170357611703612671565b046000915091506117f8565b8261ffff16600603611768578367ffffffffffffffff1660000361173957506000905060016117f8565b8367ffffffffffffffff168567ffffffffffffffff168161175c5761175c612671565b066000915091506117f8565b8261ffff1660070361178057505082821660006117f8565b8261ffff1660080361179857505082821760006117f8565b8261ffff166009036117b057505082821860006117f8565b60405162461bcd60e51b815260206004820152601660248201527f494e56414c49445f47454e455249435f42494e5f4f50000000000000000000006044820152606401610316565b935093915050565b602081015160009060018351600681111561181d5761181d611da0565b1461186a5760405162461bcd60e51b815260206004820152600760248201527f4e4f545f493634000000000000000000000000000000000000000000000000006044820152606401610316565b680100000000000000008110610ffa5760405162461bcd60e51b815260206004820152600760248201527f4241445f493634000000000000000000000000000000000000000000000000006044820152606401610316565b60006118cf604083612733565b91506118dc82604061274e565b67ffffffffffffffff168367ffffffffffffffff16901c8267ffffffffffffffff168467ffffffffffffffff16901b17905092915050565b6000611921604083612733565b915061192e82604061274e565b67ffffffffffffffff168367ffffffffffffffff16901b8267ffffffffffffffff168467ffffffffffffffff16901c17905092915050565b60408051808201909152600080825260208201528151805161198a90600190612687565b8151811061199a5761199a61276f565b60200260200101519050600060018360000151516119b89190612687565b67ffffffffffffffff8111156119d0576119d0612065565b604051908082528060200260200182016040528015611a1557816020015b60408051808201909152600080825260208201528152602001906001900390816119ee5790505b50905060005b8151811015611a70578351805182908110611a3857611a3861276f565b6020026020010151828281518110611a5257611a5261276f565b60200260200101819052508080611a6890612785565b915050611a1b565b5090915290565b815151600090611a8890600161279f565b67ffffffffffffffff811115611aa057611aa0612065565b604051908082528060200260200182016040528015611ae557816020015b6040805180820190915260008082526020820152815260200190600190039081611abe5790505b50905060005b835151811015611b41578351805182908110611b0957611b0961276f565b6020026020010151828281518110611b2357611b2361276f565b60200260200101819052508080611b3990612785565b915050611aeb565b50818184600001515181518110611b5a57611b5a61276f565b602090810291909101015290915250565b6040805161018081019091528060008152602001611ba060408051606080820183529181019182529081526000602082015290565b8152604080518082018252600080825260208083019190915283015201611bde60408051606080820183529181019182529081526000602082015290565b8152602001611c03604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b611c546127b2565b565b60008083601f840112611c6857600080fd5b50813567ffffffffffffffff811115611c8057600080fd5b602083019150836020828501011115611c9857600080fd5b9250929050565b6000806000806000808688036101e0811215611cba57600080fd5b6060811215611cc857600080fd5b879650606088013567ffffffffffffffff80821115611ce657600080fd5b818a0191506101c080838d031215611cfd57600080fd5b8298506101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8085011215611d3157600080fd5b60808b01975060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8085011215611d6757600080fd5b6101808b0196508a0135925080831115611d8057600080fd5b5050611d8e89828a01611c56565b979a9699509497509295939492505050565b634e487b7160e01b600052602160045260246000fd5b60038110611dc657611dc6611da0565b9052565b805160078110611ddc57611ddc611da0565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611e3957611e25828651611dca565b938201936001939093019290850190611e12565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015611eb8578451611e84858251611dca565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101611e6f565b509687015197909601969096525093949350505050565b6000610120808352611ee48184018651611db6565b60208501516101c06101408181870152611f026102e0870184611de9565b92506040880151610160611f228189018380518252602090810151910152565b60608a015191507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee080898703016101a08a0152611f5f8684611de9565b955060808b015192508089870301858a015250611f7c8583611e4d565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c086015250915061138c9050602083018480518252602081015167ffffffffffffffff80825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561209e5761209e612065565b60405290565b6040516020810167ffffffffffffffff8111828210171561209e5761209e612065565b6040516080810167ffffffffffffffff8111828210171561209e5761209e612065565b604051610180810167ffffffffffffffff8111828210171561209e5761209e612065565b60405160c0810167ffffffffffffffff8111828210171561209e5761209e612065565b6040516060810167ffffffffffffffff8111828210171561209e5761209e612065565b604051601f8201601f1916810167ffffffffffffffff8111828210171561217d5761217d612065565b604052919050565b80356003811061140257600080fd5b600067ffffffffffffffff8211156121ae576121ae612065565b5060051b60200190565b6000604082840312156121ca57600080fd5b6121d261207b565b90508135600781106121e357600080fd5b808252506020820135602082015292915050565b6000604080838503121561220a57600080fd5b61221261207b565b9150823567ffffffffffffffff8082111561222c57600080fd5b8185019150602080838803121561224257600080fd5b61224a6120a4565b83358381111561225957600080fd5b80850194505087601f85011261226e57600080fd5b8335925061228361227e84612194565b612154565b83815260069390931b840182019282810190898511156122a257600080fd5b948301945b848610156122c8576122b98a876121b8565b825294860194908301906122a7565b8252508552948501359484019490945250909392505050565b6000604082840312156122f357600080fd5b6122fb61207b565b9050813581526020820135602082015292915050565b803563ffffffff8116811461140257600080fd5b6000604080838503121561233857600080fd5b61234061207b565b9150823567ffffffffffffffff81111561235957600080fd5b8301601f8101851361236a57600080fd5b8035602061237a61227e83612194565b82815260a0928302840182019282820191908985111561239957600080fd5b948301945b848610156124025780868b0312156123b65760008081fd5b6123be6120c7565b6123c88b886121b8565b8152878701358582015260606123df818901612311565b898301526123ef60808901612311565b908201528352948501949183019161239e565b50808752505080860135818601525050505092915050565b60006101c0823603121561242d57600080fd5b6124356120ea565b61243e83612185565b8152602083013567ffffffffffffffff8082111561245b57600080fd5b612467368387016121f7565b602084015261247936604087016122e1565b6040840152608085013591508082111561249257600080fd5b61249e368387016121f7565b606084015260a08501359150808211156124b757600080fd5b506124c436828601612325565b6080830152506124d73660c085016122e1565b60a08201526101008084013560c08301526101206124f6818601612311565b60e0840152610140612509818701612311565b83850152610160925061251d838701612311565b91840191909152610180850135908301526101a090930135928101929092525090565b803567ffffffffffffffff8116811461140257600080fd5b600081830361010081121561256c57600080fd5b61257461210e565b833581526060601f198301121561258a57600080fd5b612592612131565b91506125a060208501612540565b82526125ae60408501612540565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c084013560808201526125ed60e08501612311565b60a0820152949350505050565b60006020828403121561260c57600080fd5b813561ffff8116811461138c57600080fd5b634e487b7160e01b600052601160045260246000fd5b61ffff81811683821601908082111561264f5761264f61261e565b5092915050565b61ffff82811682821603908082111561264f5761264f61261e565b634e487b7160e01b600052601260045260246000fd5b81810381811115610ffa57610ffa61261e565b60ff8281168282160390811115610ffa57610ffa61261e565b63ffffffff82811682821603908082111561264f5761264f61261e565b63ffffffff81811683821601908082111561264f5761264f61261e565b600063ffffffff8083168181036127065761270661261e565b6001019392505050565b600063ffffffff8084168061272757612727612671565b92169190910692915050565b600067ffffffffffffffff8084168061272757612727612671565b67ffffffffffffffff82811682821603908082111561264f5761264f61261e565b634e487b7160e01b600052603260045260246000fd5b600060001982036127985761279861261e565b5060010190565b80820180821115610ffa57610ffa61261e565b634e487b7160e01b600052605160045260246000fdfea26469706673582212209f51f2d771696cc1565d970e89bf70009ec9898e61ea5b7dd1b8166a89a6479d64736f6c63430008110033",
}

// OneStepProverMathABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMathMetaData.ABI instead.
var OneStepProverMathABI = OneStepProverMathMetaData.ABI

// OneStepProverMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMathMetaData.Bin instead.
var OneStepProverMathBin = OneStepProverMathMetaData.Bin

// DeployOneStepProverMath deploys a new Ethereum contract, binding an instance of OneStepProverMath to it.
func DeployOneStepProverMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMath, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// OneStepProverMath is an auto generated Go binding around an Ethereum contract.
type OneStepProverMath struct {
	OneStepProverMathCaller     // Read-only binding to the contract
	OneStepProverMathTransactor // Write-only binding to the contract
	OneStepProverMathFilterer   // Log filterer for contract events
}

// OneStepProverMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMathSession struct {
	Contract     *OneStepProverMath // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProverMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMathCallerSession struct {
	Contract *OneStepProverMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProverMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMathTransactorSession struct {
	Contract     *OneStepProverMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProverMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMathRaw struct {
	Contract *OneStepProverMath // Generic contract binding to access the raw methods on
}

// OneStepProverMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMathCallerRaw struct {
	Contract *OneStepProverMathCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactorRaw struct {
	Contract *OneStepProverMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMath creates a new instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMath(address common.Address, backend bind.ContractBackend) (*OneStepProverMath, error) {
	contract, err := bindOneStepProverMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// NewOneStepProverMathCaller creates a new read-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMathCaller, error) {
	contract, err := bindOneStepProverMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathCaller{contract: contract}, nil
}

// NewOneStepProverMathTransactor creates a new write-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMathTransactor, error) {
	contract, err := bindOneStepProverMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathTransactor{contract: contract}, nil
}

// NewOneStepProverMathFilterer creates a new log filterer instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMathFilterer, error) {
	contract, err := bindOneStepProverMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathFilterer{contract: contract}, nil
}

// bindOneStepProverMath binds a generic wrapper to an already deployed contract.
func bindOneStepProverMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.OneStepProverMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMath.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverMemoryMetaData contains all meta data concerning the OneStepProverMemory contract.
var OneStepProverMemoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506120e0806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a92cb50114610030575b600080fd5b61004361003e36600461157e565b61005a565b6040516100519291906117ae565b60405180910390f35b61006261144a565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526100b587611cfe565b91506100c636879003870187611e3c565b905060006100d76020870187611ede565b905061152b602861ffff8316108015906100f65750603561ffff831611155b1561010457506101ff6101e1565b603661ffff83161080159061011e5750603e61ffff831611155b1561012c57506106656101e1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc161ffff8316016101605750610a1c6101e1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc061ffff8316016101945750610a766101e1565b60405162461bcd60e51b815260206004820152601560248201527f494e56414c49445f4d454d4f52595f4f50434f4445000000000000000000000060448201526064015b60405180910390fd5b6101f284848989898663ffffffff16565b5050965096945050505050565b6000808060286102126020880188611ede565b61ffff160361022a5750600091506004905081610472565b60296102396020880188611ede565b61ffff1603610252575060019150600890506000610472565b602a6102616020880188611ede565b61ffff160361027a575060029150600490506000610472565b602b6102896020880188611ede565b61ffff16036102a2575060039150600890506000610472565b602c6102b16020880188611ede565b61ffff16036102c95750600091506001905080610472565b602d6102d86020880188611ede565b61ffff16036102f05750600091506001905081610472565b602e6102ff6020880188611ede565b61ffff1603610318575060009150600290506001610472565b602f6103276020880188611ede565b61ffff160361033f5750600091506002905081610472565b603061034e6020880188611ede565b61ffff160361036557506001915081905080610472565b60316103746020880188611ede565b61ffff160361038c5750600191508190506000610472565b603261039b6020880188611ede565b61ffff16036103b35750600191506002905081610472565b60336103c26020880188611ede565b61ffff16036103db575060019150600290506000610472565b60346103ea6020880188611ede565b61ffff16036104025750600191506004905081610472565b60356104116020880188611ede565b61ffff160361042a575060019150600490506000610472565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f4d454d4f52595f4c4f41445f4f50434f444500000000000060448201526064016101d8565b60006104896104848a60200151610b79565b610b9e565b61049d9063ffffffff166020890135611f18565b602089015190915060009081906104b89084878b8b86610c5b565b509150915081156104d3575050600289525061065e92505050565b80841561061a578560011480156104fb575060008760068111156104f9576104f961167f565b145b15610511578060000b63ffffffff16905061061a565b856001148015610532575060018760068111156105305761053061167f565b145b1561053f5760000b61061a565b8560021480156105605750600087600681111561055e5761055e61167f565b145b15610576578060010b63ffffffff16905061061a565b856002148015610597575060018760068111156105955761059561167f565b145b156105a45760010b61061a565b8560041480156105c5575060018760068111156105c3576105c361167f565b145b156105d25760030b61061a565b60405162461bcd60e51b815260206004820152601560248201527f4241445f524541445f42595445535f5349474e4544000000000000000000000060448201526064016101d8565b610656604051806040016040528089600681111561063a5761063a61167f565b815267ffffffffffffffff84166020918201528e015190610d34565b505050505050505b5050505050565b6000808060366106786020880188611ede565b61ffff160361068d57506004915060006107f4565b603761069c6020880188611ede565b61ffff16036106b157506008915060016107f4565b60386106c06020880188611ede565b61ffff16036106d557506004915060026107f4565b60396106e46020880188611ede565b61ffff16036106f957506008915060036107f4565b603a6107086020880188611ede565b61ffff160361071d57506001915060006107f4565b603b61072c6020880188611ede565b61ffff160361074157506002915060006107f4565b603c6107506020880188611ede565b61ffff1603610764575060019150816107f4565b603d6107736020880188611ede565b61ffff160361078857506002915060016107f4565b603e6107976020880188611ede565b61ffff16036107ac57506004915060016107f4565b60405162461bcd60e51b815260206004820152601b60248201527f494e56414c49445f4d454d4f52595f53544f52455f4f50434f4445000000000060448201526064016101d8565b60006108038960200151610b79565b90508160068111156108175761081761167f565b8151600681111561082a5761082a61167f565b146108775760405162461bcd60e51b815260206004820152600e60248201527f4241445f53544f52455f5459504500000000000000000000000000000000000060448201526064016101d8565b8060200151925060088467ffffffffffffffff1610156108c557600161089e856008611f2b565b67ffffffffffffffff16600167ffffffffffffffff16901b6108c09190611f57565b831692505b505060006108d96104848960200151610b79565b6108ed9063ffffffff166020880135611f18565b905086602001516000015167ffffffffffffffff168367ffffffffffffffff16826109189190611f18565b111561092a575050600286525061065e565b604080516020810190915260608152600090600019906000805b8767ffffffffffffffff168110156109f95760006109628288611f18565b90506000610971602083611f95565b90508581146109b65760001986146109985761098e858786610d44565b60208f0151604001525b6109a98e60200151828e8e8b610de0565b9098509196509094509250845b60006109c3602084611fa9565b90506109d085828c610e89565b945060088a67ffffffffffffffff16901c995050505080806109f190611fbd565b915050610944565b50610a05828483610d44565b60208c015160400152505050505050505050505050565b602084015151600090610a33906201000090611fd7565b604080518082018252600080825260209182018190528251808401909352825263ffffffff831682820152880151919250610a6e9190610d34565b505050505050565b602084015151600090610a8d906201000090611fd7565b90506000610aa16104848860200151610b79565b90506000610ab863ffffffff808416908516611f18565b905086602001516020015167ffffffffffffffff168111610b3d57610ae06201000082611ffe565b602088015167ffffffffffffffff9091169052610b38610b2d84604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b60208a015190610d34565b610b6f565b604080518082018252600080825260209182018190528251808401909352825263ffffffff90820152610b6f90610b2d565b5050505050505050565b60408051808201909152600080825260208201528151610b9890610f16565b92915050565b60208101516000908183516006811115610bba57610bba61167f565b14610c075760405162461bcd60e51b815260206004820152600760248201527f4e4f545f4933320000000000000000000000000000000000000000000000000060448201526064016101d8565b6401000000008110610b985760405162461bcd60e51b815260206004820152600760248201527f4241445f4933320000000000000000000000000000000000000000000000000060448201526064016101d8565b85516000908190819067ffffffffffffffff16610c78888a611f18565b1115610c8d5750600191506000905082610d28565b600019600080805b8a811015610d1b576000610ca9828e611f18565b90506000610cb8602083611f95565b9050858114610cd857610cce8f828e8e8e610de0565b509a509095509350845b6000610ce5602084611fa9565b9050610cf2846008611ffe565b610cfc8783611027565b60ff16901b851794505050508080610d1390611fbd565b915050610c95565b5060009550935085925050505b96509650969350505050565b8151610d4090826110a8565b5050565b6040517f4d656d6f7279206c6561663a00000000000000000000000000000000000000006020820152602c81018290526000908190604c01604051602081830303815290604052805190602001209050610dd58585836040518060400160405280601381526020017f4d656d6f7279206d65726b6c6520747265653a0000000000000000000000000081525061119c565b9150505b9392505050565b600080610df96040518060200160405280606081525090565b839150610e078686846112bf565b9093509150610e178686846112db565b925090506000610e28828986610d44565b905088604001518114610e7d5760405162461bcd60e51b815260206004820152600e60248201527f57524f4e475f4d454d5f524f4f5400000000000000000000000000000000000060448201526064016101d8565b50955095509592505050565b600060208310610edb5760405162461bcd60e51b815260206004820152601560248201527f4241445f5345545f4c4541465f425954455f494458000000000000000000000060448201526064016101d8565b600083610eea60016020612015565b610ef49190612015565b610eff906008611ffe565b60ff848116821b911b198616179150509392505050565b604080518082019091526000808252602082015281518051610f3a90600190612015565b81518110610f4a57610f4a612028565b6020026020010151905060006001836000015151610f689190612015565b67ffffffffffffffff811115610f8057610f80611944565b604051908082528060200260200182016040528015610fc557816020015b6040805180820190915260008082526020820152815260200190600190039081610f9e5790505b50905060005b8151811015611020578351805182908110610fe857610fe8612028565b602002602001015182828151811061100257611002612028565b6020026020010181905250808061101890611fbd565b915050610fcb565b5090915290565b6000602082106110795760405162461bcd60e51b815260206004820152601660248201527f4241445f50554c4c5f4c4541465f425954455f4944580000000000000000000060448201526064016101d8565b60008261108860016020612015565b6110929190612015565b61109d906008611ffe565b9390931c9392505050565b8151516000906110b9906001611f18565b67ffffffffffffffff8111156110d1576110d1611944565b60405190808252806020026020018201604052801561111657816020015b60408051808201909152600080825260208201528152602001906001900390816110ef5790505b50905060005b83515181101561117257835180518290811061113a5761113a612028565b602002602001015182828151811061115457611154612028565b6020026020010181905250808061116a90611fbd565b91505061111c565b5081818460000151518151811061118b5761118b612028565b602090810291909101015290915250565b8160005b8551518110156112685784600116600003611204578282876000015183815181106111cd576111cd612028565b60200260200101516040516020016111e79392919061203e565b60405160208183030381529060405280519060200120915061124f565b828660000151828151811061121b5761121b612028565b6020026020010151836040516020016112369392919061203e565b6040516020818303038152906040528051906020012091505b60019490941c938061126081611fbd565b9150506111a0565b5083156112b75760405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f5254000000000000000000000000000000000060448201526064016101d8565b949350505050565b600081816112ce8686846113b6565b9097909650945050505050565b6040805160208101909152606081528160006112f8868684611414565b92509050600060ff821667ffffffffffffffff81111561131a5761131a611944565b604051908082528060200260200182016040528015611343578160200160208202803683370190505b50905060005b8260ff168160ff16101561139a576113628888866112bf565b838360ff168151811061137757611377612028565b60200260200101819650828152505050808061139290612075565b915050611349565b5060405180602001604052808281525093505050935093915050565b600081815b602081101561140b57600883901b92508585838181106113dd576113dd612028565b919091013560f81c939093179250816113f581611fbd565b925050808061140390611fbd565b9150506113bb565b50935093915050565b60008184848281811061142957611429612028565b919091013560f81c925081905061143f81611fbd565b915050935093915050565b604080516101808101909152806000815260200161147f60408051606080820183529181019182529081526000602082015290565b81526040805180820182526000808252602080830191909152830152016114bd60408051606080820183529181019182529081526000602082015290565b81526020016114e2604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b611533612094565b565b60008083601f84011261154757600080fd5b50813567ffffffffffffffff81111561155f57600080fd5b60208301915083602082850101111561157757600080fd5b9250929050565b6000806000806000808688036101e081121561159957600080fd5b60608112156115a757600080fd5b879650606088013567ffffffffffffffff808211156115c557600080fd5b818a0191506101c080838d0312156115dc57600080fd5b8298506101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808501121561161057600080fd5b60808b01975060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe808501121561164657600080fd5b6101808b0196508a013592508083111561165f57600080fd5b505061166d89828a01611535565b979a9699509497509295939492505050565b634e487b7160e01b600052602160045260246000fd5b600381106116a5576116a561167f565b9052565b8051600781106116bb576116bb61167f565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611718576117048286516116a9565b9382019360019390930192908501906116f1565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b828110156117975784516117638582516116a9565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a09093019260010161174e565b509687015197909601969096525093949350505050565b60006101208083526117c38184018651611695565b60208501516101c061014081818701526117e16102e08701846116c8565b925060408801516101606118018189018380518252602090810151910152565b60608a015191507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee080898703016101a08a015261183e86846116c8565b955060808b015192508089870301858a01525061185b858361172c565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c0860152509150610dd99050602083018480518252602081015167ffffffffffffffff80825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561197d5761197d611944565b60405290565b6040516020810167ffffffffffffffff8111828210171561197d5761197d611944565b6040516080810167ffffffffffffffff8111828210171561197d5761197d611944565b604051610180810167ffffffffffffffff8111828210171561197d5761197d611944565b60405160c0810167ffffffffffffffff8111828210171561197d5761197d611944565b6040516060810167ffffffffffffffff8111828210171561197d5761197d611944565b604051601f8201601f1916810167ffffffffffffffff81118282101715611a5c57611a5c611944565b604052919050565b803560038110611a7357600080fd5b919050565b600067ffffffffffffffff821115611a9257611a92611944565b5060051b60200190565b600060408284031215611aae57600080fd5b611ab661195a565b9050813560078110611ac757600080fd5b808252506020820135602082015292915050565b60006040808385031215611aee57600080fd5b611af661195a565b9150823567ffffffffffffffff80821115611b1057600080fd5b81850191506020808388031215611b2657600080fd5b611b2e611983565b833583811115611b3d57600080fd5b80850194505087601f850112611b5257600080fd5b83359250611b67611b6284611a78565b611a33565b83815260069390931b84018201928281019089851115611b8657600080fd5b948301945b84861015611bac57611b9d8a87611a9c565b82529486019490830190611b8b565b8252508552948501359484019490945250909392505050565b600060408284031215611bd757600080fd5b611bdf61195a565b9050813581526020820135602082015292915050565b803563ffffffff81168114611a7357600080fd5b60006040808385031215611c1c57600080fd5b611c2461195a565b9150823567ffffffffffffffff811115611c3d57600080fd5b8301601f81018513611c4e57600080fd5b80356020611c5e611b6283611a78565b82815260a09283028401820192828201919089851115611c7d57600080fd5b948301945b84861015611ce65780868b031215611c9a5760008081fd5b611ca26119a6565b611cac8b88611a9c565b815287870135858201526060611cc3818901611bf5565b89830152611cd360808901611bf5565b9082015283529485019491830191611c82565b50808752505080860135818601525050505092915050565b60006101c08236031215611d1157600080fd5b611d196119c9565b611d2283611a64565b8152602083013567ffffffffffffffff80821115611d3f57600080fd5b611d4b36838701611adb565b6020840152611d5d3660408701611bc5565b60408401526080850135915080821115611d7657600080fd5b611d8236838701611adb565b606084015260a0850135915080821115611d9b57600080fd5b50611da836828601611c09565b608083015250611dbb3660c08501611bc5565b60a08201526101008084013560c0830152610120611dda818601611bf5565b60e0840152610140611ded818701611bf5565b838501526101609250611e01838701611bf5565b91840191909152610180850135908301526101a090930135928101929092525090565b803567ffffffffffffffff81168114611a7357600080fd5b6000818303610100811215611e5057600080fd5b611e586119ed565b833581526060601f1983011215611e6e57600080fd5b611e76611a10565b9150611e8460208501611e24565b8252611e9260408501611e24565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c08401356080820152611ed160e08501611bf5565b60a0820152949350505050565b600060208284031215611ef057600080fd5b813561ffff81168114610dd957600080fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610b9857610b98611f02565b67ffffffffffffffff818116838216028082169190828114611f4f57611f4f611f02565b505092915050565b67ffffffffffffffff828116828216039080821115611f7857611f78611f02565b5092915050565b634e487b7160e01b600052601260045260246000fd5b600082611fa457611fa4611f7f565b500490565b600082611fb857611fb8611f7f565b500690565b60006000198203611fd057611fd0611f02565b5060010190565b600067ffffffffffffffff80841680611ff257611ff2611f7f565b92169190910492915050565b8082028115828204841417610b9857610b98611f02565b81810381811115610b9857610b98611f02565b634e487b7160e01b600052603260045260246000fd5b6000845160005b8181101561205f5760208188018101518583015201612045565b5091909101928352506020820152604001919050565b600060ff821660ff810361208b5761208b611f02565b60010192915050565b634e487b7160e01b600052605160045260246000fdfea26469706673582212200b673d48d0bc6d32021f93260161f3386855f9ba2403d8cdd723e8406aaf6d3e64736f6c63430008110033",
}

// OneStepProverMemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMemoryMetaData.ABI instead.
var OneStepProverMemoryABI = OneStepProverMemoryMetaData.ABI

// OneStepProverMemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMemoryMetaData.Bin instead.
var OneStepProverMemoryBin = OneStepProverMemoryMetaData.Bin

// DeployOneStepProverMemory deploys a new Ethereum contract, binding an instance of OneStepProverMemory to it.
func DeployOneStepProverMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMemory, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// OneStepProverMemory is an auto generated Go binding around an Ethereum contract.
type OneStepProverMemory struct {
	OneStepProverMemoryCaller     // Read-only binding to the contract
	OneStepProverMemoryTransactor // Write-only binding to the contract
	OneStepProverMemoryFilterer   // Log filterer for contract events
}

// OneStepProverMemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMemorySession struct {
	Contract     *OneStepProverMemory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverMemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMemoryCallerSession struct {
	Contract *OneStepProverMemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverMemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMemoryTransactorSession struct {
	Contract     *OneStepProverMemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverMemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMemoryRaw struct {
	Contract *OneStepProverMemory // Generic contract binding to access the raw methods on
}

// OneStepProverMemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCallerRaw struct {
	Contract *OneStepProverMemoryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactorRaw struct {
	Contract *OneStepProverMemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMemory creates a new instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemory(address common.Address, backend bind.ContractBackend) (*OneStepProverMemory, error) {
	contract, err := bindOneStepProverMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// NewOneStepProverMemoryCaller creates a new read-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMemoryCaller, error) {
	contract, err := bindOneStepProverMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryCaller{contract: contract}, nil
}

// NewOneStepProverMemoryTransactor creates a new write-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMemoryTransactor, error) {
	contract, err := bindOneStepProverMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryTransactor{contract: contract}, nil
}

// NewOneStepProverMemoryFilterer creates a new log filterer instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMemoryFilterer, error) {
	contract, err := bindOneStepProverMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryFilterer{contract: contract}, nil
}

// bindOneStepProverMemory binds a generic wrapper to an already deployed contract.
func bindOneStepProverMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.OneStepProverMemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMemory.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemorySession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ReferenceDAProofValidatorMetaData contains all meta data concerning the ReferenceDAProofValidator contract.
var ReferenceDAProofValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"validateReadPreimage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"preimageChunk\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106fc806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d2f26f5214610030575b600080fd5b61004361003e3660046104e2565b610059565b6040516100509190610578565b60405180910390f35b606060528210156100b15760405162461bcd60e51b815260206004820152600f60248201527f50726f6f6620746f6f2073686f7274000000000000000000000000000000000060448201526064015b60405180910390fd5b8235602084013560c01c60006100cb6049602887896105ab565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525084519495508594909350151591506101179050576101176105d5565b6020910101517fff00000000000000000000000000000000000000000000000000000000000000167f0100000000000000000000000000000000000000000000000000000000000000146101ad5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c69642063657274696669636174652068656164657200000000000060448201526064016100a8565b8051602082012083146102025760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642063657274696669636174652068617368000000000000000060448201526064016100a8565b60218101518686604981811061021a5761021a6105d5565b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916600160f81b146102985760405162461bcd60e51b815260206004820152601960248201527f556e737570706f727465642070726f6f662076657273696f6e0000000000000060448201526064016100a8565b604a87013560c01c6102ab816052610633565b87146102f95760405162461bcd60e51b815260206004820152601460248201527f496e76616c69642070726f6f66206c656e67746800000000000000000000000060448201526064016100a8565b6000610308886052818c6105ab565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060405192935085926002925061035591508490602001610646565b60408051601f198184030181529082905261036f91610646565b602060405180830381855afa15801561038c573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103af9190610662565b146103fc5760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420707265696d6167652068617368000000000000000000000060448201526064016100a8565b84600061040a826020610633565b90508251811115610419575081515b6000610425838361067b565b60408051602080825281830190925291925060208201818036833701905050995080156104d25760005b818110156104d057846104628286610633565b81518110610472576104726105d5565b602001015160f81c60f81b8b828151811061048f5761048f6105d5565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350806104c88161068e565b91505061044f565b505b5050505050505050505b92915050565b600080602083850312156104f557600080fd5b823567ffffffffffffffff8082111561050d57600080fd5b818501915085601f83011261052157600080fd5b81358181111561053057600080fd5b86602082850101111561054257600080fd5b60209290920196919550909350505050565b60005b8381101561056f578181015183820152602001610557565b50506000910152565b6020815260008251806020840152610597816040850160208701610554565b601f01601f19169190910160400192915050565b600080858511156105bb57600080fd5b838611156105c857600080fd5b5050820193919092039150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104dc576104dc610604565b60008251610658818460208701610554565b9190910192915050565b60006020828403121561067457600080fd5b5051919050565b818103818111156104dc576104dc610604565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036106bf576106bf610604565b506001019056fea26469706673582212203d8302ec52d239050396db544aa6b16d3ebe888fe52cbec12a2cff4f449f9e5964736f6c63430008110033",
}

// ReferenceDAProofValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ReferenceDAProofValidatorMetaData.ABI instead.
var ReferenceDAProofValidatorABI = ReferenceDAProofValidatorMetaData.ABI

// ReferenceDAProofValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReferenceDAProofValidatorMetaData.Bin instead.
var ReferenceDAProofValidatorBin = ReferenceDAProofValidatorMetaData.Bin

// DeployReferenceDAProofValidator deploys a new Ethereum contract, binding an instance of ReferenceDAProofValidator to it.
func DeployReferenceDAProofValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReferenceDAProofValidator, error) {
	parsed, err := ReferenceDAProofValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReferenceDAProofValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReferenceDAProofValidator{ReferenceDAProofValidatorCaller: ReferenceDAProofValidatorCaller{contract: contract}, ReferenceDAProofValidatorTransactor: ReferenceDAProofValidatorTransactor{contract: contract}, ReferenceDAProofValidatorFilterer: ReferenceDAProofValidatorFilterer{contract: contract}}, nil
}

// ReferenceDAProofValidator is an auto generated Go binding around an Ethereum contract.
type ReferenceDAProofValidator struct {
	ReferenceDAProofValidatorCaller     // Read-only binding to the contract
	ReferenceDAProofValidatorTransactor // Write-only binding to the contract
	ReferenceDAProofValidatorFilterer   // Log filterer for contract events
}

// ReferenceDAProofValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReferenceDAProofValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferenceDAProofValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReferenceDAProofValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferenceDAProofValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReferenceDAProofValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferenceDAProofValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReferenceDAProofValidatorSession struct {
	Contract     *ReferenceDAProofValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReferenceDAProofValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReferenceDAProofValidatorCallerSession struct {
	Contract *ReferenceDAProofValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ReferenceDAProofValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReferenceDAProofValidatorTransactorSession struct {
	Contract     *ReferenceDAProofValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ReferenceDAProofValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReferenceDAProofValidatorRaw struct {
	Contract *ReferenceDAProofValidator // Generic contract binding to access the raw methods on
}

// ReferenceDAProofValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReferenceDAProofValidatorCallerRaw struct {
	Contract *ReferenceDAProofValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ReferenceDAProofValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReferenceDAProofValidatorTransactorRaw struct {
	Contract *ReferenceDAProofValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReferenceDAProofValidator creates a new instance of ReferenceDAProofValidator, bound to a specific deployed contract.
func NewReferenceDAProofValidator(address common.Address, backend bind.ContractBackend) (*ReferenceDAProofValidator, error) {
	contract, err := bindReferenceDAProofValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReferenceDAProofValidator{ReferenceDAProofValidatorCaller: ReferenceDAProofValidatorCaller{contract: contract}, ReferenceDAProofValidatorTransactor: ReferenceDAProofValidatorTransactor{contract: contract}, ReferenceDAProofValidatorFilterer: ReferenceDAProofValidatorFilterer{contract: contract}}, nil
}

// NewReferenceDAProofValidatorCaller creates a new read-only instance of ReferenceDAProofValidator, bound to a specific deployed contract.
func NewReferenceDAProofValidatorCaller(address common.Address, caller bind.ContractCaller) (*ReferenceDAProofValidatorCaller, error) {
	contract, err := bindReferenceDAProofValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReferenceDAProofValidatorCaller{contract: contract}, nil
}

// NewReferenceDAProofValidatorTransactor creates a new write-only instance of ReferenceDAProofValidator, bound to a specific deployed contract.
func NewReferenceDAProofValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ReferenceDAProofValidatorTransactor, error) {
	contract, err := bindReferenceDAProofValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReferenceDAProofValidatorTransactor{contract: contract}, nil
}

// NewReferenceDAProofValidatorFilterer creates a new log filterer instance of ReferenceDAProofValidator, bound to a specific deployed contract.
func NewReferenceDAProofValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ReferenceDAProofValidatorFilterer, error) {
	contract, err := bindReferenceDAProofValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReferenceDAProofValidatorFilterer{contract: contract}, nil
}

// bindReferenceDAProofValidator binds a generic wrapper to an already deployed contract.
func bindReferenceDAProofValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReferenceDAProofValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferenceDAProofValidator.Contract.ReferenceDAProofValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferenceDAProofValidator.Contract.ReferenceDAProofValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferenceDAProofValidator.Contract.ReferenceDAProofValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferenceDAProofValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferenceDAProofValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferenceDAProofValidator.Contract.contract.Transact(opts, method, params...)
}

// ValidateReadPreimage is a free data retrieval call binding the contract method 0xd2f26f52.
//
// Solidity: function validateReadPreimage(bytes proof) pure returns(bytes preimageChunk)
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorCaller) ValidateReadPreimage(opts *bind.CallOpts, proof []byte) ([]byte, error) {
	var out []interface{}
	err := _ReferenceDAProofValidator.contract.Call(opts, &out, "validateReadPreimage", proof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateReadPreimage is a free data retrieval call binding the contract method 0xd2f26f52.
//
// Solidity: function validateReadPreimage(bytes proof) pure returns(bytes preimageChunk)
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorSession) ValidateReadPreimage(proof []byte) ([]byte, error) {
	return _ReferenceDAProofValidator.Contract.ValidateReadPreimage(&_ReferenceDAProofValidator.CallOpts, proof)
}

// ValidateReadPreimage is a free data retrieval call binding the contract method 0xd2f26f52.
//
// Solidity: function validateReadPreimage(bytes proof) pure returns(bytes preimageChunk)
func (_ReferenceDAProofValidator *ReferenceDAProofValidatorCallerSession) ValidateReadPreimage(proof []byte) ([]byte, error) {
	return _ReferenceDAProofValidator.Contract.ValidateReadPreimage(&_ReferenceDAProofValidator.CallOpts, proof)
}
