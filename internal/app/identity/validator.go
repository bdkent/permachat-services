package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
	"golang.org/x/crypto/sha3"
)

// QueryResult =
type QueryResult struct {
	actualToken    string
	actualUsername string
	evidenceURI    string
}

// NewQueryResult =
func NewQueryResult(token string, name string, uri string) *QueryResult {
	return &QueryResult{actualToken: token, actualUsername: name, evidenceURI: uri}
}

// Validator =
type Validator interface {
	initialize(config *viper.Viper) (bool, error)
	query(expectedUsername string, identifier string) (*QueryResult, error)
}

func sha3Hex(input string) (string, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", stacktrace.Propagate(err, fmt.Sprintf("unable to hash: %v", input))
	}
	var buf []byte
	buf = hash.Sum(buf)
	return "0x" + hex.EncodeToString(buf), nil
}

func validate(validator Validator, address string, provider string, username string, identitifer string) (bool, error) {
	log.Printf("validate(%#v, %#v, %#v, %#v)\n", address, provider, username, identitifer)
	expectedToken, err := sha3Hex(provider + username + address)
	if err != nil {
		return false, err
	}
	log.Printf("expectedToken: %v\n", expectedToken)
	result, err := validator.query(username, identitifer)
	if err != nil {
		return false, err
	}
	log.Printf("result: %v\n", result)

	success := result.actualUsername == username && result.actualToken == expectedToken
	log.Printf("success: %v\n", success)

	return success, nil
}

// ValidatorRegistry =
type ValidatorRegistry struct {
	mapping map[string]Validator
	config  *viper.Viper
}

// NewValidatorRegistry =
func NewValidatorRegistry(config *viper.Viper) *ValidatorRegistry {
	return &ValidatorRegistry{mapping: make(map[string]Validator), config: config}
}

// NewValidatorRegistryWithMapping =
func NewValidatorRegistryWithMapping(config *viper.Viper, mapping map[string]Validator) (*ValidatorRegistry, error) {
	registry := NewValidatorRegistry(config)

	for provider, validator := range mapping {
		result, err := registry.register(provider, validator)
		if err != nil {
			return nil, stacktrace.Propagate(err, fmt.Sprintf("error during registration of provider: %v", provider))
		}
		if !result {
			return nil, stacktrace.Propagate(err, fmt.Sprintf("unable to register provider: %v", provider))
		}
	}

	return registry, nil
}

func (registry ValidatorRegistry) register(provider string, validator Validator) (bool, error) {
	registry.mapping[provider] = validator
	return validator.initialize(registry.config)
}

func (registry ValidatorRegistry) validate(address string, provider string, username string, identitifer string) (bool, error) {

	validator := registry.mapping[provider]
	if validator == nil {
		return false, fmt.Errorf("unknown provider: %v", provider)
	}

	return validate(validator, address, provider, username, identitifer)
}
