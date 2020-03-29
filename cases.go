package corona

import (
	"strings"
)

// Cases is a collection of Reports.
type Cases []*Report

// FilterCountryName returns only reports in the given country name; case insensitive.
func (c Cases) FilterCountryName(country string) (Cases, error) {
	byCountry := Cases{}

	for _, r := range c {
		if strings.EqualFold(r.Country.Name, country) {
			byCountry = append(byCountry, r)
		}
	}

	if len(byCountry) == 0 {
		return nil, ErrorNoReportsFound
	}

	return byCountry, nil
}

// FilterCountryCode returns only reports in the given country code; case insensitive.
func (c Cases) FilterCountryCode(code string) (Cases, error) {
	byCountry := Cases{}

	for _, r := range c {
		if strings.EqualFold(r.Country.Code, code) {
			byCountry = append(byCountry, r)
		}
	}

	if len(byCountry) == 0 {
		return nil, ErrorNoReportsFound
	}

	return byCountry, nil
}

// FilterRegionName returns only reports in the given region name; case insensitive.
func (c Cases) FilterRegionName(region string) (Cases, error) {
	byRegion := Cases{}

	for _, r := range c {
		if strings.EqualFold(r.Region.Name, region) {
			byRegion = append(byRegion, r)
		}
	}

	if len(byRegion) == 0 {
		return nil, ErrorNoReportsFound
	}

	return byRegion, nil
}

// FilterRegionCode returns only reports in the given region code; case insensitive.
func (c Cases) FilterRegionCode(code string) (Cases, error) {
	byRegion := Cases{}

	for _, r := range c {
		if strings.EqualFold(r.Region.Code, code) {
			byRegion = append(byRegion, r)
		}
	}

	if len(byRegion) == 0 {
		return nil, ErrorNoReportsFound
	}

	return byRegion, nil
}
