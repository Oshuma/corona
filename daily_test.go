package corona

import (
	"testing"
)

// TODO: Load JSON from local file.

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
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.CountryCode != "US" {
		t.Fatalf("wrong country loaded: %s", c.CountryName)
	}
}

func TestDailyByCountryInsensitive(t *testing.T) {
	cases, err := DailyByCountry("Us")
	if err != nil {
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.CountryCode != "US" {
		t.Fatalf("wrong country loaded: %s", c.CountryName)
	}
}

func TestDailyByCountryNotFound(t *testing.T) {
	_, err := DailyByCountry("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestDailyByRegion(t *testing.T) {
	cases, err := DailyByRegion("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	if cases.RegionName != "South Carolina" {
		t.Fatalf("wrong region loaded: %s", cases.RegionName)
	}
}

func TestDailyByRegionInsensitive(t *testing.T) {
	cases, err := DailyByRegion("sOuTh CaRoLiNa")
	if err != nil {
		t.Fatal(err)
	}

	if cases.RegionName != "South Carolina" {
		t.Fatalf("wrong region loaded: %s", cases.RegionName)
	}
}

func TestDailyByRegionNotFound(t *testing.T) {
	_, err := DailyByRegion("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}
