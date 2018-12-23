package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/event"
	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
)

// RequestHandler =
type RequestHandler struct {
	events            chan *PermaChatNewRequestEvent
	contract          *PermaChat
	txContext         *TransactionContext
	sub               *event.Subscription
	validatorRegistry *ValidatorRegistry
}

// NewRequestHandler =
func NewRequestHandler(viper *viper.Viper, contract *PermaChat, txContext *TransactionContext, validators map[string]Validator) (*RequestHandler, error) {

	events := make(chan *PermaChatNewRequestEvent)

	sub, err := contract.WatchNewRequestEvent(nil, events)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Could not start watching for request events transactor")
	}

	registry, err := NewValidatorRegistryWithMapping(viper, validators)
	if err != nil {
		return nil, stacktrace.Propagate(err, "unable to create validator registry")
	}

	handler := RequestHandler{
		events:            events,
		contract:          contract,
		txContext:         txContext,
		sub:               &sub,
		validatorRegistry: registry,
	}

	go handler.awaitEvents()

	return &handler, nil
}

func (handler *RequestHandler) awaitEvents() {
	fmt.Println("nextRequestHandler")
	fmt.Printf("handler: %#v\n", handler)

	for {
		msg := <-handler.events

		fmt.Println("new message")

		err := handle(msg, handler.validatorRegistry, handler.contract, handler.txContext)
		if err != nil {
			log.Printf("unable to process request: %#v\nerror: %v", msg, err)
		}

	}
}

func handle(msg *PermaChatNewRequestEvent, validatorRegistry *ValidatorRegistry, contract *PermaChat, txCon *TransactionContext) error {
	fmt.Printf("event request id: %v\n", msg.IdentityRequestId)

	request, err := contract.GetNextRequest(nil)
	if err != nil {
		return stacktrace.Propagate(err, "Could not get next request")
	}

	fmt.Printf("next request: %v\n", request)

	isValid, err := validatorRegistry.validate(request.Requestor.String(), request.Provider, request.UserName, request.Identifier)
	if err != nil {
		return stacktrace.Propagate(err, "unable to validate request")
	}

	auth, err := txCon.nextTransactor()
	if err != nil {
		return stacktrace.Propagate(err, "unable to create next transaction context")
	}

	if isValid {
		_, err := contract.Approve(auth, request.RequestId)
		if err != nil {
			return stacktrace.Propagate(err, "unable to approve request")
		}
	} else {
		_, err := contract.Reject(auth, request.RequestId, 0)
		if err != nil {
			return stacktrace.Propagate(err, "unable to reject request")
		}
	}

	return nil
}
