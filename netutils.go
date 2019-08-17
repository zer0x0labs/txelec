package txelec

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetLatestDLURL returns the most recent download URL from Ercot archive page
func GetLatestDLURL(url string) (string, error) {

}

// GetDataFromZippedCSV returns a slice of string mapped strings from a zipped CSV
func GetDataFromZippedCSV(url string) ([]map[string]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	zbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	zr := bytes.NewReader(zbody)

	r, err := zip.NewReader(zr, int64(len(zbody)))
	if err != nil {
		logrus.Fatal(err)
	}

	var recmap []map[string]string

	for _, f := range r.File {
		logrus.Info(f.Name)
		fh, err := f.Open()
		if err != nil {
			return nil, err
		}
		csvr := csv.NewReader(fh)
		recs, err := csvr.ReadAll()
		if err != nil {
			return nil, err
		}

		headers := recs[0]

		for _, row := range recs[1:] {
			rm := make(map[string]string)
			for k, v := range row {
				rm[headers[k]] = v
			}
			recmap = append(recmap, rm)
		}

	}
	return recmap, err
}
