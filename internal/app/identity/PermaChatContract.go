// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ChatModelABI is the input ABI used to generate the binding from.
const ChatModelABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"nextActionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextCommentIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPostIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"name\":\"targetId\",\"type\":\"uint256\"},{\"name\":\"targetType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewPostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"parentPostId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewReplyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"commenter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"NewCommentaryEvent\",\"type\":\"event\"}]"

// ChatModelBin is the compiled bytecode used for deploying new contracts.
const ChatModelBin = `0x`

// DeployChatModel deploys a new Ethereum contract, binding an instance of ChatModel to it.
func DeployChatModel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChatModel, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatModelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChatModelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChatModel{ChatModelCaller: ChatModelCaller{contract: contract}, ChatModelTransactor: ChatModelTransactor{contract: contract}, ChatModelFilterer: ChatModelFilterer{contract: contract}}, nil
}

// ChatModel is an auto generated Go binding around an Ethereum contract.
type ChatModel struct {
	ChatModelCaller     // Read-only binding to the contract
	ChatModelTransactor // Write-only binding to the contract
	ChatModelFilterer   // Log filterer for contract events
}

// ChatModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChatModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChatModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChatModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChatModelSession struct {
	Contract     *ChatModel        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChatModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChatModelCallerSession struct {
	Contract *ChatModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChatModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChatModelTransactorSession struct {
	Contract     *ChatModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChatModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChatModelRaw struct {
	Contract *ChatModel // Generic contract binding to access the raw methods on
}

// ChatModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChatModelCallerRaw struct {
	Contract *ChatModelCaller // Generic read-only contract binding to access the raw methods on
}

// ChatModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChatModelTransactorRaw struct {
	Contract *ChatModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChatModel creates a new instance of ChatModel, bound to a specific deployed contract.
func NewChatModel(address common.Address, backend bind.ContractBackend) (*ChatModel, error) {
	contract, err := bindChatModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChatModel{ChatModelCaller: ChatModelCaller{contract: contract}, ChatModelTransactor: ChatModelTransactor{contract: contract}, ChatModelFilterer: ChatModelFilterer{contract: contract}}, nil
}

// NewChatModelCaller creates a new read-only instance of ChatModel, bound to a specific deployed contract.
func NewChatModelCaller(address common.Address, caller bind.ContractCaller) (*ChatModelCaller, error) {
	contract, err := bindChatModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatModelCaller{contract: contract}, nil
}

// NewChatModelTransactor creates a new write-only instance of ChatModel, bound to a specific deployed contract.
func NewChatModelTransactor(address common.Address, transactor bind.ContractTransactor) (*ChatModelTransactor, error) {
	contract, err := bindChatModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChatModelTransactor{contract: contract}, nil
}

// NewChatModelFilterer creates a new log filterer instance of ChatModel, bound to a specific deployed contract.
func NewChatModelFilterer(address common.Address, filterer bind.ContractFilterer) (*ChatModelFilterer, error) {
	contract, err := bindChatModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChatModelFilterer{contract: contract}, nil
}

// bindChatModel binds a generic wrapper to an already deployed contract.
func bindChatModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatModel *ChatModelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatModel.Contract.ChatModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatModel *ChatModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatModel.Contract.ChatModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatModel *ChatModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatModel.Contract.ChatModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatModel *ChatModelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatModel *ChatModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatModel *ChatModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatModel.Contract.contract.Transact(opts, method, params...)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatModel *ChatModelCaller) GetAction(opts *bind.CallOpts, actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	ret := new(struct {
		TargetId   *big.Int
		TargetType uint8
	})
	out := ret
	err := _ChatModel.contract.Call(opts, out, "getAction", actionId)
	return *ret, err
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatModel *ChatModelSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _ChatModel.Contract.GetAction(&_ChatModel.CallOpts, actionId)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatModel *ChatModelCallerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _ChatModel.Contract.GetAction(&_ChatModel.CallOpts, actionId)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatModel *ChatModelCaller) NextActionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatModel.contract.Call(opts, out, "nextActionId")
	return *ret0, err
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatModel *ChatModelSession) NextActionId() (*big.Int, error) {
	return _ChatModel.Contract.NextActionId(&_ChatModel.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatModel *ChatModelCallerSession) NextActionId() (*big.Int, error) {
	return _ChatModel.Contract.NextActionId(&_ChatModel.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatModel *ChatModelCaller) NextCommentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatModel.contract.Call(opts, out, "nextCommentIndex")
	return *ret0, err
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatModel *ChatModelSession) NextCommentIndex() (*big.Int, error) {
	return _ChatModel.Contract.NextCommentIndex(&_ChatModel.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatModel *ChatModelCallerSession) NextCommentIndex() (*big.Int, error) {
	return _ChatModel.Contract.NextCommentIndex(&_ChatModel.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatModel *ChatModelCaller) NextPostIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatModel.contract.Call(opts, out, "nextPostIndex")
	return *ret0, err
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatModel *ChatModelSession) NextPostIndex() (*big.Int, error) {
	return _ChatModel.Contract.NextPostIndex(&_ChatModel.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatModel *ChatModelCallerSession) NextPostIndex() (*big.Int, error) {
	return _ChatModel.Contract.NextPostIndex(&_ChatModel.CallOpts)
}

// ChatModelNewCommentaryEventIterator is returned from FilterNewCommentaryEvent and is used to iterate over the raw logs and unpacked data for NewCommentaryEvent events raised by the ChatModel contract.
type ChatModelNewCommentaryEventIterator struct {
	Event *ChatModelNewCommentaryEvent // Event containing the contract specifics and raw log

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
func (it *ChatModelNewCommentaryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatModelNewCommentaryEvent)
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
		it.Event = new(ChatModelNewCommentaryEvent)
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
func (it *ChatModelNewCommentaryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatModelNewCommentaryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatModelNewCommentaryEvent represents a NewCommentaryEvent event raised by the ChatModel contract.
type ChatModelNewCommentaryEvent struct {
	PostId          *big.Int
	CommentaryIndex *big.Int
	CommentId       *big.Int
	CommentaryType  uint8
	Commenter       common.Address
	Timestamp       *big.Int
	Tip             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCommentaryEvent is a free log retrieval operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatModel *ChatModelFilterer) FilterNewCommentaryEvent(opts *bind.FilterOpts) (*ChatModelNewCommentaryEventIterator, error) {

	logs, sub, err := _ChatModel.contract.FilterLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return &ChatModelNewCommentaryEventIterator{contract: _ChatModel.contract, event: "NewCommentaryEvent", logs: logs, sub: sub}, nil
}

// WatchNewCommentaryEvent is a free log subscription operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatModel *ChatModelFilterer) WatchNewCommentaryEvent(opts *bind.WatchOpts, sink chan<- *ChatModelNewCommentaryEvent) (event.Subscription, error) {

	logs, sub, err := _ChatModel.contract.WatchLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatModelNewCommentaryEvent)
				if err := _ChatModel.contract.UnpackLog(event, "NewCommentaryEvent", log); err != nil {
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

// ChatModelNewPostEventIterator is returned from FilterNewPostEvent and is used to iterate over the raw logs and unpacked data for NewPostEvent events raised by the ChatModel contract.
type ChatModelNewPostEventIterator struct {
	Event *ChatModelNewPostEvent // Event containing the contract specifics and raw log

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
func (it *ChatModelNewPostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatModelNewPostEvent)
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
		it.Event = new(ChatModelNewPostEvent)
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
func (it *ChatModelNewPostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatModelNewPostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatModelNewPostEvent represents a NewPostEvent event raised by the ChatModel contract.
type ChatModelNewPostEvent struct {
	PostId      *big.Int
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPostEvent is a free log retrieval operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatModel *ChatModelFilterer) FilterNewPostEvent(opts *bind.FilterOpts) (*ChatModelNewPostEventIterator, error) {

	logs, sub, err := _ChatModel.contract.FilterLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return &ChatModelNewPostEventIterator{contract: _ChatModel.contract, event: "NewPostEvent", logs: logs, sub: sub}, nil
}

// WatchNewPostEvent is a free log subscription operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatModel *ChatModelFilterer) WatchNewPostEvent(opts *bind.WatchOpts, sink chan<- *ChatModelNewPostEvent) (event.Subscription, error) {

	logs, sub, err := _ChatModel.contract.WatchLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatModelNewPostEvent)
				if err := _ChatModel.contract.UnpackLog(event, "NewPostEvent", log); err != nil {
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

// ChatModelNewReplyEventIterator is returned from FilterNewReplyEvent and is used to iterate over the raw logs and unpacked data for NewReplyEvent events raised by the ChatModel contract.
type ChatModelNewReplyEventIterator struct {
	Event *ChatModelNewReplyEvent // Event containing the contract specifics and raw log

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
func (it *ChatModelNewReplyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatModelNewReplyEvent)
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
		it.Event = new(ChatModelNewReplyEvent)
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
func (it *ChatModelNewReplyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatModelNewReplyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatModelNewReplyEvent represents a NewReplyEvent event raised by the ChatModel contract.
type ChatModelNewReplyEvent struct {
	ParentPostId *big.Int
	PostId       *big.Int
	IpfsHash     string
	Poster       common.Address
	BlockNumber  *big.Int
	Timestamp    *big.Int
	ContentType  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewReplyEvent is a free log retrieval operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatModel *ChatModelFilterer) FilterNewReplyEvent(opts *bind.FilterOpts) (*ChatModelNewReplyEventIterator, error) {

	logs, sub, err := _ChatModel.contract.FilterLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return &ChatModelNewReplyEventIterator{contract: _ChatModel.contract, event: "NewReplyEvent", logs: logs, sub: sub}, nil
}

// WatchNewReplyEvent is a free log subscription operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatModel *ChatModelFilterer) WatchNewReplyEvent(opts *bind.WatchOpts, sink chan<- *ChatModelNewReplyEvent) (event.Subscription, error) {

	logs, sub, err := _ChatModel.contract.WatchLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatModelNewReplyEvent)
				if err := _ChatModel.contract.UnpackLog(event, "NewReplyEvent", log); err != nil {
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

// ChatServiceABI is the input ABI used to generate the binding from.
const ChatServiceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addDownvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"pinPost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextActionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextCommentIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identityAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPostIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"parentPostId\",\"type\":\"uint256\"},{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"newReply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"newPost\",\"outputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addFlag\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"name\":\"targetId\",\"type\":\"uint256\"},{\"name\":\"targetType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addUpvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityRequestId\",\"type\":\"uint256\"}],\"name\":\"NewRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NewApprovedProviderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewPostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"parentPostId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewReplyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"commenter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"NewCommentaryEvent\",\"type\":\"event\"}]"

// ChatServiceBin is the compiled bytecode used for deploying new contracts.
const ChatServiceBin = `0x60806040526001600081905560028190556009819055600d819055600e819055601055600b8054600160a060020a0319163317905561101a806100436000396000f3fe6080604052600436106100df577c0100000000000000000000000000000000000000000000000000000000600035046301c6090781146100e45780631604f9ea14610110578063251cf4ea1461013757806346abe73a146101615780634d625d991461017657806361165d081461018b57806375efbecf146101a05780637e889537146101b55780638e26c7e6146101e65780638ef5d2b9146102035780638faa6dad146102185780639e28cf9614610359578063ae2cded614610493578063b6e76873146104bd578063f1de9ea414610512578063fbd5b4e11461053c575b600080fd5b3480156100f057600080fd5b5061010e6004803603602081101561010757600080fd5b5035610551565b005b34801561011c57600080fd5b50610125610561565b60408051918252519081900360200190f35b34801561014357600080fd5b5061010e6004803603602081101561015a57600080fd5b5035610567565b34801561016d57600080fd5b506101256105da565b34801561018257600080fd5b506101256105e0565b34801561019757600080fd5b506101256105e6565b3480156101ac57600080fd5b506101256105ec565b3480156101c157600080fd5b506101ca6105f2565b60408051600160a060020a039092168252519081900360200190f35b61010e600480360360208110156101fc57600080fd5b5035610601565b34801561020f57600080fd5b50610125610659565b34801561022457600080fd5b506101256004803603606081101561023b57600080fd5b8135919081019060408101602082013564010000000081111561025d57600080fd5b82018360208201111561026f57600080fd5b8035906020019184600183028401116401000000008311171561029157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156102e457600080fd5b8201836020820111156102f657600080fd5b8035906020019184600183028401116401000000008311171561031857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061065f945050505050565b34801561036557600080fd5b506101256004803603604081101561037c57600080fd5b81019060208101813564010000000081111561039757600080fd5b8201836020820111156103a957600080fd5b803590602001918460018302840111640100000000831117156103cb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561041e57600080fd5b82018360208201111561043057600080fd5b8035906020019184600183028401116401000000008311171561045257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108e0945050505050565b34801561049f57600080fd5b5061010e600480360360208110156104b657600080fd5b5035610b56565b3480156104c957600080fd5b506104e7600480360360208110156104e057600080fd5b5035610b63565b604051808381526020018260018111156104fd57fe5b60ff1681526020019250505060405180910390f35b34801561051e57600080fd5b5061010e6004803603602081101561053557600080fd5b5035610b83565b34801561054857600080fd5b5061010e610b90565b61055e8160026000610be3565b50565b600c5481565b6002548190811061057757600080fd5b3360009081526012602090815260408083205483526011909152812060010154116105a157600080fd5b600082815260036020526040902060020154600160a060020a031633146105c757600080fd5b5033600090815260066020526040902055565b60005481565b600d5481565b60095481565b600e5481565b600b54600160a060020a031681565b61060d81600034610be3565b600081815260036020526040808220600201549051600160a060020a03909116913480156108fc02929091818181858888f19350505050158015610655573d6000803e3d6000fd5b5050565b60025481565b3360009081526012602090815260408083205483526011909152812060010154811061068a57600080fd5b6002548490811061069a57600080fd5b60006106a685856108e0565b9050600460008781526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190915055508560056000838152602001908152602001600020819055507f378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3866003600084815260200190815260200160002060000154600360008581526020019081526020016000206001016003600086815260200190815260200160002060020160009054906101000a9004600160a060020a0316600360008781526020019081526020016000206003015460036000888152602001908152602001600020600401548a604051808881526020018781526020018060200186600160a060020a0316600160a060020a03168152602001858152602001848152602001806020018381038352888181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156108625780601f1061083757610100808354040283529160200191610862565b820191906000526020600020905b81548152906001019060200180831161084557829003601f168201915b5050838103825284518152845160209182019186019080838360005b8381101561089657818101518382015260200161087e565b50505050905090810190601f1680156108c35780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a195945050505050565b3360009081526012602090815260408083205483526011909152812060010154811061090b57600080fd5b6002546040805160c081018252828152602080820187815233838501524360608401526103e84202608084015260a083018790526000858152600383529390932082518155925180519293926109679260018501920190610f21565b50604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055606082015160038201556080820151600482015560a082015180516109ca916005840191602090910190610f21565b5050600280546001019055506109e1816000610ea0565b60008181526003602081815260409283902080546002808301549483015460048401548751848152600160a060020a03909716978701889052606087018290526080870181905260c095870186815260019586018054600019978116156101000297909701909616939093049587018690527fab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac275497939694959394919390928b929160a083019060e084019089908015610adb5780601f10610ab057610100808354040283529160200191610adb565b820191906000526020600020905b815481529060010190602001808311610abe57829003601f168201915b5050838103825284518152845160209182019186019080838360005b83811015610b0f578181015183820152602001610af7565b50505050905090810190601f168015610b3c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a19392505050565b61055e8160036000610be3565b60009081526001602081905260409091208054910154909160ff90911690565b61055e8160016000610be3565b336000908152601260209081526040808320548352601190915281206001015411610bba57600080fd5b336000908152600660205260408120541115610be157336000908152600660205260408120555b565b336000908152601260209081526040808320548352601190915281206001015411610c0d57600080fd5b60025483908110610c1d57600080fd5b60008481526008602052604081208591610c3633610f06565b815260208101919091526040016000205460ff1615610c5457600080fd5b6009546000868152600760209081526040808320546008909252822090916103e8420291600191610c8433610f06565b81526020810191909152604001600020805460ff1916911515919091179055610cab610f9f565b608060405190810160405280896003811115610cc357fe5b815233602080830191909152604080830186905260609092018a905260008c81526007825291822080546001818101808455928552929093208451600394850290910180549596509194869492939192849260ff1990921691908490811115610d2857fe5b02179055506020828101518254600160a060020a039091166101000274ffffffffffffffffffffffffffffffffffffffff00199091161782556040808401516001808501919091556060909401516002909301929092556000888152600a909152208351815485945091929091839160ff1990911690836003811115610daa57fe5b021790555060208201518154600160a060020a039091166101000274ffffffffffffffffffffffffffffffffffffffff00199091161781556040820151600180830191909155606090920151600290910155610e07908590610ea0565b60016009600082825401925050819055507fa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc158984868b33878d60405180888152602001878152602001868152602001856003811115610e6257fe5b60ff168152600160a060020a0390941660208501525060408084019290925260608301525190819003608001945092505050a1505050505050505050565b6040805190810160405280838152602001826001811115610ebd57fe5b9052600080548152600160208181526040909220835181559183015182820180549192909160ff1916908381811115610ef257fe5b021790555050600080546001019055505050565b600160a060020a031660009081526012602052604090205490565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f6257805160ff1916838001178555610f8f565b82800160010185558215610f8f579182015b82811115610f8f578251825591602001919060010190610f74565b50610f9b929150610fd1565b5090565b604080516080810190915280600081526020016000600160a060020a0316815260200160008152602001600081525090565b610feb91905b80821115610f9b5760008155600101610fd7565b9056fea165627a7a72305820920634ef3809d9d094aeb85f4c128ef48d6cbc4a7570b74e87089ce8dfb4a6050029`

// DeployChatService deploys a new Ethereum contract, binding an instance of ChatService to it.
func DeployChatService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChatService, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChatServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChatService{ChatServiceCaller: ChatServiceCaller{contract: contract}, ChatServiceTransactor: ChatServiceTransactor{contract: contract}, ChatServiceFilterer: ChatServiceFilterer{contract: contract}}, nil
}

// ChatService is an auto generated Go binding around an Ethereum contract.
type ChatService struct {
	ChatServiceCaller     // Read-only binding to the contract
	ChatServiceTransactor // Write-only binding to the contract
	ChatServiceFilterer   // Log filterer for contract events
}

// ChatServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChatServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChatServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChatServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChatServiceSession struct {
	Contract     *ChatService      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChatServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChatServiceCallerSession struct {
	Contract *ChatServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChatServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChatServiceTransactorSession struct {
	Contract     *ChatServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChatServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChatServiceRaw struct {
	Contract *ChatService // Generic contract binding to access the raw methods on
}

// ChatServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChatServiceCallerRaw struct {
	Contract *ChatServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ChatServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChatServiceTransactorRaw struct {
	Contract *ChatServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChatService creates a new instance of ChatService, bound to a specific deployed contract.
func NewChatService(address common.Address, backend bind.ContractBackend) (*ChatService, error) {
	contract, err := bindChatService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChatService{ChatServiceCaller: ChatServiceCaller{contract: contract}, ChatServiceTransactor: ChatServiceTransactor{contract: contract}, ChatServiceFilterer: ChatServiceFilterer{contract: contract}}, nil
}

// NewChatServiceCaller creates a new read-only instance of ChatService, bound to a specific deployed contract.
func NewChatServiceCaller(address common.Address, caller bind.ContractCaller) (*ChatServiceCaller, error) {
	contract, err := bindChatService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatServiceCaller{contract: contract}, nil
}

// NewChatServiceTransactor creates a new write-only instance of ChatService, bound to a specific deployed contract.
func NewChatServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ChatServiceTransactor, error) {
	contract, err := bindChatService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChatServiceTransactor{contract: contract}, nil
}

// NewChatServiceFilterer creates a new log filterer instance of ChatService, bound to a specific deployed contract.
func NewChatServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ChatServiceFilterer, error) {
	contract, err := bindChatService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChatServiceFilterer{contract: contract}, nil
}

// bindChatService binds a generic wrapper to an already deployed contract.
func bindChatService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatService *ChatServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatService.Contract.ChatServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatService *ChatServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatService.Contract.ChatServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatService *ChatServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatService.Contract.ChatServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatService *ChatServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatService *ChatServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatService *ChatServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatService.Contract.contract.Transact(opts, method, params...)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatService *ChatServiceCaller) GetAction(opts *bind.CallOpts, actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	ret := new(struct {
		TargetId   *big.Int
		TargetType uint8
	})
	out := ret
	err := _ChatService.contract.Call(opts, out, "getAction", actionId)
	return *ret, err
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatService *ChatServiceSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _ChatService.Contract.GetAction(&_ChatService.CallOpts, actionId)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatService *ChatServiceCallerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _ChatService.Contract.GetAction(&_ChatService.CallOpts, actionId)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_ChatService *ChatServiceCaller) IdentityAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "identityAdmin")
	return *ret0, err
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_ChatService *ChatServiceSession) IdentityAdmin() (common.Address, error) {
	return _ChatService.Contract.IdentityAdmin(&_ChatService.CallOpts)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_ChatService *ChatServiceCallerSession) IdentityAdmin() (common.Address, error) {
	return _ChatService.Contract.IdentityAdmin(&_ChatService.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatService *ChatServiceCaller) NextActionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "nextActionId")
	return *ret0, err
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatService *ChatServiceSession) NextActionId() (*big.Int, error) {
	return _ChatService.Contract.NextActionId(&_ChatService.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatService *ChatServiceCallerSession) NextActionId() (*big.Int, error) {
	return _ChatService.Contract.NextActionId(&_ChatService.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatService *ChatServiceCaller) NextCommentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "nextCommentIndex")
	return *ret0, err
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatService *ChatServiceSession) NextCommentIndex() (*big.Int, error) {
	return _ChatService.Contract.NextCommentIndex(&_ChatService.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatService *ChatServiceCallerSession) NextCommentIndex() (*big.Int, error) {
	return _ChatService.Contract.NextCommentIndex(&_ChatService.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_ChatService *ChatServiceCaller) NextIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "nextIdentityRequestIndex")
	return *ret0, err
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_ChatService *ChatServiceSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _ChatService.Contract.NextIdentityRequestIndex(&_ChatService.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_ChatService *ChatServiceCallerSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _ChatService.Contract.NextIdentityRequestIndex(&_ChatService.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatService *ChatServiceCaller) NextPostIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "nextPostIndex")
	return *ret0, err
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatService *ChatServiceSession) NextPostIndex() (*big.Int, error) {
	return _ChatService.Contract.NextPostIndex(&_ChatService.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatService *ChatServiceCallerSession) NextPostIndex() (*big.Int, error) {
	return _ChatService.Contract.NextPostIndex(&_ChatService.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_ChatService *ChatServiceCaller) PendingIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "pendingIdentityRequestIndex")
	return *ret0, err
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_ChatService *ChatServiceSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _ChatService.Contract.PendingIdentityRequestIndex(&_ChatService.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_ChatService *ChatServiceCallerSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _ChatService.Contract.PendingIdentityRequestIndex(&_ChatService.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_ChatService *ChatServiceCaller) RequestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatService.contract.Call(opts, out, "requestPrice")
	return *ret0, err
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_ChatService *ChatServiceSession) RequestPrice() (*big.Int, error) {
	return _ChatService.Contract.RequestPrice(&_ChatService.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_ChatService *ChatServiceCallerSession) RequestPrice() (*big.Int, error) {
	return _ChatService.Contract.RequestPrice(&_ChatService.CallOpts)
}

// AddDownvote is a paid mutator transaction binding the contract method 0x01c60907.
//
// Solidity: function addDownvote(postId uint256) returns()
func (_ChatService *ChatServiceTransactor) AddDownvote(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "addDownvote", postId)
}

// AddDownvote is a paid mutator transaction binding the contract method 0x01c60907.
//
// Solidity: function addDownvote(postId uint256) returns()
func (_ChatService *ChatServiceSession) AddDownvote(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddDownvote(&_ChatService.TransactOpts, postId)
}

// AddDownvote is a paid mutator transaction binding the contract method 0x01c60907.
//
// Solidity: function addDownvote(postId uint256) returns()
func (_ChatService *ChatServiceTransactorSession) AddDownvote(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddDownvote(&_ChatService.TransactOpts, postId)
}

// AddFlag is a paid mutator transaction binding the contract method 0xae2cded6.
//
// Solidity: function addFlag(postId uint256) returns()
func (_ChatService *ChatServiceTransactor) AddFlag(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "addFlag", postId)
}

// AddFlag is a paid mutator transaction binding the contract method 0xae2cded6.
//
// Solidity: function addFlag(postId uint256) returns()
func (_ChatService *ChatServiceSession) AddFlag(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddFlag(&_ChatService.TransactOpts, postId)
}

// AddFlag is a paid mutator transaction binding the contract method 0xae2cded6.
//
// Solidity: function addFlag(postId uint256) returns()
func (_ChatService *ChatServiceTransactorSession) AddFlag(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddFlag(&_ChatService.TransactOpts, postId)
}

// AddTip is a paid mutator transaction binding the contract method 0x8e26c7e6.
//
// Solidity: function addTip(postId uint256) returns()
func (_ChatService *ChatServiceTransactor) AddTip(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "addTip", postId)
}

// AddTip is a paid mutator transaction binding the contract method 0x8e26c7e6.
//
// Solidity: function addTip(postId uint256) returns()
func (_ChatService *ChatServiceSession) AddTip(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddTip(&_ChatService.TransactOpts, postId)
}

// AddTip is a paid mutator transaction binding the contract method 0x8e26c7e6.
//
// Solidity: function addTip(postId uint256) returns()
func (_ChatService *ChatServiceTransactorSession) AddTip(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddTip(&_ChatService.TransactOpts, postId)
}

// AddUpvote is a paid mutator transaction binding the contract method 0xf1de9ea4.
//
// Solidity: function addUpvote(postId uint256) returns()
func (_ChatService *ChatServiceTransactor) AddUpvote(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "addUpvote", postId)
}

// AddUpvote is a paid mutator transaction binding the contract method 0xf1de9ea4.
//
// Solidity: function addUpvote(postId uint256) returns()
func (_ChatService *ChatServiceSession) AddUpvote(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddUpvote(&_ChatService.TransactOpts, postId)
}

// AddUpvote is a paid mutator transaction binding the contract method 0xf1de9ea4.
//
// Solidity: function addUpvote(postId uint256) returns()
func (_ChatService *ChatServiceTransactorSession) AddUpvote(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.AddUpvote(&_ChatService.TransactOpts, postId)
}

// NewPost is a paid mutator transaction binding the contract method 0x9e28cf96.
//
// Solidity: function newPost(ipfsHash string, contentType string) returns(postId uint256)
func (_ChatService *ChatServiceTransactor) NewPost(opts *bind.TransactOpts, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "newPost", ipfsHash, contentType)
}

// NewPost is a paid mutator transaction binding the contract method 0x9e28cf96.
//
// Solidity: function newPost(ipfsHash string, contentType string) returns(postId uint256)
func (_ChatService *ChatServiceSession) NewPost(ipfsHash string, contentType string) (*types.Transaction, error) {
	return _ChatService.Contract.NewPost(&_ChatService.TransactOpts, ipfsHash, contentType)
}

// NewPost is a paid mutator transaction binding the contract method 0x9e28cf96.
//
// Solidity: function newPost(ipfsHash string, contentType string) returns(postId uint256)
func (_ChatService *ChatServiceTransactorSession) NewPost(ipfsHash string, contentType string) (*types.Transaction, error) {
	return _ChatService.Contract.NewPost(&_ChatService.TransactOpts, ipfsHash, contentType)
}

// NewReply is a paid mutator transaction binding the contract method 0x8faa6dad.
//
// Solidity: function newReply(parentPostId uint256, ipfsHash string, contentType string) returns(uint256)
func (_ChatService *ChatServiceTransactor) NewReply(opts *bind.TransactOpts, parentPostId *big.Int, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "newReply", parentPostId, ipfsHash, contentType)
}

// NewReply is a paid mutator transaction binding the contract method 0x8faa6dad.
//
// Solidity: function newReply(parentPostId uint256, ipfsHash string, contentType string) returns(uint256)
func (_ChatService *ChatServiceSession) NewReply(parentPostId *big.Int, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _ChatService.Contract.NewReply(&_ChatService.TransactOpts, parentPostId, ipfsHash, contentType)
}

// NewReply is a paid mutator transaction binding the contract method 0x8faa6dad.
//
// Solidity: function newReply(parentPostId uint256, ipfsHash string, contentType string) returns(uint256)
func (_ChatService *ChatServiceTransactorSession) NewReply(parentPostId *big.Int, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _ChatService.Contract.NewReply(&_ChatService.TransactOpts, parentPostId, ipfsHash, contentType)
}

// PinPost is a paid mutator transaction binding the contract method 0x251cf4ea.
//
// Solidity: function pinPost(postId uint256) returns()
func (_ChatService *ChatServiceTransactor) PinPost(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "pinPost", postId)
}

// PinPost is a paid mutator transaction binding the contract method 0x251cf4ea.
//
// Solidity: function pinPost(postId uint256) returns()
func (_ChatService *ChatServiceSession) PinPost(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.PinPost(&_ChatService.TransactOpts, postId)
}

// PinPost is a paid mutator transaction binding the contract method 0x251cf4ea.
//
// Solidity: function pinPost(postId uint256) returns()
func (_ChatService *ChatServiceTransactorSession) PinPost(postId *big.Int) (*types.Transaction, error) {
	return _ChatService.Contract.PinPost(&_ChatService.TransactOpts, postId)
}

// Unpin is a paid mutator transaction binding the contract method 0xfbd5b4e1.
//
// Solidity: function unpin() returns()
func (_ChatService *ChatServiceTransactor) Unpin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatService.contract.Transact(opts, "unpin")
}

// Unpin is a paid mutator transaction binding the contract method 0xfbd5b4e1.
//
// Solidity: function unpin() returns()
func (_ChatService *ChatServiceSession) Unpin() (*types.Transaction, error) {
	return _ChatService.Contract.Unpin(&_ChatService.TransactOpts)
}

// Unpin is a paid mutator transaction binding the contract method 0xfbd5b4e1.
//
// Solidity: function unpin() returns()
func (_ChatService *ChatServiceTransactorSession) Unpin() (*types.Transaction, error) {
	return _ChatService.Contract.Unpin(&_ChatService.TransactOpts)
}

// ChatServiceNewApprovedProviderEventIterator is returned from FilterNewApprovedProviderEvent and is used to iterate over the raw logs and unpacked data for NewApprovedProviderEvent events raised by the ChatService contract.
type ChatServiceNewApprovedProviderEventIterator struct {
	Event *ChatServiceNewApprovedProviderEvent // Event containing the contract specifics and raw log

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
func (it *ChatServiceNewApprovedProviderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatServiceNewApprovedProviderEvent)
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
		it.Event = new(ChatServiceNewApprovedProviderEvent)
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
func (it *ChatServiceNewApprovedProviderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatServiceNewApprovedProviderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatServiceNewApprovedProviderEvent represents a NewApprovedProviderEvent event raised by the ChatService contract.
type ChatServiceNewApprovedProviderEvent struct {
	IdentityId *big.Int
	ProviderId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewApprovedProviderEvent is a free log retrieval operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_ChatService *ChatServiceFilterer) FilterNewApprovedProviderEvent(opts *bind.FilterOpts) (*ChatServiceNewApprovedProviderEventIterator, error) {

	logs, sub, err := _ChatService.contract.FilterLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return &ChatServiceNewApprovedProviderEventIterator{contract: _ChatService.contract, event: "NewApprovedProviderEvent", logs: logs, sub: sub}, nil
}

// WatchNewApprovedProviderEvent is a free log subscription operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_ChatService *ChatServiceFilterer) WatchNewApprovedProviderEvent(opts *bind.WatchOpts, sink chan<- *ChatServiceNewApprovedProviderEvent) (event.Subscription, error) {

	logs, sub, err := _ChatService.contract.WatchLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatServiceNewApprovedProviderEvent)
				if err := _ChatService.contract.UnpackLog(event, "NewApprovedProviderEvent", log); err != nil {
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

// ChatServiceNewCommentaryEventIterator is returned from FilterNewCommentaryEvent and is used to iterate over the raw logs and unpacked data for NewCommentaryEvent events raised by the ChatService contract.
type ChatServiceNewCommentaryEventIterator struct {
	Event *ChatServiceNewCommentaryEvent // Event containing the contract specifics and raw log

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
func (it *ChatServiceNewCommentaryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatServiceNewCommentaryEvent)
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
		it.Event = new(ChatServiceNewCommentaryEvent)
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
func (it *ChatServiceNewCommentaryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatServiceNewCommentaryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatServiceNewCommentaryEvent represents a NewCommentaryEvent event raised by the ChatService contract.
type ChatServiceNewCommentaryEvent struct {
	PostId          *big.Int
	CommentaryIndex *big.Int
	CommentId       *big.Int
	CommentaryType  uint8
	Commenter       common.Address
	Timestamp       *big.Int
	Tip             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCommentaryEvent is a free log retrieval operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatService *ChatServiceFilterer) FilterNewCommentaryEvent(opts *bind.FilterOpts) (*ChatServiceNewCommentaryEventIterator, error) {

	logs, sub, err := _ChatService.contract.FilterLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return &ChatServiceNewCommentaryEventIterator{contract: _ChatService.contract, event: "NewCommentaryEvent", logs: logs, sub: sub}, nil
}

// WatchNewCommentaryEvent is a free log subscription operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatService *ChatServiceFilterer) WatchNewCommentaryEvent(opts *bind.WatchOpts, sink chan<- *ChatServiceNewCommentaryEvent) (event.Subscription, error) {

	logs, sub, err := _ChatService.contract.WatchLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatServiceNewCommentaryEvent)
				if err := _ChatService.contract.UnpackLog(event, "NewCommentaryEvent", log); err != nil {
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

// ChatServiceNewPostEventIterator is returned from FilterNewPostEvent and is used to iterate over the raw logs and unpacked data for NewPostEvent events raised by the ChatService contract.
type ChatServiceNewPostEventIterator struct {
	Event *ChatServiceNewPostEvent // Event containing the contract specifics and raw log

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
func (it *ChatServiceNewPostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatServiceNewPostEvent)
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
		it.Event = new(ChatServiceNewPostEvent)
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
func (it *ChatServiceNewPostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatServiceNewPostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatServiceNewPostEvent represents a NewPostEvent event raised by the ChatService contract.
type ChatServiceNewPostEvent struct {
	PostId      *big.Int
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPostEvent is a free log retrieval operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatService *ChatServiceFilterer) FilterNewPostEvent(opts *bind.FilterOpts) (*ChatServiceNewPostEventIterator, error) {

	logs, sub, err := _ChatService.contract.FilterLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return &ChatServiceNewPostEventIterator{contract: _ChatService.contract, event: "NewPostEvent", logs: logs, sub: sub}, nil
}

// WatchNewPostEvent is a free log subscription operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatService *ChatServiceFilterer) WatchNewPostEvent(opts *bind.WatchOpts, sink chan<- *ChatServiceNewPostEvent) (event.Subscription, error) {

	logs, sub, err := _ChatService.contract.WatchLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatServiceNewPostEvent)
				if err := _ChatService.contract.UnpackLog(event, "NewPostEvent", log); err != nil {
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

// ChatServiceNewReplyEventIterator is returned from FilterNewReplyEvent and is used to iterate over the raw logs and unpacked data for NewReplyEvent events raised by the ChatService contract.
type ChatServiceNewReplyEventIterator struct {
	Event *ChatServiceNewReplyEvent // Event containing the contract specifics and raw log

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
func (it *ChatServiceNewReplyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatServiceNewReplyEvent)
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
		it.Event = new(ChatServiceNewReplyEvent)
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
func (it *ChatServiceNewReplyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatServiceNewReplyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatServiceNewReplyEvent represents a NewReplyEvent event raised by the ChatService contract.
type ChatServiceNewReplyEvent struct {
	ParentPostId *big.Int
	PostId       *big.Int
	IpfsHash     string
	Poster       common.Address
	BlockNumber  *big.Int
	Timestamp    *big.Int
	ContentType  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewReplyEvent is a free log retrieval operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatService *ChatServiceFilterer) FilterNewReplyEvent(opts *bind.FilterOpts) (*ChatServiceNewReplyEventIterator, error) {

	logs, sub, err := _ChatService.contract.FilterLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return &ChatServiceNewReplyEventIterator{contract: _ChatService.contract, event: "NewReplyEvent", logs: logs, sub: sub}, nil
}

// WatchNewReplyEvent is a free log subscription operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatService *ChatServiceFilterer) WatchNewReplyEvent(opts *bind.WatchOpts, sink chan<- *ChatServiceNewReplyEvent) (event.Subscription, error) {

	logs, sub, err := _ChatService.contract.WatchLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatServiceNewReplyEvent)
				if err := _ChatService.contract.UnpackLog(event, "NewReplyEvent", log); err != nil {
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

// ChatServiceNewRequestEventIterator is returned from FilterNewRequestEvent and is used to iterate over the raw logs and unpacked data for NewRequestEvent events raised by the ChatService contract.
type ChatServiceNewRequestEventIterator struct {
	Event *ChatServiceNewRequestEvent // Event containing the contract specifics and raw log

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
func (it *ChatServiceNewRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatServiceNewRequestEvent)
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
		it.Event = new(ChatServiceNewRequestEvent)
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
func (it *ChatServiceNewRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatServiceNewRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatServiceNewRequestEvent represents a NewRequestEvent event raised by the ChatService contract.
type ChatServiceNewRequestEvent struct {
	IdentityRequestId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRequestEvent is a free log retrieval operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_ChatService *ChatServiceFilterer) FilterNewRequestEvent(opts *bind.FilterOpts) (*ChatServiceNewRequestEventIterator, error) {

	logs, sub, err := _ChatService.contract.FilterLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return &ChatServiceNewRequestEventIterator{contract: _ChatService.contract, event: "NewRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewRequestEvent is a free log subscription operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_ChatService *ChatServiceFilterer) WatchNewRequestEvent(opts *bind.WatchOpts, sink chan<- *ChatServiceNewRequestEvent) (event.Subscription, error) {

	logs, sub, err := _ChatService.contract.WatchLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatServiceNewRequestEvent)
				if err := _ChatService.contract.UnpackLog(event, "NewRequestEvent", log); err != nil {
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

// ChatViewsABI is the input ABI used to generate the binding from.
const ChatViewsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"replyId\",\"type\":\"uint256\"}],\"name\":\"getReplyParent\",\"outputs\":[{\"name\":\"parentPostId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"getPost\",\"outputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"poster\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"contentType\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextActionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextCommentIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"hasCommented\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"},{\"name\":\"commentaryIndex\",\"type\":\"uint256\"}],\"name\":\"getCommentary\",\"outputs\":[{\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"name\":\"commenter\",\"type\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPostIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"commentId\",\"type\":\"uint256\"}],\"name\":\"getCommentary\",\"outputs\":[{\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"name\":\"commenter\",\"type\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPinnedPost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"getReplies\",\"outputs\":[{\"name\":\"replyIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"getCommentaryCount\",\"outputs\":[{\"name\":\"commentaryCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"name\":\"targetId\",\"type\":\"uint256\"},{\"name\":\"targetType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewPostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"parentPostId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewReplyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"commenter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"NewCommentaryEvent\",\"type\":\"event\"}]"

// ChatViewsBin is the compiled bytecode used for deploying new contracts.
const ChatViewsBin = `0x`

// DeployChatViews deploys a new Ethereum contract, binding an instance of ChatViews to it.
func DeployChatViews(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChatViews, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatViewsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChatViewsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChatViews{ChatViewsCaller: ChatViewsCaller{contract: contract}, ChatViewsTransactor: ChatViewsTransactor{contract: contract}, ChatViewsFilterer: ChatViewsFilterer{contract: contract}}, nil
}

// ChatViews is an auto generated Go binding around an Ethereum contract.
type ChatViews struct {
	ChatViewsCaller     // Read-only binding to the contract
	ChatViewsTransactor // Write-only binding to the contract
	ChatViewsFilterer   // Log filterer for contract events
}

// ChatViewsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChatViewsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatViewsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChatViewsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatViewsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChatViewsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatViewsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChatViewsSession struct {
	Contract     *ChatViews        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChatViewsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChatViewsCallerSession struct {
	Contract *ChatViewsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChatViewsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChatViewsTransactorSession struct {
	Contract     *ChatViewsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChatViewsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChatViewsRaw struct {
	Contract *ChatViews // Generic contract binding to access the raw methods on
}

// ChatViewsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChatViewsCallerRaw struct {
	Contract *ChatViewsCaller // Generic read-only contract binding to access the raw methods on
}

// ChatViewsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChatViewsTransactorRaw struct {
	Contract *ChatViewsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChatViews creates a new instance of ChatViews, bound to a specific deployed contract.
func NewChatViews(address common.Address, backend bind.ContractBackend) (*ChatViews, error) {
	contract, err := bindChatViews(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChatViews{ChatViewsCaller: ChatViewsCaller{contract: contract}, ChatViewsTransactor: ChatViewsTransactor{contract: contract}, ChatViewsFilterer: ChatViewsFilterer{contract: contract}}, nil
}

// NewChatViewsCaller creates a new read-only instance of ChatViews, bound to a specific deployed contract.
func NewChatViewsCaller(address common.Address, caller bind.ContractCaller) (*ChatViewsCaller, error) {
	contract, err := bindChatViews(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatViewsCaller{contract: contract}, nil
}

// NewChatViewsTransactor creates a new write-only instance of ChatViews, bound to a specific deployed contract.
func NewChatViewsTransactor(address common.Address, transactor bind.ContractTransactor) (*ChatViewsTransactor, error) {
	contract, err := bindChatViews(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChatViewsTransactor{contract: contract}, nil
}

// NewChatViewsFilterer creates a new log filterer instance of ChatViews, bound to a specific deployed contract.
func NewChatViewsFilterer(address common.Address, filterer bind.ContractFilterer) (*ChatViewsFilterer, error) {
	contract, err := bindChatViews(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChatViewsFilterer{contract: contract}, nil
}

// bindChatViews binds a generic wrapper to an already deployed contract.
func bindChatViews(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatViewsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatViews *ChatViewsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatViews.Contract.ChatViewsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatViews *ChatViewsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatViews.Contract.ChatViewsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatViews *ChatViewsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatViews.Contract.ChatViewsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatViews *ChatViewsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatViews.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatViews *ChatViewsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatViews.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatViews *ChatViewsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatViews.Contract.contract.Transact(opts, method, params...)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatViews *ChatViewsCaller) GetAction(opts *bind.CallOpts, actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	ret := new(struct {
		TargetId   *big.Int
		TargetType uint8
	})
	out := ret
	err := _ChatViews.contract.Call(opts, out, "getAction", actionId)
	return *ret, err
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatViews *ChatViewsSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _ChatViews.Contract.GetAction(&_ChatViews.CallOpts, actionId)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_ChatViews *ChatViewsCallerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _ChatViews.Contract.GetAction(&_ChatViews.CallOpts, actionId)
}

// GetCommentary is a free data retrieval call binding the contract method 0x980d9ada.
//
// Solidity: function getCommentary(commentId uint256) constant returns(commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatViews *ChatViewsCaller) GetCommentary(opts *bind.CallOpts, commentId *big.Int) (struct {
	CommentaryType uint8
	Commenter      common.Address
	Timestamp      *big.Int
	Tip            *big.Int
}, error) {
	ret := new(struct {
		CommentaryType uint8
		Commenter      common.Address
		Timestamp      *big.Int
		Tip            *big.Int
	})
	out := ret
	err := _ChatViews.contract.Call(opts, out, "getCommentary", commentId)
	return *ret, err
}

// GetCommentary is a free data retrieval call binding the contract method 0x980d9ada.
//
// Solidity: function getCommentary(commentId uint256) constant returns(commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatViews *ChatViewsSession) GetCommentary(commentId *big.Int) (struct {
	CommentaryType uint8
	Commenter      common.Address
	Timestamp      *big.Int
	Tip            *big.Int
}, error) {
	return _ChatViews.Contract.GetCommentary(&_ChatViews.CallOpts, commentId)
}

// GetCommentary is a free data retrieval call binding the contract method 0x980d9ada.
//
// Solidity: function getCommentary(commentId uint256) constant returns(commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatViews *ChatViewsCallerSession) GetCommentary(commentId *big.Int) (struct {
	CommentaryType uint8
	Commenter      common.Address
	Timestamp      *big.Int
	Tip            *big.Int
}, error) {
	return _ChatViews.Contract.GetCommentary(&_ChatViews.CallOpts, commentId)
}

// GetCommentaryCount is a free data retrieval call binding the contract method 0xb65167b2.
//
// Solidity: function getCommentaryCount(postId uint256) constant returns(commentaryCount uint256)
func (_ChatViews *ChatViewsCaller) GetCommentaryCount(opts *bind.CallOpts, postId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "getCommentaryCount", postId)
	return *ret0, err
}

// GetCommentaryCount is a free data retrieval call binding the contract method 0xb65167b2.
//
// Solidity: function getCommentaryCount(postId uint256) constant returns(commentaryCount uint256)
func (_ChatViews *ChatViewsSession) GetCommentaryCount(postId *big.Int) (*big.Int, error) {
	return _ChatViews.Contract.GetCommentaryCount(&_ChatViews.CallOpts, postId)
}

// GetCommentaryCount is a free data retrieval call binding the contract method 0xb65167b2.
//
// Solidity: function getCommentaryCount(postId uint256) constant returns(commentaryCount uint256)
func (_ChatViews *ChatViewsCallerSession) GetCommentaryCount(postId *big.Int) (*big.Int, error) {
	return _ChatViews.Contract.GetCommentaryCount(&_ChatViews.CallOpts, postId)
}

// GetPinnedPost is a free data retrieval call binding the contract method 0xa46325df.
//
// Solidity: function getPinnedPost() constant returns(uint256)
func (_ChatViews *ChatViewsCaller) GetPinnedPost(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "getPinnedPost")
	return *ret0, err
}

// GetPinnedPost is a free data retrieval call binding the contract method 0xa46325df.
//
// Solidity: function getPinnedPost() constant returns(uint256)
func (_ChatViews *ChatViewsSession) GetPinnedPost() (*big.Int, error) {
	return _ChatViews.Contract.GetPinnedPost(&_ChatViews.CallOpts)
}

// GetPinnedPost is a free data retrieval call binding the contract method 0xa46325df.
//
// Solidity: function getPinnedPost() constant returns(uint256)
func (_ChatViews *ChatViewsCallerSession) GetPinnedPost() (*big.Int, error) {
	return _ChatViews.Contract.GetPinnedPost(&_ChatViews.CallOpts)
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(postId uint256) constant returns(ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsCaller) GetPost(opts *bind.CallOpts, postId *big.Int) (struct {
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
}, error) {
	ret := new(struct {
		IpfsHash    string
		Poster      common.Address
		BlockNumber *big.Int
		Timestamp   *big.Int
		ContentType string
	})
	out := ret
	err := _ChatViews.contract.Call(opts, out, "getPost", postId)
	return *ret, err
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(postId uint256) constant returns(ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsSession) GetPost(postId *big.Int) (struct {
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
}, error) {
	return _ChatViews.Contract.GetPost(&_ChatViews.CallOpts, postId)
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(postId uint256) constant returns(ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsCallerSession) GetPost(postId *big.Int) (struct {
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
}, error) {
	return _ChatViews.Contract.GetPost(&_ChatViews.CallOpts, postId)
}

// GetReplies is a free data retrieval call binding the contract method 0xae58449d.
//
// Solidity: function getReplies(postId uint256) constant returns(replyIds uint256[])
func (_ChatViews *ChatViewsCaller) GetReplies(opts *bind.CallOpts, postId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "getReplies", postId)
	return *ret0, err
}

// GetReplies is a free data retrieval call binding the contract method 0xae58449d.
//
// Solidity: function getReplies(postId uint256) constant returns(replyIds uint256[])
func (_ChatViews *ChatViewsSession) GetReplies(postId *big.Int) ([]*big.Int, error) {
	return _ChatViews.Contract.GetReplies(&_ChatViews.CallOpts, postId)
}

// GetReplies is a free data retrieval call binding the contract method 0xae58449d.
//
// Solidity: function getReplies(postId uint256) constant returns(replyIds uint256[])
func (_ChatViews *ChatViewsCallerSession) GetReplies(postId *big.Int) ([]*big.Int, error) {
	return _ChatViews.Contract.GetReplies(&_ChatViews.CallOpts, postId)
}

// GetReplyParent is a free data retrieval call binding the contract method 0x2637adbf.
//
// Solidity: function getReplyParent(replyId uint256) constant returns(parentPostId uint256)
func (_ChatViews *ChatViewsCaller) GetReplyParent(opts *bind.CallOpts, replyId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "getReplyParent", replyId)
	return *ret0, err
}

// GetReplyParent is a free data retrieval call binding the contract method 0x2637adbf.
//
// Solidity: function getReplyParent(replyId uint256) constant returns(parentPostId uint256)
func (_ChatViews *ChatViewsSession) GetReplyParent(replyId *big.Int) (*big.Int, error) {
	return _ChatViews.Contract.GetReplyParent(&_ChatViews.CallOpts, replyId)
}

// GetReplyParent is a free data retrieval call binding the contract method 0x2637adbf.
//
// Solidity: function getReplyParent(replyId uint256) constant returns(parentPostId uint256)
func (_ChatViews *ChatViewsCallerSession) GetReplyParent(replyId *big.Int) (*big.Int, error) {
	return _ChatViews.Contract.GetReplyParent(&_ChatViews.CallOpts, replyId)
}

// HasCommented is a free data retrieval call binding the contract method 0x7af2d1af.
//
// Solidity: function hasCommented(postId uint256) constant returns(bool)
func (_ChatViews *ChatViewsCaller) HasCommented(opts *bind.CallOpts, postId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "hasCommented", postId)
	return *ret0, err
}

// HasCommented is a free data retrieval call binding the contract method 0x7af2d1af.
//
// Solidity: function hasCommented(postId uint256) constant returns(bool)
func (_ChatViews *ChatViewsSession) HasCommented(postId *big.Int) (bool, error) {
	return _ChatViews.Contract.HasCommented(&_ChatViews.CallOpts, postId)
}

// HasCommented is a free data retrieval call binding the contract method 0x7af2d1af.
//
// Solidity: function hasCommented(postId uint256) constant returns(bool)
func (_ChatViews *ChatViewsCallerSession) HasCommented(postId *big.Int) (bool, error) {
	return _ChatViews.Contract.HasCommented(&_ChatViews.CallOpts, postId)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatViews *ChatViewsCaller) NextActionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "nextActionId")
	return *ret0, err
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatViews *ChatViewsSession) NextActionId() (*big.Int, error) {
	return _ChatViews.Contract.NextActionId(&_ChatViews.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_ChatViews *ChatViewsCallerSession) NextActionId() (*big.Int, error) {
	return _ChatViews.Contract.NextActionId(&_ChatViews.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatViews *ChatViewsCaller) NextCommentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "nextCommentIndex")
	return *ret0, err
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatViews *ChatViewsSession) NextCommentIndex() (*big.Int, error) {
	return _ChatViews.Contract.NextCommentIndex(&_ChatViews.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_ChatViews *ChatViewsCallerSession) NextCommentIndex() (*big.Int, error) {
	return _ChatViews.Contract.NextCommentIndex(&_ChatViews.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatViews *ChatViewsCaller) NextPostIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChatViews.contract.Call(opts, out, "nextPostIndex")
	return *ret0, err
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatViews *ChatViewsSession) NextPostIndex() (*big.Int, error) {
	return _ChatViews.Contract.NextPostIndex(&_ChatViews.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_ChatViews *ChatViewsCallerSession) NextPostIndex() (*big.Int, error) {
	return _ChatViews.Contract.NextPostIndex(&_ChatViews.CallOpts)
}

// ChatViewsNewCommentaryEventIterator is returned from FilterNewCommentaryEvent and is used to iterate over the raw logs and unpacked data for NewCommentaryEvent events raised by the ChatViews contract.
type ChatViewsNewCommentaryEventIterator struct {
	Event *ChatViewsNewCommentaryEvent // Event containing the contract specifics and raw log

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
func (it *ChatViewsNewCommentaryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatViewsNewCommentaryEvent)
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
		it.Event = new(ChatViewsNewCommentaryEvent)
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
func (it *ChatViewsNewCommentaryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatViewsNewCommentaryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatViewsNewCommentaryEvent represents a NewCommentaryEvent event raised by the ChatViews contract.
type ChatViewsNewCommentaryEvent struct {
	PostId          *big.Int
	CommentaryIndex *big.Int
	CommentId       *big.Int
	CommentaryType  uint8
	Commenter       common.Address
	Timestamp       *big.Int
	Tip             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCommentaryEvent is a free log retrieval operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatViews *ChatViewsFilterer) FilterNewCommentaryEvent(opts *bind.FilterOpts) (*ChatViewsNewCommentaryEventIterator, error) {

	logs, sub, err := _ChatViews.contract.FilterLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return &ChatViewsNewCommentaryEventIterator{contract: _ChatViews.contract, event: "NewCommentaryEvent", logs: logs, sub: sub}, nil
}

// WatchNewCommentaryEvent is a free log subscription operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_ChatViews *ChatViewsFilterer) WatchNewCommentaryEvent(opts *bind.WatchOpts, sink chan<- *ChatViewsNewCommentaryEvent) (event.Subscription, error) {

	logs, sub, err := _ChatViews.contract.WatchLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatViewsNewCommentaryEvent)
				if err := _ChatViews.contract.UnpackLog(event, "NewCommentaryEvent", log); err != nil {
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

// ChatViewsNewPostEventIterator is returned from FilterNewPostEvent and is used to iterate over the raw logs and unpacked data for NewPostEvent events raised by the ChatViews contract.
type ChatViewsNewPostEventIterator struct {
	Event *ChatViewsNewPostEvent // Event containing the contract specifics and raw log

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
func (it *ChatViewsNewPostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatViewsNewPostEvent)
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
		it.Event = new(ChatViewsNewPostEvent)
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
func (it *ChatViewsNewPostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatViewsNewPostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatViewsNewPostEvent represents a NewPostEvent event raised by the ChatViews contract.
type ChatViewsNewPostEvent struct {
	PostId      *big.Int
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPostEvent is a free log retrieval operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsFilterer) FilterNewPostEvent(opts *bind.FilterOpts) (*ChatViewsNewPostEventIterator, error) {

	logs, sub, err := _ChatViews.contract.FilterLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return &ChatViewsNewPostEventIterator{contract: _ChatViews.contract, event: "NewPostEvent", logs: logs, sub: sub}, nil
}

// WatchNewPostEvent is a free log subscription operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsFilterer) WatchNewPostEvent(opts *bind.WatchOpts, sink chan<- *ChatViewsNewPostEvent) (event.Subscription, error) {

	logs, sub, err := _ChatViews.contract.WatchLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatViewsNewPostEvent)
				if err := _ChatViews.contract.UnpackLog(event, "NewPostEvent", log); err != nil {
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

// ChatViewsNewReplyEventIterator is returned from FilterNewReplyEvent and is used to iterate over the raw logs and unpacked data for NewReplyEvent events raised by the ChatViews contract.
type ChatViewsNewReplyEventIterator struct {
	Event *ChatViewsNewReplyEvent // Event containing the contract specifics and raw log

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
func (it *ChatViewsNewReplyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatViewsNewReplyEvent)
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
		it.Event = new(ChatViewsNewReplyEvent)
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
func (it *ChatViewsNewReplyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatViewsNewReplyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatViewsNewReplyEvent represents a NewReplyEvent event raised by the ChatViews contract.
type ChatViewsNewReplyEvent struct {
	ParentPostId *big.Int
	PostId       *big.Int
	IpfsHash     string
	Poster       common.Address
	BlockNumber  *big.Int
	Timestamp    *big.Int
	ContentType  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewReplyEvent is a free log retrieval operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsFilterer) FilterNewReplyEvent(opts *bind.FilterOpts) (*ChatViewsNewReplyEventIterator, error) {

	logs, sub, err := _ChatViews.contract.FilterLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return &ChatViewsNewReplyEventIterator{contract: _ChatViews.contract, event: "NewReplyEvent", logs: logs, sub: sub}, nil
}

// WatchNewReplyEvent is a free log subscription operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_ChatViews *ChatViewsFilterer) WatchNewReplyEvent(opts *bind.WatchOpts, sink chan<- *ChatViewsNewReplyEvent) (event.Subscription, error) {

	logs, sub, err := _ChatViews.contract.WatchLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatViewsNewReplyEvent)
				if err := _ChatViews.contract.UnpackLog(event, "NewReplyEvent", log); err != nil {
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

// CommonModelABI is the input ABI used to generate the binding from.
const CommonModelABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"nextActionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"name\":\"targetId\",\"type\":\"uint256\"},{\"name\":\"targetType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CommonModelBin is the compiled bytecode used for deploying new contracts.
const CommonModelBin = `0x6080604052600160005534801561001557600080fd5b5061010d806100256000396000f3fe6080604052600436106042577c0100000000000000000000000000000000000000000000000000000000600035046346abe73a81146047578063b6e7687314606b575b600080fd5b348015605257600080fd5b50605960bb565b60408051918252519081900360200190f35b348015607657600080fd5b50609160048036036020811015608b57600080fd5b503560c1565b6040518083815260200182600181111560a657fe5b60ff1681526020019250505060405180910390f35b60005481565b60009081526001602081905260409091208054910154909160ff9091169056fea165627a7a72305820dcd320ec78eaa00a0b297a9fe50d3676dab2a2d87ab0170252801b9bd927978e0029`

// DeployCommonModel deploys a new Ethereum contract, binding an instance of CommonModel to it.
func DeployCommonModel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CommonModel, error) {
	parsed, err := abi.JSON(strings.NewReader(CommonModelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommonModelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommonModel{CommonModelCaller: CommonModelCaller{contract: contract}, CommonModelTransactor: CommonModelTransactor{contract: contract}, CommonModelFilterer: CommonModelFilterer{contract: contract}}, nil
}

// CommonModel is an auto generated Go binding around an Ethereum contract.
type CommonModel struct {
	CommonModelCaller     // Read-only binding to the contract
	CommonModelTransactor // Write-only binding to the contract
	CommonModelFilterer   // Log filterer for contract events
}

// CommonModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommonModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommonModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommonModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommonModelSession struct {
	Contract     *CommonModel      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommonModelCallerSession struct {
	Contract *CommonModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CommonModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommonModelTransactorSession struct {
	Contract     *CommonModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CommonModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommonModelRaw struct {
	Contract *CommonModel // Generic contract binding to access the raw methods on
}

// CommonModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommonModelCallerRaw struct {
	Contract *CommonModelCaller // Generic read-only contract binding to access the raw methods on
}

// CommonModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommonModelTransactorRaw struct {
	Contract *CommonModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommonModel creates a new instance of CommonModel, bound to a specific deployed contract.
func NewCommonModel(address common.Address, backend bind.ContractBackend) (*CommonModel, error) {
	contract, err := bindCommonModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommonModel{CommonModelCaller: CommonModelCaller{contract: contract}, CommonModelTransactor: CommonModelTransactor{contract: contract}, CommonModelFilterer: CommonModelFilterer{contract: contract}}, nil
}

// NewCommonModelCaller creates a new read-only instance of CommonModel, bound to a specific deployed contract.
func NewCommonModelCaller(address common.Address, caller bind.ContractCaller) (*CommonModelCaller, error) {
	contract, err := bindCommonModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommonModelCaller{contract: contract}, nil
}

// NewCommonModelTransactor creates a new write-only instance of CommonModel, bound to a specific deployed contract.
func NewCommonModelTransactor(address common.Address, transactor bind.ContractTransactor) (*CommonModelTransactor, error) {
	contract, err := bindCommonModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommonModelTransactor{contract: contract}, nil
}

// NewCommonModelFilterer creates a new log filterer instance of CommonModel, bound to a specific deployed contract.
func NewCommonModelFilterer(address common.Address, filterer bind.ContractFilterer) (*CommonModelFilterer, error) {
	contract, err := bindCommonModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommonModelFilterer{contract: contract}, nil
}

// bindCommonModel binds a generic wrapper to an already deployed contract.
func bindCommonModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommonModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommonModel *CommonModelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CommonModel.Contract.CommonModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommonModel *CommonModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommonModel.Contract.CommonModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommonModel *CommonModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommonModel.Contract.CommonModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommonModel *CommonModelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CommonModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommonModel *CommonModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommonModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommonModel *CommonModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommonModel.Contract.contract.Transact(opts, method, params...)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_CommonModel *CommonModelCaller) GetAction(opts *bind.CallOpts, actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	ret := new(struct {
		TargetId   *big.Int
		TargetType uint8
	})
	out := ret
	err := _CommonModel.contract.Call(opts, out, "getAction", actionId)
	return *ret, err
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_CommonModel *CommonModelSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _CommonModel.Contract.GetAction(&_CommonModel.CallOpts, actionId)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_CommonModel *CommonModelCallerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _CommonModel.Contract.GetAction(&_CommonModel.CallOpts, actionId)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_CommonModel *CommonModelCaller) NextActionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CommonModel.contract.Call(opts, out, "nextActionId")
	return *ret0, err
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_CommonModel *CommonModelSession) NextActionId() (*big.Int, error) {
	return _CommonModel.Contract.NextActionId(&_CommonModel.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_CommonModel *CommonModelCallerSession) NextActionId() (*big.Int, error) {
	return _CommonModel.Contract.NextActionId(&_CommonModel.CallOpts)
}

// DatabaseIndexerABI is the input ABI used to generate the binding from.
const DatabaseIndexerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"latestDatabaseIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextActionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"setNewDB\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestDBHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"currentDatabaseIndex\",\"type\":\"uint256\"}],\"name\":\"getDB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"name\":\"targetId\",\"type\":\"uint256\"},{\"name\":\"targetType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestDBIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentDatabaseIndex\",\"type\":\"uint256\"}],\"name\":\"DBUpdateEvent\",\"type\":\"event\"}]"

// DatabaseIndexerBin is the compiled bytecode used for deploying new contracts.
const DatabaseIndexerBin = `0x60806040526001600055600060035534801561001a57600080fd5b5060028054600160a060020a0319163317905561056f8061003c6000396000f3fe60806040526004361061007c577c0100000000000000000000000000000000000000000000000000000000600035046339fa083d811461008157806346abe73a146100a8578063474c49dd146100bd5780637bda06c21461017257806398c01f25146101fc578063b6e7687314610226578063ca4454e51461027b575b600080fd5b34801561008d57600080fd5b50610096610290565b60408051918252519081900360200190f35b3480156100b457600080fd5b50610096610296565b3480156100c957600080fd5b50610170600480360360208110156100e057600080fd5b8101906020810181356401000000008111156100fb57600080fd5b82018360208201111561010d57600080fd5b8035906020019184600183028401116401000000008311171561012f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061029c945050505050565b005b34801561017e57600080fd5b50610187610323565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c15781810151838201526020016101a9565b50505050905090810190601f1680156101ee5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020857600080fd5b506101876004803603602081101561021f57600080fd5b50356103c6565b34801561023257600080fd5b506102506004803603602081101561024957600080fd5b5035610485565b6040518083815260200182600181111561026657fe5b60ff1681526020019250505060405180910390f35b34801561028757600080fd5b506100966104a5565b60035481565b60005481565b60025473ffffffffffffffffffffffffffffffffffffffff1633146102c057600080fd5b6003805460010190819055600090815260046020908152604090912082516102ea928401906104ab565b5060035460408051918252517f50634acbe4d09fc2a5c9ee5d100551a5a9e0162ed4bf7cb06c3115a7cfb25ed99181900360200190a150565b60035460009081526004602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103bb5780601f10610390576101008083540402835291602001916103bb565b820191906000526020600020905b81548152906001019060200180831161039e57829003601f168201915b505050505090505b90565b6060600082101580156103db57506003548211155b15156103e657600080fd5b60008281526004602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156104795780601f1061044e57610100808354040283529160200191610479565b820191906000526020600020905b81548152906001019060200180831161045c57829003601f168201915b50505050509050919050565b60009081526001602081905260409091208054910154909160ff90911690565b60035490565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104ec57805160ff1916838001178555610519565b82800160010185558215610519579182015b828111156105195782518255916020019190600101906104fe565b50610525929150610529565b5090565b6103c391905b80821115610525576000815560010161052f56fea165627a7a72305820e3b2456a037613501923423fba2b2cb49abde852a03071ece1c776ed68066f640029`

// DeployDatabaseIndexer deploys a new Ethereum contract, binding an instance of DatabaseIndexer to it.
func DeployDatabaseIndexer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DatabaseIndexer, error) {
	parsed, err := abi.JSON(strings.NewReader(DatabaseIndexerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatabaseIndexerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DatabaseIndexer{DatabaseIndexerCaller: DatabaseIndexerCaller{contract: contract}, DatabaseIndexerTransactor: DatabaseIndexerTransactor{contract: contract}, DatabaseIndexerFilterer: DatabaseIndexerFilterer{contract: contract}}, nil
}

// DatabaseIndexer is an auto generated Go binding around an Ethereum contract.
type DatabaseIndexer struct {
	DatabaseIndexerCaller     // Read-only binding to the contract
	DatabaseIndexerTransactor // Write-only binding to the contract
	DatabaseIndexerFilterer   // Log filterer for contract events
}

// DatabaseIndexerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatabaseIndexerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatabaseIndexerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatabaseIndexerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatabaseIndexerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatabaseIndexerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatabaseIndexerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatabaseIndexerSession struct {
	Contract     *DatabaseIndexer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatabaseIndexerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatabaseIndexerCallerSession struct {
	Contract *DatabaseIndexerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DatabaseIndexerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatabaseIndexerTransactorSession struct {
	Contract     *DatabaseIndexerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DatabaseIndexerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatabaseIndexerRaw struct {
	Contract *DatabaseIndexer // Generic contract binding to access the raw methods on
}

// DatabaseIndexerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatabaseIndexerCallerRaw struct {
	Contract *DatabaseIndexerCaller // Generic read-only contract binding to access the raw methods on
}

// DatabaseIndexerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatabaseIndexerTransactorRaw struct {
	Contract *DatabaseIndexerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatabaseIndexer creates a new instance of DatabaseIndexer, bound to a specific deployed contract.
func NewDatabaseIndexer(address common.Address, backend bind.ContractBackend) (*DatabaseIndexer, error) {
	contract, err := bindDatabaseIndexer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DatabaseIndexer{DatabaseIndexerCaller: DatabaseIndexerCaller{contract: contract}, DatabaseIndexerTransactor: DatabaseIndexerTransactor{contract: contract}, DatabaseIndexerFilterer: DatabaseIndexerFilterer{contract: contract}}, nil
}

// NewDatabaseIndexerCaller creates a new read-only instance of DatabaseIndexer, bound to a specific deployed contract.
func NewDatabaseIndexerCaller(address common.Address, caller bind.ContractCaller) (*DatabaseIndexerCaller, error) {
	contract, err := bindDatabaseIndexer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatabaseIndexerCaller{contract: contract}, nil
}

// NewDatabaseIndexerTransactor creates a new write-only instance of DatabaseIndexer, bound to a specific deployed contract.
func NewDatabaseIndexerTransactor(address common.Address, transactor bind.ContractTransactor) (*DatabaseIndexerTransactor, error) {
	contract, err := bindDatabaseIndexer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatabaseIndexerTransactor{contract: contract}, nil
}

// NewDatabaseIndexerFilterer creates a new log filterer instance of DatabaseIndexer, bound to a specific deployed contract.
func NewDatabaseIndexerFilterer(address common.Address, filterer bind.ContractFilterer) (*DatabaseIndexerFilterer, error) {
	contract, err := bindDatabaseIndexer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatabaseIndexerFilterer{contract: contract}, nil
}

// bindDatabaseIndexer binds a generic wrapper to an already deployed contract.
func bindDatabaseIndexer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatabaseIndexerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DatabaseIndexer *DatabaseIndexerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DatabaseIndexer.Contract.DatabaseIndexerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DatabaseIndexer *DatabaseIndexerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DatabaseIndexer.Contract.DatabaseIndexerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DatabaseIndexer *DatabaseIndexerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DatabaseIndexer.Contract.DatabaseIndexerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DatabaseIndexer *DatabaseIndexerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DatabaseIndexer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DatabaseIndexer *DatabaseIndexerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DatabaseIndexer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DatabaseIndexer *DatabaseIndexerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DatabaseIndexer.Contract.contract.Transact(opts, method, params...)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_DatabaseIndexer *DatabaseIndexerCaller) GetAction(opts *bind.CallOpts, actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	ret := new(struct {
		TargetId   *big.Int
		TargetType uint8
	})
	out := ret
	err := _DatabaseIndexer.contract.Call(opts, out, "getAction", actionId)
	return *ret, err
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_DatabaseIndexer *DatabaseIndexerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _DatabaseIndexer.Contract.GetAction(&_DatabaseIndexer.CallOpts, actionId)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_DatabaseIndexer *DatabaseIndexerCallerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _DatabaseIndexer.Contract.GetAction(&_DatabaseIndexer.CallOpts, actionId)
}

// GetDB is a free data retrieval call binding the contract method 0x98c01f25.
//
// Solidity: function getDB(currentDatabaseIndex uint256) constant returns(string)
func (_DatabaseIndexer *DatabaseIndexerCaller) GetDB(opts *bind.CallOpts, currentDatabaseIndex *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DatabaseIndexer.contract.Call(opts, out, "getDB", currentDatabaseIndex)
	return *ret0, err
}

// GetDB is a free data retrieval call binding the contract method 0x98c01f25.
//
// Solidity: function getDB(currentDatabaseIndex uint256) constant returns(string)
func (_DatabaseIndexer *DatabaseIndexerSession) GetDB(currentDatabaseIndex *big.Int) (string, error) {
	return _DatabaseIndexer.Contract.GetDB(&_DatabaseIndexer.CallOpts, currentDatabaseIndex)
}

// GetDB is a free data retrieval call binding the contract method 0x98c01f25.
//
// Solidity: function getDB(currentDatabaseIndex uint256) constant returns(string)
func (_DatabaseIndexer *DatabaseIndexerCallerSession) GetDB(currentDatabaseIndex *big.Int) (string, error) {
	return _DatabaseIndexer.Contract.GetDB(&_DatabaseIndexer.CallOpts, currentDatabaseIndex)
}

// GetLatestDBHash is a free data retrieval call binding the contract method 0x7bda06c2.
//
// Solidity: function getLatestDBHash() constant returns(string)
func (_DatabaseIndexer *DatabaseIndexerCaller) GetLatestDBHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DatabaseIndexer.contract.Call(opts, out, "getLatestDBHash")
	return *ret0, err
}

// GetLatestDBHash is a free data retrieval call binding the contract method 0x7bda06c2.
//
// Solidity: function getLatestDBHash() constant returns(string)
func (_DatabaseIndexer *DatabaseIndexerSession) GetLatestDBHash() (string, error) {
	return _DatabaseIndexer.Contract.GetLatestDBHash(&_DatabaseIndexer.CallOpts)
}

// GetLatestDBHash is a free data retrieval call binding the contract method 0x7bda06c2.
//
// Solidity: function getLatestDBHash() constant returns(string)
func (_DatabaseIndexer *DatabaseIndexerCallerSession) GetLatestDBHash() (string, error) {
	return _DatabaseIndexer.Contract.GetLatestDBHash(&_DatabaseIndexer.CallOpts)
}

// GetLatestDBIndex is a free data retrieval call binding the contract method 0xca4454e5.
//
// Solidity: function getLatestDBIndex() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerCaller) GetLatestDBIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DatabaseIndexer.contract.Call(opts, out, "getLatestDBIndex")
	return *ret0, err
}

// GetLatestDBIndex is a free data retrieval call binding the contract method 0xca4454e5.
//
// Solidity: function getLatestDBIndex() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerSession) GetLatestDBIndex() (*big.Int, error) {
	return _DatabaseIndexer.Contract.GetLatestDBIndex(&_DatabaseIndexer.CallOpts)
}

// GetLatestDBIndex is a free data retrieval call binding the contract method 0xca4454e5.
//
// Solidity: function getLatestDBIndex() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerCallerSession) GetLatestDBIndex() (*big.Int, error) {
	return _DatabaseIndexer.Contract.GetLatestDBIndex(&_DatabaseIndexer.CallOpts)
}

// LatestDatabaseIndex is a free data retrieval call binding the contract method 0x39fa083d.
//
// Solidity: function latestDatabaseIndex() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerCaller) LatestDatabaseIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DatabaseIndexer.contract.Call(opts, out, "latestDatabaseIndex")
	return *ret0, err
}

// LatestDatabaseIndex is a free data retrieval call binding the contract method 0x39fa083d.
//
// Solidity: function latestDatabaseIndex() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerSession) LatestDatabaseIndex() (*big.Int, error) {
	return _DatabaseIndexer.Contract.LatestDatabaseIndex(&_DatabaseIndexer.CallOpts)
}

// LatestDatabaseIndex is a free data retrieval call binding the contract method 0x39fa083d.
//
// Solidity: function latestDatabaseIndex() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerCallerSession) LatestDatabaseIndex() (*big.Int, error) {
	return _DatabaseIndexer.Contract.LatestDatabaseIndex(&_DatabaseIndexer.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerCaller) NextActionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DatabaseIndexer.contract.Call(opts, out, "nextActionId")
	return *ret0, err
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerSession) NextActionId() (*big.Int, error) {
	return _DatabaseIndexer.Contract.NextActionId(&_DatabaseIndexer.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_DatabaseIndexer *DatabaseIndexerCallerSession) NextActionId() (*big.Int, error) {
	return _DatabaseIndexer.Contract.NextActionId(&_DatabaseIndexer.CallOpts)
}

// SetNewDB is a paid mutator transaction binding the contract method 0x474c49dd.
//
// Solidity: function setNewDB(ipfsHash string) returns()
func (_DatabaseIndexer *DatabaseIndexerTransactor) SetNewDB(opts *bind.TransactOpts, ipfsHash string) (*types.Transaction, error) {
	return _DatabaseIndexer.contract.Transact(opts, "setNewDB", ipfsHash)
}

// SetNewDB is a paid mutator transaction binding the contract method 0x474c49dd.
//
// Solidity: function setNewDB(ipfsHash string) returns()
func (_DatabaseIndexer *DatabaseIndexerSession) SetNewDB(ipfsHash string) (*types.Transaction, error) {
	return _DatabaseIndexer.Contract.SetNewDB(&_DatabaseIndexer.TransactOpts, ipfsHash)
}

// SetNewDB is a paid mutator transaction binding the contract method 0x474c49dd.
//
// Solidity: function setNewDB(ipfsHash string) returns()
func (_DatabaseIndexer *DatabaseIndexerTransactorSession) SetNewDB(ipfsHash string) (*types.Transaction, error) {
	return _DatabaseIndexer.Contract.SetNewDB(&_DatabaseIndexer.TransactOpts, ipfsHash)
}

// DatabaseIndexerDBUpdateEventIterator is returned from FilterDBUpdateEvent and is used to iterate over the raw logs and unpacked data for DBUpdateEvent events raised by the DatabaseIndexer contract.
type DatabaseIndexerDBUpdateEventIterator struct {
	Event *DatabaseIndexerDBUpdateEvent // Event containing the contract specifics and raw log

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
func (it *DatabaseIndexerDBUpdateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatabaseIndexerDBUpdateEvent)
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
		it.Event = new(DatabaseIndexerDBUpdateEvent)
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
func (it *DatabaseIndexerDBUpdateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatabaseIndexerDBUpdateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatabaseIndexerDBUpdateEvent represents a DBUpdateEvent event raised by the DatabaseIndexer contract.
type DatabaseIndexerDBUpdateEvent struct {
	CurrentDatabaseIndex *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDBUpdateEvent is a free log retrieval operation binding the contract event 0x50634acbe4d09fc2a5c9ee5d100551a5a9e0162ed4bf7cb06c3115a7cfb25ed9.
//
// Solidity: e DBUpdateEvent(currentDatabaseIndex uint256)
func (_DatabaseIndexer *DatabaseIndexerFilterer) FilterDBUpdateEvent(opts *bind.FilterOpts) (*DatabaseIndexerDBUpdateEventIterator, error) {

	logs, sub, err := _DatabaseIndexer.contract.FilterLogs(opts, "DBUpdateEvent")
	if err != nil {
		return nil, err
	}
	return &DatabaseIndexerDBUpdateEventIterator{contract: _DatabaseIndexer.contract, event: "DBUpdateEvent", logs: logs, sub: sub}, nil
}

// WatchDBUpdateEvent is a free log subscription operation binding the contract event 0x50634acbe4d09fc2a5c9ee5d100551a5a9e0162ed4bf7cb06c3115a7cfb25ed9.
//
// Solidity: e DBUpdateEvent(currentDatabaseIndex uint256)
func (_DatabaseIndexer *DatabaseIndexerFilterer) WatchDBUpdateEvent(opts *bind.WatchOpts, sink chan<- *DatabaseIndexerDBUpdateEvent) (event.Subscription, error) {

	logs, sub, err := _DatabaseIndexer.contract.WatchLogs(opts, "DBUpdateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatabaseIndexerDBUpdateEvent)
				if err := _DatabaseIndexer.contract.UnpackLog(event, "DBUpdateEvent", log); err != nil {
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

// IdentityAdminServiceABI is the input ABI used to generate the binding from.
const IdentityAdminServiceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"requestPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setRequestPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identityAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityRequestId\",\"type\":\"uint256\"}],\"name\":\"NewRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NewApprovedProviderEvent\",\"type\":\"event\"}]"

// IdentityAdminServiceBin is the compiled bytecode used for deploying new contracts.
const IdentityAdminServiceBin = `0x608060405260016002819055600381905560055560008054600160a060020a03191633179055610a27806100346000396000f3fe60806040526004361061007c577c010000000000000000000000000000000000000000000000000000000060003504631604f9ea81146100815780633fcf7ca1146100a85780634d625d99146100d457806375efbecf146100e95780637e889537146100fe5780639119c8d71461012f578063b759f95414610162575b600080fd5b34801561008d57600080fd5b5061009661018c565b60408051918252519081900360200190f35b3480156100b457600080fd5b506100d2600480360360208110156100cb57600080fd5b5035610192565b005b3480156100e057600080fd5b506100966101ae565b3480156100f557600080fd5b506100966101b4565b34801561010a57600080fd5b506101136101ba565b60408051600160a060020a039092168252519081900360200190f35b34801561013b57600080fd5b506100d26004803603604081101561015257600080fd5b508035906020013560ff166101c9565b34801561016e57600080fd5b506100d26004803603602081101561018557600080fd5b50356104a8565b60015481565b600054600160a060020a031633146101a957600080fd5b600155565b60025481565b60035481565b600054600160a060020a031681565b600054600160a060020a031633146101e057600080fd5b6101e86109bb565b600083815260046020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f810186900486028301860187528083529295939493860193919290918301828280156102a55780601f1061027a576101008083540402835291602001916102a5565b820191906000526020600020905b81548152906001019060200180831161028857829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103395780601f1061030e57610100808354040283529160200191610339565b820191906000526020600020905b81548152906001019060200180831161031c57829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103cd5780601f106103a2576101008083540402835291602001916103cd565b820191906000526020600020905b8154815290600101906020018083116103b057829003601f168201915b5050509183525050600591909101546020918201528101516040820151606083015192935090916000610401848484610887565b905060066000828152602001908152602001600020600201604080519081016040528089815260200188600381111561043657fe5b905281546001818101808555600094855260209485902084516002909402019283559383015182820180549192909160ff19169083600381111561047657fe5b02179055505050600160a060020a03909416600090815260076020526040902055505060038054600101905550505050565b600054600160a060020a031633146104bf57600080fd5b6104c76109bb565b600082815260046020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f810186900486028301860187528083529295939493860193919290918301828280156105845780601f1061055957610100808354040283529160200191610584565b820191906000526020600020905b81548152906001019060200180831161056757829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156106185780601f106105ed57610100808354040283529160200191610618565b820191906000526020600020905b8154815290600101906020018083116105fb57829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156106ac5780601f10610681576101008083540402835291602001916106ac565b820191906000526020600020905b81548152906001019060200180831161068f57829003601f168201915b50505091835250506005919091015460209182015281015160408201516060830151929350909160006106e0848484610887565b600081815260066020908152604080832081516060810183528b8152426103e8028185019081526001828501818152938101805480830182559088528688209351600390910290930192835590519082015590516002909101805460ff1916911515919091179055600160a060020a0388168352600782529182902083905590518551929350839260089287929182918401908083835b602083106107965780518252601f199092019160209182019101610777565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842087519094889450925082918401908083835b602083106107f25780518252601f1990920191602091820191016107d3565b51815160001960209485036101000a8101918216911992909216179091529390910195865260408051968790038201872097909755600380546001908101909155600089815260068352889020015488875290920191850191909152505082517f32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae9799392819003909201919050a1505050505050565b60008054600160a060020a0316331461089f57600080fd5b600160a060020a038416600090815260076020526040812054908111156108c75790506109b4565b6008846040518082805190602001908083835b602083106108f95780518252601f1990920191602091820191016108da565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842087519094889450925082918401908083835b602083106109555780518252601f199092019160209182019101610936565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220549250505060008111156109955790506109b4565b5050600580546000818152600660205260409020819055600181019091555b9392505050565b60c060405190810160405280600081526020016000600160a060020a0316815260200160608152602001606081526020016060815260200160008152509056fea165627a7a72305820640497cc1904096407795cbc8eb5e79457659f62a1145bd4abdc374ffa47fb800029`

// DeployIdentityAdminService deploys a new Ethereum contract, binding an instance of IdentityAdminService to it.
func DeployIdentityAdminService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityAdminService, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityAdminServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityAdminServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityAdminService{IdentityAdminServiceCaller: IdentityAdminServiceCaller{contract: contract}, IdentityAdminServiceTransactor: IdentityAdminServiceTransactor{contract: contract}, IdentityAdminServiceFilterer: IdentityAdminServiceFilterer{contract: contract}}, nil
}

// IdentityAdminService is an auto generated Go binding around an Ethereum contract.
type IdentityAdminService struct {
	IdentityAdminServiceCaller     // Read-only binding to the contract
	IdentityAdminServiceTransactor // Write-only binding to the contract
	IdentityAdminServiceFilterer   // Log filterer for contract events
}

// IdentityAdminServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityAdminServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityAdminServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityAdminServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityAdminServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityAdminServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityAdminServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityAdminServiceSession struct {
	Contract     *IdentityAdminService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IdentityAdminServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityAdminServiceCallerSession struct {
	Contract *IdentityAdminServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IdentityAdminServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityAdminServiceTransactorSession struct {
	Contract     *IdentityAdminServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IdentityAdminServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityAdminServiceRaw struct {
	Contract *IdentityAdminService // Generic contract binding to access the raw methods on
}

// IdentityAdminServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityAdminServiceCallerRaw struct {
	Contract *IdentityAdminServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityAdminServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityAdminServiceTransactorRaw struct {
	Contract *IdentityAdminServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityAdminService creates a new instance of IdentityAdminService, bound to a specific deployed contract.
func NewIdentityAdminService(address common.Address, backend bind.ContractBackend) (*IdentityAdminService, error) {
	contract, err := bindIdentityAdminService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityAdminService{IdentityAdminServiceCaller: IdentityAdminServiceCaller{contract: contract}, IdentityAdminServiceTransactor: IdentityAdminServiceTransactor{contract: contract}, IdentityAdminServiceFilterer: IdentityAdminServiceFilterer{contract: contract}}, nil
}

// NewIdentityAdminServiceCaller creates a new read-only instance of IdentityAdminService, bound to a specific deployed contract.
func NewIdentityAdminServiceCaller(address common.Address, caller bind.ContractCaller) (*IdentityAdminServiceCaller, error) {
	contract, err := bindIdentityAdminService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityAdminServiceCaller{contract: contract}, nil
}

// NewIdentityAdminServiceTransactor creates a new write-only instance of IdentityAdminService, bound to a specific deployed contract.
func NewIdentityAdminServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityAdminServiceTransactor, error) {
	contract, err := bindIdentityAdminService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityAdminServiceTransactor{contract: contract}, nil
}

// NewIdentityAdminServiceFilterer creates a new log filterer instance of IdentityAdminService, bound to a specific deployed contract.
func NewIdentityAdminServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityAdminServiceFilterer, error) {
	contract, err := bindIdentityAdminService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityAdminServiceFilterer{contract: contract}, nil
}

// bindIdentityAdminService binds a generic wrapper to an already deployed contract.
func bindIdentityAdminService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityAdminServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityAdminService *IdentityAdminServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityAdminService.Contract.IdentityAdminServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityAdminService *IdentityAdminServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.IdentityAdminServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityAdminService *IdentityAdminServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.IdentityAdminServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityAdminService *IdentityAdminServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityAdminService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityAdminService *IdentityAdminServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityAdminService *IdentityAdminServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.contract.Transact(opts, method, params...)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityAdminService *IdentityAdminServiceCaller) IdentityAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityAdminService.contract.Call(opts, out, "identityAdmin")
	return *ret0, err
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityAdminService *IdentityAdminServiceSession) IdentityAdmin() (common.Address, error) {
	return _IdentityAdminService.Contract.IdentityAdmin(&_IdentityAdminService.CallOpts)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityAdminService *IdentityAdminServiceCallerSession) IdentityAdmin() (common.Address, error) {
	return _IdentityAdminService.Contract.IdentityAdmin(&_IdentityAdminService.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceCaller) NextIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityAdminService.contract.Call(opts, out, "nextIdentityRequestIndex")
	return *ret0, err
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _IdentityAdminService.Contract.NextIdentityRequestIndex(&_IdentityAdminService.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceCallerSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _IdentityAdminService.Contract.NextIdentityRequestIndex(&_IdentityAdminService.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceCaller) PendingIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityAdminService.contract.Call(opts, out, "pendingIdentityRequestIndex")
	return *ret0, err
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _IdentityAdminService.Contract.PendingIdentityRequestIndex(&_IdentityAdminService.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceCallerSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _IdentityAdminService.Contract.PendingIdentityRequestIndex(&_IdentityAdminService.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceCaller) RequestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityAdminService.contract.Call(opts, out, "requestPrice")
	return *ret0, err
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceSession) RequestPrice() (*big.Int, error) {
	return _IdentityAdminService.Contract.RequestPrice(&_IdentityAdminService.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityAdminService *IdentityAdminServiceCallerSession) RequestPrice() (*big.Int, error) {
	return _IdentityAdminService.Contract.RequestPrice(&_IdentityAdminService.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_IdentityAdminService *IdentityAdminServiceTransactor) Approve(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _IdentityAdminService.contract.Transact(opts, "approve", requestId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_IdentityAdminService *IdentityAdminServiceSession) Approve(requestId *big.Int) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.Approve(&_IdentityAdminService.TransactOpts, requestId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_IdentityAdminService *IdentityAdminServiceTransactorSession) Approve(requestId *big.Int) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.Approve(&_IdentityAdminService.TransactOpts, requestId)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_IdentityAdminService *IdentityAdminServiceTransactor) Reject(opts *bind.TransactOpts, requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _IdentityAdminService.contract.Transact(opts, "reject", requestId, reason)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_IdentityAdminService *IdentityAdminServiceSession) Reject(requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.Reject(&_IdentityAdminService.TransactOpts, requestId, reason)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_IdentityAdminService *IdentityAdminServiceTransactorSession) Reject(requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.Reject(&_IdentityAdminService.TransactOpts, requestId, reason)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_IdentityAdminService *IdentityAdminServiceTransactor) SetRequestPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _IdentityAdminService.contract.Transact(opts, "setRequestPrice", price)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_IdentityAdminService *IdentityAdminServiceSession) SetRequestPrice(price *big.Int) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.SetRequestPrice(&_IdentityAdminService.TransactOpts, price)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_IdentityAdminService *IdentityAdminServiceTransactorSession) SetRequestPrice(price *big.Int) (*types.Transaction, error) {
	return _IdentityAdminService.Contract.SetRequestPrice(&_IdentityAdminService.TransactOpts, price)
}

// IdentityAdminServiceNewApprovedProviderEventIterator is returned from FilterNewApprovedProviderEvent and is used to iterate over the raw logs and unpacked data for NewApprovedProviderEvent events raised by the IdentityAdminService contract.
type IdentityAdminServiceNewApprovedProviderEventIterator struct {
	Event *IdentityAdminServiceNewApprovedProviderEvent // Event containing the contract specifics and raw log

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
func (it *IdentityAdminServiceNewApprovedProviderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityAdminServiceNewApprovedProviderEvent)
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
		it.Event = new(IdentityAdminServiceNewApprovedProviderEvent)
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
func (it *IdentityAdminServiceNewApprovedProviderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityAdminServiceNewApprovedProviderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityAdminServiceNewApprovedProviderEvent represents a NewApprovedProviderEvent event raised by the IdentityAdminService contract.
type IdentityAdminServiceNewApprovedProviderEvent struct {
	IdentityId *big.Int
	ProviderId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewApprovedProviderEvent is a free log retrieval operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_IdentityAdminService *IdentityAdminServiceFilterer) FilterNewApprovedProviderEvent(opts *bind.FilterOpts) (*IdentityAdminServiceNewApprovedProviderEventIterator, error) {

	logs, sub, err := _IdentityAdminService.contract.FilterLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return &IdentityAdminServiceNewApprovedProviderEventIterator{contract: _IdentityAdminService.contract, event: "NewApprovedProviderEvent", logs: logs, sub: sub}, nil
}

// WatchNewApprovedProviderEvent is a free log subscription operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_IdentityAdminService *IdentityAdminServiceFilterer) WatchNewApprovedProviderEvent(opts *bind.WatchOpts, sink chan<- *IdentityAdminServiceNewApprovedProviderEvent) (event.Subscription, error) {

	logs, sub, err := _IdentityAdminService.contract.WatchLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityAdminServiceNewApprovedProviderEvent)
				if err := _IdentityAdminService.contract.UnpackLog(event, "NewApprovedProviderEvent", log); err != nil {
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

// IdentityAdminServiceNewRequestEventIterator is returned from FilterNewRequestEvent and is used to iterate over the raw logs and unpacked data for NewRequestEvent events raised by the IdentityAdminService contract.
type IdentityAdminServiceNewRequestEventIterator struct {
	Event *IdentityAdminServiceNewRequestEvent // Event containing the contract specifics and raw log

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
func (it *IdentityAdminServiceNewRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityAdminServiceNewRequestEvent)
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
		it.Event = new(IdentityAdminServiceNewRequestEvent)
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
func (it *IdentityAdminServiceNewRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityAdminServiceNewRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityAdminServiceNewRequestEvent represents a NewRequestEvent event raised by the IdentityAdminService contract.
type IdentityAdminServiceNewRequestEvent struct {
	IdentityRequestId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRequestEvent is a free log retrieval operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_IdentityAdminService *IdentityAdminServiceFilterer) FilterNewRequestEvent(opts *bind.FilterOpts) (*IdentityAdminServiceNewRequestEventIterator, error) {

	logs, sub, err := _IdentityAdminService.contract.FilterLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return &IdentityAdminServiceNewRequestEventIterator{contract: _IdentityAdminService.contract, event: "NewRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewRequestEvent is a free log subscription operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_IdentityAdminService *IdentityAdminServiceFilterer) WatchNewRequestEvent(opts *bind.WatchOpts, sink chan<- *IdentityAdminServiceNewRequestEvent) (event.Subscription, error) {

	logs, sub, err := _IdentityAdminService.contract.WatchLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityAdminServiceNewRequestEvent)
				if err := _IdentityAdminService.contract.UnpackLog(event, "NewRequestEvent", log); err != nil {
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

// IdentityModelABI is the input ABI used to generate the binding from.
const IdentityModelABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"requestPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identityAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityRequestId\",\"type\":\"uint256\"}],\"name\":\"NewRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NewApprovedProviderEvent\",\"type\":\"event\"}]"

// IdentityModelBin is the compiled bytecode used for deploying new contracts.
const IdentityModelBin = `0x608060405260016002556001600355600160055534801561001f57600080fd5b5060008054600160a060020a03191633179055610149806100416000396000f3fe60806040526004361061005b577c010000000000000000000000000000000000000000000000000000000060003504631604f9ea81146100605780634d625d991461008757806375efbecf1461009c5780637e889537146100b1575b600080fd5b34801561006c57600080fd5b506100756100ef565b60408051918252519081900360200190f35b34801561009357600080fd5b506100756100f5565b3480156100a857600080fd5b506100756100fb565b3480156100bd57600080fd5b506100c6610101565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60015481565b60025481565b60035481565b60005473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a7230582082ebe0d3be343bb0e94146df6bf09a2ae324be3972747d5be89ca7276904cc3a0029`

// DeployIdentityModel deploys a new Ethereum contract, binding an instance of IdentityModel to it.
func DeployIdentityModel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityModel, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityModelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityModelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityModel{IdentityModelCaller: IdentityModelCaller{contract: contract}, IdentityModelTransactor: IdentityModelTransactor{contract: contract}, IdentityModelFilterer: IdentityModelFilterer{contract: contract}}, nil
}

// IdentityModel is an auto generated Go binding around an Ethereum contract.
type IdentityModel struct {
	IdentityModelCaller     // Read-only binding to the contract
	IdentityModelTransactor // Write-only binding to the contract
	IdentityModelFilterer   // Log filterer for contract events
}

// IdentityModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityModelSession struct {
	Contract     *IdentityModel    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityModelCallerSession struct {
	Contract *IdentityModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IdentityModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityModelTransactorSession struct {
	Contract     *IdentityModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IdentityModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityModelRaw struct {
	Contract *IdentityModel // Generic contract binding to access the raw methods on
}

// IdentityModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityModelCallerRaw struct {
	Contract *IdentityModelCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityModelTransactorRaw struct {
	Contract *IdentityModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityModel creates a new instance of IdentityModel, bound to a specific deployed contract.
func NewIdentityModel(address common.Address, backend bind.ContractBackend) (*IdentityModel, error) {
	contract, err := bindIdentityModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityModel{IdentityModelCaller: IdentityModelCaller{contract: contract}, IdentityModelTransactor: IdentityModelTransactor{contract: contract}, IdentityModelFilterer: IdentityModelFilterer{contract: contract}}, nil
}

// NewIdentityModelCaller creates a new read-only instance of IdentityModel, bound to a specific deployed contract.
func NewIdentityModelCaller(address common.Address, caller bind.ContractCaller) (*IdentityModelCaller, error) {
	contract, err := bindIdentityModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityModelCaller{contract: contract}, nil
}

// NewIdentityModelTransactor creates a new write-only instance of IdentityModel, bound to a specific deployed contract.
func NewIdentityModelTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityModelTransactor, error) {
	contract, err := bindIdentityModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityModelTransactor{contract: contract}, nil
}

// NewIdentityModelFilterer creates a new log filterer instance of IdentityModel, bound to a specific deployed contract.
func NewIdentityModelFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityModelFilterer, error) {
	contract, err := bindIdentityModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityModelFilterer{contract: contract}, nil
}

// bindIdentityModel binds a generic wrapper to an already deployed contract.
func bindIdentityModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityModel *IdentityModelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityModel.Contract.IdentityModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityModel *IdentityModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityModel.Contract.IdentityModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityModel *IdentityModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityModel.Contract.IdentityModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityModel *IdentityModelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityModel *IdentityModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityModel *IdentityModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityModel.Contract.contract.Transact(opts, method, params...)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityModel *IdentityModelCaller) IdentityAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityModel.contract.Call(opts, out, "identityAdmin")
	return *ret0, err
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityModel *IdentityModelSession) IdentityAdmin() (common.Address, error) {
	return _IdentityModel.Contract.IdentityAdmin(&_IdentityModel.CallOpts)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityModel *IdentityModelCallerSession) IdentityAdmin() (common.Address, error) {
	return _IdentityModel.Contract.IdentityAdmin(&_IdentityModel.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityModel *IdentityModelCaller) NextIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityModel.contract.Call(opts, out, "nextIdentityRequestIndex")
	return *ret0, err
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityModel *IdentityModelSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _IdentityModel.Contract.NextIdentityRequestIndex(&_IdentityModel.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityModel *IdentityModelCallerSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _IdentityModel.Contract.NextIdentityRequestIndex(&_IdentityModel.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityModel *IdentityModelCaller) PendingIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityModel.contract.Call(opts, out, "pendingIdentityRequestIndex")
	return *ret0, err
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityModel *IdentityModelSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _IdentityModel.Contract.PendingIdentityRequestIndex(&_IdentityModel.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityModel *IdentityModelCallerSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _IdentityModel.Contract.PendingIdentityRequestIndex(&_IdentityModel.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityModel *IdentityModelCaller) RequestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityModel.contract.Call(opts, out, "requestPrice")
	return *ret0, err
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityModel *IdentityModelSession) RequestPrice() (*big.Int, error) {
	return _IdentityModel.Contract.RequestPrice(&_IdentityModel.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityModel *IdentityModelCallerSession) RequestPrice() (*big.Int, error) {
	return _IdentityModel.Contract.RequestPrice(&_IdentityModel.CallOpts)
}

// IdentityModelNewApprovedProviderEventIterator is returned from FilterNewApprovedProviderEvent and is used to iterate over the raw logs and unpacked data for NewApprovedProviderEvent events raised by the IdentityModel contract.
type IdentityModelNewApprovedProviderEventIterator struct {
	Event *IdentityModelNewApprovedProviderEvent // Event containing the contract specifics and raw log

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
func (it *IdentityModelNewApprovedProviderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityModelNewApprovedProviderEvent)
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
		it.Event = new(IdentityModelNewApprovedProviderEvent)
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
func (it *IdentityModelNewApprovedProviderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityModelNewApprovedProviderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityModelNewApprovedProviderEvent represents a NewApprovedProviderEvent event raised by the IdentityModel contract.
type IdentityModelNewApprovedProviderEvent struct {
	IdentityId *big.Int
	ProviderId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewApprovedProviderEvent is a free log retrieval operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_IdentityModel *IdentityModelFilterer) FilterNewApprovedProviderEvent(opts *bind.FilterOpts) (*IdentityModelNewApprovedProviderEventIterator, error) {

	logs, sub, err := _IdentityModel.contract.FilterLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return &IdentityModelNewApprovedProviderEventIterator{contract: _IdentityModel.contract, event: "NewApprovedProviderEvent", logs: logs, sub: sub}, nil
}

// WatchNewApprovedProviderEvent is a free log subscription operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_IdentityModel *IdentityModelFilterer) WatchNewApprovedProviderEvent(opts *bind.WatchOpts, sink chan<- *IdentityModelNewApprovedProviderEvent) (event.Subscription, error) {

	logs, sub, err := _IdentityModel.contract.WatchLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityModelNewApprovedProviderEvent)
				if err := _IdentityModel.contract.UnpackLog(event, "NewApprovedProviderEvent", log); err != nil {
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

// IdentityModelNewRequestEventIterator is returned from FilterNewRequestEvent and is used to iterate over the raw logs and unpacked data for NewRequestEvent events raised by the IdentityModel contract.
type IdentityModelNewRequestEventIterator struct {
	Event *IdentityModelNewRequestEvent // Event containing the contract specifics and raw log

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
func (it *IdentityModelNewRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityModelNewRequestEvent)
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
		it.Event = new(IdentityModelNewRequestEvent)
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
func (it *IdentityModelNewRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityModelNewRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityModelNewRequestEvent represents a NewRequestEvent event raised by the IdentityModel contract.
type IdentityModelNewRequestEvent struct {
	IdentityRequestId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRequestEvent is a free log retrieval operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_IdentityModel *IdentityModelFilterer) FilterNewRequestEvent(opts *bind.FilterOpts) (*IdentityModelNewRequestEventIterator, error) {

	logs, sub, err := _IdentityModel.contract.FilterLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return &IdentityModelNewRequestEventIterator{contract: _IdentityModel.contract, event: "NewRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewRequestEvent is a free log subscription operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_IdentityModel *IdentityModelFilterer) WatchNewRequestEvent(opts *bind.WatchOpts, sink chan<- *IdentityModelNewRequestEvent) (event.Subscription, error) {

	logs, sub, err := _IdentityModel.contract.WatchLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityModelNewRequestEvent)
				if err := _IdentityModel.contract.UnpackLog(event, "NewRequestEvent", log); err != nil {
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

// IdentityRegistryABI is the input ABI used to generate the binding from.
const IdentityRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"requestPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setRequestPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"submitRequest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityId\",\"type\":\"uint256\"}],\"name\":\"getIdentityInfo\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"providerCount\",\"type\":\"uint256\"},{\"name\":\"rejectionCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identityAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextRequest\",\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityId\",\"type\":\"uint256\"},{\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"getProviderInfo\",\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"identityAddress\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"getMyProviderInfo\",\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"identityAddress\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"getIdentityInfoByAddress\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"providerCount\",\"type\":\"uint256\"},{\"name\":\"rejectionCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMyIdentityInfo\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"providerCount\",\"type\":\"uint256\"},{\"name\":\"rejectionCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityRequestId\",\"type\":\"uint256\"}],\"name\":\"NewRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NewApprovedProviderEvent\",\"type\":\"event\"}]"

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
const IdentityRegistryBin = `0x608060405260016002819055600381905560055560008054600160a060020a03191633179055611a1b806100346000396000f3fe6080604052600436106100d4577c010000000000000000000000000000000000000000000000000000000060003504631604f9ea81146100d957806326757b73146101005780633fcf7ca114610295578063470c153b146102c15780634d625d9914610475578063632859c31461048a57806375efbecf146104d25780637e889537146104e75780639119c8d71461051857806395dfb8961461054b578063b759f95414610560578063ba438b3b1461058a578063d64c73a6146106ca578063e0774281146106f4578063f388ad2314610727575b600080fd5b3480156100e557600080fd5b506100ee61073c565b60408051918252519081900360200190f35b34801561010c57600080fd5b5061012a6004803603602081101561012357600080fd5b5035610742565b6040518087815260200186600160a060020a0316600160a060020a03168152602001806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b8381101561019357818101518382015260200161017b565b50505050905090810190601f1680156101c05780820380516001836020036101000a031916815260200191505b50848103835287518152875160209182019189019080838360005b838110156101f35781810151838201526020016101db565b50505050905090810190601f1680156102205780820380516001836020036101000a031916815260200191505b50848103825286518152865160209182019188019080838360005b8381101561025357818101518382015260200161023b565b50505050905090810190601f1680156102805780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3480156102a157600080fd5b506102bf600480360360208110156102b857600080fd5b5035610951565b005b6102bf600480360360608110156102d757600080fd5b8101906020810181356401000000008111156102f257600080fd5b82018360208201111561030457600080fd5b8035906020019184600183028401116401000000008311171561032657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561037957600080fd5b82018360208201111561038b57600080fd5b803590602001918460018302840111640100000000831117156103ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561040057600080fd5b82018360208201111561041257600080fd5b8035906020019184600183028401116401000000008311171561043457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061096d945050505050565b34801561048157600080fd5b506100ee610afa565b34801561049657600080fd5b506104b4600480360360208110156104ad57600080fd5b5035610b00565b60408051938452602084019290925282820152519081900360600190f35b3480156104de57600080fd5b506100ee610c51565b3480156104f357600080fd5b506104fc610c57565b60408051600160a060020a039092168252519081900360200190f35b34801561052457600080fd5b506102bf6004803603604081101561053b57600080fd5b508035906020013560ff16610c66565b34801561055757600080fd5b5061012a610f45565b34801561056c57600080fd5b506102bf6004803603602081101561058357600080fd5b5035610f7f565b34801561059657600080fd5b506105ba600480360360408110156105ad57600080fd5b508035906020013561135e565b6040518087815260200186600160a060020a0316600160a060020a03168152602001806020018060200185815260200184151515158152602001838103835287818151815260200191508051906020019080838360005b83811015610629578181015183820152602001610611565b50505050905090810190601f1680156106565780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b83811015610689578181015183820152602001610671565b50505050905090810190601f1680156106b65780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156106d657600080fd5b506105ba600480360360208110156106ed57600080fd5b5035611706565b34801561070057600080fd5b506104b46004803603602081101561071757600080fd5b5035600160a060020a031661173f565b34801561073357600080fd5b506104b4611772565b60015481565b60008060608060606000866002548110151561075d57600080fd5b60008881526004602081815260409283902060018082015460058301546002808501805489516101009682161596909602600019011691909104601f81018790048702850187019098528784528f97600160a060020a03909316969095600386019590910193919286918301828280156108185780601f106107ed57610100808354040283529160200191610818565b820191906000526020600020905b8154815290600101906020018083116107fb57829003601f168201915b5050865460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959950889450925084019050828280156108a65780601f1061087b576101008083540402835291602001916108a6565b820191906000526020600020905b81548152906001019060200180831161088957829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959850879450925084019050828280156109345780601f1061090957610100808354040283529160200191610934565b820191906000526020600020905b81548152906001019060200180831161091757829003601f168201915b505050505091509650965096509650965096505091939550919395565b600054600160a060020a0316331461096857600080fd5b600155565b826000815111151561097e57600080fd5b826000815111151561098f57600080fd5b82600081511115156109a057600080fd5b600280546040805160c0810182528281523360208083019182528284018c8152606084018c9052608084018b9052426103e80260a08501526000868152600483529490942083518155915160018301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905592518051949592949193610a32938501929101906118ce565b5060608201518051610a4e9160038401916020909101906118ce565b5060808201518051610a6a9160048401916020909101906118ce565b5060a09190910151600590910155600280546001908101909155600080549154604051600160a060020a039093169281156108fc0292818181858888f19350505050158015610abd573d6000803e3d6000fd5b506040805182815290517f6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d9181900360200190a150505050505050565b60025481565b6000806000610b0d61194c565b60008581526006602090815260408083208151606081018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b82821015610ba05760008481526020908190206040805160608101825260038602909201805483526001808201548486015260029091015460ff161515918301919091529083529092019101610b4f565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015610c2e5760008481526020908190206040805180820190915260028502909101805482526001810154919290919083019060ff166003811115610c1057fe5b6003811115610c1b57fe5b8152505081526020019060010190610bcd565b505050915250508051602082015151604090920151519097919650945092505050565b60035481565b600054600160a060020a031681565b600054600160a060020a03163314610c7d57600080fd5b610c8561196e565b600083815260046020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f81018690048602830186018752808352929593949386019391929091830182828015610d425780601f10610d1757610100808354040283529160200191610d42565b820191906000526020600020905b815481529060010190602001808311610d2557829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015610dd65780601f10610dab57610100808354040283529160200191610dd6565b820191906000526020600020905b815481529060010190602001808311610db957829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015610e6a5780601f10610e3f57610100808354040283529160200191610e6a565b820191906000526020600020905b815481529060010190602001808311610e4d57829003601f168201915b5050509183525050600591909101546020918201528101516040820151606083015192935090916000610e9e84848461179a565b9050600660008281526020019081526020016000206002016040805190810160405280898152602001886003811115610ed357fe5b905281546001818101808555600094855260209485902084516002909402019283559383015182820180549192909160ff191690836003811115610f1357fe5b02179055505050600160a060020a03909416600090815260076020526040902055505060038054600101905550505050565b60008060608060606000600254600354101515610f6157600080fd5b610f6c600354610742565b949b939a50919850965094509092509050565b600054600160a060020a03163314610f9657600080fd5b610f9e61196e565b600082815260046020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f8101869004860283018601875280835292959394938601939192909183018282801561105b5780601f106110305761010080835404028352916020019161105b565b820191906000526020600020905b81548152906001019060200180831161103e57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156110ef5780601f106110c4576101008083540402835291602001916110ef565b820191906000526020600020905b8154815290600101906020018083116110d257829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156111835780601f1061115857610100808354040283529160200191611183565b820191906000526020600020905b81548152906001019060200180831161116657829003601f168201915b50505091835250506005919091015460209182015281015160408201516060830151929350909160006111b784848461179a565b600081815260066020908152604080832081516060810183528b8152426103e8028185019081526001828501818152938101805480830182559088528688209351600390910290930192835590519082015590516002909101805460ff1916911515919091179055600160a060020a0388168352600782529182902083905590518551929350839260089287929182918401908083835b6020831061126d5780518252601f19909201916020918201910161124e565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842087519094889450925082918401908083835b602083106112c95780518252601f1990920191602091820191016112aa565b51815160001960209485036101000a8101918216911992909216179091529390910195865260408051968790038201872097909755600380546001908101909155600089815260068352889020015488875290920191850191909152505082517f32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae9799392819003909201919050a1505050505050565b60008060608060008061136f61194c565b60008981526006602090815260408083208151606081018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b828210156114025760008481526020908190206040805160608101825260038602909201805483526001808201548486015260029091015460ff1615159183019190915290835290920191016113b1565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156114905760008481526020908190206040805180820190915260028502909101805482526001810154919290919083019060ff16600381111561147257fe5b600381111561147d57fe5b815250508152602001906001019061142f565b5050509152505060208101515190915088106114ab57600080fd5b6114b36119ae565b602082015180518a9081106114c457fe5b602090810290910101518051985090506114dc61196e565b600089815260046020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f810186900486028301860187528083529295939493860193919290918301828280156115995780601f1061156e57610100808354040283529160200191611599565b820191906000526020600020905b81548152906001019060200180831161157c57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561162d5780601f106116025761010080835404028352916020019161162d565b820191906000526020600020905b81548152906001019060200180831161161057829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156116c15780601f10611696576101008083540402835291602001916116c1565b820191906000526020600020905b8154815290600101906020018083116116a457829003601f168201915b50505050508152602001600582015481525050905080602001519750806040015196508060600151955081602001519450816040015193505050509295509295509295565b33600090815260076020526040812054819060609081908390819061172b908861135e565b949c939b5091995097509550909350915050565b600160a060020a0381166000908152600760205260408120548190819061176590610b00565b9250925092509193909250565b336000908152600760205260408120548190819061178f90610b00565b925092509250909192565b60008054600160a060020a031633146117b257600080fd5b600160a060020a038416600090815260076020526040812054908111156117da5790506118c7565b6008846040518082805190602001908083835b6020831061180c5780518252601f1990920191602091820191016117ed565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842087519094889450925082918401908083835b602083106118685780518252601f199092019160209182019101611849565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220549250505060008111156118a85790506118c7565b5050600580546000818152600660205260409020819055600181019091555b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061190f57805160ff191683800117855561193c565b8280016001018555821561193c579182015b8281111561193c578251825591602001919060010190611921565b506119489291506119d2565b5090565b6060604051908101604052806000815260200160608152602001606081525090565b60c060405190810160405280600081526020016000600160a060020a03168152602001606081526020016060815260200160608152602001600081525090565b60606040519081016040528060008152602001600081526020016000151581525090565b6119ec91905b8082111561194857600081556001016119d8565b9056fea165627a7a72305820f458a6005d0633e1e1f760f89252f81c8bb5bec41dca2bef9fd1ff49db0c10590029`

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetIdentityInfo is a free data retrieval call binding the contract method 0x632859c3.
//
// Solidity: function getIdentityInfo(identityId uint256) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistryCaller) GetIdentityInfo(opts *bind.CallOpts, identityId *big.Int) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	ret := new(struct {
		Id             *big.Int
		ProviderCount  *big.Int
		RejectionCount *big.Int
	})
	out := ret
	err := _IdentityRegistry.contract.Call(opts, out, "getIdentityInfo", identityId)
	return *ret, err
}

// GetIdentityInfo is a free data retrieval call binding the contract method 0x632859c3.
//
// Solidity: function getIdentityInfo(identityId uint256) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistrySession) GetIdentityInfo(identityId *big.Int) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetIdentityInfo(&_IdentityRegistry.CallOpts, identityId)
}

// GetIdentityInfo is a free data retrieval call binding the contract method 0x632859c3.
//
// Solidity: function getIdentityInfo(identityId uint256) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetIdentityInfo(identityId *big.Int) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetIdentityInfo(&_IdentityRegistry.CallOpts, identityId)
}

// GetIdentityInfoByAddress is a free data retrieval call binding the contract method 0xe0774281.
//
// Solidity: function getIdentityInfoByAddress(identityAddress address) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistryCaller) GetIdentityInfoByAddress(opts *bind.CallOpts, identityAddress common.Address) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	ret := new(struct {
		Id             *big.Int
		ProviderCount  *big.Int
		RejectionCount *big.Int
	})
	out := ret
	err := _IdentityRegistry.contract.Call(opts, out, "getIdentityInfoByAddress", identityAddress)
	return *ret, err
}

// GetIdentityInfoByAddress is a free data retrieval call binding the contract method 0xe0774281.
//
// Solidity: function getIdentityInfoByAddress(identityAddress address) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistrySession) GetIdentityInfoByAddress(identityAddress common.Address) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetIdentityInfoByAddress(&_IdentityRegistry.CallOpts, identityAddress)
}

// GetIdentityInfoByAddress is a free data retrieval call binding the contract method 0xe0774281.
//
// Solidity: function getIdentityInfoByAddress(identityAddress address) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetIdentityInfoByAddress(identityAddress common.Address) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetIdentityInfoByAddress(&_IdentityRegistry.CallOpts, identityAddress)
}

// GetMyIdentityInfo is a free data retrieval call binding the contract method 0xf388ad23.
//
// Solidity: function getMyIdentityInfo() constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistryCaller) GetMyIdentityInfo(opts *bind.CallOpts) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	ret := new(struct {
		Id             *big.Int
		ProviderCount  *big.Int
		RejectionCount *big.Int
	})
	out := ret
	err := _IdentityRegistry.contract.Call(opts, out, "getMyIdentityInfo")
	return *ret, err
}

// GetMyIdentityInfo is a free data retrieval call binding the contract method 0xf388ad23.
//
// Solidity: function getMyIdentityInfo() constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistrySession) GetMyIdentityInfo() (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetMyIdentityInfo(&_IdentityRegistry.CallOpts)
}

// GetMyIdentityInfo is a free data retrieval call binding the contract method 0xf388ad23.
//
// Solidity: function getMyIdentityInfo() constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetMyIdentityInfo() (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetMyIdentityInfo(&_IdentityRegistry.CallOpts)
}

// GetMyProviderInfo is a free data retrieval call binding the contract method 0xd64c73a6.
//
// Solidity: function getMyProviderInfo(providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_IdentityRegistry *IdentityRegistryCaller) GetMyProviderInfo(opts *bind.CallOpts, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	ret := new(struct {
		RequestId       *big.Int
		IdentityAddress common.Address
		Provider        string
		UserName        string
		Timestamp       *big.Int
		Active          bool
	})
	out := ret
	err := _IdentityRegistry.contract.Call(opts, out, "getMyProviderInfo", providerId)
	return *ret, err
}

// GetMyProviderInfo is a free data retrieval call binding the contract method 0xd64c73a6.
//
// Solidity: function getMyProviderInfo(providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_IdentityRegistry *IdentityRegistrySession) GetMyProviderInfo(providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _IdentityRegistry.Contract.GetMyProviderInfo(&_IdentityRegistry.CallOpts, providerId)
}

// GetMyProviderInfo is a free data retrieval call binding the contract method 0xd64c73a6.
//
// Solidity: function getMyProviderInfo(providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetMyProviderInfo(providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _IdentityRegistry.Contract.GetMyProviderInfo(&_IdentityRegistry.CallOpts, providerId)
}

// GetNextRequest is a free data retrieval call binding the contract method 0x95dfb896.
//
// Solidity: function getNextRequest() constant returns(requestId uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_IdentityRegistry *IdentityRegistryCaller) GetNextRequest(opts *bind.CallOpts) (struct {
	RequestId  *big.Int
	Requestor  common.Address
	Provider   string
	UserName   string
	Identifier string
	Timestamp  *big.Int
}, error) {
	ret := new(struct {
		RequestId  *big.Int
		Requestor  common.Address
		Provider   string
		UserName   string
		Identifier string
		Timestamp  *big.Int
	})
	out := ret
	err := _IdentityRegistry.contract.Call(opts, out, "getNextRequest")
	return *ret, err
}

// GetNextRequest is a free data retrieval call binding the contract method 0x95dfb896.
//
// Solidity: function getNextRequest() constant returns(requestId uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_IdentityRegistry *IdentityRegistrySession) GetNextRequest() (struct {
	RequestId  *big.Int
	Requestor  common.Address
	Provider   string
	UserName   string
	Identifier string
	Timestamp  *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetNextRequest(&_IdentityRegistry.CallOpts)
}

// GetNextRequest is a free data retrieval call binding the contract method 0x95dfb896.
//
// Solidity: function getNextRequest() constant returns(requestId uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetNextRequest() (struct {
	RequestId  *big.Int
	Requestor  common.Address
	Provider   string
	UserName   string
	Identifier string
	Timestamp  *big.Int
}, error) {
	return _IdentityRegistry.Contract.GetNextRequest(&_IdentityRegistry.CallOpts)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(identityId uint256, providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_IdentityRegistry *IdentityRegistryCaller) GetProviderInfo(opts *bind.CallOpts, identityId *big.Int, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	ret := new(struct {
		RequestId       *big.Int
		IdentityAddress common.Address
		Provider        string
		UserName        string
		Timestamp       *big.Int
		Active          bool
	})
	out := ret
	err := _IdentityRegistry.contract.Call(opts, out, "getProviderInfo", identityId, providerId)
	return *ret, err
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(identityId uint256, providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_IdentityRegistry *IdentityRegistrySession) GetProviderInfo(identityId *big.Int, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _IdentityRegistry.Contract.GetProviderInfo(&_IdentityRegistry.CallOpts, identityId, providerId)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(identityId uint256, providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetProviderInfo(identityId *big.Int, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _IdentityRegistry.Contract.GetProviderInfo(&_IdentityRegistry.CallOpts, identityId, providerId)
}

// GetRequestById is a free data retrieval call binding the contract method 0x26757b73.
//
// Solidity: function getRequestById(requestId uint256) constant returns(uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_IdentityRegistry *IdentityRegistryCaller) GetRequestById(opts *bind.CallOpts, requestId *big.Int) (*big.Int, common.Address, string, string, string, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(string)
		ret3 = new(string)
		ret4 = new(string)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _IdentityRegistry.contract.Call(opts, out, "getRequestById", requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestById is a free data retrieval call binding the contract method 0x26757b73.
//
// Solidity: function getRequestById(requestId uint256) constant returns(uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_IdentityRegistry *IdentityRegistrySession) GetRequestById(requestId *big.Int) (*big.Int, common.Address, string, string, string, *big.Int, error) {
	return _IdentityRegistry.Contract.GetRequestById(&_IdentityRegistry.CallOpts, requestId)
}

// GetRequestById is a free data retrieval call binding the contract method 0x26757b73.
//
// Solidity: function getRequestById(requestId uint256) constant returns(uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetRequestById(requestId *big.Int) (*big.Int, common.Address, string, string, string, *big.Int, error) {
	return _IdentityRegistry.Contract.GetRequestById(&_IdentityRegistry.CallOpts, requestId)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) IdentityAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "identityAdmin")
	return *ret0, err
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) IdentityAdmin() (common.Address, error) {
	return _IdentityRegistry.Contract.IdentityAdmin(&_IdentityRegistry.CallOpts)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) IdentityAdmin() (common.Address, error) {
	return _IdentityRegistry.Contract.IdentityAdmin(&_IdentityRegistry.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) NextIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "nextIdentityRequestIndex")
	return *ret0, err
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) NextIdentityRequestIndex() (*big.Int, error) {
	return _IdentityRegistry.Contract.NextIdentityRequestIndex(&_IdentityRegistry.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _IdentityRegistry.Contract.NextIdentityRequestIndex(&_IdentityRegistry.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) PendingIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "pendingIdentityRequestIndex")
	return *ret0, err
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _IdentityRegistry.Contract.PendingIdentityRequestIndex(&_IdentityRegistry.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _IdentityRegistry.Contract.PendingIdentityRequestIndex(&_IdentityRegistry.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) RequestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "requestPrice")
	return *ret0, err
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) RequestPrice() (*big.Int, error) {
	return _IdentityRegistry.Contract.RequestPrice(&_IdentityRegistry.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) RequestPrice() (*big.Int, error) {
	return _IdentityRegistry.Contract.RequestPrice(&_IdentityRegistry.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Approve(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "approve", requestId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_IdentityRegistry *IdentityRegistrySession) Approve(requestId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Approve(&_IdentityRegistry.TransactOpts, requestId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Approve(requestId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Approve(&_IdentityRegistry.TransactOpts, requestId)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Reject(opts *bind.TransactOpts, requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "reject", requestId, reason)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_IdentityRegistry *IdentityRegistrySession) Reject(requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Reject(&_IdentityRegistry.TransactOpts, requestId, reason)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Reject(requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Reject(&_IdentityRegistry.TransactOpts, requestId, reason)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetRequestPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setRequestPrice", price)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetRequestPrice(price *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetRequestPrice(&_IdentityRegistry.TransactOpts, price)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetRequestPrice(price *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetRequestPrice(&_IdentityRegistry.TransactOpts, price)
}

// SubmitRequest is a paid mutator transaction binding the contract method 0x470c153b.
//
// Solidity: function submitRequest(provider string, userName string, identifier string) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SubmitRequest(opts *bind.TransactOpts, provider string, userName string, identifier string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "submitRequest", provider, userName, identifier)
}

// SubmitRequest is a paid mutator transaction binding the contract method 0x470c153b.
//
// Solidity: function submitRequest(provider string, userName string, identifier string) returns()
func (_IdentityRegistry *IdentityRegistrySession) SubmitRequest(provider string, userName string, identifier string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SubmitRequest(&_IdentityRegistry.TransactOpts, provider, userName, identifier)
}

// SubmitRequest is a paid mutator transaction binding the contract method 0x470c153b.
//
// Solidity: function submitRequest(provider string, userName string, identifier string) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SubmitRequest(provider string, userName string, identifier string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SubmitRequest(&_IdentityRegistry.TransactOpts, provider, userName, identifier)
}

// IdentityRegistryNewApprovedProviderEventIterator is returned from FilterNewApprovedProviderEvent and is used to iterate over the raw logs and unpacked data for NewApprovedProviderEvent events raised by the IdentityRegistry contract.
type IdentityRegistryNewApprovedProviderEventIterator struct {
	Event *IdentityRegistryNewApprovedProviderEvent // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryNewApprovedProviderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryNewApprovedProviderEvent)
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
		it.Event = new(IdentityRegistryNewApprovedProviderEvent)
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
func (it *IdentityRegistryNewApprovedProviderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryNewApprovedProviderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryNewApprovedProviderEvent represents a NewApprovedProviderEvent event raised by the IdentityRegistry contract.
type IdentityRegistryNewApprovedProviderEvent struct {
	IdentityId *big.Int
	ProviderId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewApprovedProviderEvent is a free log retrieval operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterNewApprovedProviderEvent(opts *bind.FilterOpts) (*IdentityRegistryNewApprovedProviderEventIterator, error) {

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryNewApprovedProviderEventIterator{contract: _IdentityRegistry.contract, event: "NewApprovedProviderEvent", logs: logs, sub: sub}, nil
}

// WatchNewApprovedProviderEvent is a free log subscription operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchNewApprovedProviderEvent(opts *bind.WatchOpts, sink chan<- *IdentityRegistryNewApprovedProviderEvent) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryNewApprovedProviderEvent)
				if err := _IdentityRegistry.contract.UnpackLog(event, "NewApprovedProviderEvent", log); err != nil {
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

// IdentityRegistryNewRequestEventIterator is returned from FilterNewRequestEvent and is used to iterate over the raw logs and unpacked data for NewRequestEvent events raised by the IdentityRegistry contract.
type IdentityRegistryNewRequestEventIterator struct {
	Event *IdentityRegistryNewRequestEvent // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryNewRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryNewRequestEvent)
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
		it.Event = new(IdentityRegistryNewRequestEvent)
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
func (it *IdentityRegistryNewRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryNewRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryNewRequestEvent represents a NewRequestEvent event raised by the IdentityRegistry contract.
type IdentityRegistryNewRequestEvent struct {
	IdentityRequestId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRequestEvent is a free log retrieval operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterNewRequestEvent(opts *bind.FilterOpts) (*IdentityRegistryNewRequestEventIterator, error) {

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryNewRequestEventIterator{contract: _IdentityRegistry.contract, event: "NewRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewRequestEvent is a free log subscription operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchNewRequestEvent(opts *bind.WatchOpts, sink chan<- *IdentityRegistryNewRequestEvent) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryNewRequestEvent)
				if err := _IdentityRegistry.contract.UnpackLog(event, "NewRequestEvent", log); err != nil {
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

// PermaChatABI is the input ABI used to generate the binding from.
const PermaChatABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addDownvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"pinPost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"replyId\",\"type\":\"uint256\"}],\"name\":\"getReplyParent\",\"outputs\":[{\"name\":\"parentPostId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestDatabaseIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setRequestPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"getPost\",\"outputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"poster\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"contentType\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextActionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"submitRequest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"setNewDB\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextCommentIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityId\",\"type\":\"uint256\"}],\"name\":\"getIdentityInfo\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"providerCount\",\"type\":\"uint256\"},{\"name\":\"rejectionCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingIdentityRequestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"hasCommented\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestDBHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identityAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"},{\"name\":\"commentaryIndex\",\"type\":\"uint256\"}],\"name\":\"getCommentary\",\"outputs\":[{\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"name\":\"commenter\",\"type\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPostIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"parentPostId\",\"type\":\"uint256\"},{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"newReply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextRequest\",\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"commentId\",\"type\":\"uint256\"}],\"name\":\"getCommentary\",\"outputs\":[{\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"name\":\"commenter\",\"type\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"currentDatabaseIndex\",\"type\":\"uint256\"}],\"name\":\"getDB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"newPost\",\"outputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPinnedPost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addFlag\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"getReplies\",\"outputs\":[{\"name\":\"replyIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"getCommentaryCount\",\"outputs\":[{\"name\":\"commentaryCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"name\":\"targetId\",\"type\":\"uint256\"},{\"name\":\"targetType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityId\",\"type\":\"uint256\"},{\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"getProviderInfo\",\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"identityAddress\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestDBIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"getMyProviderInfo\",\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"identityAddress\",\"type\":\"address\"},{\"name\":\"provider\",\"type\":\"string\"},{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"getIdentityInfoByAddress\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"providerCount\",\"type\":\"uint256\"},{\"name\":\"rejectionCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"addUpvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMyIdentityInfo\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"providerCount\",\"type\":\"uint256\"},{\"name\":\"rejectionCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentDatabaseIndex\",\"type\":\"uint256\"}],\"name\":\"DBUpdateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityRequestId\",\"type\":\"uint256\"}],\"name\":\"NewRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"identityId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NewApprovedProviderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewPostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"parentPostId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"poster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contentType\",\"type\":\"string\"}],\"name\":\"NewReplyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commentaryType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"commenter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"NewCommentaryEvent\",\"type\":\"event\"}]"

// PermaChatBin is the compiled bytecode used for deploying new contracts.
const PermaChatBin = `0x60806040526001600081815560028290556009829055600d829055600e829055601091909155601555600b8054600160a060020a03199081163390811790925560148054909116909117905561335a8061005a6000396000f3fe6080604052600436106101e7577c0100000000000000000000000000000000000000000000000000000000600035046301c6090781146101ec5780631604f9ea14610218578063251cf4ea1461023f5780632637adbf1461026957806326757b731461029357806339fa083d146104285780633fcf7ca11461043d57806340731c241461046757806346abe73a14610596578063470c153b146105ab578063474c49dd1461075f5780634d625d991461081257806361165d0814610827578063632859c31461083c57806375efbecf146108845780637af2d1af146108995780637bda06c2146108d75780637e889537146109615780638df0c08f146109925780638e26c7e614610a025780638ef5d2b914610a1f5780638faa6dad14610a345780639119c8d714610b7557806395dfb89614610ba8578063980d9ada14610bbd57806398c01f2514610be75780639e28cf9614610c11578063a46325df14610d4b578063ae2cded614610d60578063ae58449d14610d8a578063b65167b214610e04578063b6e7687314610e2e578063b759f95414610e83578063ba438b3b14610ead578063ca4454e514610fed578063d64c73a614611002578063e07742811461102c578063f1de9ea41461105f578063f388ad2314611089578063fbd5b4e11461109e575b600080fd5b3480156101f857600080fd5b506102166004803603602081101561020f57600080fd5b50356110b3565b005b34801561022457600080fd5b5061022d6110c3565b60408051918252519081900360200190f35b34801561024b57600080fd5b506102166004803603602081101561026257600080fd5b50356110c9565b34801561027557600080fd5b5061022d6004803603602081101561028c57600080fd5b503561113c565b34801561029f57600080fd5b506102bd600480360360208110156102b657600080fd5b5035611163565b6040518087815260200186600160a060020a0316600160a060020a03168152602001806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b8381101561032657818101518382015260200161030e565b50505050905090810190601f1680156103535780820380516001836020036101000a031916815260200191505b50848103835287518152875160209182019189019080838360005b8381101561038657818101518382015260200161036e565b50505050905090810190601f1680156103b35780820380516001836020036101000a031916815260200191505b50848103825286518152865160209182019188019080838360005b838110156103e65781810151838201526020016103ce565b50505050905090810190601f1680156104135780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34801561043457600080fd5b5061022d611375565b34801561044957600080fd5b506102166004803603602081101561046057600080fd5b503561137b565b34801561047357600080fd5b506104916004803603602081101561048a57600080fd5b5035611397565b604051808060200186600160a060020a0316600160a060020a0316815260200185815260200184815260200180602001838103835288818151815260200191508051906020019080838360005b838110156104f65781810151838201526020016104de565b50505050905090810190601f1680156105235780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561055657818101518382015260200161053e565b50505050905090810190601f1680156105835780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156105a257600080fd5b5061022d61152d565b610216600480360360608110156105c157600080fd5b8101906020810181356401000000008111156105dc57600080fd5b8201836020820111156105ee57600080fd5b8035906020019184600183028401116401000000008311171561061057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561066357600080fd5b82018360208201111561067557600080fd5b8035906020019184600183028401116401000000008311171561069757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156106ea57600080fd5b8201836020820111156106fc57600080fd5b8035906020019184600183028401116401000000008311171561071e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611533945050505050565b34801561076b57600080fd5b506102166004803603602081101561078257600080fd5b81019060208101813564010000000081111561079d57600080fd5b8201836020820111156107af57600080fd5b803590602001918460018302840111640100000000831117156107d157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116be945050505050565b34801561081e57600080fd5b5061022d611738565b34801561083357600080fd5b5061022d61173e565b34801561084857600080fd5b506108666004803603602081101561085f57600080fd5b5035611744565b60408051938452602084019290925282820152519081900360600190f35b34801561089057600080fd5b5061022d611895565b3480156108a557600080fd5b506108c3600480360360208110156108bc57600080fd5b503561189b565b604080519115158252519081900360200190f35b3480156108e357600080fd5b506108ec6118e1565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561092657818101518382015260200161090e565b50505050905090810190601f1680156109535780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561096d57600080fd5b50610976611984565b60408051600160a060020a039092168252519081900360200190f35b34801561099e57600080fd5b506109c2600480360360408110156109b557600080fd5b5080359060200135611993565b604051808560038111156109d257fe5b60ff168152600160a060020a03909416602085015250604080840192909252606083015251908190036080019150f35b61021660048036036020811015610a1857600080fd5b5035611a6f565b348015610a2b57600080fd5b5061022d611ac7565b348015610a4057600080fd5b5061022d60048036036060811015610a5757600080fd5b81359190810190604081016020820135640100000000811115610a7957600080fd5b820183602082011115610a8b57600080fd5b80359060200191846001830284011164010000000083111715610aad57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050640100000000811115610b0057600080fd5b820183602082011115610b1257600080fd5b80359060200191846001830284011164010000000083111715610b3457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611acd945050505050565b348015610b8157600080fd5b5061021660048036036040811015610b9857600080fd5b508035906020013560ff16611d4e565b348015610bb457600080fd5b506102bd61202d565b348015610bc957600080fd5b506109c260048036036020811015610be057600080fd5b5035612067565b348015610bf357600080fd5b506108ec60048036036020811015610c0a57600080fd5b50356120bf565b348015610c1d57600080fd5b5061022d60048036036040811015610c3457600080fd5b810190602081018135640100000000811115610c4f57600080fd5b820183602082011115610c6157600080fd5b80359060200191846001830284011164010000000083111715610c8357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050640100000000811115610cd657600080fd5b820183602082011115610ce857600080fd5b80359060200191846001830284011164010000000083111715610d0a57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061217e945050505050565b348015610d5757600080fd5b5061022d6123f4565b348015610d6c57600080fd5b5061021660048036036020811015610d8357600080fd5b5035612407565b348015610d9657600080fd5b50610db460048036036020811015610dad57600080fd5b5035612414565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610df0578181015183820152602001610dd8565b505050509050019250505060405180910390f35b348015610e1057600080fd5b5061022d60048036036020811015610e2757600080fd5b5035612488565b348015610e3a57600080fd5b50610e5860048036036020811015610e5157600080fd5b50356124af565b60405180838152602001826001811115610e6e57fe5b60ff1681526020019250505060405180910390f35b348015610e8f57600080fd5b5061021660048036036020811015610ea657600080fd5b50356124cf565b348015610eb957600080fd5b50610edd60048036036040811015610ed057600080fd5b50803590602001356128ae565b6040518087815260200186600160a060020a0316600160a060020a03168152602001806020018060200185815260200184151515158152602001838103835287818151815260200191508051906020019080838360005b83811015610f4c578181015183820152602001610f34565b50505050905090810190601f168015610f795780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b83811015610fac578181015183820152602001610f94565b50505050905090810190601f168015610fd95780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610ff957600080fd5b5061022d612c56565b34801561100e57600080fd5b50610edd6004803603602081101561102557600080fd5b5035612c5c565b34801561103857600080fd5b506108666004803603602081101561104f57600080fd5b5035600160a060020a0316612c95565b34801561106b57600080fd5b506102166004803603602081101561108257600080fd5b5035612cc8565b34801561109557600080fd5b50610866612cd5565b3480156110aa57600080fd5b50610216612cfd565b6110c08160026000612d50565b50565b600c5481565b600254819081106110d957600080fd5b33600090815260126020908152604080832054835260119091528120600101541161110357600080fd5b600082815260036020526040902060020154600160a060020a0316331461112957600080fd5b5033600090815260066020526040902055565b6000816002548110151561114f57600080fd5b505060009081526005602052604090205490565b6000806060806060600086600d548110151561117e57600080fd5b6000888152600f602090815260409182902060018082015460058301546002808501805488516101009682161596909602600019011691909104601f81018790048702850187019097528684528e96600160a060020a039093169590946003810194600490910193909186919083018282801561123c5780601f106112115761010080835404028352916020019161123c565b820191906000526020600020905b81548152906001019060200180831161121f57829003601f168201915b5050865460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959950889450925084019050828280156112ca5780601f1061129f576101008083540402835291602001916112ca565b820191906000526020600020905b8154815290600101906020018083116112ad57829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959850879450925084019050828280156113585780601f1061132d57610100808354040283529160200191611358565b820191906000526020600020905b81548152906001019060200180831161133b57829003601f168201915b505050505091509650965096509650965096505091939550919395565b60155481565b600b54600160a060020a0316331461139257600080fd5b600c55565b60606000806000606085600254811015156113b157600080fd5b600360008881526020019081526020016000206001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561145b5780601f106114305761010080835404028352916020019161145b565b820191906000526020600020905b81548152906001019060200180831161143e57829003601f168201915b50505060008a8152600360208181526040928390206002808201549382015460048301546005909301805487516101006001831615026000190190911693909304601f8101869004860284018601909752868352989e50600160a060020a039094169c50929a5098509094935090915083018282801561151c5780601f106114f15761010080835404028352916020019161151c565b820191906000526020600020905b8154815290600101906020018083116114ff57829003601f168201915b505050505091505091939590929450565b60005481565b826000815111151561154457600080fd5b826000815111151561155557600080fd5b826000815111151561156657600080fd5b600d546040805160c0810182528281523360208083019182528284018b8152606084018b9052608084018a9052426103e80260a08501526000868152600f83529490942083518155915160018301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905592518051929391926115f692600285019201906131de565b50606082015180516116129160038401916020909101906131de565b506080820151805161162e9160048401916020909101906131de565b5060a09190910151600590910155600d80546001019055600b54600c54604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015611681573d6000803e3d6000fd5b506040805182815290517f6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d9181900360200190a150505050505050565b601454600160a060020a031633146116d557600080fd5b6015805460010190819055600090815260166020908152604090912082516116ff928401906131de565b5060155460408051918252517f50634acbe4d09fc2a5c9ee5d100551a5a9e0162ed4bf7cb06c3115a7cfb25ed99181900360200190a150565b600d5481565b60095481565b600080600061175161325c565b60008581526011602090815260408083208151606081018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b828210156117e45760008481526020908190206040805160608101825260038602909201805483526001808201548486015260029091015460ff161515918301919091529083529092019101611793565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156118725760008481526020908190206040805180820190915260028502909101805482526001810154919290919083019060ff16600381111561185457fe5b600381111561185f57fe5b8152505081526020019060010190611811565b505050915250508051602082015151604090920151519097919650945092505050565b600e5481565b600081600254811015156118ae57600080fd5b6000838152600860205260408120906118c63361300d565b815260208101919091526040016000205460ff169392505050565b60155460009081526016602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156119795780601f1061194e57610100808354040283529160200191611979565b820191906000526020600020905b81548152906001019060200180831161195c57829003601f168201915b505050505090505b90565b600b54600160a060020a031681565b60008060008085600254811015156119aa57600080fd5b6000878152600760205260409020548790879081116119c857600080fd5b60008981526007602052604090208054611a5c91908a9081106119e757fe5b9060005260206000209060030201608060405190810160405290816000820160009054906101000a900460ff166003811115611a1f57fe5b6003811115611a2a57fe5b815281546101009004600160a060020a0316602082015260018201546040820152600290910154606090910152613028565b929c919b50995090975095505050505050565b611a7b81600034612d50565b600081815260036020526040808220600201549051600160a060020a03909116913480156108fc02929091818181858888f19350505050158015611ac3573d6000803e3d6000fd5b5050565b60025481565b33600090815260126020908152604080832054835260119091528120600101548110611af857600080fd5b60025484908110611b0857600080fd5b6000611b14858561217e565b9050600460008781526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190915055508560056000838152602001908152602001600020819055507f378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3866003600084815260200190815260200160002060000154600360008581526020019081526020016000206001016003600086815260200190815260200160002060020160009054906101000a9004600160a060020a0316600360008781526020019081526020016000206003015460036000888152602001908152602001600020600401548a604051808881526020018781526020018060200186600160a060020a0316600160a060020a0316815260200185815260200184815260200180602001838103835288818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611cd05780601f10611ca557610100808354040283529160200191611cd0565b820191906000526020600020905b815481529060010190602001808311611cb357829003601f168201915b5050838103825284518152845160209182019186019080838360005b83811015611d04578181015183820152602001611cec565b50505050905090810190601f168015611d315780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a195945050505050565b600b54600160a060020a03163314611d6557600080fd5b611d6d61327e565b6000838152600f6020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f81018690048602830186018752808352929593949386019391929091830182828015611e2a5780601f10611dff57610100808354040283529160200191611e2a565b820191906000526020600020905b815481529060010190602001808311611e0d57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015611ebe5780601f10611e9357610100808354040283529160200191611ebe565b820191906000526020600020905b815481529060010190602001808311611ea157829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015611f525780601f10611f2757610100808354040283529160200191611f52565b820191906000526020600020905b815481529060010190602001808311611f3557829003601f168201915b5050509183525050600591909101546020918201528101516040820151606083015192935090916000611f86848484613042565b9050601160008281526020019081526020016000206002016040805190810160405280898152602001886003811115611fbb57fe5b905281546001818101808555600094855260209485902084516002909402019283559383015182820180549192909160ff191690836003811115611ffb57fe5b02179055505050600160a060020a039094166000908152601260205260409020555050600e8054600101905550505050565b60008060608060606000600d54600e5410151561204957600080fd5b612054600e54611163565b949b939a50919850965094509092509050565b600080600080846009548110151561207e57600080fd5b6000868152600a60205260409081902081516080810190925280546120af929190829060ff166003811115611a1f57fe5b9450945094509450509193509193565b6060600082101580156120d457506015548211155b15156120df57600080fd5b60008281526016602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156121725780601f1061214757610100808354040283529160200191612172565b820191906000526020600020905b81548152906001019060200180831161215557829003601f168201915b50505050509050919050565b336000908152601260209081526040808320548352601190915281206001015481106121a957600080fd5b6002546040805160c081018252828152602080820187815233838501524360608401526103e84202608084015260a0830187905260008581526003835293909320825181559251805192939261220592600185019201906131de565b50604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055606082015160038201556080820151600482015560a082015180516122689160058401916020909101906131de565b50506002805460010190555061227f816000613178565b60008181526003602081815260409283902080546002808301549483015460048401548751848152600160a060020a03909716978701889052606087018290526080870181905260c095870186815260019586018054600019978116156101000297909701909616939093049587018690527fab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac275497939694959394919390928b929160a083019060e0840190899080156123795780601f1061234e57610100808354040283529160200191612379565b820191906000526020600020905b81548152906001019060200180831161235c57829003601f168201915b5050838103825284518152845160209182019186019080838360005b838110156123ad578181015183820152602001612395565b50505050905090810190601f1680156123da5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a19392505050565b3360009081526006602052604090205490565b6110c08160036000612d50565b6060816002548110151561242757600080fd5b6000838152600460209081526040918290208054835181840281018401909452808452909183018282801561247b57602002820191906000526020600020905b815481526020019060010190808311612467575b5050505050915050919050565b6000816002548110151561249b57600080fd5b505060009081526007602052604090205490565b60009081526001602081905260409091208054910154909160ff90911690565b600b54600160a060020a031633146124e657600080fd5b6124ee61327e565b6000828152600f6020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f810186900486028301860187528083529295939493860193919290918301828280156125ab5780601f10612580576101008083540402835291602001916125ab565b820191906000526020600020905b81548152906001019060200180831161258e57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561263f5780601f106126145761010080835404028352916020019161263f565b820191906000526020600020905b81548152906001019060200180831161262257829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156126d35780601f106126a8576101008083540402835291602001916126d3565b820191906000526020600020905b8154815290600101906020018083116126b657829003601f168201915b5050509183525050600591909101546020918201528101516040820151606083015192935090916000612707848484613042565b600081815260116020908152604080832081516060810183528b8152426103e8028185019081526001828501818152938101805480830182559088528688209351600390910290930192835590519082015590516002909101805460ff1916911515919091179055600160a060020a0388168352601282529182902083905590518551929350839260139287929182918401908083835b602083106127bd5780518252601f19909201916020918201910161279e565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842087519094889450925082918401908083835b602083106128195780518252601f1990920191602091820191016127fa565b51815160001960209485036101000a8101918216911992909216179091529390910195865260408051968790038201872097909755600e80546001908101909155600089815260118352889020015488875290920191850191909152505082517f32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae9799392819003909201919050a1505050505050565b6000806060806000806128bf61325c565b60008981526011602090815260408083208151606081018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b828210156129525760008481526020908190206040805160608101825260038602909201805483526001808201548486015260029091015460ff161515918301919091529083529092019101612901565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156129e05760008481526020908190206040805180820190915260028502909101805482526001810154919290919083019060ff1660038111156129c257fe5b60038111156129cd57fe5b815250508152602001906001019061297f565b5050509152505060208101515190915088106129fb57600080fd5b612a036132be565b602082015180518a908110612a1457fe5b60209081029091010151805198509050612a2c61327e565b6000898152600f6020908152604091829020825160c08101845281548152600180830154600160a060020a0316828501526002808401805487516101009482161594909402600019011691909104601f81018690048602830186018752808352929593949386019391929091830182828015612ae95780601f10612abe57610100808354040283529160200191612ae9565b820191906000526020600020905b815481529060010190602001808311612acc57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015612b7d5780601f10612b5257610100808354040283529160200191612b7d565b820191906000526020600020905b815481529060010190602001808311612b6057829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015612c115780601f10612be657610100808354040283529160200191612c11565b820191906000526020600020905b815481529060010190602001808311612bf457829003601f168201915b50505050508152602001600582015481525050905080602001519750806040015196508060600151955081602001519450816040015193505050509295509295509295565b60155490565b336000908152601260205260408120548190606090819083908190612c8190886128ae565b949c939b5091995097509550909350915050565b600160a060020a03811660009081526012602052604081205481908190612cbb90611744565b9250925092509193909250565b6110c08160016000612d50565b3360009081526012602052604081205481908190612cf290611744565b925092509250909192565b336000908152601260209081526040808320548352601190915281206001015411612d2757600080fd5b336000908152600660205260408120541115612d4e57336000908152600660205260408120555b565b336000908152601260209081526040808320548352601190915281206001015411612d7a57600080fd5b60025483908110612d8a57600080fd5b60008481526008602052604081208591612da33361300d565b815260208101919091526040016000205460ff1615612dc157600080fd5b6009546000868152600760209081526040808320546008909252822090916103e8420291600191612df13361300d565b81526020810191909152604001600020805460ff1916911515919091179055612e186132e2565b608060405190810160405280896003811115612e3057fe5b815233602080830191909152604080830186905260609092018a905260008c81526007825291822080546001818101808455928552929093208451600394850290910180549596509194869492939192849260ff1990921691908490811115612e9557fe5b02179055506020828101518254600160a060020a039091166101000274ffffffffffffffffffffffffffffffffffffffff00199091161782556040808401516001808501919091556060909401516002909301929092556000888152600a909152208351815485945091929091839160ff1990911690836003811115612f1757fe5b021790555060208201518154600160a060020a039091166101000274ffffffffffffffffffffffffffffffffffffffff00199091161781556040820151600180830191909155606090920151600290910155612f74908590613178565b60016009600082825401925050819055507fa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc158984868b33878d60405180888152602001878152602001868152602001856003811115612fcf57fe5b60ff168152600160a060020a0390941660208501525060408084019290925260608301525190819003608001945092505050a1505050505050505050565b600160a060020a031660009081526012602052604090205490565b805160208201516040830151606090930151919390929190565b600b54600090600160a060020a0316331461305c57600080fd5b600160a060020a03841660009081526012602052604081205490811115613084579050613171565b6013846040518082805190602001908083835b602083106130b65780518252601f199092019160209182019101613097565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842087519094889450925082918401908083835b602083106131125780518252601f1990920191602091820191016130f3565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054925050506000811115613152579050613171565b5050601080546000818152601160205260409020819055600181019091555b9392505050565b604080519081016040528083815260200182600181111561319557fe5b9052600080548152600160208181526040909220835181559183015182820180549192909160ff19169083818111156131ca57fe5b021790555050600080546001019055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061321f57805160ff191683800117855561324c565b8280016001018555821561324c579182015b8281111561324c578251825591602001919060010190613231565b50613258929150613314565b5090565b6060604051908101604052806000815260200160608152602001606081525090565b60c060405190810160405280600081526020016000600160a060020a03168152602001606081526020016060815260200160608152602001600081525090565b60606040519081016040528060008152602001600081526020016000151581525090565b604080516080810190915280600081526020016000600160a060020a0316815260200160008152602001600081525090565b61198191905b80821115613258576000815560010161331a56fea165627a7a7230582099d4ed60091ec5f2b362d777a59927f8b12a9b64cc569bb9047d4de25ec11b0d0029`

// DeployPermaChat deploys a new Ethereum contract, binding an instance of PermaChat to it.
func DeployPermaChat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PermaChat, error) {
	parsed, err := abi.JSON(strings.NewReader(PermaChatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PermaChatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermaChat{PermaChatCaller: PermaChatCaller{contract: contract}, PermaChatTransactor: PermaChatTransactor{contract: contract}, PermaChatFilterer: PermaChatFilterer{contract: contract}}, nil
}

// PermaChat is an auto generated Go binding around an Ethereum contract.
type PermaChat struct {
	PermaChatCaller     // Read-only binding to the contract
	PermaChatTransactor // Write-only binding to the contract
	PermaChatFilterer   // Log filterer for contract events
}

// PermaChatCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermaChatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermaChatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermaChatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermaChatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermaChatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermaChatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermaChatSession struct {
	Contract     *PermaChat        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermaChatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermaChatCallerSession struct {
	Contract *PermaChatCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PermaChatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermaChatTransactorSession struct {
	Contract     *PermaChatTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PermaChatRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermaChatRaw struct {
	Contract *PermaChat // Generic contract binding to access the raw methods on
}

// PermaChatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermaChatCallerRaw struct {
	Contract *PermaChatCaller // Generic read-only contract binding to access the raw methods on
}

// PermaChatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermaChatTransactorRaw struct {
	Contract *PermaChatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermaChat creates a new instance of PermaChat, bound to a specific deployed contract.
func NewPermaChat(address common.Address, backend bind.ContractBackend) (*PermaChat, error) {
	contract, err := bindPermaChat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermaChat{PermaChatCaller: PermaChatCaller{contract: contract}, PermaChatTransactor: PermaChatTransactor{contract: contract}, PermaChatFilterer: PermaChatFilterer{contract: contract}}, nil
}

// NewPermaChatCaller creates a new read-only instance of PermaChat, bound to a specific deployed contract.
func NewPermaChatCaller(address common.Address, caller bind.ContractCaller) (*PermaChatCaller, error) {
	contract, err := bindPermaChat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermaChatCaller{contract: contract}, nil
}

// NewPermaChatTransactor creates a new write-only instance of PermaChat, bound to a specific deployed contract.
func NewPermaChatTransactor(address common.Address, transactor bind.ContractTransactor) (*PermaChatTransactor, error) {
	contract, err := bindPermaChat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermaChatTransactor{contract: contract}, nil
}

// NewPermaChatFilterer creates a new log filterer instance of PermaChat, bound to a specific deployed contract.
func NewPermaChatFilterer(address common.Address, filterer bind.ContractFilterer) (*PermaChatFilterer, error) {
	contract, err := bindPermaChat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermaChatFilterer{contract: contract}, nil
}

// bindPermaChat binds a generic wrapper to an already deployed contract.
func bindPermaChat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermaChatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermaChat *PermaChatRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermaChat.Contract.PermaChatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermaChat *PermaChatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermaChat.Contract.PermaChatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermaChat *PermaChatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermaChat.Contract.PermaChatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermaChat *PermaChatCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermaChat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermaChat *PermaChatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermaChat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermaChat *PermaChatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermaChat.Contract.contract.Transact(opts, method, params...)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_PermaChat *PermaChatCaller) GetAction(opts *bind.CallOpts, actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	ret := new(struct {
		TargetId   *big.Int
		TargetType uint8
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getAction", actionId)
	return *ret, err
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_PermaChat *PermaChatSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _PermaChat.Contract.GetAction(&_PermaChat.CallOpts, actionId)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(actionId uint256) constant returns(targetId uint256, targetType uint8)
func (_PermaChat *PermaChatCallerSession) GetAction(actionId *big.Int) (struct {
	TargetId   *big.Int
	TargetType uint8
}, error) {
	return _PermaChat.Contract.GetAction(&_PermaChat.CallOpts, actionId)
}

// GetCommentary is a free data retrieval call binding the contract method 0x980d9ada.
//
// Solidity: function getCommentary(commentId uint256) constant returns(commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_PermaChat *PermaChatCaller) GetCommentary(opts *bind.CallOpts, commentId *big.Int) (struct {
	CommentaryType uint8
	Commenter      common.Address
	Timestamp      *big.Int
	Tip            *big.Int
}, error) {
	ret := new(struct {
		CommentaryType uint8
		Commenter      common.Address
		Timestamp      *big.Int
		Tip            *big.Int
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getCommentary", commentId)
	return *ret, err
}

// GetCommentary is a free data retrieval call binding the contract method 0x980d9ada.
//
// Solidity: function getCommentary(commentId uint256) constant returns(commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_PermaChat *PermaChatSession) GetCommentary(commentId *big.Int) (struct {
	CommentaryType uint8
	Commenter      common.Address
	Timestamp      *big.Int
	Tip            *big.Int
}, error) {
	return _PermaChat.Contract.GetCommentary(&_PermaChat.CallOpts, commentId)
}

// GetCommentary is a free data retrieval call binding the contract method 0x980d9ada.
//
// Solidity: function getCommentary(commentId uint256) constant returns(commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_PermaChat *PermaChatCallerSession) GetCommentary(commentId *big.Int) (struct {
	CommentaryType uint8
	Commenter      common.Address
	Timestamp      *big.Int
	Tip            *big.Int
}, error) {
	return _PermaChat.Contract.GetCommentary(&_PermaChat.CallOpts, commentId)
}

// GetCommentaryCount is a free data retrieval call binding the contract method 0xb65167b2.
//
// Solidity: function getCommentaryCount(postId uint256) constant returns(commentaryCount uint256)
func (_PermaChat *PermaChatCaller) GetCommentaryCount(opts *bind.CallOpts, postId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getCommentaryCount", postId)
	return *ret0, err
}

// GetCommentaryCount is a free data retrieval call binding the contract method 0xb65167b2.
//
// Solidity: function getCommentaryCount(postId uint256) constant returns(commentaryCount uint256)
func (_PermaChat *PermaChatSession) GetCommentaryCount(postId *big.Int) (*big.Int, error) {
	return _PermaChat.Contract.GetCommentaryCount(&_PermaChat.CallOpts, postId)
}

// GetCommentaryCount is a free data retrieval call binding the contract method 0xb65167b2.
//
// Solidity: function getCommentaryCount(postId uint256) constant returns(commentaryCount uint256)
func (_PermaChat *PermaChatCallerSession) GetCommentaryCount(postId *big.Int) (*big.Int, error) {
	return _PermaChat.Contract.GetCommentaryCount(&_PermaChat.CallOpts, postId)
}

// GetDB is a free data retrieval call binding the contract method 0x98c01f25.
//
// Solidity: function getDB(currentDatabaseIndex uint256) constant returns(string)
func (_PermaChat *PermaChatCaller) GetDB(opts *bind.CallOpts, currentDatabaseIndex *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getDB", currentDatabaseIndex)
	return *ret0, err
}

// GetDB is a free data retrieval call binding the contract method 0x98c01f25.
//
// Solidity: function getDB(currentDatabaseIndex uint256) constant returns(string)
func (_PermaChat *PermaChatSession) GetDB(currentDatabaseIndex *big.Int) (string, error) {
	return _PermaChat.Contract.GetDB(&_PermaChat.CallOpts, currentDatabaseIndex)
}

// GetDB is a free data retrieval call binding the contract method 0x98c01f25.
//
// Solidity: function getDB(currentDatabaseIndex uint256) constant returns(string)
func (_PermaChat *PermaChatCallerSession) GetDB(currentDatabaseIndex *big.Int) (string, error) {
	return _PermaChat.Contract.GetDB(&_PermaChat.CallOpts, currentDatabaseIndex)
}

// GetIdentityInfo is a free data retrieval call binding the contract method 0x632859c3.
//
// Solidity: function getIdentityInfo(identityId uint256) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatCaller) GetIdentityInfo(opts *bind.CallOpts, identityId *big.Int) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	ret := new(struct {
		Id             *big.Int
		ProviderCount  *big.Int
		RejectionCount *big.Int
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getIdentityInfo", identityId)
	return *ret, err
}

// GetIdentityInfo is a free data retrieval call binding the contract method 0x632859c3.
//
// Solidity: function getIdentityInfo(identityId uint256) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatSession) GetIdentityInfo(identityId *big.Int) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _PermaChat.Contract.GetIdentityInfo(&_PermaChat.CallOpts, identityId)
}

// GetIdentityInfo is a free data retrieval call binding the contract method 0x632859c3.
//
// Solidity: function getIdentityInfo(identityId uint256) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatCallerSession) GetIdentityInfo(identityId *big.Int) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _PermaChat.Contract.GetIdentityInfo(&_PermaChat.CallOpts, identityId)
}

// GetIdentityInfoByAddress is a free data retrieval call binding the contract method 0xe0774281.
//
// Solidity: function getIdentityInfoByAddress(identityAddress address) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatCaller) GetIdentityInfoByAddress(opts *bind.CallOpts, identityAddress common.Address) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	ret := new(struct {
		Id             *big.Int
		ProviderCount  *big.Int
		RejectionCount *big.Int
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getIdentityInfoByAddress", identityAddress)
	return *ret, err
}

// GetIdentityInfoByAddress is a free data retrieval call binding the contract method 0xe0774281.
//
// Solidity: function getIdentityInfoByAddress(identityAddress address) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatSession) GetIdentityInfoByAddress(identityAddress common.Address) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _PermaChat.Contract.GetIdentityInfoByAddress(&_PermaChat.CallOpts, identityAddress)
}

// GetIdentityInfoByAddress is a free data retrieval call binding the contract method 0xe0774281.
//
// Solidity: function getIdentityInfoByAddress(identityAddress address) constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatCallerSession) GetIdentityInfoByAddress(identityAddress common.Address) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _PermaChat.Contract.GetIdentityInfoByAddress(&_PermaChat.CallOpts, identityAddress)
}

// GetLatestDBHash is a free data retrieval call binding the contract method 0x7bda06c2.
//
// Solidity: function getLatestDBHash() constant returns(string)
func (_PermaChat *PermaChatCaller) GetLatestDBHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getLatestDBHash")
	return *ret0, err
}

// GetLatestDBHash is a free data retrieval call binding the contract method 0x7bda06c2.
//
// Solidity: function getLatestDBHash() constant returns(string)
func (_PermaChat *PermaChatSession) GetLatestDBHash() (string, error) {
	return _PermaChat.Contract.GetLatestDBHash(&_PermaChat.CallOpts)
}

// GetLatestDBHash is a free data retrieval call binding the contract method 0x7bda06c2.
//
// Solidity: function getLatestDBHash() constant returns(string)
func (_PermaChat *PermaChatCallerSession) GetLatestDBHash() (string, error) {
	return _PermaChat.Contract.GetLatestDBHash(&_PermaChat.CallOpts)
}

// GetLatestDBIndex is a free data retrieval call binding the contract method 0xca4454e5.
//
// Solidity: function getLatestDBIndex() constant returns(uint256)
func (_PermaChat *PermaChatCaller) GetLatestDBIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getLatestDBIndex")
	return *ret0, err
}

// GetLatestDBIndex is a free data retrieval call binding the contract method 0xca4454e5.
//
// Solidity: function getLatestDBIndex() constant returns(uint256)
func (_PermaChat *PermaChatSession) GetLatestDBIndex() (*big.Int, error) {
	return _PermaChat.Contract.GetLatestDBIndex(&_PermaChat.CallOpts)
}

// GetLatestDBIndex is a free data retrieval call binding the contract method 0xca4454e5.
//
// Solidity: function getLatestDBIndex() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) GetLatestDBIndex() (*big.Int, error) {
	return _PermaChat.Contract.GetLatestDBIndex(&_PermaChat.CallOpts)
}

// GetMyIdentityInfo is a free data retrieval call binding the contract method 0xf388ad23.
//
// Solidity: function getMyIdentityInfo() constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatCaller) GetMyIdentityInfo(opts *bind.CallOpts) (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	ret := new(struct {
		Id             *big.Int
		ProviderCount  *big.Int
		RejectionCount *big.Int
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getMyIdentityInfo")
	return *ret, err
}

// GetMyIdentityInfo is a free data retrieval call binding the contract method 0xf388ad23.
//
// Solidity: function getMyIdentityInfo() constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatSession) GetMyIdentityInfo() (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _PermaChat.Contract.GetMyIdentityInfo(&_PermaChat.CallOpts)
}

// GetMyIdentityInfo is a free data retrieval call binding the contract method 0xf388ad23.
//
// Solidity: function getMyIdentityInfo() constant returns(id uint256, providerCount uint256, rejectionCount uint256)
func (_PermaChat *PermaChatCallerSession) GetMyIdentityInfo() (struct {
	Id             *big.Int
	ProviderCount  *big.Int
	RejectionCount *big.Int
}, error) {
	return _PermaChat.Contract.GetMyIdentityInfo(&_PermaChat.CallOpts)
}

// GetMyProviderInfo is a free data retrieval call binding the contract method 0xd64c73a6.
//
// Solidity: function getMyProviderInfo(providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_PermaChat *PermaChatCaller) GetMyProviderInfo(opts *bind.CallOpts, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	ret := new(struct {
		RequestId       *big.Int
		IdentityAddress common.Address
		Provider        string
		UserName        string
		Timestamp       *big.Int
		Active          bool
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getMyProviderInfo", providerId)
	return *ret, err
}

// GetMyProviderInfo is a free data retrieval call binding the contract method 0xd64c73a6.
//
// Solidity: function getMyProviderInfo(providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_PermaChat *PermaChatSession) GetMyProviderInfo(providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _PermaChat.Contract.GetMyProviderInfo(&_PermaChat.CallOpts, providerId)
}

// GetMyProviderInfo is a free data retrieval call binding the contract method 0xd64c73a6.
//
// Solidity: function getMyProviderInfo(providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_PermaChat *PermaChatCallerSession) GetMyProviderInfo(providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _PermaChat.Contract.GetMyProviderInfo(&_PermaChat.CallOpts, providerId)
}

// GetNextRequest is a free data retrieval call binding the contract method 0x95dfb896.
//
// Solidity: function getNextRequest() constant returns(requestId uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_PermaChat *PermaChatCaller) GetNextRequest(opts *bind.CallOpts) (struct {
	RequestId  *big.Int
	Requestor  common.Address
	Provider   string
	UserName   string
	Identifier string
	Timestamp  *big.Int
}, error) {
	ret := new(struct {
		RequestId  *big.Int
		Requestor  common.Address
		Provider   string
		UserName   string
		Identifier string
		Timestamp  *big.Int
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getNextRequest")
	return *ret, err
}

// GetNextRequest is a free data retrieval call binding the contract method 0x95dfb896.
//
// Solidity: function getNextRequest() constant returns(requestId uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_PermaChat *PermaChatSession) GetNextRequest() (struct {
	RequestId  *big.Int
	Requestor  common.Address
	Provider   string
	UserName   string
	Identifier string
	Timestamp  *big.Int
}, error) {
	return _PermaChat.Contract.GetNextRequest(&_PermaChat.CallOpts)
}

// GetNextRequest is a free data retrieval call binding the contract method 0x95dfb896.
//
// Solidity: function getNextRequest() constant returns(requestId uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_PermaChat *PermaChatCallerSession) GetNextRequest() (struct {
	RequestId  *big.Int
	Requestor  common.Address
	Provider   string
	UserName   string
	Identifier string
	Timestamp  *big.Int
}, error) {
	return _PermaChat.Contract.GetNextRequest(&_PermaChat.CallOpts)
}

// GetPinnedPost is a free data retrieval call binding the contract method 0xa46325df.
//
// Solidity: function getPinnedPost() constant returns(uint256)
func (_PermaChat *PermaChatCaller) GetPinnedPost(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getPinnedPost")
	return *ret0, err
}

// GetPinnedPost is a free data retrieval call binding the contract method 0xa46325df.
//
// Solidity: function getPinnedPost() constant returns(uint256)
func (_PermaChat *PermaChatSession) GetPinnedPost() (*big.Int, error) {
	return _PermaChat.Contract.GetPinnedPost(&_PermaChat.CallOpts)
}

// GetPinnedPost is a free data retrieval call binding the contract method 0xa46325df.
//
// Solidity: function getPinnedPost() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) GetPinnedPost() (*big.Int, error) {
	return _PermaChat.Contract.GetPinnedPost(&_PermaChat.CallOpts)
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(postId uint256) constant returns(ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatCaller) GetPost(opts *bind.CallOpts, postId *big.Int) (struct {
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
}, error) {
	ret := new(struct {
		IpfsHash    string
		Poster      common.Address
		BlockNumber *big.Int
		Timestamp   *big.Int
		ContentType string
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getPost", postId)
	return *ret, err
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(postId uint256) constant returns(ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatSession) GetPost(postId *big.Int) (struct {
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
}, error) {
	return _PermaChat.Contract.GetPost(&_PermaChat.CallOpts, postId)
}

// GetPost is a free data retrieval call binding the contract method 0x40731c24.
//
// Solidity: function getPost(postId uint256) constant returns(ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatCallerSession) GetPost(postId *big.Int) (struct {
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
}, error) {
	return _PermaChat.Contract.GetPost(&_PermaChat.CallOpts, postId)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(identityId uint256, providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_PermaChat *PermaChatCaller) GetProviderInfo(opts *bind.CallOpts, identityId *big.Int, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	ret := new(struct {
		RequestId       *big.Int
		IdentityAddress common.Address
		Provider        string
		UserName        string
		Timestamp       *big.Int
		Active          bool
	})
	out := ret
	err := _PermaChat.contract.Call(opts, out, "getProviderInfo", identityId, providerId)
	return *ret, err
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(identityId uint256, providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_PermaChat *PermaChatSession) GetProviderInfo(identityId *big.Int, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _PermaChat.Contract.GetProviderInfo(&_PermaChat.CallOpts, identityId, providerId)
}

// GetProviderInfo is a free data retrieval call binding the contract method 0xba438b3b.
//
// Solidity: function getProviderInfo(identityId uint256, providerId uint256) constant returns(requestId uint256, identityAddress address, provider string, userName string, timestamp uint256, active bool)
func (_PermaChat *PermaChatCallerSession) GetProviderInfo(identityId *big.Int, providerId *big.Int) (struct {
	RequestId       *big.Int
	IdentityAddress common.Address
	Provider        string
	UserName        string
	Timestamp       *big.Int
	Active          bool
}, error) {
	return _PermaChat.Contract.GetProviderInfo(&_PermaChat.CallOpts, identityId, providerId)
}

// GetReplies is a free data retrieval call binding the contract method 0xae58449d.
//
// Solidity: function getReplies(postId uint256) constant returns(replyIds uint256[])
func (_PermaChat *PermaChatCaller) GetReplies(opts *bind.CallOpts, postId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getReplies", postId)
	return *ret0, err
}

// GetReplies is a free data retrieval call binding the contract method 0xae58449d.
//
// Solidity: function getReplies(postId uint256) constant returns(replyIds uint256[])
func (_PermaChat *PermaChatSession) GetReplies(postId *big.Int) ([]*big.Int, error) {
	return _PermaChat.Contract.GetReplies(&_PermaChat.CallOpts, postId)
}

// GetReplies is a free data retrieval call binding the contract method 0xae58449d.
//
// Solidity: function getReplies(postId uint256) constant returns(replyIds uint256[])
func (_PermaChat *PermaChatCallerSession) GetReplies(postId *big.Int) ([]*big.Int, error) {
	return _PermaChat.Contract.GetReplies(&_PermaChat.CallOpts, postId)
}

// GetReplyParent is a free data retrieval call binding the contract method 0x2637adbf.
//
// Solidity: function getReplyParent(replyId uint256) constant returns(parentPostId uint256)
func (_PermaChat *PermaChatCaller) GetReplyParent(opts *bind.CallOpts, replyId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "getReplyParent", replyId)
	return *ret0, err
}

// GetReplyParent is a free data retrieval call binding the contract method 0x2637adbf.
//
// Solidity: function getReplyParent(replyId uint256) constant returns(parentPostId uint256)
func (_PermaChat *PermaChatSession) GetReplyParent(replyId *big.Int) (*big.Int, error) {
	return _PermaChat.Contract.GetReplyParent(&_PermaChat.CallOpts, replyId)
}

// GetReplyParent is a free data retrieval call binding the contract method 0x2637adbf.
//
// Solidity: function getReplyParent(replyId uint256) constant returns(parentPostId uint256)
func (_PermaChat *PermaChatCallerSession) GetReplyParent(replyId *big.Int) (*big.Int, error) {
	return _PermaChat.Contract.GetReplyParent(&_PermaChat.CallOpts, replyId)
}

// GetRequestById is a free data retrieval call binding the contract method 0x26757b73.
//
// Solidity: function getRequestById(requestId uint256) constant returns(uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_PermaChat *PermaChatCaller) GetRequestById(opts *bind.CallOpts, requestId *big.Int) (*big.Int, common.Address, string, string, string, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(string)
		ret3 = new(string)
		ret4 = new(string)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _PermaChat.contract.Call(opts, out, "getRequestById", requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestById is a free data retrieval call binding the contract method 0x26757b73.
//
// Solidity: function getRequestById(requestId uint256) constant returns(uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_PermaChat *PermaChatSession) GetRequestById(requestId *big.Int) (*big.Int, common.Address, string, string, string, *big.Int, error) {
	return _PermaChat.Contract.GetRequestById(&_PermaChat.CallOpts, requestId)
}

// GetRequestById is a free data retrieval call binding the contract method 0x26757b73.
//
// Solidity: function getRequestById(requestId uint256) constant returns(uint256, requestor address, provider string, userName string, identifier string, timestamp uint256)
func (_PermaChat *PermaChatCallerSession) GetRequestById(requestId *big.Int) (*big.Int, common.Address, string, string, string, *big.Int, error) {
	return _PermaChat.Contract.GetRequestById(&_PermaChat.CallOpts, requestId)
}

// HasCommented is a free data retrieval call binding the contract method 0x7af2d1af.
//
// Solidity: function hasCommented(postId uint256) constant returns(bool)
func (_PermaChat *PermaChatCaller) HasCommented(opts *bind.CallOpts, postId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "hasCommented", postId)
	return *ret0, err
}

// HasCommented is a free data retrieval call binding the contract method 0x7af2d1af.
//
// Solidity: function hasCommented(postId uint256) constant returns(bool)
func (_PermaChat *PermaChatSession) HasCommented(postId *big.Int) (bool, error) {
	return _PermaChat.Contract.HasCommented(&_PermaChat.CallOpts, postId)
}

// HasCommented is a free data retrieval call binding the contract method 0x7af2d1af.
//
// Solidity: function hasCommented(postId uint256) constant returns(bool)
func (_PermaChat *PermaChatCallerSession) HasCommented(postId *big.Int) (bool, error) {
	return _PermaChat.Contract.HasCommented(&_PermaChat.CallOpts, postId)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_PermaChat *PermaChatCaller) IdentityAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "identityAdmin")
	return *ret0, err
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_PermaChat *PermaChatSession) IdentityAdmin() (common.Address, error) {
	return _PermaChat.Contract.IdentityAdmin(&_PermaChat.CallOpts)
}

// IdentityAdmin is a free data retrieval call binding the contract method 0x7e889537.
//
// Solidity: function identityAdmin() constant returns(address)
func (_PermaChat *PermaChatCallerSession) IdentityAdmin() (common.Address, error) {
	return _PermaChat.Contract.IdentityAdmin(&_PermaChat.CallOpts)
}

// LatestDatabaseIndex is a free data retrieval call binding the contract method 0x39fa083d.
//
// Solidity: function latestDatabaseIndex() constant returns(uint256)
func (_PermaChat *PermaChatCaller) LatestDatabaseIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "latestDatabaseIndex")
	return *ret0, err
}

// LatestDatabaseIndex is a free data retrieval call binding the contract method 0x39fa083d.
//
// Solidity: function latestDatabaseIndex() constant returns(uint256)
func (_PermaChat *PermaChatSession) LatestDatabaseIndex() (*big.Int, error) {
	return _PermaChat.Contract.LatestDatabaseIndex(&_PermaChat.CallOpts)
}

// LatestDatabaseIndex is a free data retrieval call binding the contract method 0x39fa083d.
//
// Solidity: function latestDatabaseIndex() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) LatestDatabaseIndex() (*big.Int, error) {
	return _PermaChat.Contract.LatestDatabaseIndex(&_PermaChat.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_PermaChat *PermaChatCaller) NextActionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "nextActionId")
	return *ret0, err
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_PermaChat *PermaChatSession) NextActionId() (*big.Int, error) {
	return _PermaChat.Contract.NextActionId(&_PermaChat.CallOpts)
}

// NextActionId is a free data retrieval call binding the contract method 0x46abe73a.
//
// Solidity: function nextActionId() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) NextActionId() (*big.Int, error) {
	return _PermaChat.Contract.NextActionId(&_PermaChat.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_PermaChat *PermaChatCaller) NextCommentIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "nextCommentIndex")
	return *ret0, err
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_PermaChat *PermaChatSession) NextCommentIndex() (*big.Int, error) {
	return _PermaChat.Contract.NextCommentIndex(&_PermaChat.CallOpts)
}

// NextCommentIndex is a free data retrieval call binding the contract method 0x61165d08.
//
// Solidity: function nextCommentIndex() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) NextCommentIndex() (*big.Int, error) {
	return _PermaChat.Contract.NextCommentIndex(&_PermaChat.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_PermaChat *PermaChatCaller) NextIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "nextIdentityRequestIndex")
	return *ret0, err
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_PermaChat *PermaChatSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _PermaChat.Contract.NextIdentityRequestIndex(&_PermaChat.CallOpts)
}

// NextIdentityRequestIndex is a free data retrieval call binding the contract method 0x4d625d99.
//
// Solidity: function nextIdentityRequestIndex() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) NextIdentityRequestIndex() (*big.Int, error) {
	return _PermaChat.Contract.NextIdentityRequestIndex(&_PermaChat.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_PermaChat *PermaChatCaller) NextPostIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "nextPostIndex")
	return *ret0, err
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_PermaChat *PermaChatSession) NextPostIndex() (*big.Int, error) {
	return _PermaChat.Contract.NextPostIndex(&_PermaChat.CallOpts)
}

// NextPostIndex is a free data retrieval call binding the contract method 0x8ef5d2b9.
//
// Solidity: function nextPostIndex() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) NextPostIndex() (*big.Int, error) {
	return _PermaChat.Contract.NextPostIndex(&_PermaChat.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_PermaChat *PermaChatCaller) PendingIdentityRequestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "pendingIdentityRequestIndex")
	return *ret0, err
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_PermaChat *PermaChatSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _PermaChat.Contract.PendingIdentityRequestIndex(&_PermaChat.CallOpts)
}

// PendingIdentityRequestIndex is a free data retrieval call binding the contract method 0x75efbecf.
//
// Solidity: function pendingIdentityRequestIndex() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) PendingIdentityRequestIndex() (*big.Int, error) {
	return _PermaChat.Contract.PendingIdentityRequestIndex(&_PermaChat.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_PermaChat *PermaChatCaller) RequestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermaChat.contract.Call(opts, out, "requestPrice")
	return *ret0, err
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_PermaChat *PermaChatSession) RequestPrice() (*big.Int, error) {
	return _PermaChat.Contract.RequestPrice(&_PermaChat.CallOpts)
}

// RequestPrice is a free data retrieval call binding the contract method 0x1604f9ea.
//
// Solidity: function requestPrice() constant returns(uint256)
func (_PermaChat *PermaChatCallerSession) RequestPrice() (*big.Int, error) {
	return _PermaChat.Contract.RequestPrice(&_PermaChat.CallOpts)
}

// AddDownvote is a paid mutator transaction binding the contract method 0x01c60907.
//
// Solidity: function addDownvote(postId uint256) returns()
func (_PermaChat *PermaChatTransactor) AddDownvote(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "addDownvote", postId)
}

// AddDownvote is a paid mutator transaction binding the contract method 0x01c60907.
//
// Solidity: function addDownvote(postId uint256) returns()
func (_PermaChat *PermaChatSession) AddDownvote(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddDownvote(&_PermaChat.TransactOpts, postId)
}

// AddDownvote is a paid mutator transaction binding the contract method 0x01c60907.
//
// Solidity: function addDownvote(postId uint256) returns()
func (_PermaChat *PermaChatTransactorSession) AddDownvote(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddDownvote(&_PermaChat.TransactOpts, postId)
}

// AddFlag is a paid mutator transaction binding the contract method 0xae2cded6.
//
// Solidity: function addFlag(postId uint256) returns()
func (_PermaChat *PermaChatTransactor) AddFlag(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "addFlag", postId)
}

// AddFlag is a paid mutator transaction binding the contract method 0xae2cded6.
//
// Solidity: function addFlag(postId uint256) returns()
func (_PermaChat *PermaChatSession) AddFlag(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddFlag(&_PermaChat.TransactOpts, postId)
}

// AddFlag is a paid mutator transaction binding the contract method 0xae2cded6.
//
// Solidity: function addFlag(postId uint256) returns()
func (_PermaChat *PermaChatTransactorSession) AddFlag(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddFlag(&_PermaChat.TransactOpts, postId)
}

// AddTip is a paid mutator transaction binding the contract method 0x8e26c7e6.
//
// Solidity: function addTip(postId uint256) returns()
func (_PermaChat *PermaChatTransactor) AddTip(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "addTip", postId)
}

// AddTip is a paid mutator transaction binding the contract method 0x8e26c7e6.
//
// Solidity: function addTip(postId uint256) returns()
func (_PermaChat *PermaChatSession) AddTip(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddTip(&_PermaChat.TransactOpts, postId)
}

// AddTip is a paid mutator transaction binding the contract method 0x8e26c7e6.
//
// Solidity: function addTip(postId uint256) returns()
func (_PermaChat *PermaChatTransactorSession) AddTip(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddTip(&_PermaChat.TransactOpts, postId)
}

// AddUpvote is a paid mutator transaction binding the contract method 0xf1de9ea4.
//
// Solidity: function addUpvote(postId uint256) returns()
func (_PermaChat *PermaChatTransactor) AddUpvote(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "addUpvote", postId)
}

// AddUpvote is a paid mutator transaction binding the contract method 0xf1de9ea4.
//
// Solidity: function addUpvote(postId uint256) returns()
func (_PermaChat *PermaChatSession) AddUpvote(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddUpvote(&_PermaChat.TransactOpts, postId)
}

// AddUpvote is a paid mutator transaction binding the contract method 0xf1de9ea4.
//
// Solidity: function addUpvote(postId uint256) returns()
func (_PermaChat *PermaChatTransactorSession) AddUpvote(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.AddUpvote(&_PermaChat.TransactOpts, postId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_PermaChat *PermaChatTransactor) Approve(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "approve", requestId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_PermaChat *PermaChatSession) Approve(requestId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.Approve(&_PermaChat.TransactOpts, requestId)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(requestId uint256) returns()
func (_PermaChat *PermaChatTransactorSession) Approve(requestId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.Approve(&_PermaChat.TransactOpts, requestId)
}

// NewPost is a paid mutator transaction binding the contract method 0x9e28cf96.
//
// Solidity: function newPost(ipfsHash string, contentType string) returns(postId uint256)
func (_PermaChat *PermaChatTransactor) NewPost(opts *bind.TransactOpts, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "newPost", ipfsHash, contentType)
}

// NewPost is a paid mutator transaction binding the contract method 0x9e28cf96.
//
// Solidity: function newPost(ipfsHash string, contentType string) returns(postId uint256)
func (_PermaChat *PermaChatSession) NewPost(ipfsHash string, contentType string) (*types.Transaction, error) {
	return _PermaChat.Contract.NewPost(&_PermaChat.TransactOpts, ipfsHash, contentType)
}

// NewPost is a paid mutator transaction binding the contract method 0x9e28cf96.
//
// Solidity: function newPost(ipfsHash string, contentType string) returns(postId uint256)
func (_PermaChat *PermaChatTransactorSession) NewPost(ipfsHash string, contentType string) (*types.Transaction, error) {
	return _PermaChat.Contract.NewPost(&_PermaChat.TransactOpts, ipfsHash, contentType)
}

// NewReply is a paid mutator transaction binding the contract method 0x8faa6dad.
//
// Solidity: function newReply(parentPostId uint256, ipfsHash string, contentType string) returns(uint256)
func (_PermaChat *PermaChatTransactor) NewReply(opts *bind.TransactOpts, parentPostId *big.Int, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "newReply", parentPostId, ipfsHash, contentType)
}

// NewReply is a paid mutator transaction binding the contract method 0x8faa6dad.
//
// Solidity: function newReply(parentPostId uint256, ipfsHash string, contentType string) returns(uint256)
func (_PermaChat *PermaChatSession) NewReply(parentPostId *big.Int, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _PermaChat.Contract.NewReply(&_PermaChat.TransactOpts, parentPostId, ipfsHash, contentType)
}

// NewReply is a paid mutator transaction binding the contract method 0x8faa6dad.
//
// Solidity: function newReply(parentPostId uint256, ipfsHash string, contentType string) returns(uint256)
func (_PermaChat *PermaChatTransactorSession) NewReply(parentPostId *big.Int, ipfsHash string, contentType string) (*types.Transaction, error) {
	return _PermaChat.Contract.NewReply(&_PermaChat.TransactOpts, parentPostId, ipfsHash, contentType)
}

// PinPost is a paid mutator transaction binding the contract method 0x251cf4ea.
//
// Solidity: function pinPost(postId uint256) returns()
func (_PermaChat *PermaChatTransactor) PinPost(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "pinPost", postId)
}

// PinPost is a paid mutator transaction binding the contract method 0x251cf4ea.
//
// Solidity: function pinPost(postId uint256) returns()
func (_PermaChat *PermaChatSession) PinPost(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.PinPost(&_PermaChat.TransactOpts, postId)
}

// PinPost is a paid mutator transaction binding the contract method 0x251cf4ea.
//
// Solidity: function pinPost(postId uint256) returns()
func (_PermaChat *PermaChatTransactorSession) PinPost(postId *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.PinPost(&_PermaChat.TransactOpts, postId)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_PermaChat *PermaChatTransactor) Reject(opts *bind.TransactOpts, requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "reject", requestId, reason)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_PermaChat *PermaChatSession) Reject(requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _PermaChat.Contract.Reject(&_PermaChat.TransactOpts, requestId, reason)
}

// Reject is a paid mutator transaction binding the contract method 0x9119c8d7.
//
// Solidity: function reject(requestId uint256, reason uint8) returns()
func (_PermaChat *PermaChatTransactorSession) Reject(requestId *big.Int, reason uint8) (*types.Transaction, error) {
	return _PermaChat.Contract.Reject(&_PermaChat.TransactOpts, requestId, reason)
}

// SetNewDB is a paid mutator transaction binding the contract method 0x474c49dd.
//
// Solidity: function setNewDB(ipfsHash string) returns()
func (_PermaChat *PermaChatTransactor) SetNewDB(opts *bind.TransactOpts, ipfsHash string) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "setNewDB", ipfsHash)
}

// SetNewDB is a paid mutator transaction binding the contract method 0x474c49dd.
//
// Solidity: function setNewDB(ipfsHash string) returns()
func (_PermaChat *PermaChatSession) SetNewDB(ipfsHash string) (*types.Transaction, error) {
	return _PermaChat.Contract.SetNewDB(&_PermaChat.TransactOpts, ipfsHash)
}

// SetNewDB is a paid mutator transaction binding the contract method 0x474c49dd.
//
// Solidity: function setNewDB(ipfsHash string) returns()
func (_PermaChat *PermaChatTransactorSession) SetNewDB(ipfsHash string) (*types.Transaction, error) {
	return _PermaChat.Contract.SetNewDB(&_PermaChat.TransactOpts, ipfsHash)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_PermaChat *PermaChatTransactor) SetRequestPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "setRequestPrice", price)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_PermaChat *PermaChatSession) SetRequestPrice(price *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.SetRequestPrice(&_PermaChat.TransactOpts, price)
}

// SetRequestPrice is a paid mutator transaction binding the contract method 0x3fcf7ca1.
//
// Solidity: function setRequestPrice(price uint256) returns()
func (_PermaChat *PermaChatTransactorSession) SetRequestPrice(price *big.Int) (*types.Transaction, error) {
	return _PermaChat.Contract.SetRequestPrice(&_PermaChat.TransactOpts, price)
}

// SubmitRequest is a paid mutator transaction binding the contract method 0x470c153b.
//
// Solidity: function submitRequest(provider string, userName string, identifier string) returns()
func (_PermaChat *PermaChatTransactor) SubmitRequest(opts *bind.TransactOpts, provider string, userName string, identifier string) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "submitRequest", provider, userName, identifier)
}

// SubmitRequest is a paid mutator transaction binding the contract method 0x470c153b.
//
// Solidity: function submitRequest(provider string, userName string, identifier string) returns()
func (_PermaChat *PermaChatSession) SubmitRequest(provider string, userName string, identifier string) (*types.Transaction, error) {
	return _PermaChat.Contract.SubmitRequest(&_PermaChat.TransactOpts, provider, userName, identifier)
}

// SubmitRequest is a paid mutator transaction binding the contract method 0x470c153b.
//
// Solidity: function submitRequest(provider string, userName string, identifier string) returns()
func (_PermaChat *PermaChatTransactorSession) SubmitRequest(provider string, userName string, identifier string) (*types.Transaction, error) {
	return _PermaChat.Contract.SubmitRequest(&_PermaChat.TransactOpts, provider, userName, identifier)
}

// Unpin is a paid mutator transaction binding the contract method 0xfbd5b4e1.
//
// Solidity: function unpin() returns()
func (_PermaChat *PermaChatTransactor) Unpin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermaChat.contract.Transact(opts, "unpin")
}

// Unpin is a paid mutator transaction binding the contract method 0xfbd5b4e1.
//
// Solidity: function unpin() returns()
func (_PermaChat *PermaChatSession) Unpin() (*types.Transaction, error) {
	return _PermaChat.Contract.Unpin(&_PermaChat.TransactOpts)
}

// Unpin is a paid mutator transaction binding the contract method 0xfbd5b4e1.
//
// Solidity: function unpin() returns()
func (_PermaChat *PermaChatTransactorSession) Unpin() (*types.Transaction, error) {
	return _PermaChat.Contract.Unpin(&_PermaChat.TransactOpts)
}

// PermaChatDBUpdateEventIterator is returned from FilterDBUpdateEvent and is used to iterate over the raw logs and unpacked data for DBUpdateEvent events raised by the PermaChat contract.
type PermaChatDBUpdateEventIterator struct {
	Event *PermaChatDBUpdateEvent // Event containing the contract specifics and raw log

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
func (it *PermaChatDBUpdateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermaChatDBUpdateEvent)
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
		it.Event = new(PermaChatDBUpdateEvent)
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
func (it *PermaChatDBUpdateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermaChatDBUpdateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermaChatDBUpdateEvent represents a DBUpdateEvent event raised by the PermaChat contract.
type PermaChatDBUpdateEvent struct {
	CurrentDatabaseIndex *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDBUpdateEvent is a free log retrieval operation binding the contract event 0x50634acbe4d09fc2a5c9ee5d100551a5a9e0162ed4bf7cb06c3115a7cfb25ed9.
//
// Solidity: e DBUpdateEvent(currentDatabaseIndex uint256)
func (_PermaChat *PermaChatFilterer) FilterDBUpdateEvent(opts *bind.FilterOpts) (*PermaChatDBUpdateEventIterator, error) {

	logs, sub, err := _PermaChat.contract.FilterLogs(opts, "DBUpdateEvent")
	if err != nil {
		return nil, err
	}
	return &PermaChatDBUpdateEventIterator{contract: _PermaChat.contract, event: "DBUpdateEvent", logs: logs, sub: sub}, nil
}

// WatchDBUpdateEvent is a free log subscription operation binding the contract event 0x50634acbe4d09fc2a5c9ee5d100551a5a9e0162ed4bf7cb06c3115a7cfb25ed9.
//
// Solidity: e DBUpdateEvent(currentDatabaseIndex uint256)
func (_PermaChat *PermaChatFilterer) WatchDBUpdateEvent(opts *bind.WatchOpts, sink chan<- *PermaChatDBUpdateEvent) (event.Subscription, error) {

	logs, sub, err := _PermaChat.contract.WatchLogs(opts, "DBUpdateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermaChatDBUpdateEvent)
				if err := _PermaChat.contract.UnpackLog(event, "DBUpdateEvent", log); err != nil {
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

// PermaChatNewApprovedProviderEventIterator is returned from FilterNewApprovedProviderEvent and is used to iterate over the raw logs and unpacked data for NewApprovedProviderEvent events raised by the PermaChat contract.
type PermaChatNewApprovedProviderEventIterator struct {
	Event *PermaChatNewApprovedProviderEvent // Event containing the contract specifics and raw log

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
func (it *PermaChatNewApprovedProviderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermaChatNewApprovedProviderEvent)
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
		it.Event = new(PermaChatNewApprovedProviderEvent)
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
func (it *PermaChatNewApprovedProviderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermaChatNewApprovedProviderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermaChatNewApprovedProviderEvent represents a NewApprovedProviderEvent event raised by the PermaChat contract.
type PermaChatNewApprovedProviderEvent struct {
	IdentityId *big.Int
	ProviderId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewApprovedProviderEvent is a free log retrieval operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_PermaChat *PermaChatFilterer) FilterNewApprovedProviderEvent(opts *bind.FilterOpts) (*PermaChatNewApprovedProviderEventIterator, error) {

	logs, sub, err := _PermaChat.contract.FilterLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return &PermaChatNewApprovedProviderEventIterator{contract: _PermaChat.contract, event: "NewApprovedProviderEvent", logs: logs, sub: sub}, nil
}

// WatchNewApprovedProviderEvent is a free log subscription operation binding the contract event 0x32cade7c48dc67e9c10083ef64ab0a0ae73e2eaf20d00cec79f41eb3af5ae979.
//
// Solidity: e NewApprovedProviderEvent(identityId uint256, providerId uint256)
func (_PermaChat *PermaChatFilterer) WatchNewApprovedProviderEvent(opts *bind.WatchOpts, sink chan<- *PermaChatNewApprovedProviderEvent) (event.Subscription, error) {

	logs, sub, err := _PermaChat.contract.WatchLogs(opts, "NewApprovedProviderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermaChatNewApprovedProviderEvent)
				if err := _PermaChat.contract.UnpackLog(event, "NewApprovedProviderEvent", log); err != nil {
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

// PermaChatNewCommentaryEventIterator is returned from FilterNewCommentaryEvent and is used to iterate over the raw logs and unpacked data for NewCommentaryEvent events raised by the PermaChat contract.
type PermaChatNewCommentaryEventIterator struct {
	Event *PermaChatNewCommentaryEvent // Event containing the contract specifics and raw log

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
func (it *PermaChatNewCommentaryEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermaChatNewCommentaryEvent)
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
		it.Event = new(PermaChatNewCommentaryEvent)
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
func (it *PermaChatNewCommentaryEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermaChatNewCommentaryEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermaChatNewCommentaryEvent represents a NewCommentaryEvent event raised by the PermaChat contract.
type PermaChatNewCommentaryEvent struct {
	PostId          *big.Int
	CommentaryIndex *big.Int
	CommentId       *big.Int
	CommentaryType  uint8
	Commenter       common.Address
	Timestamp       *big.Int
	Tip             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCommentaryEvent is a free log retrieval operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_PermaChat *PermaChatFilterer) FilterNewCommentaryEvent(opts *bind.FilterOpts) (*PermaChatNewCommentaryEventIterator, error) {

	logs, sub, err := _PermaChat.contract.FilterLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return &PermaChatNewCommentaryEventIterator{contract: _PermaChat.contract, event: "NewCommentaryEvent", logs: logs, sub: sub}, nil
}

// WatchNewCommentaryEvent is a free log subscription operation binding the contract event 0xa795a9ff0736cf8985b849f99bf89309aa6fef525f19ec5a5c7617f27794bc15.
//
// Solidity: e NewCommentaryEvent(postId uint256, commentaryIndex uint256, commentId uint256, commentaryType uint8, commenter address, timestamp uint256, tip uint256)
func (_PermaChat *PermaChatFilterer) WatchNewCommentaryEvent(opts *bind.WatchOpts, sink chan<- *PermaChatNewCommentaryEvent) (event.Subscription, error) {

	logs, sub, err := _PermaChat.contract.WatchLogs(opts, "NewCommentaryEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermaChatNewCommentaryEvent)
				if err := _PermaChat.contract.UnpackLog(event, "NewCommentaryEvent", log); err != nil {
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

// PermaChatNewPostEventIterator is returned from FilterNewPostEvent and is used to iterate over the raw logs and unpacked data for NewPostEvent events raised by the PermaChat contract.
type PermaChatNewPostEventIterator struct {
	Event *PermaChatNewPostEvent // Event containing the contract specifics and raw log

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
func (it *PermaChatNewPostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermaChatNewPostEvent)
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
		it.Event = new(PermaChatNewPostEvent)
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
func (it *PermaChatNewPostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermaChatNewPostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermaChatNewPostEvent represents a NewPostEvent event raised by the PermaChat contract.
type PermaChatNewPostEvent struct {
	PostId      *big.Int
	IpfsHash    string
	Poster      common.Address
	BlockNumber *big.Int
	Timestamp   *big.Int
	ContentType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPostEvent is a free log retrieval operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatFilterer) FilterNewPostEvent(opts *bind.FilterOpts) (*PermaChatNewPostEventIterator, error) {

	logs, sub, err := _PermaChat.contract.FilterLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return &PermaChatNewPostEventIterator{contract: _PermaChat.contract, event: "NewPostEvent", logs: logs, sub: sub}, nil
}

// WatchNewPostEvent is a free log subscription operation binding the contract event 0xab818809a8dbeb592e34e43ffab44637a97fdeb7a2fa51b01db3d2e3aaac2754.
//
// Solidity: e NewPostEvent(postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatFilterer) WatchNewPostEvent(opts *bind.WatchOpts, sink chan<- *PermaChatNewPostEvent) (event.Subscription, error) {

	logs, sub, err := _PermaChat.contract.WatchLogs(opts, "NewPostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermaChatNewPostEvent)
				if err := _PermaChat.contract.UnpackLog(event, "NewPostEvent", log); err != nil {
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

// PermaChatNewReplyEventIterator is returned from FilterNewReplyEvent and is used to iterate over the raw logs and unpacked data for NewReplyEvent events raised by the PermaChat contract.
type PermaChatNewReplyEventIterator struct {
	Event *PermaChatNewReplyEvent // Event containing the contract specifics and raw log

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
func (it *PermaChatNewReplyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermaChatNewReplyEvent)
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
		it.Event = new(PermaChatNewReplyEvent)
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
func (it *PermaChatNewReplyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermaChatNewReplyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermaChatNewReplyEvent represents a NewReplyEvent event raised by the PermaChat contract.
type PermaChatNewReplyEvent struct {
	ParentPostId *big.Int
	PostId       *big.Int
	IpfsHash     string
	Poster       common.Address
	BlockNumber  *big.Int
	Timestamp    *big.Int
	ContentType  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewReplyEvent is a free log retrieval operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatFilterer) FilterNewReplyEvent(opts *bind.FilterOpts) (*PermaChatNewReplyEventIterator, error) {

	logs, sub, err := _PermaChat.contract.FilterLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return &PermaChatNewReplyEventIterator{contract: _PermaChat.contract, event: "NewReplyEvent", logs: logs, sub: sub}, nil
}

// WatchNewReplyEvent is a free log subscription operation binding the contract event 0x378296500f6b4fa4c53cd02c0cccb2694b4798d5d1e2077c027fb6214cccffa3.
//
// Solidity: e NewReplyEvent(parentPostId uint256, postId uint256, ipfsHash string, poster address, blockNumber uint256, timestamp uint256, contentType string)
func (_PermaChat *PermaChatFilterer) WatchNewReplyEvent(opts *bind.WatchOpts, sink chan<- *PermaChatNewReplyEvent) (event.Subscription, error) {

	logs, sub, err := _PermaChat.contract.WatchLogs(opts, "NewReplyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermaChatNewReplyEvent)
				if err := _PermaChat.contract.UnpackLog(event, "NewReplyEvent", log); err != nil {
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

// PermaChatNewRequestEventIterator is returned from FilterNewRequestEvent and is used to iterate over the raw logs and unpacked data for NewRequestEvent events raised by the PermaChat contract.
type PermaChatNewRequestEventIterator struct {
	Event *PermaChatNewRequestEvent // Event containing the contract specifics and raw log

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
func (it *PermaChatNewRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermaChatNewRequestEvent)
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
		it.Event = new(PermaChatNewRequestEvent)
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
func (it *PermaChatNewRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermaChatNewRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermaChatNewRequestEvent represents a NewRequestEvent event raised by the PermaChat contract.
type PermaChatNewRequestEvent struct {
	IdentityRequestId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRequestEvent is a free log retrieval operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_PermaChat *PermaChatFilterer) FilterNewRequestEvent(opts *bind.FilterOpts) (*PermaChatNewRequestEventIterator, error) {

	logs, sub, err := _PermaChat.contract.FilterLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return &PermaChatNewRequestEventIterator{contract: _PermaChat.contract, event: "NewRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewRequestEvent is a free log subscription operation binding the contract event 0x6b0584cd69ce9b6c815448eca200446bfcb1f1ee17b73b97af1e6bdb4f45f04d.
//
// Solidity: e NewRequestEvent(identityRequestId uint256)
func (_PermaChat *PermaChatFilterer) WatchNewRequestEvent(opts *bind.WatchOpts, sink chan<- *PermaChatNewRequestEvent) (event.Subscription, error) {

	logs, sub, err := _PermaChat.contract.WatchLogs(opts, "NewRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermaChatNewRequestEvent)
				if err := _PermaChat.contract.UnpackLog(event, "NewRequestEvent", log); err != nil {
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
