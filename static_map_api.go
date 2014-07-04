package mapquest

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/url"
	"strings"
)

var _ = log.Print

const (
	// StaticMapPathPrefix is the default path prefix for the Static Map API.
	StaticMapPathPrefix = "/staticmap/v4"
)

// StaticMapAPI enables users to request static map images via the
// MapQuest API. See http://open.mapquestapi.com/staticmap/ for details.
type StaticMapAPI struct {
	c *Client
}

// Get returns an image of static map by querying MapQuest.
func (api *StaticMapAPI) Get(req *StaticMapRequest) (image.Image, error) {
	u, err := api.buildURL(req)
	if err != nil {
		return nil, err
	}

	res, err := api.c.getResponse(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// buildURL returns the complete URL for the request,
// including the key to query the MapQuest API.
func (api *StaticMapAPI) buildURL(req *StaticMapRequest) (string, error) {
	urls := fmt.Sprintf("%s%s/getmap", api.c.BaseURL(), StaticMapPathPrefix)
	u, err := url.Parse(urls)
	if err != nil {
		return "", err
	}

	// Add key and other parameters to the query string
	q := u.Query()
	if req.Center != nil {
		pt := *req.Center
		q.Set("center", fmt.Sprintf("%f,%f", pt.Latitude, pt.Longitude))
	}
	if req.Bestfit != nil {
		box := *req.Bestfit
		q.Set("bestfit", fmt.Sprintf("%f,%f,%f,%f",
			box.A.Latitude, box.A.Longitude,
			box.B.Latitude, box.B.Longitude))
	}
	if req.Margin > 0 {
		q.Set("margin", fmt.Sprintf("%d", req.Margin))
	}
	q.Set("size", fmt.Sprintf("%d,%d", req.Width, req.Height))
	if req.Zoom > 0 {
		q.Set("zoom", fmt.Sprintf("%d", req.Zoom))
	}
	if req.Scale > 0 {
		q.Set("scale", fmt.Sprintf("%d", req.Scale))
	}
	if req.Type != "" {
		q.Set("type", req.Type)
	}
	if req.Format != "" {
		q.Set("imagetype", req.Format)
	}
	if len(req.PointsOfInterest) > 0 {
		parts := make([]string, len(req.PointsOfInterest))
		for i, poi := range req.PointsOfInterest {
			if poi.OffsetX > 0 || poi.OffsetY > 0 {
				parts[i] = fmt.Sprintf("%s,%f,%f,%d,%d",
					poi.Label,
					poi.Latitude,
					poi.Longitude,
					poi.OffsetX,
					poi.OffsetY)
			} else {
				parts[i] = fmt.Sprintf("%s,%f,%f",
					poi.Label,
					poi.Latitude,
					poi.Longitude)
			}
		}
		q.Set("pois", strings.Join(parts, "|"))
	}

	// Key has to be handled specifically here, because
	// the MapQuest API seems to not like the key URL-encoded
	u.RawQuery = fmt.Sprintf("key=%s&%s", api.c.key, q.Encode())
	return u.String(), nil
}

type StaticMapRequest struct {
	// Center defines the center point for the map image.
	Center *GeoPoint

	// Bestfit defines a bounding box to be used to specify
	// the area of the map to be shown. This can be used
	// instead of Center.
	Bestfit *GeoBox

	// Margin can adjust the zoom level accordingly when
	// you are out of bounds of the map. Use this in
	// combination with Bestfit.
	Margin int

	// Width of the map. The width must not be greater than 3840.
	Width int

	// Height of the map. The height must not be greater than 3840.
	Height int

	// Zoom specifies the zoom level, which is in the
	// range of 1 (world view) to 18 (most details).
	// See http://open.mapquestapi.com/staticmap/zoomToScale.html
	// for details and scale.
	Zoom int

	// Scale specifies the display scale for the map image,
	// in the range of 1-18 (see Zoom).
	// You must specify a scale when you use the Center property
	// and you do not specify a zoom level.
	Scale int

	// Type specifies the map mode. It can be "map", "sat", or "hyb".
	// The default is "map".
	Type string

	// Format specifies the image format. Valid values are
	// "jpeg", "jpg", "gif", and "png". The default is "jpg".
	Format string

	// PointsOfInterest enlists the various points of interest to be
	// displayed on the map.
	PointsOfInterest []*PointOfInterest
}

// PointOfInterest defines an interesting point to be displayed on a map.
type PointOfInterest struct {
	// Label is the name of the icon to display.
	// See http://open.mapquestapi.com/staticmap/icons.html for
	// the list of valid icons.
	Label string

	// Latitude of the point of interest.
	Latitude float64

	// Longitude of the point of interest.
	Longitude float64

	// OffsetX is the offset on the x axis. It is optional.
	OffsetX int

	// OffsetY is the offset on the y axis. It is optional.
	OffsetY int
}
