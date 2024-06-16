package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init loads the configuration for a given environment and provides
// a global Viper instance.
func Init(environment string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.AddConfigPath("internal/config/")
	config.SetConfigName(environment)
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file error for env: %s, %w", environment, err))
	}
	if err := config.MergeConfigMap(config.AllSettings()); err != nil {
		panic("unable to merge config")
	}
}

func GetConfig() *viper.Viper {
	return config
}
