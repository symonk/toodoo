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
	config.SetConfigName("default")
	config.AddConfigPath("config/")
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("unable to load the configuration for environment %s", environment))
	}
	config.MergeConfigMap(config.AllSettings())
}

func GetConfig() *viper.Viper {
	return config
}
