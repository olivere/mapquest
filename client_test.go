package mapquest

import (
	"io/ioutil"
	"strings"
	"testing"
)

func readKey(t *testing.T) (string, error) {
	key, err := ioutil.ReadFile("ACCESS_KEY")
	if err != nil {
		t.Fatalf("mapquest: please put a file called ACCESS_KEY in the " +
			"directory of the package and paste your MapQuest key into it")
		return "", err
	}
	return strings.TrimSpace(string(key)), nil
}

func TestDefaults(t *testing.T) {
	expected := "open.mapquestapi.com"
	if DefaultHost != expected {
		t.Errorf("expected default host of %q, got: %q", expected, DefaultHost)
	}
}

func TestHTTPClient(t *testing.T) {
	c := NewClient("my-key")
	hc := c.HTTPClient()
	if hc == nil {
		t.Fatalf("expected HTTPClient() to never return nil, got: %v", hc)
	}
}

func TestBaseURL(t *testing.T) {
	c := NewClient("my-key")
	if c.HTTPS() {
		t.Error("expected HTTP scheme by default, got: HTTPS")
	}
	expected := "http://open.mapquestapi.com"
	got := c.BaseURL()
	if got != expected {
		t.Errorf("expeced base URL of %q, got: %q", expected, got)
	}

	c = NewClient("my-key")
	c.SetHTTPS(true)
	if !c.HTTPS() {
		t.Error("expected HTTPS scheme, got: HTTP")
	}
	expected = "https://open.mapquestapi.com"
	got = c.BaseURL()
	if got != expected {
		t.Errorf("expeced base URL of %q, got: %q", expected, got)
	}
}
