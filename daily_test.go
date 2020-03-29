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

func TestDailyByCountryName(t *testing.T) {
	cases, err := DailyByCountryName("United States of America")
	if err != nil {
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.Country.Name != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.Country.Code)
	}
}

func TestDailyByCountryNameInsensitive(t *testing.T) {
	cases, err := DailyByCountryName("UnItEd StAtEs Of aMeRiCa")
	if err != nil {
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.Country.Name != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.Country.Code)
	}
}

func TestDailyByCountryNameNotFound(t *testing.T) {
	_, err := DailyByCountryName("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTotalByCountryName(t *testing.T) {
	r, err := TotalByCountryName("United States of America")
	if err != nil {
		t.Fatal(err)
	}

	if r.Country.Name != "United States of America" {
		t.Fatalf("wrong country loaded: %s", r.Country.Name)
	}
}

func TestDailyByCountryCode(t *testing.T) {
	cases, err := DailyByCountryCode("US")
	if err != nil {
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.Country.Code != "US" {
		t.Fatalf("wrong country loaded: %s", c.Country.Code)
	}
}

func TestDailyByCountryCodeInsensitive(t *testing.T) {
	cases, err := DailyByCountryCode("Us")
	if err != nil {
		t.Fatal(err)
	}

	if len(cases) == 0 {
		t.Fatal("no cases loaded")
	}

	c := cases[0]
	if c.Country.Code != "US" {
		t.Fatalf("wrong country loaded: %s", c.Country.Code)
	}
}

func TestDailyByCountryCodeNotFound(t *testing.T) {
	_, err := DailyByCountryCode("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTotalByCountryCode(t *testing.T) {
	r, err := TotalByCountryCode("US")
	if err != nil {
		t.Fatal(err)
	}

	if r.Country.Code != "US" {
		t.Fatalf("wrong country loaded: %s", r.Country.Code)
	}
}

func TestDailyByRegionName(t *testing.T) {
	cases, err := DailyByRegionName("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range cases {
		if r.Region.Name != "South Carolina" {
			t.Fatalf("wrong region loaded: %s", r.Region.Name)
		}
	}
}

func TestDailyByRegionNameInsensitive(t *testing.T) {
	cases, err := DailyByRegionName("sOuTh CaRoLiNa")
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range cases {
		if r.Region.Name != "South Carolina" {
			t.Fatalf("wrong region loaded: %s", r.Region.Name)
		}
	}
}

func TestDailyByRegionNameNotFound(t *testing.T) {
	_, err := DailyByRegionName("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestDailyByRegionCode(t *testing.T) {
	cases, err := DailyByRegionCode("NY")
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range cases {
		if r.Region.Code != "NY" {
			t.Fatalf("wrong region loaded: %s", r.Region.Code)
		}
	}
}

func TestDailyByRegionCodeInsensitive(t *testing.T) {
	cases, err := DailyByRegionCode("Ny")
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range cases {
		if r.Region.Code != "NY" {
			t.Fatalf("wrong region loaded: %s", r.Region.Code)
		}
	}
}

func TestDailyByRegionCodeNotFound(t *testing.T) {
	_, err := DailyByRegionCode("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}
