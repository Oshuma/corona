package corona

import (
	"strings"
)

const (
	// DailyURL is the URL where daily data is pulled.
	DailyURL = "https://open-covid-19.github.io/data/data_latest.json"
)

// DailyWorldwide returns all known worldwide cases.
func DailyWorldwide() ([]*Cases, error) {
	cases, err := getCases(DailyURL)
	if err != nil {
		return nil, err
	}

	return cases, nil
}

// DailyByCountry returns all known cases by country; case insensitive.
func DailyByCountry(country string) ([]*Cases, error) {
	cases, err := getCases(DailyURL)
	if err != nil {
		return nil, err
	}

	return filterCountry(cases, country)
}

// DailyByRegion returns all known cases by region (province/state); case insensitive.
func DailyByRegion(region string) (*Cases, error) {
	cases, err := getCases(DailyURL)
	if err != nil {
		return nil, err
	}

	for _, c := range cases {
		if strings.EqualFold(c.RegionName, region) {
			return c, nil
		}
	}

	return nil, ErrorNoCasesFound
}
