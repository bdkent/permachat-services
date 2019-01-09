package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("START")
	viper := initializeConfig("permachat")
	// fmt.Println(viper)

	fmt.Println("initializing contract")
	contract, txContext, err := initializeContract(viper.Sub("contract.private"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("updating price")
	err1 := updateRequestPrice(contract, txContext)
	if err1 != nil {
		fmt.Println("yuck")
		log.Fatal(err1)
	}

	validators := map[string]Validator{
		"twitter": newTwitterValidator(),
		"github":  newGitHubValidator(),
	}

	fmt.Println("initializing next request handler")
	handler, err := NewRequestHandler(viper, contract, txContext, validators)
	if err != nil {
		fmt.Println("yug")
		log.Fatal(err)
	}

	fmt.Println("WAITING FOR EVENTS")

	handler.handlePending()

	fmt.Println("PROCESSED PENDING")

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	fmt.Println("exiting")

	sub := *handler.sub
	sub.Unsubscribe()

	fmt.Println("END")
}
