package corona

const (
	// HistoricalURL is the URL where historical data is pulled.
	HistoricalURL = "https://open-covid-19.github.io/data/data.json"
)

// TimeSeries returns a time series of all reported cases.
func TimeSeries() ([]*Cases, error) {
	return getCases(HistoricalURL)
}

// TimeSeriesByCountryName returns a time series of all reported cases in the given country.
func TimeSeriesByCountryName(country string) ([]*Cases, error) {
	cases, err := TimeSeries()
	if err != nil {
		return nil, err
	}

	return filterCountryName(cases, country)
}

func TimeSeriesByCountryCode(code string) ([]*Cases, error) {
	cases, err := TimeSeries()
	if err != nil {
		return nil, err
	}

	return filterCountryCode(cases, code)
}

// TimeSeriesByRegionName returns a time series of all reported cases in the given region (province/state).
func TimeSeriesByRegionName(region string) ([]*Cases, error) {
	cases, err := TimeSeries()
	if err != nil {
		return nil, err
	}

	return filterRegionName(cases, region)
}

func TimeSeriesByRegionCode(code string) ([]*Cases, error) {
	cases, err := TimeSeries()
	if err != nil {
		return nil, err
	}

	return filterRegionCode(cases, code)
}
