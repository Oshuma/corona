package corona

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Cases stores case information.
type Cases struct {
	Date       time.Time `json:"Date"`
	Country    *Country  `json:"-"`
	Region     *Region   `json:"-"`
	Confirmed  int       `json:"Confirmed"`
	Deaths     int       `json:"Deaths"`
	Population int       `json:"Population"`
	Latitude   float64   `json:"Latitude"`
	Longitude  float64   `json:"Longitude"`
}

type Country struct {
	Code string
	Name string
}

type Region struct {
	Code string
	Name string
}

func (c *Cases) UnmarshalJSON(input []byte) error {
	if c.Country == nil {
		c.Country = &Country{}
	}

	if c.Region == nil {
		c.Region = &Region{}
	}

	var data map[string]string
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	for k, v := range data {
		if v == "" || v == "null" {
			continue
		}

		switch k {
		case "CountryCode":
			c.Country.Code = v
		case "CountryName":
			c.Country.Name = v
		case "RegionCode":
			c.Region.Code = v
		case "RegionName":
			c.Region.Name = v
		case "Date":
			t, err := time.Parse("2006-01-02", v)
			if err != nil {
				return err
			}
			c.Date = t
		case "Confirmed":
			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			c.Confirmed = i
		case "Deaths":
			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			c.Deaths = i
		case "Population":
			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			c.Population = i
		case "Latitude":
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			c.Latitude = f
		case "Longitude":
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			c.Longitude = f
		}
	}

	return nil
}

func getCases(url string) ([]*Cases, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cases := []*Cases{}
	err = json.Unmarshal(content, &cases)
	if err != nil {
		return nil, err
	}

	return cases, nil
}

func filterCountryName(cases []*Cases, country string) ([]*Cases, error) {
	byCountry := []*Cases{}

	for _, c := range cases {
		if strings.EqualFold(c.Country.Name, country) {
			byCountry = append(byCountry, c)
		}
	}

	if len(byCountry) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byCountry, nil
}

func filterCountryCode(cases []*Cases, code string) ([]*Cases, error) {
	byCountry := []*Cases{}

	for _, c := range cases {
		if strings.EqualFold(c.Country.Code, code) {
			byCountry = append(byCountry, c)
		}
	}

	if len(byCountry) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byCountry, nil
}

func filterRegionName(cases []*Cases, region string) ([]*Cases, error) {
	byRegion := []*Cases{}

	for _, c := range cases {
		if strings.EqualFold(c.Region.Name, region) {
			byRegion = append(byRegion, c)
		}
	}

	if len(byRegion) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byRegion, nil
}

func filterRegionCode(cases []*Cases, code string) ([]*Cases, error) {
	byRegion := []*Cases{}

	for _, c := range cases {
		if strings.EqualFold(c.Region.Code, code) {
			byRegion = append(byRegion, c)
		}
	}

	if len(byRegion) == 0 {
		return nil, ErrorNoCasesFound
	}

	return byRegion, nil
}
