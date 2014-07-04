package mapquest

import (
	"bytes"
	"log"
	"testing"
)

func TestGeocodingBuildAddressURLs(t *testing.T) {
	testKey, err := readKey(t)
	if err != nil {
		t.Fail()
		return
	}

	tests := []struct {
		Request *GeocodingAddressRequest
		URL     string
	}{
		{
			Request: &GeocodingAddressRequest{
				Location: &GeocodingLocation{
					Street:     "1090 N Charlotte St",
					City:       "Lancaster",
					State:      "PA",
					PostalCode: "17603",
				},
			},
			URL: "http://open.mapquestapi.com/geocoding/v1/address?key=" + testKey + "&inFormat=json&json=%7B%22location%22%3A%7B%22street%22%3A%221090+N+Charlotte+St%22%2C%22city%22%3A%22Lancaster%22%2C%22state%22%3A%22PA%22%2C%22postalCode%22%3A%2217603%22%7D%7D&outFormat=json",
		},
	}

	client := NewClient(testKey)
	for _, test := range tests {
		got, err := client.Geocoding().buildAddressURL(test.Request)
		if err != nil {
			t.Fatalf("expeced no error, got: %v", err)
		}
		if got != test.URL {
			t.Errorf("expected %q, got: %q", test.URL, got)
		}
	}
}

func TestGeocodingAddress(t *testing.T) {
	key, err := readKey(t)
	if err != nil {
		t.Fail()
		return
	}

	client := NewClient(key)

	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	client.SetLogger(logger)

	req := &GeocodingAddressRequest{
		Location: &GeocodingLocation{
			Street:     "1090 N Charlotte St",
			City:       "Lancaster",
			State:      "PA",
			PostalCode: "17603",
		},
	}
	res, err := client.Geocoding().Address(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if res == nil {
		t.Fatalf("expected result, got: %v", res)
	}

	//t.Logf("%v", buf.String())

	logging := buf.String()
	if len(logging) == 0 {
		t.Errorf("expected client to do loggin, got: %v", logging)
	}

	if len(res.Results) != 1 {
		t.Fatalf("expected 1 result, got: %v", len(res.Results))
	}

	result := res.Results[0]
	if result.ProvidedLocation == nil {
		t.Fatalf("expected provided location, got: %v", result.ProvidedLocation)
	}
	if result.ProvidedLocation.City != req.Location.City {
		t.Errorf("expected city %q, got: %q", req.Location.City, result.ProvidedLocation.City)
	}

	if len(result.Locations) != 1 {
		t.Fatalf("expected 1 location in result, got: %v", len(result.Locations))
	}

	location := result.Locations[0]
	if location.AdminArea1 != "US" {
		t.Errorf("expected AdminArea1 %q, got: %q", "US", location.AdminArea1)
	}
	if location.AdminArea1Type != "Country" {
		t.Errorf("expected AdminArea1Type %q, got: %q", "Country", location.AdminArea1Type)
	}
	if location.AdminArea3 != "PA" {
		t.Errorf("expected AdminArea3 %q, got: %q", "PA", location.AdminArea3)
	}
	if location.AdminArea3Type != "State" {
		t.Errorf("expected AdminArea3Type %q, got: %q", "State", location.AdminArea3Type)
	}
	if location.AdminArea4 != "Lancaster County" {
		t.Errorf("expected AdminArea4 %q, got: %q", "Lancaster County", location.AdminArea4)
	}
	if location.AdminArea4Type != "County" {
		t.Errorf("expected AdminArea4Type %q, got: %q", "County", location.AdminArea4Type)
	}
	if location.AdminArea5 != "Lancaster" {
		t.Errorf("expected AdminArea5 %q, got: %q", "Lancaster", location.AdminArea5)
	}
	if location.AdminArea5Type != "City" {
		t.Errorf("expected AdminArea5Type %q, got: %q", "City", location.AdminArea5Type)
	}
	if location.LatLng == nil {
		t.Fatalf("expected LatLng, got: %q", location.LatLng)
	}
	if location.LatLng.Latitude != float64(40.05305) {
		t.Errorf("expected LatLng.Latitude %q, got: %q", float64(40.05305), location.LatLng.Latitude)
	}
	if location.LatLng.Longitude != float64(-76.314707) {
		t.Errorf("expected LatLng.Longitude %q, got: %q", float64(-76.314707), location.LatLng.Longitude)
	}
	if location.MapUrl == "" {
		t.Errorf("expected MapUrl, got: %q", location.MapUrl)
	}
	if location.PostalCode != "17603" {
		t.Errorf("expected PostalCode %q, got: %q", "17603", location.PostalCode)
	}
	if location.SideOfStreet != "N" {
		t.Errorf("expected SideOfStreet %q, got: %q", "N", location.SideOfStreet)
	}
	if location.Street != "1090 North Charlotte Street" {
		t.Errorf("expected PostalCode %q, got: %q", "1090 North Charlotte Street", location.Street)
	}
	if location.Type != "s" {
		t.Errorf("expected Type %q, got: %q", "s", location.Type)
	}
}
