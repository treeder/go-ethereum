// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross

import (
	"math/big"
	"strings"

	"github.com/gochain/gochain/v3"
	"github.com/gochain/gochain/v3/accounts/abi"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = gochain.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressSetABI is the input ABI used to generate the binding from.
const AddressSetABI = "[]"

// AddressSetFuncSigs maps the 4-byte function signature to its string representation.
var AddressSetFuncSigs = map[string]string{
	"e207783c": "majority(AddressSet.Set storage)",
}

// AddressSetBin is the compiled bytecode used for deploying new contracts.
const AddressSetBin = "0x60a4610024600b82828239805160001a607314601757fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063e207783c146038575b600080fd5b605260048036036020811015604c57600080fd5b50356064565b60408051918252519081900360200190f35b54600290046001019056fea265627a7a72315820267f258c55183a50510a66009d670f9c40830992e35ffde2d5e5f914b05a2b1d64736f6c63430005100032"

// DeployAddressSet deploys a new GoChain contract, binding an instance of AddressSet to it.
func DeployAddressSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressSet, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	addressSetBin := AddressSetBin

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(addressSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressSet{AddressSetCaller: AddressSetCaller{contract: contract}, AddressSetTransactor: AddressSetTransactor{contract: contract}, AddressSetFilterer: AddressSetFilterer{contract: contract}}, nil
}

// AddressSet is an auto generated Go binding around an GoChain contract.
type AddressSet struct {
	AddressSetCaller     // Read-only binding to the contract
	AddressSetTransactor // Write-only binding to the contract
	AddressSetFilterer   // Log filterer for contract events
}

// AddressSetCaller is an auto generated read-only Go binding around an GoChain contract.
type AddressSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSetTransactor is an auto generated write-only Go binding around an GoChain contract.
type AddressSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSetFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type AddressSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSetSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type AddressSetSession struct {
	Contract     *AddressSet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressSetCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type AddressSetCallerSession struct {
	Contract *AddressSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AddressSetTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type AddressSetTransactorSession struct {
	Contract     *AddressSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressSetRaw is an auto generated low-level Go binding around an GoChain contract.
type AddressSetRaw struct {
	Contract *AddressSet // Generic contract binding to access the raw methods on
}

// AddressSetCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type AddressSetCallerRaw struct {
	Contract *AddressSetCaller // Generic read-only contract binding to access the raw methods on
}

// AddressSetTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type AddressSetTransactorRaw struct {
	Contract *AddressSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressSet creates a new instance of AddressSet, bound to a specific deployed contract.
func NewAddressSet(address common.Address, backend bind.ContractBackend) (*AddressSet, error) {
	contract, err := bindAddressSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressSet{AddressSetCaller: AddressSetCaller{contract: contract}, AddressSetTransactor: AddressSetTransactor{contract: contract}, AddressSetFilterer: AddressSetFilterer{contract: contract}}, nil
}

// NewAddressSetCaller creates a new read-only instance of AddressSet, bound to a specific deployed contract.
func NewAddressSetCaller(address common.Address, caller bind.ContractCaller) (*AddressSetCaller, error) {
	contract, err := bindAddressSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressSetCaller{contract: contract}, nil
}

// NewAddressSetTransactor creates a new write-only instance of AddressSet, bound to a specific deployed contract.
func NewAddressSetTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressSetTransactor, error) {
	contract, err := bindAddressSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressSetTransactor{contract: contract}, nil
}

// NewAddressSetFilterer creates a new log filterer instance of AddressSet, bound to a specific deployed contract.
func NewAddressSetFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressSetFilterer, error) {
	contract, err := bindAddressSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressSetFilterer{contract: contract}, nil
}

// bindAddressSet binds a generic wrapper to an already deployed contract.
func bindAddressSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressSet *AddressSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressSet.Contract.AddressSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressSet *AddressSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressSet.Contract.AddressSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressSet *AddressSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressSet.Contract.AddressSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressSet *AddressSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressSet *AddressSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressSet *AddressSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressSet.Contract.contract.Transact(opts, method, params...)
}

// ConfirmationsABI is the input ABI used to generate the binding from.
const ConfirmationsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"Confirmed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_confirmedGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"pendingStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invalid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"shouldConfirm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumIConfirmations.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalConfirmGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ConfirmationsFuncSigs maps the 4-byte function signature to its string representation.
var ConfirmationsFuncSigs = map[string]string{
	"2ea6ece8": "_confirmedGas()",
	"8f7fd0d6": "confirm(uint256,uint256,bytes32,bool)",
	"42715aec": "confirmGas()",
	"7849ed5a": "count(address,bool,bool)",
	"3ffefe4e": "getSigner(uint256)",
	"d07bff0c": "getVoter(uint256)",
	"7df73e27": "isSigner(address)",
	"a7771ee3": "isVoter(address)",
	"03aca792": "pendingList(uint256)",
	"26b60630": "pendingListLength()",
	"6e06676b": "pendingStatus(uint256,uint256,bytes32)",
	"32030803": "request(uint256,uint256,bytes32)",
	"48c5000d": "setVote(address,bool,bool)",
	"859f32c2": "shouldConfirm(uint256,uint256,bytes32)",
	"41f684f3": "signersLength()",
	"fad3ffd6": "status(uint256,uint256,bytes32)",
	"5c5f6b39": "totalConfirmGas()",
	"6c6c39fb": "votersLength()",
	"d8bff5a5": "votes(address)",
}

// ConfirmationsBin is the compiled bytecode used for deploying new contracts.
const ConfirmationsBin = "0x60806040523480156200001157600080fd5b5060405162001bf338038062001bf3833981810160405260208110156200003757600080fd5b5051806200005360008262000073602090811b620012c917901c565b6200006b6002826200007360201b620012c91760201c565b505062000114565b6001600160a01b038116600090815260018301602052604090205415620000d2576040805162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481a5b881cd95d60921b604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b611acf80620001246000396000f3fe6080604052600436106101145760003560e01c80636c6c39fb116100a05780638f7fd0d6116100645780638f7fd0d6146103dc578063a7771ee31461041a578063d07bff0c1461044d578063d8bff5a514610477578063fad3ffd6146104d457610114565b80636c6c39fb146102905780636e06676b146102a55780637849ed5a146103015780637df73e2714610344578063859f32c21461038b57610114565b80633ffefe4e116100e75780633ffefe4e146101c857806341f684f31461020e57806342715aec1461022357806348c5000d146102385780635c5f6b391461027b57610114565b806303aca7921461011957806326b60630146101615780632ea6ece814610188578063320308031461019d575b600080fd5b34801561012557600080fd5b506101436004803603602081101561013c57600080fd5b503561052e565b60408051938452602084019290925282820152519081900360600190f35b34801561016d57600080fd5b5061017661055e565b60408051918252519081900360200190f35b34801561019457600080fd5b50610176610565565b6101c6600480360360608110156101b357600080fd5b508035906020810135906040013561056c565b005b3480156101d457600080fd5b506101f2600480360360208110156101eb57600080fd5b50356107ea565b604080516001600160a01b039092168252519081900360200190f35b34801561021a57600080fd5b50610176610817565b34801561022f57600080fd5b5061017661081d565b34801561024457600080fd5b506101c66004803603606081101561025b57600080fd5b506001600160a01b03813516906020810135151590604001351515610824565b34801561028757600080fd5b50610176610b05565b34801561029c57600080fd5b50610176610b0c565b3480156102b157600080fd5b506102db600480360360608110156102c857600080fd5b5080359060208101359060400135610b12565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561030d57600080fd5b506101766004803603606081101561032457600080fd5b506001600160a01b03813516906020810135151590604001351515610ba1565b34801561035057600080fd5b506103776004803603602081101561036757600080fd5b50356001600160a01b0316610bc4565b604080519115158252519081900360200190f35b34801561039757600080fd5b506103c1600480360360608110156103ae57600080fd5b5080359060208101359060400135610be1565b60408051921515835260208301919091528051918290030190f35b3480156103e857600080fd5b506101c6600480360360808110156103ff57600080fd5b50803590602081013590604081013590606001351515610cf4565b34801561042657600080fd5b506103776004803603602081101561043d57600080fd5b50356001600160a01b0316610efc565b34801561045957600080fd5b506101f26004803603602081101561047057600080fd5b5035610f19565b34801561048357600080fd5b506104aa6004803603602081101561049a57600080fd5b50356001600160a01b0316610f2a565b604080516001600160a01b0390941684529115156020840152151582820152519081900360600190f35b3480156104e057600080fd5b5061050a600480360360608110156104f757600080fd5b5080359060208101359060400135610f5b565b6040518082600381111561051a57fe5b60ff16815260200191505060405180910390f35b6008818154811061053b57fe5b600091825260209091206003909102018054600182015460029092015490925083565b6008545b90565b6202710081565b6000838152600660209081526040808320858452825280832084845290915281205460ff169081600381111561059e57fe5b14806105b5575060018160038111156105b357fe5b145b610606576040805162461bcd60e51b815260206004820152601e60248201527f537461747573206d757374206265204e6f6e65206f722050656e64696e670000604482015290519081900360640190fd5b622625a03a023481111561064b5760405162461bcd60e51b8152600401808060200182810382526027815260200180611a746027913960400191505060405180910390fd5b803411156106845760405133903483900380156108fc02916000818181858888f19350505050158015610682573d6000803e3d6000fd5b505b600082600381111561069257fe5b1415610784576000858152600660209081526040808320878452825280832086845282528083208054600160ff19909116811790915581516060810183528981528084018981528184018981526008805494850180825590885292517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee360039095029485015590517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee4840155517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee590920191909155888452600983528184208885528352818420878552909252909120555b6000858152600760209081526040808320878452825280832086845282529182902060040180543a01905581518581529151869288927f2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b0392918290030190a35050505050565b6000600260000182815481106107fc57fe5b6000918252602090912001546001600160a01b031692915050565b60025490565b62015f9081565b3360009081526001602052604090205461087a576040805162461bcd60e51b815260206004820152601260248201527129b2b73232b9103737ba1030903b37ba32b960711b604482015290519081900360640190fd5b33600090815260046020526040902080546001600160a01b03161561094e5780546001600160a01b0385811691161480156108c45750805460ff600160a01b909104161515831515145b80156108df5750805460ff600160a81b909104161515821515145b156108ea5750610b00565b80546001600160a01b038116600090815260056020908152604080832060ff600160a01b8604811615158552908352818420600160a81b90950416151583529281528282208054600019019055338252600490522080546001600160b01b03191690555b6001600160a01b0384166109625750610b00565b821561099f5781156109875761097784610efc565b156109825750610b00565b61099a565b61099084610efc565b61099a5750610b00565b6109cc565b81156109b9576109ae84610bc4565b1561099a5750610b00565b6109c284610bc4565b6109cc5750610b00565b604080516060810182526001600160a01b03808716808352861515602080850182815288151586880181815233600090815260048086528a822099518a54955193511515600160a81b0260ff60a81b19941515600160a01b0260ff60a01b1992909b166001600160a01b03199097169690961716989098179190911692909217909655928352600581528583209183529081528482209382529283528381208054600101908190558451633881de0f60e21b815292830191909152925173__$1c537defa8d8abb880144579cecc12eeda$__9263e207783c9260248082019391829003018186803b158015610ac057600080fd5b505af4158015610ad4573d6000803e3d6000fd5b505050506040513d6020811015610aea57600080fd5b50518110610afd57610afd858585610f81565b50505b505050565b622625a081565b60005490565b6000808080600160008881526006602090815260408083208a8452825280832089845290915290205460ff166003811115610b4957fe5b14610b5f57506000925082915081905080610b98565b50505060008481526007602090815260408083208684528252808320858452909152902060048101546002805483549190930154919350905b93509350935093565b600560209081526000938452604080852082529284528284209052825290205481565b6001600160a01b0316600090815260036020526040902054151590565b336000908152600360205260408120548190610c3a576040805162461bcd60e51b815260206004820152601360248201527229b2b73232b9103737ba10309039b4b3b732b960691b604482015290519081900360640190fd5b60016000868152600660209081526040808320888452825280832087845290915290205460ff166003811115610c6c57fe5b14610c7c57506000905080610cec565b600085815260076020908152604080832087845282528083208684528252808320338452600181019092529091205415610cbd575060009150819050610cec565b33600090815260038201602052604090205415610ce1575060009150819050610cec565b600401546001925090505b935093915050565b33600090815260036020526040902054610d4b576040805162461bcd60e51b815260206004820152601360248201527229b2b73232b9103737ba10309039b4b3b732b960691b604482015290519081900360640190fd5b60016000858152600660209081526040808320878452825280832086845290915290205460ff166003811115610d7d57fe5b14610dc4576040805162461bcd60e51b8152602060048201526012602482015271537461747573206e6f742050656e64696e6760701b604482015290519081900360640190fd5b600084815260076020908152604080832086845282528083208584529091529020808215610e155733600090815260038301602052604090205415610e1057610e1082600201336111b0565b610e3c565b50336000908152600182016020526040902054600282019015610e3c57610e3c82336111b0565b336000908152600182016020526040902054610ef357610e5c81336112c9565b73__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610eac57600080fd5b505af4158015610ec0573d6000803e3d6000fd5b505050506040513d6020811015610ed657600080fd5b505181541015610ee7575050610ef6565b610ef386868686611369565b50505b50505050565b6001600160a01b0316600090815260016020526040902054151590565b60008060000182815481106107fc57fe5b6004602052600090815260409020546001600160a01b0381169060ff600160a01b8204811691600160a81b90041683565b600660209081526000938452604080852082529284528284209052825290205460ff1681565b60005b60005481101561103e576000806000018281548110610f9f57fe5b60009182526020808320909101546001600160a01b0390811680845260049092526040909220805491935091908116908716148015610fed5750805460ff600160a01b909104161515851515145b80156110085750805460ff600160a81b909104161515841515145b15611034576001600160a01b038216600090815260046020526040902080546001600160b01b03191690555b5050600101610f84565b506001600160a01b03831660009081526005602090815260408083208515801585529083528184208515158552909252822091909155829061107d5750805b156110d7576001600160a01b0383166000908152600360205260409020546110aa576110aa6002846112c9565b6001600160a01b0383166000908152600160205260409020546110d2576110d26000846112c9565b610b00565b8180156110e2575080155b1561110e576001600160a01b038316600090815260016020526040902054156110d2576110d283611647565b811580156111195750805b15611146576001600160a01b0383166000908152600360205260409020546110d2576110d26002846112c9565b81158015611152575080155b15610b00576001600160a01b0383166000908152600160205260409020541561117e5761117e83611647565b6001600160a01b03831660009081526003602052604090205415610b00576111a76002846111b0565b610b0083611838565b6001600160a01b03811660009081526001830160205260409020548061120a576040805162461bcd60e51b815260206004820152600a602482015269139bdd081a5b881cd95d60b21b604482015290519081900360640190fd5b6001600160a01b0382166000908152600184016020526040812055825483600019820182811061123657fe5b60009182526020909120015484546001600160a01b03909116908590600019850190811061126057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508360000160018203815481106112a057fe5b600091825260209091200180546001600160a01b03191690558354610afd8560001983016119c0565b6001600160a01b038116600090815260018301602052604090205415611327576040805162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481a5b881cd95d60921b604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b600081156113795750600361137d565b5060025b600085815260066020908152604080832087845282528083208684529091529020805482919060ff191660018360038111156113b557fe5b021790555060008581526007602090815260408083208784528252808320868452909152902080836113e75750600281015b805460048301546223b4a00260008282816113fe57fe5b04905060005b8381101561146a57600085600001828154811061141d57fe5b60009182526020822001546040516001600160a01b039091169250829185156108fc02918691818181858888f19350505050158015611460573d6000803e3d6000fd5b5050600101611404565b506004850154604051848302840391620271000290339082840180156108fc02916000818181858888f193505050501580156114aa573d6000803e3d6000fd5b5060008c81526007602090815260408083208e845282528083208d845290915281209081816114d982826119e4565b50506002820160006114eb82826119e4565b50506000600492909201829055508c81526009602090815260408083208e845282528083208d84529091528120805491905560085460018111156115b15760006008600183038154811061153b57fe5b90600052602060002090600302019050806008600185038154811061155c57fe5b6000918252602080832084546003909302019182556001808501548184015560029485015492850192909255845483526009815260408084209286015484529181528183209490930154825292909152208290555b600860018203815481106115c157fe5b600091825260208220600390910201818155600181018290556002015560088054906115f1906000198301611a02565b508c8e7fdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c8e8e60405180838152602001821515151581526020019250505060405180910390a35050505050505050505050505050565b6116526000826111b0565b6001600160a01b0380821660009081526004602052604090208054909116156116dd5780546001600160a01b03808216600090815260056020908152604080832060ff600160a01b8704811615158552908352818420600160a81b90960416151583529381528382208054600019019055918516815260049091522080546001600160b01b03191690555b600073__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60006040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561172f57600080fd5b505af4158015611743573d6000803e3d6000fd5b505050506040513d602081101561175957600080fd5b5051905060005b600054811015610ef6576004600080600001838154811061177d57fe5b60009182526020808320909101546001600160a01b03908116845290830193909352604090910190208054909450161561182d5782546001600160a01b038116600090815260056020908152604080832060ff600160a01b8604811615158552908352818420600160a81b90950416151583529290522054821161182d578254611825906001600160a01b0381169060ff600160a01b8204811691600160a81b900416610f81565b505050611835565b600101611760565b50565b600073__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561188a57600080fd5b505af415801561189e573d6000803e3d6000fd5b505050506040513d60208110156118b457600080fd5b5051905060005b600854811015610b00576000600882815481106118d457fe5b6000918252602080832060039092029091018054835260078252604080842060018084015486529084528185206002840154865284528185206001600160a01b038a16865290810190935290922054919250901561193b5761193681866111b0565b611968565b6001600160a01b0385166000908152600382016020526040902054156119685761196881600201866111b0565b8054841161198e576119898260000154836001015484600201546001611369565b6119b9565b600281015484116119b2576119898260000154836001015484600201546000611369565b6001909201915b50506118bb565b815481835581811115610b0057600083815260209020610b00918101908301611a2e565b50805460008255906000526020600020908101906118359190611a2e565b815481835581811115610b0057600302816003028360005260206000209182019101610b009190611a4c565b61056291905b80821115611a485760008155600101611a34565b5090565b61056291905b80821115611a48576000808255600182018190556002820155600301611a5256fe54782076616c756520646f65736e277420636f76657220636f6e6669726d6174696f6e20666565a265627a7a723158200f4a6531c3f1a6ac425f5d60ff7b7f61b9aa82df1357b8c1b75c009d433475f864736f6c63430005100032"

// DeployConfirmations deploys a new GoChain contract, binding an instance of Confirmations to it.
func DeployConfirmations(auth *bind.TransactOpts, backend bind.ContractBackend, voter common.Address) (common.Address, *types.Transaction, *Confirmations, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	confirmationsBin := ConfirmationsBin

	if auth.Nonce == nil {
		if err := bind.EnsureNonce(auth, backend); err != nil {
			return common.Address{}, nil, nil, err
		}
	}

	addressSetAddr, _, _, err := DeployAddressSet(auth, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	confirmationsBin = strings.Replace(confirmationsBin, "__$1c537defa8d8abb880144579cecc12eeda$__", addressSetAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(confirmationsBin), backend, voter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Confirmations{ConfirmationsCaller: ConfirmationsCaller{contract: contract}, ConfirmationsTransactor: ConfirmationsTransactor{contract: contract}, ConfirmationsFilterer: ConfirmationsFilterer{contract: contract}}, nil
}

// Confirmations is an auto generated Go binding around an GoChain contract.
type Confirmations struct {
	ConfirmationsCaller     // Read-only binding to the contract
	ConfirmationsTransactor // Write-only binding to the contract
	ConfirmationsFilterer   // Log filterer for contract events
}

// ConfirmationsCaller is an auto generated read-only Go binding around an GoChain contract.
type ConfirmationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsTransactor is an auto generated write-only Go binding around an GoChain contract.
type ConfirmationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type ConfirmationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type ConfirmationsSession struct {
	Contract     *Confirmations    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfirmationsCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type ConfirmationsCallerSession struct {
	Contract *ConfirmationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConfirmationsTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type ConfirmationsTransactorSession struct {
	Contract     *ConfirmationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConfirmationsRaw is an auto generated low-level Go binding around an GoChain contract.
type ConfirmationsRaw struct {
	Contract *Confirmations // Generic contract binding to access the raw methods on
}

// ConfirmationsCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type ConfirmationsCallerRaw struct {
	Contract *ConfirmationsCaller // Generic read-only contract binding to access the raw methods on
}

// ConfirmationsTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type ConfirmationsTransactorRaw struct {
	Contract *ConfirmationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfirmations creates a new instance of Confirmations, bound to a specific deployed contract.
func NewConfirmations(address common.Address, backend bind.ContractBackend) (*Confirmations, error) {
	contract, err := bindConfirmations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Confirmations{ConfirmationsCaller: ConfirmationsCaller{contract: contract}, ConfirmationsTransactor: ConfirmationsTransactor{contract: contract}, ConfirmationsFilterer: ConfirmationsFilterer{contract: contract}}, nil
}

// NewConfirmationsCaller creates a new read-only instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsCaller(address common.Address, caller bind.ContractCaller) (*ConfirmationsCaller, error) {
	contract, err := bindConfirmations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsCaller{contract: contract}, nil
}

// NewConfirmationsTransactor creates a new write-only instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfirmationsTransactor, error) {
	contract, err := bindConfirmations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsTransactor{contract: contract}, nil
}

// NewConfirmationsFilterer creates a new log filterer instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfirmationsFilterer, error) {
	contract, err := bindConfirmations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsFilterer{contract: contract}, nil
}

// bindConfirmations binds a generic wrapper to an already deployed contract.
func bindConfirmations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirmations *ConfirmationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Confirmations.Contract.ConfirmationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirmations *ConfirmationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirmations.Contract.ConfirmationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirmations *ConfirmationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirmations.Contract.ConfirmationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirmations *ConfirmationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Confirmations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirmations *ConfirmationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirmations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirmations *ConfirmationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirmations.Contract.contract.Transact(opts, method, params...)
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) ConfirmedGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "_confirmedGas")
	return *ret0, err
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) ConfirmedGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmedGas(&_Confirmations.CallOpts)
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) ConfirmedGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmedGas(&_Confirmations.CallOpts)
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) ConfirmGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "confirmGas")
	return *ret0, err
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) ConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmGas(&_Confirmations.CallOpts)
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) ConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmGas(&_Confirmations.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) Count(opts *bind.CallOpts, arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "count", arg0, arg1, arg2)
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Confirmations.Contract.Count(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Confirmations.Contract.Count(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCaller) GetSigner(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "getSigner", i)
	return *ret0, err
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetSigner(&_Confirmations.CallOpts, i)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCallerSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetSigner(&_Confirmations.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCaller) GetVoter(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "getVoter", i)
	return *ret0, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetVoter(&_Confirmations.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCallerSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetVoter(&_Confirmations.CallOpts, i)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCaller) IsSigner(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "isSigner", voter)
	return *ret0, err
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsSession) IsSigner(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsSigner(&_Confirmations.CallOpts, voter)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCallerSession) IsSigner(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsSigner(&_Confirmations.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCaller) IsVoter(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "isVoter", voter)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsSession) IsVoter(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsVoter(&_Confirmations.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCallerSession) IsVoter(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsVoter(&_Confirmations.CallOpts, voter)
}

// PendingList is a free data retrieval call binding the contract method 0x03aca792.
//
// Solidity: function pendingList(uint256 ) constant returns(uint256 blockNum, uint256 logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsCaller) PendingList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
}, error) {
	ret := new(struct {
		BlockNum  *big.Int
		LogIndex  *big.Int
		EventHash [32]byte
	})
	out := ret
	err := _Confirmations.contract.Call(opts, out, "pendingList", arg0)
	return *ret, err
}

// PendingList is a free data retrieval call binding the contract method 0x03aca792.
//
// Solidity: function pendingList(uint256 ) constant returns(uint256 blockNum, uint256 logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsSession) PendingList(arg0 *big.Int) (struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
}, error) {
	return _Confirmations.Contract.PendingList(&_Confirmations.CallOpts, arg0)
}

// PendingList is a free data retrieval call binding the contract method 0x03aca792.
//
// Solidity: function pendingList(uint256 ) constant returns(uint256 blockNum, uint256 logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsCallerSession) PendingList(arg0 *big.Int) (struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
}, error) {
	return _Confirmations.Contract.PendingList(&_Confirmations.CallOpts, arg0)
}

// PendingListLength is a free data retrieval call binding the contract method 0x26b60630.
//
// Solidity: function pendingListLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) PendingListLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "pendingListLength")
	return *ret0, err
}

// PendingListLength is a free data retrieval call binding the contract method 0x26b60630.
//
// Solidity: function pendingListLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) PendingListLength() (*big.Int, error) {
	return _Confirmations.Contract.PendingListLength(&_Confirmations.CallOpts)
}

// PendingListLength is a free data retrieval call binding the contract method 0x26b60630.
//
// Solidity: function pendingListLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) PendingListLength() (*big.Int, error) {
	return _Confirmations.Contract.PendingListLength(&_Confirmations.CallOpts)
}

// PendingStatus is a free data retrieval call binding the contract method 0x6e06676b.
//
// Solidity: function pendingStatus(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(uint256 gasPrice, uint256 total, uint256 valid, uint256 invalid)
func (_Confirmations *ConfirmationsCaller) PendingStatus(opts *bind.CallOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (struct {
	GasPrice *big.Int
	Total    *big.Int
	Valid    *big.Int
	Invalid  *big.Int
}, error) {
	ret := new(struct {
		GasPrice *big.Int
		Total    *big.Int
		Valid    *big.Int
		Invalid  *big.Int
	})
	out := ret
	err := _Confirmations.contract.Call(opts, out, "pendingStatus", blockNum, logIndex, eventHash)
	return *ret, err
}

// PendingStatus is a free data retrieval call binding the contract method 0x6e06676b.
//
// Solidity: function pendingStatus(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(uint256 gasPrice, uint256 total, uint256 valid, uint256 invalid)
func (_Confirmations *ConfirmationsSession) PendingStatus(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (struct {
	GasPrice *big.Int
	Total    *big.Int
	Valid    *big.Int
	Invalid  *big.Int
}, error) {
	return _Confirmations.Contract.PendingStatus(&_Confirmations.CallOpts, blockNum, logIndex, eventHash)
}

// PendingStatus is a free data retrieval call binding the contract method 0x6e06676b.
//
// Solidity: function pendingStatus(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(uint256 gasPrice, uint256 total, uint256 valid, uint256 invalid)
func (_Confirmations *ConfirmationsCallerSession) PendingStatus(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (struct {
	GasPrice *big.Int
	Total    *big.Int
	Valid    *big.Int
	Invalid  *big.Int
}, error) {
	return _Confirmations.Contract.PendingStatus(&_Confirmations.CallOpts, blockNum, logIndex, eventHash)
}

// ShouldConfirm is a free data retrieval call binding the contract method 0x859f32c2.
//
// Solidity: function shouldConfirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(bool, uint256)
func (_Confirmations *ConfirmationsCaller) ShouldConfirm(opts *bind.CallOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Confirmations.contract.Call(opts, out, "shouldConfirm", blockNum, logIndex, eventHash)
	return *ret0, *ret1, err
}

// ShouldConfirm is a free data retrieval call binding the contract method 0x859f32c2.
//
// Solidity: function shouldConfirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(bool, uint256)
func (_Confirmations *ConfirmationsSession) ShouldConfirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (bool, *big.Int, error) {
	return _Confirmations.Contract.ShouldConfirm(&_Confirmations.CallOpts, blockNum, logIndex, eventHash)
}

// ShouldConfirm is a free data retrieval call binding the contract method 0x859f32c2.
//
// Solidity: function shouldConfirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(bool, uint256)
func (_Confirmations *ConfirmationsCallerSession) ShouldConfirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (bool, *big.Int, error) {
	return _Confirmations.Contract.ShouldConfirm(&_Confirmations.CallOpts, blockNum, logIndex, eventHash)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "signersLength")
	return *ret0, err
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) SignersLength() (*big.Int, error) {
	return _Confirmations.Contract.SignersLength(&_Confirmations.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) SignersLength() (*big.Int, error) {
	return _Confirmations.Contract.SignersLength(&_Confirmations.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsCaller) Status(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "status", arg0, arg1, arg2)
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsSession) Status(arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	return _Confirmations.Contract.Status(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsCallerSession) Status(arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	return _Confirmations.Contract.Status(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) TotalConfirmGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "totalConfirmGas")
	return *ret0, err
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) TotalConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.TotalConfirmGas(&_Confirmations.CallOpts)
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) TotalConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.TotalConfirmGas(&_Confirmations.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) VotersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "votersLength")
	return *ret0, err
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) VotersLength() (*big.Int, error) {
	return _Confirmations.Contract.VotersLength(&_Confirmations.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) VotersLength() (*big.Int, error) {
	return _Confirmations.Contract.VotersLength(&_Confirmations.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	ret := new(struct {
		Addr  common.Address
		Voter bool
		Add   bool
	})
	out := ret
	err := _Confirmations.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Confirmations.Contract.Votes(&_Confirmations.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsCallerSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Confirmations.Contract.Votes(&_Confirmations.CallOpts, arg0)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsTransactor) Confirm(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "confirm", blockNum, logIndex, eventHash, valid)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsSession) Confirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.Contract.Confirm(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash, valid)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsTransactorSession) Confirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.Contract.Confirm(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash, valid)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsTransactor) Request(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "request", blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.Contract.Request(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsTransactorSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.Contract.Request(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsTransactor) SetVote(opts *bind.TransactOpts, addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "setVote", addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.Contract.SetVote(&_Confirmations.TransactOpts, addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsTransactorSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.Contract.SetVote(&_Confirmations.TransactOpts, addr, voter, add)
}

// ConfirmationsConfirmationRequestedIterator is returned from FilterConfirmationRequested and is used to iterate over the raw logs and unpacked data for ConfirmationRequested events raised by the Confirmations contract.
type ConfirmationsConfirmationRequestedIterator struct {
	Event *ConfirmationsConfirmationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfirmationsConfirmationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmationsConfirmationRequested)
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
		it.Event = new(ConfirmationsConfirmationRequested)
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
func (it *ConfirmationsConfirmationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmationsConfirmationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmationsConfirmationRequested represents a ConfirmationRequested event raised by the Confirmations contract.
type ConfirmationsConfirmationRequested struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmationRequested is a free log retrieval operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) FilterConfirmationRequested(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*ConfirmationsConfirmationRequestedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.FilterLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsConfirmationRequestedIterator{contract: _Confirmations.contract, event: "ConfirmationRequested", logs: logs, sub: sub}, nil
}

// WatchConfirmationRequested is a free log subscription operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) WatchConfirmationRequested(opts *bind.WatchOpts, sink chan<- *ConfirmationsConfirmationRequested, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.WatchLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmationsConfirmationRequested)
				if err := _Confirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
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

// ParseConfirmationRequested is a log parse operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) ParseConfirmationRequested(log types.Log) (*ConfirmationsConfirmationRequested, error) {
	event := new(ConfirmationsConfirmationRequested)
	if err := _Confirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConfirmationsConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the Confirmations contract.
type ConfirmationsConfirmedIterator struct {
	Event *ConfirmationsConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfirmationsConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmationsConfirmed)
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
		it.Event = new(ConfirmationsConfirmed)
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
func (it *ConfirmationsConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmationsConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmationsConfirmed represents a Confirmed event raised by the Confirmations contract.
type ConfirmationsConfirmed struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Valid     bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) FilterConfirmed(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*ConfirmationsConfirmedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.FilterLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsConfirmedIterator{contract: _Confirmations.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *ConfirmationsConfirmed, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.WatchLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmationsConfirmed)
				if err := _Confirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// ParseConfirmed is a log parse operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) ParseConfirmed(log types.Log) (*ConfirmationsConfirmed, error) {
	event := new(ConfirmationsConfirmed)
	if err := _Confirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}
