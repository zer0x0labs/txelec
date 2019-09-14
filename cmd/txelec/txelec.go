package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zer0x0labs/txelec"
)

func main() {
	err := txelec.LoadConfiguration()
	if err != nil {
		logrus.Fatal(err)
	}

	prices := txelec.LatestPrices()
	region := txelec.DefaultRegion()

	pricemap := make(map[string]txelec.Price)

	for _, price := range prices {
		pricemap[price.Region] = price
	}

	if len(os.Args) > 2 {
		if _, ok := pricemap[os.Args[1]]; ok {
			region = os.Args[1]
		}
	}

	fmt.Printf("%s %.2f c/kWh", region, (pricemap[region].PriceMWh+pricemap[region].Adder.All)/10)
}
