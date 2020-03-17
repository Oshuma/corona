package corona

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
)

const (
	// BaseDailyURL is the URL where data is pulled; date should be formatted as: MM-DD-YYYY
	BaseDailyURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/%s.csv"
)

// LastUpdate is used as a time.Time wrapper to parse the LastUpdate CSV field.
type LastUpdate struct {
	time.Time
}

// UnmarshalCSV is used to convert the LastUpdate CSV field to a time.Time.
func (lu *LastUpdate) UnmarshalCSV(csv string) error {
	t, err := time.Parse("2006-01-02T15:04:05", csv)
	lu.Time = t.UTC()
	return err
}

// Cases stores case information.
type Cases struct {
	ProvinceState string     `csv:"Province/State"`
	CountryRegion string     `csv:"Country/Region"`
	LastUpdate    LastUpdate `csv:"Last Update"`
	Confirmed     int        `csv:"Confirmed"`
	Deaths        int        `csv:"Deaths"`
	Recovered     int        `csv:"Recovered"`
	Latitude      float64    `csv:"Latitude"`
	Longitude     float64    `csv:"Longitude"`
}

// DailyWorldwide returns all known worldwide cases.
func DailyWorldwide() ([]*Cases, error) {
	cases, err := getCases()
	if err != nil {
		return nil, err
	}

	return cases, nil
}

// DailyByCountry returns all known cases by country.
func DailyByCountry(country string) ([]*Cases, error) {
	cases, err := getCases()
	if err != nil {
		return nil, err
	}

	byCountry := []*Cases{}
	for _, c := range cases {
		if c.CountryRegion == country {
			byCountry = append(byCountry, c)
		}
	}

	if len(byCountry) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byCountry, nil
}

// DailyByState returns all known cases by state.
func DailyByState(state string) ([]*Cases, error) {
	return dailyByProvinceOrState(state)
}

// DailyByProvinceState returns all known cases by province.
func DailyByProvinceState(province string) ([]*Cases, error) {
	return dailyByProvinceOrState(province)
}

func dailyByProvinceOrState(ps string) ([]*Cases, error) {
	cases, err := getCases()
	if err != nil {
		return nil, err
	}

	byPorS := []*Cases{}
	for _, c := range cases {
		if c.ProvinceState == ps {
			byPorS = append(byPorS, c)
		}
	}

	if len(byPorS) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byPorS, nil
}

func getCSV(date time.Time) (*http.Response, error) {
	url := fmt.Sprintf(BaseDailyURL, date.Format("01-02-2006"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func getCases() ([]*Cases, error) {
	date := time.Now()

	csv, err := getCSV(date)
	if err != nil {
		return nil, err
	}
	if csv.StatusCode == http.StatusNotFound {
		date = date.AddDate(0, 0, -1)
		csv, err = getCSV(date)
		if err != nil {
			return nil, err
		}
	}
	defer csv.Body.Close()

	cases := []*Cases{}
	err = gocsv.Unmarshal(csv.Body, &cases)
	if err != nil {
		return nil, err
	}

	return cases, nil
}
