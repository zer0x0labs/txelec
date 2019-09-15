package txelec

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Price is a price per MWh
type Price struct {
	PriceMWh     float64
	RepeatedHour bool
	Timestamp    time.Time
	Region       string
	Adder        Adder
}

// Adder is a price adder per MWh
type Adder struct {
	RTORDPA      float64
	RTORPA       float64
	All          float64
	RepeatedHour bool
	Timestamp    time.Time
}

// Prices converts the raw k,v data to a list of prices
func Prices(prices []map[string]string) ([]Price, error) {
	p := make([]Price, 0)
	for _, line := range prices {
		export := false
		t := Price{}
		for k, v := range line {
			switch strings.ToLower(k) {
			case "lmp":
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					continue
				}
				export = true
				t.PriceMWh = f
			case "repeatedhourflag":
				t.RepeatedHour = parseRepeatedHourFlag(v)
			case "scedtimestamp":
				ts, err := parseTimestamp(v)
				if err != nil {
					return p, err
				}
				t.Timestamp = ts
			case "settlementpoint":
				t.Region = v
			}
		}
		if export {
			p = append(p, t)
		}
	}
	return p, nil
}

// Adders converts raw k,v data to a list of adders
func Adders(adders []map[string]string) ([]Adder, error) {
	ad := make([]Adder, 0)
	for _, line := range adders {
		export := false
		a := Adder{}
		for k, v := range line {
			switch strings.ToLower(k) {
			case "rtorpa":
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					continue
				}
				export = true
				a.RTORPA = f
				a.All += f
			case "rtordpa":
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					continue
				}
				export = true
				a.RTORDPA = f
				a.All += f
			case "repeatedhourflag":
				a.RepeatedHour = parseRepeatedHourFlag(v)
			case "scedtimestamp":
				ts, err := parseTimestamp(v)
				if err != nil {
					return ad, err
				}
				a.Timestamp = ts
			}
		}
		if export {
			ad = append(ad, a)
		}
	}
	return ad, nil
}

func parseRepeatedHourFlag(v string) bool {
	if strings.ToLower(v) == "n" {
		return false
	}
	return true
}

func parseTimestamp(v string) (time.Time, error) {
	var ts time.Time
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return ts, err
	}

	ts, err = time.ParseInLocation("01/02/2006 15:04:05", v, loc)
	if err != nil {
		return ts, err
	}

	return ts, err
}

// PricesWithAdders converts a list of k,v prices and adders into prices with matching adders
func PricesWithAdders(prices, adders []map[string]string) ([]Price, error) {
	a, err := Adders(adders)
	if err != nil {
		return nil, err
	}
	p, err := Prices(prices)
	if err != nil {
		return nil, err
	}
	for i, v := range p {
		for _, av := range a {
			if av.Timestamp.Hour() == v.Timestamp.Hour() && av.Timestamp.Minute() == v.Timestamp.Minute() {
				p[i].Adder = av
			} else {
				return nil, fmt.Errorf("price timestamp %+v and adder timestamp %+v do not match", v.Timestamp, av.Timestamp)
			}
		}
	}
	return p, nil
}
