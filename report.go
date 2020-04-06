package corona

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Report stores reported case information.
type Report struct {
	Date       time.Time `json:"Date"`
	Key        string    `json:"Key"`
	Country    *Country  `json:"-"`
	Region     *Region   `json:"-"`
	Confirmed  float64   `json:"Confirmed"`
	Deaths     float64   `json:"Deaths"`
	Population float64   `json:"Population"`
	Latitude   float64   `json:"Latitude"`
	Longitude  float64   `json:"Longitude"`
}

// Country stores country information.
type Country struct {
	Code string
	Name string
}

// Region stores province/state information.
type Region struct {
	Code string
	Name string
}

// UnmarshalJSON implements json.Unmarshaler to parse the data response.
func (c *Report) UnmarshalJSON(input []byte) error {
	if c.Country == nil {
		c.Country = &Country{}
	}

	if c.Region == nil {
		c.Region = &Region{}
	}

	var data map[string]interface{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	for k, v := range data {
		if v == "" || v == "null" {
			continue
		}

		switch k {
		case "Key":
			if key, ok := v.(string); ok {
				c.Key = key
			}
		case "CountryCode":
			if code, ok := v.(string); ok {
				c.Country.Code = code
			}
		case "CountryName":
			if name, ok := v.(string); ok {
				c.Country.Name = name
			}
		case "RegionCode":
			if code, ok := v.(string); ok {
				c.Region.Code = code
			}
		case "RegionName":
			if name, ok := v.(string); ok {
				c.Region.Name = name
			}
		case "Date":
			if date, ok := v.(string); ok {
				t, err := time.Parse("2006-01-02", date)
				if err != nil {
					return err
				}
				c.Date = t
			}
		case "Confirmed":
			if confirmed, ok := v.(float64); ok {
				c.Confirmed = confirmed
			}
		case "Deaths":
			if deaths, ok := v.(float64); ok {
				c.Deaths = deaths
			}
		case "Population":
			if pop, ok := v.(float64); ok {
				c.Population = pop
			}
		case "Latitude":
			if lat, ok := v.(string); ok {
				f, err := strconv.ParseFloat(lat, 64)
				if err != nil {
					return err
				}
				c.Latitude = f
			}
		case "Longitude":
			if lon, ok := v.(string); ok {
				f, err := strconv.ParseFloat(lon, 64)
				if err != nil {
					return err
				}
				c.Longitude = f
			}
		}
	}

	return nil
}

func getReports(url string) (Cases, error) {
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

	cases := Cases{}
	err = json.Unmarshal(content, &cases)
	if err != nil {
		return nil, err
	}

	return cases, nil
}
