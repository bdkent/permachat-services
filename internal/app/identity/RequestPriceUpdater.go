package main

import (
	"fmt"
	"math/big"

	"github.com/palantir/stacktrace"
)

// UpdateRequestPrice is
func updateRequestPrice(contract *PermaChat, txContext *TransactionContext) error {
	fmt.Printf("updateRequestPrice")
	price := big.NewInt(10000000000000000) // 10 Finney

	res1, err := contract.RequestPrice(nil)
	if err != nil {
		fmt.Printf("hmmm")
		return stacktrace.Propagate(err, "Could not get current request price")
	}

	fmt.Printf("transaction: %#v \n", res1)

	auth1, err := txContext.nextTransactor()
	if err != nil {
		return stacktrace.Propagate(err, "Could not create transactor")
	}

	res2, err := contract.SetRequestPrice(auth1, price)
	if err != nil {
		return stacktrace.Propagate(err, "Could not set new request price")
	}

	fmt.Println(res2)

	return nil
}
