package main

import (
	"os"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func logger() error {
	level := timber.StringToLevel(viper.GetString("timber.cli.level"))
	if level < 1 {
		return nil
	}

	return timber.New(
		os.Stdout,
		level,
		viper.GetString("timber.cli.timestamp"),
		timber.StringToFlags(viper.GetString("timber.cli.flags")),
	)
}
