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

func TestTimeSeriesByCountryName(t *testing.T) {
	ts, err := TimeSeriesByCountryName("United States of America")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Country.Name != "United States of America" {
				t.Fatalf("wrong country loaded: %s", r.Country.Name)
			}
		}
	}
}

func TestTimeSeriesByCountryNameInsensitive(t *testing.T) {
	ts, err := TimeSeriesByCountryName("UnItEd StAtEs Of aMeRiCa")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Country.Name != "United States of America" {
				t.Fatalf("wrong country loaded: %s", r.Country.Name)
			}
		}
	}
}

func TestTimeSeriesByCountryNameNotFound(t *testing.T) {
	_, err := TimeSeriesByCountryName("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesByCountryCode(t *testing.T) {
	ts, err := TimeSeriesByCountryCode("US")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Country.Code != "US" {
				t.Fatalf("wrong country loaded: %s", r.Country.Code)
			}
		}
	}
}

func TestTimeSeriesByCountryCodeInsensitive(t *testing.T) {
	ts, err := TimeSeriesByCountryCode("Us")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Country.Code != "US" {
				t.Fatalf("wrong country loaded: %s", r.Country.Code)
			}
		}
	}
}

func TestTimeSeriesByCountryCodeNotFound(t *testing.T) {
	_, err := TimeSeriesByCountryCode("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesByRegionName(t *testing.T) {
	ts, err := TimeSeriesByRegionName("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Region.Name != "South Carolina" {
				t.Fatalf("wrong region loaded: %s", r.Region.Name)
			}
		}
	}
}

func TestTimeSeriesByRegionNameInsensitive(t *testing.T) {
	ts, err := TimeSeriesByRegionName("SoUtH cArOlInA")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Region.Name != "South Carolina" {
				t.Fatalf("wrong region loaded: %s", r.Region.Name)
			}
		}
	}
}

func TestTimeSeriesByRegionNameNotFound(t *testing.T) {
	_, err := TimeSeriesByRegionName("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesByRegionCode(t *testing.T) {
	ts, err := TimeSeriesByRegionCode("NY")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Region.Code != "NY" {
				t.Fatalf("wrong region loaded: %s", r.Region.Code)
			}
		}
	}
}

func TestTimeSeriesByRegionCodeInsensitive(t *testing.T) {
	ts, err := TimeSeriesByRegionCode("Ny")
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}

	for _, cases := range ts {
		for _, r := range cases {
			if r.Region.Code != "NY" {
				t.Fatalf("wrong region loaded: %s", r.Region.Code)
			}
		}
	}
}

func TestTimeSeriesByRegionCodeNotFound(t *testing.T) {
	_, err := TimeSeriesByRegionCode("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}
