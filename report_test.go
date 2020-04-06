package corona

import (
	"testing"

	"encoding/json"
	"fmt"
	"time"
)

var reportJSON = []byte(`
	{
		"Date": "2020-03-24",
		"Key": "US_SC",
		"CountryCode": "US",
		"CountryName": "United States of America",
		"RegionCode": "SC",
		"RegionName": "South Carolina",
		"Confirmed": 298,
		"Deaths": 5,
		"Latitude": "33.8569",
		"Longitude": "-80.945",
		"Population": 12345
	}
`)

var reportJSONNull = []byte(`
	{
		"Date": null,
		"Key": null,
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
	var r Report
	err := json.Unmarshal(reportJSON, &r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Country.Code != "US" {
		t.Fatalf("error parsing Country.Code: %+v", r.Country)
	}

	if r.Country.Name != "United States of America" {
		t.Fatalf("error parsing Country.Name: %+v", r.Country.Name)
	}

	if r.Region.Code != "SC" {
		t.Fatalf("error parsing Region.Code: %+v", r.Region.Code)
	}

	if r.Region.Name != "South Carolina" {
		t.Fatalf("error parsing Region.Name: %+v", r.Region.Name)
	}

	if r.Confirmed != 298 {
		t.Fatalf("%+v", r)
		t.Fatalf("error parsing Confirmed: %+v", r.Confirmed)
	}

	if r.Deaths != 5 {
		t.Fatalf("error parsing Deaths: %+v", r.Deaths)
	}

	if r.Population != 12345 {
		t.Fatalf("error parsing Population: %+v", r.Population)
	}

	if r.Latitude != 33.8569 {
		t.Fatalf("error parsing Latitude: %+v", r.Latitude)
	}

	if r.Longitude != -80.945 {
		t.Fatalf("error parsing Longitude: %+v", r.Longitude)
	}
}

func TestParseNull(t *testing.T) {
	var r Report
	err := json.Unmarshal(reportJSONNull, &r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseDate(t *testing.T) {
	var r Report
	err := json.Unmarshal(reportJSON, &r)
	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2020, time.March, 24, 0, 0, 0, 0, time.UTC)
	if r.Date != expected {
		t.Fatalf("wrong date; got=%+v  expected=%s", r.Date, expected)
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
