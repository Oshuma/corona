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

	c := ts[0]
	if c.Country.Name != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.Country.Name)
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

	c := ts[0]
	if c.Country.Name != "United States of America" {
		t.Fatalf("wrong country loaded: %s", c.Country.Name)
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

	c := ts[0]
	if c.Country.Code != "US" {
		t.Fatalf("wrong country loaded: %s", c.Country.Code)
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

	c := ts[0]
	if c.Country.Code != "US" {
		t.Fatalf("wrong country loaded: %s", c.Country.Code)
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

	c := ts[0]
	if c.Region.Name != "South Carolina" {
		t.Fatalf("wrong region loaded: %s", c.Region.Name)
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

	c := ts[0]
	if c.Region.Name != "South Carolina" {
		t.Fatalf("wrong region loaded: %s", c.Region.Name)
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

	c := ts[0]
	if c.Region.Code != "NY" {
		t.Fatalf("wrong region loaded: %s", c.Region.Code)
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

	c := ts[0]
	if c.Region.Code != "NY" {
		t.Fatalf("wrong region loaded: %s", c.Region.Code)
	}
}

func TestTimeSeriesByRegionCodeNotFound(t *testing.T) {
	_, err := TimeSeriesByRegionCode("foo")
	if err != ErrorNoReportsFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}
