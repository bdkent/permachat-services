package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

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

func printEnvKeys() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

func initializeContract(config *viper.Viper) (*PermaChat, *TransactionContext, error) {

	ethereumAddress := config.GetString("permachat_ethereum_ws_address")
	if ethereumAddress == "" {
		printEnvKeys()
		return nil, nil, fmt.Errorf("expected config value for ethereum address")
	}
	ethereumPort := config.GetString("permachat_ethereum_ws_port")
	if ethereumPort == "" {
		printEnvKeys()
		return nil, nil, fmt.Errorf("expected config value for ethereum port")
	}

	privateKeyHex := config.GetString("permachat_identity_admin_privatekeyhex")
	if privateKeyHex == "" {
		printEnvKeys()
		return nil, nil, fmt.Errorf("expected config value for private key")
	}

	contractAddressHex := config.GetString("permachat_identity_contract_address")
	if contractAddressHex == "" {
		printEnvKeys()
		return nil, nil, fmt.Errorf("expected config value for contract address")
	}

	clientURI := fmt.Sprintf("ws://%v:%v", ethereumAddress, ethereumPort)

	client, err := ethclient.Dial(clientURI)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("could not connect to ethereum client: %v", clientURI))
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("could not create private key"))
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, stacktrace.Propagate(err, fmt.Sprintf("error casting public key to ECDSA"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	contractAddress := common.HexToAddress(contractAddressHex)

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
