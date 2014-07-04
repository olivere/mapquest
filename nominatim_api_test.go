package mapquest

import (
	"bytes"
	"log"
	"testing"
)

func TestNominatimBuildSearchURLs(t *testing.T) {
	testKey, err := readKey(t)
	if err != nil {
		t.Fail()
		return
	}

	tests := []struct {
		Request *NominatimSearchRequest
		URL     string
	}{
		{
			Request: &NominatimSearchRequest{
				Query: "Marienplatz 2a, München, DE",
				Limit: 1,
			},
			URL: "http://open.mapquestapi.com/nominatim/v1/search.php?addressdetails=1&format=json&limit=1&q=Marienplatz+2a%2C+M%C3%BCnchen%2C+DE",
		},
	}

	client := NewClient(testKey)
	for _, test := range tests {
		got, err := client.Nominatim().buildSearchURL(test.Request)
		if err != nil {
			t.Fatalf("expeced no error, got: %v", err)
		}
		if got != test.URL {
			t.Errorf("expected %q, got: %q", test.URL, got)
		}
	}
}

func TestNominatimSearch(t *testing.T) {
	key, err := readKey(t)
	if err != nil {
		t.Fail()
		return
	}

	client := NewClient(key)

	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	client.SetLogger(logger)

	req := &NominatimSearchRequest{
		Query: "Unter den Linden 117, Berlin, DE",
		Limit: 1,
	}
	res, err := client.Nominatim().Search(req)
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
		t.Fatalf("expected 1 result, got: %v\nresponse: %v", len(res.Results), logging)
	}

	result := res.Results[0]
	if result.Address == nil {
		t.Fatalf("expected address, got: %v", result.Address)
	}
	if result.Address.City != "Berlin" {
		t.Errorf("expected city %q, got: %q", "Berlin", result.Address.City)
	}
	if result.Address.CityDistrict != "Mitte" {
		t.Errorf("expected city district %q, got: %q", "Mitte", result.Address.CityDistrict)
	}
	if result.Address.Continent != "European Union" {
		t.Errorf("expected continent %q, got: %q", "European Union", result.Address.Continent)
	}
	if result.Address.Country != "Deutschland" {
		t.Errorf("expected country %q, got: %q", "Deutschland", result.Address.Country)
	}
	if result.Address.CountryCode != "de" {
		t.Errorf("expected country code %q, got: %q", "de", result.Address.CountryCode)
	}
	if result.Address.Neighbourhood != "Scheunenviertel" {
		t.Errorf("expected neighbourhood %q, got: %q", "Scheunenviertel", result.Address.Neighbourhood)
	}
	if result.Address.PostCode != "10117" {
		t.Errorf("expected postcode %q, got: %q", "10117", result.Address.PostCode)
	}
	if result.Address.Road != "Unter den Linden" {
		t.Errorf("expected road %q, got: %q", "Unter den Linden", result.Address.Road)
	}
	if result.Address.State != "Berlin" {
		t.Errorf("expected state %q, got: %q", "Berlin", result.Address.State)
	}
	if result.Address.Suburb != "Mitte" {
		t.Errorf("expected suburb %q, got: %q", "Mitte", result.Address.Suburb)
	}

	if result.Class != "highway" {
		t.Errorf("expected class %q, got: %q", "highway", result.Class)
	}
	if result.DisplayName != "Unter den Linden, Scheunenviertel, Mitte, Berlin, 10117, Deutschland, European Union" {
		t.Errorf("expected display name %q, got: %q", "Unter den Linden, Scheunenviertel, Mitte, Berlin, 10117, Deutschland, European Union", result.DisplayName)
	}
	if result.Importance != float64(1.2) {
		t.Errorf("expected importance %f, got: %f", float64(1.2), result.Importance)
	}
	if result.Latitude != float64(52.5173324) {
		t.Errorf("expected latitude %f, got: %f", float64(52.5173324), result.Latitude)
	}
	if result.Longitude != float64(13.3932632) {
		t.Errorf("expected longitude %f, got: %f", float64(13.3932632), result.Longitude)
	}
	if result.OSMId != "110676319" {
		t.Errorf("expected OSM id %q, got: %q", "110676319", result.OSMId)
	}
	if result.OSMType != "way" {
		t.Errorf("expected OSM type %q, got: %q", "way", result.OSMType)
	}
	if result.PlaceId != "70421736" {
		t.Errorf("expected place id %q, got: %q", "70421736", result.PlaceId)
	}
	if result.Type != "primary" {
		t.Errorf("expected type %q, got: %q", "primary", result.Type)
	}
	if result.License != "Data © OpenStreetMap contributors, ODbL 1.0. http://www.openstreetmap.org/copyright" {
		t.Errorf("expected license %q, got: %q", "Data © OpenStreetMap contributors, ODbL 1.0. http://www.openstreetmap.org/copyright", result.License)
	}
}
