package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zer0x0labs/txelec"
)

func main() {
	err := txelec.LoadConfiguration()
	if err != nil {
		logrus.Fatal(err)
	}

	url, err := txelec.GetLatestDLURL(viper.GetString("sources.lmp_index"))
	if err != nil {
		logrus.Fatal(err)
	}

	prices, err := txelec.GetDataFromZippedCSV(url.URL)
	if err != nil {
		logrus.Fatal(err)
	}

	url, err = txelec.GetLatestDLURL(viper.GetString("sources.adders_index"))
	if err != nil {
		logrus.Fatal(err)
	}

	adders, err := txelec.GetDataFromZippedCSV(url.URL)
	if err != nil {
		logrus.Fatal(err)
	}

	pa := txelec.PricesWithAdders(prices, adders)
	logrus.Info(pa)

}
