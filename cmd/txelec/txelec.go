package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zer0x0labs/txelec"
)

func main() {
	err := txelec.LoadConfiguration()
	if err != nil {
		logrus.Fatal(err)
	}

	prices := txelec.LatestPrices()
	logrus.Info(prices)
}
