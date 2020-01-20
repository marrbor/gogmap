package gogmap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiURL = "https://maps.googleapis.com/maps/api/geocode/json"
)

//// import from googlemaps.github.io/maps
type (
	response struct {
		Results      []GeocodingResult `json:"results"`
		Status       string            `json:"status"`
		ErrorMessage string            `json:"error_message"`
	}

	// GeocodingResult is a single geocoded address
	GeocodingResult struct {
		AddressComponents []AddressComponent `json:"address_components"`
		FormattedAddress  string             `json:"formatted_address"`
		Geometry          AddressGeometry    `json:"geometry"`
		Types             []string           `json:"types"`
		PlaceID           string             `json:"place_id"`
	}

	// AddressComponent is a part of an address
	AddressComponent struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}

	// AddressGeometry is the location of a an address
	AddressGeometry struct {
		Location     LatLng       `json:"location"`
		LocationType string       `json:"location_type"`
		Bounds       LatLngBounds `json:"bounds"`
		Viewport     LatLngBounds `json:"viewport"`
		Types        []string     `json:"types"`
	}

	// LatLngBounds represents a bounded square area on the Earth.
	LatLngBounds struct {
		NorthEast LatLng `json:"northeast"`
		SouthWest LatLng `json:"southwest"`
	}

	// LatLng represents a location on the Earth.
	LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
)

// GeoCoding
func GeoCoding(dst, apiKey string) (*LatLng, error) {
	v := url.Values{}
	v.Add("address", dst)
	v.Add("key", apiKey)
	uri := apiURL + "?" + v.Encode()
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("google map error:%d", res.StatusCode)
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if len(resp.Results) < 1 {
		return nil, fmt.Errorf("no result found")
	}

	ll := resp.Results[0].Geometry.Location
	return &LatLng{Lat: ll.Lat, Lng: ll.Lng}, nil
}
