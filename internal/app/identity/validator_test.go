package main

import (
	"testing"
)

func TestSha3Hex(t *testing.T) {
	address := "0x6A257ff275ca3ca4D238c6502fe538f65D0d3C84"
	provider := "github"
	username := "bdkent"
	token := provider + username + address
	expected := "0xdfdc0fdf09a1a93dd197d7b648b6151dfdf0183e7790778821bcffb91c9b9b39"
	actual, err := sha3Hex(token)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if expected != actual {
		t.Errorf("Expected sha3 hash to work!\nActual: %v\nExpected: %v", actual, expected)
	}
}
