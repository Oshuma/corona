package corona

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Cases stores case information.
type Cases struct {
	Date        Date       `json:"Date"`
	CountryCode string     `json:"CountryCode"`
	CountryName string     `json:"CountryName"`
	RegionCode  string     `json:"RegionCode"`
	RegionName  string     `json:"RegionName"`
	Confirmed   Confirmed  `json:"Confirmed"`
	Deaths      Deaths     `json:"Deaths"`
	Latitude    Latitude   `json:"Latitude"`
	Longitude   Longitude  `json:"Longitude"`
	Population  Population `json:"Population"`
}

// Date is a time.Time wrapper used to unmarshal and parse the JSON response.
type Date struct {
	time.Time
}

// UnmarshalJSON unmarshals the time format in the JSON response.
func (d Date) UnmarshalJSON(input []byte) error {
	s, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	d.Time = t
	return err
}

type stringToInt int

func (si stringToInt) UnmarshalJSON(input []byte) error {
	if string(input) == "null" {
		return nil
	}

	s, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	si = stringToInt(i)
	return nil
}

// Confirmed is an int wrapper used to unmarshal from a JSON string.
type Confirmed struct {
	stringToInt
}

// Deaths is an int wrapper used to unmarshal from a JSON string.
type Deaths struct {
	stringToInt
}

// Population is an int wrapper used to unmarshal from a JSON string.
type Population struct {
	stringToInt
}

type stringToFloat64 float64

func (sf stringToFloat64) UnmarshalJSON(input []byte) error {
	if string(input) == "null" {
		return nil
	}

	s, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	sf = stringToFloat64(f)
	return nil
}

// Latitude is a float64 wrapper used to unmarshal from a JSON string.
type Latitude struct {
	stringToFloat64
}

// Longitude is a float64 wrapper used to unmarshal from a JSON string.
type Longitude struct {
	stringToFloat64
}

func getCases(url string) ([]*Cases, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cases := []*Cases{}
	err = json.Unmarshal(content, &cases)
	if err != nil {
		return nil, err
	}

	return cases, nil
}

func filterCountry(cases []*Cases, country string) ([]*Cases, error) {
	byCountry := []*Cases{}

	for _, c := range cases {
		if strings.EqualFold(c.CountryName, country) || strings.EqualFold(c.CountryCode, country) {
			byCountry = append(byCountry, c)
		}
	}

	if len(byCountry) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byCountry, nil
}
