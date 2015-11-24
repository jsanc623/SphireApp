package geofence

import "encoding/json"

const (
	degree_per_lat_mile = 0.0144697
	degree_per_lng_mile = 0.0144812
)

type (
	Bounds struct {
		Lat BoundsLatLng `json:"lat"`
		Lng BoundsLatLng `json:"lng"`
		Mid LatLng       `json:"mid"`
	}

	LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	BoundsLatLng struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	}
)

func BoundingBox(latitude float64, longitude float64, miles float64) string {
	var lat_bnd float64 = degree_per_lat_mile * miles
	var lng_bnd float64 = degree_per_lng_mile * miles

	output, _ := json.Marshal(&Bounds{
		Lat: BoundsLatLng{Min: latitude - lat_bnd, Max: latitude + lat_bnd},
		Lng: BoundsLatLng{Min: longitude - lng_bnd, Max: longitude + lng_bnd},
		Mid: LatLng{Lat: latitude, Lng: longitude},
	})

	return string(output)
}
