package main

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/viper"
)

func Test_fromDir(t *testing.T) {

	dir, _ := ioutil.TempDir("", "config-dir-")
	ioutil.WriteFile(dir+"/x", []byte("hello"), 0644)
	ioutil.WriteFile(dir+"/a.b", []byte("world"), 0644)

	v := viper.New()
	fromDir(v, dir)

	if !v.IsSet("x") {
		t.Errorf("expected simple key to exist")
	}

	if v.GetString("x") != "hello" {
		t.Errorf("expected simple key to retrieve value")
	}

	if !v.IsSet("a.b") && v.Sub("a") != nil && v.Sub("a").IsSet("b") {
		t.Errorf("expected key path to exist")
	}

	if v.GetString("a.b") != "world" && v.Sub("a").GetString("b") != "world" {
		t.Errorf("expected key path to retrieve value")
	}

	_, err := fromDir(v, "/potato")
	if err == nil {
		t.Error("expected bad directory to return error")
	}

}
