package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olivere/mapquest"
)

var _ = os.DevNull

func main() {
	key := os.Getenv("MAPQUEST_KEY")
	if key == "" {
		log.Fatal("MAPQUEST_KEY environment variable not specified")
	}
	if len(os.Args) != 2 {
		return
	}

	client := mapquest.NewClient(key)
	//logger := log.New(os.Stdout, "", 0)
	//client.SetLogger(logger)

	req := &mapquest.NominatimSearchRequest{
		Query: os.Args[1],
	}
	res, err := client.Nominatim().Search(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, result := range res.Results {
		if result.Address != nil {
			if result.Address.City != "" {
				fmt.Printf("%20s: %s\n", "City", result.Address.City)
			}
			if result.Address.CityDistrict != "" {
				fmt.Printf("%20s: %s\n", "City district", result.Address.CityDistrict)
			}
			if result.Address.Continent != "" {
				fmt.Printf("%20s: %s\n", "Continent", result.Address.Continent)
			}
			if result.Address.Country != "" {
				fmt.Printf("%20s: %s\n", "Country", result.Address.Country)
			}
			if result.Address.CountryCode != "" {
				fmt.Printf("%20s: %s\n", "Country code", result.Address.CountryCode)
			}
			if result.Address.County != "" {
				fmt.Printf("%20s: %s\n", "County", result.Address.County)
			}
			if result.Address.PostCode != "" {
				fmt.Printf("%20s: %s\n", "Postcode", result.Address.PostCode)
			}
			if result.Address.Road != "" {
				fmt.Printf("%20s: %s\n", "Road", result.Address.Road)
			}
			if result.Address.State != "" {
				fmt.Printf("%20s: %s\n", "State", result.Address.State)
			}
			if result.Address.StateDistrict != "" {
				fmt.Printf("%20s: %s\n", "State district", result.Address.StateDistrict)
			}
			if result.Address.Suburb != "" {
				fmt.Printf("%20s: %s\n", "Suburb", result.Address.Suburb)
			}
			fmt.Printf("%20s: [%.6f,%.6f]\n", "Lat/Lon", result.Latitude, result.Longitude)
			if result.Class != "" {
				fmt.Printf("%20s: %s\n", "Class", result.Class)
			}
			if result.Type != "" {
				fmt.Printf("%20s: %s\n", "Type", result.Type)
			}
			fmt.Printf("%20s: %f\n", "Importance", result.Importance)
			if result.OSMId != "" {
				fmt.Printf("%20s: %s\n", "OSM ID", result.OSMId)
			}
			if result.OSMType != "" {
				fmt.Printf("%20s: %s\n", "OSM Type", result.OSMType)
			}
			if result.PlaceId != "" {
				fmt.Printf("%20s: %s\n", "Place ID", result.PlaceId)
			}
			if result.DisplayName != "" {
				fmt.Printf("%20s: %s\n", "Display name", result.DisplayName)
			}
			fmt.Println()
		} else {
			fmt.Printf("%v\n", result)
		}
	}
}
