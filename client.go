package mapquest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	DefaultHost = "open.mapquestapi.com"
	UserAgent   = "MapQuest Open Data API Google Go Client v0.1"
)

// Client is the entry point to all services of the MapQuest Open Data API.
// See http://developer.mapquest.com/web/products/open for details about
// what you can do with the MapQuest API.
type Client struct {
	httpClient *http.Client
	https      bool
	key        string
	log        *log.Logger
}

// NewClient creates a new client for accessing the MapQuest API. You need
// to specify your AppKey here.
func NewClient(key string) *Client {
	return &Client{
		https:      false,
		key:        key,
		httpClient: http.DefaultClient,
	}
}

// SetHTTPClient allows the caller to specify a special http.Client for
// invoking the MapQuest API. If you do not specify a http.Client, the
// http.DefaultClient from net/http is used.
func (c *Client) SetHTTPClient(client *http.Client) {
	if client == nil {
		client = http.DefaultClient
	}
	c.httpClient = client
}

// HTTPClient returns the registered http.Client. Notice that nil can
// be returned here.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// SetHTTPS tells the client whether to use HTTPS or HTTP.
func (c *Client) SetHTTPS(https bool) {
	c.https = https
}

// HTTPS returns true if the client is configured to use HTTPS,
// false otherwise.
func (c *Client) HTTPS() bool {
	return c.https
}

// SetLogger sets the logger to use when e.g. debugging requests.
// Set to nil to disable logging (the default).
func (c *Client) SetLogger(logger *log.Logger) {
	c.log = logger
}

// BaseURL returns the base URL to access the MapQuest API.
// Example: https://open.mapquestapi.com (without the trailing slash).
func (c *Client) BaseURL() string {
	if c.https {
		return fmt.Sprintf("https://%s", DefaultHost)
	}
	return fmt.Sprintf("http://%s", DefaultHost)
}

// StaticMap gives access to the MapQuest static map API
// described here: http://open.mapquestapi.com/staticmap/
func (c *Client) StaticMap() *StaticMapAPI {
	return &StaticMapAPI{c: c}
}

// Geocoding gives access to the MapQuest geocoding API
// described here: http://open.mapquestapi.com/geocoding/
func (c *Client) Geocoding() *GeocodingAPI {
	return &GeocodingAPI{c: c}
}

// Nominatim is a gateway to the Nominatim API provided by MapQuest.
// See http://open.mapquestapi.com/nominatim/ for details.
func (c *Client) Nominatim() *NominatimAPI {
	return &NominatimAPI{c: c}
}

// -- Helper functions --

func (c *Client) logRequest(r *http.Request) error {
	if c.log != nil {
		out, err := httputil.DumpRequestOut(r, true)
		if err != nil {
			return err
		}
		c.log.Printf("Request: %s", string(out))
	}
	return nil
}

func (c *Client) logResponse(r *http.Response, body bool) error {
	if c.log != nil {
		out, err := httputil.DumpResponse(r, body)
		if err != nil {
			return err
		}
		c.log.Printf("Response: %s", string(out))
	}
	return nil
}

// getResponse returns the HTTP response to the caller.
// Warning: The caller is responsible for closing the
// Body via e.g. `defer res.Body.Close()`.
func (c *Client) getResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)

	if err := c.logRequest(req); err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := c.logResponse(res, false); err != nil {
		return nil, err
	}

	return res, nil
}

// getJSON performs a HTTP GET request to the specified URL,
// decodes the result into v and returns nil.
func (c *Client) getJSON(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", UserAgent)

	if err := c.logRequest(req); err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := c.logResponse(res, true); err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
