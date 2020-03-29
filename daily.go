package corona

const (
	// DailyURL is the URL where daily data is pulled.
	DailyURL = "https://open-covid-19.github.io/data/data_latest.json"
)

// DailyWorldwide returns all known worldwide reports.
func DailyWorldwide() (Cases, error) {
	cases, err := getReports(DailyURL)
	if err != nil {
		return nil, err
	}

	return cases, nil
}

// DailyByCountryName returns all known reports by country; case insensitive.
func DailyByCountryName(country string) (Cases, error) {
	cases, err := getReports(DailyURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterCountryName(country)
}

// TotalByCountryName returns a single cumulative Report for the given country name; case insensitive.
func TotalByCountryName(country string) (*Report, error) {
	cases, err := DailyByCountryName(country)
	if err != nil {
		return nil, err
	}

	total, err := cases.FilterRegionCode("")
	if err != nil {
		return nil, err
	}

	return total[0], nil
}

// DailyByCountryCode returns all known reports by country code; case insensitive.
func DailyByCountryCode(code string) (Cases, error) {
	cases, err := getReports(DailyURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterCountryCode(code)
}

// TotalByCountryCode returns a single cumulative Report for the given country name; case insensitive.
func TotalByCountryCode(code string) (*Report, error) {
	cases, err := DailyByCountryCode(code)
	if err != nil {
		return nil, err
	}

	total, err := cases.FilterRegionCode("")
	if err != nil {
		return nil, err
	}

	return total[0], nil
}

// DailyByRegionName returns all known reports by region (province/state); case insensitive.
func DailyByRegionName(region string) (Cases, error) {
	cases, err := getReports(DailyURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterRegionName(region)
}

// DailyByRegionCode returns all known reports by region code; case insensitive.
func DailyByRegionCode(code string) (Cases, error) {
	cases, err := getReports(DailyURL)
	if err != nil {
		return nil, err
	}

	return cases.FilterRegionCode(code)
}
