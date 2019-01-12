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

// ContractConfiguration is
type ContractConfiguration struct {
	Ethereum struct {
		WS struct {
			Address string `json:"address"`
			Port    string `json:"port"`
		} `json:"ws"`
	} `json:"ethereum"`

	Identity struct {
		Contract struct {
			Address string `json:"address"`
		} `json:"contract"`

		Admin struct {
			PrivateKeyHex string `json:"privateKeyHex"`
		} `json:"admin"`
	} `json:"identity"`
}

func initializeContract(viper *viper.Viper) (*PermaChat, *TransactionContext, error) {

	var config ContractConfiguration

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "Could not load contract config")
	}

	clientURI := "ws://" + config.Ethereum.WS.Address + ":" + config.Ethereum.WS.Port

	client, err := ethclient.Dial(clientURI)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("could not connect to ethereum client: %v", clientURI))
	}

	privateKey, err := crypto.HexToECDSA(config.Identity.Admin.PrivateKeyHex)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("could not create private key: %v", config.Identity.Admin.PrivateKeyHex))
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("error casting public key to ECDSA"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contractAddress := common.HexToAddress(config.Identity.Contract.Address)

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
