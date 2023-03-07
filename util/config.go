package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DriverName     string `mapstructure:"DB_DRIVER"`
	DataSourceName string `mapstructure:"DB_SOURCE"`
	Address        string `mapstructure:"SERVER_ADDRESS"`
}

func ParseToConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
