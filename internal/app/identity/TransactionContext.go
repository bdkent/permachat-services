package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
)

//TransactionContext is
type TransactionContext struct {
	client      *ethclient.Client
	privateKey  *ecdsa.PrivateKey
	fromAddress *common.Address
}

func initializeContract(config *viper.Viper) (*PermaChat, *TransactionContext, error) {

	clientURI := "ws://" + config.GetString("permachat_ethereum_ws_address") + ":" + config.GetString("permachat_ethereum_ws_port")

	client, err := ethclient.Dial(clientURI)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("could not connect to ethereum client: %v", clientURI))
	}

	privateKey, err := crypto.HexToECDSA(config.GetString("permachat_identity_admin_privatekeyhex"))
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("could not create private key"))
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("error casting public key to ECDSA"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contractAddress := common.HexToAddress(config.GetString("permachat_identity_contract_address"))

	instance, err := NewPermaChat(contractAddress, client)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "could not connect to contract")
	}

	txCon := TransactionContext{privateKey: privateKey, client: client, fromAddress: &fromAddress}

	return instance, &txCon, nil
}

func (txCon *TransactionContext) nextTransactor() (*bind.TransactOpts, error) {
	nonce, err := txCon.client.PendingNonceAt(context.Background(), *txCon.fromAddress)
	if err != nil {
		return nil, stacktrace.Propagate(err, "could not get nonce")
	}

	gasPrice, err := txCon.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, stacktrace.Propagate(err, "could not get gass price")
	}

	auth := bind.NewKeyedTransactor(txCon.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
