package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
)

type gitHubValidator struct {
}

func newGitHubValidator() Validator {
	return gitHubValidator{}
}

func (v gitHubValidator) initialize(config *viper.Viper) (bool, error) {
	return true, nil
}

func (v gitHubValidator) query(expectedUsername string, identifier string) (*QueryResult, error) {

	evidenceURI := "https://gist.githubusercontent.com/" + expectedUsername + "/" + identifier

	resp, err := http.Get(evidenceURI + "/raw")
	if err != nil {
		return nil, stacktrace.Propagate(err, fmt.Sprintf("unable to query GitHub for evidence URI: %v", evidenceURI))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	token := buf.String()

	result := NewQueryResult(token, expectedUsername, evidenceURI)

	return result, nil
}
