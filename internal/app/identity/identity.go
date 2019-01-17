package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// TODO: upgrade to less stupid logger...try logrus?
	log.SetOutput(os.Stdout)

	log.Println("START")
	config := initializeConfig("permachat")

	log.Println("initializing contract")
	contract, txContext, err := initializeContract(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("updating price")
	err1 := updateRequestPrice(contract, txContext)
	if err1 != nil {
		log.Println("yuck")
		log.Fatal(err1)
	}

	validators := map[string]Validator{
		"twitter": newTwitterValidator(),
		"github":  newGitHubValidator(),
	}

	log.Println("initializing next request handler")
	handler, err := NewRequestHandler(config, contract, txContext, validators)
	if err != nil {
		log.Println("yug")
		log.Fatal(err)
	}

	log.Println("WAITING FOR EVENTS")

	handler.handlePending()

	log.Println("PROCESSED PENDING")

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	log.Println("exiting")

	sub := *handler.sub
	sub.Unsubscribe()

	log.Println("END")
}
