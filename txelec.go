package txelec

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LatestPrices returns the latest prices with adders
func LatestPrices() ([]Price, error) {
	url, err := GetLatestDLURL(viper.GetString("sources.lmp_index"))
	if err != nil {
		logrus.Fatal(err)
	}

	prices, err := GetDataFromZippedCSV(url.URL)
	if err != nil {
		logrus.Fatal(err)
	}

	url, err = GetLatestDLURL(viper.GetString("sources.adders_index"))
	if err != nil {
		logrus.Fatal(err)
	}

	adders, err := GetDataFromZippedCSV(url.URL)
	if err != nil {
		logrus.Fatal(err)
	}

	return PricesWithAdders(prices, adders)
}
