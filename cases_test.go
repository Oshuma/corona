package corona

import (
	"testing"

	"encoding/json"
	"fmt"
	"time"
)

var casesJSON = []byte(`
  {
    "Date": "2020-03-24",
    "CountryCode": "US",
    "CountryName": "United States of America",
    "RegionCode": "SC",
    "RegionName": "South Carolina",
    "Confirmed": "298",
    "Deaths": "5",
    "Latitude": "33.8569",
    "Longitude": "-80.945",
    "Population": "12345"
  }
`)

var casesJSONNull = []byte(`
  {
    "Date": null,
    "CountryCode": null,
    "CountryName": null,
    "RegionCode": null,
    "RegionName": null,
    "Confirmed": null,
    "Deaths": null,
    "Latitude": null,
    "Longitude": null,
    "Population": null
  }
`)

func TestParse(t *testing.T) {
	var c Cases
	err := json.Unmarshal(casesJSON, &c)
	if err != nil {
		t.Fatal(err)
	}

	if c.Country.Code != "US" {
		t.Fatalf("error parsing Country.Code: %+v", c.Country)
	}

	if c.Country.Name != "United States of America" {
		t.Fatalf("error parsing Country.Name: %+v", c.Country.Name)
	}

	if c.Region.Code != "SC" {
		t.Fatalf("error parsing Region.Code: %+v", c.Region.Code)
	}

	if c.Region.Name != "South Carolina" {
		t.Fatalf("error parsing Region.Name: %+v", c.Region.Name)
	}

	if c.Confirmed != 298 {
		t.Fatalf("error parsing Confirmed: %+v", c.Confirmed)
	}

	if c.Deaths != 5 {
		t.Fatalf("error parsing Deaths: %+v", c.Deaths)
	}

	if c.Population != 12345 {
		t.Fatalf("error parsing Population: %+v", c.Population)
	}

	if c.Latitude != 33.8569 {
		t.Fatalf("error parsing Latitude: %+v", c.Latitude)
	}

	if c.Longitude != -80.945 {
		t.Fatalf("error parsing Longitude: %+v", c.Longitude)
	}
}

func TestParseNull(t *testing.T) {
	var c Cases
	err := json.Unmarshal(casesJSONNull, &c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseDate(t *testing.T) {
	var c Cases
	err := json.Unmarshal(casesJSON, &c)
	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2020, time.March, 24, 0, 0, 0, 0, time.UTC)
	if c.Date != expected {
		t.Fatalf("wrong date; got=%+v  expected=%s", c.Date, expected)
	}
}

func TestInvalidJSON(t *testing.T) {
	j := []byte(`{ invalid }`)

	var c Cases
	err := json.Unmarshal(j, &c)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestInvalidJSONFields(t *testing.T) {
	fields := []string{
		"CountryCode",
		"CountryName",
		"RegionCode",
		"RegionName",
		"Confirmed",
		"Deaths",
		"Latitude",
		"Longitude",
		"Population",
	}

	for _, f := range fields {
		j := []byte(fmt.Sprintf("{ \"%s\": invalid }", f))

		var c Cases
		err := json.Unmarshal(j, &c)
		if err == nil {
			t.Fatal("expected error")
		}
	}
}

func TestInvalidDate(t *testing.T) {
	j := []byte(`{ "Date": "foo" }`)

	var c Cases
	err := json.Unmarshal(j, &c)
	if err == nil {
		t.Fatal("expected error")
	}
}
