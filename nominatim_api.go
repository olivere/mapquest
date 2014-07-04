package mapquest

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

var _ = log.Print

const (
	// NominatimPathPrefix is the default path prefix for the Nominatim API.
	NominatimPathPrefix = "/nominatim/v1"
)

// NominatimAPI is a geographic search service that relies solely on the
// data contributed to OpenStreetMap.
// See http://open.mapquestapi.com/nominatim/ for details.
type NominatimAPI struct {
	c *Client
}

// Search searches for details given an address.
func (api *NominatimAPI) Search(req *NominatimSearchRequest) (*NominatimSearchResponse, error) {
	u, err := api.buildSearchURL(req)
	if err != nil {
		return nil, err
	}

	res := new(NominatimSearchResponse)
	res.Results = make([]*NominatimSearchResult, 0)

	if err := api.c.getJSON(u, &res.Results); err != nil {
		return nil, err
	}

	return res, nil
}

// buildSearchURL returns the complete URL for the request,
// including the key to query the MapQuest API.
func (api *NominatimAPI) buildSearchURL(req *NominatimSearchRequest) (string, error) {
	urls := fmt.Sprintf("%s%s/search.php", api.c.BaseURL(), NominatimPathPrefix)
	u, err := url.Parse(urls)
	if err != nil {
		return "", err
	}

	// Add key and other parameters to the query string
	q := u.Query()
	q.Set("format", "json")
	q.Set("q", req.Query)
	q.Set("addressdetails", "1")
	if req.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	if len(req.CountryCodes) > 0 {
		q.Set("countrycodes", strings.Join(req.CountryCodes, ","))
	}
	if len(req.ViewBox) == 4 {
		q.Set("viewbox", fmt.Sprintf("%f,%f,%f,%f", req.ViewBox[0], req.ViewBox[1], req.ViewBox[2], req.ViewBox[3]))
	}
	if len(req.ExcludePlaceIds) > 0 {
		q.Set("exclude_place_ids", strings.Join(req.ExcludePlaceIds, ","))
	}
	if req.Bounded != nil {
		if *req.Bounded {
			q.Set("bounded", "1")
		} else {
			q.Set("bounded", "0")
		}
	}
	// TODO(oe): routewidth
	if req.RouteWidth != nil {
		q.Set("routewidth", fmt.Sprintf("%f", *req.RouteWidth))
	}
	if req.OSMType != "" {
		q.Set("osm_type", req.OSMType)
	}
	if req.OSMId != "" {
		q.Set("osm_id", req.OSMId)
	}

	// No key here!
	u.RawQuery = q.Encode()
	return u.String(), nil
}

type NominatimSearchRequest struct {
	Query           string
	Limit           int
	CountryCodes    []string
	ViewBox         []float64
	ExcludePlaceIds []string
	Bounded         *bool
	RouteWidth      *float64
	OSMType         string
	OSMId           string
}

type NominatimSearchResponse struct {
	Results []*NominatimSearchResult
}

type NominatimSearchResult struct {
	Address *struct {
		City          string `json:"city,omitempty"`
		CityDistrict  string `json:"city_district,omitempty"`
		Continent     string `json:"continent,omitempty"`
		Country       string `json:"country,omitempty"`
		CountryCode   string `json:"country_code,omitempty"`
		County        string `json:"county,omitempty"`
		Hamlet        string `json:"hamlet,omitempty"`
		HouseNumber   string `json:"house_number,omitempty"`
		Pedestrian    string `json:"pedestrian,omitempty"`
		Neighbourhood string `json:"neighbourhood,omitempty"`
		PostCode      string `json:"postcode,omitempty"`
		Road          string `json:"road,omitempty"`
		State         string `json:"state,omitempty"`
		StateDistrict string `json:"state_district,omitempty"`
		Suburb        string `json:"suburb,omitempty"`
	} `json:"address,omitempty"`
	BoundingBox []string `json:"boundingbox,omitempty"`
	Class       string   `json:"class,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Importance  float64  `json:"importance,omitempty"`
	Latitude    float64  `json:"lat,string,omitempty"`
	Longitude   float64  `json:"lon,string,omitempty"`
	OSMId       string   `json:"osm_id,omitempty"`
	OSMType     string   `json:"osm_type,omitempty"`
	PlaceId     string   `json:"place_id,omitempty"`
	Type        string   `json:"type,omitempty"`
	License     string   `json:"licence,omitempty"` // typo in API?
}
