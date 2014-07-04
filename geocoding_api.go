package mapquest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

var _ = log.Print

const (
	// GeocodingPathPrefix is the default path prefix for the Geocoding API.
	GeocodingPathPrefix = "/geocoding/v1"
)

// GeocodingAPI enables users to take an address and get the associated
// latitude and longitude from the MapQuest API.
// See http://open.mapquestapi.com/geocoding/ for details.
type GeocodingAPI struct {
	c *Client
}

// Address returns information about a specific address.
func (api *GeocodingAPI) Address(req *GeocodingAddressRequest) (*GeocodingAddressResponse, error) {
	u, err := api.buildAddressURL(req)
	if err != nil {
		return nil, err
	}

	res := new(GeocodingAddressResponse)
	if err := api.c.getJSON(u, res); err != nil {
		return nil, err
	}

	return res, nil
}

// buildAddressURL returns the complete URL for the request,
// including the key to query the MapQuest API.
func (api *GeocodingAPI) buildAddressURL(req *GeocodingAddressRequest) (string, error) {
	urls := fmt.Sprintf("%s%s/address", api.c.BaseURL(), GeocodingPathPrefix)
	u, err := url.Parse(urls)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	// Add key and other parameters to the query string
	q := u.Query()
	q.Set("inFormat", "json")
	q.Set("json", string(jsonData))
	q.Set("outFormat", "json")

	// Key has to be handled specifically here, because
	// the MapQuest API seems to not like the key URL-encoded
	u.RawQuery = fmt.Sprintf("key=%s&%s", api.c.key, q.Encode())
	return u.String(), nil
}

type GeocodingAddressRequest struct {
	Location *GeocodingLocation `json:"location,omitempty"`
}

type GeocodingAddressResponse struct {
	Options struct {
		IgnoreLatLngInput bool `json:"ignoreLatLngInput,omitempty"`
		MaxResults        int  `json:"maxResults,omitempty"`
		ThumbMaps         bool `json:"thumbMaps,omitempty"`
	} `json:"options,omitempty"`
	Results []*GeocodingAddressResponseResults `json:"results,omitempty"`
}

type GeocodingAddressResponseResults struct {
	Locations        []*GeocodingAddressResponseLocation `json:"locations,omitempty"`
	ProvidedLocation *GeocodingLocation                  `json:"providedLocation,omitempty"`
}

type GeocodingAddressResponseLocation struct {
	AdminArea1     string `json:"adminArea1,omitempty"`
	AdminArea1Type string `json:"adminArea1Type,omitempty"`
	AdminArea2     string `json:"adminArea2,omitempty"`
	AdminArea2Type string `json:"adminArea2Type,omitempty"`
	AdminArea3     string `json:"adminArea3,omitempty"`
	AdminArea3Type string `json:"adminArea3Type,omitempty"`
	AdminArea4     string `json:"adminArea4,omitempty"`
	AdminArea4Type string `json:"adminArea4Type,omitempty"`
	AdminArea5     string `json:"adminArea5,omitempty"`
	AdminArea5Type string `json:"adminArea5Type,omitempty"`
	DisplayLatLng  *struct {
		Latitude  float64 `json:"lat,omitempty"`
		Longitude float64 `json:"lng,omitempty"`
	} `json:"displayLatLng,omitempty"`
	DragPoint          *bool  `json:"dragPoint,omitempty"`
	GeocodeQuality     string `json:"geocodeQualtity,omitempty"`
	GeocodeQualityCode string `json:"geocodeQualtityCode,omitempty"`
	LatLng             *struct {
		Latitude  float64 `json:"lat,omitempty"`
		Longitude float64 `json:"lng,omitempty"`
	} `json:"latLng,omitempty"`
	LinkId       int    `json:"linkId,omitempty"`
	MapUrl       string `json:"mapUrl,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	SideOfStreet string `json:"sideOfStreet,omitempty"`
	Street       string `json:"street,omitempty"`
	Type         string `json:"type,omitempty"`
}

type GeocodingLocation struct {
	LatLng     string `json:"latLng,omitempty"`
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	County     string `json:"county,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	Type       string `json:"type,omitempty"`
	DragPoint  *bool  `json:"dragPoint,omitempty"`
}
