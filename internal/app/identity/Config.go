package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
)

func safeSub(v *viper.Viper, path string) *viper.Viper {
	if !v.IsSet(path) {
		return viper.New()
	}

	return v.Sub(path)
}

func initializeConfig(appName string) *viper.Viper {

	config := viper.New()

	config.SetConfigName("config")
	config.AddConfigPath("/etc/" + appName + "/")
	config.AddConfigPath("$HOME/." + appName)
	config.AddConfigPath(".")

	config.AutomaticEnv()

	fromDir(config, "/run/secrets")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return config
}

func fromDir(v *viper.Viper, path string) (*viper.Viper, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, stacktrace.Propagate(err, fmt.Sprintf("expected directory of config properties: %v", path))
	}

	for _, f := range files {
		key := f.Name()
		bs, _ := ioutil.ReadFile(path + "/" + key)
		value := strings.TrimSpace(string(bs))
		v.Set(key, value)
	}

	return v, nil
}
