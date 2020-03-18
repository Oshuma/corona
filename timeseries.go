package corona

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

const (
	// ConfirmedURL is the URL where confirmed time series data is pulled.
	ConfirmedURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Confirmed.csv"

	// DeathsURL is the URL where deaths time series data is pulled.
	DeathsURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Deaths.csv"

	// RecoveredURL is the URL where recovered time series data is pulled.
	RecoveredURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Recovered.csv"
)

// DateData stores a single time series date and count.
type DateData struct {
	Date  time.Time
	Count int
}

// TimeSeries stores a single entry in the time series data.
type TimeSeries struct {
	ProvinceState string
	CountryRegion string
	Latitude      float64
	Longitude     float64
	Dates         []DateData
}

// TimeSeriesConfirmed returns a time series of all confirmed cases.
func TimeSeriesConfirmed() ([]*TimeSeries, error) {
	return timeSeriesForURL(ConfirmedURL)
}

// TimeSeriesConfirmedByCountry returns a time series of all confirmed cases in the given country; case insensitive.
func TimeSeriesConfirmedByCountry(country string) ([]*TimeSeries, error) {
	all, err := timeSeriesForURL(ConfirmedURL)
	if err != nil {
		return nil, err
	}

	confirmed := []*TimeSeries{}

	for _, ts := range all {
		if strings.EqualFold(ts.CountryRegion, country) {
			confirmed = append(confirmed, ts)
		}
	}

	if len(confirmed) == 0 {
		return nil, ErrorNoCasesFound
	}

	return confirmed, nil
}

// TimeSeriesConfirmedByState returns a time series of all confirmed cases in the given state; case insensitive.
func TimeSeriesConfirmedByState(state string) ([]*TimeSeries, error) {
	return confirmedByProvinceOrState(state)
}

// TimeSeriesConfirmedByProvince returns a time series of all confirmed cases in the given province; case insensitive.
func TimeSeriesConfirmedByProvince(province string) ([]*TimeSeries, error) {
	return confirmedByProvinceOrState(province)
}

// TimeSeriesDeaths returns a time series of all deaths.
func TimeSeriesDeaths() ([]*TimeSeries, error) {
	return timeSeriesForURL(DeathsURL)
}

// TimeSeriesDeathsByCountry returns a time series of all deaths in the given country; case insensitive.
func TimeSeriesDeathsByCountry(country string) ([]*TimeSeries, error) {
	all, err := timeSeriesForURL(DeathsURL)
	if err != nil {
		return nil, err
	}

	deaths := []*TimeSeries{}

	for _, ts := range all {
		if strings.EqualFold(ts.CountryRegion, country) {
			deaths = append(deaths, ts)
		}
	}

	if len(deaths) == 0 {
		return nil, ErrorNoCasesFound
	}

	return deaths, nil
}

// TimeSeriesDeathsByState returns a time series of all deaths in the given state; case insensitive.
func TimeSeriesDeathsByState(state string) ([]*TimeSeries, error) {
	return deathsByProvinceOrState(state)
}

// TimeSeriesDeathsByProvince returns a time series of all deaths in the given province; case insensitive.
func TimeSeriesDeathsByProvince(province string) ([]*TimeSeries, error) {
	return deathsByProvinceOrState(province)
}

// TimeSeriesRecovered returns a time series of all recovered cases.
func TimeSeriesRecovered() ([]*TimeSeries, error) {
	return timeSeriesForURL(RecoveredURL)
}

// TimeSeriesRecoveredByCountry returns a time series of all recovered cases in the given country; case insensitive.
func TimeSeriesRecoveredByCountry(country string) ([]*TimeSeries, error) {
	all, err := timeSeriesForURL(RecoveredURL)
	if err != nil {
		return nil, err
	}

	recovered := []*TimeSeries{}

	for _, ts := range all {
		if strings.EqualFold(ts.CountryRegion, country) {
			recovered = append(recovered, ts)
		}
	}

	if len(recovered) == 0 {
		return nil, ErrorNoCasesFound
	}

	return recovered, nil
}

// TimeSeriesRecoveredByState returns a time series of all recovered cases in the given state; case insensitive.
func TimeSeriesRecoveredByState(state string) ([]*TimeSeries, error) {
	return recoveredByProvinceOrState(state)
}

// TimeSeriesRecoveredByProvince returns a time series of all recovered cases in the given province; case insensitive.
func TimeSeriesRecoveredByProvince(province string) ([]*TimeSeries, error) {
	return recoveredByProvinceOrState(province)
}

func confirmedByProvinceOrState(ps string) ([]*TimeSeries, error) {
	all, err := timeSeriesForURL(ConfirmedURL)
	if err != nil {
		return nil, err
	}

	confirmed := []*TimeSeries{}

	for _, ts := range all {
		if strings.EqualFold(ts.ProvinceState, ps) {
			confirmed = append(confirmed, ts)
		}
	}

	if len(confirmed) == 0 {
		return nil, ErrorNoCasesFound
	}

	return confirmed, nil
}

func deathsByProvinceOrState(ps string) ([]*TimeSeries, error) {
	all, err := timeSeriesForURL(DeathsURL)
	if err != nil {
		return nil, err
	}

	deaths := []*TimeSeries{}

	for _, ts := range all {
		if strings.EqualFold(ts.ProvinceState, ps) {
			deaths = append(deaths, ts)
		}
	}

	if len(deaths) == 0 {
		return nil, ErrorNoCasesFound
	}

	return deaths, nil
}

func recoveredByProvinceOrState(ps string) ([]*TimeSeries, error) {
	all, err := timeSeriesForURL(RecoveredURL)
	if err != nil {
		return nil, err
	}

	recovered := []*TimeSeries{}

	for _, ts := range all {
		if strings.EqualFold(ts.ProvinceState, ps) {
			recovered = append(recovered, ts)
		}
	}

	if len(recovered) == 0 {
		return nil, ErrorNoCasesFound
	}

	return recovered, nil
}

func timeSeriesForURL(url string) ([]*TimeSeries, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rows, err := gocsv.CSVToMaps(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []*TimeSeries
	for _, r := range rows {
		var ts TimeSeries
		for k, v := range r {
			switch k {
			case "Province/State":
				ts.ProvinceState = v
			case "Country/Region":
				ts.CountryRegion = v
			case "Lat":
				lat, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, err
				}
				ts.Latitude = lat
			case "Long":
				long, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, err
				}
				ts.Longitude = long
			default:
				date, err := time.Parse("1/2/06", k)
				if err != nil {
					return nil, err
				}

				count := 0
				if v != "" {
					count, err = strconv.Atoi(v)
					if err != nil {
						return nil, err
					}
				}

				d := DateData{Date: date, Count: count}
				ts.Dates = append(ts.Dates, d)
			}
		}

		data = append(data, &ts)
	}

	return data, nil
}
