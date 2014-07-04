package mapquest

type GeoPoint struct {
	Latitude  float64
	Longitude float64
}

type GeoBox struct {
	A GeoPoint
	B GeoPoint
}
