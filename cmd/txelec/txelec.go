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

	a, err := txelec.NewAPI()
	if err != nil {
		logrus.Fatal(err)
	}

	a.Start()

}
