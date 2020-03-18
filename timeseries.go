package corona

import (
	"net/http"
	"strconv"
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

// TimeSeriesConfirmedByCountry returns a time series of all confirmed cases in the given country.
func TimeSeriesConfirmedByCountry(country string) (*TimeSeries, error) {
	confirmed, err := timeSeriesForURL(ConfirmedURL)
	if err != nil {
		return nil, err
	}

	for _, ts := range confirmed {
		if ts.CountryRegion == country {
			return ts, nil
		}
	}

	return nil, ErrorNoCasesFound
}

// TimeSeriesConfirmedByState returns a time series of all confirmed cases in the given state.
func TimeSeriesConfirmedByState(state string) (*TimeSeries, error) {
	return confirmedByProvinceOrState(state)
}

// TimeSeriesConfirmedByProvince returns a time series of all confirmed cases in the given province.
func TimeSeriesConfirmedByProvince(province string) (*TimeSeries, error) {
	return confirmedByProvinceOrState(province)
}

// TimeSeriesDeaths returns a time series of all deaths.
func TimeSeriesDeaths() ([]*TimeSeries, error) {
	return timeSeriesForURL(DeathsURL)
}

// TimeSeriesDeathsByCountry returns a time series of all deaths in the given country.
func TimeSeriesDeathsByCountry(country string) (*TimeSeries, error) {
	deaths, err := timeSeriesForURL(DeathsURL)
	if err != nil {
		return nil, err
	}

	for _, ts := range deaths {
		if ts.CountryRegion == country {
			return ts, nil
		}
	}

	return nil, ErrorNoCasesFound
}

// TimeSeriesDeathsByState returns a time series of all deaths in the given state.
func TimeSeriesDeathsByState(state string) (*TimeSeries, error) {
	return deathsByProvinceOrState(state)
}

// TimeSeriesDeathsByProvince returns a time series of all deaths in the given province.
func TimeSeriesDeathsByProvince(province string) (*TimeSeries, error) {
	return deathsByProvinceOrState(province)
}

// TimeSeriesRecovered returns a time series of all recovered cases.
func TimeSeriesRecovered() ([]*TimeSeries, error) {
	return timeSeriesForURL(RecoveredURL)
}

// TimeSeriesRecoveredByCountry returns a time series of all recovered cases in the given country.
func TimeSeriesRecoveredByCountry(country string) (*TimeSeries, error) {
	recovered, err := timeSeriesForURL(RecoveredURL)
	if err != nil {
		return nil, err
	}

	for _, ts := range recovered {
		if ts.CountryRegion == country {
			return ts, nil
		}
	}

	return nil, ErrorNoCasesFound
}

// TimeSeriesRecoveredByState returns a time series of all recovered cases in the given state.
func TimeSeriesRecoveredByState(state string) (*TimeSeries, error) {
	return recoveredByProvinceOrState(state)
}

// TimeSeriesRecoveredByProvince returns a time series of all recovered cases in the given province.
func TimeSeriesRecoveredByProvince(province string) (*TimeSeries, error) {
	return recoveredByProvinceOrState(province)
}

func confirmedByProvinceOrState(ps string) (*TimeSeries, error) {
	confirmed, err := timeSeriesForURL(ConfirmedURL)
	if err != nil {
		return nil, err
	}

	for _, ts := range confirmed {
		if ts.ProvinceState == ps {
			return ts, nil
		}
	}

	return nil, ErrorNoCasesFound
}

func deathsByProvinceOrState(ps string) (*TimeSeries, error) {
	deaths, err := timeSeriesForURL(DeathsURL)
	if err != nil {
		return nil, err
	}

	for _, ts := range deaths {
		if ts.ProvinceState == ps {
			return ts, nil
		}
	}

	return nil, ErrorNoCasesFound
}

func recoveredByProvinceOrState(ps string) (*TimeSeries, error) {
	recovered, err := timeSeriesForURL(RecoveredURL)
	if err != nil {
		return nil, err
	}

	for _, ts := range recovered {
		if ts.ProvinceState == ps {
			return ts, nil
		}
	}

	return nil, ErrorNoCasesFound
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
