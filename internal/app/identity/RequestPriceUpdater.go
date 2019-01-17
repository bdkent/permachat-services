package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/palantir/stacktrace"
)

// UpdateRequestPrice is
func updateRequestPrice(contract *PermaChat, txContext *TransactionContext) error {
	log.Printf("updateRequestPrice")
	price := big.NewInt(10000000000000000) // 10 Finney

	res1, err := contract.RequestPrice(nil)
	if err != nil {
		log.Printf("hmmm")
		return stacktrace.Propagate(err, "Could not get current request price")
	}

	log.Printf("transaction: %#v \n", res1)

	auth1, err := txContext.nextTransactor()
	if err != nil {
		return stacktrace.Propagate(err, "Could not create transactor")
	}

	res2, err := contract.SetRequestPrice(auth1, price)
	if err != nil {
		return stacktrace.Propagate(err, "Could not set new request price")
	}

	if res2 == nil {
		return fmt.Errorf("go is stupid")
	}

	return nil
}
