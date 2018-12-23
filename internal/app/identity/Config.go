package main

import (
	"log"

	"github.com/spf13/viper"
)

func safeSub(v *viper.Viper, path string) *viper.Viper {
	if !v.IsSet(path) {
		return viper.New()
	}

	return v.Sub(path)
}

func initializeConfig(appName string) *viper.Viper {

	viper := viper.New()

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/" + appName + "/")
	viper.AddConfigPath("$HOME/." + appName)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper
}
