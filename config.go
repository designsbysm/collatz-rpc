package main

import (
	"errors"

	"github.com/spf13/viper"
)

func config() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./config.yaml not found")
		} else {
			return err
		}
	}

	return nil
}
