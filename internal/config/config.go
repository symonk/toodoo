package config

var config *Config

func Init() {

}

type Config struct {
}

func GetConfig() *Config {
	return config
}
