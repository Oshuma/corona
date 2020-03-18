package corona

import (
	"testing"
)

func TestTimeSeriesConfirmed(t *testing.T) {
	ts, err := TimeSeriesConfirmed()
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}
}

func TestTimeSeriesConfirmedByCountry(t *testing.T) {
	ts, err := TimeSeriesConfirmedByCountry("US")
	if err != nil {
		t.Fatal(err)
	}

	if ts.CountryRegion != "US" {
		t.Fatal("wrong country loaded")
	}
}

func TestTimeSeriesConfirmedByCountryNotFound(t *testing.T) {
	_, err := TimeSeriesConfirmedByCountry("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesConfirmedByState(t *testing.T) {
	ts, err := TimeSeriesConfirmedByState("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	if ts.ProvinceState != "South Carolina" {
		t.Fatal("wrong state loaded")
	}
}

func TestTimeSeriesConfirmedByStateNotFound(t *testing.T) {
	_, err := TimeSeriesConfirmedByState("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesConfirmedByProvince(t *testing.T) {
	ts, err := TimeSeriesConfirmedByProvince("Hubei")
	if err != nil {
		t.Fatal(err)
	}

	if ts.ProvinceState != "Hubei" {
		t.Fatal("wrong state loaded")
	}
}

func TestTimeSeriesConfirmedByProvinceNotFound(t *testing.T) {
	_, err := TimeSeriesConfirmedByProvince("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesDeaths(t *testing.T) {
	ts, err := TimeSeriesDeaths()
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}
}

func TestTimeSeriesDeathsByCountry(t *testing.T) {
	ts, err := TimeSeriesDeathsByCountry("US")
	if err != nil {
		t.Fatal(err)
	}

	if ts.CountryRegion != "US" {
		t.Fatal("wrong country loaded")
	}
}

func TestTimeSeriesDeathsByCountryNotFound(t *testing.T) {
	_, err := TimeSeriesDeathsByCountry("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesDeathsByState(t *testing.T) {
	ts, err := TimeSeriesDeathsByState("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	if ts.ProvinceState != "South Carolina" {
		t.Fatal("wrong state loaded")
	}
}

func TestTimeSeriesDeathsByStateNotFound(t *testing.T) {
	_, err := TimeSeriesDeathsByState("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesDeathsByProvince(t *testing.T) {
	ts, err := TimeSeriesDeathsByProvince("Hubei")
	if err != nil {
		t.Fatal(err)
	}

	if ts.ProvinceState != "Hubei" {
		t.Fatal("wrong province loaded")
	}
}

func TestTimeSeriesDeathsByProvinceNotFound(t *testing.T) {
	_, err := TimeSeriesDeathsByProvince("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesRecovered(t *testing.T) {
	ts, err := TimeSeriesRecovered()
	if err != nil {
		t.Fatal(err)
	}

	if len(ts) == 0 {
		t.Fatal("no data loaded")
	}
}

func TestTimeSeriesRecoveredByCountry(t *testing.T) {
	ts, err := TimeSeriesRecoveredByCountry("US")
	if err != nil {
		t.Fatal(err)
	}

	if ts.CountryRegion != "US" {
		t.Fatal("wrong country loaded")
	}
}

func TestTimeSeriesRecoveredByCountryNotFound(t *testing.T) {
	_, err := TimeSeriesRecoveredByCountry("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesRecoveredByState(t *testing.T) {
	ts, err := TimeSeriesRecoveredByState("South Carolina")
	if err != nil {
		t.Fatal(err)
	}

	if ts.ProvinceState != "South Carolina" {
		t.Fatal("wrong state loaded")
	}
}

func TestTimeSeriesRecoveredByStateNotFound(t *testing.T) {
	_, err := TimeSeriesRecoveredByState("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}

func TestTimeSeriesRecoveredByProvince(t *testing.T) {
	ts, err := TimeSeriesRecoveredByProvince("Hubei")
	if err != nil {
		t.Fatal(err)
	}

	if ts.ProvinceState != "Hubei" {
		t.Fatal("wrong province loaded")
	}
}

func TestTimeSeriesRecoveredByProvinceNotFound(t *testing.T) {
	_, err := TimeSeriesRecoveredByProvince("foo")
	if err != ErrorNoCasesFound {
		t.Fatalf("wrong error returned: %s", err)
	}
}
