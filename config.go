package txelec

import (
	"github.com/spf13/viper"
)

// LoadConfiguration loads the configuration
func LoadConfiguration() error {
	viper.SetConfigName("txelec")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

// LMPIndexURL returns the LMP Report Index
func LMPIndexURL() string {
	return viper.GetString("sources.lmp_index")
}

// PriceAddersIndexURL returns the LMP Report Index
func PriceAddersIndexURL() string {
	return viper.GetString("sources.adders_index")
}

// SettlementIndexURL returns the Settlement Report Index
func SettlementIndexURL() string {
	return viper.GetString("sources.settlement_index")
}
