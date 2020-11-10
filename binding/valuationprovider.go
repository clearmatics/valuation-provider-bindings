// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package valuationprovider

import (
	"math/big"
	"strings"

	ethereum "github.com/clearmatics/autonity"
	"github.com/clearmatics/autonity/accounts/abi"
	"github.com/clearmatics/autonity/accounts/abi/bind"
	"github.com/clearmatics/autonity/common"
	"github.com/clearmatics/autonity/core/types"
	"github.com/clearmatics/autonity/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ICMCABI is the input ABI used to generate the binding from.
const ICMCABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"valuation\",\"type\":\"bytes\"}],\"name\":\"executeEnforcement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ICMCFuncSigs maps the 4-byte function signature to its string representation.
var ICMCFuncSigs = map[string]string{
	"e4a74092": "executeEnforcement(bytes)",
}

// ICMC is an auto generated Go binding around an Ethereum contract.
type ICMC struct {
	ICMCCaller     // Read-only binding to the contract
	ICMCTransactor // Write-only binding to the contract
	ICMCFilterer   // Log filterer for contract events
}

// ICMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICMCSession struct {
	Contract     *ICMC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICMCCallerSession struct {
	Contract *ICMCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ICMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICMCTransactorSession struct {
	Contract     *ICMCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICMCRaw struct {
	Contract *ICMC // Generic contract binding to access the raw methods on
}

// ICMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICMCCallerRaw struct {
	Contract *ICMCCaller // Generic read-only contract binding to access the raw methods on
}

// ICMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICMCTransactorRaw struct {
	Contract *ICMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICMC creates a new instance of ICMC, bound to a specific deployed contract.
func NewICMC(address common.Address, backend bind.ContractBackend) (*ICMC, error) {
	contract, err := bindICMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICMC{ICMCCaller: ICMCCaller{contract: contract}, ICMCTransactor: ICMCTransactor{contract: contract}, ICMCFilterer: ICMCFilterer{contract: contract}}, nil
}

// NewICMCCaller creates a new read-only instance of ICMC, bound to a specific deployed contract.
func NewICMCCaller(address common.Address, caller bind.ContractCaller) (*ICMCCaller, error) {
	contract, err := bindICMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICMCCaller{contract: contract}, nil
}

// NewICMCTransactor creates a new write-only instance of ICMC, bound to a specific deployed contract.
func NewICMCTransactor(address common.Address, transactor bind.ContractTransactor) (*ICMCTransactor, error) {
	contract, err := bindICMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICMCTransactor{contract: contract}, nil
}

// NewICMCFilterer creates a new log filterer instance of ICMC, bound to a specific deployed contract.
func NewICMCFilterer(address common.Address, filterer bind.ContractFilterer) (*ICMCFilterer, error) {
	contract, err := bindICMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICMCFilterer{contract: contract}, nil
}

// bindICMC binds a generic wrapper to an already deployed contract.
func bindICMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICMC *ICMCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICMC.Contract.ICMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICMC *ICMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICMC.Contract.ICMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICMC *ICMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICMC.Contract.ICMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICMC *ICMCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ICMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICMC *ICMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICMC *ICMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICMC.Contract.contract.Transact(opts, method, params...)
}

// ExecuteEnforcement is a paid mutator transaction binding the contract method 0xe4a74092.
//
// Solidity: function executeEnforcement(bytes valuation) returns()
func (_ICMC *ICMCTransactor) ExecuteEnforcement(opts *bind.TransactOpts, valuation []byte) (*types.Transaction, error) {
	return _ICMC.contract.Transact(opts, "executeEnforcement", valuation)
}

// ExecuteEnforcement is a paid mutator transaction binding the contract method 0xe4a74092.
//
// Solidity: function executeEnforcement(bytes valuation) returns()
func (_ICMC *ICMCSession) ExecuteEnforcement(valuation []byte) (*types.Transaction, error) {
	return _ICMC.Contract.ExecuteEnforcement(&_ICMC.TransactOpts, valuation)
}

// ExecuteEnforcement is a paid mutator transaction binding the contract method 0xe4a74092.
//
// Solidity: function executeEnforcement(bytes valuation) returns()
func (_ICMC *ICMCTransactorSession) ExecuteEnforcement(valuation []byte) (*types.Transaction, error) {
	return _ICMC.Contract.ExecuteEnforcement(&_ICMC.TransactOpts, valuation)
}

// IValuationProviderABI is the input ABI used to generate the binding from.
const IValuationProviderABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ValuationRequestReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"valuation\",\"type\":\"bytes\"}],\"name\":\"ValuationRequestSubmitted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestValuation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"valuation\",\"type\":\"bytes\"}],\"name\":\"submitValuation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IValuationProviderFuncSigs maps the 4-byte function signature to its string representation.
var IValuationProviderFuncSigs = map[string]string{
	"b16b1f9d": "requestValuation()",
	"a333d322": "submitValuation(address,bytes)",
}

// IValuationProvider is an auto generated Go binding around an Ethereum contract.
type IValuationProvider struct {
	IValuationProviderCaller     // Read-only binding to the contract
	IValuationProviderTransactor // Write-only binding to the contract
	IValuationProviderFilterer   // Log filterer for contract events
}

// IValuationProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IValuationProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValuationProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IValuationProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValuationProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IValuationProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValuationProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IValuationProviderSession struct {
	Contract     *IValuationProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IValuationProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IValuationProviderCallerSession struct {
	Contract *IValuationProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IValuationProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IValuationProviderTransactorSession struct {
	Contract     *IValuationProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IValuationProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IValuationProviderRaw struct {
	Contract *IValuationProvider // Generic contract binding to access the raw methods on
}

// IValuationProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IValuationProviderCallerRaw struct {
	Contract *IValuationProviderCaller // Generic read-only contract binding to access the raw methods on
}

// IValuationProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IValuationProviderTransactorRaw struct {
	Contract *IValuationProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIValuationProvider creates a new instance of IValuationProvider, bound to a specific deployed contract.
func NewIValuationProvider(address common.Address, backend bind.ContractBackend) (*IValuationProvider, error) {
	contract, err := bindIValuationProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IValuationProvider{IValuationProviderCaller: IValuationProviderCaller{contract: contract}, IValuationProviderTransactor: IValuationProviderTransactor{contract: contract}, IValuationProviderFilterer: IValuationProviderFilterer{contract: contract}}, nil
}

// NewIValuationProviderCaller creates a new read-only instance of IValuationProvider, bound to a specific deployed contract.
func NewIValuationProviderCaller(address common.Address, caller bind.ContractCaller) (*IValuationProviderCaller, error) {
	contract, err := bindIValuationProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IValuationProviderCaller{contract: contract}, nil
}

// NewIValuationProviderTransactor creates a new write-only instance of IValuationProvider, bound to a specific deployed contract.
func NewIValuationProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*IValuationProviderTransactor, error) {
	contract, err := bindIValuationProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IValuationProviderTransactor{contract: contract}, nil
}

// NewIValuationProviderFilterer creates a new log filterer instance of IValuationProvider, bound to a specific deployed contract.
func NewIValuationProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*IValuationProviderFilterer, error) {
	contract, err := bindIValuationProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IValuationProviderFilterer{contract: contract}, nil
}

// bindIValuationProvider binds a generic wrapper to an already deployed contract.
func bindIValuationProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IValuationProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValuationProvider *IValuationProviderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IValuationProvider.Contract.IValuationProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValuationProvider *IValuationProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValuationProvider.Contract.IValuationProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValuationProvider *IValuationProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValuationProvider.Contract.IValuationProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValuationProvider *IValuationProviderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IValuationProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValuationProvider *IValuationProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValuationProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValuationProvider *IValuationProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValuationProvider.Contract.contract.Transact(opts, method, params...)
}

// RequestValuation is a paid mutator transaction binding the contract method 0xb16b1f9d.
//
// Solidity: function requestValuation() returns()
func (_IValuationProvider *IValuationProviderTransactor) RequestValuation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValuationProvider.contract.Transact(opts, "requestValuation")
}

// RequestValuation is a paid mutator transaction binding the contract method 0xb16b1f9d.
//
// Solidity: function requestValuation() returns()
func (_IValuationProvider *IValuationProviderSession) RequestValuation() (*types.Transaction, error) {
	return _IValuationProvider.Contract.RequestValuation(&_IValuationProvider.TransactOpts)
}

// RequestValuation is a paid mutator transaction binding the contract method 0xb16b1f9d.
//
// Solidity: function requestValuation() returns()
func (_IValuationProvider *IValuationProviderTransactorSession) RequestValuation() (*types.Transaction, error) {
	return _IValuationProvider.Contract.RequestValuation(&_IValuationProvider.TransactOpts)
}

// SubmitValuation is a paid mutator transaction binding the contract method 0xa333d322.
//
// Solidity: function submitValuation(address to, bytes valuation) returns()
func (_IValuationProvider *IValuationProviderTransactor) SubmitValuation(opts *bind.TransactOpts, to common.Address, valuation []byte) (*types.Transaction, error) {
	return _IValuationProvider.contract.Transact(opts, "submitValuation", to, valuation)
}

// SubmitValuation is a paid mutator transaction binding the contract method 0xa333d322.
//
// Solidity: function submitValuation(address to, bytes valuation) returns()
func (_IValuationProvider *IValuationProviderSession) SubmitValuation(to common.Address, valuation []byte) (*types.Transaction, error) {
	return _IValuationProvider.Contract.SubmitValuation(&_IValuationProvider.TransactOpts, to, valuation)
}

// SubmitValuation is a paid mutator transaction binding the contract method 0xa333d322.
//
// Solidity: function submitValuation(address to, bytes valuation) returns()
func (_IValuationProvider *IValuationProviderTransactorSession) SubmitValuation(to common.Address, valuation []byte) (*types.Transaction, error) {
	return _IValuationProvider.Contract.SubmitValuation(&_IValuationProvider.TransactOpts, to, valuation)
}

// IValuationProviderValuationRequestReceivedIterator is returned from FilterValuationRequestReceived and is used to iterate over the raw logs and unpacked data for ValuationRequestReceived events raised by the IValuationProvider contract.
type IValuationProviderValuationRequestReceivedIterator struct {
	Event *IValuationProviderValuationRequestReceived // Event containing the contract specifics and raw log

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
func (it *IValuationProviderValuationRequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValuationProviderValuationRequestReceived)
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
		it.Event = new(IValuationProviderValuationRequestReceived)
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
func (it *IValuationProviderValuationRequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValuationProviderValuationRequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValuationProviderValuationRequestReceived represents a ValuationRequestReceived event raised by the IValuationProvider contract.
type IValuationProviderValuationRequestReceived struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValuationRequestReceived is a free log retrieval operation binding the contract event 0x9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af.
//
// Solidity: event ValuationRequestReceived(address from)
func (_IValuationProvider *IValuationProviderFilterer) FilterValuationRequestReceived(opts *bind.FilterOpts) (*IValuationProviderValuationRequestReceivedIterator, error) {

	logs, sub, err := _IValuationProvider.contract.FilterLogs(opts, "ValuationRequestReceived")
	if err != nil {
		return nil, err
	}
	return &IValuationProviderValuationRequestReceivedIterator{contract: _IValuationProvider.contract, event: "ValuationRequestReceived", logs: logs, sub: sub}, nil
}

// WatchValuationRequestReceived is a free log subscription operation binding the contract event 0x9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af.
//
// Solidity: event ValuationRequestReceived(address from)
func (_IValuationProvider *IValuationProviderFilterer) WatchValuationRequestReceived(opts *bind.WatchOpts, sink chan<- *IValuationProviderValuationRequestReceived) (event.Subscription, error) {

	logs, sub, err := _IValuationProvider.contract.WatchLogs(opts, "ValuationRequestReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValuationProviderValuationRequestReceived)
				if err := _IValuationProvider.contract.UnpackLog(event, "ValuationRequestReceived", log); err != nil {
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

// ParseValuationRequestReceived is a log parse operation binding the contract event 0x9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af.
//
// Solidity: event ValuationRequestReceived(address from)
func (_IValuationProvider *IValuationProviderFilterer) ParseValuationRequestReceived(log types.Log) (*IValuationProviderValuationRequestReceived, error) {
	event := new(IValuationProviderValuationRequestReceived)
	if err := _IValuationProvider.contract.UnpackLog(event, "ValuationRequestReceived", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IValuationProviderValuationRequestSubmittedIterator is returned from FilterValuationRequestSubmitted and is used to iterate over the raw logs and unpacked data for ValuationRequestSubmitted events raised by the IValuationProvider contract.
type IValuationProviderValuationRequestSubmittedIterator struct {
	Event *IValuationProviderValuationRequestSubmitted // Event containing the contract specifics and raw log

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
func (it *IValuationProviderValuationRequestSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValuationProviderValuationRequestSubmitted)
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
		it.Event = new(IValuationProviderValuationRequestSubmitted)
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
func (it *IValuationProviderValuationRequestSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValuationProviderValuationRequestSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValuationProviderValuationRequestSubmitted represents a ValuationRequestSubmitted event raised by the IValuationProvider contract.
type IValuationProviderValuationRequestSubmitted struct {
	To        common.Address
	Valuation []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValuationRequestSubmitted is a free log retrieval operation binding the contract event 0x7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d.
//
// Solidity: event ValuationRequestSubmitted(address to, bytes valuation)
func (_IValuationProvider *IValuationProviderFilterer) FilterValuationRequestSubmitted(opts *bind.FilterOpts) (*IValuationProviderValuationRequestSubmittedIterator, error) {

	logs, sub, err := _IValuationProvider.contract.FilterLogs(opts, "ValuationRequestSubmitted")
	if err != nil {
		return nil, err
	}
	return &IValuationProviderValuationRequestSubmittedIterator{contract: _IValuationProvider.contract, event: "ValuationRequestSubmitted", logs: logs, sub: sub}, nil
}

// WatchValuationRequestSubmitted is a free log subscription operation binding the contract event 0x7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d.
//
// Solidity: event ValuationRequestSubmitted(address to, bytes valuation)
func (_IValuationProvider *IValuationProviderFilterer) WatchValuationRequestSubmitted(opts *bind.WatchOpts, sink chan<- *IValuationProviderValuationRequestSubmitted) (event.Subscription, error) {

	logs, sub, err := _IValuationProvider.contract.WatchLogs(opts, "ValuationRequestSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValuationProviderValuationRequestSubmitted)
				if err := _IValuationProvider.contract.UnpackLog(event, "ValuationRequestSubmitted", log); err != nil {
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

// ParseValuationRequestSubmitted is a log parse operation binding the contract event 0x7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d.
//
// Solidity: event ValuationRequestSubmitted(address to, bytes valuation)
func (_IValuationProvider *IValuationProviderFilterer) ParseValuationRequestSubmitted(log types.Log) (*IValuationProviderValuationRequestSubmitted, error) {
	event := new(IValuationProviderValuationRequestSubmitted)
	if err := _IValuationProvider.contract.UnpackLog(event, "ValuationRequestSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValuationProviderABI is the input ABI used to generate the binding from.
const ValuationProviderABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ValuationRequestReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"valuation\",\"type\":\"bytes\"}],\"name\":\"ValuationRequestSubmitted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestValuation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"valuation\",\"type\":\"bytes\"}],\"name\":\"submitValuation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValuationProviderFuncSigs maps the 4-byte function signature to its string representation.
var ValuationProviderFuncSigs = map[string]string{
	"b16b1f9d": "requestValuation()",
	"a333d322": "submitValuation(address,bytes)",
}

// ValuationProviderBin is the compiled bytecode used for deploying new contracts.
var ValuationProviderBin = "0x608060405234801561001057600080fd5b5061023a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a333d3221461003b578063b16b1f9d146100bd575b600080fd5b6100bb6004803603604081101561005157600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561007c57600080fd5b82018360208201111561008e57600080fd5b803590602001918460018302840111640100000000831117156100b057600080fd5b5090925090506100c5565b005b6100bb6101d0565b604051637253a04960e11b8152602060048201908152602482018390526001600160a01b0385169163e4a740929185918591908190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561013957600080fd5b505af115801561014d573d6000803e3d6000fd5b505050507f7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d83838360405180846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6040805133815290517f9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af9181900360200190a156fea265627a7a723158201a05323d412ca727bd830d532d3294a0bee84059ae171798009c3c85a147f1c164736f6c63430005110032"

// DeployValuationProvider deploys a new Ethereum contract, binding an instance of ValuationProvider to it.
func DeployValuationProvider(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValuationProvider, error) {
	parsed, err := abi.JSON(strings.NewReader(ValuationProviderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValuationProviderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValuationProvider{ValuationProviderCaller: ValuationProviderCaller{contract: contract}, ValuationProviderTransactor: ValuationProviderTransactor{contract: contract}, ValuationProviderFilterer: ValuationProviderFilterer{contract: contract}}, nil
}

// ValuationProvider is an auto generated Go binding around an Ethereum contract.
type ValuationProvider struct {
	ValuationProviderCaller     // Read-only binding to the contract
	ValuationProviderTransactor // Write-only binding to the contract
	ValuationProviderFilterer   // Log filterer for contract events
}

// ValuationProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValuationProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValuationProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValuationProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValuationProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValuationProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValuationProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValuationProviderSession struct {
	Contract     *ValuationProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValuationProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValuationProviderCallerSession struct {
	Contract *ValuationProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValuationProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValuationProviderTransactorSession struct {
	Contract     *ValuationProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValuationProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValuationProviderRaw struct {
	Contract *ValuationProvider // Generic contract binding to access the raw methods on
}

// ValuationProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValuationProviderCallerRaw struct {
	Contract *ValuationProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ValuationProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValuationProviderTransactorRaw struct {
	Contract *ValuationProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValuationProvider creates a new instance of ValuationProvider, bound to a specific deployed contract.
func NewValuationProvider(address common.Address, backend bind.ContractBackend) (*ValuationProvider, error) {
	contract, err := bindValuationProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValuationProvider{ValuationProviderCaller: ValuationProviderCaller{contract: contract}, ValuationProviderTransactor: ValuationProviderTransactor{contract: contract}, ValuationProviderFilterer: ValuationProviderFilterer{contract: contract}}, nil
}

// NewValuationProviderCaller creates a new read-only instance of ValuationProvider, bound to a specific deployed contract.
func NewValuationProviderCaller(address common.Address, caller bind.ContractCaller) (*ValuationProviderCaller, error) {
	contract, err := bindValuationProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValuationProviderCaller{contract: contract}, nil
}

// NewValuationProviderTransactor creates a new write-only instance of ValuationProvider, bound to a specific deployed contract.
func NewValuationProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ValuationProviderTransactor, error) {
	contract, err := bindValuationProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValuationProviderTransactor{contract: contract}, nil
}

// NewValuationProviderFilterer creates a new log filterer instance of ValuationProvider, bound to a specific deployed contract.
func NewValuationProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ValuationProviderFilterer, error) {
	contract, err := bindValuationProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValuationProviderFilterer{contract: contract}, nil
}

// bindValuationProvider binds a generic wrapper to an already deployed contract.
func bindValuationProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValuationProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValuationProvider *ValuationProviderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValuationProvider.Contract.ValuationProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValuationProvider *ValuationProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValuationProvider.Contract.ValuationProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValuationProvider *ValuationProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValuationProvider.Contract.ValuationProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValuationProvider *ValuationProviderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValuationProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValuationProvider *ValuationProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValuationProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValuationProvider *ValuationProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValuationProvider.Contract.contract.Transact(opts, method, params...)
}

// RequestValuation is a paid mutator transaction binding the contract method 0xb16b1f9d.
//
// Solidity: function requestValuation() returns()
func (_ValuationProvider *ValuationProviderTransactor) RequestValuation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValuationProvider.contract.Transact(opts, "requestValuation")
}

// RequestValuation is a paid mutator transaction binding the contract method 0xb16b1f9d.
//
// Solidity: function requestValuation() returns()
func (_ValuationProvider *ValuationProviderSession) RequestValuation() (*types.Transaction, error) {
	return _ValuationProvider.Contract.RequestValuation(&_ValuationProvider.TransactOpts)
}

// RequestValuation is a paid mutator transaction binding the contract method 0xb16b1f9d.
//
// Solidity: function requestValuation() returns()
func (_ValuationProvider *ValuationProviderTransactorSession) RequestValuation() (*types.Transaction, error) {
	return _ValuationProvider.Contract.RequestValuation(&_ValuationProvider.TransactOpts)
}

// SubmitValuation is a paid mutator transaction binding the contract method 0xa333d322.
//
// Solidity: function submitValuation(address to, bytes valuation) returns()
func (_ValuationProvider *ValuationProviderTransactor) SubmitValuation(opts *bind.TransactOpts, to common.Address, valuation []byte) (*types.Transaction, error) {
	return _ValuationProvider.contract.Transact(opts, "submitValuation", to, valuation)
}

// SubmitValuation is a paid mutator transaction binding the contract method 0xa333d322.
//
// Solidity: function submitValuation(address to, bytes valuation) returns()
func (_ValuationProvider *ValuationProviderSession) SubmitValuation(to common.Address, valuation []byte) (*types.Transaction, error) {
	return _ValuationProvider.Contract.SubmitValuation(&_ValuationProvider.TransactOpts, to, valuation)
}

// SubmitValuation is a paid mutator transaction binding the contract method 0xa333d322.
//
// Solidity: function submitValuation(address to, bytes valuation) returns()
func (_ValuationProvider *ValuationProviderTransactorSession) SubmitValuation(to common.Address, valuation []byte) (*types.Transaction, error) {
	return _ValuationProvider.Contract.SubmitValuation(&_ValuationProvider.TransactOpts, to, valuation)
}

// ValuationProviderValuationRequestReceivedIterator is returned from FilterValuationRequestReceived and is used to iterate over the raw logs and unpacked data for ValuationRequestReceived events raised by the ValuationProvider contract.
type ValuationProviderValuationRequestReceivedIterator struct {
	Event *ValuationProviderValuationRequestReceived // Event containing the contract specifics and raw log

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
func (it *ValuationProviderValuationRequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValuationProviderValuationRequestReceived)
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
		it.Event = new(ValuationProviderValuationRequestReceived)
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
func (it *ValuationProviderValuationRequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValuationProviderValuationRequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValuationProviderValuationRequestReceived represents a ValuationRequestReceived event raised by the ValuationProvider contract.
type ValuationProviderValuationRequestReceived struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValuationRequestReceived is a free log retrieval operation binding the contract event 0x9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af.
//
// Solidity: event ValuationRequestReceived(address from)
func (_ValuationProvider *ValuationProviderFilterer) FilterValuationRequestReceived(opts *bind.FilterOpts) (*ValuationProviderValuationRequestReceivedIterator, error) {

	logs, sub, err := _ValuationProvider.contract.FilterLogs(opts, "ValuationRequestReceived")
	if err != nil {
		return nil, err
	}
	return &ValuationProviderValuationRequestReceivedIterator{contract: _ValuationProvider.contract, event: "ValuationRequestReceived", logs: logs, sub: sub}, nil
}

// WatchValuationRequestReceived is a free log subscription operation binding the contract event 0x9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af.
//
// Solidity: event ValuationRequestReceived(address from)
func (_ValuationProvider *ValuationProviderFilterer) WatchValuationRequestReceived(opts *bind.WatchOpts, sink chan<- *ValuationProviderValuationRequestReceived) (event.Subscription, error) {

	logs, sub, err := _ValuationProvider.contract.WatchLogs(opts, "ValuationRequestReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValuationProviderValuationRequestReceived)
				if err := _ValuationProvider.contract.UnpackLog(event, "ValuationRequestReceived", log); err != nil {
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

// ParseValuationRequestReceived is a log parse operation binding the contract event 0x9112d470c4799e9a8541023cffabd2adbd83b00209706a1516542e741af417af.
//
// Solidity: event ValuationRequestReceived(address from)
func (_ValuationProvider *ValuationProviderFilterer) ParseValuationRequestReceived(log types.Log) (*ValuationProviderValuationRequestReceived, error) {
	event := new(ValuationProviderValuationRequestReceived)
	if err := _ValuationProvider.contract.UnpackLog(event, "ValuationRequestReceived", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValuationProviderValuationRequestSubmittedIterator is returned from FilterValuationRequestSubmitted and is used to iterate over the raw logs and unpacked data for ValuationRequestSubmitted events raised by the ValuationProvider contract.
type ValuationProviderValuationRequestSubmittedIterator struct {
	Event *ValuationProviderValuationRequestSubmitted // Event containing the contract specifics and raw log

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
func (it *ValuationProviderValuationRequestSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValuationProviderValuationRequestSubmitted)
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
		it.Event = new(ValuationProviderValuationRequestSubmitted)
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
func (it *ValuationProviderValuationRequestSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValuationProviderValuationRequestSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValuationProviderValuationRequestSubmitted represents a ValuationRequestSubmitted event raised by the ValuationProvider contract.
type ValuationProviderValuationRequestSubmitted struct {
	To        common.Address
	Valuation []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValuationRequestSubmitted is a free log retrieval operation binding the contract event 0x7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d.
//
// Solidity: event ValuationRequestSubmitted(address to, bytes valuation)
func (_ValuationProvider *ValuationProviderFilterer) FilterValuationRequestSubmitted(opts *bind.FilterOpts) (*ValuationProviderValuationRequestSubmittedIterator, error) {

	logs, sub, err := _ValuationProvider.contract.FilterLogs(opts, "ValuationRequestSubmitted")
	if err != nil {
		return nil, err
	}
	return &ValuationProviderValuationRequestSubmittedIterator{contract: _ValuationProvider.contract, event: "ValuationRequestSubmitted", logs: logs, sub: sub}, nil
}

// WatchValuationRequestSubmitted is a free log subscription operation binding the contract event 0x7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d.
//
// Solidity: event ValuationRequestSubmitted(address to, bytes valuation)
func (_ValuationProvider *ValuationProviderFilterer) WatchValuationRequestSubmitted(opts *bind.WatchOpts, sink chan<- *ValuationProviderValuationRequestSubmitted) (event.Subscription, error) {

	logs, sub, err := _ValuationProvider.contract.WatchLogs(opts, "ValuationRequestSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValuationProviderValuationRequestSubmitted)
				if err := _ValuationProvider.contract.UnpackLog(event, "ValuationRequestSubmitted", log); err != nil {
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

// ParseValuationRequestSubmitted is a log parse operation binding the contract event 0x7e33467a61041e8a9bfaede503107cae7a4fe823f782886635119ffb9f68f36d.
//
// Solidity: event ValuationRequestSubmitted(address to, bytes valuation)
func (_ValuationProvider *ValuationProviderFilterer) ParseValuationRequestSubmitted(log types.Log) (*ValuationProviderValuationRequestSubmitted, error) {
	event := new(ValuationProviderValuationRequestSubmitted)
	if err := _ValuationProvider.contract.UnpackLog(event, "ValuationRequestSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}
