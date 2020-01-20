package gogmap_test

import (
	"github.com/marrbor/gogmap"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type (
	GeoCodingTestData struct {
		Address string
		LatLng  gogmap.LatLng
	}
)

var testdata = []GeoCodingTestData{
	{Address: "新宿駅", LatLng: gogmap.LatLng{Lat: 35.690921, Lng: 139.70025799999996}},
}

func TestGeoCoding(t *testing.T) {
	apikey := os.Getenv("APIKEY")
	for _, d := range testdata {
		ll, err := gogmap.GeoCoding(d.Address, apikey)
		assert.NoError(t, err)
		assert.EqualValues(t, d.LatLng.Lat, ll.Lat)
		assert.EqualValues(t, d.LatLng.Lng, ll.Lng)
	}
}
