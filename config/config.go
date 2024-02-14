package config

import (
	"errors"

	"github.com/spf13/viper"
)

var (
	ZConfig = zConfig{}
)

type zConfig struct {
}

func Init(cfgFileName string) {
	viper.SetConfigFile(cfgFileName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("fatal error to read config file"))
	}

	viper.WatchConfig()
}

func (c *zConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (c *zConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (c *zConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (c *zConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}
