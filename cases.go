package corona

import (
	"strconv"
	"time"
)

// Cases stores case information.
type Cases struct {
	Date        Date      `json:"Date"`
	CountryCode string    `json:"CountryCode"`
	CountryName string    `json:"CountryName"`
	RegionCode  string    `json:"RegionCode"`
	RegionName  string    `json:"RegionName"`
	Confirmed   Confirmed `json:"Confirmed"`
	Deaths      Deaths    `json:"Deaths"`
	Latitude    Latitude  `json:"Latitude"`
	Longitude   Longitude `json:"Longitude"`
	// Population  int    `json:"Population"`  // TODO: Not yet implemented in response.
}

// Date is a time.Time wrapper used to unmarshal and parse the JSON response.
type Date struct {
	time.Time
}

// UnmarshalJSON unmarshals the time format in the JSON response.
func (d Date) UnmarshalJSON(input []byte) error {
	s, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	d.Time = t
	return err
}

type stringToInt int

func (si stringToInt) UnmarshalJSON(input []byte) error {
	if string(input) == "null" {
		return nil
	}

	s, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	si = stringToInt(i)
	return nil
}

// Confirmed is an int wrapper used to unmarshal from a JSON string.
type Confirmed struct {
	stringToInt
}

// Deaths is an int wrapper used to unmarshal from a JSON string.
type Deaths struct {
	stringToInt
}

type stringToFloat64 float64

func (sf stringToFloat64) UnmarshalJSON(input []byte) error {
	if string(input) == "null" {
		return nil
	}

	s, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	sf = stringToFloat64(f)
	return nil
}

// Latitude is a float64 wrapper used to unmarshal from a JSON string.
type Latitude struct {
	stringToFloat64
}

// Longitude is a float64 wrapper used to unmarshal from a JSON string.
type Longitude struct {
	stringToFloat64
}
