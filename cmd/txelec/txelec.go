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

	recs, err := txelec.GetData("http://mis.ercot.com/misdownload/servlets/mirDownload?mimic_duns=&doclookupId=674550800")

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(recs)

}
