package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the structure for storing configuration.
type Config struct {
	Local  string
	Remote string
	Key    string
}

// LoadConfig is a function that load the config file.
func LoadConfig(path string) *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &Config{
		Local:  viper.GetString("address.local"),
		Remote: viper.GetString("address.remote"),
		Key:    viper.GetString("crypto.key"),
	}
}
