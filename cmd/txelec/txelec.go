package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zer0x0labs/txelec"
)

func main() {
	err := txelec.LoadConfiguration()
	if err != nil {
		logrus.Fatal(err)
	}

	prices, err := txelec.LatestPrices()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%+v", prices)
}
