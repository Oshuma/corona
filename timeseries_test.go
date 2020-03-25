package corona

import (
	"testing"
)

func TestTimeSeries(t *testing.T) {
	ts, err := TimeSeries()
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}
}

func TestTimeSeriesByCountry(t *testing.T) {
	ts, err := TimeSeriesByCountry("United States of America")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	c := ts[0]
	if c.CountryName != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.CountryName)
	}
}

func TestTimeSeriesByCountryInsensitive(t *testing.T) {
	ts, err := TimeSeriesByCountry("UnItEd StAtEs Of aMeRiCa")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	c := ts[0]
	if c.CountryName != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.CountryName)
	}
}

func TestTimeSeriesByCountryCode(t *testing.T) {
	ts, err := TimeSeriesByCountry("US")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	c := ts[0]
	if c.CountryName != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.CountryName)
	}
}

func TestTimeSeriesByCountryCodeInsensitive(t *testing.T) {
	ts, err := TimeSeriesByCountry("Us")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	c := ts[0]
	if c.CountryName != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.CountryName)
	}
}

func TestTimeSeriesByRegion(t *testing.T) {
	ts, err := TimeSeriesByRegion("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	c := ts[0]
	if c.RegionName != "South Carolina" {
		t.Fatalf("wrong region loaded: %s", c.RegionName)
	}
}

func TestTimeSeriesByRegionInsensitive(t *testing.T) {
	ts, err := TimeSeriesByRegion("SoUtH cArOlInA")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	c := ts[0]
	if c.RegionName != "South Carolina" {
		t.Fatalf("wrong region loaded: %s", c.RegionName)
	}
}
