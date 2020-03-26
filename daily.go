package corona

// TODO: The region funcs probably need to be scoped by country.

import (
	"strings"
)

const (
	// DailyURL is the URL where daily data is pulled.
	DailyURL = "https://open-covid-19.github.io/data/data_latest.json"
)

// DailyWorldwide returns all known worldwide cases.
func DailyWorldwide() ([]*Cases, error) {
	return getCases(DailyURL)
}

// DailyByCountryName returns all known cases by country; case insensitive.
func DailyByCountryName(country string) ([]*Cases, error) {
	cases, err := DailyWorldwide()
	if err != nil {
		return nil, err
	}

	return filterCountryName(cases, country)
}

func DailyByCountryCode(code string) ([]*Cases, error) {
	cases, err := DailyWorldwide()
	if err != nil {
		return nil, err
	}

	return filterCountryCode(cases, code)
}

// DailyByRegionName returns all known cases by region (province/state); case insensitive.
func DailyByRegionName(region string) (*Cases, error) {
	cases, err := DailyWorldwide()
	if err != nil {
		return nil, err
	}

	for _, c := range cases {
		if strings.EqualFold(c.Region.Name, region) {
			return c, nil
		}
	}

	return nil, ErrorNoCasesFound
}

func DailyByRegionCode(code string) (*Cases, error) {
	cases, err := DailyWorldwide()
	if err != nil {
		return nil, err
	}

	for _, c := range cases {
		if strings.EqualFold(c.Region.Code, code) {
			return c, nil
		}
	}

	return nil, ErrorNoCasesFound
}
