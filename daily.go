package corona

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// DailyURL is the URL where data is pulled.
	DailyURL = "https://open-covid-19.github.io/data/data_latest.json"
)

// DailyWorldwide returns all known worldwide cases.
func DailyWorldwide() ([]*Cases, error) {
	cases, err := getCases()
	if err != nil {
		return nil, err
	}

	return cases, nil
}

// DailyByCountry returns all known cases by country; case insensitive.
func DailyByCountry(country string) ([]*Cases, error) {
	cases, err := getCases()
	if err != nil {
		return nil, err
	}

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

// DailyByRegion returns all known cases by region (province/state); case insensitive.
func DailyByRegion(region string) (*Cases, error) {
	cases, err := getCases()
	if err != nil {
		return nil, err
	}

	for _, c := range cases {
		if strings.EqualFold(c.RegionName, region) || strings.EqualFold(c.RegionCode, region) {
			return c, nil
		}
	}

	return nil, ErrorNoCasesFound
}

func getCases() ([]*Cases, error) {
	req, err := http.NewRequest("GET", DailyURL, nil)
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
