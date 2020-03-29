package corona

const (
	// HistoricalURL is the URL where historical data is pulled.
	HistoricalURL = "https://open-covid-19.github.io/data/data.json"
)

// TimeSeries returns a time series of all reported cases.
func TimeSeries() (Cases, error) {
	return getReports(HistoricalURL)
}

// TimeSeriesByCountryName returns a time series of all reported cases in the given country.
func TimeSeriesByCountryName(country string) (Cases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterCountryName(country)
}

// TimeSeriesByCountryCode returns a time series of all reported cases in the given country code.
func TimeSeriesByCountryCode(code string) (Cases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterCountryCode(code)
}

// TimeSeriesByRegionName returns a time series of all reported cases in the given region (province/state).
func TimeSeriesByRegionName(region string) (Cases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterRegionName(region)
}

// TimeSeriesByRegionCode returns a time series of all reported cases in the given region (province/state) code.
func TimeSeriesByRegionCode(code string) (Cases, error) {
	cases, err := getReports(HistoricalURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterRegionCode(code)
}
