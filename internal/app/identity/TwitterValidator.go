package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
)

type twitterValidator struct {
	bearerToken string
}

func newTwitterValidator() Validator {
	return twitterValidator{bearerToken: ""}
}

// TwitterValidatorParams =
type TwitterValidatorParams struct {
	APIKey       string `json:"apiKey"`
	APISecretKey string `json:"apiSecretKey"`
}

func (v twitterValidator) initialize(config *viper.Viper) (bool, error) {
	var params TwitterValidatorParams
	err := config.Unmarshal(&params)
	if err != nil {
		return false, stacktrace.Propagate(err, fmt.Sprintf("unable to obtain config info for Twitter initialization: %v", config))
	}

	encodedKey, _ := url.Parse(params.APIKey)
	encodedSecret, _ := url.Parse(params.APISecretKey)
	token := encodedKey.String() + ":" + encodedSecret.String()

	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))

	client := http.DefaultClient
	request, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return false, stacktrace.Propagate(err, fmt.Sprintf("unable to create Twitter API request"))
	}

	request.Header.Add("Authorization", "Basic "+encodedToken)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(request)
	if err != nil {
		return false, stacktrace.Propagate(err, fmt.Sprintf("unable to obtain Twitter access token"))
	}
	if resp.StatusCode != http.StatusOK {
		return false, stacktrace.Propagate(err, fmt.Sprintf("bad status code while trying to obtain Twitter access token: %v", resp.StatusCode))
	}

	var result map[string]string
	err1 := json.NewDecoder(resp.Body).Decode(&result)
	if err1 != nil {
		return false, stacktrace.Propagate(err, fmt.Sprintf("unable to decode Twitter access token response"))
	}

	v.bearerToken = string(result["access_token"])

	return true, nil
}

type tweetResult struct {
	User struct {
		ScreenName string `json:"screen_name"`
	}
	Text string `json:"text"`
}

func (v twitterValidator) query(expectedUsername string, identifier string) (*QueryResult, error) {

	request, err := http.NewRequest("GET", "https://api.twitter.com/1.1/statuses/show.json?id="+identifier, nil)
	if err != nil {
		return nil, stacktrace.Propagate(err, fmt.Sprintf("unable to create Twitter API request"))
	}

	request.Header.Add("Authorization", "Bearer "+v.bearerToken)

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, stacktrace.Propagate(err, fmt.Sprintf("unable to request tweet: %v", identifier))
	}

	var body tweetResult
	err1 := json.NewDecoder(response.Body).Decode(&body)
	if err1 != nil {
		return nil, stacktrace.Propagate(err, fmt.Sprintf("unable to decode Twitter API response"))
	}

	token := body.Text
	actualUserName := body.User.ScreenName
	evidenceURI := "https://twitter.com/" + actualUserName + "/status/" + identifier
	result := NewQueryResult(token, actualUserName, evidenceURI)

	return result, nil
}
