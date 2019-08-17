package txelec

import (
	"github.com/spf13/viper"
)

func LoadConfiguration() error {
	viper.SetConfigName("txelec")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}
