package corona

import (
	"testing"

	"reflect"
)

// TODO: Load CSV from local file.

func TestDailyWorldwide(t *testing.T) {
	cases, err := DailyWorldwide()
	if err != nil {
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}
}

func TestDailyByCountry(t *testing.T) {
	cases, err := DailyByCountry("US")
	if err != nil {
		t.Fatalf("DailyByCountry(): %s", err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.CountryRegion != "US" {
		t.Fatalf("wrong country loaded: %s", c.CountryRegion)
	}
}

func TestDailyByCountryNotFound(t *testing.T) {
	_, err := DailyByCountry("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestDailyByState(t *testing.T) {
	cases, err := DailyByState("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	c := cases[0]
	if c.ProvinceState != "South Carolina" {
		t.Fatalf("wrong state loaded: %s", c.ProvinceState)
	}
}

func TestDailyByStateNotFound(t *testing.T) {
	_, err := DailyByState("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestDailyByProvince(t *testing.T) {
	cases, err := DailyByProvinceState("Hubei")
	if err != nil {
		t.Fatal(err)
	}

	c := cases[0]
	if c.ProvinceState != "Hubei" {
		t.Fatalf("wrong province loaded: %s", c.ProvinceState)
	}
}

func TestDailyByProvinceNotFound(t *testing.T) {
	_, err := DailyByProvinceState("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestParseLastUpdate(t *testing.T) {
	cases, err := DailyWorldwide()
	if err != nil {
		t.Fatal(err)
	}

	c := cases[0]
	instance := reflect.TypeOf(c.LastUpdate)
	if instance.Name() != "LastUpdate" {
		t.Fatal("could not parse LastUpdate")
	}

	if c.LastUpdate.IsZero() {
		t.Fatal("error parsing LastUpdate")
	}
}
