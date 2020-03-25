package corona

import (
	"strings"
)

const (
	// HistoricalURL is the URL where historical data is pulled.
	HistoricalURL = "https://open-covid-19.github.io/data/data.json"
)

// TimeSeries returns a time series of all reported cases.
func TimeSeries() ([]*Cases, error) {
	return getCases(HistoricalURL)
}

// TimeSeriesByCountry returns a time series of all reported cases in the given country.
func TimeSeriesByCountry(country string) ([]*Cases, error) {
	cases, err := TimeSeries()
	if err != nil {
		return nil, err
	}

	return filterCountry(cases, country)
}

// TimeSeriesByRegion returns a time series of all reported cases in the given region (province/state).
func TimeSeriesByRegion(region string) ([]*Cases, error) {
	cases, err := TimeSeries()
	if err != nil {
		return nil, err
	}

	byRegion := []*Cases{}

	for _, c := range cases {
		if strings.EqualFold(c.RegionName, region) {
			byRegion = append(byRegion, c)
		}
	}

	if len(byRegion) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byRegion, nil
}
