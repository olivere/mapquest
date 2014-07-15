package mapquest

import (
	_ "bytes"
	_ "image/png"
	_ "io/ioutil"
	"testing"
)

func TestStaticMapBuildURLs(t *testing.T) {
	testKey, err := readKey(t)
	if err != nil {
		t.Fail()
		return
	}

	tests := []struct {
		Request *StaticMapRequest
		URL     string
	}{
		{
			Request: &StaticMapRequest{
				Center: &GeoPoint{
					Longitude: 11.54165,
					Latitude:  48.151313,
				},
				Zoom:   9,
				Width:  500,
				Height: 300,
				Format: "png",
			},
			URL: "http://open.mapquestapi.com/staticmap/v4/getmap?center=48.151313,11.541650&size=500,300&zoom=9&imagetype=png&key=" + testKey,
		},
	}

	client := NewClient(testKey)
	for _, test := range tests {
		got, err := client.StaticMap().buildURL(test.Request)
		if err != nil {
			t.Fatalf("expeced no error, got: %v", err)
		}
		if got != test.URL {
			t.Errorf("expected %q, got: %q", test.URL, got)
		}
	}
}

func TestStaticMapGet(t *testing.T) {
	key, err := readKey(t)
	if err != nil {
		t.Fail()
		return
	}

	client := NewClient(key)
	req := &StaticMapRequest{
		Center: &GeoPoint{
			Longitude: 11.54165,
			Latitude:  48.151313,
		},
		Zoom:   9,
		Width:  500,
		Height: 300,
		Format: "png",
	}
	img, err := client.StaticMap().Get(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if img == nil {
		t.Fatalf("expected image data, got: %v", img)
	}
	bounds := img.Bounds()
	if bounds.Empty() {
		t.Fatal("expected a non-empty image")
	}
}
