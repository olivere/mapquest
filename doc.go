/*
Package mapquest enables access to the Open MapQuest APIs.
For further details, see http://open.mapquestapi.com/.

To get started, you need to create a client:

	client := mapquest.NewClient("<your-app-key>")

    // To use HTTPS, use:
    client.SetHTTPS(true)

    // To use your own http.Client:
    client.SetHTTPClient(myClient)

    // To log request and response, set a logger:
    logger := log.New(os.Stderr, "", 0)
    client.SetLogger(logger)

Now that you have a client, you can use the APIs.

Here's an example of how to use the MapQuest static map API:

    req := &mapquest.StaticMapRequest{
      Center: &mapquest.GeoPoint{
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
      panic(err)
	}

To use the Geocoding API, issue a request like this:

    req := &mapquest.GeocodingAddressRequest{
      Location: &mapquest.GeocodingLocation{
        Street:     "1090 N Charlotte St",
        City:       "Lancaster",
        State:      "PA",
        PostalCode: "17603",
      },
    }
    res, err := client.Geocoding().Address(req)
    if err != nil {
      panic(err)
    }

The Nominatim API can be used as follows:

    req := &mapquest.NominatimSearchRequest{
      Query: "Unter den Linden 117, Berlin, DE",
      Limit: 1,
    }
    res, err := client.Nominatim().Search(req)
    if err != nil {
      panic(err)
    }
*/
package mapquest
