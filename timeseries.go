package corona

import (
	"sort"
	"time"
)

const (
	// HistoricalURL is the URL where historical data is pulled.
	HistoricalURL = "https://open-covid-19.github.io/data/data.json"
)

// TimeSeriesCases is a date (Time) keyed collection of Reports.
type TimeSeriesCases map[time.Time][]*Report

// SortedDates returns an array of sorted time.Time elements; can be used to iterate over the series chronologically.
//
//	for _, date := range ts.SortedDates() {
//	  for _, report := range ts[date] {
//	    fmt.Printf("%+v\n", report)
//	  }
//	}
func (ts TimeSeriesCases) SortedDates() []time.Time {
	dates := make([]time.Time, len(ts))
	for d := range ts {
		dates = append(dates, d)
	}

	sort.Sort(SortDate(dates))
	return dates
}

// TimeSeries returns a time series of all reported cases.
func TimeSeries() (TimeSeriesCases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	return buildTimeSeriesCases(cases), nil
}

// TimeSeriesByCountryName returns a time series of all reported cases in the given country.
func TimeSeriesByCountryName(country string) (TimeSeriesCases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	cases, err = cases.FilterCountryName(country)
	if err != nil {
		return nil, err
	}

	return buildTimeSeriesCases(cases), nil
}

// TimeSeriesByCountryCode returns a time series of all reported cases in the given country code.
func TimeSeriesByCountryCode(code string) (TimeSeriesCases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	cases, err = cases.FilterCountryCode(code)
	if err != nil {
		return nil, err
	}

	return buildTimeSeriesCases(cases), nil
}

// TimeSeriesByRegionName returns a time series of all reported cases in the given region (province/state).
func TimeSeriesByRegionName(region string) (TimeSeriesCases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	cases, err = cases.FilterRegionName(region)
	if err != nil {
		return nil, err
	}

	return buildTimeSeriesCases(cases), nil
}

// TimeSeriesByRegionCode returns a time series of all reported cases in the given region (province/state) code.
func TimeSeriesByRegionCode(code string) (TimeSeriesCases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	cases, err = cases.FilterRegionCode(code)
	if err != nil {
		return nil, err
	}

	return buildTimeSeriesCases(cases), nil
}

func buildTimeSeriesCases(cases Cases) TimeSeriesCases {
	ts := TimeSeriesCases{}

	for _, r := range cases {
		ts[r.Date] = append(ts[r.Date], r)
	}

	return ts
}
