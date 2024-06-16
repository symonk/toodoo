package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init loads the configuration for a given environment and provides
// a global Viper instance.
func Init(environment string) {
	config = viper.New()
	config.SetConfigType("yaml")
	base, _ := os.Executable()
	dir := filepath.Dir(base) + "/internal/config/"
	config.AddConfigPath(dir)
	config.SetConfigFile("dev")
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("unable to load the configuration for environment %s, %s", environment, err.Error()))
	}
	config.MergeConfigMap(config.AllSettings())
}

func GetConfig() *viper.Viper {
	return config
}
